# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/pci/pci.proto](#onos/pci/pci.proto)
    - [GetMetricRequest](#onos.pci.GetMetricRequest)
    - [GetMetricResponse](#onos.pci.GetMetricResponse)
    - [GetMetricResponse.MetricsEntry](#onos.pci.GetMetricResponse.MetricsEntry)
    - [GetNeigborsRequest](#onos.pci.GetNeigborsRequest)
    - [GetNeigborsResponse](#onos.pci.GetNeigborsResponse)
    - [GetNumConflictsReponse](#onos.pci.GetNumConflictsReponse)
    - [GetNumConflictsRequest](#onos.pci.GetNumConflictsRequest)
    - [GetPciPoolRequest](#onos.pci.GetPciPoolRequest)
    - [GetPciPoolResponse](#onos.pci.GetPciPoolResponse)
    - [GetPciPoolResponse.PoolsEntry](#onos.pci.GetPciPoolResponse.PoolsEntry)
    - [Metrics](#onos.pci.Metrics)
    - [PciRange](#onos.pci.PciRange)
  
    - [CellType](#onos.pci.CellType)
  
    - [Pci](#onos.pci.Pci)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos/pci/pci.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/pci/pci.proto



<a name="onos.pci.GetMetricRequest"></a>

### GetMetricRequest
if cell id is not specified, will return all metrics


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cell_id | [uint64](#uint64) |  |  |






<a name="onos.pci.GetMetricResponse"></a>

### GetMetricResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metrics | [GetMetricResponse.MetricsEntry](#onos.pci.GetMetricResponse.MetricsEntry) | repeated |  |






<a name="onos.pci.GetMetricResponse.MetricsEntry"></a>

### GetMetricResponse.MetricsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [uint64](#uint64) |  |  |
| value | [Metrics](#onos.pci.Metrics) |  |  |






<a name="onos.pci.GetNeigborsRequest"></a>

### GetNeigborsRequest
must specify cell id: will only return a single cell&#39;s neigbors


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cell_id | [uint64](#uint64) |  |  |






<a name="onos.pci.GetNeigborsResponse"></a>

### GetNeigborsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| neigbor_ids | [uint64](#uint64) | repeated |  |






<a name="onos.pci.GetNumConflictsReponse"></a>

### GetNumConflictsReponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [uint64](#uint64) |  |  |






<a name="onos.pci.GetNumConflictsRequest"></a>

### GetNumConflictsRequest
if cell id is not specified, will return total number of conflicts


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cell_id | [uint64](#uint64) |  |  |






<a name="onos.pci.GetPciPoolRequest"></a>

### GetPciPoolRequest
if cell id is not specified, will return all


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cell_id | [uint64](#uint64) |  |  |






<a name="onos.pci.GetPciPoolResponse"></a>

### GetPciPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pools | [GetPciPoolResponse.PoolsEntry](#onos.pci.GetPciPoolResponse.PoolsEntry) | repeated |  |






<a name="onos.pci.GetPciPoolResponse.PoolsEntry"></a>

### GetPciPoolResponse.PoolsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [uint64](#uint64) |  |  |
| value | [PciRange](#onos.pci.PciRange) |  |  |






<a name="onos.pci.Metrics"></a>

### Metrics
metrics for a cell


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dlearfcn | [uint32](#uint32) |  |  |
| cell_type | [CellType](#onos.pci.CellType) |  |  |
| pci | [uint32](#uint32) |  |  |






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
| GetNumConflicts | [GetNumConflictsRequest](#onos.pci.GetNumConflictsRequest) | [GetNumConflictsReponse](#onos.pci.GetNumConflictsReponse) |  |
| GetNeighbors | [GetNeigborsRequest](#onos.pci.GetNeigborsRequest) | [GetNeigborsResponse](#onos.pci.GetNeigborsResponse) |  |
| GetMetric | [GetMetricRequest](#onos.pci.GetMetricRequest) | [GetMetricResponse](#onos.pci.GetMetricResponse) |  |
| GetPci | [GetPciPoolRequest](#onos.pci.GetPciPoolRequest) | [GetPciPoolResponse](#onos.pci.GetPciPoolResponse) |  |

 



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

