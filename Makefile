# Copyright (c) 2022, KUNLUNXIN CORPORATION.  All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

DOCKER ?= docker
GOLANG_VERSION ?= 1.15
C_FOR_GO_TAG ?= 8eeee8c3b71f9c3c90c4a73db54ed08b0bba971d
BUILDIMAGE ?= kunlunxin/go-xpuml-devel:$(GOLANG_VERSION)-$(C_FOR_GO_TAG)

PWD := $(shell pwd)
GEN_DIR := $(PWD)/gen
PKG_DIR := $(PWD)/pkg
GEN_BINDINGS_DIR := $(GEN_DIR)/xpuml
PKG_BINDINGS_DIR := $(PKG_DIR)/xpuml

SOURCES = $(shell find $(GEN_BINDINGS_DIR) -type f)

EXAMPLES := $(patsubst ./examples/%/,%,$(sort $(dir $(wildcard ./examples/*/))))
EXAMPLE_TARGETS := $(patsubst %,example-%, $(EXAMPLES))

ifeq ($(shell uname),Darwin)
	SED := $(DOCKER) run -it --rm -v "$(PWD):$(PWD)" -w "$(PWD)" alpine:latest sed
else
	SED := sed
endif

TARGETS := all test clean bindings test-bindings clean-bindings patch-xpuml-h examples $(EXAMPLE_TARGETS)
DOCKER_TARGETS := $(patsubst %, docker-%, $(TARGETS))
.PHONY: $(TARGETS) $(DOCKER_TARGETS)

.DEFAULT_GOAL = all

all: bindings
test: test-bindings
clean: clean-bindings

GOOS := linux

examples: $(EXAMPLE_TARGETS)
$(EXAMPLE_TARGETS): example-%: bindings
	GOOS=$(GOOS) go build ./examples/$(*)

$(PKG_BINDINGS_DIR):
	mkdir -p $(@)

patch-xpuml-h: $(PKG_BINDINGS_DIR)/xpuml.h
$(PKG_BINDINGS_DIR)/xpuml.h: $(GEN_BINDINGS_DIR)/xpuml.h | $(PKG_BINDINGS_DIR)
	$(SED) -E 's#(typedef\s+struct)\s+(xpuml.*_st\*)\s+(xpuml.*_t);#\1\n{\n    struct \2 handle;\n} \3;#g' $(<) > $(@)

bindings: .create-bindings .strip-autogen-comment .strip-xpuml-h-linenumber

.create-bindings: $(PKG_BINDINGS_DIR)/xpuml.h $(SOURCES) | $(PKG_BINDINGS_DIR)
	cp $(GEN_BINDINGS_DIR)/*.go $(PKG_BINDINGS_DIR)
	c-for-go -out $(PKG_DIR) $(GEN_BINDINGS_DIR)/xpuml.yml
	cd $(PKG_BINDINGS_DIR); \
		go tool cgo -godefs types.go > types_gen.go; \
		go fmt types_gen.go; \
	cd -> /dev/null
	rm -rf $(PKG_BINDINGS_DIR)/types.go $(PKG_BINDINGS_DIR)/_obj
# add this cp step to pass pre-commit golangci check
	cd $(PKG_BINDINGS_DIR); \
		cp `ls | grep -v xpuml.h|xargs` $(GEN_BINDINGS_DIR)

.strip-autogen-comment: SED_SEARCH_STRING := // WARNING: This file has automatically been generated on
.strip-autogen-comment: SED_REPLACE_STRING := // WARNING: THIS FILE WAS AUTOMATICALLY GENERATED.
.strip-autogen-comment: | .create-bindings
	grep -l -R "$(SED_SEARCH_STRING)" pkg \
		| xargs $(SED) -i -E 's#$(SED_SEARCH_STRING).*$$#$(SED_REPLACE_STRING)#g'

.strip-xpuml-h-linenumber: SED_SEARCH_STRING := // (.*) xpuml/xpuml.h:[0-9]+
.strip-xpuml-h-linenumber: SED_REPLACE_STRING := // \1 xpuml/xpuml.h
.strip-xpuml-h-linenumber: | .create-bindings
	grep -l -RE "$(SED_SEARCH_STRING)" pkg \
		| xargs $(SED) -i -E 's#$(SED_SEARCH_STRING)$$#$(SED_REPLACE_STRING)#g'

test-bindings: bindings
	cd $(PKG_BINDINGS_DIR); \
		go test -v .; \
	cd -> /dev/null

clean-bindings:
	rm -rf $(PKG_BINDINGS_DIR)

# Generate an image for containerized builds
# Note: This image is local only
.build-image: Dockerfile
	$(DOCKER) build \
		--progress=plain \
		--build-arg GOLANG_VERSION="$(GOLANG_VERSION)" \
		--build-arg C_FOR_GO_TAG="$(C_FOR_GO_TAG)" \
		--tag $(BUILDIMAGE) \
		.

# A target for executing make targets in a docker container
$(DOCKER_TARGETS): docker-%: .build-image
	@echo "Running 'make $(*)' in docker container $(BUILDIMAGE)"
	$(DOCKER) run \
		--rm \
		-e GOCACHE=/tmp/.cache \
		-v $(PWD):$(PWD) \
		-w $(PWD) \
		--user $$(id -u):$$(id -g) \
		$(BUILDIMAGE) \
			make $(*)

# A make target to set up an interactive docker environment as the current user.
# This is useful for debugging issues in the make process in the container.
.docker-shell: .build-image
	$(DOCKER) run \
		-ti \
		--rm \
		-e GOCACHE=/tmp/.cache \
		-v $(PWD):$(PWD) \
		-w $(PWD) \
		--user $$(id -u):$$(id -g) \
		$(BUILDIMAGE) \
			bash

# A make target to set up an interactive docker environment as root.
# This is useful for debugging issues in dependencies or the container build process
.docker-root-shell: .build-image
	$(DOCKER) run \
		-ti \
		--rm \
		-e GOCACHE=/tmp/.cache \
		-v $(PWD):$(PWD) \
		-w $(PWD) \
		$(BUILDIMAGE) \
			bash

# Run markdownlint (https://github.com/markdownlint/markdownlint) for README.md
# Note: Tabs are preferred for Golang code blocks
markdownlint: MDL := $(DOCKER) run --rm -v "$(PWD):$(PWD)" -w "$(PWD)" markdownlint/markdownlint:latest
markdownlint:
	@$(MDL) --rules=~no-hard-tabs,~line-length README.md
