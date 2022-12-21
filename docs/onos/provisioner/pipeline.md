# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/provisioner/pipeline.proto](#onos_provisioner_pipeline-proto)
    - [AddPipelineConfigRequest](#onos-provisioner-AddPipelineConfigRequest)
    - [AddPipelineConfigResponse](#onos-provisioner-AddPipelineConfigResponse)
    - [DeletePipelineConfigRequest](#onos-provisioner-DeletePipelineConfigRequest)
    - [DeletePipelineConfigResponse](#onos-provisioner-DeletePipelineConfigResponse)
    - [GetPipelineConfigRequest](#onos-provisioner-GetPipelineConfigRequest)
    - [GetPipelineConfigResponse](#onos-provisioner-GetPipelineConfigResponse)
    - [ListPipelineConfigsRequest](#onos-provisioner-ListPipelineConfigsRequest)
    - [ListPipelineConfigsResponse](#onos-provisioner-ListPipelineConfigsResponse)
    - [PipelineConfig](#onos-provisioner-PipelineConfig)
  
    - [PipelineConfigService](#onos-provisioner-PipelineConfigService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos_provisioner_pipeline-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/provisioner/pipeline.proto



<a name="onos-provisioner-AddPipelineConfigRequest"></a>

### AddPipelineConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [PipelineConfig](#onos-provisioner-PipelineConfig) |  |  |






<a name="onos-provisioner-AddPipelineConfigResponse"></a>

### AddPipelineConfigResponse







<a name="onos-provisioner-DeletePipelineConfigRequest"></a>

### DeletePipelineConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pipeline_config | [PipelineConfig](#onos-provisioner-PipelineConfig) |  | config_id and revision must be present; all others optional |






<a name="onos-provisioner-DeletePipelineConfigResponse"></a>

### DeletePipelineConfigResponse







<a name="onos-provisioner-GetPipelineConfigRequest"></a>

### GetPipelineConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config_id | [string](#string) |  |  |






<a name="onos-provisioner-GetPipelineConfigResponse"></a>

### GetPipelineConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [PipelineConfig](#onos-provisioner-PipelineConfig) |  |  |






<a name="onos-provisioner-ListPipelineConfigsRequest"></a>

### ListPipelineConfigsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| watch | [bool](#bool) |  |  |






<a name="onos-provisioner-ListPipelineConfigsResponse"></a>

### ListPipelineConfigsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [PipelineConfig](#onos-provisioner-PipelineConfig) |  |  |






<a name="onos-provisioner-PipelineConfig"></a>

### PipelineConfig
PipelineConfig


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config_id | [string](#string) |  |  |
| revision | [uint64](#uint64) |  |  |
| p4_info | [bytes](#bytes) |  |  |
| p4_binary | [bytes](#bytes) |  |  |





 

 

 


<a name="onos-provisioner-PipelineConfigService"></a>

### PipelineConfigService
PipelineConfigService

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Add | [AddPipelineConfigRequest](#onos-provisioner-AddPipelineConfigRequest) | [AddPipelineConfigResponse](#onos-provisioner-AddPipelineConfigResponse) | Add registers new pipeline configuration |
| Delete | [DeletePipelineConfigRequest](#onos-provisioner-DeletePipelineConfigRequest) | [DeletePipelineConfigResponse](#onos-provisioner-DeletePipelineConfigResponse) | Delete removes a pipeline configuration |
| Get | [GetPipelineConfigRequest](#onos-provisioner-GetPipelineConfigRequest) | [GetPipelineConfigResponse](#onos-provisioner-GetPipelineConfigResponse) | Get returns pipeline configuration based on a given ID |
| List | [ListPipelineConfigsRequest](#onos-provisioner-ListPipelineConfigsRequest) | [ListPipelineConfigsResponse](#onos-provisioner-ListPipelineConfigsResponse) stream | List returns all registered pipelines |

 



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

