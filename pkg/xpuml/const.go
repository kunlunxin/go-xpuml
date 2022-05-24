// Copyright (c) 2022, KUNLUNXIN CORPORATION. All rights reserved.
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

// WARNING: THIS FILE WAS AUTOMATICALLY GENERATED.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package xpuml

/*
#cgo LDFLAGS: -Wl,--unresolved-symbols=ignore-in-object-files
#cgo CFLAGS: -DXPUML_NO_UNVERSIONED_FUNC_DEFS=1
#include "xpuml.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// NO_UNVERSIONED_FUNC_DEFS as defined in go-xpuml/<predefine>:24
	NO_UNVERSIONED_FUNC_DEFS = 1
	// API_VERSION as defined in xpuml/xpuml.h
	API_VERSION = 1
	// API_VERSION_STR as defined in xpuml/xpuml.h
	API_VERSION_STR = "1"
	// VALUE_NOT_AVAILABLE as defined in xpuml/xpuml.h
	VALUE_NOT_AVAILABLE = -1
	// DEVICE_PCI_BUS_ID_BUFFER_SIZE as defined in xpuml/xpuml.h
	DEVICE_PCI_BUS_ID_BUFFER_SIZE = 32
	// DEVICE_PCI_BUS_ID_FMT as defined in xpuml/xpuml.h
	DEVICE_PCI_BUS_ID_FMT = "%08X:%02X:%02X.%01X"
	// TOPOLOGY_CPU as defined in xpuml/xpuml.h
	TOPOLOGY_CPU = 0
	// FlagDefault as defined in xpuml/xpuml.h
	FlagDefault = 0
	// FlagForce as defined in xpuml/xpuml.h
	FlagForce = 1
	// VXPU_NAME_BUFFER_SIZE as defined in xpuml/xpuml.h
	VXPU_NAME_BUFFER_SIZE = 64
	// BUS_TYPE_UNKNOWN as defined in xpuml/xpuml.h
	BUS_TYPE_UNKNOWN = 0
	// BUS_TYPE_PCI as defined in xpuml/xpuml.h
	BUS_TYPE_PCI = 1
	// BUS_TYPE_PCIE as defined in xpuml/xpuml.h
	BUS_TYPE_PCIE = 2
	// BUS_TYPE_FPCI as defined in xpuml/xpuml.h
	BUS_TYPE_FPCI = 3
	// BUS_TYPE_AGP as defined in xpuml/xpuml.h
	BUS_TYPE_AGP = 4
	// FI_DEV_ECC_CURRENT as defined in xpuml/xpuml.h
	FI_DEV_ECC_CURRENT = 1
	// FI_MAX as defined in xpuml/xpuml.h
	FI_MAX = 161
	// INIT_FLAG_NO_XPUS as defined in xpuml/xpuml.h
	INIT_FLAG_NO_XPUS = 1
	// INIT_FLAG_NO_ATTACH as defined in xpuml/xpuml.h
	INIT_FLAG_NO_ATTACH = 2
	// DEVICE_UUID_BUFFER_SIZE as defined in xpuml/xpuml.h
	DEVICE_UUID_BUFFER_SIZE = 96
	// DEVICE_PART_NUMBER_BUFFER_SIZE as defined in xpuml/xpuml.h
	DEVICE_PART_NUMBER_BUFFER_SIZE = 80
	// SYSTEM_DRIVER_VERSION_BUFFER_SIZE as defined in xpuml/xpuml.h
	SYSTEM_DRIVER_VERSION_BUFFER_SIZE = 80
	// SYSTEM_XPUML_VERSION_BUFFER_SIZE as defined in xpuml/xpuml.h
	SYSTEM_XPUML_VERSION_BUFFER_SIZE = 80
	// DEVICE_NAME_BUFFER_SIZE as defined in xpuml/xpuml.h
	DEVICE_NAME_BUFFER_SIZE = 96
	// DEVICE_SERIAL_BUFFER_SIZE as defined in xpuml/xpuml.h
	DEVICE_SERIAL_BUFFER_SIZE = 30
	// DEVICE_MAX_PROCESS_COUNT as defined in xpuml/xpuml.h
	DEVICE_MAX_PROCESS_COUNT = 100
)

// XPUTopologyLevel as declared in xpuml/xpuml.h
type XPUTopologyLevel int32

// XPUTopologyLevel enumeration from xpuml/xpuml.h
const (
	TOPOLOGY_INTERNAL   XPUTopologyLevel = iota
	TOPOLOGY_SINGLE     XPUTopologyLevel = 10
	TOPOLOGY_MULTIPLE   XPUTopologyLevel = 20
	TOPOLOGY_HOSTBRIDGE XPUTopologyLevel = 30
	TOPOLOGY_NODE       XPUTopologyLevel = 40
	TOPOLOGY_SYSTEM     XPUTopologyLevel = 50
)

// XPUP2PStatus as declared in xpuml/xpuml.h
type XPUP2PStatus int32

// XPUP2PStatus enumeration from xpuml/xpuml.h
const (
	P2P_STATUS_OK                         XPUP2PStatus = iota
	P2P_STATUS_CHIPSET_NOT_SUPPORED       XPUP2PStatus = 1
	P2P_STATUS_XPU_NOT_SUPPORTED          XPUP2PStatus = 2
	P2P_STATUS_IOH_TOPOLOGY_NOT_SUPPORTED XPUP2PStatus = 3
	P2P_STATUS_DISABLED_BY_REGKEY         XPUP2PStatus = 4
	P2P_STATUS_NOT_SUPPORTED              XPUP2PStatus = 5
	P2P_STATUS_UNKNOWN                    XPUP2PStatus = 6
)

// XPUP2PCapsIndex as declared in xpuml/xpuml.h
type XPUP2PCapsIndex int32

// XPUP2PCapsIndex enumeration from xpuml/xpuml.h
const (
	P2P_CAPS_INDEX_READ    XPUP2PCapsIndex = iota
	P2P_CAPS_INDEX_WRITE   XPUP2PCapsIndex = 1
	P2P_CAPS_INDEX_KLINK   XPUP2PCapsIndex = 2
	P2P_CAPS_INDEX_ATOMICS XPUP2PCapsIndex = 3
	P2P_CAPS_INDEX_PROP    XPUP2PCapsIndex = 4
	P2P_CAPS_INDEX_UNKNOWN XPUP2PCapsIndex = 5
)

// SamplingType as declared in xpuml/xpuml.h
type SamplingType int32

// SamplingType enumeration from xpuml/xpuml.h
const (
	TOTAL_POWER_SAMPLES        SamplingType = iota
	XPU_UTILIZATION_SAMPLES    SamplingType = 1
	MEMORY_UTILIZATION_SAMPLES SamplingType = 2
	ENC_UTILIZATION_SAMPLES    SamplingType = 3
	DEC_UTILIZATION_SAMPLES    SamplingType = 4
	PROCESSOR_CLK_SAMPLES      SamplingType = 5
	MEMORY_CLK_SAMPLES         SamplingType = 6
	SAMPLINGTYPE_COUNT         SamplingType = 7
)

// PcieUtilCounter as declared in xpuml/xpuml.h
type PcieUtilCounter int32

// PcieUtilCounter enumeration from xpuml/xpuml.h
const (
	PCIE_UTIL_TX_BYTES PcieUtilCounter = iota
	PCIE_UTIL_RX_BYTES PcieUtilCounter = 1
	PCIE_UTIL_COUNT    PcieUtilCounter = 2
)

// ValueType as declared in xpuml/xpuml.h
type ValueType int32

// ValueType enumeration from xpuml/xpuml.h
const (
	VALUE_TYPE_DOUBLE             ValueType = iota
	VALUE_TYPE_UNSIGNED_INT       ValueType = 1
	VALUE_TYPE_UNSIGNED_LONG      ValueType = 2
	VALUE_TYPE_UNSIGNED_LONG_LONG ValueType = 3
	VALUE_TYPE_SIGNED_LONG_LONG   ValueType = 4
	VALUE_TYPE_COUNT              ValueType = 5
)

// PerfPolicyType as declared in xpuml/xpuml.h
type PerfPolicyType int32

// PerfPolicyType enumeration from xpuml/xpuml.h
const (
	PERF_POLICY_POWER             PerfPolicyType = iota
	PERF_POLICY_THERMAL           PerfPolicyType = 1
	PERF_POLICY_SYNC_BOOST        PerfPolicyType = 2
	PERF_POLICY_BOARD_LIMIT       PerfPolicyType = 3
	PERF_POLICY_LOW_UTILIZATION   PerfPolicyType = 4
	PERF_POLICY_RELIABILITY       PerfPolicyType = 5
	PERF_POLICY_TOTAL_APP_CLOCKS  PerfPolicyType = 10
	PERF_POLICY_TOTAL_BASE_CLOCKS PerfPolicyType = 11
	PERF_POLICY_COUNT             PerfPolicyType = 12
)

// EnableState as declared in xpuml/xpuml.h
type EnableState int32

// EnableState enumeration from xpuml/xpuml.h
const (
	FEATURE_DISABLED EnableState = iota
	FEATURE_ENABLED  EnableState = 1
)

// DeviceArch as declared in xpuml/xpuml.h
type DeviceArch int32

// DeviceArch enumeration from xpuml/xpuml.h
const (
	DEVICE_ARCH_UNKNOWN DeviceArch = iota
	DEVICE_ARCH_KL1     DeviceArch = 1
	DEVICE_ARCH_KL2     DeviceArch = 2
	DEVICE_ARCH_KL3     DeviceArch = 3
	DEVICE_ARCH_COUNT   DeviceArch = 4
)

// DeviceBoard as declared in xpuml/xpuml.h
type DeviceBoard int32

// DeviceBoard enumeration from xpuml/xpuml.h
const (
	DEVICE_BOARD_UNKNOWN     DeviceBoard = iota
	DEVICE_BOARD_KL1_K100    DeviceBoard = 1
	DEVICE_BOARD_KL1_K200    DeviceBoard = 2
	DEVICE_BOARD_KL2_R200    DeviceBoard = 100
	DEVICE_BOARD_KL2_R300    DeviceBoard = 101
	DEVICE_BOARD_KL2_R200_8F DeviceBoard = 102
	DEVICE_BOARD_COUNT       DeviceBoard = 103
)

// ClockType as declared in xpuml/xpuml.h
type ClockType int32

// ClockType enumeration from xpuml/xpuml.h
const (
	CLOCK_XPU   ClockType = iota
	CLOCK_COUNT ClockType = 1
)

// DeviceModel as declared in xpuml/xpuml.h
type DeviceModel int32

// DeviceModel enumeration from xpuml/xpuml.h
const (
	DEVICE_MODEL_UNKNOWN                        DeviceModel = 9999
	DEVICE_MODEL_KL1_K200                       DeviceModel = 0
	DEVICE_MODEL_KL1_K100                       DeviceModel = 1
	DEVICE_MODEL_KL2_R200                       DeviceModel = 2
	DEVICE_MODEL_KL2_R300                       DeviceModel = 3
	DEVICE_MODEL_KL2_R200_8F                    DeviceModel = 4
	DEVICE_MODEL_KL2_R200_SRIOV_PF              DeviceModel = 200
	DEVICE_MODEL_KL2_R200_SRIOV_VF_ONE_OF_ONE   DeviceModel = 201
	DEVICE_MODEL_KL2_R200_SRIOV_VF_ONE_OF_TWO   DeviceModel = 202
	DEVICE_MODEL_KL2_R200_SRIOV_VF_ONE_OF_THREE DeviceModel = 203
	DEVICE_MODEL_COUNT                          DeviceModel = 204
)

// TemperatureThresholds as declared in xpuml/xpuml.h
type TemperatureThresholds int32

// TemperatureThresholds enumeration from xpuml/xpuml.h
const (
	TEMPERATURE_THRESHOLD_SHUTDOWN      TemperatureThresholds = iota
	TEMPERATURE_THRESHOLD_SLOWDOWN      TemperatureThresholds = 1
	TEMPERATURE_THRESHOLD_MEM_MAX       TemperatureThresholds = 2
	TEMPERATURE_THRESHOLD_XPU_MAX       TemperatureThresholds = 3
	TEMPERATURE_THRESHOLD_ACOUSTIC_MIN  TemperatureThresholds = 4
	TEMPERATURE_THRESHOLD_ACOUSTIC_CURR TemperatureThresholds = 5
	TEMPERATURE_THRESHOLD_ACOUSTIC_MAX  TemperatureThresholds = 6
	TEMPERATURE_THRESHOLD_COUNT         TemperatureThresholds = 7
)

// TemperatureSensors as declared in xpuml/xpuml.h
type TemperatureSensors int32

// TemperatureSensors enumeration from xpuml/xpuml.h
const (
	TEMPERATURE_XPU   TemperatureSensors = iota
	TEMPERATURE_COUNT TemperatureSensors = 1
)

// Return as declared in xpuml/xpuml.h
type Return int32

// Return enumeration from xpuml/xpuml.h
const (
	SUCCESS                       Return = iota
	ERROR_UNINITIALIZED           Return = 1
	ERROR_INVALID_ARGUMENT        Return = 2
	ERROR_NOT_SUPPORTED           Return = 3
	ERROR_NO_PERMISSION           Return = 4
	ERROR_ALREADY_INITIALIZED     Return = 5
	ERROR_NOT_FOUND               Return = 6
	ERROR_INSUFFICIENT_SIZE       Return = 7
	ERROR_INSUFFICIENT_POWER      Return = 8
	ERROR_DRIVER_NOT_LOADED       Return = 9
	ERROR_TIMEOUT                 Return = 10
	ERROR_IRQ_ISSUE               Return = 11
	ERROR_LIBRARY_NOT_FOUND       Return = 12
	ERROR_FUNCTION_NOT_FOUND      Return = 13
	ERROR_CORRUPTED_INFOROM       Return = 14
	ERROR_XPU_IS_LOST             Return = 15
	ERROR_RESET_REQUIRED          Return = 16
	ERROR_OPERATING_SYSTEM        Return = 17
	ERROR_LIB_RM_VERSION_MISMATCH Return = 18
	ERROR_IN_USE                  Return = 19
	ERROR_MEMORY                  Return = 20
	ERROR_NO_DATA                 Return = 21
	ERROR_VXPU_ECC_NOT_SUPPORTED  Return = 22
	ERROR_INSUFFICIENT_RESOURCES  Return = 23
	ERROR_FREQ_NOT_SUPPORTED      Return = 24
	ERROR_UNKNOWN                 Return = 999
)

// MemoryLocation as declared in xpuml/xpuml.h
type MemoryLocation int32

// MemoryLocation enumeration from xpuml/xpuml.h
const (
	MEMORY_LOCATION_GLOBAL_MEMORY MemoryLocation = iota
	MEMORY_LOCATION_L3_MEMORY     MemoryLocation = 1
	MEMORY_LOCATION_COUNT         MemoryLocation = 2
)

// XPUVirtualizationMode as declared in xpuml/xpuml.h
type XPUVirtualizationMode int32

// XPUVirtualizationMode enumeration from xpuml/xpuml.h
const (
	XPU_VIRTUALIZATION_MODE_NONE        XPUVirtualizationMode = iota
	XPU_VIRTUALIZATION_MODE_PASSTHROUGH XPUVirtualizationMode = 1
	XPU_VIRTUALIZATION_MODE_VXPU        XPUVirtualizationMode = 2
	XPU_VIRTUALIZATION_MODE_HOST_VXPU   XPUVirtualizationMode = 3
	XPU_VIRTUALIZATION_MODE_HOST_VSGA   XPUVirtualizationMode = 4
)

// HostVxpuMode as declared in xpuml/xpuml.h
type HostVxpuMode int32

// HostVxpuMode enumeration from xpuml/xpuml.h
const (
	HOST_VXPU_MODE_NON_SRIOV HostVxpuMode = iota
	HOST_VXPU_MODE_SRIOV_OFF HostVxpuMode = 1
	HOST_VXPU_MODE_SRIOV_ON  HostVxpuMode = 2
)

// VxpuVmIdType as declared in xpuml/xpuml.h
type VxpuVmIdType int32

// VxpuVmIdType enumeration from xpuml/xpuml.h
const (
	VXPU_VM_ID_DOMAIN_ID VxpuVmIdType = iota
	VXPU_VM_ID_UUID      VxpuVmIdType = 1
)

// VxpuGuestInfoState as declared in xpuml/xpuml.h
type VxpuGuestInfoState int32

// VxpuGuestInfoState enumeration from xpuml/xpuml.h
const (
	VXPU_INSTANCE_GUEST_INFO_STATE_UNINITIALIZED VxpuGuestInfoState = iota
	VXPU_INSTANCE_GUEST_INFO_STATE_INITIALIZED   VxpuGuestInfoState = 1
)

// DeviceState as declared in xpuml/xpuml.h
type DeviceState int32

// DeviceState enumeration from xpuml/xpuml.h
const (
	DEVICE_STATE_UNUSED       DeviceState = iota
	DEVICE_STATE_RUNNING      DeviceState = 201
	DEVICE_STATE_IN_RESET     DeviceState = 205
	DEVICE_STATE_ERROR        DeviceState = 207
	DEVICE_STATE_IN_EXCEPTION DeviceState = 208
)

// DetachXPUState as declared in xpuml/xpuml.h
type DetachXPUState int32

// DetachXPUState enumeration from xpuml/xpuml.h
const (
	DETACH_XPU_KEEP   DetachXPUState = iota
	DETACH_XPU_REMOVE DetachXPUState = 1
)

// PcieLinkState as declared in xpuml/xpuml.h
type PcieLinkState int32

// PcieLinkState enumeration from xpuml/xpuml.h
const (
	PCIE_LINK_KEEP      PcieLinkState = iota
	PCIE_LINK_SHUT_DOWN PcieLinkState = 1
)
