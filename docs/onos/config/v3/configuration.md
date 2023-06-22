# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/config/v3/configuration.proto](#onos_config_v3_configuration-proto)
    - [AppliedConfiguration](#onos-config-v3-AppliedConfiguration)
    - [AppliedConfiguration.ValuesEntry](#onos-config-v3-AppliedConfiguration-ValuesEntry)
    - [CommittedConfiguration](#onos-config-v3-CommittedConfiguration)
    - [CommittedConfiguration.ValuesEntry](#onos-config-v3-CommittedConfiguration-ValuesEntry)
    - [Configuration](#onos-config-v3-Configuration)
    - [ConfigurationEvent](#onos-config-v3-ConfigurationEvent)
    - [ConfigurationID](#onos-config-v3-ConfigurationID)
    - [ConfigurationStatus](#onos-config-v3-ConfigurationStatus)
    - [MastershipStatus](#onos-config-v3-MastershipStatus)
  
    - [ConfigurationEvent.EventType](#onos-config-v3-ConfigurationEvent-EventType)
    - [ConfigurationStatus.State](#onos-config-v3-ConfigurationStatus-State)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos_config_v3_configuration-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/config/v3/configuration.proto



<a name="onos-config-v3-AppliedConfiguration"></a>

### AppliedConfiguration



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| index | [uint64](#uint64) |  |  |
| ordinal | [uint64](#uint64) |  |  |
| revision | [uint64](#uint64) |  |  |
| target | [uint64](#uint64) |  |  |
| term | [uint64](#uint64) |  |  |
| values | [AppliedConfiguration.ValuesEntry](#onos-config-v3-AppliedConfiguration-ValuesEntry) | repeated |  |






<a name="onos-config-v3-AppliedConfiguration-ValuesEntry"></a>

### AppliedConfiguration.ValuesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [PathValue](#onos-config-v3-PathValue) |  |  |






<a name="onos-config-v3-CommittedConfiguration"></a>

### CommittedConfiguration



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| index | [uint64](#uint64) |  |  |
| ordinal | [uint64](#uint64) |  |  |
| revision | [uint64](#uint64) |  |  |
| target | [uint64](#uint64) |  |  |
| change | [uint64](#uint64) |  |  |
| values | [CommittedConfiguration.ValuesEntry](#onos-config-v3-CommittedConfiguration-ValuesEntry) | repeated |  |






<a name="onos-config-v3-CommittedConfiguration-ValuesEntry"></a>

### CommittedConfiguration.ValuesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [PathValue](#onos-config-v3-PathValue) |  |  |






<a name="onos-config-v3-Configuration"></a>

### Configuration
Configuration represents complete desired target configuration


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meta | [ObjectMeta](#onos-config-v3-ObjectMeta) |  |  |
| id | [ConfigurationID](#onos-config-v3-ConfigurationID) |  | &#39;id&#39; is the unique identifier of the configuration |
| committed | [CommittedConfiguration](#onos-config-v3-CommittedConfiguration) |  |  |
| applied | [AppliedConfiguration](#onos-config-v3-AppliedConfiguration) |  |  |
| status | [ConfigurationStatus](#onos-config-v3-ConfigurationStatus) |  | &#39;ConfigurationStatus&#39; is the current lifecycle status of the configuration |






<a name="onos-config-v3-ConfigurationEvent"></a>

### ConfigurationEvent
ConfigurationEvent configuration store event


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [ConfigurationEvent.EventType](#onos-config-v3-ConfigurationEvent-EventType) |  | EventType configuration event type |
| configuration | [Configuration](#onos-config-v3-Configuration) |  |  |






<a name="onos-config-v3-ConfigurationID"></a>

### ConfigurationID



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| target | [Target](#onos-config-v3-Target) |  |  |






<a name="onos-config-v3-ConfigurationStatus"></a>

### ConfigurationStatus
ConfigurationStatus is the status of a Configuration


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| state | [ConfigurationStatus.State](#onos-config-v3-ConfigurationStatus-State) |  | &#39;state&#39; is the configuration state |
| mastership | [MastershipStatus](#onos-config-v3-MastershipStatus) |  |  |






<a name="onos-config-v3-MastershipStatus"></a>

### MastershipStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| master | [string](#string) |  |  |
| term | [uint64](#uint64) |  |  |





 


<a name="onos-config-v3-ConfigurationEvent-EventType"></a>

### ConfigurationEvent.EventType
EventType configuration event types for configuration store

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 | UNKNOWN indicates unknown configuration store event |
| CREATED | 1 | CREATED indicates the configuration entry in the store is created |
| UPDATED | 2 | UPDATED indicates the configuration entry in the store is updated |
| DELETED | 3 | DELETED indicates the configuration entry in the store is deleted |
| REPLAYED | 4 | REPLAYED |



<a name="onos-config-v3-ConfigurationStatus-State"></a>

### ConfigurationStatus.State
State is the configuration state

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 |  |
| SYNCHRONIZING | 1 |  |
| SYNCHRONIZED | 2 |  |
| PERSISTED | 3 |  |


 

 

 



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

