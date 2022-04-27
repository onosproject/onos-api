# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/ransim/metrics/metrics.proto](#onos_ransim_metrics_metrics-proto)
    - [DeleteAllRequest](#onos-ransim-metrics-DeleteAllRequest)
    - [DeleteAllResponse](#onos-ransim-metrics-DeleteAllResponse)
    - [DeleteRequest](#onos-ransim-metrics-DeleteRequest)
    - [DeleteResponse](#onos-ransim-metrics-DeleteResponse)
    - [GetRequest](#onos-ransim-metrics-GetRequest)
    - [GetResponse](#onos-ransim-metrics-GetResponse)
    - [ListRequest](#onos-ransim-metrics-ListRequest)
    - [ListResponse](#onos-ransim-metrics-ListResponse)
    - [Metric](#onos-ransim-metrics-Metric)
    - [SetRequest](#onos-ransim-metrics-SetRequest)
    - [SetResponse](#onos-ransim-metrics-SetResponse)
    - [WatchRequest](#onos-ransim-metrics-WatchRequest)
    - [WatchResponse](#onos-ransim-metrics-WatchResponse)
  
    - [EventType](#onos-ransim-metrics-EventType)
  
    - [MetricsService](#onos-ransim-metrics-MetricsService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos_ransim_metrics_metrics-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/ransim/metrics/metrics.proto



<a name="onos-ransim-metrics-DeleteAllRequest"></a>

### DeleteAllRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entityid | [uint64](#uint64) |  |  |






<a name="onos-ransim-metrics-DeleteAllResponse"></a>

### DeleteAllResponse







<a name="onos-ransim-metrics-DeleteRequest"></a>

### DeleteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entityid | [uint64](#uint64) |  |  |
| name | [string](#string) |  |  |






<a name="onos-ransim-metrics-DeleteResponse"></a>

### DeleteResponse







<a name="onos-ransim-metrics-GetRequest"></a>

### GetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entityid | [uint64](#uint64) |  |  |
| name | [string](#string) |  |  |






<a name="onos-ransim-metrics-GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metric | [Metric](#onos-ransim-metrics-Metric) |  |  |






<a name="onos-ransim-metrics-ListRequest"></a>

### ListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entityid | [uint64](#uint64) |  |  |






<a name="onos-ransim-metrics-ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entityid | [uint64](#uint64) |  |  |
| metrics | [Metric](#onos-ransim-metrics-Metric) | repeated |  |






<a name="onos-ransim-metrics-Metric"></a>

### Metric



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entityid | [uint64](#uint64) |  |  |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |
| type | [string](#string) |  | intX, uintX, floatX, string; X := {8|16|32|64} |






<a name="onos-ransim-metrics-SetRequest"></a>

### SetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metric | [Metric](#onos-ransim-metrics-Metric) |  |  |






<a name="onos-ransim-metrics-SetResponse"></a>

### SetResponse







<a name="onos-ransim-metrics-WatchRequest"></a>

### WatchRequest







<a name="onos-ransim-metrics-WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metric | [Metric](#onos-ransim-metrics-Metric) |  |  |
| type | [EventType](#onos-ransim-metrics-EventType) |  |  |





 


<a name="onos-ransim-metrics-EventType"></a>

### EventType
Change event type

| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 | NONE indicates unknown event type |
| UPDATED | 1 | UPDATED indicates a metric value was set (updated) |
| DELETED | 2 | DELETED indicates a metric was deleted |


 

 


<a name="onos-ransim-metrics-MetricsService"></a>

### MetricsService
Model provides means to create, delete and read RAN simulation model.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| List | [ListRequest](#onos-ransim-metrics-ListRequest) | [ListResponse](#onos-ransim-metrics-ListResponse) | List returns an array of all metrics for the specified entity (Node, Cell or UE) |
| Set | [SetRequest](#onos-ransim-metrics-SetRequest) | [SetResponse](#onos-ransim-metrics-SetResponse) | Set sets value of the named metric for the specified entity |
| Get | [GetRequest](#onos-ransim-metrics-GetRequest) | [GetResponse](#onos-ransim-metrics-GetResponse) | Get retrieves the named metric for the specified entity |
| Delete | [DeleteRequest](#onos-ransim-metrics-DeleteRequest) | [DeleteResponse](#onos-ransim-metrics-DeleteResponse) | Delete deletes the the named metric for the specified entity |
| DeleteAll | [DeleteAllRequest](#onos-ransim-metrics-DeleteAllRequest) | [DeleteAllResponse](#onos-ransim-metrics-DeleteAllResponse) | DeleteAll deletes all metrics for the specified entity |
| Watch | [WatchRequest](#onos-ransim-metrics-WatchRequest) | [WatchResponse](#onos-ransim-metrics-WatchResponse) stream | Watch returns a stream of ongoing changes to the metrics |

 



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

