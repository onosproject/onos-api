# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/discovery/discovery.proto](#onos_discovery_discovery-proto)
    - [AddPodRequest](#onos-discovery-AddPodRequest)
    - [AddPodResponse](#onos-discovery-AddPodResponse)
    - [AddRackRequest](#onos-discovery-AddRackRequest)
    - [AddRackResponse](#onos-discovery-AddRackResponse)
    - [AddServerIPURequest](#onos-discovery-AddServerIPURequest)
    - [AddServerIPUResponse](#onos-discovery-AddServerIPUResponse)
    - [AddSwitchRequest](#onos-discovery-AddSwitchRequest)
    - [AddSwitchResponse](#onos-discovery-AddSwitchResponse)
    - [InjectedLink](#onos-discovery-InjectedLink)
  
    - [DiscoveryService](#onos-discovery-DiscoveryService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos_discovery_discovery-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/discovery/discovery.proto



<a name="onos-discovery-AddPodRequest"></a>

### AddPodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="onos-discovery-AddPodResponse"></a>

### AddPodResponse







<a name="onos-discovery-AddRackRequest"></a>

### AddRackRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| pod_id | [string](#string) |  |  |






<a name="onos-discovery-AddRackResponse"></a>

### AddRackResponse







<a name="onos-discovery-AddServerIPURequest"></a>

### AddServerIPURequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| pod_id | [string](#string) |  |  |
| rack_id | [string](#string) |  |  |
| p4_endpoint | [string](#string) |  |  |
| gnmi_endpoint | [string](#string) |  |  |
| pipeline_config_id | [string](#string) |  |  |
| chassis_config_id | [string](#string) |  |  |
| links | [InjectedLink](#onos-discovery-InjectedLink) | repeated | Provisional feature used to inject links until dynamic discovery is available |






<a name="onos-discovery-AddServerIPUResponse"></a>

### AddServerIPUResponse







<a name="onos-discovery-AddSwitchRequest"></a>

### AddSwitchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| pod_id | [string](#string) |  |  |
| rack_id | [string](#string) |  |  |
| p4_endpoint | [string](#string) |  |  |
| gnmi_endpoint | [string](#string) |  |  |
| pipeline_config_id | [string](#string) |  |  |
| chassis_config_id | [string](#string) |  |  |






<a name="onos-discovery-AddSwitchResponse"></a>

### AddSwitchResponse







<a name="onos-discovery-InjectedLink"></a>

### InjectedLink



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| port | [uint64](#uint64) |  |  |
| remote_port | [string](#string) |  |  |





 

 

 


<a name="onos-discovery-DiscoveryService"></a>

### DiscoveryService
DiscoveryService allows injection of topology objects to act as seeds for the topology
discovery.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddPod | [AddPodRequest](#onos-discovery-AddPodRequest) | [AddPodResponse](#onos-discovery-AddPodResponse) | AddPod adds a new POD entity with the requisite aspects |
| AddRack | [AddRackRequest](#onos-discovery-AddRackRequest) | [AddRackResponse](#onos-discovery-AddRackResponse) | AddRack adds a new rack entity with the requisite aspects as part of a POD |
| AddSwitch | [AddSwitchRequest](#onos-discovery-AddSwitchRequest) | [AddSwitchResponse](#onos-discovery-AddSwitchResponse) | AddSwitch adds a new switch entity with the requisite aspects into a rack |
| AddServerIPU | [AddServerIPURequest](#onos-discovery-AddServerIPURequest) | [AddServerIPUResponse](#onos-discovery-AddServerIPUResponse) | AddServerIPU adds a new server entity and an associated IPU entity, both with the requisite aspects into a rack |

 



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

