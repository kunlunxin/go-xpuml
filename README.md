# Go Bindings for the XPU Management Library (XPUML)

## Quick Start

导入 xpuml package 后调用`xpuml.init()`即可，可参考`examples/devices/main.go`，该示例实现了调用xpuml接口查看昆仑设备内存信息的功能(注意运行时需要链接`libxpuml.so`)。

输出类似：

```console
$ go run examples/devices/main.go
Memory info of device at index 0: {PageSizeGlobalMemory:131072 TotalGlobalMemory:8455716864 FreeGlobalMemory:8455716864 UsedGlobalMemory:0 PageSizeL3Memory:4096 TotalL3Memory:16777216 FreeL3Memory:8455716864 UsedL3Memory:0 Reserved:[0 0 0 0 0 0 0 0]}
Memory info of device at index 1: {PageSizeGlobalMemory:131072 TotalGlobalMemory:8455716864 FreeGlobalMemory:8455716864 UsedGlobalMemory:0 PageSizeL3Memory:4096 TotalL3Memory:16777216 FreeL3Memory:8455716864 UsedL3Memory:0 Reserved:[0 0 0 0 0 0 0 0]}
Memory info of device at index 2: {PageSizeGlobalMemory:131072 TotalGlobalMemory:8455716864 FreeGlobalMemory:8455716864 UsedGlobalMemory:0 PageSizeL3Memory:4096 TotalL3Memory:16777216 FreeL3Memory:8455716864 UsedL3Memory:0 Reserved:[0 0 0 0 0 0 0 0]}
Memory info of device at index 3: {PageSizeGlobalMemory:131072 TotalGlobalMemory:8455716864 FreeGlobalMemory:8455716864 UsedGlobalMemory:0 PageSizeL3Memory:4096 TotalL3Memory:16777216 FreeL3Memory:8455716864 UsedL3Memory:0 Reserved:[0 0 0 0 0 0 0 0]}
```

## Code Structure

- `pkg`目录为最终的package产出。
- `pkg/dl`负责在运行时链接libxpuml.so。
- `gen/xpuml/xpuml.yml`为 c-for-go 的配置文件，c-for-go 基于 xpuml.h 自动生成了一些函数接口。
- `gen/xpuml/`下的文件，是基于这些自动生成的函数接口，构造的更符合最终用户使用习惯的调用接口。
- `gen/xpuml/init.go`负责初始化和接口的版本管理。

## How To Contribute

1. 安装 c-for-go

```console
$ go install github.com/xlab/c-for-go@8eeee8c3
```

2. 安装 pre-commit hook

```console
$ pip install pre-commit
$ pre-commit install .pre-commit-config.yaml
```

3. 编译

```console
$ make
```

注意编译后在`gen/xpuml`目录下会产生一些自动生成的文件，这些是 c-for-go 根据 xpuml.h 产生的 func wrapper，这些自动生成文件以外的 go 文件，是人工地基于这些文件，编写的更符合使用者习惯的函数接口(可参考`gen/xpuml/device.go`)。

```console
Untracked files:
  (use "git add <file>..." to include in what will be committed)

        gen/xpuml/cgo_helpers.go
        gen/xpuml/cgo_helpers.h
        gen/xpuml/const.go
        gen/xpuml/doc.go
        gen/xpuml/types_gen.go
        gen/xpuml/xpuml.go
```

>注意编译阶段不需要链接 libxpuml.so 库, 但后续运行阶段需要。

4. 测试

```console
$ make test
=== RUN   TestInit
    xpuml_test.go:26: Init: 0
    xpuml_test.go:33: Shutdown: 0
--- PASS: TestInit (0.00s)
=== RUN   TestSystem
    xpuml_test.go:45: SystemGetDriverVersion: 0
    xpuml_test.go:46:   version: 4.0.third_unknown.1 commit: commit_unknown_git_not_installed
    xpuml_test.go:53: SystemGetXPUMLVersion: 0
    xpuml_test.go:54:   version: 4.0.7.1 commit: 34377fd717001d7b15db36d4285c98abae372cee
--- PASS: TestSystem (0.00s)
=== RUN   TestDevice
    xpuml_test.go:66: DeviceGetCount: 0
    xpuml_test.go:67:   count: 4
    xpuml_test.go:79: DeviceGetHandleByIndex index: 0, 0
    xpuml_test.go:86: device index: 0, attr: {ArchId:1 BoardId:2 ModelId:0 ClusterCount:4294967295 SdnnCount:4294967295 DmaEngineCount:4294967295 DecoderCount:4294967295 EncoderCount:4294967295 ImgProcCo$nt:4294967295 MaxVFCount:4294967295 GlobalMemorySizeMB:8064 L3MemorySizeMB:16 Reserved:[0 0 0 0 0 0 0 0 0]}
    xpuml_test.go:93: device index: 0, meminfo: {PageSizeGlobalMemory:131072 TotalGlobalMemory:8455716864 FreeGlobalMemory:8455716864 UsedGlobalMemory:0 PageSizeL3Memory:4096 TotalL3Memory:16777216 FreeL$Memory:8455716864 UsedL3Memory:0 Reserved:[0 0 0 0 0 0 0 0]}
    xpuml_test.go:100: device index: 0, boardId: 2
    xpuml_test.go:107: Found 1 processes on device 0
    xpuml_test.go:109:  [ 0] ProcessInfo: {Pid:21100 UsedGlobalMemory:0 UsedL3Memory:0 Reserved:[0 0 0 0 0 0 0 0 0 0 0 0 0]}
    xpuml_test.go:79: DeviceGetHandleByIndex index: 1, 0
    xpuml_test.go:86: device index: 1, attr: {ArchId:1 BoardId:2 ModelId:0 ClusterCount:4294967295 SdnnCount:4294967295 DmaEngineCount:4294967295 DecoderCount:4294967295 EncoderCount:4294967295 ImgProcCo$nt:4294967295 MaxVFCount:4294967295 GlobalMemorySizeMB:8064 L3MemorySizeMB:16 Reserved:[0 0 0 0 0 0 0 0 0]}
    xpuml_test.go:93: device index: 1, meminfo: {PageSizeGlobalMemory:131072 TotalGlobalMemory:8455716864 FreeGlobalMemory:8455716864 UsedGlobalMemory:0 PageSizeL3Memory:4096 TotalL3Memory:16777216 FreeL$Memory:8455716864 UsedL3Memory:0 Reserved:[0 0 0 0 0 0 0 0]}
    xpuml_test.go:100: device index: 1, boardId: 2
    xpuml_test.go:107: Found 1 processes on device 1
    xpuml_test.go:109:  [ 0] ProcessInfo: {Pid:21100 UsedGlobalMemory:0 UsedL3Memory:0 Reserved:[0 0 0 0 0 0 0 0 0 0 0 0 0]}
    xpuml_test.go:79: DeviceGetHandleByIndex index: 2, 0
    xpuml_test.go:86: device index: 2, attr: {ArchId:1 BoardId:2 ModelId:0 ClusterCount:4294967295 SdnnCount:4294967295 DmaEngineCount:4294967295 DecoderCount:4294967295 EncoderCount:4294967295 ImgProcCou
nt:4294967295 MaxVFCount:4294967295 GlobalMemorySizeMB:8064 L3MemorySizeMB:16 Reserved:[0 0 0 0 0 0 0 0 0]}
    xpuml_test.go:93: device index: 2, meminfo: {PageSizeGlobalMemory:131072 TotalGlobalMemory:8455716864 FreeGlobalMemory:8455716864 UsedGlobalMemory:0 PageSizeL3Memory:4096 TotalL3Memory:16777216 FreeL3
Memory:8455716864 UsedL3Memory:0 Reserved:[0 0 0 0 0 0 0 0]}
    xpuml_test.go:100: device index: 2, boardId: 2
    xpuml_test.go:107: Found 1 processes on device 2
    xpuml_test.go:109:  [ 0] ProcessInfo: {Pid:21100 UsedGlobalMemory:0 UsedL3Memory:0 Reserved:[0 0 0 0 0 0 0 0 0 0 0 0 0]}
    xpuml_test.go:79: DeviceGetHandleByIndex index: 3, 0
    xpuml_test.go:86: device index: 3, attr: {ArchId:1 BoardId:2 ModelId:0 ClusterCount:4294967295 SdnnCount:4294967295 DmaEngineCount:4294967295 DecoderCount:4294967295 EncoderCount:4294967295 ImgProcCou
nt:4294967295 MaxVFCount:4294967295 GlobalMemorySizeMB:8064 L3MemorySizeMB:16 Reserved:[0 0 0 0 0 0 0 0 0]}
    xpuml_test.go:93: device index: 3, meminfo: {PageSizeGlobalMemory:131072 TotalGlobalMemory:8455716864 FreeGlobalMemory:8455716864 UsedGlobalMemory:0 PageSizeL3Memory:4096 TotalL3Memory:16777216 FreeL3
Memory:8455716864 UsedL3Memory:0 Reserved:[0 0 0 0 0 0 0 0]}
    xpuml_test.go:100: device index: 3, boardId: 2
    xpuml_test.go:107: Found 1 processes on device 3
    xpuml_test.go:109:  [ 0] ProcessInfo: {Pid:21100 UsedGlobalMemory:0 UsedL3Memory:0 Reserved:[0 0 0 0 0 0 0 0 0 0 0 0 0]}
--- PASS: TestDevice (0.00s)
PASS
ok      kunlun-xre/go-xpuml/pkg/xpuml   (cached)
```
