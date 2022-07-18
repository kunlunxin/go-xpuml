/*
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
*/

package main

import (
	"fmt"
	"log"

	"github.com/kunlunxin/go-xpuml/pkg/xpuml"
)

func main() {
	ret := xpuml.Init()
	if ret != xpuml.SUCCESS {
		log.Fatalf("Unable to initialize XPUML: %v", xpuml.ErrorString(ret))
	}
	defer func() {
		ret := xpuml.Shutdown()
		if ret != xpuml.SUCCESS {
			log.Fatalf("Unable to shutdown XPUML: %v", xpuml.ErrorString(ret))
		}
	}()

	count, ret := xpuml.DeviceGetCount()
	if ret != xpuml.SUCCESS {
		log.Fatalf("Unable to get device count: %v", xpuml.ErrorString(ret))
	}

	for i := 0; i < count; i++ {
		device, ret := xpuml.DeviceGetHandleByIndex(i)
		if ret != xpuml.SUCCESS {
			log.Fatalf("Unable to get device at index %d: %v", i, xpuml.ErrorString(ret))
		}

		serial, ret := device.GetSerial()
		if ret != xpuml.SUCCESS {
			log.Fatalf("Unable to get serial of device at index %d: %v", i, xpuml.ErrorString(ret))
		}

		fmt.Printf("Serial of device at index %d: %+v\n", i, serial)

		devAttr, ret := device.GetAttributes()
		if ret != xpuml.SUCCESS {
			log.Fatalf("Unable to get attributes of device at index %d: %v", i, xpuml.ErrorString(ret))
		}

		fmt.Printf("Attributes of device at index %d: %+v\n", i, devAttr)

		mem_info, ret := device.GetMemoryInfo()
		if ret != xpuml.SUCCESS {
			log.Fatalf("Unable to get memory info of device at index %d: %v", i, xpuml.ErrorString(ret))
		}

		fmt.Printf("Memory info of device at index %d: %+v\n", i, mem_info)

		utilization, ret := device.GetUtilizationRates()
		if ret != xpuml.SUCCESS {
			log.Fatalf("Unable to get utilization rates of device at index %d: %v", i, xpuml.ErrorString(ret))
		}

		fmt.Printf("utilization rates of device at index %d: %+v\n", i, utilization)

		clock, ret := device.GetClockInfo(xpuml.CLOCK_XPU)
		if ret != xpuml.SUCCESS {
			log.Fatalf("Unable to get xpu clock of device at index %d: %v", i, xpuml.ErrorString(ret))
		}
		fmt.Printf("xpu clock of device at index %d: %+v\n", i, clock)

		temp, ret := device.GetTemperature(xpuml.TEMPERATURE_XPU)
		if ret != xpuml.SUCCESS {
			log.Fatalf("Unable to get xpu temperature of device at index %d: %v", i, xpuml.ErrorString(ret))
		}
		fmt.Printf("temperature of device at index %d: %+v\n", i, temp)

		power, ret := device.GetPowerUsage()
		if ret != xpuml.SUCCESS {
			log.Fatalf("Unable to get xpu power usage of device at index %d: %v", i, xpuml.ErrorString(ret))
		}
		fmt.Printf("power usage of device at index %d: %+v\n", i, power)

		ProcessInfo, _ := device.GetComputeRunningProcesses()
		fmt.Printf("ProcessInfo of device %d: %+v\n", i, ProcessInfo)
	}
}
