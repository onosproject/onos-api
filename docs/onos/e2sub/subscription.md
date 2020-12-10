# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/e2sub/subscription/subscription.proto](#onos/e2sub/subscription/subscription.proto)
    - [Action](#onos.e2sub.subscription.Action)
    - [AddSubscriptionRequest](#onos.e2sub.subscription.AddSubscriptionRequest)
    - [AddSubscriptionResponse](#onos.e2sub.subscription.AddSubscriptionResponse)
    - [Event](#onos.e2sub.subscription.Event)
    - [EventTrigger](#onos.e2sub.subscription.EventTrigger)
    - [GetSubscriptionRequest](#onos.e2sub.subscription.GetSubscriptionRequest)
    - [GetSubscriptionResponse](#onos.e2sub.subscription.GetSubscriptionResponse)
    - [Lifecycle](#onos.e2sub.subscription.Lifecycle)
    - [ListSubscriptionsRequest](#onos.e2sub.subscription.ListSubscriptionsRequest)
    - [ListSubscriptionsResponse](#onos.e2sub.subscription.ListSubscriptionsResponse)
    - [Payload](#onos.e2sub.subscription.Payload)
    - [RemoveSubscriptionRequest](#onos.e2sub.subscription.RemoveSubscriptionRequest)
    - [RemoveSubscriptionResponse](#onos.e2sub.subscription.RemoveSubscriptionResponse)
    - [ServiceModel](#onos.e2sub.subscription.ServiceModel)
    - [Subscription](#onos.e2sub.subscription.Subscription)
    - [SubscriptionDetails](#onos.e2sub.subscription.SubscriptionDetails)
    - [SubsequentAction](#onos.e2sub.subscription.SubsequentAction)
    - [WatchSubscriptionsRequest](#onos.e2sub.subscription.WatchSubscriptionsRequest)
    - [WatchSubscriptionsResponse](#onos.e2sub.subscription.WatchSubscriptionsResponse)
  
    - [ActionType](#onos.e2sub.subscription.ActionType)
    - [Encoding](#onos.e2sub.subscription.Encoding)
    - [EventType](#onos.e2sub.subscription.EventType)
    - [Status](#onos.e2sub.subscription.Status)
    - [SubsequentActionType](#onos.e2sub.subscription.SubsequentActionType)
    - [TimeToWait](#onos.e2sub.subscription.TimeToWait)
  
    - [E2SubscriptionService](#onos.e2sub.subscription.E2SubscriptionService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos/e2sub/subscription/subscription.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/e2sub/subscription/subscription.proto



<a name="onos.e2sub.subscription.Action"></a>

### Action



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| type | [ActionType](#onos.e2sub.subscription.ActionType) |  |  |
| payload | [Payload](#onos.e2sub.subscription.Payload) |  |  |
| subsequent_action | [SubsequentAction](#onos.e2sub.subscription.SubsequentAction) |  |  |






<a name="onos.e2sub.subscription.AddSubscriptionRequest"></a>

### AddSubscriptionRequest
AddSubscriptionRequest a subscription request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscription | [Subscription](#onos.e2sub.subscription.Subscription) |  |  |






<a name="onos.e2sub.subscription.AddSubscriptionResponse"></a>

### AddSubscriptionResponse
AddSubscriptionResponse a subscription response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscription | [Subscription](#onos.e2sub.subscription.Subscription) |  |  |






<a name="onos.e2sub.subscription.Event"></a>

### Event
Event is a subscription event


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [EventType](#onos.e2sub.subscription.EventType) |  |  |
| subscription | [Subscription](#onos.e2sub.subscription.Subscription) |  |  |






<a name="onos.e2sub.subscription.EventTrigger"></a>

### EventTrigger



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payload | [Payload](#onos.e2sub.subscription.Payload) |  |  |






<a name="onos.e2sub.subscription.GetSubscriptionRequest"></a>

### GetSubscriptionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="onos.e2sub.subscription.GetSubscriptionResponse"></a>

### GetSubscriptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscription | [Subscription](#onos.e2sub.subscription.Subscription) |  |  |






<a name="onos.e2sub.subscription.Lifecycle"></a>

### Lifecycle
Lifecycle is the subscription lifecycle


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#onos.e2sub.subscription.Status) |  |  |






<a name="onos.e2sub.subscription.ListSubscriptionsRequest"></a>

### ListSubscriptionsRequest







<a name="onos.e2sub.subscription.ListSubscriptionsResponse"></a>

### ListSubscriptionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscriptions | [Subscription](#onos.e2sub.subscription.Subscription) | repeated |  |






<a name="onos.e2sub.subscription.Payload"></a>

### Payload



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| encoding | [Encoding](#onos.e2sub.subscription.Encoding) |  |  |
| data | [bytes](#bytes) |  |  |






<a name="onos.e2sub.subscription.RemoveSubscriptionRequest"></a>

### RemoveSubscriptionRequest
RemoveSubscriptionRequest a subscription delete request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="onos.e2sub.subscription.RemoveSubscriptionResponse"></a>

### RemoveSubscriptionResponse
RemoveSubscriptionResponse a subscription delete response






<a name="onos.e2sub.subscription.ServiceModel"></a>

### ServiceModel
ServiceModel is a service model definition


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="onos.e2sub.subscription.Subscription"></a>

### Subscription
Subscription is a subscription state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| revision | [uint64](#uint64) |  |  |
| details | [SubscriptionDetails](#onos.e2sub.subscription.SubscriptionDetails) |  |  |
| lifecycle | [Lifecycle](#onos.e2sub.subscription.Lifecycle) |  |  |






<a name="onos.e2sub.subscription.SubscriptionDetails"></a>

### SubscriptionDetails



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| app_id | [string](#string) |  |  |
| e2_node_id | [string](#string) |  |  |
| service_model | [ServiceModel](#onos.e2sub.subscription.ServiceModel) |  |  |
| event_trigger | [EventTrigger](#onos.e2sub.subscription.EventTrigger) |  |  |
| actions | [Action](#onos.e2sub.subscription.Action) | repeated |  |






<a name="onos.e2sub.subscription.SubsequentAction"></a>

### SubsequentAction
sequence from e2ap-v01.00.00.asn1:1132


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [SubsequentActionType](#onos.e2sub.subscription.SubsequentActionType) |  |  |
| time_to_wait | [TimeToWait](#onos.e2sub.subscription.TimeToWait) |  |  |






<a name="onos.e2sub.subscription.WatchSubscriptionsRequest"></a>

### WatchSubscriptionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| noreplay | [bool](#bool) |  |  |






<a name="onos.e2sub.subscription.WatchSubscriptionsResponse"></a>

### WatchSubscriptionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#onos.e2sub.subscription.Event) |  |  |





 


<a name="onos.e2sub.subscription.ActionType"></a>

### ActionType


| Name | Number | Description |
| ---- | ------ | ----------- |
| ACTION_TYPE_REPORT | 0 |  |
| ACTION_TYPE_INSERT | 1 |  |
| ACTION_TYPE_POLICY | 2 |  |



<a name="onos.e2sub.subscription.Encoding"></a>

### Encoding
Encoding indicates a payload encoding

| Name | Number | Description |
| ---- | ------ | ----------- |
| ENCODING_ASN1 | 0 |  |
| ENCODING_PROTO | 1 |  |



<a name="onos.e2sub.subscription.EventType"></a>

### EventType
EventType is a subscription event type

| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 |  |
| ADDED | 1 |  |
| UPDATED | 2 |  |
| REMOVED | 3 |  |



<a name="onos.e2sub.subscription.Status"></a>

### Status
Status is a subscription status

| Name | Number | Description |
| ---- | ------ | ----------- |
| ACTIVE | 0 |  |
| PENDING_DELETE | 1 |  |



<a name="onos.e2sub.subscription.SubsequentActionType"></a>

### SubsequentActionType


| Name | Number | Description |
| ---- | ------ | ----------- |
| SUBSEQUENT_ACTION_TYPE_CONTINUE | 0 |  |
| SUBSEQUENT_ACTION_TYPE_WAIT | 1 |  |



<a name="onos.e2sub.subscription.TimeToWait"></a>

### TimeToWait


| Name | Number | Description |
| ---- | ------ | ----------- |
| TIME_TO_WAIT_ZERO | 0 |  |
| TIME_TO_WAIT_W1MS | 1 |  |
| TIME_TO_WAIT_W2MS | 2 |  |
| TIME_TO_WAIT_W5MS | 3 |  |
| TIME_TO_WAIT_W10MS | 4 |  |
| TIME_TO_WAIT_W20MS | 5 |  |
| TIME_TO_WAIT_W30MS | 6 |  |
| TIME_TO_WAIT_W40MS | 7 |  |
| TIME_TO_WAIT_W50MS | 8 |  |
| TIME_TO_WAIT_W100MS | 9 |  |
| TIME_TO_WAIT_W200MS | 10 |  |
| TIME_TO_WAIT_W500MS | 11 |  |
| TIME_TO_WAIT_W1S | 12 |  |
| TIME_TO_WAIT_W2S | 13 |  |
| TIME_TO_WAIT_W5S | 14 |  |
| TIME_TO_WAIT_W10S | 15 |  |
| TIME_TO_WAIT_W20S | 16 |  |
| TIME_TO_WAIT_W60S | 17 |  |


 

 


<a name="onos.e2sub.subscription.E2SubscriptionService"></a>

### E2SubscriptionService
SubscriptionService manages subscription and subscription delete requests

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddSubscription | [AddSubscriptionRequest](#onos.e2sub.subscription.AddSubscriptionRequest) | [AddSubscriptionResponse](#onos.e2sub.subscription.AddSubscriptionResponse) | AddSubscription establishes E2 subscriptions on E2 Node. |
| RemoveSubscription | [RemoveSubscriptionRequest](#onos.e2sub.subscription.RemoveSubscriptionRequest) | [RemoveSubscriptionResponse](#onos.e2sub.subscription.RemoveSubscriptionResponse) | RemoveSubscription removes E2 subscriptions on E2 Node. |
| GetSubscription | [GetSubscriptionRequest](#onos.e2sub.subscription.GetSubscriptionRequest) | [GetSubscriptionResponse](#onos.e2sub.subscription.GetSubscriptionResponse) | GetSubscription retrieves information about a specific subscription in the list of existing subscriptions |
| ListSubscriptions | [ListSubscriptionsRequest](#onos.e2sub.subscription.ListSubscriptionsRequest) | [ListSubscriptionsResponse](#onos.e2sub.subscription.ListSubscriptionsResponse) | ListSubscriptions returns the list of current existing subscriptions |
| WatchSubscriptions | [WatchSubscriptionsRequest](#onos.e2sub.subscription.WatchSubscriptionsRequest) | [WatchSubscriptionsResponse](#onos.e2sub.subscription.WatchSubscriptionsResponse) stream | WatchSubscriptions returns a stream of subscription changes |

 



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

