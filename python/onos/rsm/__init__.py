# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: onos/rsm/rsm.proto
# plugin: python-betterproto
from dataclasses import dataclass
from typing import Dict, List, Optional

import betterproto
from betterproto.grpc.grpclib_server import ServiceBase
import grpclib


class SliceType(betterproto.Enum):
    SLICE_TYPE_DL_SLICE = 0
    SLICE_TYPE_UL_SLICE = 1


class SchedulerType(betterproto.Enum):
    SCHEDULER_TYPE_ROUND_ROBIN = 0
    SCHEDULER_TYPE_PROPORTIONALLY_FAIR = 1
    SCHEDULER_TYPE_QOS_BASED = 2


class UeIdType(betterproto.Enum):
    UE_ID_TYPE_CU_UE_F1_AP_ID = 0
    UE_ID_TYPE_DU_UE_F1_AP_ID = 1
    UE_ID_TYPE_RAN_UE_NGAP_ID = 2
    UE_ID_TYPE_AMF_UE_NGAP_ID = 3
    UE_ID_TYPE_ENB_UE_S1_AP_ID = 4


@dataclass(eq=False, repr=False)
class SliceItem(betterproto.Message):
    e2_node_id: str = betterproto.string_field(1)
    slice_ids: List[str] = betterproto.string_field(2)


@dataclass(eq=False, repr=False)
class Ack(betterproto.Message):
    success: bool = betterproto.bool_field(1)
    cause: str = betterproto.string_field(2)


@dataclass(eq=False, repr=False)
class CreateSliceRequest(betterproto.Message):
    e2_node_id: str = betterproto.string_field(1)
    slice_id: str = betterproto.string_field(2)
    scheduler_type: "SchedulerType" = betterproto.enum_field(3)
    weight: str = betterproto.string_field(4)
    slice_type: "SliceType" = betterproto.enum_field(5)


@dataclass(eq=False, repr=False)
class CreateSliceResponse(betterproto.Message):
    ack: "Ack" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class UpdateSliceRequest(betterproto.Message):
    e2_node_id: str = betterproto.string_field(1)
    slice_id: str = betterproto.string_field(2)
    scheduler_type: "SchedulerType" = betterproto.enum_field(3)
    weight: str = betterproto.string_field(4)
    slice_type: "SliceType" = betterproto.enum_field(5)


@dataclass(eq=False, repr=False)
class UpdateSliceResponse(betterproto.Message):
    ack: "Ack" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class DeleteSliceRequest(betterproto.Message):
    e2_node_id: str = betterproto.string_field(1)
    slice_id: str = betterproto.string_field(2)
    slice_type: "SliceType" = betterproto.enum_field(3)


@dataclass(eq=False, repr=False)
class DeleteSliceResponse(betterproto.Message):
    ack: "Ack" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class SliceAssocItem(betterproto.Message):
    ue_slice_assoc_id: str = betterproto.string_field(1)
    e2_node_id: str = betterproto.string_field(2)
    ue_id: List["UeIdType"] = betterproto.enum_field(3)
    slice_id: str = betterproto.string_field(4)


@dataclass(eq=False, repr=False)
class UeId(betterproto.Message):
    ue_id: str = betterproto.string_field(1)
    type: "UeIdType" = betterproto.enum_field(2)


@dataclass(eq=False, repr=False)
class SetUeSliceAssociationRequest(betterproto.Message):
    e2_node_id: str = betterproto.string_field(1)
    ue_id: List["UeId"] = betterproto.message_field(2)
    dl_slice_id: str = betterproto.string_field(3)
    ul_slice_id: str = betterproto.string_field(4)
    drb_id: str = betterproto.string_field(5)


@dataclass(eq=False, repr=False)
class SetUeSliceAssociationResponse(betterproto.Message):
    ack: "Ack" = betterproto.message_field(1)
    assigned_ue_slice_assoc_id: str = betterproto.string_field(2)


class RsmStub(betterproto.ServiceStub):
    async def create_slice(
        self,
        *,
        e2_node_id: str = "",
        slice_id: str = "",
        scheduler_type: "SchedulerType" = None,
        weight: str = "",
        slice_type: "SliceType" = None,
    ) -> "CreateSliceResponse":

        request = CreateSliceRequest()
        request.e2_node_id = e2_node_id
        request.slice_id = slice_id
        request.scheduler_type = scheduler_type
        request.weight = weight
        request.slice_type = slice_type

        return await self._unary_unary(
            "/onos.rsm.Rsm/CreateSlice", request, CreateSliceResponse
        )

    async def update_slice(
        self,
        *,
        e2_node_id: str = "",
        slice_id: str = "",
        scheduler_type: "SchedulerType" = None,
        weight: str = "",
        slice_type: "SliceType" = None,
    ) -> "UpdateSliceResponse":

        request = UpdateSliceRequest()
        request.e2_node_id = e2_node_id
        request.slice_id = slice_id
        request.scheduler_type = scheduler_type
        request.weight = weight
        request.slice_type = slice_type

        return await self._unary_unary(
            "/onos.rsm.Rsm/UpdateSlice", request, UpdateSliceResponse
        )

    async def delete_slice(
        self,
        *,
        e2_node_id: str = "",
        slice_id: str = "",
        slice_type: "SliceType" = None,
    ) -> "DeleteSliceResponse":

        request = DeleteSliceRequest()
        request.e2_node_id = e2_node_id
        request.slice_id = slice_id
        request.slice_type = slice_type

        return await self._unary_unary(
            "/onos.rsm.Rsm/DeleteSlice", request, DeleteSliceResponse
        )

    async def set_ue_slice_association(
        self,
        *,
        e2_node_id: str = "",
        ue_id: Optional[List["UeId"]] = None,
        dl_slice_id: str = "",
        ul_slice_id: str = "",
        drb_id: str = "",
    ) -> "SetUeSliceAssociationResponse":
        ue_id = ue_id or []

        request = SetUeSliceAssociationRequest()
        request.e2_node_id = e2_node_id
        if ue_id is not None:
            request.ue_id = ue_id
        request.dl_slice_id = dl_slice_id
        request.ul_slice_id = ul_slice_id
        request.drb_id = drb_id

        return await self._unary_unary(
            "/onos.rsm.Rsm/SetUeSliceAssociation",
            request,
            SetUeSliceAssociationResponse,
        )


class RsmBase(ServiceBase):
    async def create_slice(
        self,
        e2_node_id: str,
        slice_id: str,
        scheduler_type: "SchedulerType",
        weight: str,
        slice_type: "SliceType",
    ) -> "CreateSliceResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def update_slice(
        self,
        e2_node_id: str,
        slice_id: str,
        scheduler_type: "SchedulerType",
        weight: str,
        slice_type: "SliceType",
    ) -> "UpdateSliceResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def delete_slice(
        self, e2_node_id: str, slice_id: str, slice_type: "SliceType"
    ) -> "DeleteSliceResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def set_ue_slice_association(
        self,
        e2_node_id: str,
        ue_id: Optional[List["UeId"]],
        dl_slice_id: str,
        ul_slice_id: str,
        drb_id: str,
    ) -> "SetUeSliceAssociationResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def __rpc_create_slice(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "e2_node_id": request.e2_node_id,
            "slice_id": request.slice_id,
            "scheduler_type": request.scheduler_type,
            "weight": request.weight,
            "slice_type": request.slice_type,
        }

        response = await self.create_slice(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_update_slice(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "e2_node_id": request.e2_node_id,
            "slice_id": request.slice_id,
            "scheduler_type": request.scheduler_type,
            "weight": request.weight,
            "slice_type": request.slice_type,
        }

        response = await self.update_slice(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_delete_slice(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "e2_node_id": request.e2_node_id,
            "slice_id": request.slice_id,
            "slice_type": request.slice_type,
        }

        response = await self.delete_slice(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_set_ue_slice_association(
        self, stream: grpclib.server.Stream
    ) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "e2_node_id": request.e2_node_id,
            "ue_id": request.ue_id,
            "dl_slice_id": request.dl_slice_id,
            "ul_slice_id": request.ul_slice_id,
            "drb_id": request.drb_id,
        }

        response = await self.set_ue_slice_association(**request_kwargs)
        await stream.send_message(response)

    def __mapping__(self) -> Dict[str, grpclib.const.Handler]:
        return {
            "/onos.rsm.Rsm/CreateSlice": grpclib.const.Handler(
                self.__rpc_create_slice,
                grpclib.const.Cardinality.UNARY_UNARY,
                CreateSliceRequest,
                CreateSliceResponse,
            ),
            "/onos.rsm.Rsm/UpdateSlice": grpclib.const.Handler(
                self.__rpc_update_slice,
                grpclib.const.Cardinality.UNARY_UNARY,
                UpdateSliceRequest,
                UpdateSliceResponse,
            ),
            "/onos.rsm.Rsm/DeleteSlice": grpclib.const.Handler(
                self.__rpc_delete_slice,
                grpclib.const.Cardinality.UNARY_UNARY,
                DeleteSliceRequest,
                DeleteSliceResponse,
            ),
            "/onos.rsm.Rsm/SetUeSliceAssociation": grpclib.const.Handler(
                self.__rpc_set_ue_slice_association,
                grpclib.const.Cardinality.UNARY_UNARY,
                SetUeSliceAssociationRequest,
                SetUeSliceAssociationResponse,
            ),
        }
