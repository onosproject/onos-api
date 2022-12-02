# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/config/admin/admin.proto](#onos_config_admin_admin-proto)
    - [GetConfigurationRequest](#onos-config-admin-GetConfigurationRequest)
    - [GetConfigurationResponse](#onos-config-admin-GetConfigurationResponse)
    - [GetTransactionRequest](#onos-config-admin-GetTransactionRequest)
    - [GetTransactionResponse](#onos-config-admin-GetTransactionResponse)
    - [LeafSelectionQueryRequest](#onos-config-admin-LeafSelectionQueryRequest)
    - [LeafSelectionQueryResponse](#onos-config-admin-LeafSelectionQueryResponse)
    - [ListConfigurationsRequest](#onos-config-admin-ListConfigurationsRequest)
    - [ListConfigurationsResponse](#onos-config-admin-ListConfigurationsResponse)
    - [ListModelsRequest](#onos-config-admin-ListModelsRequest)
    - [ListTransactionsRequest](#onos-config-admin-ListTransactionsRequest)
    - [ListTransactionsResponse](#onos-config-admin-ListTransactionsResponse)
    - [ModelInfo](#onos-config-admin-ModelInfo)
    - [ModelInfoRequest](#onos-config-admin-ModelInfoRequest)
    - [ModelInfoResponse](#onos-config-admin-ModelInfoResponse)
    - [ModelPlugin](#onos-config-admin-ModelPlugin)
    - [Namespace](#onos-config-admin-Namespace)
    - [PathValuesRequest](#onos-config-admin-PathValuesRequest)
    - [PathValuesResponse](#onos-config-admin-PathValuesResponse)
    - [ReadOnlyPath](#onos-config-admin-ReadOnlyPath)
    - [ReadOnlySubPath](#onos-config-admin-ReadOnlySubPath)
    - [ReadWritePath](#onos-config-admin-ReadWritePath)
    - [RollbackRequest](#onos-config-admin-RollbackRequest)
    - [RollbackResponse](#onos-config-admin-RollbackResponse)
    - [ValidateConfigRequest](#onos-config-admin-ValidateConfigRequest)
    - [ValidateConfigResponse](#onos-config-admin-ValidateConfigResponse)
    - [ValueSelectionRequest](#onos-config-admin-ValueSelectionRequest)
    - [ValueSelectionResponse](#onos-config-admin-ValueSelectionResponse)
    - [WatchConfigurationsRequest](#onos-config-admin-WatchConfigurationsRequest)
    - [WatchConfigurationsResponse](#onos-config-admin-WatchConfigurationsResponse)
    - [WatchTransactionsRequest](#onos-config-admin-WatchTransactionsRequest)
    - [WatchTransactionsResponse](#onos-config-admin-WatchTransactionsResponse)
  
    - [ConfigAdminService](#onos-config-admin-ConfigAdminService)
    - [ConfigurationService](#onos-config-admin-ConfigurationService)
    - [ModelPluginService](#onos-config-admin-ModelPluginService)
    - [TransactionService](#onos-config-admin-TransactionService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos_config_admin_admin-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/config/admin/admin.proto



<a name="onos-config-admin-GetConfigurationRequest"></a>

### GetConfigurationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| configuration_id | [string](#string) |  |  |






<a name="onos-config-admin-GetConfigurationResponse"></a>

### GetConfigurationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| configuration | [onos.config.v2.Configuration](#onos-config-v2-Configuration) |  |  |






<a name="onos-config-admin-GetTransactionRequest"></a>

### GetTransactionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | ID of transaction to get |
| index | [uint64](#uint64) |  | index of transaction to get; leave 0 for lookup by ID; if specified takes precedence |






<a name="onos-config-admin-GetTransactionResponse"></a>

### GetTransactionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction | [onos.config.v2.Transaction](#onos-config-v2-Transaction) |  |  |






<a name="onos-config-admin-LeafSelectionQueryRequest"></a>

### LeafSelectionQueryRequest
LeafSelectionQueryRequest carries request for the selection of leaf values


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| target | [string](#string) |  | target is the name of the target (device) to perform the query on |
| selectionPath | [string](#string) |  | selectionPath is a configuration path to a leaf in the format: /a/b[key1=index][key2=index2]/c/d where d is a leaf node |
| changeContext | [gnmi.SetRequest](#gnmi-SetRequest) |  | changeContext is the set of changes from the GUI form that have to be superimposed on the current configuration before the leaf selection can be made All the changes in this should be for the same target |






<a name="onos-config-admin-LeafSelectionQueryResponse"></a>

### LeafSelectionQueryResponse
LeafSelectionQueryResponse carries response for the model information query


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| selection | [string](#string) | repeated | selection is an array of string values |






<a name="onos-config-admin-ListConfigurationsRequest"></a>

### ListConfigurationsRequest







<a name="onos-config-admin-ListConfigurationsResponse"></a>

### ListConfigurationsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| configuration | [onos.config.v2.Configuration](#onos-config-v2-Configuration) |  |  |






<a name="onos-config-admin-ListModelsRequest"></a>

### ListModelsRequest
ListModelsRequest carries data for querying registered model plugins.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| verbose | [bool](#bool) |  | verbose option causes all of the ReadWrite and ReadOnly paths to be included. |
| model_name | [string](#string) |  | An optional filter on the name of the model plugins to list. |
| model_version | [string](#string) |  | An optional filter on the version of the model plugins to list |






<a name="onos-config-admin-ListTransactionsRequest"></a>

### ListTransactionsRequest







<a name="onos-config-admin-ListTransactionsResponse"></a>

### ListTransactionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction | [onos.config.v2.Transaction](#onos-config-v2-Transaction) |  |  |






<a name="onos-config-admin-ModelInfo"></a>

### ModelInfo
ModelInfo is general information about a model plugin.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | name is the name given to the model plugin - no spaces and title case. |
| version | [string](#string) |  | version is the semantic version of the Plugin e.g. 1.0.0. |
| model_data | [gnmi.ModelData](#gnmi-ModelData) | repeated | model_data is a set of metadata about the YANG files that went in to generating the model plugin. It includes name, version and organization for each YANG file, similar to how they are represented in gNMI Capabilities. |
| module | [string](#string) |  | **Deprecated.** module is no longer used |
| getStateMode | [uint32](#uint32) |  | getStateMode is flag that defines how the &#34;get state&#34; operation works. 0) means that no retrieval of state is attempted 1) means that the synchronizer will make 2 requests to the device - one for Get with State and another for Get with Operational. 2) means that the synchronizer will do a Get request comprising of each one of the ReadOnlyPaths and their sub paths. If there is a `list` in any one of these paths it will be sent down as is, expecting the devices implementation of gNMI will be able to expand wildcards. 3) means that the synchronizer will do a Get request comprising of each one of the ReadOnlyPaths and their sub paths. If there is a `list` in any one of these paths, a separate call will be made first to find all the instances in the list and a Get including these expanded wildcards will be sent down to the device. |
| read_only_path | [ReadOnlyPath](#onos-config-admin-ReadOnlyPath) | repeated | read_only_path is all of the read only paths for the model plugin. |
| read_write_path | [ReadWritePath](#onos-config-admin-ReadWritePath) | repeated | read_write_path is all of the read write paths for the model plugin. |
| supported_encodings | [gnmi.Encoding](#gnmi-Encoding) | repeated |  |
| namespace_mappings | [Namespace](#onos-config-admin-Namespace) | repeated | namespace_mappings is a set of all prefix to module name mapping in the model |
| southboundUsePrefix | [bool](#bool) |  | southboundUsePrefix indicates that the southbound should add prefixes in gNMI paths |






<a name="onos-config-admin-ModelInfoRequest"></a>

### ModelInfoRequest
ModelInfoRequest carries request for the model information






<a name="onos-config-admin-ModelInfoResponse"></a>

### ModelInfoResponse
ModelInfoResponse carries response for the model information query


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| modelInfo | [ModelInfo](#onos-config-admin-ModelInfo) |  |  |






<a name="onos-config-admin-ModelPlugin"></a>

### ModelPlugin



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| endpoint | [string](#string) |  |  |
| info | [ModelInfo](#onos-config-admin-ModelInfo) |  |  |
| status | [string](#string) |  |  |
| error | [string](#string) |  |  |






<a name="onos-config-admin-Namespace"></a>

### Namespace
Namespace is a mapping between a module name and its shorthand prefix


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| module | [string](#string) |  |  |
| prefix | [string](#string) |  |  |






<a name="onos-config-admin-PathValuesRequest"></a>

### PathValuesRequest
PathValuesRequest carries configuration change as a JSON blob


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pathPrefix | [string](#string) |  |  |
| json | [bytes](#bytes) |  |  |






<a name="onos-config-admin-PathValuesResponse"></a>

### PathValuesResponse
PathValuesResponse carries a list of typed path values


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pathValues | [onos.config.v2.PathValue](#onos-config-v2-PathValue) | repeated |  |






<a name="onos-config-admin-ReadOnlyPath"></a>

### ReadOnlyPath
ReadOnlyPath extracted from the model plugin as the definition of a tree of read only items.
In YANG models items are defined as ReadOnly with the `config false` keyword.
This can be applied to single items (leafs) or collections (containers or lists).
When this `config false` is applied to an object every item beneath it will
also become readonly - here these are shown as subpaths.
The complete read only path then will be a concatenation of both e.g.
/cont1a/cont1b-state/list2b/index and the type is defined in the SubPath as UInt8.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| path | [string](#string) |  | path of the topmost `config false` object e.g. /cont1a/cont1b-state |
| sub_path | [ReadOnlySubPath](#onos-config-admin-ReadOnlySubPath) | repeated | ReadOnlySubPath is a set of children of the path including an entry for the type of the topmost object with subpath `/` An example is /list2b/index |






<a name="onos-config-admin-ReadOnlySubPath"></a>

### ReadOnlySubPath
ReadOnlySubPath is an extension to the ReadOnlyPath to define the datatype of
the subpath


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sub_path | [string](#string) |  | sub_path is the relative path of a child object e.g. /list2b/index |
| value_type | [onos.config.v2.ValueType](#onos-config-v2-ValueType) |  | value_type is the datatype of the read only path |
| type_opts | [uint64](#uint64) | repeated |  |
| description | [string](#string) |  |  |
| units | [string](#string) |  |  |
| IsAKey | [bool](#bool) |  |  |
| AttrName | [string](#string) |  |  |






<a name="onos-config-admin-ReadWritePath"></a>

### ReadWritePath
ReadWritePath is extracted from the model plugin as the definition of a writeable attributes.
In YANG models items are writable by default unless they are specified as `config false` or
have an item with `config false` as a parent.
Each configurable item has metadata with meanings taken from the YANG specification RFC 6020.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| path | [string](#string) |  | path is the full path to the attribute (leaf or leaf-list) |
| value_type | [onos.config.v2.ValueType](#onos-config-v2-ValueType) |  | value_type is the data type of the attribute |
| units | [string](#string) |  | units is the unit of measurement e.g. dB, mV |
| description | [string](#string) |  | description is an explaination of the meaning of the attribute |
| mandatory | [bool](#bool) |  | mandatory shows whether the attribute is optional (false) or required (true) |
| default | [string](#string) |  | **Deprecated.** default is a default value used with optional attributes. Replaced by &#39;defaults&#39; below |
| range | [string](#string) | repeated | range is definition of the range of values a value is allowed |
| length | [string](#string) | repeated | length is a defintion of the length restrictions for the attribute |
| type_opts | [uint64](#uint64) | repeated |  |
| IsAKey | [bool](#bool) |  |  |
| AttrName | [string](#string) |  |  |
| defaults | [string](#string) | repeated | defaults is a default value(s) used with optional attributes. For leaf-list can have repeated values replaces the &#39;default&#39; attribute above |






<a name="onos-config-admin-RollbackRequest"></a>

### RollbackRequest
RollbackRequest carries the index of the configuration change transaction to rollback.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| index | [uint64](#uint64) |  | index of the transaction that should be rolled back |






<a name="onos-config-admin-RollbackResponse"></a>

### RollbackResponse
RollbackResponse carries the response of the rollback operation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | ID of the rollback transaction |
| index | [uint64](#uint64) |  | index of the rollback transaction |






<a name="onos-config-admin-ValidateConfigRequest"></a>

### ValidateConfigRequest
ValidateConfigRequest carries configuration data to be validated as a JSON blob


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| json | [bytes](#bytes) |  |  |






<a name="onos-config-admin-ValidateConfigResponse"></a>

### ValidateConfigResponse
ValidateConfigResponse carries the result of the validation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| valid | [bool](#bool) |  |  |
| message | [string](#string) |  |  |






<a name="onos-config-admin-ValueSelectionRequest"></a>

### ValueSelectionRequest
ValueSelectionRequest carries the necessary parts to form a selection context


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| selectionPath | [string](#string) |  | selectionPath is a configuration path to a leaf in the format: /a/b[key1=index][key2=index2]/c/d where d is a leaf node |
| configJson | [bytes](#bytes) |  | configJson is a JSON tree view of the complete Configuration for a Target |






<a name="onos-config-admin-ValueSelectionResponse"></a>

### ValueSelectionResponse
ValueSelectionResponse returns the result of applying the selection rules to the selection context


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| selection | [string](#string) | repeated | selection is an array of string values |






<a name="onos-config-admin-WatchConfigurationsRequest"></a>

### WatchConfigurationsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| configuration_id | [string](#string) |  |  |
| noreplay | [bool](#bool) |  |  |






<a name="onos-config-admin-WatchConfigurationsResponse"></a>

### WatchConfigurationsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [onos.config.v2.ConfigurationEvent](#onos-config-v2-ConfigurationEvent) |  |  |






<a name="onos-config-admin-WatchTransactionsRequest"></a>

### WatchTransactionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| noreplay | [bool](#bool) |  |  |






<a name="onos-config-admin-WatchTransactionsResponse"></a>

### WatchTransactionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [onos.config.v2.TransactionEvent](#onos-config-v2-TransactionEvent) |  |  |





 

 

 


<a name="onos-config-admin-ConfigAdminService"></a>

### ConfigAdminService
ConfigAdminService provides means for enhanced interactions with the configuration subsystem.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListRegisteredModels | [ListModelsRequest](#onos-config-admin-ListModelsRequest) | [ModelPlugin](#onos-config-admin-ModelPlugin) stream | ListRegisteredModels returns a stream of registered models. |
| RollbackTransaction | [RollbackRequest](#onos-config-admin-RollbackRequest) | [RollbackResponse](#onos-config-admin-RollbackResponse) | RollbackTransaction rolls back the specified configuration change transaction. |
| LeafSelectionQuery | [LeafSelectionQueryRequest](#onos-config-admin-LeafSelectionQueryRequest) | [LeafSelectionQueryResponse](#onos-config-admin-LeafSelectionQueryResponse) | LeafSelectionQuery selects values allowable for leaf. It supports the ROC GUI by supplying a list of valid leaf values based off an XPath query defined in a &#39;leaf-selection&#39; YANG extension Calls on GetValueSelection RPC on Model Plugin |


<a name="onos-config-admin-ConfigurationService"></a>

### ConfigurationService
ConfigurationService provides means to inspect the contents of the internal configurations store.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetConfiguration | [GetConfigurationRequest](#onos-config-admin-GetConfigurationRequest) | [GetConfigurationResponse](#onos-config-admin-GetConfigurationResponse) | Get configuration by its target ID |
| ListConfigurations | [ListConfigurationsRequest](#onos-config-admin-ListConfigurationsRequest) | [ListConfigurationsResponse](#onos-config-admin-ListConfigurationsResponse) stream | List returns all target configurations |
| WatchConfigurations | [WatchConfigurationsRequest](#onos-config-admin-WatchConfigurationsRequest) | [WatchConfigurationsResponse](#onos-config-admin-WatchConfigurationsResponse) stream | Watch returns a stream of configuration change notifications |


<a name="onos-config-admin-ModelPluginService"></a>

### ModelPluginService
ModelPluginService is to be implemented by model plugin sidecar

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetModelInfo | [ModelInfoRequest](#onos-config-admin-ModelInfoRequest) | [ModelInfoResponse](#onos-config-admin-ModelInfoResponse) | GetModelInfo provides information about the model |
| ValidateConfig | [ValidateConfigRequest](#onos-config-admin-ValidateConfigRequest) | [ValidateConfigResponse](#onos-config-admin-ValidateConfigResponse) | ValidateConfig validates the provided configuration data against the model |
| GetPathValues | [PathValuesRequest](#onos-config-admin-PathValuesRequest) | [PathValuesResponse](#onos-config-admin-PathValuesResponse) | GetPathValues produces list of typed path value entries from the specified configuration change JSON tree |
| GetValueSelection | [ValueSelectionRequest](#onos-config-admin-ValueSelectionRequest) | [ValueSelectionResponse](#onos-config-admin-ValueSelectionResponse) | GetValueSelection gets a list of valid options for a leaf by applying selection rules in YANG. The selection rules should be defined as an XPath expression, as an argument to a `leaf-selection` extension in the YANG model (Used to support the ROC GUI) |


<a name="onos-config-admin-TransactionService"></a>

### TransactionService
TransactionService provides means to inspect the contents of the internal transactions store.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetTransaction | [GetTransactionRequest](#onos-config-admin-GetTransactionRequest) | [GetTransactionResponse](#onos-config-admin-GetTransactionResponse) | Get transaction by its ID or index |
| ListTransactions | [ListTransactionsRequest](#onos-config-admin-ListTransactionsRequest) | [ListTransactionsResponse](#onos-config-admin-ListTransactionsResponse) stream | List returns all configuration transactions |
| WatchTransactions | [WatchTransactionsRequest](#onos-config-admin-WatchTransactionsRequest) | [WatchTransactionsResponse](#onos-config-admin-WatchTransactionsResponse) stream | Watch returns a stream of configuration transaction change notifications |

 



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

