# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/e2sub/subscription.proto](#api/e2sub/subscription.proto)
    - [AddSubscriptionRequest](#e2sub.subscription.AddSubscriptionRequest)
    - [AddSubscriptionResponse](#e2sub.subscription.AddSubscriptionResponse)
    - [Event](#e2sub.subscription.Event)
    - [GetSubscriptionRequest](#e2sub.subscription.GetSubscriptionRequest)
    - [GetSubscriptionResponse](#e2sub.subscription.GetSubscriptionResponse)
    - [Lifecycle](#e2sub.subscription.Lifecycle)
    - [ListSubscriptionsRequest](#e2sub.subscription.ListSubscriptionsRequest)
    - [ListSubscriptionsResponse](#e2sub.subscription.ListSubscriptionsResponse)
    - [Payload](#e2sub.subscription.Payload)
    - [RemoveSubscriptionRequest](#e2sub.subscription.RemoveSubscriptionRequest)
    - [RemoveSubscriptionResponse](#e2sub.subscription.RemoveSubscriptionResponse)
    - [ServiceModel](#e2sub.subscription.ServiceModel)
    - [Subscription](#e2sub.subscription.Subscription)
    - [WatchSubscriptionsRequest](#e2sub.subscription.WatchSubscriptionsRequest)
    - [WatchSubscriptionsResponse](#e2sub.subscription.WatchSubscriptionsResponse)
  
    - [Encoding](#e2sub.subscription.Encoding)
    - [EventType](#e2sub.subscription.EventType)
    - [Status](#e2sub.subscription.Status)
  
    - [E2SubscriptionService](#e2sub.subscription.E2SubscriptionService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api/e2sub/subscription.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/e2sub/subscription.proto



<a name="e2sub.subscription.AddSubscriptionRequest"></a>

### AddSubscriptionRequest
AddSubscriptionRequest a subscription request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscription | [Subscription](#e2sub.subscription.Subscription) |  |  |






<a name="e2sub.subscription.AddSubscriptionResponse"></a>

### AddSubscriptionResponse
AddSubscriptionResponse a subscription response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscription | [Subscription](#e2sub.subscription.Subscription) |  |  |






<a name="e2sub.subscription.Event"></a>

### Event
Event is a subscription event


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [EventType](#e2sub.subscription.EventType) |  |  |
| subscription | [Subscription](#e2sub.subscription.Subscription) |  |  |






<a name="e2sub.subscription.GetSubscriptionRequest"></a>

### GetSubscriptionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="e2sub.subscription.GetSubscriptionResponse"></a>

### GetSubscriptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscription | [Subscription](#e2sub.subscription.Subscription) |  |  |






<a name="e2sub.subscription.Lifecycle"></a>

### Lifecycle
Lifecycle is the subscription lifecycle


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#e2sub.subscription.Status) |  |  |






<a name="e2sub.subscription.ListSubscriptionsRequest"></a>

### ListSubscriptionsRequest







<a name="e2sub.subscription.ListSubscriptionsResponse"></a>

### ListSubscriptionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscriptions | [Subscription](#e2sub.subscription.Subscription) | repeated |  |






<a name="e2sub.subscription.Payload"></a>

### Payload
Payload is a subscription payload


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| encoding | [Encoding](#e2sub.subscription.Encoding) |  |  |
| bytes | [bytes](#bytes) |  |  |






<a name="e2sub.subscription.RemoveSubscriptionRequest"></a>

### RemoveSubscriptionRequest
RemoveSubscriptionRequest a subscription delete request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="e2sub.subscription.RemoveSubscriptionResponse"></a>

### RemoveSubscriptionResponse
RemoveSubscriptionResponse a subscription delete response






<a name="e2sub.subscription.ServiceModel"></a>

### ServiceModel
ServiceModel is a service model definition


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="e2sub.subscription.Subscription"></a>

### Subscription
Subscription is a subscription state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| revision | [uint64](#uint64) |  |  |
| app_id | [string](#string) |  |  |
| e2_node_id | [string](#string) |  |  |
| service_model | [ServiceModel](#e2sub.subscription.ServiceModel) |  |  |
| payload | [Payload](#e2sub.subscription.Payload) |  |  |
| lifecycle | [Lifecycle](#e2sub.subscription.Lifecycle) |  |  |






<a name="e2sub.subscription.WatchSubscriptionsRequest"></a>

### WatchSubscriptionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| noreplay | [bool](#bool) |  |  |






<a name="e2sub.subscription.WatchSubscriptionsResponse"></a>

### WatchSubscriptionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#e2sub.subscription.Event) |  |  |





 


<a name="e2sub.subscription.Encoding"></a>

### Encoding
Encoding indicates a payload encoding

| Name | Number | Description |
| ---- | ------ | ----------- |
| ENCODING_ASN1 | 0 |  |
| ENCODING_PROTO | 1 |  |



<a name="e2sub.subscription.EventType"></a>

### EventType
EventType is a subscription event type

| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 |  |
| ADDED | 1 |  |
| UPDATED | 2 |  |
| REMOVED | 3 |  |



<a name="e2sub.subscription.Status"></a>

### Status
Status is a subscription status

| Name | Number | Description |
| ---- | ------ | ----------- |
| ACTIVE | 0 |  |
| PENDING_DELETE | 1 |  |


 

 


<a name="e2sub.subscription.E2SubscriptionService"></a>

### E2SubscriptionService
SubscriptionService manages subscription and subscription delete requests

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddSubscription | [AddSubscriptionRequest](#e2sub.subscription.AddSubscriptionRequest) | [AddSubscriptionResponse](#e2sub.subscription.AddSubscriptionResponse) | AddSubscription establishes E2 subscriptions on E2 Node. |
| RemoveSubscription | [RemoveSubscriptionRequest](#e2sub.subscription.RemoveSubscriptionRequest) | [RemoveSubscriptionResponse](#e2sub.subscription.RemoveSubscriptionResponse) | RemoveSubscription removes E2 subscriptions on E2 Node. |
| GetSubscription | [GetSubscriptionRequest](#e2sub.subscription.GetSubscriptionRequest) | [GetSubscriptionResponse](#e2sub.subscription.GetSubscriptionResponse) | GetSubscription retrieves information about a specific subscription in the list of existing subscriptions |
| ListSubscriptions | [ListSubscriptionsRequest](#e2sub.subscription.ListSubscriptionsRequest) | [ListSubscriptionsResponse](#e2sub.subscription.ListSubscriptionsResponse) | ListSubscriptions returns the list of current existing subscriptions |
| WatchSubscriptions | [WatchSubscriptionsRequest](#e2sub.subscription.WatchSubscriptionsRequest) | [WatchSubscriptionsResponse](#e2sub.subscription.WatchSubscriptionsResponse) stream | WatchSubscriptions returns a stream of subscription changes |

 



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

