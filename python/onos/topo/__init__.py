# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: onos/topo/config.proto, onos/topo/ran.proto, onos/topo/topo.proto
# plugin: python-betterproto
import warnings
from dataclasses import dataclass
from datetime import datetime, timedelta
from typing import AsyncIterator, Dict, List

import betterproto
from betterproto.grpc.grpclib_server import ServiceBase
import grpclib


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


class RanEntityKinds(betterproto.Enum):
    """
    TODO: Deprecate! Use language specific constant definitions instead. Kinds
    of RAN entities
    """

    E2NODE = 0
    E2CELL = 1
    E2T = 3
    XAPP = 4
    A1T = 5


class RanRelationKinds(betterproto.Enum):
    """
    TODO: Deprecate! Use language specific constant definitions instead. Kinds
    of RAN relations
    """

    CONTROLS = 0
    CONTAINS = 1
    NEIGHBORS = 2


class CellGlobalIdType(betterproto.Enum):
    NRCGI = 0
    ECGI = 1


class NodeType(betterproto.Enum):
    """
    NodeType is for use in encoding node type in topo Entity URI identifier;
    see section 9.2.6 of ORAN E2AP v02.00 spec
    """

    NT_NONE = 0
    NT_GNB = 1
    NT_EN_GNB = 2
    NT_NG_ENB = 3
    NT_ENB = 4


class ComponentType(betterproto.Enum):
    """
    ComponentType is for use in encoding node component type in topo Entity URI
    identifier; see section 9.2.6 of ORAN E2AP v02.00 spec
    """

    CT_NONE = 0
    CT_CU = 1
    CT_CU_UP = 2
    CT_DU = 3
    CT_ENB = 4


class E2SmRsmCommand(betterproto.Enum):
    E2_SM_RSM_COMMAND_SLICE_CREATE = 0
    E2_SM_RSM_COMMAND_SLICE_UPDATE = 1
    E2_SM_RSM_COMMAND_SLICE_DELETE = 2
    E2_SM_RSM_COMMAND_UE_ASSOCIATE = 3
    E2_SM_RSM_COMMAND_EVENT_TRIGGERS = 4


class RsmSlicingType(betterproto.Enum):
    SLICING_TYPE_STATIC = 0
    SLICING_TYPE_DYNAMIC = 1


class RsmSchedulerType(betterproto.Enum):
    SCHEDULER_TYPE_ROUND_ROBIN = 0
    SCHEDULER_TYPE_PROPORTIONALLY_FAIR = 1
    SCHEDULER_TYPE_QOS_BASED = 2


class RsmSliceType(betterproto.Enum):
    SLICE_TYPE_DL_SLICE = 0
    SLICE_TYPE_UL_SLICE = 1


class UeIdType(betterproto.Enum):
    UE_ID_TYPE_CU_UE_F1_AP_ID = 0
    UE_ID_TYPE_DU_UE_F1_AP_ID = 1
    UE_ID_TYPE_RAN_UE_NGAP_ID = 2
    UE_ID_TYPE_AMF_UE_NGAP_ID = 3
    UE_ID_TYPE_ENB_UE_S1_AP_ID = 4


class InterfaceType(betterproto.Enum):
    INTERFACE_UNKNOWN = 0
    INTERFACE_E2T = 1
    INTERFACE_E2AP101 = 2
    INTERFACE_E2AP200 = 3
    INTERFACE_A1_XAPP = 4
    INTERFACE_A1AP = 5


class EventType(betterproto.Enum):
    """EventType is a topo operation event type"""

    NONE = 0
    ADDED = 1
    UPDATED = 2
    REMOVED = 3


class RelationFilterScope(betterproto.Enum):
    TARGETS_ONLY = 0
    ALL = 1
    SOURCE_AND_TARGETS = 2
    RELATIONS_ONLY = 3
    RELATIONS_AND_TARGETS = 4


class SortOrder(betterproto.Enum):
    UNORDERED = 0
    ASCENDING = 1
    DESCENDING = 2


class ObjectType(betterproto.Enum):
    UNSPECIFIED = 0
    ENTITY = 1
    RELATION = 2
    KIND = 3


@dataclass(eq=False, repr=False)
class Asset(betterproto.Message):
    """Basic asset information"""

    name: str = betterproto.string_field(1)
    manufacturer: str = betterproto.string_field(2)
    model: str = betterproto.string_field(3)
    serial: str = betterproto.string_field(4)
    asset: str = betterproto.string_field(5)
    sw_version: str = betterproto.string_field(6)
    role: str = betterproto.string_field(8)


@dataclass(eq=False, repr=False)
class Configurable(betterproto.Message):
    """Configurable device aspect"""

    type: str = betterproto.string_field(1)
    address: str = betterproto.string_field(2)
    target: str = betterproto.string_field(3)
    version: str = betterproto.string_field(4)
    timeout: timedelta = betterproto.message_field(5)
    persistent: bool = betterproto.bool_field(6)


@dataclass(eq=False, repr=False)
class MastershipState(betterproto.Message):
    """Aspect for tracking device mastership"""

    term: int = betterproto.uint64_field(1)
    node_id: str = betterproto.string_field(2)


@dataclass(eq=False, repr=False)
class TlsOptions(betterproto.Message):
    """TLS connectivity aspect"""

    insecure: bool = betterproto.bool_field(1)
    plain: bool = betterproto.bool_field(2)
    key: str = betterproto.string_field(3)
    ca_cert: str = betterproto.string_field(4)
    cert: str = betterproto.string_field(5)


@dataclass(eq=False, repr=False)
class AdHoc(betterproto.Message):
    """Aspect for ad-hoc properties"""

    properties: Dict[str, str] = betterproto.map_field(
        1, betterproto.TYPE_STRING, betterproto.TYPE_STRING
    )


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


@dataclass(eq=False, repr=False)
class Protocols(betterproto.Message):
    """Protocols"""

    state: List["ProtocolState"] = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class Location(betterproto.Message):
    # Legacy Geographical location; expected value type of "location" aspect
    lat: float = betterproto.double_field(1)
    lng: float = betterproto.double_field(2)
    wgs84: "Wgs84Location" = betterproto.message_field(3, group="ext")
    cartesian: "CartesianLocation" = betterproto.message_field(4, group="ext")

    def __post_init__(self) -> None:
        super().__post_init__()
        if self.lat:
            warnings.warn("Location.lat is deprecated", DeprecationWarning)
        if self.lng:
            warnings.warn("Location.lng is deprecated", DeprecationWarning)


@dataclass(eq=False, repr=False)
class Wgs84Location(betterproto.Message):
    # Latitude in degrees, between +/- 90.
    latitude_deg: float = betterproto.double_field(1)
    # Longitude in degrees, between +/- 180.
    longitude_deg: float = betterproto.double_field(2)
    # Altitude in meters, from the surface of the WGS-84 ellipsoid.
    altitude_m: float = betterproto.double_field(3)


@dataclass(eq=False, repr=False)
class CartesianLocation(betterproto.Message):
    """ECEF-compatible cartesian coordinates."""

    # X, Y, and Z coordinates in meters.
    x_m: float = betterproto.double_field(1)
    y_m: float = betterproto.double_field(2)
    z_m: float = betterproto.double_field(3)


@dataclass(eq=False, repr=False)
class AzElOrientation(betterproto.Message):
    # Azimuth angle in degrees.
    azimuth_deg: float = betterproto.double_field(1)
    # Elevation angle in degrees.
    elevation_deg: float = betterproto.double_field(2)
    # Rotation around boresight in degrees.
    rotation_deg: float = betterproto.double_field(3)


@dataclass(eq=False, repr=False)
class YprOrientation(betterproto.Message):
    # Yaw angle around z-axis.
    yaw_deg: float = betterproto.double_field(1)
    # Pitch angle around y-axix.
    pitch_deg: float = betterproto.double_field(2)
    # Roll angle around x-axis.
    roll_deg: float = betterproto.double_field(3)


@dataclass(eq=False, repr=False)
class Orientation(betterproto.Message):
    azel: "AzElOrientation" = betterproto.message_field(1, group="orientation")
    ypr: "YprOrientation" = betterproto.message_field(2, group="orientation")


@dataclass(eq=False, repr=False)
class Waypoint(betterproto.Message):
    time: datetime = betterproto.message_field(1)
    location: "Location" = betterproto.message_field(2)


@dataclass(eq=False, repr=False)
class Waypoints(betterproto.Message):
    """A set of waypoints that can be interpolated."""

    waypoint: List["Waypoint"] = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class OrbitData(betterproto.Message):
    """
    Minimal satellite orbit data, e.g. from a TLE, that can be propagated.
    """

    epoch: datetime = betterproto.message_field(1)
    # Orbit inclination in degrees.
    inclination_deg: float = betterproto.double_field(2)
    # Right ascension of the ascending node in degrees.
    raan_deg: float = betterproto.double_field(3)
    # Orbit eccentricity.
    e: float = betterproto.double_field(4)
    # Argument of perigee in degrees.
    argument_deg: float = betterproto.double_field(5)
    # Mean anomaly in degrees.
    anomaly_deg: float = betterproto.double_field(6)
    # Mean motion in revolutions per day.
    mean_motion: float = betterproto.double_field(7)


@dataclass(eq=False, repr=False)
class Motion(betterproto.Message):
    fixed_location: "Location" = betterproto.message_field(1, group="motion")
    waypoints: "Waypoints" = betterproto.message_field(2, group="motion")
    orbit: "OrbitData" = betterproto.message_field(3, group="motion")


@dataclass(eq=False, repr=False)
class Coverage(betterproto.Message):
    """Area of coverage; expected value type of "coverage" aspect"""

    height: int = betterproto.int32_field(1)
    arc_width: int = betterproto.int32_field(2)
    azimuth: int = betterproto.int32_field(3)
    tilt: int = betterproto.int32_field(4)


@dataclass(eq=False, repr=False)
class E2Node(betterproto.Message):
    """
    E2Node aspect; expected value type of "E2NODE" aspect and expected on
    entities of "E2NODE" kind
    """

    service_models: Dict[str, "ServiceModelInfo"] = betterproto.map_field(
        1, betterproto.TYPE_STRING, betterproto.TYPE_MESSAGE
    )


@dataclass(eq=False, repr=False)
class E2NodeConfig(betterproto.Message):
    """E2NodeConfig aspect tracks the managed configuration for an E2 node"""

    connections: List["Interface"] = betterproto.message_field(1)
    version: int = betterproto.uint64_field(2)


@dataclass(eq=False, repr=False)
class Lease(betterproto.Message):
    """Lease aspect with an expiration timestamp for RAN entities"""

    expiration: datetime = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class Interface(betterproto.Message):
    """Interface determines address and type of an endpoint interface"""

    type: "InterfaceType" = betterproto.enum_field(1)
    ip: str = betterproto.string_field(2)
    port: int = betterproto.uint32_field(3)


@dataclass(eq=False, repr=False)
class E2TInfo(betterproto.Message):
    """E2TInfo E2T aspect;"""

    interfaces: List["Interface"] = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class XAppInfo(betterproto.Message):
    interfaces: List["Interface"] = betterproto.message_field(1)
    a1_policy_types: List["A1PolicyType"] = betterproto.message_field(2)


@dataclass(eq=False, repr=False)
class A1PolicyType(betterproto.Message):
    id: str = betterproto.string_field(1)
    name: str = betterproto.string_field(2)
    version: str = betterproto.string_field(3)
    description: str = betterproto.string_field(4)


@dataclass(eq=False, repr=False)
class A1TInfo(betterproto.Message):
    interfaces: List["Interface"] = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class CellGlobalId(betterproto.Message):
    value: str = betterproto.string_field(1)
    type: "CellGlobalIdType" = betterproto.enum_field(2)


@dataclass(eq=False, repr=False)
class NeighborCellId(betterproto.Message):
    cell_global_id: "CellGlobalId" = betterproto.message_field(1)
    plmn_id: str = betterproto.string_field(2)


@dataclass(eq=False, repr=False)
class E2Cell(betterproto.Message):
    """
    E2Cell aspect; expected value type of "E2CELL" aspect and expected on
    entities of "E2CELL" kind
    """

    cell_object_id: str = betterproto.string_field(1)
    cell_global_id: "CellGlobalId" = betterproto.message_field(2)
    antenna_count: int = betterproto.uint32_field(3)
    earfcn: int = betterproto.uint32_field(4)
    cell_type: str = betterproto.string_field(5)
    pci: int = betterproto.uint32_field(6)
    kpi_reports: Dict[str, int] = betterproto.map_field(
        7, betterproto.TYPE_STRING, betterproto.TYPE_UINT32
    )
    neighbor_cell_ids: List["NeighborCellId"] = betterproto.message_field(8)


@dataclass(eq=False, repr=False)
class ServiceModelInfo(betterproto.Message):
    oid: str = betterproto.string_field(1)
    name: str = betterproto.string_field(2)
    ran_functions: List[
        "betterproto_lib_google_protobuf.Any"
    ] = betterproto.message_field(3)
    ran_function_i_ds: List[int] = betterproto.uint32_field(4)


@dataclass(eq=False, repr=False)
class RcRanFunction(betterproto.Message):
    id: str = betterproto.string_field(1)
    report_styles: List["RcReportStyle"] = betterproto.message_field(2)
    insert_styles: List["RcInsertStyle"] = betterproto.message_field(3)
    event_trigger_styles: List["RcEventTriggerStyle"] = betterproto.message_field(4)


@dataclass(eq=False, repr=False)
class MhoRanFunction(betterproto.Message):
    id: str = betterproto.string_field(1)
    report_styles: List["MhoReportStyle"] = betterproto.message_field(2)


@dataclass(eq=False, repr=False)
class KpmRanFunction(betterproto.Message):
    id: str = betterproto.string_field(1)
    report_styles: List["KpmReportStyle"] = betterproto.message_field(2)


@dataclass(eq=False, repr=False)
class RsmRanFunction(betterproto.Message):
    id: str = betterproto.string_field(1)
    ric_slicing_node_capability_list: List[
        "RsmNodeSlicingCapabilityItem"
    ] = betterproto.message_field(2)


@dataclass(eq=False, repr=False)
class RcEventTriggerStyle(betterproto.Message):
    name: str = betterproto.string_field(1)
    type: int = betterproto.int32_field(2)
    format_type: int = betterproto.int32_field(3)


@dataclass(eq=False, repr=False)
class RcReportStyle(betterproto.Message):
    name: str = betterproto.string_field(1)
    type: int = betterproto.int32_field(2)
    ran_parameters: List["RanParameter"] = betterproto.message_field(3)


@dataclass(eq=False, repr=False)
class RcInsertStyle(betterproto.Message):
    name: str = betterproto.string_field(1)
    type: int = betterproto.int32_field(2)
    insert_indications: List["InsertIndication"] = betterproto.message_field(3)


@dataclass(eq=False, repr=False)
class InsertIndication(betterproto.Message):
    id: int = betterproto.int32_field(1)
    name: str = betterproto.string_field(2)
    ran_parameters: List["RanParameter"] = betterproto.message_field(3)


@dataclass(eq=False, repr=False)
class RanParameter(betterproto.Message):
    id: int = betterproto.int32_field(1)
    name: str = betterproto.string_field(2)


@dataclass(eq=False, repr=False)
class KpmReportStyle(betterproto.Message):
    name: str = betterproto.string_field(1)
    type: int = betterproto.int32_field(2)
    measurements: List["KpmMeasurement"] = betterproto.message_field(3)


@dataclass(eq=False, repr=False)
class MhoReportStyle(betterproto.Message):
    name: str = betterproto.string_field(1)
    type: int = betterproto.int32_field(2)


@dataclass(eq=False, repr=False)
class KpmMeasurement(betterproto.Message):
    id: str = betterproto.string_field(1)
    name: str = betterproto.string_field(2)


@dataclass(eq=False, repr=False)
class RsmNodeSlicingCapabilityItem(betterproto.Message):
    max_number_of_slices_dl: int = betterproto.int32_field(1)
    max_number_of_slices_ul: int = betterproto.int32_field(2)
    slicing_type: "RsmSlicingType" = betterproto.enum_field(3)
    max_number_of_ues_per_slice: int = betterproto.int32_field(4)
    supported_config: List["RsmSupportedSlicingConfigItem"] = betterproto.message_field(
        5
    )


@dataclass(eq=False, repr=False)
class RsmSupportedSlicingConfigItem(betterproto.Message):
    slicing_config_type: "E2SmRsmCommand" = betterproto.enum_field(1)


@dataclass(eq=False, repr=False)
class RsmSliceItemList(betterproto.Message):
    rsm_slice_list: List["RsmSlicingItem"] = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class RsmSlicingItem(betterproto.Message):
    id: str = betterproto.string_field(1)
    slice_desc: str = betterproto.string_field(2)
    slice_parameters: "RsmSliceParameters" = betterproto.message_field(3)
    slice_type: "RsmSliceType" = betterproto.enum_field(4)
    ue_id_list: List["UeIdentity"] = betterproto.message_field(5)


@dataclass(eq=False, repr=False)
class RsmSliceParameters(betterproto.Message):
    scheduler_type: "RsmSchedulerType" = betterproto.enum_field(1)
    weight: int = betterproto.int32_field(2)
    qos_level: int = betterproto.int32_field(3)


@dataclass(eq=False, repr=False)
class DuUeF1ApId(betterproto.Message):
    value: int = betterproto.int64_field(1)


@dataclass(eq=False, repr=False)
class CuUeF1ApId(betterproto.Message):
    value: int = betterproto.int64_field(1)


@dataclass(eq=False, repr=False)
class RanUeNgapId(betterproto.Message):
    value: int = betterproto.int64_field(1)


@dataclass(eq=False, repr=False)
class EnbUeS1ApId(betterproto.Message):
    value: int = betterproto.int32_field(1)


@dataclass(eq=False, repr=False)
class AmfUeNgapId(betterproto.Message):
    value: int = betterproto.int64_field(1)


@dataclass(eq=False, repr=False)
class UeIdentity(betterproto.Message):
    du_ue_f1_ap_id: "DuUeF1ApId" = betterproto.message_field(1)
    cu_ue_f1_ap_id: "CuUeF1ApId" = betterproto.message_field(2)
    ran_ue_ngap_id: "RanUeNgapId" = betterproto.message_field(3)
    enb_ue_s1_ap_id: "EnbUeS1ApId" = betterproto.message_field(4)
    amf_ue_ngap_id: "AmfUeNgapId" = betterproto.message_field(5)
    preferred_id_type: "UeIdType" = betterproto.enum_field(6)
    drb_id: "DrbId" = betterproto.message_field(7)


@dataclass(eq=False, repr=False)
class DrbId(betterproto.Message):
    four_gdrb_id: "FourGDrbId" = betterproto.message_field(1, group="drb_id")
    five_gdrb_id: "FiveGDrbId" = betterproto.message_field(2, group="drb_id")


@dataclass(eq=False, repr=False)
class FiveGDrbId(betterproto.Message):
    value: int = betterproto.int32_field(1)
    qfi: "Qfi" = betterproto.message_field(2)
    flows_map_to_drb: List["QoSflowLevelParameters"] = betterproto.message_field(3)


@dataclass(eq=False, repr=False)
class Qfi(betterproto.Message):
    value: int = betterproto.int32_field(1)


@dataclass(eq=False, repr=False)
class QoSflowLevelParameters(betterproto.Message):
    dynamic_five_qi: "DynamicFiveQi" = betterproto.message_field(
        1, group="qos_flow_level_parameters"
    )
    non_dynamic_five_qi: "NonDynamicFiveQi" = betterproto.message_field(
        2, group="qos_flow_level_parameters"
    )


@dataclass(eq=False, repr=False)
class DynamicFiveQi(betterproto.Message):
    priority_level: int = betterproto.int32_field(1)
    packet_delay_budge: int = betterproto.int32_field(2)
    packet_error_rate: int = betterproto.int32_field(3)


@dataclass(eq=False, repr=False)
class NonDynamicFiveQi(betterproto.Message):
    five_qi: "FiveQi" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class FiveQi(betterproto.Message):
    value: int = betterproto.int32_field(1)


@dataclass(eq=False, repr=False)
class FourGDrbId(betterproto.Message):
    value: int = betterproto.int32_field(1)
    qci: "Qci" = betterproto.message_field(2)


@dataclass(eq=False, repr=False)
class Qci(betterproto.Message):
    value: int = betterproto.int32_field(1)


@dataclass(eq=False, repr=False)
class Event(betterproto.Message):
    """Event is a topo operation event"""

    type: "EventType" = betterproto.enum_field(1)
    object: "Object" = betterproto.message_field(2)


@dataclass(eq=False, repr=False)
class CreateRequest(betterproto.Message):
    object: "Object" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class CreateResponse(betterproto.Message):
    object: "Object" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class GetRequest(betterproto.Message):
    id: str = betterproto.string_field(1)


@dataclass(eq=False, repr=False)
class GetResponse(betterproto.Message):
    object: "Object" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class UpdateRequest(betterproto.Message):
    object: "Object" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class UpdateResponse(betterproto.Message):
    object: "Object" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class DeleteRequest(betterproto.Message):
    id: str = betterproto.string_field(1)
    revision: int = betterproto.uint64_field(2)


@dataclass(eq=False, repr=False)
class DeleteResponse(betterproto.Message):
    pass


@dataclass(eq=False, repr=False)
class Filter(betterproto.Message):
    equal: "EqualFilter" = betterproto.message_field(1, group="filter")
    not_: "NotFilter" = betterproto.message_field(2, group="filter")
    in_: "InFilter" = betterproto.message_field(3, group="filter")
    key: str = betterproto.string_field(4)


@dataclass(eq=False, repr=False)
class EqualFilter(betterproto.Message):
    value: str = betterproto.string_field(1)


@dataclass(eq=False, repr=False)
class InFilter(betterproto.Message):
    values: List[str] = betterproto.string_field(1)


@dataclass(eq=False, repr=False)
class NotFilter(betterproto.Message):
    inner: "Filter" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class RelationFilter(betterproto.Message):
    """
    Filter for targets of given relation kinds and given source ids;
    optionally, filters by specified target kind
    """

    src_id: str = betterproto.string_field(1)
    relation_kind: str = betterproto.string_field(2)
    target_kind: str = betterproto.string_field(3)
    scope: "RelationFilterScope" = betterproto.enum_field(4)
    target_id: str = betterproto.string_field(5)


@dataclass(eq=False, repr=False)
class Filters(betterproto.Message):
    kind_filter: "Filter" = betterproto.message_field(1)
    label_filters: List["Filter"] = betterproto.message_field(2)
    relation_filter: "RelationFilter" = betterproto.message_field(3)
    object_types: List["ObjectType"] = betterproto.enum_field(4)
    with_aspects: List[str] = betterproto.string_field(6)


@dataclass(eq=False, repr=False)
class ListRequest(betterproto.Message):
    filters: "Filters" = betterproto.message_field(1)
    sort_order: "SortOrder" = betterproto.enum_field(2)


@dataclass(eq=False, repr=False)
class ListResponse(betterproto.Message):
    objects: List["Object"] = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class WatchRequest(betterproto.Message):
    filters: "Filters" = betterproto.message_field(1)
    noreplay: bool = betterproto.bool_field(2)


@dataclass(eq=False, repr=False)
class WatchResponse(betterproto.Message):
    event: "Event" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class Object(betterproto.Message):
    """
    Object is an one of the following: a kind (archetype of entity or
    relation), an entity, a relation
    """

    uuid: str = betterproto.string_field(9)
    id: str = betterproto.string_field(1)
    revision: int = betterproto.uint64_field(2)
    type: "ObjectType" = betterproto.enum_field(3)
    entity: "Entity" = betterproto.message_field(4, group="obj")
    relation: "Relation" = betterproto.message_field(5, group="obj")
    kind: "Kind" = betterproto.message_field(6, group="obj")
    # Map of aspects as typed values; for kind, these represent expected aspects
    # and their default values
    aspects: Dict[str, "betterproto_lib_google_protobuf.Any"] = betterproto.map_field(
        7, betterproto.TYPE_STRING, betterproto.TYPE_MESSAGE
    )
    # Arbitrary labels for classification/search
    labels: Dict[str, str] = betterproto.map_field(
        8, betterproto.TYPE_STRING, betterproto.TYPE_STRING
    )


@dataclass(eq=False, repr=False)
class Entity(betterproto.Message):
    """Entity represents any "thing" that is represented in the topology"""

    # user-defined entity kind
    kind_id: str = betterproto.string_field(1)
    # these lists are maintained by the system and are provided as read-only
    # values for clients
    src_relation_ids: List[str] = betterproto.string_field(2)
    tgt_relation_ids: List[str] = betterproto.string_field(3)


@dataclass(eq=False, repr=False)
class Relation(betterproto.Message):
    """
    Relation represents any "relation" between two entitites in the topology.
    """

    # user defined relation kind
    kind_id: str = betterproto.string_field(1)
    src_entity_id: str = betterproto.string_field(2)
    tgt_entity_id: str = betterproto.string_field(3)


@dataclass(eq=False, repr=False)
class Kind(betterproto.Message):
    """Kind represents an archetype of an object, i.e. entity or relation"""

    name: str = betterproto.string_field(1)


class TopoStub(betterproto.ServiceStub):
    async def create(self, *, object: "Object" = None) -> "CreateResponse":

        request = CreateRequest()
        if object is not None:
            request.object = object

        return await self._unary_unary(
            "/onos.topo.Topo/Create", request, CreateResponse
        )

    async def get(self, *, id: str = "") -> "GetResponse":

        request = GetRequest()
        request.id = id

        return await self._unary_unary("/onos.topo.Topo/Get", request, GetResponse)

    async def update(self, *, object: "Object" = None) -> "UpdateResponse":

        request = UpdateRequest()
        if object is not None:
            request.object = object

        return await self._unary_unary(
            "/onos.topo.Topo/Update", request, UpdateResponse
        )

    async def delete(self, *, id: str = "", revision: int = 0) -> "DeleteResponse":

        request = DeleteRequest()
        request.id = id
        request.revision = revision

        return await self._unary_unary(
            "/onos.topo.Topo/Delete", request, DeleteResponse
        )

    async def list(
        self, *, filters: "Filters" = None, sort_order: "SortOrder" = None
    ) -> "ListResponse":

        request = ListRequest()
        if filters is not None:
            request.filters = filters
        request.sort_order = sort_order

        return await self._unary_unary("/onos.topo.Topo/List", request, ListResponse)

    async def watch(
        self, *, filters: "Filters" = None, noreplay: bool = False
    ) -> AsyncIterator["WatchResponse"]:

        request = WatchRequest()
        if filters is not None:
            request.filters = filters
        request.noreplay = noreplay

        async for response in self._unary_stream(
            "/onos.topo.Topo/Watch",
            request,
            WatchResponse,
        ):
            yield response


class TopoBase(ServiceBase):
    async def create(self, object: "Object") -> "CreateResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def get(self, id: str) -> "GetResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def update(self, object: "Object") -> "UpdateResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def delete(self, id: str, revision: int) -> "DeleteResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def list(self, filters: "Filters", sort_order: "SortOrder") -> "ListResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def watch(
        self, filters: "Filters", noreplay: bool
    ) -> AsyncIterator["WatchResponse"]:
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def __rpc_create(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "object": request.object,
        }

        response = await self.create(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_get(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "id": request.id,
        }

        response = await self.get(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_update(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "object": request.object,
        }

        response = await self.update(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_delete(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "id": request.id,
            "revision": request.revision,
        }

        response = await self.delete(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_list(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "filters": request.filters,
            "sort_order": request.sort_order,
        }

        response = await self.list(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_watch(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "filters": request.filters,
            "noreplay": request.noreplay,
        }

        await self._call_rpc_handler_server_stream(
            self.watch,
            stream,
            request_kwargs,
        )

    def __mapping__(self) -> Dict[str, grpclib.const.Handler]:
        return {
            "/onos.topo.Topo/Create": grpclib.const.Handler(
                self.__rpc_create,
                grpclib.const.Cardinality.UNARY_UNARY,
                CreateRequest,
                CreateResponse,
            ),
            "/onos.topo.Topo/Get": grpclib.const.Handler(
                self.__rpc_get,
                grpclib.const.Cardinality.UNARY_UNARY,
                GetRequest,
                GetResponse,
            ),
            "/onos.topo.Topo/Update": grpclib.const.Handler(
                self.__rpc_update,
                grpclib.const.Cardinality.UNARY_UNARY,
                UpdateRequest,
                UpdateResponse,
            ),
            "/onos.topo.Topo/Delete": grpclib.const.Handler(
                self.__rpc_delete,
                grpclib.const.Cardinality.UNARY_UNARY,
                DeleteRequest,
                DeleteResponse,
            ),
            "/onos.topo.Topo/List": grpclib.const.Handler(
                self.__rpc_list,
                grpclib.const.Cardinality.UNARY_UNARY,
                ListRequest,
                ListResponse,
            ),
            "/onos.topo.Topo/Watch": grpclib.const.Handler(
                self.__rpc_watch,
                grpclib.const.Cardinality.UNARY_STREAM,
                WatchRequest,
                WatchResponse,
            ),
        }


import betterproto.lib.google.protobuf as betterproto_lib_google_protobuf
