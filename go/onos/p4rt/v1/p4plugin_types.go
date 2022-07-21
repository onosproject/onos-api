// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package p4rt

import "fmt"

// P4PluginID p4 plugin ID
type P4PluginID string

// NewP4PluginID creates a P4 plugin ID based on a given name, version, and architecture
func NewP4PluginID(name, version, arch string) P4PluginID {
	return P4PluginID(fmt.Sprintf("%s-%s-%s", name, version, arch))
}
