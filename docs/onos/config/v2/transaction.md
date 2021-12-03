# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/config/v2/transaction.proto](#onos/config/v2/transaction.proto)
    - [Change](#onos.config.v2.Change)
    - [ChangeValue](#onos.config.v2.ChangeValue)
    - [Transaction](#onos.config.v2.Transaction)
    - [TransactionEvent](#onos.config.v2.TransactionEvent)
    - [TransactionRef](#onos.config.v2.TransactionRef)
    - [TransactionStatus](#onos.config.v2.TransactionStatus)
  
    - [TransactionEventType](#onos.config.v2.TransactionEventType)
    - [TransactionPhase](#onos.config.v2.TransactionPhase)
    - [TransactionState](#onos.config.v2.TransactionState)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos/config/v2/transaction.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/config/v2/transaction.proto



<a name="onos.config.v2.Change"></a>

### Change
Change represents a configuration change to a single target


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| target_id | [string](#string) |  | &#39;target_id&#39; is the identifier of the target to which this change applies |
| target_version | [string](#string) |  | &#39;target_version&#39; is an optional target version to which to apply this change |
| target_type | [string](#string) |  | &#39;target_type&#39; is an optional target type to which to apply this change |
| values | [ChangeValue](#onos.config.v2.ChangeValue) | repeated | &#39;values&#39; is a set of change values to apply |






<a name="onos.config.v2.ChangeValue"></a>

### ChangeValue
ChangeValue is an individual Path/Value and removed flag combination in a Change


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| path | [string](#string) |  | &#39;path&#39; is the path to change |
| value | [TypedValue](#onos.config.v2.TypedValue) |  | &#39;value&#39; is the change value |
| removed | [bool](#bool) |  | &#39;removed&#39; indicates whether this is a delete |






<a name="onos.config.v2.Transaction"></a>

### Transaction
Transaction refers to a multi-target transactional change


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | &#39;id&#39; is the unique identifier of the transaction This field should be set prior to persisting the object. |
| index | [uint64](#uint64) |  | &#39;index&#39; is a monotonically increasing, globally unique index of the change The index is provided by the store, is static and unique for each unique change identifier, and should not be modified by client code. |
| revision | [uint64](#uint64) |  | &#39;revision&#39; is the change revision number The revision number is provided by the store and should not be modified by client code. Each unique state of the change will be assigned a unique revision number which can be used for optimistic concurrency control when updating or deleting the change state. |
| status | [TransactionStatus](#onos.config.v2.TransactionStatus) |  | &#39;status&#39; is the current lifecycle status of the transaction |
| created | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | &#39;created&#39; is the time at which the transaction was created |
| updated | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | &#39;updated&#39; is the time at which the transaction was last updated |
| changes | [Change](#onos.config.v2.Change) | repeated | &#39;changes&#39; is a set of changes to apply to targets The list of changes should contain only a single change per target/version pair. |
| deleted | [bool](#bool) |  | &#39;deleted&#39; is a flag indicating whether this transaction is being deleted by a snapshot |
| dependency | [TransactionRef](#onos.config.v2.TransactionRef) |  | &#39;dependency&#39; is a reference to the transaction on which this transaction is dependent |
| dependents | [TransactionRef](#onos.config.v2.TransactionRef) | repeated | &#39;dependents&#39; is a list of references to transactions that depend on this transaction |
| username | [string](#string) |  | &#39;username&#39; is the name of the user that made the transaction |






<a name="onos.config.v2.TransactionEvent"></a>

### TransactionEvent
TransactionEvent transaction store event


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [TransactionEventType](#onos.config.v2.TransactionEventType) |  |  |
| transaction | [Transaction](#onos.config.v2.Transaction) |  |  |






<a name="onos.config.v2.TransactionRef"></a>

### TransactionRef
TransactionRef is a reference to a transaction


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| none | [google.protobuf.Empty](#google.protobuf.Empty) |  |  |
| transaction_id | [string](#string) |  |  |






<a name="onos.config.v2.TransactionStatus"></a>

### TransactionStatus
TransactionStatus is the status of a Transaction


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phase | [TransactionPhase](#onos.config.v2.TransactionPhase) |  | &#39;phase&#39; is the current phase of the |
| state | [TransactionState](#onos.config.v2.TransactionState) |  | &#39;state&#39; is the state of the transaction within a Phase |





 


<a name="onos.config.v2.TransactionEventType"></a>

### TransactionEventType
TransactionEventType transaction event types for transaction store

| Name | Number | Description |
| ---- | ------ | ----------- |
| TRANSACTION_EVENT_UNKNOWN | 0 |  |
| TRANSACTION_CREATED | 1 |  |
| TRANSACTION_UPDATED | 2 |  |
| TRANSACTION_DELETED | 3 |  |
| TRANSACTION_REPLAYED | 4 |  |



<a name="onos.config.v2.TransactionPhase"></a>

### TransactionPhase
TransactionPhase is the phase of a Transaction

| Name | Number | Description |
| ---- | ------ | ----------- |
| TRANSACTION_CHANGE | 0 | TRANSACTION_CHANGE indicates the transaction has been requested |
| TRANSACTION_ROLLBACK | 1 | TRANSACTION_ROLLBACK indicates a rollback has been requested for the transaction |



<a name="onos.config.v2.TransactionState"></a>

### TransactionState
TransactionState is the transaction state of a transaction phase

| Name | Number | Description |
| ---- | ------ | ----------- |
| TRANSACTION_PENDING | 0 | TRANSACTION_PENDING indicates the transaction is pending |
| TRANSACTION_COMPLETE | 2 | TRANSACTION_COMPLETE indicates the transaction is complete |
| TRANSACTION_FAILED | 3 | TRANSACTION_FAILED indicates the transaction failed |
| TRANSACTION_VALIDATING | 4 | TRANSACTION_VALIDATING indicates the transaction is in the validating state |
| TRANSACTION_VALIDATED | 5 | TRANSACTION_VALIDATED indicates the transaction is validated successfully |
| TRANSACTION_VALIDATION_FAILED | 6 | TRANSACTION_VALIDATION_FAILED indicates the transaction validation is failed |


 

 

 



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

