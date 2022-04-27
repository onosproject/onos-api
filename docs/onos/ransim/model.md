# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/ransim/model/model.proto](#onos_ransim_model_model-proto)
    - [AgentControlRequest](#onos-ransim-model-AgentControlRequest)
    - [AgentControlResponse](#onos-ransim-model-AgentControlResponse)
    - [ClearRequest](#onos-ransim-model-ClearRequest)
    - [ClearResponse](#onos-ransim-model-ClearResponse)
    - [CreateCellRequest](#onos-ransim-model-CreateCellRequest)
    - [CreateCellResponse](#onos-ransim-model-CreateCellResponse)
    - [CreateNodeRequest](#onos-ransim-model-CreateNodeRequest)
    - [CreateNodeResponse](#onos-ransim-model-CreateNodeResponse)
    - [CreateRouteRequest](#onos-ransim-model-CreateRouteRequest)
    - [CreateRouteResponse](#onos-ransim-model-CreateRouteResponse)
    - [DataSet](#onos-ransim-model-DataSet)
    - [DeleteCellRequest](#onos-ransim-model-DeleteCellRequest)
    - [DeleteCellResponse](#onos-ransim-model-DeleteCellResponse)
    - [DeleteNodeRequest](#onos-ransim-model-DeleteNodeRequest)
    - [DeleteNodeResponse](#onos-ransim-model-DeleteNodeResponse)
    - [DeleteRouteRequest](#onos-ransim-model-DeleteRouteRequest)
    - [DeleteRouteResponse](#onos-ransim-model-DeleteRouteResponse)
    - [DeleteUERequest](#onos-ransim-model-DeleteUERequest)
    - [DeleteUEResponse](#onos-ransim-model-DeleteUEResponse)
    - [GetCellRequest](#onos-ransim-model-GetCellRequest)
    - [GetCellResponse](#onos-ransim-model-GetCellResponse)
    - [GetNodeRequest](#onos-ransim-model-GetNodeRequest)
    - [GetNodeResponse](#onos-ransim-model-GetNodeResponse)
    - [GetRouteRequest](#onos-ransim-model-GetRouteRequest)
    - [GetRouteResponse](#onos-ransim-model-GetRouteResponse)
    - [GetUECountRequest](#onos-ransim-model-GetUECountRequest)
    - [GetUECountResponse](#onos-ransim-model-GetUECountResponse)
    - [GetUERequest](#onos-ransim-model-GetUERequest)
    - [GetUEResponse](#onos-ransim-model-GetUEResponse)
    - [ListCellsRequest](#onos-ransim-model-ListCellsRequest)
    - [ListCellsResponse](#onos-ransim-model-ListCellsResponse)
    - [ListNodesRequest](#onos-ransim-model-ListNodesRequest)
    - [ListNodesResponse](#onos-ransim-model-ListNodesResponse)
    - [ListRoutesRequest](#onos-ransim-model-ListRoutesRequest)
    - [ListRoutesResponse](#onos-ransim-model-ListRoutesResponse)
    - [ListUEsRequest](#onos-ransim-model-ListUEsRequest)
    - [ListUEsResponse](#onos-ransim-model-ListUEsResponse)
    - [LoadRequest](#onos-ransim-model-LoadRequest)
    - [LoadResponse](#onos-ransim-model-LoadResponse)
    - [MoveToCellRequest](#onos-ransim-model-MoveToCellRequest)
    - [MoveToCellResponse](#onos-ransim-model-MoveToCellResponse)
    - [MoveToLocationRequest](#onos-ransim-model-MoveToLocationRequest)
    - [MoveToLocationResponse](#onos-ransim-model-MoveToLocationResponse)
    - [PlmnIDRequest](#onos-ransim-model-PlmnIDRequest)
    - [PlmnIDResponse](#onos-ransim-model-PlmnIDResponse)
    - [SetUECountRequest](#onos-ransim-model-SetUECountRequest)
    - [SetUECountResponse](#onos-ransim-model-SetUECountResponse)
    - [UpdateCellRequest](#onos-ransim-model-UpdateCellRequest)
    - [UpdateCellResponse](#onos-ransim-model-UpdateCellResponse)
    - [UpdateNodeRequest](#onos-ransim-model-UpdateNodeRequest)
    - [UpdateNodeResponse](#onos-ransim-model-UpdateNodeResponse)
    - [WatchCellsRequest](#onos-ransim-model-WatchCellsRequest)
    - [WatchCellsResponse](#onos-ransim-model-WatchCellsResponse)
    - [WatchNodesRequest](#onos-ransim-model-WatchNodesRequest)
    - [WatchNodesResponse](#onos-ransim-model-WatchNodesResponse)
    - [WatchRoutesRequest](#onos-ransim-model-WatchRoutesRequest)
    - [WatchRoutesResponse](#onos-ransim-model-WatchRoutesResponse)
    - [WatchUEsRequest](#onos-ransim-model-WatchUEsRequest)
    - [WatchUEsResponse](#onos-ransim-model-WatchUEsResponse)
  
    - [EventType](#onos-ransim-model-EventType)
  
    - [CellModel](#onos-ransim-model-CellModel)
    - [ModelService](#onos-ransim-model-ModelService)
    - [NodeModel](#onos-ransim-model-NodeModel)
    - [RouteModel](#onos-ransim-model-RouteModel)
    - [UEModel](#onos-ransim-model-UEModel)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos_ransim_model_model-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/ransim/model/model.proto



<a name="onos-ransim-model-AgentControlRequest"></a>

### AgentControlRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enbid | [uint64](#uint64) |  |  |
| command | [string](#string) |  | start, stop, drop, reconnect, etc. |
| args | [string](#string) | repeated | optional command parameters |






<a name="onos-ransim-model-AgentControlResponse"></a>

### AgentControlResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node | [onos.ransim.types.Node](#onos-ransim-types-Node) |  |  |






<a name="onos-ransim-model-ClearRequest"></a>

### ClearRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resume | [bool](#bool) |  |  |






<a name="onos-ransim-model-ClearResponse"></a>

### ClearResponse







<a name="onos-ransim-model-CreateCellRequest"></a>

### CreateCellRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cell | [onos.ransim.types.Cell](#onos-ransim-types-Cell) |  |  |






<a name="onos-ransim-model-CreateCellResponse"></a>

### CreateCellResponse







<a name="onos-ransim-model-CreateNodeRequest"></a>

### CreateNodeRequest
CreateNodeRequest create a node request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node | [onos.ransim.types.Node](#onos-ransim-types-Node) |  |  |






<a name="onos-ransim-model-CreateNodeResponse"></a>

### CreateNodeResponse
CreateNodeResponse create a node response






<a name="onos-ransim-model-CreateRouteRequest"></a>

### CreateRouteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route | [onos.ransim.types.Route](#onos-ransim-types-Route) |  |  |






<a name="onos-ransim-model-CreateRouteResponse"></a>

### CreateRouteResponse







<a name="onos-ransim-model-DataSet"></a>

### DataSet



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| data | [bytes](#bytes) |  |  |






<a name="onos-ransim-model-DeleteCellRequest"></a>

### DeleteCellRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enbid | [uint64](#uint64) |  |  |






<a name="onos-ransim-model-DeleteCellResponse"></a>

### DeleteCellResponse







<a name="onos-ransim-model-DeleteNodeRequest"></a>

### DeleteNodeRequest
DeleteNodeRequest delete a node request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enbid | [uint64](#uint64) |  |  |






<a name="onos-ransim-model-DeleteNodeResponse"></a>

### DeleteNodeResponse
DeleteNodeResponse delete a node response






<a name="onos-ransim-model-DeleteRouteRequest"></a>

### DeleteRouteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enbid | [uint64](#uint64) |  |  |






<a name="onos-ransim-model-DeleteRouteResponse"></a>

### DeleteRouteResponse







<a name="onos-ransim-model-DeleteUERequest"></a>

### DeleteUERequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| imsi | [uint32](#uint32) |  |  |






<a name="onos-ransim-model-DeleteUEResponse"></a>

### DeleteUEResponse







<a name="onos-ransim-model-GetCellRequest"></a>

### GetCellRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ecgi | [uint64](#uint64) |  |  |






<a name="onos-ransim-model-GetCellResponse"></a>

### GetCellResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cell | [onos.ransim.types.Cell](#onos-ransim-types-Cell) |  |  |






<a name="onos-ransim-model-GetNodeRequest"></a>

### GetNodeRequest
GetNodeRequest get a node request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enbid | [uint64](#uint64) |  |  |






<a name="onos-ransim-model-GetNodeResponse"></a>

### GetNodeResponse
GetNodeResponse get a node response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node | [onos.ransim.types.Node](#onos-ransim-types-Node) |  |  |






<a name="onos-ransim-model-GetRouteRequest"></a>

### GetRouteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| imsi | [uint32](#uint32) |  |  |






<a name="onos-ransim-model-GetRouteResponse"></a>

### GetRouteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route | [onos.ransim.types.Route](#onos-ransim-types-Route) |  |  |






<a name="onos-ransim-model-GetUECountRequest"></a>

### GetUECountRequest







<a name="onos-ransim-model-GetUECountResponse"></a>

### GetUECountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [uint32](#uint32) |  |  |






<a name="onos-ransim-model-GetUERequest"></a>

### GetUERequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| imsi | [uint32](#uint32) |  |  |






<a name="onos-ransim-model-GetUEResponse"></a>

### GetUEResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ue | [onos.ransim.types.Ue](#onos-ransim-types-Ue) |  |  |






<a name="onos-ransim-model-ListCellsRequest"></a>

### ListCellsRequest







<a name="onos-ransim-model-ListCellsResponse"></a>

### ListCellsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Cell | [onos.ransim.types.Cell](#onos-ransim-types-Cell) |  |  |






<a name="onos-ransim-model-ListNodesRequest"></a>

### ListNodesRequest







<a name="onos-ransim-model-ListNodesResponse"></a>

### ListNodesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node | [onos.ransim.types.Node](#onos-ransim-types-Node) |  |  |






<a name="onos-ransim-model-ListRoutesRequest"></a>

### ListRoutesRequest







<a name="onos-ransim-model-ListRoutesResponse"></a>

### ListRoutesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route | [onos.ransim.types.Route](#onos-ransim-types-Route) |  |  |






<a name="onos-ransim-model-ListUEsRequest"></a>

### ListUEsRequest







<a name="onos-ransim-model-ListUEsResponse"></a>

### ListUEsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ue | [onos.ransim.types.Ue](#onos-ransim-types-Ue) |  |  |






<a name="onos-ransim-model-LoadRequest"></a>

### LoadRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataSet | [DataSet](#onos-ransim-model-DataSet) | repeated |  |
| resume | [bool](#bool) |  |  |






<a name="onos-ransim-model-LoadResponse"></a>

### LoadResponse







<a name="onos-ransim-model-MoveToCellRequest"></a>

### MoveToCellRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| imsi | [uint32](#uint32) |  |  |
| ecgi | [uint32](#uint32) |  |  |






<a name="onos-ransim-model-MoveToCellResponse"></a>

### MoveToCellResponse







<a name="onos-ransim-model-MoveToLocationRequest"></a>

### MoveToLocationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| imsi | [uint32](#uint32) |  |  |
| location | [onos.ransim.types.Point](#onos-ransim-types-Point) |  |  |
| heading | [uint32](#uint32) |  |  |






<a name="onos-ransim-model-MoveToLocationResponse"></a>

### MoveToLocationResponse







<a name="onos-ransim-model-PlmnIDRequest"></a>

### PlmnIDRequest







<a name="onos-ransim-model-PlmnIDResponse"></a>

### PlmnIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| plmnid | [uint32](#uint32) |  |  |






<a name="onos-ransim-model-SetUECountRequest"></a>

### SetUECountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [uint32](#uint32) |  |  |






<a name="onos-ransim-model-SetUECountResponse"></a>

### SetUECountResponse







<a name="onos-ransim-model-UpdateCellRequest"></a>

### UpdateCellRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cell | [onos.ransim.types.Cell](#onos-ransim-types-Cell) |  |  |






<a name="onos-ransim-model-UpdateCellResponse"></a>

### UpdateCellResponse







<a name="onos-ransim-model-UpdateNodeRequest"></a>

### UpdateNodeRequest
UpdateNodeRequest update a node request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node | [onos.ransim.types.Node](#onos-ransim-types-Node) |  |  |






<a name="onos-ransim-model-UpdateNodeResponse"></a>

### UpdateNodeResponse
UpdateNodeResponse update a node response






<a name="onos-ransim-model-WatchCellsRequest"></a>

### WatchCellsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| no_replay | [bool](#bool) |  |  |
| no_subscribe | [bool](#bool) |  |  |






<a name="onos-ransim-model-WatchCellsResponse"></a>

### WatchCellsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cell | [onos.ransim.types.Cell](#onos-ransim-types-Cell) |  |  |
| type | [EventType](#onos-ransim-model-EventType) |  |  |






<a name="onos-ransim-model-WatchNodesRequest"></a>

### WatchNodesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| no_replay | [bool](#bool) |  |  |
| no_subscribe | [bool](#bool) |  |  |






<a name="onos-ransim-model-WatchNodesResponse"></a>

### WatchNodesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node | [onos.ransim.types.Node](#onos-ransim-types-Node) |  |  |
| type | [EventType](#onos-ransim-model-EventType) |  |  |






<a name="onos-ransim-model-WatchRoutesRequest"></a>

### WatchRoutesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| no_replay | [bool](#bool) |  |  |
| no_subscribe | [bool](#bool) |  |  |






<a name="onos-ransim-model-WatchRoutesResponse"></a>

### WatchRoutesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route | [onos.ransim.types.Route](#onos-ransim-types-Route) |  |  |
| type | [EventType](#onos-ransim-model-EventType) |  |  |






<a name="onos-ransim-model-WatchUEsRequest"></a>

### WatchUEsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| no_replay | [bool](#bool) |  |  |
| no_subscribe | [bool](#bool) |  |  |






<a name="onos-ransim-model-WatchUEsResponse"></a>

### WatchUEsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ue | [onos.ransim.types.Ue](#onos-ransim-types-Ue) |  |  |
| type | [EventType](#onos-ransim-model-EventType) |  |  |





 


<a name="onos-ransim-model-EventType"></a>

### EventType
Change event type

| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 | NONE indicates this response represents a pre-existing entity |
| CREATED | 1 | CREATED indicates a new entity was created |
| UPDATED | 2 | UPDATED indicates an existing entity was updated |
| DELETED | 3 | DELETED indicates an entity was deleted |


 

 


<a name="onos-ransim-model-CellModel"></a>

### CellModel
CellModel provides means to create, delete and read simulated RAN cells.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateCell | [CreateCellRequest](#onos-ransim-model-CreateCellRequest) | [CreateCellResponse](#onos-ransim-model-CreateCellResponse) |  |
| DeleteCell | [DeleteCellRequest](#onos-ransim-model-DeleteCellRequest) | [DeleteCellResponse](#onos-ransim-model-DeleteCellResponse) |  |
| UpdateCell | [UpdateCellRequest](#onos-ransim-model-UpdateCellRequest) | [UpdateCellResponse](#onos-ransim-model-UpdateCellResponse) |  |
| GetCell | [GetCellRequest](#onos-ransim-model-GetCellRequest) | [GetCellResponse](#onos-ransim-model-GetCellResponse) |  |
| WatchCells | [WatchCellsRequest](#onos-ransim-model-WatchCellsRequest) | [WatchCellsResponse](#onos-ransim-model-WatchCellsResponse) stream |  |
| ListCells | [ListCellsRequest](#onos-ransim-model-ListCellsRequest) | [ListCellsResponse](#onos-ransim-model-ListCellsResponse) stream |  |


<a name="onos-ransim-model-ModelService"></a>

### ModelService
ModelService provides means to clear and load node and cell model in bulk

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Load | [LoadRequest](#onos-ransim-model-LoadRequest) | [LoadResponse](#onos-ransim-model-LoadResponse) |  |
| Clear | [ClearRequest](#onos-ransim-model-ClearRequest) | [ClearResponse](#onos-ransim-model-ClearResponse) |  |


<a name="onos-ransim-model-NodeModel"></a>

### NodeModel
NodeModel provides means to create, delete and read simulated RAN nodes.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetPlmnID | [PlmnIDRequest](#onos-ransim-model-PlmnIDRequest) | [PlmnIDResponse](#onos-ransim-model-PlmnIDResponse) |  |
| CreateNode | [CreateNodeRequest](#onos-ransim-model-CreateNodeRequest) | [CreateNodeResponse](#onos-ransim-model-CreateNodeResponse) |  |
| GetNode | [GetNodeRequest](#onos-ransim-model-GetNodeRequest) | [GetNodeResponse](#onos-ransim-model-GetNodeResponse) |  |
| UpdateNode | [UpdateNodeRequest](#onos-ransim-model-UpdateNodeRequest) | [UpdateNodeResponse](#onos-ransim-model-UpdateNodeResponse) |  |
| DeleteNode | [DeleteNodeRequest](#onos-ransim-model-DeleteNodeRequest) | [DeleteNodeResponse](#onos-ransim-model-DeleteNodeResponse) |  |
| WatchNodes | [WatchNodesRequest](#onos-ransim-model-WatchNodesRequest) | [WatchNodesResponse](#onos-ransim-model-WatchNodesResponse) stream |  |
| ListNodes | [ListNodesRequest](#onos-ransim-model-ListNodesRequest) | [ListNodesResponse](#onos-ransim-model-ListNodesResponse) stream |  |
| AgentControl | [AgentControlRequest](#onos-ransim-model-AgentControlRequest) | [AgentControlResponse](#onos-ransim-model-AgentControlResponse) |  |


<a name="onos-ransim-model-RouteModel"></a>

### RouteModel
RouteModel provides means to create, delete and read simulated mobile UE routes.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateRoute | [CreateRouteRequest](#onos-ransim-model-CreateRouteRequest) | [CreateRouteResponse](#onos-ransim-model-CreateRouteResponse) |  |
| DeleteRoute | [DeleteRouteRequest](#onos-ransim-model-DeleteRouteRequest) | [DeleteRouteResponse](#onos-ransim-model-DeleteRouteResponse) |  |
| GetRoute | [GetRouteRequest](#onos-ransim-model-GetRouteRequest) | [GetRouteResponse](#onos-ransim-model-GetRouteResponse) |  |
| WatchRoutes | [WatchRoutesRequest](#onos-ransim-model-WatchRoutesRequest) | [WatchRoutesResponse](#onos-ransim-model-WatchRoutesResponse) stream |  |
| ListRoutes | [ListRoutesRequest](#onos-ransim-model-ListRoutesRequest) | [ListRoutesResponse](#onos-ransim-model-ListRoutesResponse) stream |  |


<a name="onos-ransim-model-UEModel"></a>

### UEModel
UEModel provides means to simulate mobile UEs.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetUE | [GetUERequest](#onos-ransim-model-GetUERequest) | [GetUEResponse](#onos-ransim-model-GetUEResponse) |  |
| MoveToCell | [MoveToCellRequest](#onos-ransim-model-MoveToCellRequest) | [MoveToCellResponse](#onos-ransim-model-MoveToCellResponse) |  |
| MoveToLocation | [MoveToLocationRequest](#onos-ransim-model-MoveToLocationRequest) | [MoveToLocationResponse](#onos-ransim-model-MoveToLocationResponse) |  |
| DeleteUE | [DeleteUERequest](#onos-ransim-model-DeleteUERequest) | [DeleteUEResponse](#onos-ransim-model-DeleteUEResponse) |  |
| WatchUEs | [WatchUEsRequest](#onos-ransim-model-WatchUEsRequest) | [WatchUEsResponse](#onos-ransim-model-WatchUEsResponse) stream |  |
| ListUEs | [ListUEsRequest](#onos-ransim-model-ListUEsRequest) | [ListUEsResponse](#onos-ransim-model-ListUEsResponse) stream |  |
| GetUECount | [GetUECountRequest](#onos-ransim-model-GetUECountRequest) | [GetUECountResponse](#onos-ransim-model-GetUECountResponse) |  |
| SetUECount | [SetUECountRequest](#onos-ransim-model-SetUECountRequest) | [SetUECountResponse](#onos-ransim-model-SetUECountResponse) |  |

 



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

