# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: onos/topo/topo.proto
# plugin: python-betterproto
from dataclasses import dataclass
from typing import AsyncIterator, Dict, List, Optional

import betterproto
import grpclib


class EventType(betterproto.Enum):
    """EventType is a topo operation event type"""

    NONE = 0
    ADDED = 1
    UPDATED = 2
    REMOVED = 3


class Protocol(betterproto.Enum):
    """Protocol to interact with a device"""

    # UNKNOWN_PROTOCOL constant needed to go around proto3 nullifying the 0
    # values
    UNKNOWN_PROTOCOL = 0
    # GNMI protocol reference
    GNMI = 1
    # P4RUNTIME protocol reference
    P4RUNTIME = 2
    # GNOI protocol reference
    GNOI = 3
    # E2 Control Plane Protocol
    E2AP = 4


class ConnectivityState(betterproto.Enum):
    """
    ConnectivityState represents the L3 reachability of a device from the
    service container (e.g. enos-config), independently of gRPC or the service
    itself (e.g. gNMI)
    """

    # UNKNOWN_CONNECTIVITY_STATE constant needed to go around proto3 nullifying
    # the 0 values
    UNKNOWN_CONNECTIVITY_STATE = 0
    # REACHABLE indicates the the service can reach the device at L3
    REACHABLE = 1
    # UNREACHABLE indicates the the service can't reach the device at L3
    UNREACHABLE = 2


class ChannelState(betterproto.Enum):
    """
    ConnectivityState represents the state of a gRPC channel to the device from
    the service container
    """

    # UNKNOWN_CHANNEL_STATE constant needed to go around proto3 nullifying the 0
    # values
    UNKNOWN_CHANNEL_STATE = 0
    # CONNECTED indicates the corresponding grpc channel is connected on this
    # device
    CONNECTED = 1
    # DISCONNECTED indicates the corresponding grpc channel is not connected on
    # this device
    DISCONNECTED = 2


class ServiceState(betterproto.Enum):
    """
    ServiceState represents the state of the gRPC service (e.g. gNMI) to the
    device from the service container
    """

    # UNKNOWN_SERVICE_STATE constant needed to go around proto3 nullifying the 0
    # values
    UNKNOWN_SERVICE_STATE = 0
    # AVAILABLE indicates the corresponding grpc service is available
    AVAILABLE = 1
    # UNAVAILABLE indicates the corresponding grpc service is not available
    UNAVAILABLE = 2
    # CONNECTING indicates the corresponding protocol is in the connecting phase
    # on this device
    CONNECTING = 3


class ObjectType(betterproto.Enum):
    UNSPECIFIED = 0
    ENTITY = 1
    RELATION = 2
    KIND = 3


@dataclass(eq=False, repr=False)
class Event(betterproto.Message):
    """Event is a topo operation event"""

    type: "EventType" = betterproto.enum_field(1)
    object: "Object" = betterproto.message_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class CreateRequest(betterproto.Message):
    object: "Object" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class CreateResponse(betterproto.Message):
    object: "Object" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class GetRequest(betterproto.Message):
    id: str = betterproto.string_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class GetResponse(betterproto.Message):
    object: "Object" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class UpdateRequest(betterproto.Message):
    object: "Object" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class UpdateResponse(betterproto.Message):
    object: "Object" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class DeleteRequest(betterproto.Message):
    id: str = betterproto.string_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class DeleteResponse(betterproto.Message):
    pass

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class ListRequest(betterproto.Message):
    pass

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class ListResponse(betterproto.Message):
    objects: List["Object"] = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class WatchRequest(betterproto.Message):
    noreplay: bool = betterproto.bool_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class WatchResponse(betterproto.Message):
    event: "Event" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class Object(betterproto.Message):
    id: str = betterproto.string_field(1)
    revision: int = betterproto.uint64_field(2)
    type: "ObjectType" = betterproto.enum_field(3)
    entity: "Entity" = betterproto.message_field(4, group="obj")
    relation: "Relation" = betterproto.message_field(5, group="obj")
    kind: "Kind" = betterproto.message_field(6, group="obj")
    attributes: Dict[str, str] = betterproto.map_field(
        7, betterproto.TYPE_STRING, betterproto.TYPE_STRING
    )

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class Entity(betterproto.Message):
    """Entity represents any "thing" that is represented in the topology"""

    # user-defined entity kind
    kind_id: str = betterproto.string_field(1)
    protocols: List["ProtocolState"] = betterproto.message_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class Relation(betterproto.Message):
    # user defined relation kind
    kind_id: str = betterproto.string_field(1)
    src_entity_id: str = betterproto.string_field(2)
    tgt_entity_id: str = betterproto.string_field(3)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class Kind(betterproto.Message):
    name: str = betterproto.string_field(1)
    # Map of attributes and their default values for this Kind
    attributes: Dict[str, str] = betterproto.map_field(
        2, betterproto.TYPE_STRING, betterproto.TYPE_STRING
    )

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class ProtocolState(betterproto.Message):
    """
    ProtocolState contains information related to service and connectivity to a
    device
    """

    # The protocol to which state relates
    protocol: "Protocol" = betterproto.enum_field(1)
    # ConnectivityState contains the L3 connectivity information
    connectivity_state: "ConnectivityState" = betterproto.enum_field(2)
    # ChannelState relates to the availability of the gRPC channel
    channel_state: "ChannelState" = betterproto.enum_field(3)
    # ServiceState indicates the availability of the gRPC servic on top of the
    # channel
    service_state: "ServiceState" = betterproto.enum_field(4)

    def __post_init__(self) -> None:
        super().__post_init__()


class TopoStub(betterproto.ServiceStub):
    """EntityService provides an API for managing entities."""

    async def create(self, *, object: "Object" = None) -> "CreateResponse":
        """Create a new topology object"""

        request = CreateRequest()
        if object is not None:
            request.object = object

        return await self._unary_unary(
            "/onos.topo.Topo/Create", request, CreateResponse
        )

    async def get(self, *, id: str = "") -> "GetResponse":
        """Get an object from topology"""

        request = GetRequest()
        request.id = id

        return await self._unary_unary("/onos.topo.Topo/Get", request, GetResponse)

    async def update(self, *, object: "Object" = None) -> "UpdateResponse":
        """Update an existing topology object"""

        request = UpdateRequest()
        if object is not None:
            request.object = object

        return await self._unary_unary(
            "/onos.topo.Topo/Update", request, UpdateResponse
        )

    async def delete(self, *, id: str = "") -> "DeleteResponse":
        """Delete an object from topology"""

        request = DeleteRequest()
        request.id = id

        return await self._unary_unary(
            "/onos.topo.Topo/Delete", request, DeleteResponse
        )

    async def list(self) -> "ListResponse":
        """List gets a stream of requested objects"""

        request = ListRequest()

        return await self._unary_unary("/onos.topo.Topo/List", request, ListResponse)

    async def watch(self, *, noreplay: bool = False) -> AsyncIterator["WatchResponse"]:
        """Watch returns a stream of topo change notifications"""

        request = WatchRequest()
        request.noreplay = noreplay

        async for response in self._unary_stream(
            "/onos.topo.Topo/Watch",
            request,
            WatchResponse,
        ):
            yield response
