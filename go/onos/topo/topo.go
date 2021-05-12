// Copyright 2019-present Open Networking Foundation.
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
	"errors"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"google.golang.org/grpc"
)

// ID ...
type ID string

// NullID ...
const NullID = ""

// Revision is an object revision
type Revision uint64

// Attribute keys
const (
	Address     = "address"
	Target      = "target"
	Type        = "type"
	Version     = "version"
	Timeout     = "timeout"
	Role        = "role"
	Displayname = "displayname"
	User        = "user"
	Password    = "password"
	TLSCaCert   = "tls-ca-cert"
	TLSCert     = "tls-cert"
	TLSKey      = "tls-key"
	TLSPlain    = "tls-plain"
	TLSInsecure = "tls-insecure"
)

// TopoClientFactory : Default EntityServiceClient creation.
var TopoClientFactory = func(cc *grpc.ClientConn) TopoClient {
	return NewTopoClient(cc)
}

// CreateTopoClient creates and returns a new topo device client
func CreateTopoClient(cc *grpc.ClientConn) TopoClient {
	return TopoClientFactory(cc)
}

// GetAspectSafe retrieves the specified aspect value from the given object.
func (obj *Object) GetAspectSafe(key string, destValue proto.Message) (proto.Message, error) {
	any := obj.Aspects[key]
	if !types.Is(any, destValue) {
		return nil, errors.New("unexpected aspect type")
	}
	err := types.UnmarshalAny(any, destValue)
	if err != nil {
		return nil, err
	}
	return destValue, nil
}

// GetAspect retrieves the specified aspect value from the given object.
func (obj *Object) GetAspect(key string, destValue proto.Message) proto.Message {
	any := obj.Aspects[key]
	if !types.Is(any, destValue) {
		return nil
	}
	err := types.UnmarshalAny(any, destValue)
	if err != nil {
		return nil
	}
	return destValue
}

// SetAspect applies the specified aspect value to the given object.
func (obj *Object) SetAspect(key string, value proto.Message) error {
	any, err := types.MarshalAny(value)
	if err != nil {
		return err
	}
	obj.Aspects[key] = any
	return nil
}