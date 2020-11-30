# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/e2sub/endpoint/endpoint.proto](#onos/e2sub/endpoint/endpoint.proto)
    - [AddTerminationRequest](#onos.e2sub.endpoint.AddTerminationRequest)
    - [AddTerminationResponse](#onos.e2sub.endpoint.AddTerminationResponse)
    - [Event](#onos.e2sub.endpoint.Event)
    - [GetTerminationRequest](#onos.e2sub.endpoint.GetTerminationRequest)
    - [GetTerminationResponse](#onos.e2sub.endpoint.GetTerminationResponse)
    - [ListTerminationsRequest](#onos.e2sub.endpoint.ListTerminationsRequest)
    - [ListTerminationsResponse](#onos.e2sub.endpoint.ListTerminationsResponse)
    - [RemoveTerminationRequest](#onos.e2sub.endpoint.RemoveTerminationRequest)
    - [RemoveTerminationResponse](#onos.e2sub.endpoint.RemoveTerminationResponse)
    - [TerminationEndpoint](#onos.e2sub.endpoint.TerminationEndpoint)
    - [WatchTerminationsRequest](#onos.e2sub.endpoint.WatchTerminationsRequest)
    - [WatchTerminationsResponse](#onos.e2sub.endpoint.WatchTerminationsResponse)
  
    - [EventType](#onos.e2sub.endpoint.EventType)
  
    - [E2RegistryService](#onos.e2sub.endpoint.E2RegistryService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos/e2sub/endpoint/endpoint.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/e2sub/endpoint/endpoint.proto



<a name="onos.e2sub.endpoint.AddTerminationRequest"></a>

### AddTerminationRequest
AddTerminationRequest is a request for adding a new termination point


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| endpoint | [TerminationEndpoint](#onos.e2sub.endpoint.TerminationEndpoint) |  |  |






<a name="onos.e2sub.endpoint.AddTerminationResponse"></a>

### AddTerminationResponse
AddTerminationResponse is a response to adding a new termination point






<a name="onos.e2sub.endpoint.Event"></a>

### Event
Event is an end-point event


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [EventType](#onos.e2sub.endpoint.EventType) |  |  |
| endpoint | [TerminationEndpoint](#onos.e2sub.endpoint.TerminationEndpoint) |  |  |






<a name="onos.e2sub.endpoint.GetTerminationRequest"></a>

### GetTerminationRequest
GetTerminationRequest is a request for getting existing termination point


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="onos.e2sub.endpoint.GetTerminationResponse"></a>

### GetTerminationResponse
GetTerminationResponse is a response with invormation about a requested termination point


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| endpoint | [TerminationEndpoint](#onos.e2sub.endpoint.TerminationEndpoint) |  |  |






<a name="onos.e2sub.endpoint.ListTerminationsRequest"></a>

### ListTerminationsRequest
ListTerminationsRequest is a request to list all available E2 terminations






<a name="onos.e2sub.endpoint.ListTerminationsResponse"></a>

### ListTerminationsResponse
ListTerminationsResponse is a response to list all available E2 terminations


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| endpoints | [TerminationEndpoint](#onos.e2sub.endpoint.TerminationEndpoint) | repeated |  |






<a name="onos.e2sub.endpoint.RemoveTerminationRequest"></a>

### RemoveTerminationRequest
RemoveTerminationRequest is a request for removing termination point


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="onos.e2sub.endpoint.RemoveTerminationResponse"></a>

### RemoveTerminationResponse
RemoveTerminationResponse is a response to removing a termination point






<a name="onos.e2sub.endpoint.TerminationEndpoint"></a>

### TerminationEndpoint
Termination is a record identifying the IP address and TCP port coordinates where the E2 termination
service is available.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| revision | [uint64](#uint64) |  |  |
| ip | [string](#string) |  |  |
| port | [uint32](#uint32) |  |  |






<a name="onos.e2sub.endpoint.WatchTerminationsRequest"></a>

### WatchTerminationsRequest
WatchTerminationsRequest is a request to receive a stream of all E2 termination changes.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| noreplay | [bool](#bool) |  |  |






<a name="onos.e2sub.endpoint.WatchTerminationsResponse"></a>

### WatchTerminationsResponse
WatchTerminationsResponse is a response indicating a change in the available E2 termination end-points.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#onos.e2sub.endpoint.Event) |  |  |





 


<a name="onos.e2sub.endpoint.EventType"></a>

### EventType
Type of change

| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 |  |
| ADDED | 1 |  |
| REMOVED | 3 |  |


 

 


<a name="onos.e2sub.endpoint.E2RegistryService"></a>

### E2RegistryService
E2RegistryService manages subscription and subscription delete requests

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddTermination | [AddTerminationRequest](#onos.e2sub.endpoint.AddTerminationRequest) | [AddTerminationResponse](#onos.e2sub.endpoint.AddTerminationResponse) | AddTermination registers new E2 termination end-point. |
| GetTermination | [GetTerminationRequest](#onos.e2sub.endpoint.GetTerminationRequest) | [GetTerminationResponse](#onos.e2sub.endpoint.GetTerminationResponse) | GetTermination retrieves information about a specific end-point |
| RemoveTermination | [RemoveTerminationRequest](#onos.e2sub.endpoint.RemoveTerminationRequest) | [RemoveTerminationResponse](#onos.e2sub.endpoint.RemoveTerminationResponse) | RemoveTermination removes the specified E2 termination end-point. |
| ListTerminations | [ListTerminationsRequest](#onos.e2sub.endpoint.ListTerminationsRequest) | [ListTerminationsResponse](#onos.e2sub.endpoint.ListTerminationsResponse) | ListTerminations returns the list of currently registered E2 terminations. |
| WatchTerminations | [WatchTerminationsRequest](#onos.e2sub.endpoint.WatchTerminationsRequest) | [WatchTerminationsResponse](#onos.e2sub.endpoint.WatchTerminationsResponse) stream | WatchTerminations returns a stream of changes in the set of available E2 terminations. |

 



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

