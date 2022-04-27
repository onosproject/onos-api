# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/kpimon/kpimon.proto](#onos_kpimon_kpimon-proto)
    - [GetRequest](#onos-kpimon-GetRequest)
    - [GetResponse](#onos-kpimon-GetResponse)
    - [GetResponse.MeasurementsEntry](#onos-kpimon-GetResponse-MeasurementsEntry)
    - [IntegerValue](#onos-kpimon-IntegerValue)
    - [MeasurementItem](#onos-kpimon-MeasurementItem)
    - [MeasurementItems](#onos-kpimon-MeasurementItems)
    - [MeasurementRecord](#onos-kpimon-MeasurementRecord)
    - [NoValue](#onos-kpimon-NoValue)
    - [RealValue](#onos-kpimon-RealValue)
  
    - [Kpimon](#onos-kpimon-Kpimon)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos_kpimon_kpimon-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/kpimon/kpimon.proto



<a name="onos-kpimon-GetRequest"></a>

### GetRequest







<a name="onos-kpimon-GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| measurements | [GetResponse.MeasurementsEntry](#onos-kpimon-GetResponse-MeasurementsEntry) | repeated |  |






<a name="onos-kpimon-GetResponse-MeasurementsEntry"></a>

### GetResponse.MeasurementsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [MeasurementItems](#onos-kpimon-MeasurementItems) |  |  |






<a name="onos-kpimon-IntegerValue"></a>

### IntegerValue



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [int64](#int64) |  |  |






<a name="onos-kpimon-MeasurementItem"></a>

### MeasurementItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| measurement_records | [MeasurementRecord](#onos-kpimon-MeasurementRecord) | repeated |  |






<a name="onos-kpimon-MeasurementItems"></a>

### MeasurementItems



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| measurement_items | [MeasurementItem](#onos-kpimon-MeasurementItem) | repeated |  |






<a name="onos-kpimon-MeasurementRecord"></a>

### MeasurementRecord



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| timestamp | [uint64](#uint64) |  |  |
| measurement_name | [string](#string) |  |  |
| measurement_value | [google.protobuf.Any](#google-protobuf-Any) |  |  |






<a name="onos-kpimon-NoValue"></a>

### NoValue



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [int32](#int32) |  |  |






<a name="onos-kpimon-RealValue"></a>

### RealValue



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [double](#double) |  |  |





 

 

 


<a name="onos-kpimon-Kpimon"></a>

### Kpimon


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListMeasurements | [GetRequest](#onos-kpimon-GetRequest) | [GetResponse](#onos-kpimon-GetResponse) |  |
| WatchMeasurements | [GetRequest](#onos-kpimon-GetRequest) | [GetResponse](#onos-kpimon-GetResponse) stream |  |

 



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

