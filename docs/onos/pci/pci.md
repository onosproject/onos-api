# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/pci/pci.proto](#onos/pci/pci.proto)
    - [GetCellRequest](#onos.pci.GetCellRequest)
    - [GetCellResponse](#onos.pci.GetCellResponse)
    - [GetCellsRequest](#onos.pci.GetCellsRequest)
    - [GetCellsResponse](#onos.pci.GetCellsResponse)
    - [GetConflictsRequest](#onos.pci.GetConflictsRequest)
    - [GetConflictsResponse](#onos.pci.GetConflictsResponse)
    - [PciCell](#onos.pci.PciCell)
    - [PciRange](#onos.pci.PciRange)
  
    - [CellType](#onos.pci.CellType)
  
    - [Pci](#onos.pci.Pci)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos/pci/pci.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/pci/pci.proto



<a name="onos.pci.GetCellRequest"></a>

### GetCellRequest
cell id required


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cell_id | [uint64](#uint64) |  |  |






<a name="onos.pci.GetCellResponse"></a>

### GetCellResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cell | [PciCell](#onos.pci.PciCell) |  |  |






<a name="onos.pci.GetCellsRequest"></a>

### GetCellsRequest
cell id required






<a name="onos.pci.GetCellsResponse"></a>

### GetCellsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cells | [PciCell](#onos.pci.PciCell) | repeated |  |






<a name="onos.pci.GetConflictsRequest"></a>

### GetConflictsRequest
if cell id is not specified, will return all cells with conflicts


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cell_id | [uint64](#uint64) |  |  |






<a name="onos.pci.GetConflictsResponse"></a>

### GetConflictsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cells | [PciCell](#onos.pci.PciCell) | repeated |  |






<a name="onos.pci.PciCell"></a>

### PciCell



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| node_id | [string](#string) |  |  |
| dlearfcn | [uint32](#uint32) |  |  |
| cell_type | [CellType](#onos.pci.CellType) |  |  |
| pci | [uint32](#uint32) |  |  |
| pci_pool | [PciRange](#onos.pci.PciRange) | repeated |  |
| neighbor_ids | [uint64](#uint64) | repeated |  |






<a name="onos.pci.PciRange"></a>

### PciRange



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| min | [uint32](#uint32) |  |  |
| max | [uint32](#uint32) |  |  |





 


<a name="onos.pci.CellType"></a>

### CellType


| Name | Number | Description |
| ---- | ------ | ----------- |
| FEMTO | 0 |  |
| ENTERPRISE | 1 |  |
| OUTDOOR_SMALL | 2 |  |
| MACRO | 3 |  |


 

 


<a name="onos.pci.Pci"></a>

### Pci


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetConflicts | [GetConflictsRequest](#onos.pci.GetConflictsRequest) | [GetConflictsResponse](#onos.pci.GetConflictsResponse) |  |
| GetCell | [GetCellRequest](#onos.pci.GetCellRequest) | [GetCellResponse](#onos.pci.GetCellResponse) |  |
| GetCells | [GetCellsRequest](#onos.pci.GetCellsRequest) | [GetCellsResponse](#onos.pci.GetCellsResponse) |  |

 



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

