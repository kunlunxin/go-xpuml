// Copyright (c) 2022, KUNLUNXIN CORPORATION.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package xpuml

import (
	"testing"
)

func TestInit(t *testing.T) {
	ret := Init()
	if ret != SUCCESS {
		t.Errorf("Init: %v", ret)
	} else {
		t.Logf("Init: %v", ret)
	}

	ret = Shutdown()
	if ret != SUCCESS {
		t.Errorf("Shutdown: %v", ret)
	} else {
		t.Logf("Shutdown: %v", ret)
	}
}

func TestSystem(t *testing.T) {
	Init()
	defer Shutdown()

	driverVersion, ret := SystemGetDriverVersion()
	if ret != SUCCESS {
		t.Errorf("SystemGetDriverVersion: %v", ret)
	} else {
		t.Logf("SystemGetDriverVersion: %v", ret)
		t.Logf("  version: %v", driverVersion)
	}

	xpumlVersion, ret := SystemGetXPUMLVersion()
	if ret != SUCCESS {
		t.Errorf("SystemGetXPUMLVersion: %v", ret)
	} else {
		t.Logf("SystemGetXPUMLVersion: %v", ret)
		t.Logf("  version: %v", xpumlVersion)
	}
}

func TestDevice(t *testing.T) {
	Init()
	defer Shutdown()

	devCount, ret := DeviceGetCount()
	if ret != SUCCESS {
		t.Errorf("DeviceGetCount: %v", ret)
	} else {
		t.Logf("DeviceGetCount: %v", ret)
		t.Logf("  count: %v", devCount)
	}

	if devCount == 0 {
		t.Skip("Skipping test with no Devices.")
	}

	for i := 0; i < devCount; i++ {
		device, ret := DeviceGetHandleByIndex(i)
		if ret != SUCCESS {
			t.Errorf("DeviceGetHandleByIndex: %v", ret)
		} else {
			t.Logf("DeviceGetHandleByIndex index: %d, %+v", i, ret)
		}

		attr, ret := device.GetAttributes()
		if ret != SUCCESS {
			t.Errorf("GetAttributes: %v", ret)
		} else {
			t.Logf("device index: %d, attr: %+v", i, attr)
		}

		meminfo, ret := device.GetMemoryInfo()
		if ret != SUCCESS {
			t.Errorf("GetMemoryInfo: %v", ret)
		} else {
			t.Logf("device index: %d, meminfo: %+v", i, meminfo)
		}

		utilization, ret := device.GetUtilizationRates()
		if ret != SUCCESS {
			t.Errorf("GetUtilizaionRates: %v", ret)
		} else {
			t.Logf("device index: %d, device utilization: %+v", i, utilization)
		}

		boardId, ret := device.GetBoardId()
		if ret != SUCCESS {
			t.Errorf("GetBoardId: %v", ret)
		} else {
			t.Logf("device index: %d, boardId: %+v", i, boardId)
		}

		pInfo, ret := device.GetComputeRunningProcesses()
		if ret != SUCCESS {
			t.Errorf("GetComputeRunningProcesses: %v", ret)
		} else {
			t.Logf("Found %d processes on device %d\n", len(pInfo), i)
			for pi, processInfo := range pInfo {
				t.Logf("\t[%2d] ProcessInfo: %+v\n", pi, processInfo)
			}
		}
	}
}
