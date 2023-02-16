// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package topo

import (
	"bytes"
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

// String convert ID to string
func (id ID) String() string {
	return string(id)
}

// NullID represents a null/empty/omitted identifier; usually an indicator for system to generate one.
const NullID = ""

// Revision is an object revision
type Revision uint64

// DEPRECATED Entity and Relation Kinds
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

// TODO UPPERCASE entity kinds and relations should be replaced gradually with CamelCase ones
const (
	ControlsKind   = "controls"
	ContainsKind   = "contains"
	HasKind        = "has"
	TerminatesKind = "terminates"
	OriginatesKind = "originates"
	NeighborsKind  = "neighbors"
	ConnectionKind = "connection"

	// Fabric Entity kinds
	PodKind          = "pod"
	RackKind         = "rack"
	NetworkLayerKind = "network-layer"
	SwitchKind       = "switch"
	ServerKind       = "server"
	IPUKind          = "ipu"
	HostKind         = "host"
	RouterKind       = "router"
	PortKind         = "port"
	InterfaceKind    = "interface"
	LinkKind         = "link"
	ControllerKind   = "controller"
	ServiceKind      = "service"

	// onos-config entity
	OnosConfigKind = "onos-config"

	// RAN Entitiy kinds
	E2NodeKind = "e2node"
	E2CellKind = "e2cell"
	E2tKind    = "e2t"
	XappKind   = "xapp"
	A1tKind    = "a1t"
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

// NewEntity allocates a new topology entity using the specified ID and kind.
func NewEntity(id ID, kind ID, aspects ...proto.Message) *Object {
	return &Object{ID: id, Type: Object_ENTITY, Obj: &Object_Entity{Entity: &Entity{KindID: kind}}}
}

// NewRelation allocates a new topology relation using the specified source, target, and kind.
func NewRelation(source ID, target ID, kind ID, aspects ...proto.Message) *Object {
	return &Object{
		ID:   RelationID(source, kind, target),
		Type: Object_RELATION,
		Obj:  &Object_Relation{Relation: &Relation{SrcEntityID: source, TgtEntityID: target, KindID: kind}},
	}
}

// WithAspects applies the given aspects to the object.
func (obj *Object) WithAspects(aspects ...proto.Message) (*Object, error) {
	for _, aspect := range aspects {
		if err := obj.SetAspect(aspect); err != nil {
			return obj, err
		}
	}
	return obj, nil
}

// ToAny provides a convenience utility to convert an aspect message to protobuf types.Any
func ToAny(value proto.Message) *types.Any {
	jm := jsonpb.Marshaler{}
	writer := bytes.Buffer{}
	if err := jm.Marshal(&writer, value); err != nil {
		return nil
	}
	return &types.Any{
		TypeUrl: proto.MessageName(value),
		Value:   writer.Bytes(),
	}
}

// GetAspect retrieves the specified aspect value from the given object.
func (obj *Object) GetAspect(destValue proto.Message) error {
	if obj.Aspects == nil {
		return fmt.Errorf("no aspects found on %s", obj.String())
	}
	aspectType := proto.MessageName(destValue)
	any := obj.Aspects[aspectType]
	if any == nil {
		return fmt.Errorf("aspect '%s' not found in %s", aspectType, obj.String())
	}
	if any.TypeUrl != aspectType {
		return fmt.Errorf("unexpected aspect type: %s", aspectType)
	}
	reader := bytes.NewReader(any.Value)
	err := jsonpb.Unmarshal(reader, destValue)
	if err != nil {
		return fmt.Errorf("error '%s' when unmarshalling aspect %s: %v from %s",
			err.Error(), aspectType, any.Value, obj.String())
	}
	return nil
}

// GetAspectBytes applies the specified aspect as raw JSON bytes to the given object.
func (obj *Object) GetAspectBytes(aspectType string) ([]byte, error) {
	if obj.Aspects == nil {
		return nil, fmt.Errorf("no aspects found on %s", obj.String())
	}
	any := obj.Aspects[aspectType]
	if any == nil {
		return nil, fmt.Errorf("aspect '%s' not found on %s", aspectType, obj.String())
	}
	return any.Value, nil
}

// SetAspect applies the specified aspect value to the given object.
func (obj *Object) SetAspect(value proto.Message) error {
	jm := jsonpb.Marshaler{}
	writer := bytes.Buffer{}
	err := jm.Marshal(&writer, value)
	if err != nil {
		return fmt.Errorf("error '%s' marshaling aspect %v on to %s",
			err.Error(), value, obj.String())
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
