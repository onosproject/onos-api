# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/provisioner/aspects.proto](#onos_provisioner_aspects-proto)
    - [ChassisConfigState](#onos-provisioner-ChassisConfigState)
    - [DeviceConfig](#onos-provisioner-DeviceConfig)
    - [PipelineConfigState](#onos-provisioner-PipelineConfigState)
  
- [onos/provisioner/config.proto](#onos_provisioner_config-proto)
    - [AddConfigRequest](#onos-provisioner-AddConfigRequest)
    - [AddConfigResponse](#onos-provisioner-AddConfigResponse)
    - [Config](#onos-provisioner-Config)
    - [Config.ArtifactsEntry](#onos-provisioner-Config-ArtifactsEntry)
    - [ConfigRecord](#onos-provisioner-ConfigRecord)
    - [DeleteConfigRequest](#onos-provisioner-DeleteConfigRequest)
    - [DeleteConfigResponse](#onos-provisioner-DeleteConfigResponse)
    - [GetConfigRequest](#onos-provisioner-GetConfigRequest)
    - [GetConfigResponse](#onos-provisioner-GetConfigResponse)
    - [ListConfigsRequest](#onos-provisioner-ListConfigsRequest)
    - [ListConfigsResponse](#onos-provisioner-ListConfigsResponse)
  
    - [ProvisionerService](#onos-provisioner-ProvisionerService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos_provisioner_aspects-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/provisioner/aspects.proto



<a name="onos-provisioner-ChassisConfigState"></a>

### ChassisConfigState
ChassisConfigState is a topology entity aspect used to indicate what chassis config a device has presently applied to it


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config_id | [string](#string) |  |  |
| updated | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="onos-provisioner-DeviceConfig"></a>

### DeviceConfig
DeviceConfig is a topology entity aspect used to specify what pipeline and chassis config a device should have applied to it


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pipeline_config_id | [string](#string) |  |  |
| chassis_config_id | [string](#string) |  |  |






<a name="onos-provisioner-PipelineConfigState"></a>

### PipelineConfigState
PipelineConfigState is a topology entity aspect used to indicate what pipeline config a device has presently applied to it


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config_id | [string](#string) |  |  |
| cookie | [uint64](#uint64) |  |  |
| updated | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |





 

 

 

 



<a name="onos_provisioner_config-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/provisioner/config.proto



<a name="onos-provisioner-AddConfigRequest"></a>

### AddConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [Config](#onos-provisioner-Config) |  |  |






<a name="onos-provisioner-AddConfigResponse"></a>

### AddConfigResponse







<a name="onos-provisioner-Config"></a>

### Config
Config represents


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| record | [ConfigRecord](#onos-provisioner-ConfigRecord) |  |  |
| artifacts | [Config.ArtifactsEntry](#onos-provisioner-Config-ArtifactsEntry) | repeated |  |






<a name="onos-provisioner-Config-ArtifactsEntry"></a>

### Config.ArtifactsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [bytes](#bytes) |  |  |






<a name="onos-provisioner-ConfigRecord"></a>

### ConfigRecord
ConfigRecord is used for storing a record of the pipeline or chassis configuration with
references to its related file artifact(s)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config_id | [string](#string) |  |  |
| kind | [string](#string) |  | pipeline, chassis, etc. |
| artifacts | [string](#string) | repeated | list of associated artifact types |






<a name="onos-provisioner-DeleteConfigRequest"></a>

### DeleteConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config_id | [string](#string) |  |  |






<a name="onos-provisioner-DeleteConfigResponse"></a>

### DeleteConfigResponse







<a name="onos-provisioner-GetConfigRequest"></a>

### GetConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config_id | [string](#string) |  |  |
| include_artifacts | [bool](#bool) |  |  |






<a name="onos-provisioner-GetConfigResponse"></a>

### GetConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [Config](#onos-provisioner-Config) |  |  |






<a name="onos-provisioner-ListConfigsRequest"></a>

### ListConfigsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kind | [string](#string) |  |  |
| include_artifacts | [bool](#bool) |  |  |
| watch | [bool](#bool) |  |  |






<a name="onos-provisioner-ListConfigsResponse"></a>

### ListConfigsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [Config](#onos-provisioner-Config) |  |  |





 

 

 


<a name="onos-provisioner-ProvisionerService"></a>

### ProvisionerService
ProvisionerService allows managing inventory of various device configuratoins, e.g., pipeline, chassis.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Add | [AddConfigRequest](#onos-provisioner-AddConfigRequest) | [AddConfigResponse](#onos-provisioner-AddConfigResponse) | Add registers new pipeline configuration |
| Delete | [DeleteConfigRequest](#onos-provisioner-DeleteConfigRequest) | [DeleteConfigResponse](#onos-provisioner-DeleteConfigResponse) | Delete removes a pipeline configuration |
| Get | [GetConfigRequest](#onos-provisioner-GetConfigRequest) | [GetConfigResponse](#onos-provisioner-GetConfigResponse) | Get returns pipeline configuration based on a given ID |
| List | [ListConfigsRequest](#onos-provisioner-ListConfigsRequest) | [ListConfigsResponse](#onos-provisioner-ListConfigsResponse) stream | List returns all registered pipelines |

 



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

