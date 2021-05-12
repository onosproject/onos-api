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

	err := o.SetAspect(&Location{Lat: 1.23, Lng: 3.21})
	assert.NoError(t, err)

	loc := o.GetAspect(&Location{}).(*Location)
	assert.NotNil(t, loc)
	assert.Equal(t, 1.23, loc.Lat)
	assert.Equal(t, 3.21, loc.Lng)

	err = o.SetAspectJSON("onos.topo.Location", bytes.NewBufferString("{\"lat\":3.14,\"lng\":6.28}").Bytes())
	assert.NoError(t, err)

	msg, err := o.GetAspectSafe(&Location{})
	assert.NoError(t, err)
	loc = msg.(*Location)
	assert.Equal(t, 3.14, loc.Lat)
	assert.Equal(t, 6.28, loc.Lng)
}
