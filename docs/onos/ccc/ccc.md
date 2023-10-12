# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/ccc/ccc.proto](#onos_ccc_ccc-proto)
    - [Ack](#onos-ccc-Ack)
    - [ORRmpolicyRatio](#onos-ccc-ORRmpolicyRatio)
    - [Plmnidentity](#onos-ccc-Plmnidentity)
    - [RrmPolicyMember](#onos-ccc-RrmPolicyMember)
    - [RrmPolicyMemberList](#onos-ccc-RrmPolicyMemberList)
    - [SNSsai](#onos-ccc-SNSsai)
    - [Sd](#onos-ccc-Sd)
    - [Sst](#onos-ccc-Sst)
    - [UpdateSliceRequest](#onos-ccc-UpdateSliceRequest)
    - [UpdateSliceResponse](#onos-ccc-UpdateSliceResponse)
  
    - [ResourceType](#onos-ccc-ResourceType)
    - [SchedulerType](#onos-ccc-SchedulerType)
  
    - [Ccc](#onos-ccc-Ccc)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos_ccc_ccc-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/ccc/ccc.proto



<a name="onos-ccc-Ack"></a>

### Ack



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |
| cause | [string](#string) |  |  |






<a name="onos-ccc-ORRmpolicyRatio"></a>

### ORRmpolicyRatio
sequence from e2sm_ccc.asn1:229
{O-RRMPolicyRatio}


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resource_type | [ResourceType](#onos-ccc-ResourceType) |  |  |
| scheduler_type | [SchedulerType](#onos-ccc-SchedulerType) |  |  |
| rrm_policy_member_list | [RrmPolicyMemberList](#onos-ccc-RrmPolicyMemberList) |  |  |
| rrm_policy_max_ratio | [int32](#int32) |  |  |
| rrm_policy_min_ratio | [int32](#int32) |  |  |
| rrm_policy_dedicated_ratio | [int32](#int32) |  |  |






<a name="onos-ccc-Plmnidentity"></a>

### Plmnidentity
range of Integer from e2sm_common_ies.asn1:444


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [bytes](#bytes) |  |  |






<a name="onos-ccc-RrmPolicyMember"></a>

### RrmPolicyMember
sequence from e2sm_ccc.asn1:161


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| plmn_id | [Plmnidentity](#onos-ccc-Plmnidentity) |  |  |
| snssai | [SNSsai](#onos-ccc-SNSsai) |  |  |






<a name="onos-ccc-RrmPolicyMemberList"></a>

### RrmPolicyMemberList
sequence from e2sm_ccc.asn1:167


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rrm_policy_member | [RrmPolicyMember](#onos-ccc-RrmPolicyMember) | repeated |  |






<a name="onos-ccc-SNSsai"></a>

### SNSsai
sequence from e2sm_common_ies.asn1:454


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sst | [Sst](#onos-ccc-Sst) |  |  |
| sd | [Sd](#onos-ccc-Sd) |  |  |






<a name="onos-ccc-Sd"></a>

### Sd
range of Integer from e2sm_common_ies.asn1:450


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [bytes](#bytes) |  |  |






<a name="onos-ccc-Sst"></a>

### Sst
range of Integer from e2sm_common_ies.asn1:460


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [bytes](#bytes) |  |  |






<a name="onos-ccc-UpdateSliceRequest"></a>

### UpdateSliceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| e2_node_id | [string](#string) |  |  |
| rrm_policy_ratio | [ORRmpolicyRatio](#onos-ccc-ORRmpolicyRatio) |  |  |






<a name="onos-ccc-UpdateSliceResponse"></a>

### UpdateSliceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ack | [Ack](#onos-ccc-Ack) |  |  |





 


<a name="onos-ccc-ResourceType"></a>

### ResourceType
enumerated from e2sm_ccc.asn1:141

| Name | Number | Description |
| ---- | ------ | ----------- |
| RESOURCE_TYPE_PRB_UL | 0 |  |
| RESOURCE_TYPE_PRB_DL | 1 |  |
| RESOURCE_TYPE_DRB | 2 |  |
| RESOURCE_TYPE_RRC | 3 |  |



<a name="onos-ccc-SchedulerType"></a>

### SchedulerType
enumerated from e2sm_ccc.asn1:143

| Name | Number | Description |
| ---- | ------ | ----------- |
| SCHEDULER_TYPE_ROUND_ROBIN | 0 |  |
| SCHEDULER_TYPE_PROPORTIONALLY_FAIR | 1 |  |
| SCHEDULER_TYPE_QOS_BASED | 2 |  |


 

 


<a name="onos-ccc-Ccc"></a>

### Ccc


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| UpdateSlice | [UpdateSliceRequest](#onos-ccc-UpdateSliceRequest) | [UpdateSliceResponse](#onos-ccc-UpdateSliceResponse) | Slice management |

 



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

