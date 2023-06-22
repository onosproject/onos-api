# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/config/v3/transaction.proto](#onos_config_v3_transaction-proto)
    - [Transaction](#onos-config-v3-Transaction)
    - [Transaction.ValuesEntry](#onos-config-v3-Transaction-ValuesEntry)
    - [TransactionChangeStatus](#onos-config-v3-TransactionChangeStatus)
    - [TransactionEvent](#onos-config-v3-TransactionEvent)
    - [TransactionID](#onos-config-v3-TransactionID)
    - [TransactionPhaseStatus](#onos-config-v3-TransactionPhaseStatus)
    - [TransactionRollbackStatus](#onos-config-v3-TransactionRollbackStatus)
    - [TransactionRollbackStatus.ValuesEntry](#onos-config-v3-TransactionRollbackStatus-ValuesEntry)
    - [TransactionStatus](#onos-config-v3-TransactionStatus)
    - [TransactionStrategy](#onos-config-v3-TransactionStrategy)
  
    - [TransactionEvent.EventType](#onos-config-v3-TransactionEvent-EventType)
    - [TransactionPhaseStatus.State](#onos-config-v3-TransactionPhaseStatus-State)
    - [TransactionStatus.Phase](#onos-config-v3-TransactionStatus-Phase)
    - [TransactionStrategy.Synchronicity](#onos-config-v3-TransactionStrategy-Synchronicity)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos_config_v3_transaction-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/config/v3/transaction.proto



<a name="onos-config-v3-Transaction"></a>

### Transaction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meta | [ObjectMeta](#onos-config-v3-ObjectMeta) |  |  |
| id | [TransactionID](#onos-config-v3-TransactionID) |  | &#39;id&#39; is the unique identifier of the transaction |
| values | [Transaction.ValuesEntry](#onos-config-v3-Transaction-ValuesEntry) | repeated | &#39;values&#39; is a set of changes to apply to targets |
| status | [TransactionStatus](#onos-config-v3-TransactionStatus) |  | &#39;status&#39; is the transaction status |






<a name="onos-config-v3-Transaction-ValuesEntry"></a>

### Transaction.ValuesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [PathValue](#onos-config-v3-PathValue) |  |  |






<a name="onos-config-v3-TransactionChangeStatus"></a>

### TransactionChangeStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ordinal | [uint64](#uint64) |  |  |
| commit | [TransactionPhaseStatus](#onos-config-v3-TransactionPhaseStatus) |  |  |
| apply | [TransactionPhaseStatus](#onos-config-v3-TransactionPhaseStatus) |  |  |






<a name="onos-config-v3-TransactionEvent"></a>

### TransactionEvent
TransactionEvent transaction store event


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [TransactionEvent.EventType](#onos-config-v3-TransactionEvent-EventType) |  |  |
| transaction | [Transaction](#onos-config-v3-Transaction) |  |  |






<a name="onos-config-v3-TransactionID"></a>

### TransactionID



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| target | [Target](#onos-config-v3-Target) |  |  |
| index | [uint64](#uint64) |  |  |






<a name="onos-config-v3-TransactionPhaseStatus"></a>

### TransactionPhaseStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| end | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| state | [TransactionPhaseStatus.State](#onos-config-v3-TransactionPhaseStatus-State) |  |  |
| failure | [Failure](#onos-config-v3-Failure) |  |  |






<a name="onos-config-v3-TransactionRollbackStatus"></a>

### TransactionRollbackStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ordinal | [uint64](#uint64) |  |  |
| commit | [TransactionPhaseStatus](#onos-config-v3-TransactionPhaseStatus) |  |  |
| apply | [TransactionPhaseStatus](#onos-config-v3-TransactionPhaseStatus) |  |  |
| index | [uint64](#uint64) |  |  |
| values | [TransactionRollbackStatus.ValuesEntry](#onos-config-v3-TransactionRollbackStatus-ValuesEntry) | repeated |  |






<a name="onos-config-v3-TransactionRollbackStatus-ValuesEntry"></a>

### TransactionRollbackStatus.ValuesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [PathValue](#onos-config-v3-PathValue) |  |  |






<a name="onos-config-v3-TransactionStatus"></a>

### TransactionStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phase | [TransactionStatus.Phase](#onos-config-v3-TransactionStatus-Phase) |  |  |
| change | [TransactionChangeStatus](#onos-config-v3-TransactionChangeStatus) |  |  |
| rollback | [TransactionRollbackStatus](#onos-config-v3-TransactionRollbackStatus) |  |  |






<a name="onos-config-v3-TransactionStrategy"></a>

### TransactionStrategy



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| synchronicity | [TransactionStrategy.Synchronicity](#onos-config-v3-TransactionStrategy-Synchronicity) |  | &#39;synchronicity&#39; indicates the transaction synchronicity level |





 


<a name="onos-config-v3-TransactionEvent-EventType"></a>

### TransactionEvent.EventType
EventType transaction event types for transaction store

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 |  |
| CREATED | 1 |  |
| UPDATED | 2 |  |
| DELETED | 3 |  |
| REPLAYED | 4 |  |



<a name="onos-config-v3-TransactionPhaseStatus-State"></a>

### TransactionPhaseStatus.State


| Name | Number | Description |
| ---- | ------ | ----------- |
| PENDING | 0 |  |
| IN_PROGRESS | 1 |  |
| COMPLETE | 2 |  |
| ABORTED | 3 |  |
| CANCELED | 4 |  |
| FAILED | 5 |  |



<a name="onos-config-v3-TransactionStatus-Phase"></a>

### TransactionStatus.Phase


| Name | Number | Description |
| ---- | ------ | ----------- |
| CHANGE | 0 |  |
| ROLLBACK | 1 |  |



<a name="onos-config-v3-TransactionStrategy-Synchronicity"></a>

### TransactionStrategy.Synchronicity


| Name | Number | Description |
| ---- | ------ | ----------- |
| ASYNCHRONOUS | 0 |  |
| SYNCHRONOUS | 1 |  |


 

 

 



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

