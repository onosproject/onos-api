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

// GetAttributeSafe retrieves the specified attribute value from the given object.
func GetAttributeSafe(obj *Object, key string, destValue proto.Message) (proto.Message, error) {
	err := types.UnmarshalAny(obj.Attributes[key], destValue)
	if err != nil {
		return nil, err
	}
	return destValue, nil
}

// GetAttribute retrieves the specified attribute value from the given object.
func GetAttribute(obj *Object, key string, destValue proto.Message) proto.Message {
	err := types.UnmarshalAny(obj.Attributes[key], destValue)
	if err != nil {
		return nil
	}
	return destValue
}

// SetAttribute applies the specified attribute value to the given object.
func SetAttribute(obj *Object, key string, value proto.Message) error {
	any, err := types.MarshalAny(value)
	if err != nil {
		return err
	}
	obj.Attributes[key] = any
	return nil
}