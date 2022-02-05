# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/config/v2/transaction.proto](#onos/config/v2/transaction.proto)
    - [ChangeTransaction](#onos.config.v2.ChangeTransaction)
    - [ChangeTransaction.ValuesEntry](#onos.config.v2.ChangeTransaction.ValuesEntry)
    - [RollbackTransaction](#onos.config.v2.RollbackTransaction)
    - [Transaction](#onos.config.v2.Transaction)
    - [TransactionApplyPhase](#onos.config.v2.TransactionApplyPhase)
    - [TransactionCommitPhase](#onos.config.v2.TransactionCommitPhase)
    - [TransactionEvent](#onos.config.v2.TransactionEvent)
    - [TransactionInfo](#onos.config.v2.TransactionInfo)
    - [TransactionInitializePhase](#onos.config.v2.TransactionInitializePhase)
    - [TransactionPhaseStatus](#onos.config.v2.TransactionPhaseStatus)
    - [TransactionPhases](#onos.config.v2.TransactionPhases)
    - [TransactionStatus](#onos.config.v2.TransactionStatus)
    - [TransactionStrategy](#onos.config.v2.TransactionStrategy)
    - [TransactionValidatePhase](#onos.config.v2.TransactionValidatePhase)
  
    - [TransactionApplyPhase.State](#onos.config.v2.TransactionApplyPhase.State)
    - [TransactionCommitPhase.State](#onos.config.v2.TransactionCommitPhase.State)
    - [TransactionEvent.EventType](#onos.config.v2.TransactionEvent.EventType)
    - [TransactionInitializePhase.State](#onos.config.v2.TransactionInitializePhase.State)
    - [TransactionStatus.State](#onos.config.v2.TransactionStatus.State)
    - [TransactionStrategy.Atomicity](#onos.config.v2.TransactionStrategy.Atomicity)
    - [TransactionStrategy.Isolation](#onos.config.v2.TransactionStrategy.Isolation)
    - [TransactionStrategy.Synchronicity](#onos.config.v2.TransactionStrategy.Synchronicity)
    - [TransactionValidatePhase.State](#onos.config.v2.TransactionValidatePhase.State)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos/config/v2/transaction.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/config/v2/transaction.proto



<a name="onos.config.v2.ChangeTransaction"></a>

### ChangeTransaction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| values | [ChangeTransaction.ValuesEntry](#onos.config.v2.ChangeTransaction.ValuesEntry) | repeated | &#39;values&#39; is a set of changes to apply to targets |






<a name="onos.config.v2.ChangeTransaction.ValuesEntry"></a>

### ChangeTransaction.ValuesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [PathValues](#onos.config.v2.PathValues) |  |  |






<a name="onos.config.v2.RollbackTransaction"></a>

### RollbackTransaction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rollback_index | [uint64](#uint64) |  | &#39;rollback_index&#39; is the index of the transaction to roll back |






<a name="onos.config.v2.Transaction"></a>

### Transaction
Transaction refers to a transaction change or transaction rollback


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meta | [ObjectMeta](#onos.config.v2.ObjectMeta) |  |  |
| id | [string](#string) |  | &#39;id&#39; is the unique identifier of the transaction This field should be set prior to persisting the object. |
| index | [uint64](#uint64) |  | &#39;index&#39; is a monotonically increasing, globally unique index of the change The index is provided by the store, is static and unique for each unique change identifier, and should not be modified by client code. |
| username | [string](#string) |  | &#39;username&#39; is the name of the user that made the transaction |
| strategy | [TransactionStrategy](#onos.config.v2.TransactionStrategy) |  | &#39;strategy&#39; is the transaction strategy |
| change | [ChangeTransaction](#onos.config.v2.ChangeTransaction) |  |  |
| rollback | [RollbackTransaction](#onos.config.v2.RollbackTransaction) |  |  |
| status | [TransactionStatus](#onos.config.v2.TransactionStatus) |  | &#39;status&#39; is the current lifecycle status of the transaction |






<a name="onos.config.v2.TransactionApplyPhase"></a>

### TransactionApplyPhase



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [TransactionPhaseStatus](#onos.config.v2.TransactionPhaseStatus) |  |  |
| state | [TransactionApplyPhase.State](#onos.config.v2.TransactionApplyPhase.State) |  |  |






<a name="onos.config.v2.TransactionCommitPhase"></a>

### TransactionCommitPhase



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [TransactionPhaseStatus](#onos.config.v2.TransactionPhaseStatus) |  |  |
| state | [TransactionCommitPhase.State](#onos.config.v2.TransactionCommitPhase.State) |  |  |






<a name="onos.config.v2.TransactionEvent"></a>

### TransactionEvent
TransactionEvent transaction store event


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [TransactionEvent.EventType](#onos.config.v2.TransactionEvent.EventType) |  |  |
| transaction | [Transaction](#onos.config.v2.Transaction) |  |  |






<a name="onos.config.v2.TransactionInfo"></a>

### TransactionInfo
TransactionInfo is an extension providing information about the transaction
to clients in responses.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| index | [uint64](#uint64) |  |  |






<a name="onos.config.v2.TransactionInitializePhase"></a>

### TransactionInitializePhase



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [TransactionPhaseStatus](#onos.config.v2.TransactionPhaseStatus) |  |  |
| state | [TransactionInitializePhase.State](#onos.config.v2.TransactionInitializePhase.State) |  |  |
| failure | [Failure](#onos.config.v2.Failure) |  |  |






<a name="onos.config.v2.TransactionPhaseStatus"></a>

### TransactionPhaseStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| end | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="onos.config.v2.TransactionPhases"></a>

### TransactionPhases



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| initialize | [TransactionInitializePhase](#onos.config.v2.TransactionInitializePhase) |  | &#39;initialize&#39; is the proposal initialization phase status |
| validate | [TransactionValidatePhase](#onos.config.v2.TransactionValidatePhase) |  | &#39;validate&#39; is the proposal validation phase status |
| commit | [TransactionCommitPhase](#onos.config.v2.TransactionCommitPhase) |  | &#39;commit&#39; is the proposal commit phase status |
| apply | [TransactionApplyPhase](#onos.config.v2.TransactionApplyPhase) |  | &#39;apply&#39; is the proposal apply phase status |






<a name="onos.config.v2.TransactionStatus"></a>

### TransactionStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phases | [TransactionPhases](#onos.config.v2.TransactionPhases) |  | &#39;phases&#39; is the transaction phases |
| proposals | [string](#string) | repeated | &#39;proposals&#39; is the set of proposals managed by the transaction |
| state | [TransactionStatus.State](#onos.config.v2.TransactionStatus.State) |  | &#39;state&#39; is the overall transaction state |
| failure | [Failure](#onos.config.v2.Failure) |  | &#39;failure&#39; is the transaction failure (if any) |






<a name="onos.config.v2.TransactionStrategy"></a>

### TransactionStrategy



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| synchronicity | [TransactionStrategy.Synchronicity](#onos.config.v2.TransactionStrategy.Synchronicity) |  | &#39;synchronicity&#39; indicates the transaction synchronicity level |
| atomicity | [TransactionStrategy.Atomicity](#onos.config.v2.TransactionStrategy.Atomicity) |  | &#39;atomicity&#39; indicates the transaction atomicity level |
| isolation | [TransactionStrategy.Isolation](#onos.config.v2.TransactionStrategy.Isolation) |  | &#39;isolation&#39; indicates the transaction isolation level |






<a name="onos.config.v2.TransactionValidatePhase"></a>

### TransactionValidatePhase



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [TransactionPhaseStatus](#onos.config.v2.TransactionPhaseStatus) |  |  |
| state | [TransactionValidatePhase.State](#onos.config.v2.TransactionValidatePhase.State) |  |  |
| failure | [Failure](#onos.config.v2.Failure) |  |  |





 


<a name="onos.config.v2.TransactionApplyPhase.State"></a>

### TransactionApplyPhase.State


| Name | Number | Description |
| ---- | ------ | ----------- |
| APPLYING | 0 |  |
| APPLIED | 1 |  |



<a name="onos.config.v2.TransactionCommitPhase.State"></a>

### TransactionCommitPhase.State


| Name | Number | Description |
| ---- | ------ | ----------- |
| COMMITTING | 0 |  |
| COMMITTED | 1 |  |



<a name="onos.config.v2.TransactionEvent.EventType"></a>

### TransactionEvent.EventType
EventType transaction event types for transaction store

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 |  |
| CREATED | 1 |  |
| UPDATED | 2 |  |
| DELETED | 3 |  |
| REPLAYED | 4 |  |



<a name="onos.config.v2.TransactionInitializePhase.State"></a>

### TransactionInitializePhase.State


| Name | Number | Description |
| ---- | ------ | ----------- |
| INITIALIZING | 0 |  |
| INITIALIZED | 1 |  |
| FAILED | 2 |  |



<a name="onos.config.v2.TransactionStatus.State"></a>

### TransactionStatus.State


| Name | Number | Description |
| ---- | ------ | ----------- |
| PENDING | 0 |  |
| VALIDATED | 1 |  |
| COMMITTED | 2 |  |
| APPLIED | 3 |  |
| FAILED | 4 |  |



<a name="onos.config.v2.TransactionStrategy.Atomicity"></a>

### TransactionStrategy.Atomicity


| Name | Number | Description |
| ---- | ------ | ----------- |
| ATOMIC | 0 |  |
| NONATOMIC | 1 |  |



<a name="onos.config.v2.TransactionStrategy.Isolation"></a>

### TransactionStrategy.Isolation


| Name | Number | Description |
| ---- | ------ | ----------- |
| DEFAULT | 0 |  |
| SERIALIZABLE | 1 |  |



<a name="onos.config.v2.TransactionStrategy.Synchronicity"></a>

### TransactionStrategy.Synchronicity


| Name | Number | Description |
| ---- | ------ | ----------- |
| ASYNCHRONOUS | 0 |  |
| SYNCHRONOUS | 1 |  |



<a name="onos.config.v2.TransactionValidatePhase.State"></a>

### TransactionValidatePhase.State


| Name | Number | Description |
| ---- | ------ | ----------- |
| VALIDATING | 0 |  |
| VALIDATED | 1 |  |
| FAILED | 2 |  |


 

 

 



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

