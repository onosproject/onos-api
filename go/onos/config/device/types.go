// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package device

import (
	"fmt"
	types "github.com/onosproject/onos-api/go/onos/config"
	"strings"
)

const separator = ":"

// ID is a device ID
type ID string

// Type is a device type
type Type string

// Role is a device role
type Role string

// Version is a device version
type Version string

// VersionedID is a versioned device ID
type VersionedID types.ID

// NewVersionedID returns a new versioned device identifier
func NewVersionedID(id ID, version Version) VersionedID {
	return VersionedID(fmt.Sprintf("%s%s%s", id, separator, version))
}

// GetID returns the device ID
func (i VersionedID) GetID() ID {
	return ID(i[:strings.Index(string(i), separator)])
}

// GetVersion returns the device version
func (i VersionedID) GetVersion() Version {
	return Version(i[strings.Index(string(i), separator)+1:])
}
