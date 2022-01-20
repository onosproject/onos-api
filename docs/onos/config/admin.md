# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/config/admin/configuration.proto](#onos/config/admin/configuration.proto)
    - [ConfigurationEvent](#onos.config.admin.ConfigurationEvent)
    - [GetConfigurationRequest](#onos.config.admin.GetConfigurationRequest)
    - [GetConfigurationResponse](#onos.config.admin.GetConfigurationResponse)
    - [ListConfigurationsRequest](#onos.config.admin.ListConfigurationsRequest)
    - [ListConfigurationsResponse](#onos.config.admin.ListConfigurationsResponse)
    - [WatchConfigurationsRequest](#onos.config.admin.WatchConfigurationsRequest)
    - [WatchConfigurationsResponse](#onos.config.admin.WatchConfigurationsResponse)
  
    - [ConfigurationService](#onos.config.admin.ConfigurationService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos/config/admin/configuration.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/config/admin/configuration.proto



<a name="onos.config.admin.ConfigurationEvent"></a>

### ConfigurationEvent
Event is a topo operation event


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [Type](#onos.config.admin.Type) |  |  |
| configuration | [onos.config.v2.Configuration](#onos.config.v2.Configuration) |  |  |






<a name="onos.config.admin.GetConfigurationRequest"></a>

### GetConfigurationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="onos.config.admin.GetConfigurationResponse"></a>

### GetConfigurationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction | [onos.config.v2.Configuration](#onos.config.v2.Configuration) |  |  |






<a name="onos.config.admin.ListConfigurationsRequest"></a>

### ListConfigurationsRequest







<a name="onos.config.admin.ListConfigurationsResponse"></a>

### ListConfigurationsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction | [onos.config.v2.Configuration](#onos.config.v2.Configuration) |  |  |






<a name="onos.config.admin.WatchConfigurationsRequest"></a>

### WatchConfigurationsRequest







<a name="onos.config.admin.WatchConfigurationsResponse"></a>

### WatchConfigurationsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [ConfigurationEvent](#onos.config.admin.ConfigurationEvent) |  |  |





 

 

 


<a name="onos.config.admin.ConfigurationService"></a>

### ConfigurationService
ConfigurationService provides means to inspect the contents of the internal configurations store.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetConfigurationRequest](#onos.config.admin.GetConfigurationRequest) | [GetConfigurationResponse](#onos.config.admin.GetConfigurationResponse) | Get configuration by its target ID |
| List | [ListConfigurationsRequest](#onos.config.admin.ListConfigurationsRequest) | [ListConfigurationsResponse](#onos.config.admin.ListConfigurationsResponse) stream | List returns all target configurations |
| Watch | [WatchConfigurationsRequest](#onos.config.admin.WatchConfigurationsRequest) | [WatchConfigurationsResponse](#onos.config.admin.WatchConfigurationsResponse) stream | Watch returns a stream of configuration change notifications |

 



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

