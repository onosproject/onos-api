# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/fabricsim/types.proto](#onos_fabricsim_types-proto)
    - [Device](#onos-fabricsim-Device)
    - [Host](#onos-fabricsim-Host)
    - [Link](#onos-fabricsim-Link)
    - [NetworkInterface](#onos-fabricsim-NetworkInterface)
    - [NetworkInterfaceBehavior](#onos-fabricsim-NetworkInterfaceBehavior)
    - [Port](#onos-fabricsim-Port)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos_fabricsim_types-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/fabricsim/types.proto



<a name="onos-fabricsim-Device"></a>

### Device
Device describes a simulated switch or IPU

unique_id
type - switch, IPU, etc.
p4_port
gnmi_port
list of ports
behavior?
gnmi models






<a name="onos-fabricsim-Host"></a>

### Host
Host describes a simulated host (bare metal, VM or container)

network interfaces






<a name="onos-fabricsim-Link"></a>

### Link
Link describes a simulated interdevice link, i.e a link between two device ports

source port unique id
destination port unique id
status (operational state derived from adjacent ports status)






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






<a name="onos-fabricsim-Port"></a>

### Port
Port describes a simulated device port

unique_id
type?
name
port number
sdn port number
speed
p4 behaviors? (counters, etc.)
gnmi behaviors? (counters, etc.)





 

 

 

 



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

