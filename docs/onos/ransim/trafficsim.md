# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/ransim/trafficsim/trafficsim.proto](#onos/ransim/trafficsim/trafficsim.proto)
    - [ListRoutesRequest](#onos.ransim.trafficsim.ListRoutesRequest)
    - [ListRoutesResponse](#onos.ransim.trafficsim.ListRoutesResponse)
    - [ListUesRequest](#onos.ransim.trafficsim.ListUesRequest)
    - [ListUesResponse](#onos.ransim.trafficsim.ListUesResponse)
    - [MapLayoutRequest](#onos.ransim.trafficsim.MapLayoutRequest)
    - [ResetMetricsMsg](#onos.ransim.trafficsim.ResetMetricsMsg)
    - [SetNumberUEsRequest](#onos.ransim.trafficsim.SetNumberUEsRequest)
    - [SetNumberUEsResponse](#onos.ransim.trafficsim.SetNumberUEsResponse)
    - [WatchUesRequest](#onos.ransim.trafficsim.WatchUesRequest)
    - [WatchUesResponse](#onos.ransim.trafficsim.WatchUesResponse)
  
    - [Type](#onos.ransim.trafficsim.Type)
    - [UpdateType](#onos.ransim.trafficsim.UpdateType)
  
    - [Traffic](#onos.ransim.trafficsim.Traffic)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos/ransim/trafficsim/trafficsim.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/ransim/trafficsim/trafficsim.proto



<a name="onos.ransim.trafficsim.ListRoutesRequest"></a>

### ListRoutesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| no_replay | [bool](#bool) |  |  |
| no_subscribe | [bool](#bool) |  |  |






<a name="onos.ransim.trafficsim.ListRoutesResponse"></a>

### ListRoutesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route | [onos.ransim.types.Route](#onos.ransim.types.Route) |  | route is the route change on which the event occurred |
| type | [Type](#onos.ransim.trafficsim.Type) |  | type is a qualification of the type of change being made |






<a name="onos.ransim.trafficsim.ListUesRequest"></a>

### ListUesRequest







<a name="onos.ransim.trafficsim.ListUesResponse"></a>

### ListUesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ue | [onos.ransim.types.Ue](#onos.ransim.types.Ue) |  | Ue is the UserEquipment change on which the event occurred |






<a name="onos.ransim.trafficsim.MapLayoutRequest"></a>

### MapLayoutRequest







<a name="onos.ransim.trafficsim.ResetMetricsMsg"></a>

### ResetMetricsMsg







<a name="onos.ransim.trafficsim.SetNumberUEsRequest"></a>

### SetNumberUEsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| number | [uint32](#uint32) |  |  |






<a name="onos.ransim.trafficsim.SetNumberUEsResponse"></a>

### SetNumberUEsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| number | [uint32](#uint32) |  |  |






<a name="onos.ransim.trafficsim.WatchUesRequest"></a>

### WatchUesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| no_replay | [bool](#bool) |  |  |
| no_subscribe | [bool](#bool) |  |  |






<a name="onos.ransim.trafficsim.WatchUesResponse"></a>

### WatchUesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ue | [onos.ransim.types.Ue](#onos.ransim.types.Ue) |  | Ue is the UserEquipment change on which the event occurred |
| type | [Type](#onos.ransim.trafficsim.Type) |  | type is a qualification of the type of change being made |
| update_type | [UpdateType](#onos.ransim.trafficsim.UpdateType) |  | update_type is a qualification of the type of UE change |





 


<a name="onos.ransim.trafficsim.Type"></a>

### Type
Change event type

| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 | NONE indicates this response does not represent a modification of the Change |
| ADDED | 1 | ADDED is an event which occurs when a Change is added to the topology |
| UPDATED | 2 | UPDATED is an event which occurs when a Change is updated |
| REMOVED | 3 | REMOVED is an event which occurs when a Change is removed from the configuration |



<a name="onos.ransim.trafficsim.UpdateType"></a>

### UpdateType
In the case of an update this helps qualify the update type - add items as necessary

| Name | Number | Description |
| ---- | ------ | ----------- |
| NOUPDATETYPE | 0 |  |
| POSITION | 1 | POSITION An update in position of UE only, without a change in the Tower affinity |
| TOWER | 2 | TOWER An update in non-serving tower affinity (and also position) of UE |
| HANDOVER | 3 | HANDOVER An update of the serving tower |


 

 


<a name="onos.ransim.trafficsim.Traffic"></a>

### Traffic
Traffic - provides a stream of traffic data to GUI

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetMapLayout | [MapLayoutRequest](#onos.ransim.trafficsim.MapLayoutRequest) | [.onos.ransim.types.MapLayout](#onos.ransim.types.MapLayout) |  |
| ListRoutes | [ListRoutesRequest](#onos.ransim.trafficsim.ListRoutesRequest) | [ListRoutesResponse](#onos.ransim.trafficsim.ListRoutesResponse) stream |  |
| ListUes | [ListUesRequest](#onos.ransim.trafficsim.ListUesRequest) | [ListUesResponse](#onos.ransim.trafficsim.ListUesResponse) stream |  |
| WatchUes | [WatchUesRequest](#onos.ransim.trafficsim.WatchUesRequest) | [WatchUesResponse](#onos.ransim.trafficsim.WatchUesResponse) stream |  |
| SetNumberUEs | [SetNumberUEsRequest](#onos.ransim.trafficsim.SetNumberUEsRequest) | [SetNumberUEsResponse](#onos.ransim.trafficsim.SetNumberUEsResponse) |  |
| ResetMetrics | [ResetMetricsMsg](#onos.ransim.trafficsim.ResetMetricsMsg) | [ResetMetricsMsg](#onos.ransim.trafficsim.ResetMetricsMsg) |  |

 



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

