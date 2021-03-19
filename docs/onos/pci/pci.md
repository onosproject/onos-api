# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/pci/pci.proto](#onos/pci/pci.proto)
    - [GetRequest](#onos.pci.GetRequest)
    - [GetResponse](#onos.pci.GetResponse)
    - [Object](#onos.pci.Object)
    - [Object.AttributesEntry](#onos.pci.Object.AttributesEntry)
  
    - [Pci](#onos.pci.Pci)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos/pci/pci.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/pci/pci.proto



<a name="onos.pci.GetRequest"></a>

### GetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="onos.pci.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object | [Object](#onos.pci.Object) |  |  |






<a name="onos.pci.Object"></a>

### Object



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| revision | [uint64](#uint64) |  |  |
| attributes | [Object.AttributesEntry](#onos.pci.Object.AttributesEntry) | repeated |  |






<a name="onos.pci.Object.AttributesEntry"></a>

### Object.AttributesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |





 

 

 


<a name="onos.pci.Pci"></a>

### Pci


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetNumConflicts | [GetRequest](#onos.pci.GetRequest) | [GetResponse](#onos.pci.GetResponse) |  |
| GetNumConflictsAll | [GetRequest](#onos.pci.GetRequest) | [GetResponse](#onos.pci.GetResponse) |  |
| GetNeighbors | [GetRequest](#onos.pci.GetRequest) | [GetResponse](#onos.pci.GetResponse) |  |
| GetNeighborsAll | [GetRequest](#onos.pci.GetRequest) | [GetResponse](#onos.pci.GetResponse) |  |
| GetMetric | [GetRequest](#onos.pci.GetRequest) | [GetResponse](#onos.pci.GetResponse) |  |
| GetMetricAll | [GetRequest](#onos.pci.GetRequest) | [GetResponse](#onos.pci.GetResponse) |  |
| GetPci | [GetRequest](#onos.pci.GetRequest) | [GetResponse](#onos.pci.GetResponse) |  |
| GetPciAll | [GetRequest](#onos.pci.GetRequest) | [GetResponse](#onos.pci.GetResponse) |  |

 



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

