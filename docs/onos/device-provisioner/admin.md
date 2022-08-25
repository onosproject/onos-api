# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/device-provisioner/admin/admin.proto](#onos_device-provisioner_admin_admin-proto)
    - [GetPipelineRequest](#onos-deviceprovisioner-admin-GetPipelineRequest)
    - [GetPipelineResponse](#onos-deviceprovisioner-admin-GetPipelineResponse)
    - [ListPipelinesRequest](#onos-deviceprovisioner-admin-ListPipelinesRequest)
    - [ListPipelinesResponse](#onos-deviceprovisioner-admin-ListPipelinesResponse)
    - [WatchPipelinesRequest](#onos-deviceprovisioner-admin-WatchPipelinesRequest)
    - [WatchPipelinesResponse](#onos-deviceprovisioner-admin-WatchPipelinesResponse)
  
    - [PipelineConfigService](#onos-deviceprovisioner-admin-PipelineConfigService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos_device-provisioner_admin_admin-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/device-provisioner/admin/admin.proto



<a name="onos-deviceprovisioner-admin-GetPipelineRequest"></a>

### GetPipelineRequest
GetPipelineRequest get pipeline request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pipelineconfig_id | [string](#string) |  |  |






<a name="onos-deviceprovisioner-admin-GetPipelineResponse"></a>

### GetPipelineResponse
GetPipelineResponse get pipeline response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pipelineconfig | [onos.p4rt.v1.PipelineConfig](#onos-p4rt-v1-PipelineConfig) |  |  |






<a name="onos-deviceprovisioner-admin-ListPipelinesRequest"></a>

### ListPipelinesRequest







<a name="onos-deviceprovisioner-admin-ListPipelinesResponse"></a>

### ListPipelinesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pipelineconfig | [onos.p4rt.v1.PipelineConfig](#onos-p4rt-v1-PipelineConfig) |  |  |






<a name="onos-deviceprovisioner-admin-WatchPipelinesRequest"></a>

### WatchPipelinesRequest







<a name="onos-deviceprovisioner-admin-WatchPipelinesResponse"></a>

### WatchPipelinesResponse






 

 

 


<a name="onos-deviceprovisioner-admin-PipelineConfigService"></a>

### PipelineConfigService
PipelineConfigService

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetPipeline | [GetPipelineRequest](#onos-deviceprovisioner-admin-GetPipelineRequest) | [GetPipelineResponse](#onos-deviceprovisioner-admin-GetPipelineResponse) | Get pipeline config based on a given ID |
| ListPipelines | [ListPipelinesRequest](#onos-deviceprovisioner-admin-ListPipelinesRequest) | [ListPipelinesResponse](#onos-deviceprovisioner-admin-ListPipelinesResponse) stream | List returns all target pipelines |
| WatchPipelines | [WatchPipelinesRequest](#onos-deviceprovisioner-admin-WatchPipelinesRequest) | [WatchPipelinesResponse](#onos-deviceprovisioner-admin-WatchPipelinesResponse) stream | Watch returns a stream of pipeline change notifications |

 



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

