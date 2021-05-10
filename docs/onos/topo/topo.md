# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/topo/topo.proto](#onos/topo/topo.proto)
    - [CreateRequest](#onos.topo.CreateRequest)
    - [CreateResponse](#onos.topo.CreateResponse)
    - [DeleteRequest](#onos.topo.DeleteRequest)
    - [DeleteResponse](#onos.topo.DeleteResponse)
    - [Entity](#onos.topo.Entity)
    - [Event](#onos.topo.Event)
    - [GetRequest](#onos.topo.GetRequest)
    - [GetResponse](#onos.topo.GetResponse)
    - [Kind](#onos.topo.Kind)
    - [ListRequest](#onos.topo.ListRequest)
    - [ListResponse](#onos.topo.ListResponse)
    - [Object](#onos.topo.Object)
    - [Object.AttributesEntry](#onos.topo.Object.AttributesEntry)
    - [ProtocolState](#onos.topo.ProtocolState)
    - [Relation](#onos.topo.Relation)
    - [UpdateRequest](#onos.topo.UpdateRequest)
    - [UpdateResponse](#onos.topo.UpdateResponse)
    - [Value](#onos.topo.Value)
    - [WatchRequest](#onos.topo.WatchRequest)
    - [WatchResponse](#onos.topo.WatchResponse)
  
    - [ChannelState](#onos.topo.ChannelState)
    - [ConnectivityState](#onos.topo.ConnectivityState)
    - [EventType](#onos.topo.EventType)
    - [Object.Type](#onos.topo.Object.Type)
    - [Protocol](#onos.topo.Protocol)
    - [ServiceState](#onos.topo.ServiceState)
    - [Value.Type](#onos.topo.Value.Type)
  
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






<a name="onos.topo.DeleteResponse"></a>

### DeleteResponse







<a name="onos.topo.Entity"></a>

### Entity
Entity represents any &#34;thing&#34; that is represented in the topology


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kind_id | [string](#string) |  | user-defined entity kind |
| protocols | [ProtocolState](#onos.topo.ProtocolState) | repeated |  |






<a name="onos.topo.Event"></a>

### Event
Event is a topo operation event


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [EventType](#onos.topo.EventType) |  |  |
| object | [Object](#onos.topo.Object) |  |  |






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






<a name="onos.topo.Kind"></a>

### Kind
Kind represents an archetype of an object, i.e. entity or relation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | name, a.k.a kind_id |






<a name="onos.topo.ListRequest"></a>

### ListRequest







<a name="onos.topo.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| objects | [Object](#onos.topo.Object) | repeated |  |






<a name="onos.topo.Object"></a>

### Object
Object is an one of the following: a kind (archetype of entity or relation), an entity, a relation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| revision | [uint64](#uint64) |  |  |
| type | [Object.Type](#onos.topo.Object.Type) |  |  |
| entity | [Entity](#onos.topo.Entity) |  |  |
| relation | [Relation](#onos.topo.Relation) |  |  |
| kind | [Kind](#onos.topo.Kind) |  |  |
| attributes | [Object.AttributesEntry](#onos.topo.Object.AttributesEntry) | repeated | Map of attributes as typed values; for kind, these represent expected attributed and their default values |






<a name="onos.topo.Object.AttributesEntry"></a>

### Object.AttributesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [Value](#onos.topo.Value) |  |  |






<a name="onos.topo.ProtocolState"></a>

### ProtocolState
ProtocolState contains information related to service and connectivity to a device


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| protocol | [Protocol](#onos.topo.Protocol) |  | The protocol to which state relates |
| connectivityState | [ConnectivityState](#onos.topo.ConnectivityState) |  | ConnectivityState contains the L3 connectivity information |
| channelState | [ChannelState](#onos.topo.ChannelState) |  | ChannelState relates to the availability of the gRPC channel |
| serviceState | [ServiceState](#onos.topo.ServiceState) |  | ServiceState indicates the availability of the gRPC servic on top of the channel |






<a name="onos.topo.Relation"></a>

### Relation
Relation represents any &#34;relation&#34; between two entitites in the topology.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kind_id | [string](#string) |  | user defined relation kind |
| src_entity_id | [string](#string) |  |  |
| tgt_entity_id | [string](#string) |  |  |






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






<a name="onos.topo.Value"></a>

### Value
Value is a type/value pair with the value itself being a selection dependent on the type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [Value.Type](#onos.topo.Value.Type) |  |  |
| stringValue | [string](#string) |  |  |
| uintValue | [uint64](#uint64) |  |  |
| intValue | [int64](#int64) |  |  |
| boolValue | [bool](#bool) |  |  |
| protoValue | [bytes](#bytes) |  |  |
| bytesValue | [bytes](#bytes) |  |  |






<a name="onos.topo.WatchRequest"></a>

### WatchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| noreplay | [bool](#bool) |  |  |






<a name="onos.topo.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#onos.topo.Event) |  |  |





 


<a name="onos.topo.ChannelState"></a>

### ChannelState
ConnectivityState represents the state of a gRPC channel to the device from the service container

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN_CHANNEL_STATE | 0 | UNKNOWN_CHANNEL_STATE constant needed to go around proto3 nullifying the 0 values |
| CONNECTED | 1 | CONNECTED indicates the corresponding grpc channel is connected on this device |
| DISCONNECTED | 2 | DISCONNECTED indicates the corresponding grpc channel is not connected on this device |



<a name="onos.topo.ConnectivityState"></a>

### ConnectivityState
ConnectivityState represents the L3 reachability of a device from the service container (e.g. enos-config), independently of gRPC or the service itself (e.g. gNMI)

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN_CONNECTIVITY_STATE | 0 | UNKNOWN_CONNECTIVITY_STATE constant needed to go around proto3 nullifying the 0 values |
| REACHABLE | 1 | REACHABLE indicates the the service can reach the device at L3 |
| UNREACHABLE | 2 | UNREACHABLE indicates the the service can&#39;t reach the device at L3 |



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



<a name="onos.topo.Protocol"></a>

### Protocol
Protocol to interact with a device

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN_PROTOCOL | 0 | UNKNOWN_PROTOCOL constant needed to go around proto3 nullifying the 0 values |
| GNMI | 1 | GNMI protocol reference |
| P4RUNTIME | 2 | P4RUNTIME protocol reference |
| GNOI | 3 | GNOI protocol reference |
| E2AP | 4 | E2 Control Plane Protocol |



<a name="onos.topo.ServiceState"></a>

### ServiceState
ServiceState represents the state of the gRPC service (e.g. gNMI) to the device from the service container

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN_SERVICE_STATE | 0 | UNKNOWN_SERVICE_STATE constant needed to go around proto3 nullifying the 0 values |
| AVAILABLE | 1 | AVAILABLE indicates the corresponding grpc service is available |
| UNAVAILABLE | 2 | UNAVAILABLE indicates the corresponding grpc service is not available |
| CONNECTING | 3 | CONNECTING indicates the corresponding protocol is in the connecting phase on this device |



<a name="onos.topo.Value.Type"></a>

### Value.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| STRING | 0 |  |
| UINT | 1 |  |
| INT | 2 |  |
| BOOLEAN | 3 |  |
| PROTO | 4 |  |
| BYTES | 5 |  |


 

 


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

