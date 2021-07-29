# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: onos/config/diags/diags.proto
# plugin: python-betterproto
from dataclasses import dataclass
from typing import AsyncIterator, Dict

import betterproto
from betterproto.grpc.grpclib_server import ServiceBase
import grpclib


class Type(betterproto.Enum):
    """Change (Network or Device) event type"""

    # NONE indicates this response does not represent a modification of the
    # Change
    NONE = 0
    # ADDED is an event which occurs when a Change is added to the topology
    ADDED = 1
    # UPDATED is an event which occurs when a Change is updated
    UPDATED = 2
    # REMOVED is an event which occurs when a Change is removed from the
    # configuration
    REMOVED = 3


@dataclass(eq=False, repr=False)
class OpStateRequest(betterproto.Message):
    """
    OpStateRequest is a message for specifying GetOpState query parameters.
    """

    # The request is always in the context of a Device ID. If the device does not
    # exist or is disconnected an error will be returned.
    device_id: str = betterproto.string_field(1)
    # subscribe indicates whether to subscribe to events (e.g. ADD, UPDATE, and
    # REMOVE) that occur after all paths for the device have been streamed to the
    # client
    subscribe: bool = betterproto.bool_field(2)


@dataclass(eq=False, repr=False)
class OpStateResponse(betterproto.Message):
    # type is the type of the event
    type: "_admin__.Type" = betterproto.enum_field(1)
    # device is the device on which the event occurred
    pathvalue: "_change_device__.PathValue" = betterproto.message_field(2)


@dataclass(eq=False, repr=False)
class ListNetworkChangeRequest(betterproto.Message):
    """
    ListNetworkChangeRequest requests a stream of changes and updates to them
    By default, the request requests a stream of all changes that are present
    in the topology when the request is received by the service. However, if
    `subscribe` is `true`, the stream will remain open after all changes have
    been sent and events that occur following the last changes will be streamed
    to the client until the stream is closed. If "withoutReplay" is true then
    only changes that happen after the call will be returned
    """

    # subscribe indicates whether to subscribe to events (e.g. ADD, UPDATE, and
    # REMOVE) that occur after all devices have been streamed to the client
    subscribe: bool = betterproto.bool_field(1)
    # option to specify a specific network change - if blank or '*' then select
    # all Can support `*` (match many chars) or '?' (match one char) as wildcard
    changeid: str = betterproto.string_field(2)
    # option to request only changes that happen after the call
    without_replay: bool = betterproto.bool_field(3)


@dataclass(eq=False, repr=False)
class ListNetworkChangeResponse(betterproto.Message):
    """ListNetworkChangeResponse carries a single network change event"""

    # change is the network change on which the event occurred
    change: "_change_network__.NetworkChange" = betterproto.message_field(1)
    # type is a qualification of the type of change being made
    type: "Type" = betterproto.enum_field(2)


@dataclass(eq=False, repr=False)
class ListDeviceChangeRequest(betterproto.Message):
    """
    ListDeviceChangeRequest requests a stream of changes and updates to them By
    default, the request requests a stream of all changes that are present in
    the topology when the request is received by the service. However, if
    `subscribe` is `true`, the stream will remain open after all changes have
    been sent and events that occur following the last changes will be streamed
    to the client until the stream is closed. If "withoutReplay" is true then
    only changes that happen after the call will be returned
    """

    # subscribe indicates whether to subscribe to events (e.g. ADD, UPDATE, and
    # REMOVE) that occur after all devices have been streamed to the client
    subscribe: bool = betterproto.bool_field(1)
    # option to specify a specific device change - if blank or '*' then select
    # all Can support `*` (match many chars) or '?' (match one char) as wildcard
    device_id: str = betterproto.string_field(2)
    # device_version is an optional device version
    device_version: str = betterproto.string_field(3)
    # option to request only changes that happen after the call
    without_replay: bool = betterproto.bool_field(4)


@dataclass(eq=False, repr=False)
class ListDeviceChangeResponse(betterproto.Message):
    """ListDeviceChangeResponse carries a single network change event"""

    # change is the device change on which the event occurred
    change: "_change_device__.DeviceChange" = betterproto.message_field(1)
    # type is a qualification of the type of change being made
    type: "Type" = betterproto.enum_field(2)


class ChangeServiceStub(betterproto.ServiceStub):
    async def list_network_changes(
        self,
        *,
        subscribe: bool = False,
        changeid: str = "",
        without_replay: bool = False,
    ) -> AsyncIterator["ListNetworkChangeResponse"]:

        request = ListNetworkChangeRequest()
        request.subscribe = subscribe
        request.changeid = changeid
        request.without_replay = without_replay

        async for response in self._unary_stream(
            "/onos.config.diags.ChangeService/ListNetworkChanges",
            request,
            ListNetworkChangeResponse,
        ):
            yield response

    async def list_device_changes(
        self,
        *,
        subscribe: bool = False,
        device_id: str = "",
        device_version: str = "",
        without_replay: bool = False,
    ) -> AsyncIterator["ListDeviceChangeResponse"]:

        request = ListDeviceChangeRequest()
        request.subscribe = subscribe
        request.device_id = device_id
        request.device_version = device_version
        request.without_replay = without_replay

        async for response in self._unary_stream(
            "/onos.config.diags.ChangeService/ListDeviceChanges",
            request,
            ListDeviceChangeResponse,
        ):
            yield response


class OpStateDiagsStub(betterproto.ServiceStub):
    async def get_op_state(
        self, *, device_id: str = "", subscribe: bool = False
    ) -> AsyncIterator["OpStateResponse"]:

        request = OpStateRequest()
        request.device_id = device_id
        request.subscribe = subscribe

        async for response in self._unary_stream(
            "/onos.config.diags.OpStateDiags/GetOpState",
            request,
            OpStateResponse,
        ):
            yield response


class ChangeServiceBase(ServiceBase):
    async def list_network_changes(
        self, subscribe: bool, changeid: str, without_replay: bool
    ) -> AsyncIterator["ListNetworkChangeResponse"]:
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def list_device_changes(
        self, subscribe: bool, device_id: str, device_version: str, without_replay: bool
    ) -> AsyncIterator["ListDeviceChangeResponse"]:
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def __rpc_list_network_changes(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "subscribe": request.subscribe,
            "changeid": request.changeid,
            "without_replay": request.without_replay,
        }

        await self._call_rpc_handler_server_stream(
            self.list_network_changes,
            stream,
            request_kwargs,
        )

    async def __rpc_list_device_changes(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "subscribe": request.subscribe,
            "device_id": request.device_id,
            "device_version": request.device_version,
            "without_replay": request.without_replay,
        }

        await self._call_rpc_handler_server_stream(
            self.list_device_changes,
            stream,
            request_kwargs,
        )

    def __mapping__(self) -> Dict[str, grpclib.const.Handler]:
        return {
            "/onos.config.diags.ChangeService/ListNetworkChanges": grpclib.const.Handler(
                self.__rpc_list_network_changes,
                grpclib.const.Cardinality.UNARY_STREAM,
                ListNetworkChangeRequest,
                ListNetworkChangeResponse,
            ),
            "/onos.config.diags.ChangeService/ListDeviceChanges": grpclib.const.Handler(
                self.__rpc_list_device_changes,
                grpclib.const.Cardinality.UNARY_STREAM,
                ListDeviceChangeRequest,
                ListDeviceChangeResponse,
            ),
        }


class OpStateDiagsBase(ServiceBase):
    async def get_op_state(
        self, device_id: str, subscribe: bool
    ) -> AsyncIterator["OpStateResponse"]:
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def __rpc_get_op_state(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "device_id": request.device_id,
            "subscribe": request.subscribe,
        }

        await self._call_rpc_handler_server_stream(
            self.get_op_state,
            stream,
            request_kwargs,
        )

    def __mapping__(self) -> Dict[str, grpclib.const.Handler]:
        return {
            "/onos.config.diags.OpStateDiags/GetOpState": grpclib.const.Handler(
                self.__rpc_get_op_state,
                grpclib.const.Cardinality.UNARY_STREAM,
                OpStateRequest,
                OpStateResponse,
            ),
        }


from .. import admin as _admin__
from ..change import device as _change_device__
from ..change import network as _change_network__
