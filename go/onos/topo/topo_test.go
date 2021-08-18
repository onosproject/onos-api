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

package topo

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAspects(t *testing.T) {
	o := &Object{}

	// Test retrieval of non-existent aspects (before any aspects exist)
	err := o.GetAspect(&TLSOptions{})
	assert.Error(t, err)
	_, err = o.GetAspectBytes("onos.topo.TLSOptions")
	assert.Error(t, err)

	// Set an aspect
	err = o.SetAspect(&Location{Lat: 1.23, Lng: 3.21})
	assert.NoError(t, err)

	// Fetch it...
	loc := &Location{}
	err = o.GetAspect(loc)
	assert.NoError(t, err)
	assert.NotNil(t, loc)
	assert.Equal(t, 1.23, loc.Lat)
	assert.Equal(t, 3.21, loc.Lng)

	// Set it as bytes...
	err = o.SetAspectBytes("onos.topo.Location", bytes.NewBufferString("{\"lat\":3.14,\"lng\":6.28}").Bytes())
	assert.NoError(t, err)

	// Fetch it... again.
	loc = &Location{}
	err = o.GetAspect(loc)
	assert.NoError(t, err)
	assert.Equal(t, 3.14, loc.Lat)
	assert.Equal(t, 6.28, loc.Lng)

	// Fetch it... as bytes...
	b, err := o.GetAspectBytes("onos.topo.Location")
	assert.NoError(t, err)
	assert.Equal(t, "{\"lat\":3.14,\"lng\":6.28}", bytes.NewBuffer(b).String())

	// Test retrieval of non-existent aspects (after some aspects exist)
	tls := &TLSOptions{}
	err = o.GetAspect(tls)
	assert.Error(t, err)
	_, err = o.GetAspectBytes("onos.topo.TLSOptions")
	assert.Error(t, err)
}

func TestRelationID(t *testing.T) {
	id := RelationID("foo", "is", "bar")
	assert.Equal(t, ID("foo-is-bar"), id)

	id = MultiRelationID("foo", "implies", "bar", 42)
	assert.Equal(t, ID("foo-implies-bar-42"), id)
}
