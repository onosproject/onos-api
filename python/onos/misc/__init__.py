# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: onos/misc/misc.proto
# plugin: python-betterproto
from dataclasses import dataclass

import betterproto


@dataclass(eq=False, repr=False)
class GridPosition(betterproto.Message):
    """
    GridPosition indicates where on a grid an entity should be located; used
    for visualization purposes
    """

    x: int = betterproto.int32_field(1)
    y: int = betterproto.int32_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class Connection(betterproto.Message):
    """
    Connection contains information about a single gRPC client connection
    """

    from_address: str = betterproto.string_field(1)
    protocol: str = betterproto.string_field(2)
    time: int = betterproto.int64_field(3)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class IoStats(betterproto.Message):
    """IOStats represents I/O statistics for a single device agent"""

    in_bytes: int = betterproto.uint32_field(1)
    in_messages: int = betterproto.uint32_field(2)
    out_bytes: int = betterproto.uint32_field(3)
    out_messages: int = betterproto.uint32_field(4)
    first_update_time: int = betterproto.uint64_field(5)
    last_update_time: int = betterproto.uint64_field(6)

    def __post_init__(self) -> None:
        super().__post_init__()
