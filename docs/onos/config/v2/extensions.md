# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/config/v2/extensions.proto](#onos/config/v2/extensions.proto)
    - [TransactionControl](#onos.config.v2.TransactionControl)
    - [TransactionInfo](#onos.config.v2.TransactionInfo)
    - [TransactionMode](#onos.config.v2.TransactionMode)
  
    - [TransactionCommand](#onos.config.v2.TransactionCommand)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos/config/v2/extensions.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/config/v2/extensions.proto



<a name="onos.config.v2.TransactionControl"></a>

### TransactionControl
TransactionControl is a extension that if supported by targets enables atomic
transactions across multiple targets.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| command | [TransactionCommand](#onos.config.v2.TransactionCommand) |  |  |






<a name="onos.config.v2.TransactionInfo"></a>

### TransactionInfo
TransactionInfo is a bi-directional extension carrying transaction information between the
client and onos-config.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| index | [uint64](#uint64) |  |  |






<a name="onos.config.v2.TransactionMode"></a>

### TransactionMode
TransactionMode is an extension for constraining the execution of a transaction for
stronger consistency guarantees.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sync | [bool](#bool) |  |  |
| atomic | [bool](#bool) |  |  |





 


<a name="onos.config.v2.TransactionCommand"></a>

### TransactionCommand
TransactionCommand describes phases of the two-phase transaction commit protocol.

| Name | Number | Description |
| ---- | ------ | ----------- |
| PREPARE | 0 |  |
| COMMIT | 1 |  |
| ROLLBACK | 2 |  |


 

 

 



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

