# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: onos/fabricsim/devices.proto, onos/fabricsim/hosts.proto, onos/fabricsim/links.proto
# plugin: python-betterproto
from dataclasses import dataclass
from typing import List, Optional

import betterproto
import grpclib


class DeviceType(betterproto.Enum):
    """DeviceType represents type of a device, i.e. switch, IPU, etc."""

    SWITCH = 0
    IPU = 1


class StopMode(betterproto.Enum):
    """
    StopMode indicates whether to simulate orderly (administrative) or chaotic
    (power off) shutdown
    """

    ORDERLY_STOP = 0
    CHAOTIC_STOP = 1


class LinkStatus(betterproto.Enum):
    """DeviceType represents type of a device, i.e. switch, IPU, etc."""

    LINK_DOWN = 0
    LINK_UP = 1


@dataclass(eq=False, repr=False)
class Device(betterproto.Message):
    """Device describes a simulated switch or IPU"""

    # unique device id and device type
    id: str = betterproto.string_field(1)
    type: "DeviceType" = betterproto.enum_field(2)
    # list of ports
    ports: List["Port"] = betterproto.message_field(3)
    # control port for p4 and gnmi simulation
    control_port: int = betterproto.int32_field(4)
    # unique chassis ID
    chassis_id: int = betterproto.uint64_field(5)
    # forwarding pipeline information
    pipeline_info: "PipelineInfo" = betterproto.message_field(6)
    pos: "GridPosition" = betterproto.message_field(7)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class Port(betterproto.Message):
    """Port describes a simulated device port"""

    # unique port id and port type
    id: str = betterproto.string_field(1)
    # display/friendly name
    name: str = betterproto.string_field(3)
    # port number
    number: int = betterproto.uint32_field(4)
    # sdn/internal port number
    internal_number: int = betterproto.uint32_field(5)
    # speed
    speed: str = betterproto.string_field(6)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class PipelineInfo(betterproto.Message):
    """
    PipelineInfo provides information about the currently deployed forwarding
    pipeline
    """

    cookie: int = betterproto.uint64_field(1)
    p4_info: bytes = betterproto.bytes_field(2)
    # summary information about tables, counters and meters
    tables: List["EntitiesInfo"] = betterproto.message_field(3)
    counters: List["EntitiesInfo"] = betterproto.message_field(4)
    meters: List["EntitiesInfo"] = betterproto.message_field(5)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class EntitiesInfo(betterproto.Message):
    """
    EntitiesInfo provides information about size of pipeline entities, tables,
    meters, counters
    """

    id: int = betterproto.uint32_field(1)
    size: int = betterproto.uint32_field(2)
    name: str = betterproto.string_field(3)

    def __post_init__(self) -> None:
        super().__post_init__()


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
class GetDevicesRequest(betterproto.Message):
    pass

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class GetDevicesResponse(betterproto.Message):
    devices: List["Device"] = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class GetDeviceRequest(betterproto.Message):
    id: str = betterproto.string_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class GetDeviceResponse(betterproto.Message):
    device: "Device" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class AddDeviceRequest(betterproto.Message):
    device: "Device" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class AddDeviceResponse(betterproto.Message):
    pass

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class RemoveDeviceRequest(betterproto.Message):
    id: str = betterproto.string_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class RemoveDeviceResponse(betterproto.Message):
    pass

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class StopDeviceRequest(betterproto.Message):
    id: str = betterproto.string_field(1)
    mode: "StopMode" = betterproto.enum_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class StopDeviceResponse(betterproto.Message):
    pass

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class StartDeviceRequest(betterproto.Message):
    id: str = betterproto.string_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class StartDeviceResponse(betterproto.Message):
    pass

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class DisablePortRequest(betterproto.Message):
    id: str = betterproto.string_field(1)
    mode: "StopMode" = betterproto.enum_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class DisablePortResponse(betterproto.Message):
    pass

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class EnablePortRequest(betterproto.Message):
    id: str = betterproto.string_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class EnablePortResponse(betterproto.Message):
    pass

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class Host(betterproto.Message):
    """Device describes a simulated switch or IPU"""

    # unique device id and device type
    id: str = betterproto.string_field(1)
    interfaces: List["NetworkInterface"] = betterproto.message_field(2)
    # list of ports
    pos: "GridPosition" = betterproto.message_field(3)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class NetworkInterface(betterproto.Message):
    """Port describes a simulated device port"""

    # unique port id and port type
    id: str = betterproto.string_field(1)
    # display/friendly name
    mac_address: str = betterproto.string_field(2)
    # port number
    ip_address: str = betterproto.string_field(3)
    # sdn/internal port number
    ipv6_address: str = betterproto.string_field(4)
    # speed
    behavior: "NetworkInterfaceBehavior" = betterproto.message_field(5)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class NetworkInterfaceBehavior(betterproto.Message):
    """
    PipelineInfo provides information about the currently deployed forwarding
    pipeline
    """

    pass

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class GetHostsRequest(betterproto.Message):
    """
    EntitiesInfo provides information about size of pipeline entities, tables,
    meters, counters
    """

    pass

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class GetHostsResponse(betterproto.Message):
    """
    GridPosition indicates where on a grid an entity should be located; used
    for visualization purposes
    """

    hosts: List["Host"] = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class GetHostRequest(betterproto.Message):
    id: str = betterproto.string_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class GetHostResponse(betterproto.Message):
    host: "Host" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class AddHostRequest(betterproto.Message):
    host: "Host" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class AddHostResponse(betterproto.Message):
    pass

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class RemoveHostRequest(betterproto.Message):
    id: str = betterproto.string_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class RemoveHostResponse(betterproto.Message):
    pass

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class Link(betterproto.Message):
    """Device describes a simulated switch or IPU"""

    # unique device id and device type
    id: str = betterproto.string_field(1)
    src_id: str = betterproto.string_field(2)
    # list of ports
    tgt_id: str = betterproto.string_field(3)
    # control port for p4 and gnmi simulation
    status: "LinkStatus" = betterproto.enum_field(4)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class GetLinksRequest(betterproto.Message):
    """Port describes a simulated device port"""

    pass

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class GetLinksResponse(betterproto.Message):
    """
    PipelineInfo provides information about the currently deployed forwarding
    pipeline
    """

    links: List["Link"] = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class GetLinkRequest(betterproto.Message):
    """
    EntitiesInfo provides information about size of pipeline entities, tables,
    meters, counters
    """

    id: str = betterproto.string_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class GetLinkResponse(betterproto.Message):
    """
    GridPosition indicates where on a grid an entity should be located; used
    for visualization purposes
    """

    link: "Link" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class AddLinkRequest(betterproto.Message):
    link: "Link" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class AddLinkResponse(betterproto.Message):
    pass

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class RemoveLinkRequest(betterproto.Message):
    id: str = betterproto.string_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class RemoveLinkResponse(betterproto.Message):
    pass

    def __post_init__(self) -> None:
        super().__post_init__()


class DeviceServiceStub(betterproto.ServiceStub):
    """
    DeviceService provides means to control inventory of simulated devices
    (switches and IPUs) and their ports
    """

    async def get_devices(self) -> "GetDevicesResponse":
        """
        GetDevices gets a list of all simulated devices (switches and/or IPUs)
        """

        request = GetDevicesRequest()

        return await self._unary_unary(
            "/onos.fabricsim.DeviceService/GetDevices", request, GetDevicesResponse
        )

    async def get_device(self, *, id: str = "") -> "GetDeviceResponse":
        """GetDevice gets a specific device entry"""

        request = GetDeviceRequest()
        request.id = id

        return await self._unary_unary(
            "/onos.fabricsim.DeviceService/GetDevice", request, GetDeviceResponse
        )

    async def add_device(self, *, device: "Device" = None) -> "AddDeviceResponse":
        """
        AddDevice creates a new simulated deviceand start its P4Runtime and
        gNMI services
        """

        request = AddDeviceRequest()
        if device is not None:
            request.device = device

        return await self._unary_unary(
            "/onos.fabricsim.DeviceService/AddDevice", request, AddDeviceResponse
        )

    async def remove_device(self, *, id: str = "") -> "RemoveDeviceResponse":
        """RemoveDevice removes a simulated device"""

        request = RemoveDeviceRequest()
        request.id = id

        return await self._unary_unary(
            "/onos.fabricsim.DeviceService/RemoveDevice", request, RemoveDeviceResponse
        )

    async def stop_device(
        self, *, id: str = "", mode: "StopMode" = None
    ) -> "StopDeviceResponse":
        """StopDevice stops the simulated deviceP4Runtime and gNMI services"""

        request = StopDeviceRequest()
        request.id = id
        request.mode = mode

        return await self._unary_unary(
            "/onos.fabricsim.DeviceService/StopDevice", request, StopDeviceResponse
        )

    async def start_device(self, *, id: str = "") -> "StartDeviceResponse":
        """
        StartDevice starts the simulated deviceP4Runtime and gNMI services
        """

        request = StartDeviceRequest()
        request.id = id

        return await self._unary_unary(
            "/onos.fabricsim.DeviceService/StartDevice", request, StartDeviceResponse
        )

    async def disable_port(
        self, *, id: str = "", mode: "StopMode" = None
    ) -> "DisablePortResponse":
        """DisablePort disables the specified port"""

        request = DisablePortRequest()
        request.id = id
        request.mode = mode

        return await self._unary_unary(
            "/onos.fabricsim.DeviceService/DisablePort", request, DisablePortResponse
        )

    async def enable_port(self, *, id: str = "") -> "EnablePortResponse":
        """EnablePort enables the specified port"""

        request = EnablePortRequest()
        request.id = id

        return await self._unary_unary(
            "/onos.fabricsim.DeviceService/EnablePort", request, EnablePortResponse
        )


class HostServiceStub(betterproto.ServiceStub):
    """
    DeviceService provides means to control inventory of simulated devices
    (switches and IPUs) and their ports
    """

    async def get_hosts(self) -> "GetHostsResponse":
        """
        GetDevices gets a list of all simulated devices (switches and/or IPUs)
        """

        request = GetHostsRequest()

        return await self._unary_unary(
            "/onos.fabricsim.HostService/GetHosts", request, GetHostsResponse
        )

    async def get_host(self, *, id: str = "") -> "GetHostResponse":
        """GetDevice gets a specific device entry"""

        request = GetHostRequest()
        request.id = id

        return await self._unary_unary(
            "/onos.fabricsim.HostService/GetHost", request, GetHostResponse
        )

    async def add_host(self, *, host: "Host" = None) -> "AddHostResponse":
        """
        AddDevice creates a new simulated deviceand start its P4Runtime and
        gNMI services
        """

        request = AddHostRequest()
        if host is not None:
            request.host = host

        return await self._unary_unary(
            "/onos.fabricsim.HostService/AddHost", request, AddHostResponse
        )

    async def remove_host(self, *, id: str = "") -> "RemoveHostResponse":
        """RemoveDevice removes a simulated device"""

        request = RemoveHostRequest()
        request.id = id

        return await self._unary_unary(
            "/onos.fabricsim.HostService/RemoveHost", request, RemoveHostResponse
        )


class LinkServiceStub(betterproto.ServiceStub):
    """
    DeviceService provides means to control inventory of simulated devices
    (switches and IPUs) and their ports
    """

    async def get_links(self) -> "GetLinksResponse":
        """
        GetDevices gets a list of all simulated devices (switches and/or IPUs)
        """

        request = GetLinksRequest()

        return await self._unary_unary(
            "/onos.fabricsim.LinkService/GetLinks", request, GetLinksResponse
        )

    async def get_link(self, *, id: str = "") -> "GetLinkResponse":
        """GetDevice gets a specific device entry"""

        request = GetLinkRequest()
        request.id = id

        return await self._unary_unary(
            "/onos.fabricsim.LinkService/GetLink", request, GetLinkResponse
        )

    async def add_link(self, *, link: "Link" = None) -> "AddLinkResponse":
        """
        AddDevice creates a new simulated deviceand start its P4Runtime and
        gNMI services
        """

        request = AddLinkRequest()
        if link is not None:
            request.link = link

        return await self._unary_unary(
            "/onos.fabricsim.LinkService/AddLink", request, AddLinkResponse
        )

    async def remove_link(self, *, id: str = "") -> "RemoveLinkResponse":
        """RemoveDevice removes a simulated device"""

        request = RemoveLinkRequest()
        request.id = id

        return await self._unary_unary(
            "/onos.fabricsim.LinkService/RemoveLink", request, RemoveLinkResponse
        )
