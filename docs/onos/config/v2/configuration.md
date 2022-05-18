# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/config/v2/configuration.proto](#onos_config_v2_configuration-proto)
    - [AppliedConfigurationStatus](#onos-config-v2-AppliedConfigurationStatus)
    - [AppliedConfigurationStatus.ValuesEntry](#onos-config-v2-AppliedConfigurationStatus-ValuesEntry)
    - [CommittedConfigurationStatus](#onos-config-v2-CommittedConfigurationStatus)
    - [Configuration](#onos-config-v2-Configuration)
    - [Configuration.ValuesEntry](#onos-config-v2-Configuration-ValuesEntry)
    - [ConfigurationEvent](#onos-config-v2-ConfigurationEvent)
    - [ConfigurationStatus](#onos-config-v2-ConfigurationStatus)
    - [MastershipInfo](#onos-config-v2-MastershipInfo)
    - [ProposedConfigurationStatus](#onos-config-v2-ProposedConfigurationStatus)
  
    - [ConfigurationEvent.EventType](#onos-config-v2-ConfigurationEvent-EventType)
    - [ConfigurationStatus.State](#onos-config-v2-ConfigurationStatus-State)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos_config_v2_configuration-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/config/v2/configuration.proto



<a name="onos-config-v2-AppliedConfigurationStatus"></a>

### AppliedConfigurationStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| index | [uint64](#uint64) |  |  |
| mastership | [MastershipInfo](#onos-config-v2-MastershipInfo) |  |  |
| values | [AppliedConfigurationStatus.ValuesEntry](#onos-config-v2-AppliedConfigurationStatus-ValuesEntry) | repeated |  |






<a name="onos-config-v2-AppliedConfigurationStatus-ValuesEntry"></a>

### AppliedConfigurationStatus.ValuesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [PathValue](#onos-config-v2-PathValue) |  |  |






<a name="onos-config-v2-CommittedConfigurationStatus"></a>

### CommittedConfigurationStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| index | [uint64](#uint64) |  |  |






<a name="onos-config-v2-Configuration"></a>

### Configuration
Configuration represents complete desired target configuration


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meta | [ObjectMeta](#onos-config-v2-ObjectMeta) |  |  |
| id | [string](#string) |  | &#39;id&#39; is a unique configuration identifier |
| target_id | [string](#string) |  | &#39;target_id&#39; is the target to which the desired target configuration applies |
| values | [Configuration.ValuesEntry](#onos-config-v2-Configuration-ValuesEntry) | repeated | &#39;values&#39; is a map of path/values to set |
| index | [uint64](#uint64) |  | &#39;index&#39; is the index of the configuration values |
| status | [ConfigurationStatus](#onos-config-v2-ConfigurationStatus) |  | &#39;ConfigurationStatus&#39; is the current lifecycle status of the configuration |
| target_info | [TargetTypeVersion](#onos-config-v2-TargetTypeVersion) |  |  |






<a name="onos-config-v2-Configuration-ValuesEntry"></a>

### Configuration.ValuesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [PathValue](#onos-config-v2-PathValue) |  |  |






<a name="onos-config-v2-ConfigurationEvent"></a>

### ConfigurationEvent
ConfigurationEvent configuration store event


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [ConfigurationEvent.EventType](#onos-config-v2-ConfigurationEvent-EventType) |  | EventType configuration event type |
| configuration | [Configuration](#onos-config-v2-Configuration) |  |  |






<a name="onos-config-v2-ConfigurationStatus"></a>

### ConfigurationStatus
ConfigurationStatus is the status of a Configuration


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| state | [ConfigurationStatus.State](#onos-config-v2-ConfigurationStatus-State) |  | &#39;state&#39; is the configuration state |
| mastership | [MastershipInfo](#onos-config-v2-MastershipInfo) |  | &#39;mastership&#39; is the current mastership info for the configuration |
| proposed | [ProposedConfigurationStatus](#onos-config-v2-ProposedConfigurationStatus) |  | &#39;proposed&#39; is the proposed configuration status |
| committed | [CommittedConfigurationStatus](#onos-config-v2-CommittedConfigurationStatus) |  | &#39;committed&#39; is the committed configuration status |
| applied | [AppliedConfigurationStatus](#onos-config-v2-AppliedConfigurationStatus) |  | &#39;applied&#39; is the applied configuration status |






<a name="onos-config-v2-MastershipInfo"></a>

### MastershipInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| master | [string](#string) |  |  |
| term | [uint64](#uint64) |  |  |






<a name="onos-config-v2-ProposedConfigurationStatus"></a>

### ProposedConfigurationStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| index | [uint64](#uint64) |  |  |





 


<a name="onos-config-v2-ConfigurationEvent-EventType"></a>

### ConfigurationEvent.EventType
EventType configuration event types for configuration store

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 | UNKNOWN indicates unknown configuration store event |
| CREATED | 1 | CREATED indicates the configuration entry in the store is created |
| UPDATED | 2 | UPDATED indicates the configuration entry in the store is updated |
| DELETED | 3 | DELETED indicates the configuration entry in the store is deleted |
| REPLAYED | 4 | REPLAYED |



<a name="onos-config-v2-ConfigurationStatus-State"></a>

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

