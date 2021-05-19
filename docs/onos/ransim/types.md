# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [onos/ransim/types/types.proto](#onos/ransim/types/types.proto)
    - [A3HandoverParams](#onos.ransim.types.A3HandoverParams)
    - [Cell](#onos.ransim.types.Cell)
    - [Cell.CrntiMapEntry](#onos.ransim.types.Cell.CrntiMapEntry)
    - [MapLayout](#onos.ransim.types.MapLayout)
    - [Node](#onos.ransim.types.Node)
    - [Point](#onos.ransim.types.Point)
    - [Route](#onos.ransim.types.Route)
    - [Sector](#onos.ransim.types.Sector)
    - [Ue](#onos.ransim.types.Ue)
    - [UeMetrics](#onos.ransim.types.UeMetrics)
  
- [Scalar Value Types](#scalar-value-types)



<a name="onos/ransim/types/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## onos/ransim/types/types.proto



<a name="onos.ransim.types.A3HandoverParams"></a>

### A3HandoverParams



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| A3Offset | [int32](#int32) |  |  |
| A3TimeToTrigger | [int32](#int32) |  |  |
| A3Hysteresis | [int32](#int32) |  |  |
| A3CellOffset | [int32](#int32) |  |  |
| A3FrequencyOffset | [int32](#int32) |  |  |






<a name="onos.ransim.types.Cell"></a>

### Cell



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ecgi | [uint64](#uint64) |  |  |
| location | [Point](#onos.ransim.types.Point) |  |  |
| sector | [Sector](#onos.ransim.types.Sector) |  |  |
| color | [string](#string) |  |  |
| max_ues | [uint32](#uint32) |  |  |
| neighbors | [uint64](#uint64) | repeated |  |
| tx_power_db | [double](#double) |  | The cell transmit power in decibels |
| a3_handover_params | [A3HandoverParams](#onos.ransim.types.A3HandoverParams) |  |  |
| crnti_map | [Cell.CrntiMapEntry](#onos.ransim.types.Cell.CrntiMapEntry) | repeated | crntis maps a ue&#39;s name to its crnti |
| crnti_index | [uint32](#uint32) |  |  |
| port | [uint32](#uint32) |  |  |






<a name="onos.ransim.types.Cell.CrntiMapEntry"></a>

### Cell.CrntiMapEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [uint32](#uint32) |  |  |
| value | [uint64](#uint64) |  |  |






<a name="onos.ransim.types.MapLayout"></a>

### MapLayout



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| center | [Point](#onos.ransim.types.Point) |  | Map center latitude and longitude |
| zoom | [float](#float) |  | The starting Zoom level |
| fade | [bool](#bool) |  | Show map as faded on start |
| show_routes | [bool](#bool) |  | Show routes on start |
| show_power | [bool](#bool) |  | Show power as circle on start |
| locations_scale | [float](#float) |  | Ratio of random locations diameter to tower grid width |
| min_ues | [uint32](#uint32) |  | FIXME: These are deprecated; remove Max number of UEs for complete simulation |
| max_ues | [uint32](#uint32) |  | Max number of UEs for complete simulation |
| current_routes | [uint32](#uint32) |  | the current number of routes |






<a name="onos.ransim.types.Node"></a>

### Node



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enbid | [uint32](#uint32) |  |  |
| controllers | [string](#string) | repeated |  |
| service_models | [string](#string) | repeated |  |
| cell_ecgis | [uint64](#uint64) | repeated |  |
| status | [string](#string) |  |  |






<a name="onos.ransim.types.Point"></a>

### Point



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| lat | [double](#double) |  |  |
| lng | [double](#double) |  |  |






<a name="onos.ransim.types.Route"></a>

### Route



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [uint64](#uint64) |  |  |
| waypoints | [Point](#onos.ransim.types.Point) | repeated |  |
| color | [string](#string) |  |  |
| speed_avg | [uint32](#uint32) |  |  |
| speed_stdev | [uint32](#uint32) |  |  |
| reverse | [bool](#bool) |  |  |
| next_point | [uint32](#uint32) |  |  |






<a name="onos.ransim.types.Sector"></a>

### Sector



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| azimuth | [int32](#int32) |  |  |
| arc | [int32](#int32) |  |  |
| centroid | [Point](#onos.ransim.types.Point) |  |  |






<a name="onos.ransim.types.Ue"></a>

### Ue



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| imsi | [uint64](#uint64) |  |  |
| type | [string](#string) |  |  |
| position | [Point](#onos.ransim.types.Point) |  |  |
| rotation | [uint32](#uint32) |  |  |
| serving_tower | [uint64](#uint64) |  |  |
| serving_tower_strength | [double](#double) |  |  |
| tower1 | [uint64](#uint64) |  |  |
| tower1_strength | [double](#double) |  |  |
| tower2 | [uint64](#uint64) |  |  |
| tower2_strength | [double](#double) |  |  |
| tower3 | [uint64](#uint64) |  |  |
| tower3_strength | [double](#double) |  |  |
| crnti | [uint32](#uint32) |  |  |
| admitted | [bool](#bool) |  |  |
| metrics | [UeMetrics](#onos.ransim.types.UeMetrics) |  |  |






<a name="onos.ransim.types.UeMetrics"></a>

### UeMetrics



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ho_latency | [int64](#int64) |  | Latency (in nanoseconds) of the most recent hand-over |
| ho_report_timestamp | [int64](#int64) |  | Handover report timestamp (in nanoseconds since epoch) |
| is_first | [bool](#bool) |  | flag to indicate the first measurement |





 

 

 

 



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

