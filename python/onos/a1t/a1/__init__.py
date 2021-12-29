# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: onos/a1t/a1/a1.proto, onos/a1t/a1/ei.proto, onos/a1t/a1/policy.proto
# plugin: python-betterproto
from dataclasses import dataclass
from typing import AsyncIterable, AsyncIterator, Dict, Iterable, Union

import betterproto
from betterproto.grpc.grpclib_server import ServiceBase
import grpclib


class Encoding(betterproto.Enum):
    PROTO = 0
    JSON = 1


@dataclass(eq=False, repr=False)
class Header(betterproto.Message):
    app_id: str = betterproto.string_field(1)
    app_instance_id: str = betterproto.string_field(2)
    a1_node_id: str = betterproto.string_field(3)
    encoding: "Encoding" = betterproto.enum_field(4)


@dataclass(eq=False, repr=False)
class RequestMessage(betterproto.Message):
    header: bytes = betterproto.bytes_field(1)
    payload: bytes = betterproto.bytes_field(2)


@dataclass(eq=False, repr=False)
class ResultMessage(betterproto.Message):
    header: bytes = betterproto.bytes_field(1)
    payload: bytes = betterproto.bytes_field(2)


@dataclass(eq=False, repr=False)
class StatusMessage(betterproto.Message):
    header: bytes = betterproto.bytes_field(1)
    payload: bytes = betterproto.bytes_field(2)


@dataclass(eq=False, repr=False)
class AckMessage(betterproto.Message):
    header: bytes = betterproto.bytes_field(1)
    payload: bytes = betterproto.bytes_field(2)


@dataclass(eq=False, repr=False)
class EiRequestMessage(betterproto.Message):
    header: "Header" = betterproto.message_field(1)
    message: "RequestMessage" = betterproto.message_field(2)


@dataclass(eq=False, repr=False)
class EiResultMessage(betterproto.Message):
    header: "Header" = betterproto.message_field(1)
    message: "ResultMessage" = betterproto.message_field(2)


@dataclass(eq=False, repr=False)
class EiStatusMessage(betterproto.Message):
    header: "Header" = betterproto.message_field(1)
    message: "StatusMessage" = betterproto.message_field(2)


@dataclass(eq=False, repr=False)
class EiAckMessage(betterproto.Message):
    header: "Header" = betterproto.message_field(1)
    message: "AckMessage" = betterproto.message_field(2)


@dataclass(eq=False, repr=False)
class PolicyType(betterproto.Message):
    id: str = betterproto.string_field(1)
    name: str = betterproto.string_field(2)
    version: str = betterproto.string_field(3)


@dataclass(eq=False, repr=False)
class PolicyRequestMessage(betterproto.Message):
    header: "Header" = betterproto.message_field(1)
    policy_type: "PolicyType" = betterproto.message_field(2)
    message: "RequestMessage" = betterproto.message_field(3)


@dataclass(eq=False, repr=False)
class PolicyResultMessage(betterproto.Message):
    header: "Header" = betterproto.message_field(1)
    policy_type: "PolicyType" = betterproto.message_field(2)
    message: "ResultMessage" = betterproto.message_field(3)


@dataclass(eq=False, repr=False)
class PolicyStatusMessage(betterproto.Message):
    header: "Header" = betterproto.message_field(1)
    policy_type: "PolicyType" = betterproto.message_field(2)
    message: "StatusMessage" = betterproto.message_field(3)


@dataclass(eq=False, repr=False)
class PolicyAckMessage(betterproto.Message):
    header: "Header" = betterproto.message_field(1)
    policy_type: "PolicyType" = betterproto.message_field(2)
    message: "AckMessage" = betterproto.message_field(3)


class EiServiceStub(betterproto.ServiceStub):
    async def ei_query(
        self,
        request_iterator: Union[
            AsyncIterable["EiResultMessage"], Iterable["EiResultMessage"]
        ],
    ) -> AsyncIterator["EiRequestMessage"]:

        async for response in self._stream_stream(
            "/onos.a1t.a1.EIService/EIQuery",
            request_iterator,
            EiResultMessage,
            EiRequestMessage,
        ):
            yield response

    async def ei_job_setup(
        self,
        request_iterator: Union[
            AsyncIterable["EiResultMessage"], Iterable["EiResultMessage"]
        ],
    ) -> AsyncIterator["EiRequestMessage"]:

        async for response in self._stream_stream(
            "/onos.a1t.a1.EIService/EIJobSetup",
            request_iterator,
            EiResultMessage,
            EiRequestMessage,
        ):
            yield response

    async def ei_job_update(
        self,
        request_iterator: Union[
            AsyncIterable["EiResultMessage"], Iterable["EiResultMessage"]
        ],
    ) -> AsyncIterator["EiRequestMessage"]:

        async for response in self._stream_stream(
            "/onos.a1t.a1.EIService/EIJobUpdate",
            request_iterator,
            EiResultMessage,
            EiRequestMessage,
        ):
            yield response

    async def ei_job_delete(
        self,
        request_iterator: Union[
            AsyncIterable["EiResultMessage"], Iterable["EiResultMessage"]
        ],
    ) -> AsyncIterator["EiRequestMessage"]:

        async for response in self._stream_stream(
            "/onos.a1t.a1.EIService/EIJobDelete",
            request_iterator,
            EiResultMessage,
            EiRequestMessage,
        ):
            yield response

    async def ei_job_status_query(
        self,
        request_iterator: Union[
            AsyncIterable["EiResultMessage"], Iterable["EiResultMessage"]
        ],
    ) -> AsyncIterator["EiRequestMessage"]:

        async for response in self._stream_stream(
            "/onos.a1t.a1.EIService/EIJobStatusQuery",
            request_iterator,
            EiResultMessage,
            EiRequestMessage,
        ):
            yield response

    async def ei_job_status_notify(self) -> "EiAckMessage":

        request = EiStatusMessage()

        return await self._unary_unary(
            "/onos.a1t.a1.EIService/EIJobStatusNotify", request, EiAckMessage
        )

    async def ei_job_result_delivery(self) -> "EiAckMessage":

        request = EiResultMessage()

        return await self._unary_unary(
            "/onos.a1t.a1.EIService/EIJobResultDelivery", request, EiAckMessage
        )


class PolicyServiceStub(betterproto.ServiceStub):
    async def policy_setup(
        self,
        *,
        header: "Header" = None,
        policy_type: "PolicyType" = None,
        message: "RequestMessage" = None,
    ) -> "PolicyResultMessage":

        request = PolicyRequestMessage()
        if header is not None:
            request.header = header
        if policy_type is not None:
            request.policy_type = policy_type
        if message is not None:
            request.message = message

        return await self._unary_unary(
            "/onos.a1t.a1.PolicyService/PolicySetup", request, PolicyResultMessage
        )

    async def policy_update(
        self,
        *,
        header: "Header" = None,
        policy_type: "PolicyType" = None,
        message: "RequestMessage" = None,
    ) -> "PolicyResultMessage":

        request = PolicyRequestMessage()
        if header is not None:
            request.header = header
        if policy_type is not None:
            request.policy_type = policy_type
        if message is not None:
            request.message = message

        return await self._unary_unary(
            "/onos.a1t.a1.PolicyService/PolicyUpdate", request, PolicyResultMessage
        )

    async def policy_delete(
        self,
        *,
        header: "Header" = None,
        policy_type: "PolicyType" = None,
        message: "RequestMessage" = None,
    ) -> "PolicyResultMessage":

        request = PolicyRequestMessage()
        if header is not None:
            request.header = header
        if policy_type is not None:
            request.policy_type = policy_type
        if message is not None:
            request.message = message

        return await self._unary_unary(
            "/onos.a1t.a1.PolicyService/PolicyDelete", request, PolicyResultMessage
        )

    async def policy_query(
        self,
        *,
        header: "Header" = None,
        policy_type: "PolicyType" = None,
        message: "RequestMessage" = None,
    ) -> "PolicyResultMessage":

        request = PolicyRequestMessage()
        if header is not None:
            request.header = header
        if policy_type is not None:
            request.policy_type = policy_type
        if message is not None:
            request.message = message

        return await self._unary_unary(
            "/onos.a1t.a1.PolicyService/PolicyQuery", request, PolicyResultMessage
        )

    async def policy_status(
        self,
        request_iterator: Union[
            AsyncIterable["PolicyAckMessage"], Iterable["PolicyAckMessage"]
        ],
    ) -> AsyncIterator["PolicyStatusMessage"]:

        async for response in self._stream_stream(
            "/onos.a1t.a1.PolicyService/PolicyStatus",
            request_iterator,
            PolicyAckMessage,
            PolicyStatusMessage,
        ):
            yield response


class EiServiceBase(ServiceBase):
    async def ei_query(
        self, request_iterator: AsyncIterator["EiResultMessage"]
    ) -> AsyncIterator["EiRequestMessage"]:
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def ei_job_setup(
        self, request_iterator: AsyncIterator["EiResultMessage"]
    ) -> AsyncIterator["EiRequestMessage"]:
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def ei_job_update(
        self, request_iterator: AsyncIterator["EiResultMessage"]
    ) -> AsyncIterator["EiRequestMessage"]:
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def ei_job_delete(
        self, request_iterator: AsyncIterator["EiResultMessage"]
    ) -> AsyncIterator["EiRequestMessage"]:
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def ei_job_status_query(
        self, request_iterator: AsyncIterator["EiResultMessage"]
    ) -> AsyncIterator["EiRequestMessage"]:
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def ei_job_status_notify(self) -> "EiAckMessage":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def ei_job_result_delivery(self) -> "EiAckMessage":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def __rpc_ei_query(self, stream: grpclib.server.Stream) -> None:
        request_kwargs = {"request_iterator": stream.__aiter__()}

        await self._call_rpc_handler_server_stream(
            self.ei_query,
            stream,
            request_kwargs,
        )

    async def __rpc_ei_job_setup(self, stream: grpclib.server.Stream) -> None:
        request_kwargs = {"request_iterator": stream.__aiter__()}

        await self._call_rpc_handler_server_stream(
            self.ei_job_setup,
            stream,
            request_kwargs,
        )

    async def __rpc_ei_job_update(self, stream: grpclib.server.Stream) -> None:
        request_kwargs = {"request_iterator": stream.__aiter__()}

        await self._call_rpc_handler_server_stream(
            self.ei_job_update,
            stream,
            request_kwargs,
        )

    async def __rpc_ei_job_delete(self, stream: grpclib.server.Stream) -> None:
        request_kwargs = {"request_iterator": stream.__aiter__()}

        await self._call_rpc_handler_server_stream(
            self.ei_job_delete,
            stream,
            request_kwargs,
        )

    async def __rpc_ei_job_status_query(self, stream: grpclib.server.Stream) -> None:
        request_kwargs = {"request_iterator": stream.__aiter__()}

        await self._call_rpc_handler_server_stream(
            self.ei_job_status_query,
            stream,
            request_kwargs,
        )

    async def __rpc_ei_job_status_notify(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {}

        response = await self.ei_job_status_notify(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_ei_job_result_delivery(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {}

        response = await self.ei_job_result_delivery(**request_kwargs)
        await stream.send_message(response)

    def __mapping__(self) -> Dict[str, grpclib.const.Handler]:
        return {
            "/onos.a1t.a1.EIService/EIQuery": grpclib.const.Handler(
                self.__rpc_ei_query,
                grpclib.const.Cardinality.STREAM_STREAM,
                EiResultMessage,
                EiRequestMessage,
            ),
            "/onos.a1t.a1.EIService/EIJobSetup": grpclib.const.Handler(
                self.__rpc_ei_job_setup,
                grpclib.const.Cardinality.STREAM_STREAM,
                EiResultMessage,
                EiRequestMessage,
            ),
            "/onos.a1t.a1.EIService/EIJobUpdate": grpclib.const.Handler(
                self.__rpc_ei_job_update,
                grpclib.const.Cardinality.STREAM_STREAM,
                EiResultMessage,
                EiRequestMessage,
            ),
            "/onos.a1t.a1.EIService/EIJobDelete": grpclib.const.Handler(
                self.__rpc_ei_job_delete,
                grpclib.const.Cardinality.STREAM_STREAM,
                EiResultMessage,
                EiRequestMessage,
            ),
            "/onos.a1t.a1.EIService/EIJobStatusQuery": grpclib.const.Handler(
                self.__rpc_ei_job_status_query,
                grpclib.const.Cardinality.STREAM_STREAM,
                EiResultMessage,
                EiRequestMessage,
            ),
            "/onos.a1t.a1.EIService/EIJobStatusNotify": grpclib.const.Handler(
                self.__rpc_ei_job_status_notify,
                grpclib.const.Cardinality.UNARY_UNARY,
                EiStatusMessage,
                EiAckMessage,
            ),
            "/onos.a1t.a1.EIService/EIJobResultDelivery": grpclib.const.Handler(
                self.__rpc_ei_job_result_delivery,
                grpclib.const.Cardinality.UNARY_UNARY,
                EiResultMessage,
                EiAckMessage,
            ),
        }


class PolicyServiceBase(ServiceBase):
    async def policy_setup(
        self, header: "Header", policy_type: "PolicyType", message: "RequestMessage"
    ) -> "PolicyResultMessage":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def policy_update(
        self, header: "Header", policy_type: "PolicyType", message: "RequestMessage"
    ) -> "PolicyResultMessage":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def policy_delete(
        self, header: "Header", policy_type: "PolicyType", message: "RequestMessage"
    ) -> "PolicyResultMessage":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def policy_query(
        self, header: "Header", policy_type: "PolicyType", message: "RequestMessage"
    ) -> "PolicyResultMessage":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def policy_status(
        self, request_iterator: AsyncIterator["PolicyAckMessage"]
    ) -> AsyncIterator["PolicyStatusMessage"]:
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def __rpc_policy_setup(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "header": request.header,
            "policy_type": request.policy_type,
            "message": request.message,
        }

        response = await self.policy_setup(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_policy_update(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "header": request.header,
            "policy_type": request.policy_type,
            "message": request.message,
        }

        response = await self.policy_update(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_policy_delete(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "header": request.header,
            "policy_type": request.policy_type,
            "message": request.message,
        }

        response = await self.policy_delete(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_policy_query(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "header": request.header,
            "policy_type": request.policy_type,
            "message": request.message,
        }

        response = await self.policy_query(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_policy_status(self, stream: grpclib.server.Stream) -> None:
        request_kwargs = {"request_iterator": stream.__aiter__()}

        await self._call_rpc_handler_server_stream(
            self.policy_status,
            stream,
            request_kwargs,
        )

    def __mapping__(self) -> Dict[str, grpclib.const.Handler]:
        return {
            "/onos.a1t.a1.PolicyService/PolicySetup": grpclib.const.Handler(
                self.__rpc_policy_setup,
                grpclib.const.Cardinality.UNARY_UNARY,
                PolicyRequestMessage,
                PolicyResultMessage,
            ),
            "/onos.a1t.a1.PolicyService/PolicyUpdate": grpclib.const.Handler(
                self.__rpc_policy_update,
                grpclib.const.Cardinality.UNARY_UNARY,
                PolicyRequestMessage,
                PolicyResultMessage,
            ),
            "/onos.a1t.a1.PolicyService/PolicyDelete": grpclib.const.Handler(
                self.__rpc_policy_delete,
                grpclib.const.Cardinality.UNARY_UNARY,
                PolicyRequestMessage,
                PolicyResultMessage,
            ),
            "/onos.a1t.a1.PolicyService/PolicyQuery": grpclib.const.Handler(
                self.__rpc_policy_query,
                grpclib.const.Cardinality.UNARY_UNARY,
                PolicyRequestMessage,
                PolicyResultMessage,
            ),
            "/onos.a1t.a1.PolicyService/PolicyStatus": grpclib.const.Handler(
                self.__rpc_policy_status,
                grpclib.const.Cardinality.STREAM_STREAM,
                PolicyAckMessage,
                PolicyStatusMessage,
            ),
        }
