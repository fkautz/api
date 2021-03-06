// Copyright (c) 2020 Cisco Systems, Inc.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kernel

import (
	"github.com/networkservicemesh/api/pkg/api/networkservice/mechanisms/common"
)

const (

	// MECHANISM string
	MECHANISM = "KERNEL_INTERFACE"

	// Parameters

	// LinuxIfMaxLength - Linux has a limit of 15 characters for an interface name
	LinuxIfMaxLength = 15

	// SocketFilename - socket filename memif mechanism property key
	SocketFilename = "socketfile"
	// Master - NSMgr name
	Master = "master"
	// Slave - NSMgr name
	Slave = "slave"
	// WorkspaceNSEName - NSE workspace name mechanism property key
	WorkspaceNSEName = "workspaceNseName"
	// MemifSocket - memif socket filename
	MemifSocket = "memif.sock"
	// NSMBaseDirEnv - NSM location directory
	NSMBaseDirEnv = "NSM_BASE_DIR"

	// PCIAddress string
	PCIAddress = "PCIAddress"

	// NetNSInodeKey - netns inode mechanism property key
	NetNSInodeKey = common.NetNSInodeKey

	// InterfaceNameKey - interface name mechanism property key
	InterfaceNameKey = common.InterfaceNameKey
	// InterfaceDescriptionKey - interface description mechanism property key
	InterfaceDescriptionKey = common.InterfaceNameKey
)
