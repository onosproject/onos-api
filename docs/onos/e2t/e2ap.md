# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/e2t/e2ap/control.proto](#onos/e2t/e2ap/control.proto)
    - [ControlRequest](#onos.e2t.e2ap.ControlRequest)
    - [ControlResponse](#onos.e2t.e2ap.ControlResponse)
  
    - [ControlService](#onos.e2t.e2ap.ControlService)
  
- [onos/e2t/e2ap/e2ap.proto](#onos/e2t/e2ap/e2ap.proto)
    - [Error](#onos.e2t.e2ap.Error)
    - [Error.Cause](#onos.e2t.e2ap.Error.Cause)
    - [Error.Cause.Misc](#onos.e2t.e2ap.Error.Cause.Misc)
    - [Error.Cause.Protocol](#onos.e2t.e2ap.Error.Cause.Protocol)
    - [Error.Cause.Ric](#onos.e2t.e2ap.Error.Cause.Ric)
    - [Error.Cause.RicService](#onos.e2t.e2ap.Error.Cause.RicService)
    - [Error.Cause.Transport](#onos.e2t.e2ap.Error.Cause.Transport)
    - [Error.Cause.Unknown](#onos.e2t.e2ap.Error.Cause.Unknown)
    - [Payload](#onos.e2t.e2ap.Payload)
    - [ServiceModel](#onos.e2t.e2ap.ServiceModel)
    - [Source](#onos.e2t.e2ap.Source)
    - [Target](#onos.e2t.e2ap.Target)
  
    - [EncodingType](#onos.e2t.e2ap.EncodingType)
    - [Error.Cause.Misc.Type](#onos.e2t.e2ap.Error.Cause.Misc.Type)
    - [Error.Cause.Protocol.Type](#onos.e2t.e2ap.Error.Cause.Protocol.Type)
    - [Error.Cause.Ric.Type](#onos.e2t.e2ap.Error.Cause.Ric.Type)
    - [Error.Cause.RicService.Type](#onos.e2t.e2ap.Error.Cause.RicService.Type)
    - [Error.Cause.Transport.Type](#onos.e2t.e2ap.Error.Cause.Transport.Type)
  
- [onos/e2t/e2ap/subscription.proto](#onos/e2t/e2ap/subscription.proto)
    - [Action](#onos.e2t.e2ap.Action)
    - [EventTrigger](#onos.e2t.e2ap.EventTrigger)
    - [Indication](#onos.e2t.e2ap.Indication)
    - [SubscribeRequest](#onos.e2t.e2ap.SubscribeRequest)
    - [SubscribeResponse](#onos.e2t.e2ap.SubscribeResponse)
    - [Subscription](#onos.e2t.e2ap.Subscription)
    - [SubsequentAction](#onos.e2t.e2ap.SubsequentAction)
  
    - [ActionType](#onos.e2t.e2ap.ActionType)
    - [SubsequentActionType](#onos.e2t.e2ap.SubsequentActionType)
    - [TimeToWait](#onos.e2t.e2ap.TimeToWait)
  
    - [SubscriptionService](#onos.e2t.e2ap.SubscriptionService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos/e2t/e2ap/control.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/e2t/e2ap/control.proto



<a name="onos.e2t.e2ap.ControlRequest"></a>

### ControlRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| source | [Source](#onos.e2t.e2ap.Source) |  |  |
| target | [Target](#onos.e2t.e2ap.Target) |  |  |
| service_model | [ServiceModel](#onos.e2t.e2ap.ServiceModel) |  |  |
| header | [Payload](#onos.e2t.e2ap.Payload) |  |  |
| message | [Payload](#onos.e2t.e2ap.Payload) |  |  |






<a name="onos.e2t.e2ap.ControlResponse"></a>

### ControlResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| outcome | [Payload](#onos.e2t.e2ap.Payload) |  |  |





 

 

 


<a name="onos.e2t.e2ap.ControlService"></a>

### ControlService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Control | [ControlRequest](#onos.e2t.e2ap.ControlRequest) | [ControlResponse](#onos.e2t.e2ap.ControlResponse) |  |

 



<a name="onos/e2t/e2ap/e2ap.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/e2t/e2ap/e2ap.proto



<a name="onos.e2t.e2ap.Error"></a>

### Error
Error is an E2AP protocol error


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cause | [Error.Cause](#onos.e2t.e2ap.Error.Cause) |  |  |






<a name="onos.e2t.e2ap.Error.Cause"></a>

### Error.Cause



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| unknown | [Error.Cause.Unknown](#onos.e2t.e2ap.Error.Cause.Unknown) |  |  |
| protocol | [Error.Cause.Protocol](#onos.e2t.e2ap.Error.Cause.Protocol) |  |  |
| ric | [Error.Cause.Ric](#onos.e2t.e2ap.Error.Cause.Ric) |  |  |
| ric_service | [Error.Cause.RicService](#onos.e2t.e2ap.Error.Cause.RicService) |  |  |
| transport | [Error.Cause.Transport](#onos.e2t.e2ap.Error.Cause.Transport) |  |  |






<a name="onos.e2t.e2ap.Error.Cause.Misc"></a>

### Error.Cause.Misc







<a name="onos.e2t.e2ap.Error.Cause.Protocol"></a>

### Error.Cause.Protocol







<a name="onos.e2t.e2ap.Error.Cause.Ric"></a>

### Error.Cause.Ric







<a name="onos.e2t.e2ap.Error.Cause.RicService"></a>

### Error.Cause.RicService







<a name="onos.e2t.e2ap.Error.Cause.Transport"></a>

### Error.Cause.Transport







<a name="onos.e2t.e2ap.Error.Cause.Unknown"></a>

### Error.Cause.Unknown







<a name="onos.e2t.e2ap.Payload"></a>

### Payload



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| encoding_type | [EncodingType](#onos.e2t.e2ap.EncodingType) |  |  |
| data | [bytes](#bytes) |  |  |






<a name="onos.e2t.e2ap.ServiceModel"></a>

### ServiceModel



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| version | [string](#string) |  |  |






<a name="onos.e2t.e2ap.Source"></a>

### Source



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| app_id | [string](#string) |  |  |
| app_node_id | [string](#string) |  |  |






<a name="onos.e2t.e2ap.Target"></a>

### Target



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| e2_node_id | [string](#string) |  |  |





 


<a name="onos.e2t.e2ap.EncodingType"></a>

### EncodingType


| Name | Number | Description |
| ---- | ------ | ----------- |
| PROTO | 0 |  |
| ASN1_PER | 1 |  |
| ASN1_XER | 2 |  |



<a name="onos.e2t.e2ap.Error.Cause.Misc.Type"></a>

### Error.Cause.Misc.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNSPECIFIED | 0 |  |
| CONTROL_PROCESSING_OVERLOAD | 1 |  |
| HARDWARE_FAILURE | 2 |  |
| OM_INTERVENTION | 3 |  |



<a name="onos.e2t.e2ap.Error.Cause.Protocol.Type"></a>

### Error.Cause.Protocol.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNSPECIFIED | 0 |  |
| TRANSFER_SYNTAX_ERROR | 1 |  |
| ABSTRACT_SYNTAX_ERROR_REJECT | 2 |  |
| ABSTRACT_SYNTAX_ERROR_IGNORE_AND_NOTIFY | 3 |  |
| MESSAGE_NOT_COMPATIBLE_WITH_RECEIVER_STATE | 4 |  |
| SEMANTIC_ERROR | 5 |  |
| ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE | 6 |  |



<a name="onos.e2t.e2ap.Error.Cause.Ric.Type"></a>

### Error.Cause.Ric.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNSPECIFIED | 0 |  |
| RAN_FUNCTION_ID_INVALID | 1 |  |
| ACTION_NOT_SUPPORTED | 2 |  |
| EXCESSIVE_ACTIONS | 3 |  |
| DUPLICATE_ACTION | 4 |  |
| DUPLICATE_EVENT | 5 |  |
| FUNCTION_RESOURCE_LIMIT | 6 |  |
| REQUEST_ID_UNKNOWN | 7 |  |
| INCONSISTENT_ACTION_SUBSEQUENT_ACTION_SEQUENCE | 8 |  |
| CONTROL_MESSAGE_INVALID | 9 |  |
| CALL_PROCESS_ID_INVALID | 10 |  |



<a name="onos.e2t.e2ap.Error.Cause.RicService.Type"></a>

### Error.Cause.RicService.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNSPECIFIED | 0 |  |
| FUNCTION_NOT_REQUIRED | 1 |  |
| EXCESSIVE_FUNCTIONS | 2 |  |
| RIC_RESOURCE_LIMIT | 3 |  |



<a name="onos.e2t.e2ap.Error.Cause.Transport.Type"></a>

### Error.Cause.Transport.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNSPECIFIED | 0 |  |
| TRANSPORT_RESOURCE_UNAVAILABLE | 1 |  |


 

 

 



<a name="onos/e2t/e2ap/subscription.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/e2t/e2ap/subscription.proto



<a name="onos.e2t.e2ap.Action"></a>

### Action



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| type | [ActionType](#onos.e2t.e2ap.ActionType) |  |  |
| payload | [Payload](#onos.e2t.e2ap.Payload) |  |  |
| subsequent_action | [SubsequentAction](#onos.e2t.e2ap.SubsequentAction) |  |  |






<a name="onos.e2t.e2ap.EventTrigger"></a>

### EventTrigger



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payload | [Payload](#onos.e2t.e2ap.Payload) |  |  |






<a name="onos.e2t.e2ap.Indication"></a>

### Indication



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [Payload](#onos.e2t.e2ap.Payload) |  |  |
| message | [Payload](#onos.e2t.e2ap.Payload) |  |  |






<a name="onos.e2t.e2ap.SubscribeRequest"></a>

### SubscribeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| source | [Source](#onos.e2t.e2ap.Source) |  |  |
| target | [Target](#onos.e2t.e2ap.Target) |  |  |
| subscription | [Subscription](#onos.e2t.e2ap.Subscription) |  |  |






<a name="onos.e2t.e2ap.SubscribeResponse"></a>

### SubscribeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| indication | [Indication](#onos.e2t.e2ap.Indication) |  |  |






<a name="onos.e2t.e2ap.Subscription"></a>

### Subscription



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscription_id | [string](#string) |  |  |
| event_trigger | [EventTrigger](#onos.e2t.e2ap.EventTrigger) |  |  |
| actions | [Action](#onos.e2t.e2ap.Action) | repeated |  |






<a name="onos.e2t.e2ap.SubsequentAction"></a>

### SubsequentAction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [SubsequentActionType](#onos.e2t.e2ap.SubsequentActionType) |  |  |
| time_to_wait | [TimeToWait](#onos.e2t.e2ap.TimeToWait) |  |  |





 


<a name="onos.e2t.e2ap.ActionType"></a>

### ActionType


| Name | Number | Description |
| ---- | ------ | ----------- |
| ACTION_TYPE_REPORT | 0 |  |
| ACTION_TYPE_INSERT | 1 |  |
| ACTION_TYPE_POLICY | 2 |  |



<a name="onos.e2t.e2ap.SubsequentActionType"></a>

### SubsequentActionType


| Name | Number | Description |
| ---- | ------ | ----------- |
| SUBSEQUENT_ACTION_TYPE_CONTINUE | 0 |  |
| SUBSEQUENT_ACTION_TYPE_WAIT | 1 |  |



<a name="onos.e2t.e2ap.TimeToWait"></a>

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


 

 


<a name="onos.e2t.e2ap.SubscriptionService"></a>

### SubscriptionService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Subscribe | [SubscribeRequest](#onos.e2t.e2ap.SubscribeRequest) | [SubscribeResponse](#onos.e2t.e2ap.SubscribeResponse) stream |  |

 



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

