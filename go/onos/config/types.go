// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package config

import "github.com/google/uuid"

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
