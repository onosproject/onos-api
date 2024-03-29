# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: onos/mho/mho.proto
# plugin: python-betterproto
from dataclasses import dataclass
from typing import List

import betterproto
import grpclib


@dataclass(eq=False, repr=False)
class GetRequest(betterproto.Message):
    pass

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class UeList(betterproto.Message):
    ues: List["Ue"] = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class CellList(betterproto.Message):
    cells: List["Cell"] = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class Ue(betterproto.Message):
    ue_id: str = betterproto.string_field(1)
    ho_state: str = betterproto.string_field(2)
    cgi: str = betterproto.string_field(3)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class Cell(betterproto.Message):
    cgi: str = betterproto.string_field(1)
    num_ues: int = betterproto.int64_field(2)
    cumulative_handovers_in: int = betterproto.int64_field(4)
    cumulative_handovers_out: int = betterproto.int64_field(5)

    def __post_init__(self) -> None:
        super().__post_init__()


class MhoStub(betterproto.ServiceStub):
    async def get_ues(self) -> "UeList":

        request = GetRequest()

        return await self._unary_unary("/onos.mho.mho/GetUes", request, UeList)

    async def get_cells(self) -> "CellList":

        request = GetRequest()

        return await self._unary_unary("/onos.mho.mho/GetCells", request, CellList)
