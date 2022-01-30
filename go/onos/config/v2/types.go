// Copyright 2022-present Open Networking Foundation.
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

package v2

import (
	"fmt"

	"github.com/google/uuid"
	types "github.com/onosproject/onos-api/go/onos/config"

	"strings"
)

const separator = ":"

// ID is an identifier type
type ID string

// Index is the index of an object
type Index uint64

// Revision is a revision number
type Revision uint64

// NewUUID generates a new uuid
func NewUUID() uuid.UUID {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		newUUID = uuid.New()
	}
	return newUUID
}

// ConfigurationID is a configuration identifier type
type ConfigurationID string

// TransactionID is a transaction identifier type
type TransactionID string

// TargetID is a target ID
type TargetID string

// TargetType is a target type
type TargetType string

// TargetRole is a target role
type TargetRole string

// TargetVersion is a target version
type TargetVersion string

// TargetVersionedID is a versioned target ID
type TargetVersionedID types.ID

// MastershipTerm mastership term
type MastershipTerm uint64

// NewTargetVersionedID returns a new versioned target identifier
func NewTargetVersionedID(id ID, version TargetVersion) TargetVersionedID {
	return TargetVersionedID(fmt.Sprintf("%s%s%s", id, separator, version))
}

// GetID returns the target ID
func (i TargetVersionedID) GetID() ID {
	return ID(i[:strings.Index(string(i), separator)])
}

// GetVersion returns the target version
func (i TargetVersionedID) GetVersion() TargetVersion {
	return TargetVersion(i[strings.Index(string(i), separator)+1:])
}
