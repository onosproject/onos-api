# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/ricapi/e2/v1beta1/ricapi_e2.proto](#api/ricapi/e2/v1beta1/ricapi_e2.proto)
    - [ControlRequest](#ricapi.e2.v1beta1.ControlRequest)
    - [ControlResponse](#ricapi.e2.v1beta1.ControlResponse)
    - [Indication](#ricapi.e2.v1beta1.Indication)
    - [StreamRequest](#ricapi.e2.v1beta1.StreamRequest)
    - [StreamResponse](#ricapi.e2.v1beta1.StreamResponse)
  
    - [E2TService](#ricapi.e2.v1beta1.E2TService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api/ricapi/e2/v1beta1/ricapi_e2.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/ricapi/e2/v1beta1/ricapi_e2.proto



<a name="ricapi.e2.v1beta1.ControlRequest"></a>

### ControlRequest
ControlRequest






<a name="ricapi.e2.v1beta1.ControlResponse"></a>

### ControlResponse
ControlResponse






<a name="ricapi.e2.v1beta1.Indication"></a>

### Indication
Indication an indication message






<a name="ricapi.e2.v1beta1.StreamRequest"></a>

### StreamRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| app_id | [string](#string) |  |  |
| instance_id | [string](#string) |  |  |
| subscription_id | [string](#string) |  |  |
| encoding | [ricapi.e2.headers.v1beta1.EncodingType](#ricapi.e2.headers.v1beta1.EncodingType) |  |  |






<a name="ricapi.e2.v1beta1.StreamResponse"></a>

### StreamResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| indication_header | [ricapi.e2.headers.v1beta1.ResponseHeader](#ricapi.e2.headers.v1beta1.ResponseHeader) |  |  |
| indication_message | [bytes](#bytes) |  |  |





 

 

 


<a name="ricapi.e2.v1beta1.E2TService"></a>

### E2TService
E2TService provides means for enhanced interactions with the ONOS RIC E2 Termination service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Stream | [StreamRequest](#ricapi.e2.v1beta1.StreamRequest) stream | [StreamResponse](#ricapi.e2.v1beta1.StreamResponse) stream | Stream opens an indications stream |

 



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

