# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/configmodel/registry.proto](#onos/configmodel/registry.proto)
    - [ConfigModel](#onos.configmodel.ConfigModel)
    - [ConfigModel.FilesEntry](#onos.configmodel.ConfigModel.FilesEntry)
    - [ConfigModule](#onos.configmodel.ConfigModule)
    - [DeleteModelRequest](#onos.configmodel.DeleteModelRequest)
    - [DeleteModelResponse](#onos.configmodel.DeleteModelResponse)
    - [GetModelRequest](#onos.configmodel.GetModelRequest)
    - [GetModelResponse](#onos.configmodel.GetModelResponse)
    - [ListModelsRequest](#onos.configmodel.ListModelsRequest)
    - [ListModelsResponse](#onos.configmodel.ListModelsResponse)
    - [PushModelRequest](#onos.configmodel.PushModelRequest)
    - [PushModelResponse](#onos.configmodel.PushModelResponse)
  
    - [GetStateMode](#onos.configmodel.GetStateMode)
  
    - [ConfigModelRegistryService](#onos.configmodel.ConfigModelRegistryService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos/configmodel/registry.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/configmodel/registry.proto



<a name="onos.configmodel.ConfigModel"></a>

### ConfigModel



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| version | [string](#string) |  |  |
| modules | [ConfigModule](#onos.configmodel.ConfigModule) | repeated |  |
| files | [ConfigModel.FilesEntry](#onos.configmodel.ConfigModel.FilesEntry) | repeated |  |
| get_state_mode | [GetStateMode](#onos.configmodel.GetStateMode) |  |  |






<a name="onos.configmodel.ConfigModel.FilesEntry"></a>

### ConfigModel.FilesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="onos.configmodel.ConfigModule"></a>

### ConfigModule



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| file | [string](#string) |  |  |
| revision | [string](#string) |  |  |
| organization | [string](#string) |  |  |






<a name="onos.configmodel.DeleteModelRequest"></a>

### DeleteModelRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| version | [string](#string) |  |  |






<a name="onos.configmodel.DeleteModelResponse"></a>

### DeleteModelResponse







<a name="onos.configmodel.GetModelRequest"></a>

### GetModelRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| version | [string](#string) |  |  |






<a name="onos.configmodel.GetModelResponse"></a>

### GetModelResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| model | [ConfigModel](#onos.configmodel.ConfigModel) |  |  |






<a name="onos.configmodel.ListModelsRequest"></a>

### ListModelsRequest







<a name="onos.configmodel.ListModelsResponse"></a>

### ListModelsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| models | [ConfigModel](#onos.configmodel.ConfigModel) | repeated |  |






<a name="onos.configmodel.PushModelRequest"></a>

### PushModelRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| model | [ConfigModel](#onos.configmodel.ConfigModel) |  |  |






<a name="onos.configmodel.PushModelResponse"></a>

### PushModelResponse






 


<a name="onos.configmodel.GetStateMode"></a>

### GetStateMode


| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 |  |
| OP_STATE | 1 |  |
| EXPLICIT_RO_PATHS | 2 |  |
| EXPLICIT_RO_PATHS_EXPAND_WILDCARDS | 3 |  |


 

 


<a name="onos.configmodel.ConfigModelRegistryService"></a>

### ConfigModelRegistryService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetModel | [GetModelRequest](#onos.configmodel.GetModelRequest) | [GetModelResponse](#onos.configmodel.GetModelResponse) |  |
| ListModels | [ListModelsRequest](#onos.configmodel.ListModelsRequest) | [ListModelsResponse](#onos.configmodel.ListModelsResponse) |  |
| PushModel | [PushModelRequest](#onos.configmodel.PushModelRequest) | [PushModelResponse](#onos.configmodel.PushModelResponse) |  |
| DeleteModel | [DeleteModelRequest](#onos.configmodel.DeleteModelRequest) | [DeleteModelResponse](#onos.configmodel.DeleteModelResponse) |  |

 



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

