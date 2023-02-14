# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/fabricsim/devices.proto](#onos_fabricsim_devices-proto)
    - [AddDeviceRequest](#onos-fabricsim-AddDeviceRequest)
    - [AddDeviceResponse](#onos-fabricsim-AddDeviceResponse)
    - [Device](#onos-fabricsim-Device)
    - [DisablePortRequest](#onos-fabricsim-DisablePortRequest)
    - [DisablePortResponse](#onos-fabricsim-DisablePortResponse)
    - [EmitLLDPPacketRequest](#onos-fabricsim-EmitLLDPPacketRequest)
    - [EmitLLDPPacketResponse](#onos-fabricsim-EmitLLDPPacketResponse)
    - [EnablePortRequest](#onos-fabricsim-EnablePortRequest)
    - [EnablePortResponse](#onos-fabricsim-EnablePortResponse)
    - [EntitiesInfo](#onos-fabricsim-EntitiesInfo)
    - [GetDeviceRequest](#onos-fabricsim-GetDeviceRequest)
    - [GetDeviceResponse](#onos-fabricsim-GetDeviceResponse)
    - [GetDevicesRequest](#onos-fabricsim-GetDevicesRequest)
    - [GetDevicesResponse](#onos-fabricsim-GetDevicesResponse)
    - [PipelineInfo](#onos-fabricsim-PipelineInfo)
    - [Port](#onos-fabricsim-Port)
    - [RemoveDeviceRequest](#onos-fabricsim-RemoveDeviceRequest)
    - [RemoveDeviceResponse](#onos-fabricsim-RemoveDeviceResponse)
    - [StartDeviceRequest](#onos-fabricsim-StartDeviceRequest)
    - [StartDeviceResponse](#onos-fabricsim-StartDeviceResponse)
    - [StopDeviceRequest](#onos-fabricsim-StopDeviceRequest)
    - [StopDeviceResponse](#onos-fabricsim-StopDeviceResponse)
  
    - [DeviceType](#onos-fabricsim-DeviceType)
    - [StopMode](#onos-fabricsim-StopMode)
  
    - [DeviceService](#onos-fabricsim-DeviceService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos_fabricsim_devices-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/fabricsim/devices.proto



<a name="onos-fabricsim-AddDeviceRequest"></a>

### AddDeviceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| device | [Device](#onos-fabricsim-Device) |  |  |






<a name="onos-fabricsim-AddDeviceResponse"></a>

### AddDeviceResponse







<a name="onos-fabricsim-Device"></a>

### Device
Device describes a simulated switch or IPU


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | unique device id and device type |
| type | [DeviceType](#onos-fabricsim-DeviceType) |  |  |
| ports | [Port](#onos-fabricsim-Port) | repeated | list of ports |
| control_port | [int32](#int32) |  | control port for p4 and gnmi simulation |
| chassis_id | [uint64](#uint64) |  | unique chassis ID |
| pipeline_info | [PipelineInfo](#onos-fabricsim-PipelineInfo) |  | forwarding pipeline information |
| pos | [onos.misc.GridPosition](#onos-misc-GridPosition) |  | Screen coordinates |
| connections | [onos.misc.Connection](#onos-misc-Connection) | repeated | Current connections and total connection count |
| total_connections | [int32](#int32) |  |  |
| io_stats | [onos.misc.IOStats](#onos-misc-IOStats) |  | Cumulative I/O stats for the device P4Runtime, gNMI and gNOI agent(s) |






<a name="onos-fabricsim-DisablePortRequest"></a>

### DisablePortRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| mode | [StopMode](#onos-fabricsim-StopMode) |  |  |






<a name="onos-fabricsim-DisablePortResponse"></a>

### DisablePortResponse







<a name="onos-fabricsim-EmitLLDPPacketRequest"></a>

### EmitLLDPPacketRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| port_id | [string](#string) |  |  |
| packet | [bytes](#bytes) |  |  |
| ingress_port | [uint32](#uint32) |  |  |
| role_agent_id | [uint32](#uint32) |  |  |






<a name="onos-fabricsim-EmitLLDPPacketResponse"></a>

### EmitLLDPPacketResponse







<a name="onos-fabricsim-EnablePortRequest"></a>

### EnablePortRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="onos-fabricsim-EnablePortResponse"></a>

### EnablePortResponse







<a name="onos-fabricsim-EntitiesInfo"></a>

### EntitiesInfo
EntitiesInfo provides information about size of pipeline entities, tables, meters, counters


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint32](#uint32) |  |  |
| size | [uint32](#uint32) |  |  |
| name | [string](#string) |  |  |






<a name="onos-fabricsim-GetDeviceRequest"></a>

### GetDeviceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="onos-fabricsim-GetDeviceResponse"></a>

### GetDeviceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| device | [Device](#onos-fabricsim-Device) |  |  |






<a name="onos-fabricsim-GetDevicesRequest"></a>

### GetDevicesRequest
filters?






<a name="onos-fabricsim-GetDevicesResponse"></a>

### GetDevicesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| devices | [Device](#onos-fabricsim-Device) | repeated |  |






<a name="onos-fabricsim-PipelineInfo"></a>

### PipelineInfo
PipelineInfo provides information about the currently deployed forwarding pipeline


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cookie | [uint64](#uint64) |  |  |
| p4_info | [bytes](#bytes) |  |  |
| tables | [EntitiesInfo](#onos-fabricsim-EntitiesInfo) | repeated | summary information about tables, counters, meters, groups, etc. |
| counters | [EntitiesInfo](#onos-fabricsim-EntitiesInfo) | repeated |  |
| meters | [EntitiesInfo](#onos-fabricsim-EntitiesInfo) | repeated |  |
| groups | [EntitiesInfo](#onos-fabricsim-EntitiesInfo) | repeated |  |
| multicast_groups | [EntitiesInfo](#onos-fabricsim-EntitiesInfo) | repeated |  |
| clone_sessions | [EntitiesInfo](#onos-fabricsim-EntitiesInfo) | repeated |  |






<a name="onos-fabricsim-Port"></a>

### Port
Port describes a simulated device port


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | unique port id and port type |
| name | [string](#string) |  | display/friendly name |
| number | [uint32](#uint32) |  | port number |
| internal_number | [uint32](#uint32) |  | sdn/internal port number |
| speed | [string](#string) |  | speed and status |
| enabled | [bool](#bool) |  |  |






<a name="onos-fabricsim-RemoveDeviceRequest"></a>

### RemoveDeviceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="onos-fabricsim-RemoveDeviceResponse"></a>

### RemoveDeviceResponse







<a name="onos-fabricsim-StartDeviceRequest"></a>

### StartDeviceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="onos-fabricsim-StartDeviceResponse"></a>

### StartDeviceResponse







<a name="onos-fabricsim-StopDeviceRequest"></a>

### StopDeviceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| mode | [StopMode](#onos-fabricsim-StopMode) |  |  |






<a name="onos-fabricsim-StopDeviceResponse"></a>

### StopDeviceResponse






 


<a name="onos-fabricsim-DeviceType"></a>

### DeviceType
DeviceType represents type of a device, i.e. switch, IPU, etc.

| Name | Number | Description |
| ---- | ------ | ----------- |
| SWITCH | 0 | default assumption |
| IPU | 1 |  |



<a name="onos-fabricsim-StopMode"></a>

### StopMode
StopMode indicates whether to simulate orderly (administrative) or chaotic (power off) shutdown

| Name | Number | Description |
| ---- | ------ | ----------- |
| ORDERLY_STOP | 0 |  |
| CHAOTIC_STOP | 1 |  |


 

 


<a name="onos-fabricsim-DeviceService"></a>

### DeviceService
DeviceService provides means to control inventory of simulated devices (switches and IPUs) and their ports

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetDevices | [GetDevicesRequest](#onos-fabricsim-GetDevicesRequest) | [GetDevicesResponse](#onos-fabricsim-GetDevicesResponse) | GetDevices gets a list of all simulated devices (switches and/or IPUs) |
| GetDevice | [GetDeviceRequest](#onos-fabricsim-GetDeviceRequest) | [GetDeviceResponse](#onos-fabricsim-GetDeviceResponse) | GetDevice gets a specific device entry |
| AddDevice | [AddDeviceRequest](#onos-fabricsim-AddDeviceRequest) | [AddDeviceResponse](#onos-fabricsim-AddDeviceResponse) | AddDevice creates a new simulated deviceand start its P4Runtime and gNMI services |
| RemoveDevice | [RemoveDeviceRequest](#onos-fabricsim-RemoveDeviceRequest) | [RemoveDeviceResponse](#onos-fabricsim-RemoveDeviceResponse) | RemoveDevice removes a simulated device |
| StopDevice | [StopDeviceRequest](#onos-fabricsim-StopDeviceRequest) | [StopDeviceResponse](#onos-fabricsim-StopDeviceResponse) | StopDevice stops the simulated deviceP4Runtime and gNMI services |
| StartDevice | [StartDeviceRequest](#onos-fabricsim-StartDeviceRequest) | [StartDeviceResponse](#onos-fabricsim-StartDeviceResponse) | StartDevice starts the simulated deviceP4Runtime and gNMI services |
| DisablePort | [DisablePortRequest](#onos-fabricsim-DisablePortRequest) | [DisablePortResponse](#onos-fabricsim-DisablePortResponse) | DisablePort disables the specified port |
| EnablePort | [EnablePortRequest](#onos-fabricsim-EnablePortRequest) | [EnablePortResponse](#onos-fabricsim-EnablePortResponse) | EnablePort enables the specified port |
| EmitLLDPPacket | [EmitLLDPPacketRequest](#onos-fabricsim-EmitLLDPPacketRequest) | [EmitLLDPPacketResponse](#onos-fabricsim-EmitLLDPPacketResponse) | EmitLLDPPacket emits the specified LLDP packet on a given device port. |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

