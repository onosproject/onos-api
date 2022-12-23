// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package provisioner

// ConfigID is configuration identifier
type ConfigID string

// Revision is configuration revision
type Revision uint64

const (
	// PipelineConfigKind is configuration kind for pipeline configs
	PipelineConfigKind = "pipeline"

	// ChassisConfigKind is configuration kind for chassis configs
	ChassisConfigKind = "chassis"

	// P4InfoType artifact holds p4info proto text
	P4InfoType = "p4info"
	// P4BinaryType artifact holds device P4 binary
	P4BinaryType = "p4bin"
	// ChassisType artifact holds chassis protobuf structure
	ChassisType = "chassis"
)
