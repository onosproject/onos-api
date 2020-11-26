# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/e2sub/task.proto](#api/e2sub/task.proto)
    - [Event](#e2sub.task.Event)
    - [GetSubscriptionTaskRequest](#e2sub.task.GetSubscriptionTaskRequest)
    - [GetSubscriptionTaskResponse](#e2sub.task.GetSubscriptionTaskResponse)
    - [Lifecycle](#e2sub.task.Lifecycle)
    - [ListSubscriptionTasksRequest](#e2sub.task.ListSubscriptionTasksRequest)
    - [ListSubscriptionTasksResponse](#e2sub.task.ListSubscriptionTasksResponse)
    - [SubscriptionTask](#e2sub.task.SubscriptionTask)
    - [UpdateSubscriptionTaskRequest](#e2sub.task.UpdateSubscriptionTaskRequest)
    - [UpdateSubscriptionTaskResponse](#e2sub.task.UpdateSubscriptionTaskResponse)
    - [WatchSubscriptionTasksRequest](#e2sub.task.WatchSubscriptionTasksRequest)
    - [WatchSubscriptionTasksResponse](#e2sub.task.WatchSubscriptionTasksResponse)
  
    - [EventType](#e2sub.task.EventType)
    - [Phase](#e2sub.task.Phase)
    - [Status](#e2sub.task.Status)
  
    - [E2SubscriptionTaskService](#e2sub.task.E2SubscriptionTaskService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api/e2sub/task.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/e2sub/task.proto



<a name="e2sub.task.Event"></a>

### Event
Event is a SubscriptionTask event


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [EventType](#e2sub.task.EventType) |  |  |
| task | [SubscriptionTask](#e2sub.task.SubscriptionTask) |  |  |






<a name="e2sub.task.GetSubscriptionTaskRequest"></a>

### GetSubscriptionTaskRequest
GetSubscriptionTaskRequest is a request for getting existing SubscriptionTask


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="e2sub.task.GetSubscriptionTaskResponse"></a>

### GetSubscriptionTaskResponse
GetSubscriptionTaskResponse is a response with invormation about a requested SubscriptionTask


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task | [SubscriptionTask](#e2sub.task.SubscriptionTask) |  |  |






<a name="e2sub.task.Lifecycle"></a>

### Lifecycle
Lifecycle is a subscription task status


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phase | [Phase](#e2sub.task.Phase) |  |  |
| status | [Status](#e2sub.task.Status) |  |  |






<a name="e2sub.task.ListSubscriptionTasksRequest"></a>

### ListSubscriptionTasksRequest
ListSubscriptionTasksRequest is a request to list all available SubscriptionTasks


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscription_id | [string](#string) |  |  |
| endpoint_id | [string](#string) |  |  |






<a name="e2sub.task.ListSubscriptionTasksResponse"></a>

### ListSubscriptionTasksResponse
ListSubscriptionTasksResponse is a response to list all available SubscriptionTasks


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tasks | [SubscriptionTask](#e2sub.task.SubscriptionTask) | repeated |  |






<a name="e2sub.task.SubscriptionTask"></a>

### SubscriptionTask
SubscriptionTask is a task representing a subscription between an E2 termination and an E2 node


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| revision | [uint64](#uint64) |  |  |
| subscription_id | [string](#string) |  |  |
| endpoint_id | [string](#string) |  |  |
| lifecycle | [Lifecycle](#e2sub.task.Lifecycle) |  |  |






<a name="e2sub.task.UpdateSubscriptionTaskRequest"></a>

### UpdateSubscriptionTaskRequest
UpdateSubscriptionTaskRequest is a request for updating a SubscriptionTask status


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task | [SubscriptionTask](#e2sub.task.SubscriptionTask) |  |  |






<a name="e2sub.task.UpdateSubscriptionTaskResponse"></a>

### UpdateSubscriptionTaskResponse
UpdateSubscriptionTaskResponse is a response to updating a SubscriptionTask status






<a name="e2sub.task.WatchSubscriptionTasksRequest"></a>

### WatchSubscriptionTasksRequest
WatchSubscriptionTasksRequest is a request to receive a stream of all SubscriptionTask changes.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| noreplay | [bool](#bool) |  |  |
| subscription_id | [string](#string) |  |  |
| endpoint_id | [string](#string) |  |  |






<a name="e2sub.task.WatchSubscriptionTasksResponse"></a>

### WatchSubscriptionTasksResponse
WatchSubscriptionTasksResponse is a response indicating a change in the available SubscriptionTasks.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#e2sub.task.Event) |  |  |





 


<a name="e2sub.task.EventType"></a>

### EventType
Type of change

| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 |  |
| CREATED | 1 |  |
| UPDATED | 2 |  |
| REMOVED | 3 |  |



<a name="e2sub.task.Phase"></a>

### Phase
Phase is a subscription task phase

| Name | Number | Description |
| ---- | ------ | ----------- |
| OPEN | 0 | OPEN is a subscription task open phase |
| CLOSE | 1 | CLOSE is a subscription task close phase |



<a name="e2sub.task.Status"></a>

### Status
Status is a subscription task status

| Name | Number | Description |
| ---- | ------ | ----------- |
| PENDING | 0 | PENDING indicates the subscription task phase is pending |
| COMPLETE | 1 | COMPLETE indicates the subscription task phase is complete |


 

 


<a name="e2sub.task.E2SubscriptionTaskService"></a>

### E2SubscriptionTaskService
E2SubscriptionTaskService manages subscription tasks between E2 termination points and E2 nodes

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetSubscriptionTask | [GetSubscriptionTaskRequest](#e2sub.task.GetSubscriptionTaskRequest) | [GetSubscriptionTaskResponse](#e2sub.task.GetSubscriptionTaskResponse) | GetSubscriptionTask retrieves information about a specific task |
| ListSubscriptionTasks | [ListSubscriptionTasksRequest](#e2sub.task.ListSubscriptionTasksRequest) | [ListSubscriptionTasksResponse](#e2sub.task.ListSubscriptionTasksResponse) | ListSubscriptionTasks returns the list of currently registered E2 Tasks. |
| WatchSubscriptionTasks | [WatchSubscriptionTasksRequest](#e2sub.task.WatchSubscriptionTasksRequest) | [WatchSubscriptionTasksResponse](#e2sub.task.WatchSubscriptionTasksResponse) stream | WatchSubscriptionTasks returns a stream of changes in the set of available E2 Tasks. |
| UpdateSubscriptionTask | [UpdateSubscriptionTaskRequest](#e2sub.task.UpdateSubscriptionTaskRequest) | [UpdateSubscriptionTaskResponse](#e2sub.task.UpdateSubscriptionTaskResponse) | UpdateSubscriptionTask updates a task status |

 



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

