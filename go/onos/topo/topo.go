// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package topo

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"google.golang.org/grpc"
)

// UUID represents a system-assigned unique identifier of a topology object.
type UUID string

// ID represents a client-assigned unique identifier.
type ID string

// NullID represents a null/empty/omitted identifier; usually an indicator for system to generate one.
const NullID = ""

// Revision is an object revision
type Revision uint64

// Entity and Relation Kinds
const (
	// Relations
	CONTROLS  = "controls"
	CONTAINS  = "contains"
	NEIGHBORS = "neighbors"

	// RAN Entities
	E2NODE = "e2node"
	E2CELL = "e2cell"
	E2T    = "e2t"
	XAPP   = "xapp"
	A1T    = "a1t"

	// onos-config entity
	ONOS_CONFIG = "onos-config"
)

// PolicyTypeID is an identifier of A1 policy type
type PolicyTypeID string

// PolicyTypeName is a name of A1 policy type
type PolicyTypeName string

// PolicyTypeVersion is a version of A1 policy type
type PolicyTypeVersion string

// PolicyTypeDescription describe what this A1 policy is
type PolicyTypeDescription string

// TopoClientFactory : Default EntityServiceClient creation.
var TopoClientFactory = func(cc *grpc.ClientConn) TopoClient {
	return NewTopoClient(cc)
}

// CreateTopoClient creates and returns a new topo device client
func CreateTopoClient(cc *grpc.ClientConn) TopoClient {
	return TopoClientFactory(cc)
}

// RelationID creates a unique relationship ID from the specified source, kind and target IDs
func RelationID(srcID ID, relationKind ID, tgtID ID) ID {
	return ID(fmt.Sprintf("%s-%s-%s", srcID, relationKind, tgtID))
}

// MultiRelationID creates a unique relationship ID from the specified source, kind and target IDs,
// and also from an additional discriminant to allow for multiples of same kinds of relations between
// the same two objects.
func MultiRelationID(srcID ID, relationKind ID, tgtID ID, discriminant uint8) ID {
	return ID(fmt.Sprintf("%s-%s-%s-%d", srcID, relationKind, tgtID, discriminant))
}

// GetAspect retrieves the specified aspect value from the given object.
func (obj *Object) GetAspect(destValue proto.Message) error {
	if obj.Aspects == nil {
		return errors.New("aspect not found")
	}
	aspectType := proto.MessageName(destValue)
	any := obj.Aspects[aspectType]
	if any == nil {
		return errors.New("aspect not found")
	}
	if any.TypeUrl != aspectType {
		return errors.New("unexpected aspect type")
	}
	reader := bytes.NewReader(any.Value)
	err := jsonpb.Unmarshal(reader, destValue)
	if err != nil {
		return err
	}
	return nil
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
		Value:   writer.Bytes(),
	}
	return nil
}

// SetAspectBytes applies the specified aspect as raw JSON bytes to the given object.
func (obj *Object) SetAspectBytes(aspectType string, jsonValue []byte) error {
	any := &types.Any{
		TypeUrl: aspectType,
		Value:   jsonValue,
	}
	if obj.Aspects == nil {
		obj.Aspects = make(map[string]*types.Any)
	}
	obj.Aspects[aspectType] = any
	return nil
}
