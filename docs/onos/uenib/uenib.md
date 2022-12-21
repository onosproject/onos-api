# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/uenib/uenib.proto](#onos_uenib_uenib-proto)
    - [CreateUERequest](#onos-uenib-CreateUERequest)
    - [CreateUEResponse](#onos-uenib-CreateUEResponse)
    - [DeleteUERequest](#onos-uenib-DeleteUERequest)
    - [DeleteUEResponse](#onos-uenib-DeleteUEResponse)
    - [Event](#onos-uenib-Event)
    - [GetUERequest](#onos-uenib-GetUERequest)
    - [GetUEResponse](#onos-uenib-GetUEResponse)
    - [ListUERequest](#onos-uenib-ListUERequest)
    - [ListUEResponse](#onos-uenib-ListUEResponse)
    - [UE](#onos-uenib-UE)
    - [UE.AspectsEntry](#onos-uenib-UE-AspectsEntry)
    - [UpdateUERequest](#onos-uenib-UpdateUERequest)
    - [UpdateUEResponse](#onos-uenib-UpdateUEResponse)
    - [WatchUERequest](#onos-uenib-WatchUERequest)
    - [WatchUEResponse](#onos-uenib-WatchUEResponse)
  
    - [EventType](#onos-uenib-EventType)
  
    - [UEService](#onos-uenib-UEService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos_uenib_uenib-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/uenib/uenib.proto



<a name="onos-uenib-CreateUERequest"></a>

### CreateUERequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ue | [UE](#onos-uenib-UE) |  |  |






<a name="onos-uenib-CreateUEResponse"></a>

### CreateUEResponse







<a name="onos-uenib-DeleteUERequest"></a>

### DeleteUERequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| aspect_types | [string](#string) | repeated |  |






<a name="onos-uenib-DeleteUEResponse"></a>

### DeleteUEResponse







<a name="onos-uenib-Event"></a>

### Event
Event is a record of an operation on a UE


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [EventType](#onos-uenib-EventType) |  |  |
| ue | [UE](#onos-uenib-UE) |  |  |






<a name="onos-uenib-GetUERequest"></a>

### GetUERequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| aspect_types | [string](#string) | repeated |  |






<a name="onos-uenib-GetUEResponse"></a>

### GetUEResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ue | [UE](#onos-uenib-UE) |  |  |






<a name="onos-uenib-ListUERequest"></a>

### ListUERequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| aspect_types | [string](#string) | repeated |  |






<a name="onos-uenib-ListUEResponse"></a>

### ListUEResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ue | [UE](#onos-uenib-UE) |  |  |






<a name="onos-uenib-UE"></a>

### UE
UE entity is merely an ID and a map of arbitrary aspects.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| aspects | [UE.AspectsEntry](#onos-uenib-UE-AspectsEntry) | repeated | Map of aspects as typed values |






<a name="onos-uenib-UE-AspectsEntry"></a>

### UE.AspectsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Any](#google-protobuf-Any) |  |  |






<a name="onos-uenib-UpdateUERequest"></a>

### UpdateUERequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ue | [UE](#onos-uenib-UE) |  |  |






<a name="onos-uenib-UpdateUEResponse"></a>

### UpdateUEResponse







<a name="onos-uenib-WatchUERequest"></a>

### WatchUERequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| noreplay | [bool](#bool) |  |  |
| aspect_types | [string](#string) | repeated |  |






<a name="onos-uenib-WatchUEResponse"></a>

### WatchUEResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#onos-uenib-Event) |  |  |





 


<a name="onos-uenib-EventType"></a>

### EventType
EventType is a UE operation event type

| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 |  |
| ADDED | 1 |  |
| UPDATED | 2 |  |
| REMOVED | 3 |  |


 

 


<a name="onos-uenib-UEService"></a>

### UEService
UEService provides an API for managing UEs and various aspects of
information associated with UEs.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateUE | [CreateUERequest](#onos-uenib-CreateUERequest) | [CreateUEResponse](#onos-uenib-CreateUEResponse) | Create a new UE entity and its initial set of aspects. |
| GetUE | [GetUERequest](#onos-uenib-GetUERequest) | [GetUEResponse](#onos-uenib-GetUEResponse) | Get a UE entity populated with the requested aspects. |
| UpdateUE | [UpdateUERequest](#onos-uenib-UpdateUERequest) | [UpdateUEResponse](#onos-uenib-UpdateUEResponse) | Update an existing UE entity populated with the requested aspects. Only the aspects present in the UE entity will be updated; others will be left unmodified. New aspects can be added via update. |
| DeleteUE | [DeleteUERequest](#onos-uenib-DeleteUERequest) | [DeleteUEResponse](#onos-uenib-DeleteUEResponse) | Delete the specified aspects of a UE entity. |
| ListUEs | [ListUERequest](#onos-uenib-ListUERequest) | [ListUEResponse](#onos-uenib-ListUEResponse) stream | ListUEs returns a stream of UE entities populated the requested aspects. |
| WatchUEs | [WatchUERequest](#onos-uenib-WatchUERequest) | [WatchUEResponse](#onos-uenib-WatchUEResponse) stream | WatchUEs returns a stream of UE change notifications, with each UE populated with only the requested aspects. |

 



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

