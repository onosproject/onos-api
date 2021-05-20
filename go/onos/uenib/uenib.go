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

package uenib

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

// UEServiceClientFactory : Default UEServiceClient creation.
var UEServiceClientFactory = func(cc *grpc.ClientConn) UEServiceClient {
	return NewUEServiceClient(cc)
}

// CreateUEServiceClient creates and returns a new UE service client
func CreateUEServiceClient(cc *grpc.ClientConn) UEServiceClient {
	return UEServiceClientFactory(cc)
}

// GetAspect retrieves the specified aspect value from the given UE.
func (ue *UE) GetAspect(destValue proto.Message) proto.Message {
	if ue.Aspects == nil {
		return nil
	}
	aspectType := proto.MessageName(destValue)
	any := ue.Aspects[aspectType]
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

// GetAspectSafe retrieves the specified aspect value from the given UE.
func (ue *UE) GetAspectSafe(destValue proto.Message) (proto.Message, error) {
	if ue.Aspects == nil {
		return nil, errors.New("aspect not found")
	}
	aspectType := proto.MessageName(destValue)
	any := ue.Aspects[aspectType]
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

// GetAspectBytes applies the specified aspect as raw JSON bytes to the given UE.
func (ue *UE) GetAspectBytes(aspectType string) ([]byte, error) {
	if ue.Aspects == nil {
		return nil, errors.New("aspect not found")
	}
	any := ue.Aspects[aspectType]
	if any == nil {
		return nil, errors.New("aspect not found")
	}
	return any.Value, nil
}

// SetAspect applies the specified aspect value to the given ueect.
func (ue *UE) SetAspect(value proto.Message) error {
	jm := jsonpb.Marshaler{}
	writer := bytes.Buffer{}
	err := jm.Marshal(&writer, value)
	if err != nil {
		return err
	}
	if ue.Aspects == nil {
		ue.Aspects = make(map[string]*types.Any)
	}
	ue.Aspects[proto.MessageName(value)] = &types.Any{
		TypeUrl: proto.MessageName(value),
		Value: writer.Bytes(),
	}
	return nil
}

// SetAspectBytes applies the specified aspect as raw JSON bytes to the given UE.
func (ue *UE) SetAspectBytes(aspectType string, jsonValue []byte) error {
	any := &types.Any {
		TypeUrl: aspectType,
		Value: jsonValue,
	}
	if ue.Aspects == nil {
		ue.Aspects = make(map[string]*types.Any)
	}
	ue.Aspects[aspectType] = any
	return nil
}
