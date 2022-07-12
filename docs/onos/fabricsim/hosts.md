# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/fabricsim/hosts.proto](#onos_fabricsim_hosts-proto)
    - [AddHostRequest](#onos-fabricsim-AddHostRequest)
    - [AddHostResponse](#onos-fabricsim-AddHostResponse)
    - [GetHostRequest](#onos-fabricsim-GetHostRequest)
    - [GetHostResponse](#onos-fabricsim-GetHostResponse)
    - [GetHostsRequest](#onos-fabricsim-GetHostsRequest)
    - [GetHostsResponse](#onos-fabricsim-GetHostsResponse)
    - [Host](#onos-fabricsim-Host)
    - [NetworkInterface](#onos-fabricsim-NetworkInterface)
    - [NetworkInterfaceBehavior](#onos-fabricsim-NetworkInterfaceBehavior)
    - [RemoveHostRequest](#onos-fabricsim-RemoveHostRequest)
    - [RemoveHostResponse](#onos-fabricsim-RemoveHostResponse)
  
    - [HostService](#onos-fabricsim-HostService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos_fabricsim_hosts-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/fabricsim/hosts.proto



<a name="onos-fabricsim-AddHostRequest"></a>

### AddHostRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| host | [Host](#onos-fabricsim-Host) |  |  |






<a name="onos-fabricsim-AddHostResponse"></a>

### AddHostResponse







<a name="onos-fabricsim-GetHostRequest"></a>

### GetHostRequest
unique host id






<a name="onos-fabricsim-GetHostResponse"></a>

### GetHostResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| host | [Host](#onos-fabricsim-Host) |  |  |






<a name="onos-fabricsim-GetHostsRequest"></a>

### GetHostsRequest
filters?






<a name="onos-fabricsim-GetHostsResponse"></a>

### GetHostsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hosts | [Host](#onos-fabricsim-Host) | repeated |  |






<a name="onos-fabricsim-Host"></a>

### Host
Host describes a simulated host (bare metal, VM or container)

unique host id
network interfaces






<a name="onos-fabricsim-NetworkInterface"></a>

### NetworkInterface
NetworkInterface describes simulated host&#39;s attachment to the network

mac address
ipv4 address
ipv6 address
behavior






<a name="onos-fabricsim-NetworkInterfaceBehavior"></a>

### NetworkInterfaceBehavior
NetworkInterfaceBehavior describes dynamic aspects of a simulated host interface
and how it manifests its presence on the network.

TBD






<a name="onos-fabricsim-RemoveHostRequest"></a>

### RemoveHostRequest
unique host id






<a name="onos-fabricsim-RemoveHostResponse"></a>

### RemoveHostResponse






 

 

 


<a name="onos-fabricsim-HostService"></a>

### HostService
HostService provides means to control inventory of simulated hosts

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetHosts | [GetHostsRequest](#onos-fabricsim-GetHostsRequest) | [GetHostsResponse](#onos-fabricsim-GetHostsResponse) | GetHosts gets a list of all simulated hosts |
| GetHost | [GetHostRequest](#onos-fabricsim-GetHostRequest) | [GetHostResponse](#onos-fabricsim-GetHostResponse) | GetHost gets a specific host entry |
| AddHost | [AddHostRequest](#onos-fabricsim-AddHostRequest) | [AddHostResponse](#onos-fabricsim-AddHostResponse) | AddHost adds a new simulated host |
| RemoveHost | [RemoveHostRequest](#onos-fabricsim-RemoveHostRequest) | [RemoveHostResponse](#onos-fabricsim-RemoveHostResponse) | RemoveHost removes a simulated host |

 



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

