# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/topo/topo.proto](#onos/topo/topo.proto)
    - [CreateRequest](#onos.topo.CreateRequest)
    - [CreateResponse](#onos.topo.CreateResponse)
    - [DeleteRequest](#onos.topo.DeleteRequest)
    - [DeleteResponse](#onos.topo.DeleteResponse)
    - [Entity](#onos.topo.Entity)
    - [EqualFilter](#onos.topo.EqualFilter)
    - [Event](#onos.topo.Event)
    - [Filter](#onos.topo.Filter)
    - [Filters](#onos.topo.Filters)
    - [GetRequest](#onos.topo.GetRequest)
    - [GetResponse](#onos.topo.GetResponse)
    - [InFilter](#onos.topo.InFilter)
    - [Kind](#onos.topo.Kind)
    - [ListRequest](#onos.topo.ListRequest)
    - [ListResponse](#onos.topo.ListResponse)
    - [NotFilter](#onos.topo.NotFilter)
    - [Object](#onos.topo.Object)
    - [Object.AspectsEntry](#onos.topo.Object.AspectsEntry)
    - [Object.LabelsEntry](#onos.topo.Object.LabelsEntry)
    - [Relation](#onos.topo.Relation)
    - [RelationFilter](#onos.topo.RelationFilter)
    - [UpdateRequest](#onos.topo.UpdateRequest)
    - [UpdateResponse](#onos.topo.UpdateResponse)
    - [WatchRequest](#onos.topo.WatchRequest)
    - [WatchResponse](#onos.topo.WatchResponse)
  
    - [EventType](#onos.topo.EventType)
    - [Object.Type](#onos.topo.Object.Type)
    - [RelationFilterScope](#onos.topo.RelationFilterScope)
    - [SortOrder](#onos.topo.SortOrder)
  
    - [Topo](#onos.topo.Topo)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos/topo/topo.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/topo/topo.proto



<a name="onos.topo.CreateRequest"></a>

### CreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object | [Object](#onos.topo.Object) |  |  |






<a name="onos.topo.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object | [Object](#onos.topo.Object) |  |  |






<a name="onos.topo.DeleteRequest"></a>

### DeleteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| revision | [uint64](#uint64) |  |  |






<a name="onos.topo.DeleteResponse"></a>

### DeleteResponse







<a name="onos.topo.Entity"></a>

### Entity
Entity represents any &#34;thing&#34; that is represented in the topology


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kind_id | [string](#string) |  | user-defined entity kind |
| src_relation_ids | [string](#string) | repeated | these lists are maintained by the system and are provided as read-only values for clients |
| tgt_relation_ids | [string](#string) | repeated |  |






<a name="onos.topo.EqualFilter"></a>

### EqualFilter



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) |  |  |






<a name="onos.topo.Event"></a>

### Event
Event is a topo operation event


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [EventType](#onos.topo.EventType) |  |  |
| object | [Object](#onos.topo.Object) |  |  |






<a name="onos.topo.Filter"></a>

### Filter



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| equal | [EqualFilter](#onos.topo.EqualFilter) |  |  |
| not | [NotFilter](#onos.topo.NotFilter) |  |  |
| in | [InFilter](#onos.topo.InFilter) |  |  |
| key | [string](#string) |  | optional key |






<a name="onos.topo.Filters"></a>

### Filters



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kind_filter | [Filter](#onos.topo.Filter) |  |  |
| label_filters | [Filter](#onos.topo.Filter) | repeated |  |
| relation_filter | [RelationFilter](#onos.topo.RelationFilter) |  |  |
| object_types | [Object.Type](#onos.topo.Object.Type) | repeated |  |






<a name="onos.topo.GetRequest"></a>

### GetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="onos.topo.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object | [Object](#onos.topo.Object) |  |  |






<a name="onos.topo.InFilter"></a>

### InFilter



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| values | [string](#string) | repeated |  |






<a name="onos.topo.Kind"></a>

### Kind
Kind represents an archetype of an object, i.e. entity or relation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Friendly name of the kind |






<a name="onos.topo.ListRequest"></a>

### ListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filters | [Filters](#onos.topo.Filters) |  |  |
| sort_order | [SortOrder](#onos.topo.SortOrder) |  |  |






<a name="onos.topo.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| objects | [Object](#onos.topo.Object) | repeated |  |






<a name="onos.topo.NotFilter"></a>

### NotFilter



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| inner | [Filter](#onos.topo.Filter) |  |  |






<a name="onos.topo.Object"></a>

### Object
Object is an one of the following: a kind (archetype of entity or relation), an entity, a relation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |
| id | [string](#string) |  |  |
| revision | [uint64](#uint64) |  |  |
| type | [Object.Type](#onos.topo.Object.Type) |  |  |
| entity | [Entity](#onos.topo.Entity) |  |  |
| relation | [Relation](#onos.topo.Relation) |  |  |
| kind | [Kind](#onos.topo.Kind) |  |  |
| aspects | [Object.AspectsEntry](#onos.topo.Object.AspectsEntry) | repeated | Map of aspects as typed values; for kind, these represent expected aspects and their default values |
| labels | [Object.LabelsEntry](#onos.topo.Object.LabelsEntry) | repeated | Arbitrary labels for classification/search |






<a name="onos.topo.Object.AspectsEntry"></a>

### Object.AspectsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="onos.topo.Object.LabelsEntry"></a>

### Object.LabelsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="onos.topo.Relation"></a>

### Relation
Relation represents any &#34;relation&#34; between two entitites in the topology.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kind_id | [string](#string) |  | user defined relation kind |
| src_entity_id | [string](#string) |  |  |
| tgt_entity_id | [string](#string) |  |  |






<a name="onos.topo.RelationFilter"></a>

### RelationFilter
Filter for targets of given relation kinds and given source ids; optionally, filters by specified target kind


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| src_id | [string](#string) |  |  |
| relation_kind | [string](#string) |  |  |
| target_kind | [string](#string) |  |  |
| scope | [RelationFilterScope](#onos.topo.RelationFilterScope) |  |  |






<a name="onos.topo.UpdateRequest"></a>

### UpdateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object | [Object](#onos.topo.Object) |  |  |






<a name="onos.topo.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object | [Object](#onos.topo.Object) |  |  |






<a name="onos.topo.WatchRequest"></a>

### WatchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filters | [Filters](#onos.topo.Filters) |  |  |
| noreplay | [bool](#bool) |  |  |






<a name="onos.topo.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#onos.topo.Event) |  |  |





 


<a name="onos.topo.EventType"></a>

### EventType
EventType is a topo operation event type

| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 |  |
| ADDED | 1 |  |
| UPDATED | 2 |  |
| REMOVED | 3 |  |



<a name="onos.topo.Object.Type"></a>

### Object.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNSPECIFIED | 0 |  |
| ENTITY | 1 |  |
| RELATION | 2 |  |
| KIND | 3 |  |



<a name="onos.topo.RelationFilterScope"></a>

### RelationFilterScope


| Name | Number | Description |
| ---- | ------ | ----------- |
| TARGET_ONLY | 0 |  |
| ALL | 1 |  |
| SOURCE_AND_TARGET | 2 |  |



<a name="onos.topo.SortOrder"></a>

### SortOrder


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNORDERED | 0 |  |
| ASCENDING | 1 |  |
| DESCENDING | 2 |  |


 

 


<a name="onos.topo.Topo"></a>

### Topo
EntityService provides an API for managing entities.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateRequest](#onos.topo.CreateRequest) | [CreateResponse](#onos.topo.CreateResponse) | Create a new topology object |
| Get | [GetRequest](#onos.topo.GetRequest) | [GetResponse](#onos.topo.GetResponse) | Get an object from topology |
| Update | [UpdateRequest](#onos.topo.UpdateRequest) | [UpdateResponse](#onos.topo.UpdateResponse) | Update an existing topology object |
| Delete | [DeleteRequest](#onos.topo.DeleteRequest) | [DeleteResponse](#onos.topo.DeleteResponse) | Delete an object from topology |
| List | [ListRequest](#onos.topo.ListRequest) | [ListResponse](#onos.topo.ListResponse) | List gets a stream of requested objects |
| Watch | [WatchRequest](#onos.topo.WatchRequest) | [WatchResponse](#onos.topo.WatchResponse) stream | Watch returns a stream of topo change notifications |

 



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

