// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package fabricsim

import "fmt"

// DeviceID is a unique identifier of a simulated device
type DeviceID string

// PortID is a unique identifier of a simulated device port
type PortID string

// HostID is a unique identifier of a simulated host
type HostID string

// LinkID is a unique identifier of a simulated link
type LinkID string

// NewLinkID produces a singular link ID from the specified source and target port IDs
func NewLinkID(src PortID, tgt PortID) LinkID {
	return LinkID(fmt.Sprintf("%s-%s", src, tgt))
}
