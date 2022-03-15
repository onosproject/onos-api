// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v2

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"gotest.tools/assert"
)

const testString = "This is a test"
const (
	testNegativeInt64 = -9223372036854775808
	testNegativeInt32 = -2147483648
	testZeroInt       = 0
	testPositiveInt32 = 2147483647
	testPositiveInt64 = 9223372036854775807
)
const (
	testZeroUint   = 0
	testElevenUint = 11
	testMaxUint    = uint64(18446744073709551615)
)
const (
	testPrecision3 = uint8(3)
	testPrecision6 = uint8(6)
	testDigitsZero = uint8(0)
	testPrecision0 = uint8(0)
)
const (
	testFloatNeg = -3.4E+38
	testFloatPos = 3.4E+38
)
const (
	testStringBytes    = "onos rocks!"
	testStringBytesB64 = "b25vcyByb2NrcyE="
)

var testLeafListString = []string{"abc", "def", "ghi", "with,comma"}

var testLeafListInt = []int64{testNegativeInt64, testZeroInt, testPositiveInt64}

var testLeafListUint = []uint64{testZeroUint, testElevenUint, testMaxUint}

var testLeafListBool = []bool{true, false, false, true}

var testLeafListDecimal = []int64{testNegativeInt32, testZeroInt, testPositiveInt32}

var testLeafListFloat = []float32{testFloatNeg, testZeroInt, testFloatPos}

var testByteArr0 = []byte("abc")
var testByteArr1 = []byte("defg")
var testByteArr2 = []byte("ghijk")
var testLeafListBytes = [][]byte{testByteArr0, testByteArr1, testByteArr2}

func Test_TypedValueEmpty(t *testing.T) {
	tv := NewTypedValueEmpty()

	assert.Equal(t, len(tv.Bytes), 0)
	testConversion(t, tv)
	assert.Equal(t, tv.ValueToString(), "")

	tv2, err := NewTypedValue([]byte{0x0, 0x0, 0x0}, ValueType_EMPTY, []uint8{})
	assert.NilError(t, err)
	assert.Equal(t, tv2.String(), "")
	assert.Equal(t, tv2.ValueToString(), "")
}

func Test_TypedValueString(t *testing.T) {
	tv := NewTypedValueString(testString)

	assert.Equal(t, len(tv.Bytes), 14)
	testConversion(t, tv)
	assert.Equal(t, tv.ValueToString(), testString)
}

func Test_TypedValueInt(t *testing.T) {
	tv := newInt(big.NewInt(testNegativeInt64), 64)
	assert.Equal(t, tv.String(), fmt.Sprintf("%d", testNegativeInt64))
	assert.Equal(t, len(tv.Bytes), 8)
	assert.DeepEqual(t, tv.Bytes, []uint8{0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})

	testConversion(t, (*TypedValue)(tv))

	tv = newInt(big.NewInt(12345678), 32)
	assert.Equal(t, tv.String(), fmt.Sprintf("%d", 12345678))
	assert.DeepEqual(t, tv.Bytes, []uint8{0xbc, 0x61, 0x4e})

	tv = newInt(big.NewInt(testPositiveInt64), 64)
	assert.Equal(t, tv.String(), fmt.Sprintf("%d", testPositiveInt64))
	assert.DeepEqual(t, tv.Bytes, []uint8{0x7f, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})

	tv = newInt(big.NewInt(testZeroInt), 32)
	assert.Equal(t, tv.String(), fmt.Sprintf("%d", testZeroInt))
	assert.Equal(t, (*TypedValue)(tv).ValueToString(), "0")
	assert.DeepEqual(t, tv.Bytes, []uint8{})
}

func Test_TypedValueUint(t *testing.T) {
	tv := newUint(big.NewInt(testZeroUint), 8)
	assert.Equal(t, tv.String(), fmt.Sprintf("%d", testZeroUint))
	assert.Equal(t, len(tv.Bytes), 0)

	tv = newUint(big.NewInt(12345678), 32)
	assert.Equal(t, tv.String(), fmt.Sprintf("%d", 12345678))
	assert.Equal(t, len(tv.Bytes), 3)
	assert.DeepEqual(t, tv.Bytes, []uint8{0xbc, 0x61, 0x4e})

	var bigInt big.Int
	bigInt.SetUint64(testMaxUint)
	tv = newUint(&bigInt, 64)
	assert.Equal(t, tv.String(), fmt.Sprintf("%d", testMaxUint))
	assert.DeepEqual(t, tv.Bytes, []uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})

	testConversion(t, (*TypedValue)(tv))
	assert.Equal(t, (*TypedValue)(tv).ValueToString(), "18446744073709551615")

}

func Test_TypedValueBool(t *testing.T) {
	tv := newBool(true)

	assert.Equal(t, len(tv.Bytes), 1)
	assert.Equal(t, tv.String(), "true")
	assert.DeepEqual(t, tv.Bytes, []uint8{0x01})

	testConversion(t, (*TypedValue)(tv))

	tv = newBool(false)
	assert.Equal(t, tv.String(), "false")
	assert.Equal(t, (*TypedValue)(tv).ValueToString(), "false")
}

func Test_TypedValueDecimal64(t *testing.T) {
	tv := newDecimal(big.NewInt(testNegativeInt64), testPrecision3)
	assert.Equal(t, len(tv.Bytes), 8)
	assert.Equal(t, tv.String(), "-9223372036854775.808")
	assert.Equal(t, tv.TypeOpts[0], int32(testPrecision3))
	assert.Equal(t, tv.TypeOpts[1], isNegativeTypeOpt)
	assert.DeepEqual(t, tv.Bytes, []byte{0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})

	testConversion(t, (*TypedValue)(tv))

	tv = newDecimal(big.NewInt(testPositiveInt64), testPrecision6)
	assert.Equal(t, tv.String(), "9223372036854.775807")
	assert.Equal(t, tv.TypeOpts[0], int32(testPrecision6))
	assert.Equal(t, tv.TypeOpts[1], isPositiveTypeOpt)
	assert.DeepEqual(t, tv.Bytes, []byte{0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})

	tv = newDecimal(big.NewInt(int64(testDigitsZero)), testPrecision0)
	assert.Equal(t, tv.String(), "0")
	assert.Equal(t, tv.TypeOpts[0], int32(testPrecision0))
	assert.Equal(t, tv.TypeOpts[1], isPositiveTypeOpt)
	assert.Equal(t, (*TypedValue)(tv).ValueToString(), "0")
	assert.DeepEqual(t, tv.Bytes, []byte{})
}

func Test_TypedValueFloat(t *testing.T) {
	tv := newFloat(big.NewFloat(testFloatNeg))
	assert.Equal(t, len(tv.Bytes), 18)
	assert.Equal(t, tv.String(), "-339999995214436424907732413799364296704.000000")
	assert.Equal(t, len(tv.TypeOpts), 0)
	assert.DeepEqual(t, tv.Bytes, []uint8{0x01, 0x0b, 0x00, 0x00, 0x00, 0x35, 0x00, 0x00, 0x00, 0x80, 0xff, 0xc9, 0x9e, 0x3c, 0x66, 0xfd, 0x68, 0x00})

	testConversion(t, (*TypedValue)(tv))

	tv = newFloat(big.NewFloat(testFloatPos))
	assert.Equal(t, tv.String(), "339999995214436424907732413799364296704.000000")
	assert.Equal(t, (*TypedValue)(tv).ValueToString(), "339999995214436424907732413799364296704.000000")
	assert.Equal(t, len(tv.TypeOpts), 0)
	assert.DeepEqual(t, tv.Bytes, []uint8{0x01, 0x0a, 0x00, 0x00, 0x00, 0x35, 0x00, 0x00, 0x00, 0x80, 0xff, 0xc9, 0x9e, 0x3c, 0x66, 0xfd, 0x68, 0x00})
}

func Test_TypedValueBytes(t *testing.T) {
	tv := newBytes([]byte(testStringBytes))
	assert.Equal(t, len(tv.Bytes), 11)
	assert.Equal(t, tv.String(), testStringBytesB64)

	testConversion(t, (*TypedValue)(tv))
	assert.Equal(t, (*TypedValue)(tv).ValueToString(), testStringBytesB64)
}

func Test_LeafListString(t *testing.T) {
	tv := newLeafListString(testLeafListString)
	assert.Equal(t, len(tv.Bytes), 22)
	assert.Equal(t, tv.String(), "abc,def,ghi,with,comma")
	testConversion(t, (*TypedValue)(tv))

	testArray := []string{"one"}
	tv = newLeafListString(testArray)
	assert.Equal(t, len(tv.Bytes), 3)
	assert.Equal(t, tv.String(), "one")
	assert.Equal(t, (*TypedValue)(tv).ValueToString(), "one")
}

func Test_LeafListInt64(t *testing.T) {
	tv := NewLeafListIntTv(testLeafListInt, WidthSixtyFour)
	assert.Equal(t, len(tv.Bytes), 16)
	assert.Equal(t, ValueType_LEAFLIST_INT, tv.Type)
	assert.DeepEqual(t, []int32{64, 8, 1, 0, 0, 8, 0}, tv.TypeOpts)
	testConversion(t, (*TypedValue)(tv))
	assert.Equal(t, (*TypedValue)(tv).ValueToString(), "[-9223372036854775808 0 9223372036854775807] 64")
}

func Test_LeafListUint64(t *testing.T) {
	tv := NewLeafListUintTv(testLeafListUint, WidthSixtyFour)

	assert.Equal(t, 9, len(tv.Bytes))
	assert.Equal(t, ValueType_LEAFLIST_UINT, tv.Type)
	assert.DeepEqual(t, []int32{64, 0, 1, 8}, tv.TypeOpts)

	testConversion(t, (*TypedValue)(tv))
	assert.Equal(t, (*TypedValue)(tv).ValueToString(), "[0 11 18446744073709551615] 64")
}

func Test_LeafListBool(t *testing.T) {
	tv := newLeafListBool(testLeafListBool)

	assert.Equal(t, len(tv.Bytes), 4)
	assert.Equal(t, tv.String(), "[true false false true]")

	testConversion(t, (*TypedValue)(tv))
	assert.Equal(t, (*TypedValue)(tv).ValueToString(), "[true false false true]")
}

func Test_LeafListDecimal64(t *testing.T) {
	testLeafListDecimalBi := []*big.Int{
		big.NewInt(testLeafListDecimal[0]),
		big.NewInt(testLeafListDecimal[1]),
		big.NewInt(testLeafListDecimal[2]),
	}
	tv := newLeafListDecimal(testLeafListDecimalBi, testPrecision6)

	assert.Equal(t, len(tv.Bytes), 8)
	assert.Equal(t, tv.String(), "[-2147483648 0 2147483647] 6")

	testConversion(t, (*TypedValue)(tv))
	assert.Equal(t, (*TypedValue)(tv).ValueToString(), "[-2147483648 0 2147483647] 6")
}

func Test_LeafListFloat32(t *testing.T) {
	tv := newLeafListFloat(testLeafListFloat)

	assert.Equal(t, len(tv.Bytes), 24)
	assert.Equal(t, tv.String(), "-339999995214436424907732413799364296704.000000,0.000000,339999995214436424907732413799364296704.000000")

	testConversion(t, (*TypedValue)(tv))
	assert.Equal(t, (*TypedValue)(tv).ValueToString(), "-339999995214436424907732413799364296704.000000,0.000000,339999995214436424907732413799364296704.000000")
}

func Test_LeafListBytes(t *testing.T) {
	tv := newLeafListBytes(testLeafListBytes)

	testConversion(t, (*TypedValue)(tv))

	assert.Equal(t, len(tv.Bytes), 12)
	assert.Equal(t, tv.String(), "[[97 98 99] [100 101 102 103] [103 104 105 106 107]]")
	assert.Equal(t, tv.TypeOpts[0], int32(3))
	assert.Equal(t, tv.TypeOpts[1], int32(4))
	assert.Equal(t, tv.TypeOpts[2], int32(5))
	assert.Equal(t, (*TypedValue)(tv).ValueToString(), "[[97 98 99] [100 101 102 103] [103 104 105 106 107]]")
}

func Test_JsonSerializationString(t *testing.T) {
	tv := NewTypedValueString(testString)

	jsonStr, err := json.Marshal(tv)
	assert.NilError(t, err)

	assert.Equal(t, string(jsonStr), `{"Bytes":"VGhpcyBpcyBhIHRlc3Q=","Type":1}`)

	unmarshalledTv := TypedValue{}
	err = json.Unmarshal(jsonStr, &unmarshalledTv)
	assert.NilError(t, err)

	assert.Equal(t, unmarshalledTv.Type, ValueType_STRING)
	assert.Equal(t, len(unmarshalledTv.TypeOpts), 0)
	assert.DeepEqual(t, unmarshalledTv.Bytes, []byte{0x54, 0x68, 0x69, 0x73, 0x20, 0x69, 0x73, 0x20, 0x61, 0x20, 0x74, 0x65, 0x73, 0x74})

	assert.Equal(t, (*TypedString)(&unmarshalledTv).String(), testString)
}

func Test_JsonSerializationDecimal(t *testing.T) {
	tv := NewTypedValueDecimal(1232, 6)

	jsonStr, err := json.Marshal(tv)
	assert.NilError(t, err)

	assert.Equal(t, string(jsonStr), `{"Bytes":"BNA=","Type":5,"TypeOpts":[6,0]}`)

	unmarshalledTv := TypedValue{}
	err = json.Unmarshal(jsonStr, &unmarshalledTv)
	assert.NilError(t, err)

	assert.Equal(t, unmarshalledTv.Type, ValueType_DECIMAL)
	assert.Equal(t, len(unmarshalledTv.TypeOpts), 2)
	assert.Equal(t, unmarshalledTv.TypeOpts[0], int32(6))
	assert.Equal(t, unmarshalledTv.TypeOpts[1], int32(isPositiveTypeOpt))
	assert.DeepEqual(t, unmarshalledTv.Bytes, []byte{0x04, 0xd0})

	decFloat := (*TypedDecimal)(&unmarshalledTv).Float()
	assert.Equal(t, decFloat, 0.001232)
}

// From RFC-7951: A value of the "int64", "uint64", or "decimal64" type is represented as a JSON string
func Test_JsonSerializationInt64(t *testing.T) {
	tv := NewTypedValueInt(testNegativeInt64, 64)

	jsonStr, err := json.Marshal(tv)
	assert.NilError(t, err)

	assert.Equal(t, string(jsonStr), `{"Bytes":"gAAAAAAAAAA=","Type":2,"TypeOpts":[64,1]}`)

	unmarshalledTv := TypedValue{}
	err = json.Unmarshal(jsonStr, &unmarshalledTv)
	assert.NilError(t, err)

	assert.Equal(t, ValueType_INT, unmarshalledTv.Type)
	assert.Equal(t, len(unmarshalledTv.TypeOpts), 2)
	assert.DeepEqual(t, unmarshalledTv.Bytes, []byte{0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})

	strVal := (*TypedInt)(&unmarshalledTv).String()
	assert.Equal(t, fmt.Sprintf("%d", testNegativeInt64), strVal)
	assert.Equal(t, unmarshalledTv.ValueToString(), "-9223372036854775808")
}

func Test_JsonSerializationInt32(t *testing.T) {
	tv := NewTypedValueInt(testNegativeInt32, 32)

	jsonStr, err := json.Marshal(tv)
	assert.NilError(t, err)

	assert.Equal(t, string(jsonStr), `{"Bytes":"gAAAAA==","Type":2,"TypeOpts":[32,1]}`)

	unmarshalledTv := TypedValue{}
	err = json.Unmarshal(jsonStr, &unmarshalledTv)
	assert.NilError(t, err)

	assert.Equal(t, ValueType_INT, unmarshalledTv.Type)
	assert.Equal(t, len(unmarshalledTv.TypeOpts), 2)
	assert.Equal(t, unmarshalledTv.TypeOpts[0], int32(32))
	assert.Equal(t, unmarshalledTv.TypeOpts[1], isNegativeTypeOpt)
	assert.DeepEqual(t, unmarshalledTv.Bytes, []byte{0x80, 0x00, 0x00, 0x00})

	intVal := (*TypedInt)(&unmarshalledTv).Int()
	assert.Equal(t, testNegativeInt32, intVal)
	assert.Equal(t, unmarshalledTv.ValueToString(), "-2147483648")
}

func Test_JsonSerializationUint8(t *testing.T) {
	tv := NewTypedValueUint(16, 8)

	jsonStr, err := json.Marshal(tv)
	assert.NilError(t, err)

	assert.Equal(t, string(jsonStr), `{"Bytes":"EA==","Type":3,"TypeOpts":[8]}`)

	unmarshalledTv := TypedValue{}
	err = json.Unmarshal(jsonStr, &unmarshalledTv)
	assert.NilError(t, err)

	assert.Equal(t, unmarshalledTv.Type, ValueType_UINT)
	assert.Equal(t, len(unmarshalledTv.TypeOpts), 1)
	assert.DeepEqual(t, unmarshalledTv.Bytes, []byte{0x10})

	uintVal := (*TypedUint)(&unmarshalledTv).Uint()
	assert.Equal(t, uintVal, uint(16))
}

// From RFC-7951: A value of the "int64", "uint64", or "decimal64" type is represented as a JSON string
func Test_JsonSerializationUint64(t *testing.T) {
	tv := NewTypedValueUint(testPositiveInt64, 64)

	jsonStr, err := json.Marshal(tv)
	assert.NilError(t, err)

	assert.Equal(t, string(jsonStr), `{"Bytes":"f/////////8=","Type":3,"TypeOpts":[64]}`)

	unmarshalledTv := TypedValue{}
	err = json.Unmarshal(jsonStr, &unmarshalledTv)
	assert.NilError(t, err)

	assert.Equal(t, ValueType_UINT, unmarshalledTv.Type)
	assert.Equal(t, len(unmarshalledTv.TypeOpts), 1)
	assert.DeepEqual(t, unmarshalledTv.Bytes, []byte{0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})

	strVal := (*TypedUint)(&unmarshalledTv).String()
	assert.Equal(t, fmt.Sprintf("%d", testPositiveInt64), strVal)
	assert.Equal(t, unmarshalledTv.ValueToString(), "9223372036854775807")
}

func Test_JsonSerializationBool(t *testing.T) {
	tv := NewTypedValueBool(true)

	jsonStr, err := json.Marshal(tv)
	assert.NilError(t, err)

	assert.Equal(t, string(jsonStr), `{"Bytes":"AQ==","Type":4}`)

	unmarshalledTv := TypedValue{}
	err = json.Unmarshal(jsonStr, &unmarshalledTv)
	assert.NilError(t, err)

	assert.Equal(t, unmarshalledTv.Type, ValueType_BOOL)
	assert.Equal(t, len(unmarshalledTv.TypeOpts), 0)
	assert.DeepEqual(t, unmarshalledTv.Bytes, []byte{0x1})

	b := (*TypedBool)(&unmarshalledTv).Bool()
	assert.Equal(t, b, true)
}

func Test_JsonSerializationLeafListInt(t *testing.T) {
	tv := NewLeafListIntTv(testLeafListInt, WidthSixtyFour)
	jsonStr, err := json.Marshal(tv)
	assert.NilError(t, err)

	assert.Equal(t, string(jsonStr), `{"Bytes":"gAAAAAAAAAB//////////w==","Type":9,"TypeOpts":[64,8,1,0,0,8,0]}`)

	unmarshalledTv := TypedValue{}
	err = json.Unmarshal(jsonStr, &unmarshalledTv)
	assert.NilError(t, err)

	assert.Equal(t, unmarshalledTv.Type, ValueType_LEAFLIST_INT)
	assert.Equal(t, len(unmarshalledTv.TypeOpts), 7)
	assert.Equal(t, len(unmarshalledTv.Bytes), 16)
	assert.DeepEqual(t, unmarshalledTv.Bytes, []byte{0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})

	assert.Equal(t, (*TypedLeafListInt)(&unmarshalledTv).String(), "[-9223372036854775808 0 9223372036854775807] 64")
}

func Test_JsonSerializationLeafListBytes(t *testing.T) {
	tv := NewLeafListBytesTv(testLeafListBytes)

	jsonStr, err := json.Marshal(tv)
	assert.NilError(t, err)

	assert.Equal(t, string(jsonStr), `{"Bytes":"YWJjZGVmZ2doaWpr","Type":14,"TypeOpts":[3,4,5]}`)

	unmarshalledTv := TypedValue{}
	err = json.Unmarshal(jsonStr, &unmarshalledTv)
	assert.NilError(t, err)

	assert.Equal(t, unmarshalledTv.Type, ValueType_LEAFLIST_BYTES)
	assert.Equal(t, len(unmarshalledTv.TypeOpts), 3)
	assert.Equal(t, unmarshalledTv.TypeOpts[0], int32(3))
	assert.Equal(t, unmarshalledTv.TypeOpts[1], int32(4))
	assert.Equal(t, unmarshalledTv.TypeOpts[2], int32(5))
	assert.DeepEqual(t, unmarshalledTv.Bytes, []byte{0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x67, 0x68, 0x69, 0x6a, 0x6b})

	assert.Equal(t, (*TypedLeafListBytes)(&unmarshalledTv).String(), "[[97 98 99] [100 101 102 103] [103 104 105 106 107]]")
}

func Test_CreateFromBytesInt32(t *testing.T) {
	var buf big.Int
	buf.SetInt64(int64(testPositiveInt32))

	tv, err := NewTypedValue(buf.Bytes(), ValueType_INT, []uint8{32, 0})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(tv.TypeOpts))
	assert.Equal(t, int32(32), tv.TypeOpts[0])
	assert.Equal(t, tv.ValueToString(), "2147483647")
}

func Test_CreateFromBytesNegInt32(t *testing.T) {
	var buf big.Int
	buf.SetInt64(int64(testNegativeInt32))

	tv, err := NewTypedValue(buf.Bytes(), ValueType_INT, []uint8{32, 1})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(tv.TypeOpts))
	assert.Equal(t, int32(32), tv.TypeOpts[0])
	assert.Equal(t, tv.ValueToString(), "-2147483648")
}

func Test_CreateFromBytesInt64(t *testing.T) {
	var buf big.Int
	buf.SetInt64(int64(testPositiveInt64))

	tv, err := NewTypedValue(buf.Bytes(), ValueType_INT, []uint8{64, 0})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(tv.TypeOpts))
	assert.Equal(t, tv.ValueToString(), "9223372036854775807")
}

func Test_CreateFromBytesBool(t *testing.T) {
	buf := []byte{0x01}

	tv, err := NewTypedValue(buf, ValueType_BOOL, nil)
	assert.NilError(t, err)
	assert.Equal(t, tv.ValueToString(), "true")
}

func testConversion(t *testing.T, tv *TypedValue) {

	switch tv.Type {
	case ValueType_EMPTY:
		assert.Equal(t, (*TypedEmpty)(tv).String(), "")
	case ValueType_STRING:
		assert.Equal(t, (*TypedString)(tv).String(), testString)
	case ValueType_INT:
		assert.Equal(t, (*TypedInt)(tv).Int(), testNegativeInt64)
	case ValueType_UINT:
		assert.Equal(t, (*TypedUint)(tv).Uint(), uint(testMaxUint))
	case ValueType_BOOL:
		assert.Equal(t, (*TypedBool)(tv).Bool(), true)
	case ValueType_DECIMAL:
		digits, precision := (*TypedDecimal)(tv).Decimal64()
		assert.Equal(t, digits, int64(testNegativeInt64))
		assert.Equal(t, precision, uint8(testPrecision3))
	case ValueType_FLOAT:
		assert.Equal(t, (*TypedFloat)(tv).Float32(), float32(testFloatNeg))
	case ValueType_BYTES:
		assert.Equal(t, base64.StdEncoding.EncodeToString((*TypedBytes)(tv).ByteArray()), testStringBytesB64)
		assert.Equal(t, len(tv.TypeOpts), 1)
		assert.Equal(t, tv.TypeOpts[0], int32(11))
	case ValueType_LEAFLIST_STRING:
		assert.DeepEqual(t, (*TypedLeafListString)(tv).List(), testLeafListString)
	case ValueType_LEAFLIST_INT:
		list, width := (*TypedLeafListInt)(tv).List()
		assert.Equal(t, WidthSixtyFour, width)
		assert.Equal(t, 3, len(list))
		assert.DeepEqual(t, testLeafListInt, list)
	case ValueType_LEAFLIST_UINT:
		list, width := (*TypedLeafListUint)(tv).List()
		assert.Equal(t, WidthSixtyFour, width)
		assert.Equal(t, 3, len(list))
		assert.DeepEqual(t, testLeafListUint, list)
	case ValueType_LEAFLIST_BOOL:
		assert.DeepEqual(t, (*TypedLeafListBool)(tv).List(), testLeafListBool)
	case ValueType_LEAFLIST_DECIMAL:
		digits, precision := (*TypedLeafListDecimal)(tv).List()
		for i := range testLeafListDecimal {
			assert.Equal(t, testLeafListDecimal[i], digits[i])
		}
		assert.Equal(t, precision, uint8(testPrecision6))
	case ValueType_LEAFLIST_FLOAT:
		assert.DeepEqual(t, (*TypedLeafListFloat)(tv).List(), testLeafListFloat)
	case ValueType_LEAFLIST_BYTES:
		assert.DeepEqual(t, (*TypedLeafListBytes)(tv).List(), testLeafListBytes)
		assert.Equal(t, len(tv.TypeOpts), 3)
		assert.Equal(t, tv.TypeOpts[0], int32(3))
		assert.Equal(t, tv.TypeOpts[1], int32(4))
		assert.Equal(t, tv.TypeOpts[2], int32(5))

	default:
		t.Log("Unexpected type", tv.Type)
		t.Fail()
	}

}
