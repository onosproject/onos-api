# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/config/v2/proposal.proto](#onos/config/v2/proposal.proto)
    - [ChangeProposal](#onos.config.v2.ChangeProposal)
    - [ChangeProposal.ValuesEntry](#onos.config.v2.ChangeProposal.ValuesEntry)
    - [Proposal](#onos.config.v2.Proposal)
    - [ProposalAbortPhase](#onos.config.v2.ProposalAbortPhase)
    - [ProposalApplyPhase](#onos.config.v2.ProposalApplyPhase)
    - [ProposalCommitPhase](#onos.config.v2.ProposalCommitPhase)
    - [ProposalEvent](#onos.config.v2.ProposalEvent)
    - [ProposalInitializePhase](#onos.config.v2.ProposalInitializePhase)
    - [ProposalPhaseStatus](#onos.config.v2.ProposalPhaseStatus)
    - [ProposalPhases](#onos.config.v2.ProposalPhases)
    - [ProposalStatus](#onos.config.v2.ProposalStatus)
    - [ProposalStatus.RollbackValuesEntry](#onos.config.v2.ProposalStatus.RollbackValuesEntry)
    - [ProposalValidatePhase](#onos.config.v2.ProposalValidatePhase)
    - [RollbackProposal](#onos.config.v2.RollbackProposal)
  
    - [ProposalAbortPhase.State](#onos.config.v2.ProposalAbortPhase.State)
    - [ProposalApplyPhase.State](#onos.config.v2.ProposalApplyPhase.State)
    - [ProposalCommitPhase.State](#onos.config.v2.ProposalCommitPhase.State)
    - [ProposalEvent.EventType](#onos.config.v2.ProposalEvent.EventType)
    - [ProposalInitializePhase.State](#onos.config.v2.ProposalInitializePhase.State)
    - [ProposalValidatePhase.State](#onos.config.v2.ProposalValidatePhase.State)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos/config/v2/proposal.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/config/v2/proposal.proto



<a name="onos.config.v2.ChangeProposal"></a>

### ChangeProposal



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| values | [ChangeProposal.ValuesEntry](#onos.config.v2.ChangeProposal.ValuesEntry) | repeated | &#39;changes&#39; is the proposed change values |






<a name="onos.config.v2.ChangeProposal.ValuesEntry"></a>

### ChangeProposal.ValuesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [PathValue](#onos.config.v2.PathValue) |  |  |






<a name="onos.config.v2.Proposal"></a>

### Proposal



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meta | [ObjectMeta](#onos.config.v2.ObjectMeta) |  |  |
| id | [string](#string) |  | &#39;id&#39; is the unique identifier of the proposal |
| target_id | [string](#string) |  | &#39;target_id&#39; is the proposal&#39;s target identifier |
| transaction_index | [uint64](#uint64) |  | &#39;transaction_index&#39; is the unique index of the transaction |
| change | [ChangeProposal](#onos.config.v2.ChangeProposal) |  |  |
| rollback | [RollbackProposal](#onos.config.v2.RollbackProposal) |  |  |
| status | [ProposalStatus](#onos.config.v2.ProposalStatus) |  | &#39;status&#39; is the current lifecycle status of the proposal |






<a name="onos.config.v2.ProposalAbortPhase"></a>

### ProposalAbortPhase



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [ProposalPhaseStatus](#onos.config.v2.ProposalPhaseStatus) |  |  |
| state | [ProposalAbortPhase.State](#onos.config.v2.ProposalAbortPhase.State) |  |  |






<a name="onos.config.v2.ProposalApplyPhase"></a>

### ProposalApplyPhase



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [ProposalPhaseStatus](#onos.config.v2.ProposalPhaseStatus) |  |  |
| state | [ProposalApplyPhase.State](#onos.config.v2.ProposalApplyPhase.State) |  |  |
| term | [uint64](#uint64) |  |  |
| failure | [Failure](#onos.config.v2.Failure) |  |  |






<a name="onos.config.v2.ProposalCommitPhase"></a>

### ProposalCommitPhase



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [ProposalPhaseStatus](#onos.config.v2.ProposalPhaseStatus) |  |  |
| state | [ProposalCommitPhase.State](#onos.config.v2.ProposalCommitPhase.State) |  |  |






<a name="onos.config.v2.ProposalEvent"></a>

### ProposalEvent
ProposalEvent proposal store event


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [ProposalEvent.EventType](#onos.config.v2.ProposalEvent.EventType) |  |  |
| proposal | [Proposal](#onos.config.v2.Proposal) |  |  |






<a name="onos.config.v2.ProposalInitializePhase"></a>

### ProposalInitializePhase



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [ProposalPhaseStatus](#onos.config.v2.ProposalPhaseStatus) |  |  |
| state | [ProposalInitializePhase.State](#onos.config.v2.ProposalInitializePhase.State) |  |  |






<a name="onos.config.v2.ProposalPhaseStatus"></a>

### ProposalPhaseStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| end | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="onos.config.v2.ProposalPhases"></a>

### ProposalPhases



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| initialize | [ProposalInitializePhase](#onos.config.v2.ProposalInitializePhase) |  | &#39;initialize&#39; is the proposal initialization phase status |
| validate | [ProposalValidatePhase](#onos.config.v2.ProposalValidatePhase) |  | &#39;validate&#39; is the proposal validation phase status |
| commit | [ProposalCommitPhase](#onos.config.v2.ProposalCommitPhase) |  | &#39;commit&#39; is the proposal commit phase status |
| apply | [ProposalApplyPhase](#onos.config.v2.ProposalApplyPhase) |  | &#39;apply&#39; is the proposal apply phase status |
| abort | [ProposalAbortPhase](#onos.config.v2.ProposalAbortPhase) |  | &#39;abort&#39; is the proposal abort phase status |






<a name="onos.config.v2.ProposalStatus"></a>

### ProposalStatus
ProposalStatus is the status of a Proposal


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phases | [ProposalPhases](#onos.config.v2.ProposalPhases) |  | &#39;phases&#39; is the proposal phases |
| prev_index | [uint64](#uint64) |  | &#39;prev_index&#39; is the index of the previous proposal associated with this target |
| next_index | [uint64](#uint64) |  | &#39;next_index&#39; is the index of the next proposal associated with this target |
| rollback_index | [uint64](#uint64) |  | &#39;rollback_index&#39; is a reference to the index to which to roll back |
| rollback_values | [ProposalStatus.RollbackValuesEntry](#onos.config.v2.ProposalStatus.RollbackValuesEntry) | repeated | &#39;rollback_values&#39; is the set of values to use to roll back the proposal |






<a name="onos.config.v2.ProposalStatus.RollbackValuesEntry"></a>

### ProposalStatus.RollbackValuesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [PathValue](#onos.config.v2.PathValue) |  |  |






<a name="onos.config.v2.ProposalValidatePhase"></a>

### ProposalValidatePhase



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [ProposalPhaseStatus](#onos.config.v2.ProposalPhaseStatus) |  |  |
| state | [ProposalValidatePhase.State](#onos.config.v2.ProposalValidatePhase.State) |  |  |
| failure | [Failure](#onos.config.v2.Failure) |  |  |






<a name="onos.config.v2.RollbackProposal"></a>

### RollbackProposal



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rollback_index | [uint64](#uint64) |  | &#39;rollback_index&#39; is the index of the transaction to roll back |





 


<a name="onos.config.v2.ProposalAbortPhase.State"></a>

### ProposalAbortPhase.State


| Name | Number | Description |
| ---- | ------ | ----------- |
| ABORTING | 0 |  |
| ABORTED | 1 |  |



<a name="onos.config.v2.ProposalApplyPhase.State"></a>

### ProposalApplyPhase.State


| Name | Number | Description |
| ---- | ------ | ----------- |
| APPLYING | 0 |  |
| APPLIED | 1 |  |
| FAILED | 2 |  |



<a name="onos.config.v2.ProposalCommitPhase.State"></a>

### ProposalCommitPhase.State


| Name | Number | Description |
| ---- | ------ | ----------- |
| COMMITTING | 0 |  |
| COMMITTED | 1 |  |



<a name="onos.config.v2.ProposalEvent.EventType"></a>

### ProposalEvent.EventType
EventType proposal event types for proposal store

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 |  |
| CREATED | 1 |  |
| UPDATED | 2 |  |
| DELETED | 3 |  |
| REPLAYED | 4 |  |



<a name="onos.config.v2.ProposalInitializePhase.State"></a>

### ProposalInitializePhase.State


| Name | Number | Description |
| ---- | ------ | ----------- |
| INITIALIZING | 0 |  |
| INITIALIZED | 1 |  |



<a name="onos.config.v2.ProposalValidatePhase.State"></a>

### ProposalValidatePhase.State


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

