// Copyright 2019-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package device

import (
	"encoding/binary"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"math/big"
	"testing"
)

func Test_NewChangedValue(t *testing.T) {
	path := "/a/b/c"
	badPath := "a@b@c"
	value := NewTypedValueUint(64, 32)
	const isRemove = false
	changeValueBad, errorBad := NewChangeValue(badPath, value, isRemove)
	assert.Error(t, errorBad)
	assert.Contains(t, errorBad.Error(), badPath)
	assert.Nil(t, changeValueBad)
	changeValue, err := NewChangeValue(path, value, isRemove)
	assert.NoError(t, err)
	assert.NotNil(t, changeValue)
	assert.Equal(t, changeValue.Path, path)
	assert.Equal(t, "64", changeValue.Value.ValueToString())
}

func Test_NewValueTypeString(t *testing.T) {
	const value = "xyzzy"
	typedValue, err := NewTypedValue([]byte(value), ValueType_STRING, make([]uint8, 0))
	assert.NoError(t, err)
	assert.Equal(t, typedValue.ValueToString(), value)
}

func floatToBinary(value float64, size int) []byte {
	bs := make([]byte, size)
	binary.LittleEndian.PutUint64(bs, math.Float64bits(value))
	return bs
}

func uintToBinary(value uint64) []byte {
	var bigInt big.Int
	bigInt.SetUint64(value)
	return bigInt.Bytes()
}

func intToBinary(value int64) []byte {
	var bigInt big.Int
	bigInt.SetInt64(value)
	return bigInt.Bytes()
}

func boolToBinary(value bool, size int) []byte {
	bs := make([]byte, size)
	intValue := 0
	if value {
		intValue = 1
	} else {
		intValue = 0
	}
	bs[0] = byte(intValue)
	return bs
}

func TestValueTypes(t *testing.T) {
	testCases := []struct {
		description   string
		value         []byte
		valueType     ValueType
		expectedValue string
		expectedError string
		typeOpts      []uint8
	}{
		{
			description:   "NewValueEmpty",
			value:         nil,
			valueType:     ValueType_EMPTY,
			expectedValue: "",
			expectedError: "",
		},
		{
			description:   "NewValueString",
			value:         []byte("abacab"),
			valueType:     ValueType_STRING,
			expectedValue: "abacab",
			expectedError: "",
		},
		{
			description:   "NewValueBytes",
			value:         []byte("abacab"),
			valueType:     ValueType_BYTES,
			expectedValue: "YWJhY2Fi",
			expectedError: "",
		},
		{
			description:   "NewValueTypeFloatSuccess",
			value:         floatToBinary(3.0, 8),
			valueType:     ValueType_FLOAT,
			expectedValue: fmt.Sprintf("%f", 3.0),
			expectedError: "",
		},
		{
			description:   "NewValueTypeFloatFailure",
			value:         floatToBinary(3.0, 12),
			valueType:     ValueType_FLOAT,
			expectedValue: fmt.Sprintf("%f", 3.0),
			expectedError: "expecting 8 bytes for FLOAT. Got 12",
		},
		{
			description:   "NewValueTypeUintSuccess",
			value:         uintToBinary(345678),
			valueType:     ValueType_UINT,
			typeOpts:      []uint8{32},
			expectedValue: "345678",
			expectedError: "",
		},
		{
			description:   "NewValueTypeUintFailure",
			value:         uintToBinary(345678),
			valueType:     ValueType_UINT,
			expectedError: "number width must be given for UINT as type opts. []",
		},
		{
			description:   "NewValueTypeIntSuccess",
			value:         intToBinary(345678),
			valueType:     ValueType_INT,
			typeOpts:      []uint8{32, 0},
			expectedValue: "345678",
			expectedError: "",
		},
		{
			description:   "NewValueTypeIntFailure",
			value:         intToBinary(345678),
			valueType:     ValueType_INT,
			expectedValue: "345678",
			expectedError: "number width AND sign must be given for INT as type opts. []",
		},
		{
			description:   "NewValueTypeBoolSuccess",
			value:         boolToBinary(true, 1),
			valueType:     ValueType_BOOL,
			expectedValue: "true",
			expectedError: "",
		},
		{
			description:   "NewValueTypeIntFailure",
			value:         boolToBinary(true, 2),
			valueType:     ValueType_BOOL,
			expectedValue: "",
			expectedError: "expecting 1 byte for BOOL. Got 2",
		},
		{
			description:   "NewValueTypeDecimalSuccess",
			value:         intToBinary(1234),
			valueType:     ValueType_DECIMAL,
			typeOpts:      []uint8{0, 0},
			expectedValue: "1234",
			expectedError: "",
		},
		{
			description:   "NewValueTypeDecimalNegSuccess",
			value:         intToBinary(-1234),
			valueType:     ValueType_DECIMAL,
			typeOpts:      []uint8{2, 1},
			expectedValue: "-12.34",
			expectedError: "",
		},
		{
			description:   "NewValueTypeDecimalFailure",
			value:         intToBinary(1234),
			valueType:     ValueType_DECIMAL,
			typeOpts:      []uint8{3},
			expectedValue: "",
			expectedError: "precision AND sign must be given for DECIMAL as type opts. [3]",
		},
		{
			description:   "NewValueTypeLeafListString",
			value:         []byte{'a', 0x1D, 'b', 0x1D, 'c'},
			valueType:     ValueType_LEAFLIST_STRING,
			expectedValue: "a,b,c",
			expectedError: "",
		},
		{
			description: "NewValueTypeLeafListInt",
			value: []byte{
				11,
				22,
				33,
			},
			valueType:     ValueType_LEAFLIST_INT,
			typeOpts:      []uint8{8, 1, 0, 1, 0, 1, 0},
			expectedValue: "[11 22 33] 8",
			expectedError: "",
		},
		{
			description: "NewValueTypeLeafListUint",
			value: []byte{
				3,
				5,
				7,
			},
			valueType:     ValueType_LEAFLIST_UINT,
			typeOpts:      []uint8{8, 1, 1, 1},
			expectedValue: "[3 5 7] 8",
			expectedError: "",
		},
		{
			description: "NewValueTypeLeafListBool",
			value: []byte{
				1, 0, 0, 1,
			},
			valueType:     ValueType_LEAFLIST_BOOL,
			expectedValue: "[true false false true]",
			expectedError: "",
		},
		{
			description:   "NewValueTypeLeafListDecimal",
			value:         append(intToBinary(1234), intToBinary(-4321)...),
			valueType:     ValueType_LEAFLIST_DECIMAL,
			typeOpts:      []uint8{2, 2, 0, 2, 1},
			expectedValue: "[1234 -4321] 2",
			expectedError: "",
		},
		{
			description:   "NewValueTypeLeafListBytes",
			value:         []byte("12345678"),
			valueType:     ValueType_LEAFLIST_BYTES,
			typeOpts:      []uint8{4, 4},
			expectedValue: "[[49 50 51 52] [53 54 55 56]]",
			expectedError: "",
		},
		{
			description:   "NewValueTypeLeafListFloat",
			value:         append(floatToBinary(12.34, 8), floatToBinary(-1.234, 8)...),
			valueType:     ValueType_LEAFLIST_FLOAT,
			typeOpts:      []uint8{0},
			expectedValue: "12.340000,-1.234000",
			expectedError: "",
		},
	}

	for _, testCase := range testCases {
		typedValue, err := NewTypedValue(testCase.value, testCase.valueType, testCase.typeOpts)
		if testCase.expectedError != "" {
			assert.Error(t, err, testCase.expectedError, testCase.description)
			assert.Nil(t, typedValue, testCase.description)
		} else {
			assert.NoError(t, err, testCase.description)
			s := typedValue.ValueToString()
			assert.Equal(t, s, testCase.expectedValue, testCase.description)
		}
	}
}

func TestTypedEmpty(t *testing.T) {
	empty := newEmpty()
	assert.Equal(t, empty.ValueType(), ValueType_EMPTY)
	assert.Equal(t, empty.String(), "")
}

func TestTypedInt(t *testing.T) {
	intValue := newInt(big.NewInt(112233), 32)
	assert.Equal(t, intValue.ValueType(), ValueType_INT)
	assert.Equal(t, intValue.String(), "112233")
	assert.Equal(t, intValue.Int(), 112233)
}

func TestTypedUint(t *testing.T) {
	var bigInt big.Int
	bigInt.SetUint64(112233)
	intValue := newUint(&bigInt, 32)
	assert.Equal(t, intValue.ValueType(), ValueType_UINT)
	assert.Equal(t, intValue.String(), "112233")
	assert.Equal(t, uint64(intValue.Uint()), uint64(112233))
}

func TestTypedString(t *testing.T) {
	intValue := newString("xyzzy")
	assert.Equal(t, intValue.ValueType(), ValueType_STRING)
	assert.Equal(t, intValue.String(), "xyzzy")
}

func TestTypedBool(t *testing.T) {
	boolValue := newBool(false)
	assert.Equal(t, boolValue.ValueType(), ValueType_BOOL)
	assert.Equal(t, boolValue.String(), "false")
	assert.Equal(t, boolValue.Bool(), false)
}

func TestTypedDecimal(t *testing.T) {
	decimal64Value := newDecimal(big.NewInt(1234), 2)
	assert.Equal(t, decimal64Value.ValueType(), ValueType_DECIMAL)
	assert.Equal(t, decimal64Value.String(), "12.34")
	digits, precision := decimal64Value.Decimal64()
	assert.Equal(t, int(digits), 1234)
	assert.Equal(t, precision, uint8(2))
}

func TestTypedFloat(t *testing.T) {
	float64Value := newFloat(big.NewFloat(2222.0))
	assert.Equal(t, float64Value.ValueType(), ValueType_FLOAT)
	assert.Equal(t, float64Value.String(), "2222.000000")
	assert.Equal(t, float64Value.Float32(), float32(2222.0))
}

func TestTypedByte(t *testing.T) {
	bytes := []byte("bytes")
	bytesValue := newBytes(bytes)
	assert.Equal(t, bytesValue.ValueType(), ValueType_BYTES)
	assert.Equal(t, bytesValue.String(), "Ynl0ZXM=")
	assert.Equal(t, len(bytesValue.ByteArray()), len(bytes))
}

func TestTypedLeafListString(t *testing.T) {
	values := []string{"a", "b"}
	listValue := newLeafListString(values)
	assert.Equal(t, listValue.ValueType(), ValueType_LEAFLIST_STRING)
}

func TestTypedLeafListUint(t *testing.T) {
	values := []*big.Int{big.NewInt(1), big.NewInt(2)}
	listValue := newLeafListUint(values, WidthThirtyTwo)
	assert.Equal(t, listValue.ValueType(), ValueType_LEAFLIST_UINT)
}

func TestTypedLeafListInt(t *testing.T) {
	values := []*big.Int{big.NewInt(1), big.NewInt(2)}
	listValue := newLeafListInt(values, WidthThirtyTwo)
	assert.Equal(t, listValue.ValueType(), ValueType_LEAFLIST_INT)
}

func TestTypedLeafListBool(t *testing.T) {
	values := []bool{true, false}
	listValue := newLeafListBool(values)
	assert.Equal(t, listValue.ValueType(), ValueType_LEAFLIST_BOOL)
}

func TestTypedLeafListFloat(t *testing.T) {
	values := []float32{111.0, 112.0}
	listValue := newLeafListFloat(values)
	assert.Equal(t, listValue.ValueType(), ValueType_LEAFLIST_FLOAT)
}

func TestTypedLeafListDecimal(t *testing.T) {
	digits := []*big.Int{big.NewInt(22), big.NewInt(33)}

	listValue := newLeafListDecimal(digits, 2)
	assert.Equal(t, listValue.ValueType(), ValueType_LEAFLIST_DECIMAL)
	floats := listValue.ListFloat()
	assert.Equal(t, len(floats), len(digits))
	assert.Equal(t, 0.22, floats[0])
	assert.Equal(t, 0.33, floats[1])
}

func TestTypedLeafListBytes(t *testing.T) {
	values := make([][]byte, 2)
	values[0] = []byte("abc")
	values[1] = []byte("xyz")

	listValue := newLeafListBytes(values)
	assert.Equal(t, listValue.ValueType(), ValueType_LEAFLIST_BYTES)
}

func TestLeafListBytesCrash(t *testing.T) {
	bytes := []byte("12345678")
	typeOpts := []uint8{4}
	value, err := NewTypedValue(bytes, ValueType_LEAFLIST_BYTES, typeOpts)
	assert.Nil(t, value)
	assert.Error(t, err)
}

func TestIsPathValid(t *testing.T) {
	assert.Error(t, IsPathValid("/a/b/c[name=test,address=this.is.a.test]/def"))
	assert.Error(t, IsPathValid("/a/b//c"))

	assert.NoError(t, IsPathValid("/a1/b2/c3[name=test][address=this.is.a.test]/def"))
	assert.NoError(t, IsPathValid("/a1/b2/c3[name_field=test-1_0][address=ff00::ab00]/def"))
}