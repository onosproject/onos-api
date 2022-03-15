// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package device

import (
	"gotest.tools/assert"
	"testing"
)

func Test_NewVersionedId(t *testing.T) {
	id := ID("Device1")
	version := Version("1.2.3")
	vID := NewVersionedID(id, version)
	assert.Equal(t, vID.GetID(), id)
	assert.Equal(t, vID.GetVersion(), version)
}
