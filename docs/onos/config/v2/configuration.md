# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/config/v2/configuration.proto](#onos/config/v2/configuration.proto)
    - [Configuration](#onos.config.v2.Configuration)
    - [Configuration.ValuesEntry](#onos.config.v2.Configuration.ValuesEntry)
    - [ConfigurationEvent](#onos.config.v2.ConfigurationEvent)
    - [ConfigurationStatus](#onos.config.v2.ConfigurationStatus)
    - [MastershipState](#onos.config.v2.MastershipState)
  
    - [ConfigurationEventType](#onos.config.v2.ConfigurationEventType)
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
| values | [Configuration.ValuesEntry](#onos.config.v2.Configuration.ValuesEntry) | repeated | &#39;values&#39; is a map of path/values to set |
| status | [ConfigurationStatus](#onos.config.v2.ConfigurationStatus) |  | &#39;ConfigurationStatus&#39; is the current lifecycle status of the configuration |
| revision | [uint64](#uint64) |  | revision is configuration revision |
| index | [uint64](#uint64) |  | &#39;index&#39; is a monotonically increasing, globally unique index of the configuration The index is provided by the store, is static and unique for each unique configuration identifier, and should not be modified by client code. |






<a name="onos.config.v2.Configuration.ValuesEntry"></a>

### Configuration.ValuesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [Value](#onos.config.v2.Value) |  |  |






<a name="onos.config.v2.ConfigurationEvent"></a>

### ConfigurationEvent
ConfigurationEvent configuration store event


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [ConfigurationEventType](#onos.config.v2.ConfigurationEventType) |  | ConfigurationEventType configuration event type |
| configuration | [Configuration](#onos.config.v2.Configuration) |  |  |






<a name="onos.config.v2.ConfigurationStatus"></a>

### ConfigurationStatus
ConfigurationStatus is the status of a Configuration


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| state | [ConfigurationState](#onos.config.v2.ConfigurationState) |  | &#39;state&#39; is the state of the transaction within a Phase |
| mastership_state | [MastershipState](#onos.config.v2.MastershipState) |  | mastershipState mastership info |
| transaction_index | [uint64](#uint64) |  | transaction_index highest Transaction index applied to the Configuration |
| sync_index | [uint64](#uint64) |  | sync_index highest transaction index applied to the target. |






<a name="onos.config.v2.MastershipState"></a>

### MastershipState
Mastership state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| term | [uint64](#uint64) |  |  |





 


<a name="onos.config.v2.ConfigurationEventType"></a>

### ConfigurationEventType
ConfigurationEventType configuration event types for configuration store

| Name | Number | Description |
| ---- | ------ | ----------- |
| CONFIGURATION_EVENT_UNKNOWN | 0 | CONFIGURATION_EVENT_UNKNOWN indicates unknown configuration store event |
| CONFIGURATION_CREATED | 1 | CONFIGURATION_CREATED indicates the configuration entry in the store is created |
| CONFIGURATION_UPDATED | 2 | CONFIGURATION_UPDATED indicates the configuration entry in the store is updated |
| CONFIGURATION_DELETED | 3 | CONFIGURATION_DELETED indicates the configuration entry in the store is deleted |
| CONFIGURATION_REPLAYED | 4 | CONFIGURATION_REPLAYED |



<a name="onos.config.v2.ConfigurationState"></a>

### ConfigurationState
ConfigurationState is the configuration state of a configuration phase

| Name | Number | Description |
| ---- | ------ | ----------- |
| CONFIGURATION_PENDING | 0 | CONFIGURATION_PENDING indicates the configuration is PENDING |
| CONFIGURATION_INITIALIZING | 1 | CONFIGURATION_INITIALIZING indicates the configuration is initializing |
| CONFIGURATION_UPDATING | 2 | CONFIGURATION_UPDATING indicates the configuration is updating |
| CONFIGURATION_COMPLETE | 3 | CONFIGURATION_COMPLETE indicates the configuration is complete |
| CONFIGURATION_FAILED | 4 | CONFIGURATION_FAILED indicates the configuration is failed |


 

 

 



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

