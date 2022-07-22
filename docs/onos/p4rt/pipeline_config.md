# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/p4rt/v1/pipeline_config.proto](#onos_p4rt_v1_pipeline_config-proto)
    - [ConfigurationEvent](#onos-p4rt-v1-ConfigurationEvent)
    - [PipelineConfig](#onos-p4rt-v1-PipelineConfig)
    - [PipelineConfigStatus](#onos-p4rt-v1-PipelineConfigStatus)
  
    - [ConfigurationAction](#onos-p4rt-v1-ConfigurationAction)
    - [ConfigurationEvent.EventType](#onos-p4rt-v1-ConfigurationEvent-EventType)
    - [PipelineConfigStatus.State](#onos-p4rt-v1-PipelineConfigStatus-State)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos_p4rt_v1_pipeline_config-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/p4rt/v1/pipeline_config.proto



<a name="onos-p4rt-v1-ConfigurationEvent"></a>

### ConfigurationEvent
ConfigurationEvent configuration store event


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [ConfigurationEvent.EventType](#onos-p4rt-v1-ConfigurationEvent-EventType) |  | EventType configuration event type |
| pipeline_config | [PipelineConfig](#onos-p4rt-v1-PipelineConfig) |  |  |






<a name="onos-p4rt-v1-PipelineConfig"></a>

### PipelineConfig



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meta | [ObjectMeta](#onos-p4rt-v1-ObjectMeta) |  |  |
| id | [string](#string) |  | &#39;id&#39; is a unique configuration identifier |
| target_id | [string](#string) |  | &#39;target_id&#39; is the target to which the desired target configuration applies |
| p4_device_config | [bytes](#bytes) |  | &#39;p4_device_config&#39; P4 device configuration bytes |
| action | [ConfigurationAction](#onos-p4rt-v1-ConfigurationAction) |  | &#39;ConfigurationAction&#39; |
| status | [PipelineConfigStatus](#onos-p4rt-v1-PipelineConfigStatus) |  | &#39;ConfigurationStatus&#39; is the current lifecycle status of the configuration |






<a name="onos-p4rt-v1-PipelineConfigStatus"></a>

### PipelineConfigStatus
PipelineConfigStatus


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| state | [PipelineConfigStatus.State](#onos-p4rt-v1-PipelineConfigStatus-State) |  |  |





 


<a name="onos-p4rt-v1-ConfigurationAction"></a>

### ConfigurationAction
ConfigurationAction

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNSPECIFIED | 0 |  |
| VERIFY | 1 | Verifies that the target can realize the given config. The forwarding state in the target is not modified. |
| VERIFY_AND_SAVE | 2 | Saves the config if the P4Runtime target can realize it. The forwarding state in the target is not modified. |
| VERIFY_AND_COMMIT | 3 | Saves and realizes the given config if the P4Runtime target can realize it. The forwarding state in the target is cleared. |
| COMMIT | 4 | Realizes the last saved, but not yet committed, config. The forwarding state in the target is updated by replaying the write requests to the target device since the last config was saved |
| RECONCILE_AND_COMMIT | 5 | Verifies, saves and realizes the given config, while preserving the forwarding state in the target. |



<a name="onos-p4rt-v1-ConfigurationEvent-EventType"></a>

### ConfigurationEvent.EventType
EventType configuration event types for configuration store

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 | UNKNOWN indicates unknown configuration store event |
| CREATED | 1 | CREATED indicates the configuration entry in the store is created |
| UPDATED | 2 | UPDATED indicates the configuration entry in the store is updated |
| DELETED | 3 | DELETED indicates the configuration entry in the store is deleted |
| REPLAYED | 4 | REPLAYED |



<a name="onos-p4rt-v1-PipelineConfigStatus-State"></a>

### PipelineConfigStatus.State
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

