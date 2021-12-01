# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/config/v2/configuration.proto](#onos/config/v2/configuration.proto)
    - [Configuration](#onos.config.v2.Configuration)
    - [ConfigurationStatus](#onos.config.v2.ConfigurationStatus)
    - [MastershipState](#onos.config.v2.MastershipState)
  
    - [ConfigurationState](#onos.config.v2.ConfigurationState)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos/config/v2/configuration.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/config/v2/configuration.proto



<a name="onos.config.v2.Configuration"></a>

### Configuration
Configuration represents complete desired target configuration


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | &#39;id&#39; is a unique configuration identifier |
| target_id | [string](#string) |  | &#39;target_id&#39; is the target to which the desired target configuration applies |
| target_version | [string](#string) |  | &#39;target_version&#39; is the version to which desired target configuration applies |
| target_type | [string](#string) |  | &#39;target_type&#39; is an optional target type to which to apply this desired target configuration |
| values | [PathValue](#onos.config.v2.PathValue) | repeated | &#39;values&#39; is a list of path/values to set |
| status | [ConfigurationStatus](#onos.config.v2.ConfigurationStatus) |  | &#39;ConfigurationStatus&#39; is the current lifecycle status of the configuration |






<a name="onos.config.v2.ConfigurationStatus"></a>

### ConfigurationStatus
ConfigurationStatus is the status of a Configuration


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| state | [ConfigurationState](#onos.config.v2.ConfigurationState) |  | &#39;state&#39; is the state of the transaction within a Phase |
| message | [string](#string) |  | message is a result message |
| mastershipState | [MastershipState](#onos.config.v2.MastershipState) |  | mastershipState mastership info |






<a name="onos.config.v2.MastershipState"></a>

### MastershipState
Mastership state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| term | [uint64](#uint64) |  |  |





 


<a name="onos.config.v2.ConfigurationState"></a>

### ConfigurationState
ConfigurationState is the configuration state of a configuration phase

| Name | Number | Description |
| ---- | ------ | ----------- |
| CONFIGURATION_PENDING | 0 | CONFIGURATION_PENDING indicates the configuration is PENDING |
| CONFIGURATION_COMPLETE | 2 | COMPLETE indicates the configuration is COMPLETE |


 

 

 



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

