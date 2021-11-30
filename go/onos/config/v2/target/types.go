// Copyright 2021-present Open Networking Foundation.
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

package target

import (
	"fmt"
	"strings"

	types "github.com/onosproject/onos-api/go/onos/config"
)

const separator = ":"

// ID is a target ID
type ID string

// Type is a target type
type Type string

// Role is a target role
type Role string

// Version is a target version
type Version string

// VersionedID is a versioned target ID
type VersionedID types.ID

// NewVersionedID returns a new versioned target identifier
func NewVersionedID(id ID, version Version) VersionedID {
	return VersionedID(fmt.Sprintf("%s%s%s", id, separator, version))
}

// GetID returns the target ID
func (i VersionedID) GetID() ID {
	return ID(i[:strings.Index(string(i), separator)])
}

// GetVersion returns the target version
func (i VersionedID) GetVersion() Version {
	return Version(i[strings.Index(string(i), separator)+1:])
}
