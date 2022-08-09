# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/p4rt/v1/pipeline_config.proto](#onos_p4rt_v1_pipeline_config-proto)
    - [ConfigurationEvent](#onos-p4rt-v1-ConfigurationEvent)
    - [Cookie](#onos-p4rt-v1-Cookie)
    - [MastershipInfo](#onos-p4rt-v1-MastershipInfo)
    - [PipelineConfig](#onos-p4rt-v1-PipelineConfig)
    - [PipelineConfigSpec](#onos-p4rt-v1-PipelineConfigSpec)
    - [PipelineConfigStatus](#onos-p4rt-v1-PipelineConfigStatus)
  
    - [ConfigurationAction](#onos-p4rt-v1-ConfigurationAction)
    - [ConfigurationEvent.Type](#onos-p4rt-v1-ConfigurationEvent-Type)
    - [PipelineConfigStatus.PipelineConfigState](#onos-p4rt-v1-PipelineConfigStatus-PipelineConfigState)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos_p4rt_v1_pipeline_config-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/p4rt/v1/pipeline_config.proto



<a name="onos-p4rt-v1-ConfigurationEvent"></a>

### ConfigurationEvent
ConfigurationEvent configuration store event


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [ConfigurationEvent.Type](#onos-p4rt-v1-ConfigurationEvent-Type) |  | EventType configuration event type |
| pipeline_config | [PipelineConfig](#onos-p4rt-v1-PipelineConfig) |  |  |






<a name="onos-p4rt-v1-Cookie"></a>

### Cookie
Cookie


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cookie | [uint64](#uint64) |  |  |






<a name="onos-p4rt-v1-MastershipInfo"></a>

### MastershipInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| master | [string](#string) |  |  |
| term | [uint64](#uint64) |  |  |






<a name="onos-p4rt-v1-PipelineConfig"></a>

### PipelineConfig
PipelineConfig P4 device pipeline config


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meta | [ObjectMeta](#onos-p4rt-v1-ObjectMeta) |  |  |
| id | [string](#string) |  | &#39;id&#39; is a unique configuration identifier |
| target_id | [string](#string) |  | &#39;target_id&#39; is the target to which the desired target configuration applies |
| cookie | [Cookie](#onos-p4rt-v1-Cookie) |  | &#39;cookie&#39; to uniquely identify a forwarding-pipeline configuration among others managed by the same control plane |
| status | [PipelineConfigStatus](#onos-p4rt-v1-PipelineConfigStatus) |  | &#39;ConfigurationStatus&#39; is the current lifecycle status of the configuration |
| action | [ConfigurationAction](#onos-p4rt-v1-ConfigurationAction) |  | ConfigurationAction |
| spec | [PipelineConfigSpec](#onos-p4rt-v1-PipelineConfigSpec) |  | PipelineConfigSpec device pipeline config spec |






<a name="onos-p4rt-v1-PipelineConfigSpec"></a>

### PipelineConfigSpec
PipelineConfigSpec device pipeline configuration spec


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| p4_device_config | [bytes](#bytes) |  | &#39;p4_device_config&#39; P4 device configuration bytes |
| p4_info | [bytes](#bytes) |  | TODO since p4 info is target agnostic, we can import it from P4runtime and use it as it is, a new version of proto compiler image is needed &#39;p4_info&#39; P4 info |






<a name="onos-p4rt-v1-PipelineConfigStatus"></a>

### PipelineConfigStatus
PipelineConfigStatus pipelineConfig status


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| state | [PipelineConfigStatus.PipelineConfigState](#onos-p4rt-v1-PipelineConfigStatus-PipelineConfigState) |  | &#39;PipelineConfigState&#39; pipeline config state |
| mastership | [MastershipInfo](#onos-p4rt-v1-MastershipInfo) |  | &#39;mastership&#39; is the current mastership info for the configuration |





 


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



<a name="onos-p4rt-v1-ConfigurationEvent-Type"></a>

### ConfigurationEvent.Type
Type configuration event types for configuration store

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 | UNKNOWN indicates unknown configuration store event |
| CREATED | 1 | CREATED indicates the configuration entry in the store is created |
| UPDATED | 2 | UPDATED indicates the configuration entry in the store is updated |
| DELETED | 3 | DELETED indicates the configuration entry in the store is deleted |
| REPLAYED | 4 | REPLAYED |



<a name="onos-p4rt-v1-PipelineConfigStatus-PipelineConfigState"></a>

### PipelineConfigStatus.PipelineConfigState


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 | UNKNOWN indicates the pipeline configuration state is unknown |
| PENDING | 1 | PENDING indicates the pipeline configuration state is pending |
| COMPLETE | 2 | COMPLETE indicates the pipeline configuration state is complete |
| FAILED | 3 | FAILED indicates the pipeline configuration state failed |


 

 

 



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

