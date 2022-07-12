# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/fabricsim/devices.proto](#onos_fabricsim_devices-proto)
    - [AddDeviceRequest](#onos-fabricsim-AddDeviceRequest)
    - [AddDeviceResponse](#onos-fabricsim-AddDeviceResponse)
    - [Device](#onos-fabricsim-Device)
    - [DisablePortRequest](#onos-fabricsim-DisablePortRequest)
    - [DisablePortResponse](#onos-fabricsim-DisablePortResponse)
    - [EnablePortRequest](#onos-fabricsim-EnablePortRequest)
    - [EnablePortResponse](#onos-fabricsim-EnablePortResponse)
    - [GetDeviceRequest](#onos-fabricsim-GetDeviceRequest)
    - [GetDeviceResponse](#onos-fabricsim-GetDeviceResponse)
    - [GetDevicesRequest](#onos-fabricsim-GetDevicesRequest)
    - [GetDevicesResponse](#onos-fabricsim-GetDevicesResponse)
    - [Port](#onos-fabricsim-Port)
    - [RemoveDeviceRequest](#onos-fabricsim-RemoveDeviceRequest)
    - [RemoveDeviceResponse](#onos-fabricsim-RemoveDeviceResponse)
    - [StartDeviceRequest](#onos-fabricsim-StartDeviceRequest)
    - [StartDeviceResponse](#onos-fabricsim-StartDeviceResponse)
    - [StopDeviceRequest](#onos-fabricsim-StopDeviceRequest)
    - [StopDeviceResponse](#onos-fabricsim-StopDeviceResponse)
  
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

unique device id
type - switch, IPU, etc.
p4 port
gnmi port
list of ports
behavior?
gnmi models






<a name="onos-fabricsim-DisablePortRequest"></a>

### DisablePortRequest
unique port id






<a name="onos-fabricsim-DisablePortResponse"></a>

### DisablePortResponse







<a name="onos-fabricsim-EnablePortRequest"></a>

### EnablePortRequest
unique port id






<a name="onos-fabricsim-EnablePortResponse"></a>

### EnablePortResponse







<a name="onos-fabricsim-GetDeviceRequest"></a>

### GetDeviceRequest
unique device id






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






<a name="onos-fabricsim-Port"></a>

### Port
Port describes a simulated device port

unique port id
type?
name
port number
sdn port number
speed
p4 behaviors? (counters, etc.)
gnmi behaviors? (counters, etc.)






<a name="onos-fabricsim-RemoveDeviceRequest"></a>

### RemoveDeviceRequest
unique device id






<a name="onos-fabricsim-RemoveDeviceResponse"></a>

### RemoveDeviceResponse







<a name="onos-fabricsim-StartDeviceRequest"></a>

### StartDeviceRequest
unique device id






<a name="onos-fabricsim-StartDeviceResponse"></a>

### StartDeviceResponse







<a name="onos-fabricsim-StopDeviceRequest"></a>

### StopDeviceRequest
unique device id






<a name="onos-fabricsim-StopDeviceResponse"></a>

### StopDeviceResponse






 

 

 


<a name="onos-fabricsim-DeviceService"></a>

### DeviceService
DeviceService provides means to control inventory of simulated devices (switches and IPUs) and their ports

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetDevices | [GetDevicesRequest](#onos-fabricsim-GetDevicesRequest) | [GetDevicesResponse](#onos-fabricsim-GetDevicesResponse) | GetDevices gets a list of all simulated devices (switches and/or IPUs) |
| GetDevice | [GetDeviceRequest](#onos-fabricsim-GetDeviceRequest) | [GetDeviceResponse](#onos-fabricsim-GetDeviceResponse) | GetDevice gets a specific device entry |
| AddDevice | [AddDeviceRequest](#onos-fabricsim-AddDeviceRequest) | [AddDeviceResponse](#onos-fabricsim-AddDeviceResponse) | AddDevice creates a new simulated deviceand start its P4Runtime and gNMI services |
| RemoveDevice | [AddDeviceRequest](#onos-fabricsim-AddDeviceRequest) | [RemoveDeviceResponse](#onos-fabricsim-RemoveDeviceResponse) | RemoveDevice removes a simulated device |
| StopDevice | [StopDeviceRequest](#onos-fabricsim-StopDeviceRequest) | [StopDeviceResponse](#onos-fabricsim-StopDeviceResponse) | StopDevice stops the simulated deviceP4Runtime and gNMI services |
| StartDevice | [StartDeviceRequest](#onos-fabricsim-StartDeviceRequest) | [StartDeviceResponse](#onos-fabricsim-StartDeviceResponse) | StartDevice starts the simulated deviceP4Runtime and gNMI services |
| DisablePort | [DisablePortRequest](#onos-fabricsim-DisablePortRequest) | [DisablePortResponse](#onos-fabricsim-DisablePortResponse) | DisablePort disables the specified port |
| EnablePort | [EnablePortRequest](#onos-fabricsim-EnablePortRequest) | [EnablePortResponse](#onos-fabricsim-EnablePortResponse) | EnablePort enables the specified port |

 



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

