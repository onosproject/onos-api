# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/config/v2/value.proto](#onos_config_v2_value-proto)
    - [PathValue](#onos-config-v2-PathValue)
    - [PathValues](#onos-config-v2-PathValues)
    - [PathValues.ValuesEntry](#onos-config-v2-PathValues-ValuesEntry)
    - [TypedValue](#onos-config-v2-TypedValue)
  
    - [ValueType](#onos-config-v2-ValueType)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos_config_v2_value-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/config/v2/value.proto



<a name="onos-config-v2-PathValue"></a>

### PathValue
PathValue is the state of a path/value in the configuration tree


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| path | [string](#string) |  | &#39;path&#39; is the path to change |
| value | [TypedValue](#onos-config-v2-TypedValue) |  | &#39;value&#39; is the change value |
| deleted | [bool](#bool) |  | &#39;deleted&#39; indicates whether this is a delete |






<a name="onos-config-v2-PathValues"></a>

### PathValues
PathValues is a set of path/value pairs


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| values | [PathValues.ValuesEntry](#onos-config-v2-PathValues-ValuesEntry) | repeated | &#39;values&#39; is a set of change values to apply |






<a name="onos-config-v2-PathValues-ValuesEntry"></a>

### PathValues.ValuesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [PathValue](#onos-config-v2-PathValue) |  |  |






<a name="onos-config-v2-TypedValue"></a>

### TypedValue
TypedValue is a value represented as a byte array


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bytes | [bytes](#bytes) |  | &#39;bytes&#39; is the bytes array |
| type | [ValueType](#onos-config-v2-ValueType) |  | &#39;type&#39; is the value type |
| type_opts | [int32](#int32) | repeated | &#39;type_opts&#39; is a set of type options |





 


<a name="onos-config-v2-ValueType"></a>

### ValueType
ValueType is the type for a value

| Name | Number | Description |
| ---- | ------ | ----------- |
| EMPTY | 0 |  |
| STRING | 1 |  |
| INT | 2 |  |
| UINT | 3 |  |
| BOOL | 4 |  |
| DECIMAL | 5 |  |
| FLOAT | 6 |  |
| BYTES | 7 |  |
| LEAFLIST_STRING | 8 |  |
| LEAFLIST_INT | 9 |  |
| LEAFLIST_UINT | 10 |  |
| LEAFLIST_BOOL | 11 |  |
| LEAFLIST_DECIMAL | 12 |  |
| LEAFLIST_FLOAT | 13 |  |
| LEAFLIST_BYTES | 14 |  |
| DOUBLE | 15 |  |
| LEAFLIST_DOUBLE | 16 |  |


 

 

 



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

