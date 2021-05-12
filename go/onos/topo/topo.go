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
	"bytes"
	"errors"
	"github.com/gogo/protobuf/jsonpb"
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
	Address = "address"
	Target  = "target"
	Type    = "type"
	Version = "version"
	Timeout = "timeout"
	Role    = "role"
)

// TopoClientFactory : Default EntityServiceClient creation.
var TopoClientFactory = func(cc *grpc.ClientConn) TopoClient {
	return NewTopoClient(cc)
}

// CreateTopoClient creates and returns a new topo device client
func CreateTopoClient(cc *grpc.ClientConn) TopoClient {
	return TopoClientFactory(cc)
}

// GetAspect retrieves the specified aspect value from the given object.
func (obj *Object) GetAspect(destValue proto.Message) proto.Message {
	if obj.Aspects == nil {
		return nil
	}
	aspectType := proto.MessageName(destValue)
	any := obj.Aspects[aspectType]
	if any == nil || any.TypeUrl != aspectType {
		return nil
	}
	reader := bytes.NewReader(any.Value)
	err := jsonpb.Unmarshal(reader, destValue)
	if err != nil {
		return nil
	}
	return destValue
}

// GetAspectSafe retrieves the specified aspect value from the given object.
func (obj *Object) GetAspectSafe(destValue proto.Message) (proto.Message, error) {
	if obj.Aspects == nil {
		return nil, errors.New("aspect not found")
	}
	aspectType := proto.MessageName(destValue)
	any := obj.Aspects[aspectType]
	if any == nil {
		return nil, errors.New("aspect not found")
	}
	if any.TypeUrl != aspectType {
		return nil, errors.New("unexpected aspect type")
	}
	reader := bytes.NewReader(any.Value)
	err := jsonpb.Unmarshal(reader, destValue)
	if err != nil {
		return nil, err
	}
	return destValue, nil
}

// GetAspectBytes applies the specified aspect as raw JSON bytes to the given object.
func (obj *Object) GetAspectBytes(aspectType string) ([]byte, error) {
	if obj.Aspects == nil {
		return nil, errors.New("aspect not found")
	}
	any := obj.Aspects[aspectType]
	if any == nil {
		return nil, errors.New("aspect not found")
	}
	return any.Value, nil
}

// SetAspect applies the specified aspect value to the given object.
func (obj *Object) SetAspect(value proto.Message) error {
	jm := jsonpb.Marshaler{}
	writer := bytes.Buffer{}
	err := jm.Marshal(&writer, value)
	if err != nil {
		return err
	}
	if obj.Aspects == nil {
		obj.Aspects = make(map[string]*types.Any)
	}
	obj.Aspects[proto.MessageName(value)] = &types.Any{
		TypeUrl: proto.MessageName(value),
		Value: writer.Bytes(),
	}
	return nil
}

// SetAspectBytes applies the specified aspect as raw JSON bytes to the given object.
func (obj *Object) SetAspectBytes(aspectType string, jsonValue []byte) error {
	any := &types.Any {
		TypeUrl: aspectType,
		Value: jsonValue,
	}
	if obj.Aspects == nil {
		obj.Aspects = make(map[string]*types.Any)
	}
	obj.Aspects[aspectType] = any
	return nil
}
