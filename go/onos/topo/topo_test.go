// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

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
