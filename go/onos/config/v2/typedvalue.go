// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v2

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

const (
	isPositiveTypeOpt = int32(0)
	isNegativeTypeOpt = int32(1)
)

type Width int

const (
	WidthUnknown Width = 1 << (iota + 2)
	WidthEight
	WidthSixteen
	WidthThirtyTwo
	WidthSixtyFour
)

// TypedValueMap is an alias for a map of paths and values
type TypedValueMap map[string]*TypedValue

// ValueToString is the String calculated as a Native type
// There is already a String() in the protobuf generated code that can't be
// overridden
func (tv *TypedValue) ValueToString() string {
	switch tv.Type {
	case ValueType_EMPTY:
		return ""
	case ValueType_STRING:
		return (*TypedString)(tv).String()
	case ValueType_INT:
		return (*TypedInt)(tv).String()
	case ValueType_UINT:
		return (*TypedUint)(tv).String()
	case ValueType_BOOL:
		return (*TypedBool)(tv).String()
	case ValueType_DECIMAL:
		return (*TypedDecimal)(tv).String()
	case ValueType_FLOAT:
		return (*TypedFloat)(tv).String()
	case ValueType_BYTES:
		return (*TypedBytes)(tv).String()
	case ValueType_LEAFLIST_STRING:
		return (*TypedLeafListString)(tv).String()
	case ValueType_LEAFLIST_INT:
		return (*TypedLeafListInt)(tv).String()
	case ValueType_LEAFLIST_UINT:
		return (*TypedLeafListUint)(tv).String()
	case ValueType_LEAFLIST_BOOL:
		return (*TypedLeafListBool)(tv).String()
	case ValueType_LEAFLIST_DECIMAL:
		return (*TypedLeafListDecimal)(tv).String()
	case ValueType_LEAFLIST_FLOAT:
		return (*TypedLeafListFloat)(tv).String()
	case ValueType_LEAFLIST_BYTES:
		return (*TypedLeafListBytes)(tv).String()
	}

	return ""
}

// NewTypedValue creates a TypeValue from a byte[] and type - used in changes.go
// For Int and Uint both the width and sign must be given in type opts e.g. [32, 1]
func NewTypedValue(bytes []byte, valueType ValueType, typeOpts []uint8) (*TypedValue, error) {
	switch valueType {
	case ValueType_EMPTY:
		return NewTypedValueEmpty(), nil
	case ValueType_STRING:
		return NewTypedValueString(string(bytes)), nil
	case ValueType_INT:
		if len(typeOpts) != 2 {
			return nil, fmt.Errorf("number width AND sign must be given for INT as type opts. %v", typeOpts)
		}
		if typeOpts[0]/8 < uint8(len(bytes)) {
			return nil, fmt.Errorf("number width %d must match number of bytes %d for INT", typeOpts[0], len(bytes))
		}
		var bigInt big.Int
		bigInt.SetBytes(bytes)
		if len(typeOpts) == 2 && typeOpts[1] != uint8(isPositiveTypeOpt) {
			// Negative value
			bigInt.Neg(&bigInt)
		}
		return NewTypedValueInt(int(bigInt.Int64()), Width(typeOpts[0])), nil
	case ValueType_UINT:
		if len(typeOpts) != 1 {
			return nil, fmt.Errorf("number width must be given for UINT as type opts. %v", typeOpts)
		}
		if typeOpts[0]/8 < uint8(len(bytes)) {
			return nil, fmt.Errorf("number width %d must match number of bytes %d for INT", typeOpts[0], len(bytes))
		}
		var bigInt big.Int
		bigInt.SetBytes(bytes)

		return NewTypedValueUint(uint(bigInt.Uint64()), Width(typeOpts[0])), nil
	case ValueType_BOOL:
		if len(bytes) != 1 {
			return nil, fmt.Errorf("expecting 1 byte for BOOL. Got %d", len(bytes))
		}
		value := false
		if bytes[0] == 1 {
			value = true
		}
		return NewTypedValueBool(value), nil
	case ValueType_DECIMAL:
		if len(typeOpts) != 2 {
			return nil, fmt.Errorf("precision AND sign must be given for DECIMAL as type opts. %v", typeOpts)
		}
		precision := typeOpts[0]
		var bigInt big.Int
		bigInt.SetBytes(bytes)
		if len(typeOpts) == 2 && typeOpts[1] != uint8(isPositiveTypeOpt) {
			// Negative value
			bigInt.Neg(&bigInt)
		}
		return NewTypedValueDecimal(bigInt.Int64(), precision), nil
	case ValueType_FLOAT:
		if len(bytes) != 8 {
			return nil, fmt.Errorf("expecting 8 bytes for FLOAT. Got %d", len(bytes))
		}
		return NewTypedValueFloat(float64(math.Float64frombits(binary.LittleEndian.Uint64(bytes)))), nil
	case ValueType_BYTES:
		return NewTypedValueBytes(bytes), nil
	case ValueType_LEAFLIST_STRING:
		return caseValueTypeLeafListSTRING(bytes)
	case ValueType_LEAFLIST_INT:
		return caseValueTypeLeafListINT(bytes, typeOpts)
	case ValueType_LEAFLIST_UINT:
		return caseValueTypeLeafListUINT(bytes, typeOpts)
	case ValueType_LEAFLIST_BOOL:
		return caseValueTypeLeafListBOOL(bytes)
	case ValueType_LEAFLIST_DECIMAL:
		return caseValueTypeLeafListDECIMAL(bytes, typeOpts)
	case ValueType_LEAFLIST_FLOAT:
		return caseValueTypeLeafListFLOAT(bytes)
	case ValueType_LEAFLIST_BYTES:
		return caseValueTypeLeafListBYTES(bytes, typeOpts)
	}

	return nil, fmt.Errorf("unexpected type %d", valueType)
}

// caseValueTypeLeafListSTRING is moved out of NewTypedValue because of gocyclo
func caseValueTypeLeafListSTRING(bytes []byte) (*TypedValue, error) {
	stringList := make([]string, 0)
	buf := make([]byte, 0)
	for _, b := range bytes {
		if b != 0x1D {
			buf = append(buf, b)
		} else {
			stringList = append(stringList, string(buf))
			buf = make([]byte, 0)
		}
	}
	stringList = append(stringList, string(buf))
	return NewLeafListStringTv(stringList), nil
}

// caseValueTypeLeafListINT is moved out of NewTypedValue because of gocyclo
func caseValueTypeLeafListINT(bytes []byte, typeOpts []uint8) (*TypedValue, error) {
	count := (len(typeOpts) - 1) / 2
	if count < 1 {
		return nil, fmt.Errorf("unexpected #type opts. Expect 1 for width and then a pair per entry [len bytes, negative]")
	}
	intList := make([]int64, 0)
	width := Width(typeOpts[0])
	var byteCounter uint32 = 0
	for i := 0; i < count; i++ {
		v := bytes[byteCounter : byteCounter+uint32(typeOpts[1+i*2])]
		byteCounter += uint32(typeOpts[1+i*2])
		var bigInt big.Int
		bigInt.SetBytes(v)
		negative := typeOpts[1+i*2+1]
		if negative != uint8(isPositiveTypeOpt) {
			bigInt.Neg(&bigInt)
		}
		intList = append(intList, bigInt.Int64())
	}
	return NewLeafListIntTv(intList, width), nil
}

// caseValueTypeLeafListUINT is moved out of NewTypedValue because of gocyclo
func caseValueTypeLeafListUINT(bytes []byte, typeOpts []uint8) (*TypedValue, error) {
	count := (len(typeOpts) - 1)
	if count < 1 {
		return nil, fmt.Errorf("unexpected #type opts. Expect 1 for width and then a value per entry [len bytes]")
	}
	uintList := make([]uint64, 0)
	width := Width(typeOpts[0])
	var byteCounter uint32 = 0
	for i := 0; i < count; i++ {
		v := bytes[byteCounter : byteCounter+uint32(typeOpts[1+i])]
		byteCounter += uint32(typeOpts[1+i])
		var bigInt big.Int
		bigInt.SetBytes(v)
		uintList = append(uintList, bigInt.Uint64())
	}

	return NewLeafListUintTv(uintList, width), nil
}

// caseValueTypeLeafListBOOL is moved out of NewTypedValue because of gocyclo
func caseValueTypeLeafListBOOL(bytes []byte) (*TypedValue, error) {
	count := len(bytes)
	bools := make([]bool, count)
	for i, v := range bytes {
		if v == 1 {
			bools[i] = true
		}
	}
	return NewLeafListBoolTv(bools), nil
}

// caseValueTypeLeafListDECIMAL is moved out of NewTypedValue because of gocyclo
// First typeOpt is precision. Expect a pair of type opts per entry [len bytes, negative]
func caseValueTypeLeafListDECIMAL(bytes []byte, typeOpts []uint8) (*TypedValue, error) {
	count := (len(typeOpts) - 1) / 2
	if count < 1 {
		return nil, fmt.Errorf("unexpected #type opts. Expect 1 for precision and then a pair per entry [len bytes, negative]")
	}
	digitsList := make([]int64, 0)
	precision := typeOpts[0]
	var byteCounter uint32 = 0
	for i := 0; i < count; i++ {
		v := bytes[byteCounter : byteCounter+uint32(typeOpts[1+i*2])]
		byteCounter += uint32(typeOpts[1+i*2])
		var bigInt big.Int
		bigInt.SetBytes(v)
		negative := typeOpts[1+i*2+1]
		if negative != uint8(isPositiveTypeOpt) {
			bigInt.Neg(&bigInt)
		}
		digitsList = append(digitsList, bigInt.Int64())
	}
	return NewLeafListDecimalTv(digitsList, precision), nil
}

// caseValueTypeLeafListFLOAT is moved out of NewTypedValue because of gocyclo
func caseValueTypeLeafListFLOAT(bytes []byte) (*TypedValue, error) {
	count := len(bytes) / 8
	float32s := make([]float32, 0)

	for i := 0; i < count; i++ {
		v := bytes[i*8 : i*8+8]
		float32s = append(float32s, float32(math.Float64frombits(binary.LittleEndian.Uint64(v))))
	}
	return NewLeafListFloatTv(float32s), nil
}

// caseValueTypeLeafListBYTES is moved out of NewTypedValue because of gocyclo
func caseValueTypeLeafListBYTES(bytes []byte, typeOpts []uint8) (*TypedValue, error) {
	if len(typeOpts) < 1 {
		return nil, fmt.Errorf("expecting 1 typeopt for LeafListBytes. Got %d", len(typeOpts))
	}
	byteArrays := make([][]byte, 0)
	buf := make([]byte, 0)
	idx := 0
	startAt := 0
	for i, b := range bytes {
		valueLen := typeOpts[idx]
		if i-startAt == int(valueLen) {
			byteArrays = append(byteArrays, buf)
			buf = make([]byte, 0)
			idx = idx + 1
			if idx >= len(typeOpts) {
				return nil, fmt.Errorf("not enough typeOpts provided - found %d need at least %d", len(typeOpts), idx+1)
			}
			startAt = startAt + int(valueLen)
		}
		buf = append(buf, b)
	}
	byteArrays = append(byteArrays, buf)
	return NewLeafListBytesTv(byteArrays), nil
}

////////////////////////////////////////////////////////////////////////////////
// TypedEmpty
////////////////////////////////////////////////////////////////////////////////

// TypedEmpty for an empty value
type TypedEmpty TypedValue

// NewTypedValueEmpty decodes an empty object
func NewTypedValueEmpty() *TypedValue {
	return (*TypedValue)(newEmpty())
}

// newEmpty creates an instance of the Empty type
func newEmpty() *TypedEmpty {
	typedEmpty := TypedEmpty{
		Bytes: make([]byte, 0),
		Type:  ValueType_EMPTY,
	}
	return &typedEmpty
}

// ValueType gives the value type
func (tv *TypedEmpty) ValueType() ValueType {
	return tv.Type
}

func (tv *TypedEmpty) String() string {
	return ""
}

////////////////////////////////////////////////////////////////////////////////
// TypedString
////////////////////////////////////////////////////////////////////////////////

// TypedString for a string value
type TypedString TypedValue

// NewTypedValueString decodes string value in to an object
func NewTypedValueString(value string) *TypedValue {
	return (*TypedValue)(newString(value))
}

// newString decodes string value in to a String type
func newString(value string) *TypedString {
	typedString := TypedString{
		Bytes: []byte(value),
		Type:  ValueType_STRING,
	}
	return &typedString
}

// ValueType gives the value type
func (tv *TypedString) ValueType() ValueType {
	return tv.Type
}

func (tv *TypedString) String() string {
	return string(tv.Bytes)
}

////////////////////////////////////////////////////////////////////////////////
// TypedInt
////////////////////////////////////////////////////////////////////////////////

// TypedInt for an int value
type TypedInt TypedValue

// NewTypedValueInt decodes an int value in to an object
func NewTypedValueInt(value int, width Width) *TypedValue {
	return (*TypedValue)(newInt(big.NewInt(int64(value)), width))
}

func newInt(value *big.Int, width Width) *TypedInt {
	var isNegative int32 = isPositiveTypeOpt
	if value.Sign() < 0 {
		isNegative = isNegativeTypeOpt
	}
	typedInt := TypedInt{
		Bytes:    value.Bytes(),
		Type:     ValueType_INT,
		TypeOpts: []int32{int32(width), isNegative},
	}
	return &typedInt
}

// ValueType gives the value type
func (tv *TypedInt) ValueType() ValueType {
	return tv.Type
}

func (tv *TypedInt) String() string {
	return fmt.Sprintf("%d", tv.Int())
}

// Int extracts the integer value
func (tv *TypedInt) Int() int {
	var x big.Int
	x.SetBytes(tv.Bytes)
	if len(tv.TypeOpts) > 1 && tv.TypeOpts[1] == 1 {
		x.Neg(&x)
	}
	return int(x.Int64())
}

////////////////////////////////////////////////////////////////////////////////
// TypedUint
////////////////////////////////////////////////////////////////////////////////

// TypedUint for a uint value
type TypedUint TypedValue

// NewTypedValueUint decodes a uint value in to an object
func NewTypedValueUint(value uint, width Width) *TypedValue {
	var bigInt big.Int
	bigInt.SetUint64(uint64(value))
	return (*TypedValue)(newUint(&bigInt, width))
}

// newUint decodes a uint value in to a Uint type
func newUint(value *big.Int, width Width) *TypedUint {
	typedUint64 := TypedUint{
		Bytes:    value.Bytes(),
		Type:     ValueType_UINT,
		TypeOpts: []int32{int32(width)},
	}
	return &typedUint64
}

// ValueType gives the value type
func (tv *TypedUint) ValueType() ValueType {
	return tv.Type
}

func (tv *TypedUint) String() string {
	return fmt.Sprintf("%d", tv.Uint())
}

// Uint extracts the unsigned integer value
func (tv *TypedUint) Uint() uint {
	var bigInt big.Int
	bigInt.SetBytes(tv.Bytes)
	return uint(bigInt.Uint64())
}

////////////////////////////////////////////////////////////////////////////////
// TypedBool
////////////////////////////////////////////////////////////////////////////////

// TypedBool for an int value
type TypedBool TypedValue

// NewTypedValueBool decodes a bool value in to an object
func NewTypedValueBool(value bool) *TypedValue {
	return (*TypedValue)(newBool(value))
}

// newBool decodes a bool value in to an object
func newBool(value bool) *TypedBool {
	buf := make([]byte, 1)
	if value {
		buf[0] = 1
	}
	typedBool := TypedBool{
		Bytes: buf,
		Type:  ValueType_BOOL,
	}
	return &typedBool
}

// ValueType gives the value type
func (tv *TypedBool) ValueType() ValueType {
	return tv.Type
}

func (tv *TypedBool) String() string {
	if tv.Bytes[0] == 1 {
		return "true"
	}
	return "false"
}

// Bool extracts the unsigned bool value
func (tv *TypedBool) Bool() bool {
	return tv.Bytes[0] == 1
}

////////////////////////////////////////////////////////////////////////////////
// TypedDecimal
////////////////////////////////////////////////////////////////////////////////

// TypedDecimal for a decimal64 value
type TypedDecimal TypedValue

// NewTypedValueDecimal decodes a decimal value in to an object
func NewTypedValueDecimal(digits int64, precision uint8) *TypedValue {
	return (*TypedValue)(newDecimal(big.NewInt(digits), precision))
}

// newDecimal decodes a decimal value in to a Decimal type
func newDecimal(digits *big.Int, precision uint8) *TypedDecimal {
	var isNegative int32 = isPositiveTypeOpt
	if digits.Sign() < 0 {
		isNegative = isNegativeTypeOpt
	}
	typedDecimal64 := TypedDecimal{
		Bytes:    digits.Bytes(),
		Type:     ValueType_DECIMAL,
		TypeOpts: []int32{int32(precision), isNegative},
	}
	return &typedDecimal64
}

// ValueType gives the value type
func (tv *TypedDecimal) ValueType() ValueType {
	return tv.Type
}

func (tv *TypedDecimal) String() string {
	return strDecimal64(tv.Decimal64())
}

// Decimal64 extracts the unsigned decimal value
func (tv *TypedDecimal) Decimal64() (int64, uint8) {
	if len(tv.TypeOpts) > 0 {
		precision := tv.TypeOpts[0]
		var positiveMultiplier int64 = 1
		if len(tv.TypeOpts) > 1 && tv.TypeOpts[1] == 1 {
			positiveMultiplier = -1
		}
		var value big.Int
		value.SetBytes(tv.Bytes)
		return value.Int64() * positiveMultiplier, uint8(precision)
	}
	return 0, 0
}

// Float extracts the unsigned decimal value as a float
func (tv *TypedDecimal) Float() float64 {
	floatVal, _ := strconv.ParseFloat(strDecimal64(tv.Decimal64()), 64)
	return floatVal
}

////////////////////////////////////////////////////////////////////////////////
// TypedFloat
////////////////////////////////////////////////////////////////////////////////

// TypedFloat for a float value
type TypedFloat TypedValue

// NewTypedValueFloat decodes a decimal value in to an object
func NewTypedValueFloat(value float64) *TypedValue {
	return (*TypedValue)(newFloat(big.NewFloat(value)))
}

// newFloat decodes a decimal value in to a Bool type
func newFloat(value *big.Float) *TypedFloat {
	bytes, _ := value.GobEncode()
	typedFloat := TypedFloat{
		Bytes: bytes,
		Type:  ValueType_FLOAT,
	}
	return &typedFloat
}

// ValueType gives the value type
func (tv *TypedFloat) ValueType() ValueType {
	return tv.Type
}

func (tv *TypedFloat) String() string {
	return fmt.Sprintf("%f", tv.Float32())
}

// Float32 extracts the float value
func (tv *TypedFloat) Float32() float32 {
	if len(tv.Bytes) == 0 {
		return 0.0
	}
	var value big.Float
	_ = value.GobDecode(tv.Bytes)
	flt32, _ := value.Float32()
	return flt32
}

////////////////////////////////////////////////////////////////////////////////
// TypedBytes
////////////////////////////////////////////////////////////////////////////////

// TypedBytes for a float value
type TypedBytes TypedValue

// NewTypedValueBytes decodes an array of bytes in to an object
func NewTypedValueBytes(value []byte) *TypedValue {
	return (*TypedValue)(newBytes(value))
}

// newBytes decodes an array of bytes in to a Bytes type
func newBytes(value []byte) *TypedBytes {
	typedFloat := TypedBytes{
		Bytes:    value,
		Type:     ValueType_BYTES,
		TypeOpts: []int32{int32(len(value))},
	}
	return &typedFloat
}

// ValueType gives the value type
func (tv *TypedBytes) ValueType() ValueType {
	return tv.Type
}

func (tv *TypedBytes) String() string {
	return base64.StdEncoding.EncodeToString(tv.Bytes)
}

// ByteArray extracts the bytes value
func (tv *TypedBytes) ByteArray() []byte {
	return tv.Bytes
}

////////////////////////////////////////////////////////////////////////////////
// TypedLeafListString
////////////////////////////////////////////////////////////////////////////////

// TypedLeafListString for a string leaf list
type TypedLeafListString TypedValue

// NewLeafListStringTv decodes string values in to an object
func NewLeafListStringTv(values []string) *TypedValue {
	return (*TypedValue)(newLeafListString(values))
}

// newLeafListString decodes string values in to an Leaf list type
func newLeafListString(values []string) *TypedLeafListString {
	first := true
	bytes := make([]byte, 0)
	for _, v := range values {
		if first {
			first = false
		} else {
			bytes = append(bytes, 0x1D) // Group separator
		}
		bytes = append(bytes, []byte(v)...)
	}
	typedLeafListString := TypedLeafListString{
		Bytes: bytes,
		Type:  ValueType_LEAFLIST_STRING,
	}
	return &typedLeafListString
}

// ValueType gives the value type
func (tv *TypedLeafListString) ValueType() ValueType {
	return tv.Type
}

func (tv *TypedLeafListString) String() string {
	return strings.Join(tv.List(), ",")
}

// List extracts the leaf list values
func (tv *TypedLeafListString) List() []string {
	stringList := make([]string, 0)
	buf := make([]byte, 0)
	for _, b := range tv.Bytes {
		if b != 0x1D {
			buf = append(buf, b)
		} else {
			stringList = append(stringList, string(buf))
			buf = make([]byte, 0)
		}
	}
	stringList = append(stringList, string(buf))
	return stringList
}

////////////////////////////////////////////////////////////////////////////////
// TypedLeafListInt
////////////////////////////////////////////////////////////////////////////////

// TypedLeafListInt for an int leaf list
type TypedLeafListInt TypedValue

// NewLeafListIntTv decodes int values in to an object
func NewLeafListIntTv(values []int64, width Width) *TypedValue {
	valuesBi := make([]*big.Int, 0)
	for _, v := range values {
		valuesBi = append(valuesBi, big.NewInt(v))
	}
	return (*TypedValue)(newLeafListInt(valuesBi, width))
}

// newLeafListInt decodes int values in to a Leaf list type
func newLeafListInt(values []*big.Int, width Width) *TypedLeafListInt {
	bytes := make([]byte, 0)
	typeOpts := make([]int32, 0)
	typeOpts = append(typeOpts, int32(width))
	for _, v := range values {
		var isNegative int32 = isPositiveTypeOpt
		if v.Sign() < 0 {
			isNegative = isNegativeTypeOpt
		}
		typeOpts = append(typeOpts, int32(len(v.Bytes())))
		typeOpts = append(typeOpts, isNegative)
		bytes = append(bytes, v.Bytes()...)
	}
	typedLeafListInt := TypedLeafListInt{
		Bytes:    bytes,
		Type:     ValueType_LEAFLIST_INT,
		TypeOpts: typeOpts,
	}
	return &typedLeafListInt
}

// ValueType gives the value type
func (tv *TypedLeafListInt) ValueType() ValueType {
	return tv.Type
}

func (tv *TypedLeafListInt) String() string {
	intValues, width := tv.List()
	return fmt.Sprintf("%v %d", intValues, width)
}

// List extracts the leaf list values
func (tv *TypedLeafListInt) List() ([]int64, Width) {
	count := (len(tv.TypeOpts) - 1) / 2

	intList := make([]int64, 0)
	width := tv.TypeOpts[0]
	var byteCounter int32 = 0
	for i := 0; i < count; i++ {
		v := tv.Bytes[byteCounter : byteCounter+tv.TypeOpts[1+i*2]]
		byteCounter = byteCounter + tv.TypeOpts[1+i*2]
		var bigInt big.Int
		bigInt.SetBytes(v)
		negative := tv.TypeOpts[1+i*2+1]
		if negative != isPositiveTypeOpt {
			bigInt.Neg(&bigInt)
		}
		intList = append(intList, bigInt.Int64())
	}

	return intList, Width(width)
}

////////////////////////////////////////////////////////////////////////////////
// TypedLeafListUint
////////////////////////////////////////////////////////////////////////////////

// TypedLeafListUint for an uint leaf list
type TypedLeafListUint TypedValue

// NewLeafListUintTv decodes uint values in to a Leaf list
func NewLeafListUintTv(values []uint64, width Width) *TypedValue {
	valuesBi := make([]*big.Int, 0)
	for _, v := range values {
		var bigUint big.Int
		bigUint.SetUint64(v)
		valuesBi = append(valuesBi, &bigUint)
	}

	return (*TypedValue)(newLeafListUint(valuesBi, width))
}

// newLeafListUint decodes uint values in to a Leaf list type
func newLeafListUint(values []*big.Int, width Width) *TypedLeafListUint {
	bytes := make([]byte, 0)
	typeOpts := make([]int32, 0)
	typeOpts = append(typeOpts, int32(width))
	for _, v := range values {
		typeOpts = append(typeOpts, int32(len(v.Bytes())))
		bytes = append(bytes, v.Bytes()...)
	}
	typedLeafListUint := TypedLeafListUint{
		Bytes:    bytes,
		Type:     ValueType_LEAFLIST_UINT,
		TypeOpts: typeOpts,
	}
	return &typedLeafListUint
}

// ValueType gives the value type
func (tv *TypedLeafListUint) ValueType() ValueType {
	return tv.Type
}

func (tv *TypedLeafListUint) String() string {
	intValues, width := tv.List()
	return fmt.Sprintf("%v %d", intValues, width)
}

// List extracts the leaf list values
func (tv *TypedLeafListUint) List() ([]uint64, Width) {
	count := len(tv.TypeOpts) - 1

	uintList := make([]uint64, 0)
	width := tv.TypeOpts[0]
	var byteCounter int32 = 0
	for i := 0; i < count; i++ {
		v := tv.Bytes[byteCounter : byteCounter+tv.TypeOpts[1+i]]
		byteCounter = byteCounter + tv.TypeOpts[1+i]
		var bigInt big.Int
		bigInt.SetBytes(v)
		uintList = append(uintList, bigInt.Uint64())
	}

	return uintList, Width(width)
}

////////////////////////////////////////////////////////////////////////////////
// TypedLeafListBool
////////////////////////////////////////////////////////////////////////////////

// TypedLeafListBool for an bool leaf list
type TypedLeafListBool TypedValue

// NewLeafListBoolTv decodes bool values in to an object
func NewLeafListBoolTv(values []bool) *TypedValue {
	return (*TypedValue)(newLeafListBool(values))
}

// newLeafListBool decodes bool values in to a Leaf list type
func newLeafListBool(values []bool) *TypedLeafListBool {
	count := len(values)
	bytes := make([]byte, count)
	for i, b := range values {
		// just use one byte per bool - inefficient but not worth the hassle
		var intval uint8
		if b {
			intval = 1
		}
		bytes[i] = intval
	}
	typedLeafListBool := TypedLeafListBool{
		Bytes: bytes,
		Type:  ValueType_LEAFLIST_BOOL,
	}
	return &typedLeafListBool
}

// ValueType gives the value type
func (tv *TypedLeafListBool) ValueType() ValueType {
	return tv.Type
}

func (tv *TypedLeafListBool) String() string {
	return fmt.Sprintf("%v", tv.List())
}

// List extracts the leaf list values
func (tv *TypedLeafListBool) List() []bool {
	count := len(tv.Bytes)
	bools := make([]bool, count)
	for i, v := range tv.Bytes {
		if v == 1 {
			bools[i] = true
		}
	}

	return bools
}

////////////////////////////////////////////////////////////////////////////////
// TypedLeafListDecimal
////////////////////////////////////////////////////////////////////////////////

// TypedLeafListDecimal for a decimal leaf list
type TypedLeafListDecimal TypedValue

// NewLeafListDecimalTv decodes decimal values in to a Leaf list
func NewLeafListDecimalTv(digits []int64, precision uint8) *TypedValue {
	digitsBi := make([]*big.Int, 0)
	for _, d := range digits {
		digitsBi = append(digitsBi, big.NewInt(d))
	}
	return (*TypedValue)(newLeafListDecimal(digitsBi, precision))
}

// newLeafListDecimal decodes decimal values in to a Leaf list type
func newLeafListDecimal(digits []*big.Int, precision uint8) *TypedLeafListDecimal {
	bytes := make([]byte, 0)
	typeOpts := make([]int32, 0)
	typeOpts = append(typeOpts, int32(precision))
	for _, d := range digits {
		typeOpts = append(typeOpts, int32(len(d.Bytes())))
		var isNegative int32 = isPositiveTypeOpt
		if d.Sign() < 0 {
			isNegative = isNegativeTypeOpt
		}
		typeOpts = append(typeOpts, isNegative)
		bytes = append(bytes, d.Bytes()...)
	}
	typedLeafListDecimal := TypedLeafListDecimal{
		Bytes:    bytes,
		Type:     ValueType_LEAFLIST_DECIMAL,
		TypeOpts: typeOpts,
	}
	return &typedLeafListDecimal
}

// ValueType gives the value type
func (tv *TypedLeafListDecimal) ValueType() ValueType {
	return tv.Type
}

func (tv *TypedLeafListDecimal) String() string {
	digits, precision := tv.List()
	return fmt.Sprintf("%v %d", digits, precision)
}

// List extracts the leaf list values
func (tv *TypedLeafListDecimal) List() ([]int64, uint8) {
	count := (len(tv.TypeOpts) - 1) / 2

	digitsList := make([]int64, 0)
	precision := uint8(tv.TypeOpts[0])
	var byteCounter int32 = 0
	for i := 0; i < count; i++ {
		v := tv.Bytes[byteCounter : byteCounter+tv.TypeOpts[1+i*2]]
		byteCounter = byteCounter + tv.TypeOpts[1+i*2]
		var bigInt big.Int
		bigInt.SetBytes(v)
		negative := tv.TypeOpts[1+i*2+1]
		if negative != isPositiveTypeOpt {
			bigInt.Neg(&bigInt)
		}
		digitsList = append(digitsList, bigInt.Int64())
	}

	return digitsList, precision
}

// ListFloat extracts the leaf list values as floats
func (tv *TypedLeafListDecimal) ListFloat() []float64 {
	digits, precision := tv.List()
	floatList := make([]float64, len(digits))
	for i, d := range digits {
		floatList[i] = float64(d) / math.Pow(10, float64(precision))
	}

	return floatList
}

////////////////////////////////////////////////////////////////////////////////
// TypedLeafListFloat
////////////////////////////////////////////////////////////////////////////////

// TypedLeafListFloat for a decimal leaf list
type TypedLeafListFloat TypedValue

// NewLeafListFloatTv decodes float values in to a Leaf list
func NewLeafListFloatTv(values []float32) *TypedValue {
	return (*TypedValue)(newLeafListFloat(values))
}

// newLeafListFloat decodes float values in to a Leaf list type
func newLeafListFloat(values []float32) *TypedLeafListFloat {
	bytes := make([]byte, 0)
	for _, f := range values {
		buf := make([]byte, 8)
		binary.LittleEndian.PutUint64(buf, math.Float64bits(float64(f)))
		bytes = append(bytes, buf...)
	}
	typedLeafListFloat := TypedLeafListFloat{
		Bytes: bytes,
		Type:  ValueType_LEAFLIST_FLOAT,
	}
	return &typedLeafListFloat
}

// ValueType gives the value type
func (tv *TypedLeafListFloat) ValueType() ValueType {
	return tv.Type
}

func (tv *TypedLeafListFloat) String() string {
	listStr := make([]string, 0)
	for _, f := range tv.List() {
		listStr = append(listStr, fmt.Sprintf("%f", f))
	}

	return strings.Join(listStr, ",")
}

// List extracts the leaf list values
func (tv *TypedLeafListFloat) List() []float32 {
	count := len(tv.Bytes) / 8
	float32s := make([]float32, 0)

	for i := 0; i < count; i++ {
		v := tv.Bytes[i*8 : i*8+8]
		float32s = append(float32s, float32(math.Float64frombits(binary.LittleEndian.Uint64(v))))
	}

	return float32s
}

////////////////////////////////////////////////////////////////////////////////
// TypedLeafListBytes
////////////////////////////////////////////////////////////////////////////////

// TypedLeafListBytes for an bool leaf list
type TypedLeafListBytes TypedValue

// NewLeafListBytesTv decodes byte values in to a Leaf list
func NewLeafListBytesTv(values [][]byte) *TypedValue {
	return (*TypedValue)(newLeafListBytes(values))
}

// newLeafListBytes decodes byte values in to a Leaf list type
func newLeafListBytes(values [][]byte) *TypedLeafListBytes {
	bytes := make([]byte, 0)
	typeopts := make([]int32, 0)
	for _, v := range values {
		bytes = append(bytes, v...)
		typeopts = append(typeopts, int32(len(v)))
	}
	typedLeafListBytes := TypedLeafListBytes{
		Bytes:    bytes,
		Type:     ValueType_LEAFLIST_BYTES, // Contains the lengths of each byte array in list
		TypeOpts: typeopts,
	}
	return &typedLeafListBytes
}

// ValueType gives the value type
func (tv *TypedLeafListBytes) ValueType() ValueType {
	return tv.Type
}

func (tv *TypedLeafListBytes) String() string {
	return fmt.Sprintf("%v", tv.List())
}

// List extracts the leaf list values
func (tv *TypedLeafListBytes) List() [][]byte {
	bytes := make([][]byte, 0)
	buf := make([]byte, 0)
	idx := 0
	startAt := 0
	for i, b := range tv.Bytes {
		valueLen := tv.TypeOpts[idx]
		if i-startAt == int(valueLen) {
			bytes = append(bytes, buf)
			buf = make([]byte, 0)
			idx = idx + 1
			startAt = startAt + int(valueLen)
		}
		buf = append(buf, b)
	}
	bytes = append(bytes, buf)
	return bytes
}

func strDecimal64(digits int64, precision uint8) string {
	var i, frac int64
	if precision > 0 {
		div := int64(10)
		it := precision - 1
		for it > 0 {
			div *= 10
			it--
		}
		i = digits / div
		frac = digits % div
	} else {
		i = digits
	}
	if frac < 0 {
		frac = -frac
	}
	if precision == 0 {
		return fmt.Sprintf("%d", i)
	}
	fmtString := fmt.Sprintf("%s0.%d%s", "%d.%", precision, "d")
	return fmt.Sprintf(fmtString, i, frac)
}
