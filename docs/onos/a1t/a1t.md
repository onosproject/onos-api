# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/a1t/a1t.proto](#onos/a1t/a1t.proto)
    - [CreateRequest](#onos.a1t.a1t.CreateRequest)
    - [CreateResponse](#onos.a1t.a1t.CreateResponse)
    - [DeleteRequest](#onos.a1t.a1t.DeleteRequest)
    - [DeleteResponse](#onos.a1t.a1t.DeleteResponse)
    - [EIJob](#onos.a1t.a1t.EIJob)
    - [GetRequest](#onos.a1t.a1t.GetRequest)
    - [GetResponse](#onos.a1t.a1t.GetResponse)
    - [ListResponse](#onos.a1t.a1t.ListResponse)
    - [Object](#onos.a1t.a1t.Object)
    - [Policy](#onos.a1t.a1t.Policy)
    - [Status](#onos.a1t.a1t.Status)
    - [Status.Problem](#onos.a1t.a1t.Status.Problem)
    - [Subscription](#onos.a1t.a1t.Subscription)
    - [UpdateRequest](#onos.a1t.a1t.UpdateRequest)
    - [UpdateResponse](#onos.a1t.a1t.UpdateResponse)
  
    - [Encoding](#onos.a1t.a1t.Encoding)
    - [Object.Type](#onos.a1t.a1t.Object.Type)
    - [Subscription.Type](#onos.a1t.a1t.Subscription.Type)
  
    - [A1TService](#onos.a1t.a1t.A1TService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos/a1t/a1t.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/a1t/a1t.proto



<a name="onos.a1t.a1t.CreateRequest"></a>

### CreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object | [Object](#onos.a1t.a1t.Object) |  |  |






<a name="onos.a1t.a1t.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object | [Object](#onos.a1t.a1t.Object) |  |  |






<a name="onos.a1t.a1t.DeleteRequest"></a>

### DeleteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object | [Object](#onos.a1t.a1t.Object) |  |  |






<a name="onos.a1t.a1t.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object | [Object](#onos.a1t.a1t.Object) |  |  |






<a name="onos.a1t.a1t.EIJob"></a>

### EIJob



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| typeid | [string](#string) |  |  |
| object | [bytes](#bytes) |  |  |
| status | [Status](#onos.a1t.a1t.Status) |  |  |
| encoding | [Encoding](#onos.a1t.a1t.Encoding) |  |  |






<a name="onos.a1t.a1t.GetRequest"></a>

### GetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object | [Object](#onos.a1t.a1t.Object) |  |  |






<a name="onos.a1t.a1t.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object | [Object](#onos.a1t.a1t.Object) |  |  |






<a name="onos.a1t.a1t.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| objects | [Object](#onos.a1t.a1t.Object) | repeated |  |






<a name="onos.a1t.a1t.Object"></a>

### Object



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [Object.Type](#onos.a1t.a1t.Object.Type) |  |  |
| policy | [Policy](#onos.a1t.a1t.Policy) |  |  |
| eijob | [EIJob](#onos.a1t.a1t.EIJob) |  |  |
| subscription | [Subscription](#onos.a1t.a1t.Subscription) |  |  |






<a name="onos.a1t.a1t.Policy"></a>

### Policy



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| typeid | [string](#string) |  |  |
| object | [bytes](#bytes) |  |  |
| status | [Status](#onos.a1t.a1t.Status) |  |  |
| encoding | [Encoding](#onos.a1t.a1t.Encoding) |  |  |






<a name="onos.a1t.a1t.Status"></a>

### Status



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  |  |
| problem | [Status.Problem](#onos.a1t.a1t.Status.Problem) |  |  |






<a name="onos.a1t.a1t.Status.Problem"></a>

### Status.Problem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| detail | [string](#string) |  |  |
| instance | [string](#string) |  |  |
| status | [string](#string) |  |  |
| title | [string](#string) |  |  |
| type | [string](#string) |  |  |






<a name="onos.a1t.a1t.Subscription"></a>

### Subscription



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| type | [Subscription.Type](#onos.a1t.a1t.Subscription.Type) |  |  |
| address | [string](#string) |  |  |






<a name="onos.a1t.a1t.UpdateRequest"></a>

### UpdateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object | [Object](#onos.a1t.a1t.Object) |  |  |






<a name="onos.a1t.a1t.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object | [Object](#onos.a1t.a1t.Object) |  |  |





 


<a name="onos.a1t.a1t.Encoding"></a>

### Encoding


| Name | Number | Description |
| ---- | ------ | ----------- |
| PROTO | 0 |  |
| JSON | 1 |  |



<a name="onos.a1t.a1t.Object.Type"></a>

### Object.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNSPECIFIED | 0 |  |
| POLICY | 1 |  |
| EIJOB | 2 |  |
| SUBSCRIPTION | 3 |  |



<a name="onos.a1t.a1t.Subscription.Type"></a>

### Subscription.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| POLICY | 0 |  |
| EIJOB | 1 |  |


 

 


<a name="onos.a1t.a1t.A1TService"></a>

### A1TService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateRequest](#onos.a1t.a1t.CreateRequest) | [CreateResponse](#onos.a1t.a1t.CreateResponse) |  |
| Get | [GetRequest](#onos.a1t.a1t.GetRequest) | [GetResponse](#onos.a1t.a1t.GetResponse) |  |
| Update | [UpdateRequest](#onos.a1t.a1t.UpdateRequest) | [UpdateResponse](#onos.a1t.a1t.UpdateResponse) |  |
| Delete | [DeleteRequest](#onos.a1t.a1t.DeleteRequest) | [DeleteResponse](#onos.a1t.a1t.DeleteResponse) |  |
| List | [GetRequest](#onos.a1t.a1t.GetRequest) | [ListResponse](#onos.a1t.a1t.ListResponse) |  |
| Watch | [GetRequest](#onos.a1t.a1t.GetRequest) | [GetResponse](#onos.a1t.a1t.GetResponse) stream |  |

 



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

