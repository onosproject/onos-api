# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/config/v3/proposal.proto](#onos_config_v3_proposal-proto)
    - [Proposal](#onos-config-v3-Proposal)
    - [Proposal.ValuesEntry](#onos-config-v3-Proposal-ValuesEntry)
    - [ProposalEvent](#onos-config-v3-ProposalEvent)
    - [ProposalID](#onos-config-v3-ProposalID)
    - [ProposalInfo](#onos-config-v3-ProposalInfo)
    - [ProposalPhaseStatus](#onos-config-v3-ProposalPhaseStatus)
    - [ProposalStatus](#onos-config-v3-ProposalStatus)
  
    - [Proposal.Type](#onos-config-v3-Proposal-Type)
    - [ProposalEvent.EventType](#onos-config-v3-ProposalEvent-EventType)
    - [ProposalPhaseStatus.State](#onos-config-v3-ProposalPhaseStatus-State)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos_config_v3_proposal-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/config/v3/proposal.proto



<a name="onos-config-v3-Proposal"></a>

### Proposal
Proposal refers to a proposal change or proposal rollback


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meta | [ObjectMeta](#onos-config-v3-ObjectMeta) |  |  |
| id | [ProposalID](#onos-config-v3-ProposalID) |  | &#39;id&#39; is the unique identifier of the proposal This field should be set prior to persisting the object. |
| type | [Proposal.Type](#onos-config-v3-Proposal-Type) |  | &#39;type&#39; is the proposal type |
| values | [Proposal.ValuesEntry](#onos-config-v3-Proposal-ValuesEntry) | repeated | &#39;values&#39; is a set of changes to apply to targets |
| status | [ProposalStatus](#onos-config-v3-ProposalStatus) |  | &#39;status&#39; is the current lifecycle status of the proposal |






<a name="onos-config-v3-Proposal-ValuesEntry"></a>

### Proposal.ValuesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [PathValue](#onos-config-v3-PathValue) |  |  |






<a name="onos-config-v3-ProposalEvent"></a>

### ProposalEvent
ProposalEvent proposal store event


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [ProposalEvent.EventType](#onos-config-v3-ProposalEvent-EventType) |  |  |
| proposal | [Proposal](#onos-config-v3-Proposal) |  |  |






<a name="onos-config-v3-ProposalID"></a>

### ProposalID



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| target | [Target](#onos-config-v3-Target) |  |  |
| index | [uint64](#uint64) |  |  |






<a name="onos-config-v3-ProposalInfo"></a>

### ProposalInfo
ProposalInfo is an extension providing information about the proposal
to clients in responses.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| id | [ProposalID](#onos-config-v3-ProposalID) |  |  |






<a name="onos-config-v3-ProposalPhaseStatus"></a>

### ProposalPhaseStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| end | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| state | [ProposalPhaseStatus.State](#onos-config-v3-ProposalPhaseStatus-State) |  |  |
| failure | [Failure](#onos-config-v3-Failure) |  |  |






<a name="onos-config-v3-ProposalStatus"></a>

### ProposalStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| commit | [ProposalPhaseStatus](#onos-config-v3-ProposalPhaseStatus) |  | &#39;commit&#39; is the proposal commit phase status |
| apply | [ProposalPhaseStatus](#onos-config-v3-ProposalPhaseStatus) |  | &#39;apply&#39; is the proposal apply phase status |





 


<a name="onos-config-v3-Proposal-Type"></a>

### Proposal.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| CHANGE | 0 |  |
| ROLLBACK | 1 |  |



<a name="onos-config-v3-ProposalEvent-EventType"></a>

### ProposalEvent.EventType
EventType proposal event types for proposal store

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 |  |
| CREATED | 1 |  |
| UPDATED | 2 |  |
| DELETED | 3 |  |
| REPLAYED | 4 |  |



<a name="onos-config-v3-ProposalPhaseStatus-State"></a>

### ProposalPhaseStatus.State


| Name | Number | Description |
| ---- | ------ | ----------- |
| PENDING | 0 |  |
| COMPLETE | 1 |  |
| ABORTED | 2 |  |
| FAILED | 3 |  |


 

 

 



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

