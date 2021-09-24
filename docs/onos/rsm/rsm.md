# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/rsm/rsm.proto](#onos/rsm/rsm.proto)
    - [Ack](#onos.rsm.Ack)
    - [CreateSliceRequest](#onos.rsm.CreateSliceRequest)
    - [CreateSliceResponse](#onos.rsm.CreateSliceResponse)
    - [DeleteSliceRequest](#onos.rsm.DeleteSliceRequest)
    - [DeleteSliceResponse](#onos.rsm.DeleteSliceResponse)
    - [GetSliceResponse](#onos.rsm.GetSliceResponse)
    - [GetSlicesRequest](#onos.rsm.GetSlicesRequest)
    - [GetUeSliceAssociationRequest](#onos.rsm.GetUeSliceAssociationRequest)
    - [GetUeSliceAssociationResponse](#onos.rsm.GetUeSliceAssociationResponse)
    - [SetUeSliceAssociationRequest](#onos.rsm.SetUeSliceAssociationRequest)
    - [SetUeSliceAssociationResponse](#onos.rsm.SetUeSliceAssociationResponse)
    - [SliceAssocItem](#onos.rsm.SliceAssocItem)
    - [SliceItem](#onos.rsm.SliceItem)
    - [UeId](#onos.rsm.UeId)
    - [UpdateSliceRequest](#onos.rsm.UpdateSliceRequest)
    - [UpdateSliceResponse](#onos.rsm.UpdateSliceResponse)
  
    - [SchedulerType](#onos.rsm.SchedulerType)
    - [SliceType](#onos.rsm.SliceType)
    - [UeIdType](#onos.rsm.UeIdType)
  
    - [Rsm](#onos.rsm.Rsm)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos/rsm/rsm.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/rsm/rsm.proto



<a name="onos.rsm.Ack"></a>

### Ack



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |
| cause | [string](#string) |  |  |






<a name="onos.rsm.CreateSliceRequest"></a>

### CreateSliceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| e2_node_id | [string](#string) |  |  |
| slice_id | [string](#string) |  |  |
| scheduler_type | [SchedulerType](#onos.rsm.SchedulerType) |  |  |
| weight | [string](#string) |  |  |
| sliceType | [SliceType](#onos.rsm.SliceType) |  |  |






<a name="onos.rsm.CreateSliceResponse"></a>

### CreateSliceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ack | [Ack](#onos.rsm.Ack) |  |  |






<a name="onos.rsm.DeleteSliceRequest"></a>

### DeleteSliceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| e2_node_id | [string](#string) |  |  |
| slice_id | [string](#string) |  |  |
| sliceType | [SliceType](#onos.rsm.SliceType) |  |  |






<a name="onos.rsm.DeleteSliceResponse"></a>

### DeleteSliceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ack | [Ack](#onos.rsm.Ack) |  |  |






<a name="onos.rsm.GetSliceResponse"></a>

### GetSliceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ack | [Ack](#onos.rsm.Ack) |  |  |
| slice_items | [SliceItem](#onos.rsm.SliceItem) | repeated |  |






<a name="onos.rsm.GetSlicesRequest"></a>

### GetSlicesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| e2_node_id | [string](#string) |  |  |






<a name="onos.rsm.GetUeSliceAssociationRequest"></a>

### GetUeSliceAssociationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| e2_node_id | [string](#string) |  |  |
| ue_id | [UeId](#onos.rsm.UeId) | repeated |  |
| slice_id | [string](#string) |  |  |
| ue_slice_assoc_id | [string](#string) |  |  |






<a name="onos.rsm.GetUeSliceAssociationResponse"></a>

### GetUeSliceAssociationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ack | [Ack](#onos.rsm.Ack) |  |  |
| slice_assoc_items | [SliceAssocItem](#onos.rsm.SliceAssocItem) | repeated |  |






<a name="onos.rsm.SetUeSliceAssociationRequest"></a>

### SetUeSliceAssociationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| e2_node_id | [string](#string) |  |  |
| ue_id | [UeId](#onos.rsm.UeId) | repeated |  |
| slice_id | [string](#string) |  |  |






<a name="onos.rsm.SetUeSliceAssociationResponse"></a>

### SetUeSliceAssociationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ack | [Ack](#onos.rsm.Ack) |  |  |
| assigned_ue_slice_assoc_id | [string](#string) |  |  |






<a name="onos.rsm.SliceAssocItem"></a>

### SliceAssocItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ue_slice_assoc_id | [string](#string) |  |  |
| e2_node_id | [string](#string) |  |  |
| ue_id | [UeIdType](#onos.rsm.UeIdType) | repeated |  |
| slice_id | [string](#string) |  |  |






<a name="onos.rsm.SliceItem"></a>

### SliceItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| e2_node_id | [string](#string) |  |  |
| slice_ids | [string](#string) | repeated |  |






<a name="onos.rsm.UeId"></a>

### UeId



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ue_id | [string](#string) |  |  |
| type | [UeIdType](#onos.rsm.UeIdType) |  |  |






<a name="onos.rsm.UpdateSliceRequest"></a>

### UpdateSliceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| e2_node_id | [string](#string) |  |  |
| slice_id | [string](#string) |  |  |
| scheduler_type | [SchedulerType](#onos.rsm.SchedulerType) |  |  |
| weight | [string](#string) |  |  |
| sliceType | [SliceType](#onos.rsm.SliceType) |  |  |






<a name="onos.rsm.UpdateSliceResponse"></a>

### UpdateSliceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ack | [Ack](#onos.rsm.Ack) |  |  |





 


<a name="onos.rsm.SchedulerType"></a>

### SchedulerType


| Name | Number | Description |
| ---- | ------ | ----------- |
| SCHEDULER_TYPE_ROUND_ROBIN | 0 |  |
| SCHEDULER_TYPE_PROPORTIONALLY_FAIR | 1 |  |
| SCHEDULER_TYPE_QOS_BASED | 2 |  |



<a name="onos.rsm.SliceType"></a>

### SliceType


| Name | Number | Description |
| ---- | ------ | ----------- |
| SLICE_TYPE_DL_SLICE | 0 |  |
| SLICE_TYPE_UL_SLICE | 1 |  |



<a name="onos.rsm.UeIdType"></a>

### UeIdType


| Name | Number | Description |
| ---- | ------ | ----------- |
| UE_ID_TYPE_CU_UE_F1_AP_ID | 0 |  |
| UE_ID_TYPE_DU_UE_F1_AP_ID | 1 |  |
| UE_ID_TYPE_RAN_UE_NGAP_ID | 2 |  |
| UE_ID_TYPE_AMF_UE_NGAP_ID | 3 |  |
| UE_ID_TYPE_ENB_UE_S1_AP_ID | 4 |  |


 

 


<a name="onos.rsm.Rsm"></a>

### Rsm


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetSlices | [GetSlicesRequest](#onos.rsm.GetSlicesRequest) | [GetSliceResponse](#onos.rsm.GetSliceResponse) | Slice management |
| CreateSlice | [CreateSliceRequest](#onos.rsm.CreateSliceRequest) | [CreateSliceResponse](#onos.rsm.CreateSliceResponse) |  |
| UpdateSlice | [UpdateSliceRequest](#onos.rsm.UpdateSliceRequest) | [UpdateSliceResponse](#onos.rsm.UpdateSliceResponse) |  |
| DeleteSlice | [DeleteSliceRequest](#onos.rsm.DeleteSliceRequest) | [DeleteSliceResponse](#onos.rsm.DeleteSliceResponse) |  |
| GetUeSliceAssociation | [GetUeSliceAssociationRequest](#onos.rsm.GetUeSliceAssociationRequest) | [GetUeSliceAssociationResponse](#onos.rsm.GetUeSliceAssociationResponse) | UE-Slice association |
| SetUeSliceAssociation | [SetUeSliceAssociationRequest](#onos.rsm.SetUeSliceAssociationRequest) | [SetUeSliceAssociationResponse](#onos.rsm.SetUeSliceAssociationResponse) |  |

 



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

