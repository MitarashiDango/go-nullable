package nullable

import (
	"math"
	"time"
)

// String represents a nullable string.
type String struct {
	NullableBase[string]
}

// NewString returns a non-null String containing value.
func NewString(value string) String {
	return String{
		NewNullableBase(value),
	}
}

// NewNullString returns a null String.
func NewNullString() String {
	return String{}
}

// Equal reports whether nv and value are both null or contain the same string.
func (nv String) Equal(value String) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.RawValue() == value.RawValue()
}

// Int represents a nullable int.
type Int struct {
	NullableBase[int]
}

// NewInt returns a non-null Int containing value.
func NewInt(value int) Int {
	return Int{
		NewNullableBase(value),
	}
}

// NewNullInt returns a null Int.
func NewNullInt() Int {
	return Int{}
}

// Equal reports whether nv and value are both null or contain the same int.
func (nv Int) Equal(value Int) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.RawValue() == value.RawValue()
}

// UInt represents a nullable uint.
type UInt struct {
	NullableBase[uint]
}

// NewUInt returns a non-null UInt containing value.
func NewUInt(value uint) UInt {
	return UInt{
		NewNullableBase(value),
	}
}

// NewNullUInt returns a null UInt.
func NewNullUInt() UInt {
	return UInt{}
}

// Equal reports whether nv and value are both null or contain the same uint.
func (nv UInt) Equal(value UInt) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.RawValue() == value.RawValue()
}

// Int8 represents a nullable int8.
type Int8 struct {
	NullableBase[int8]
}

// NewInt8 returns a non-null Int8 containing value.
func NewInt8(value int8) Int8 {
	return Int8{
		NewNullableBase(value),
	}
}

// NewNullInt8 returns a null Int8.
func NewNullInt8() Int8 {
	return Int8{}
}

// Equal reports whether nv and value are both null or contain the same int8.
func (nv Int8) Equal(value Int8) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.RawValue() == value.RawValue()
}

// UInt8 represents a nullable uint8.
type UInt8 struct {
	NullableBase[uint8]
}

// NewUInt8 returns a non-null UInt8 containing value.
func NewUInt8(value uint8) UInt8 {
	return UInt8{
		NewNullableBase(value),
	}
}

// NewNullUInt8 returns a null UInt8.
func NewNullUInt8() UInt8 {
	return UInt8{}
}

// Equal reports whether nv and value are both null or contain the same uint8.
func (nv UInt8) Equal(value UInt8) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.RawValue() == value.RawValue()
}

// Int16 represents a nullable int16.
type Int16 struct {
	NullableBase[int16]
}

// NewInt16 returns a non-null Int16 containing value.
func NewInt16(value int16) Int16 {
	return Int16{
		NewNullableBase(value),
	}
}

// NewNullInt16 returns a null Int16.
func NewNullInt16() Int16 {
	return Int16{}
}

// Equal reports whether nv and value are both null or contain the same int16.
func (nv Int16) Equal(value Int16) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.RawValue() == value.RawValue()
}

// UInt16 represents a nullable uint16.
type UInt16 struct {
	NullableBase[uint16]
}

// NewUInt16 returns a non-null UInt16 containing value.
func NewUInt16(value uint16) UInt16 {
	return UInt16{
		NewNullableBase(value),
	}
}

// NewNullUInt16 returns a null UInt16.
func NewNullUInt16() UInt16 {
	return UInt16{}
}

// Equal reports whether nv and value are both null or contain the same uint16.
func (nv UInt16) Equal(value UInt16) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.RawValue() == value.RawValue()
}

// Int32 represents a nullable int32.
type Int32 struct {
	NullableBase[int32]
}

// NewInt32 returns a non-null Int32 containing value.
func NewInt32(value int32) Int32 {
	return Int32{
		NewNullableBase(value),
	}
}

// NewNullInt32 returns a null Int32.
func NewNullInt32() Int32 {
	return Int32{}
}

// Equal reports whether nv and value are both null or contain the same int32.
func (nv Int32) Equal(value Int32) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.RawValue() == value.RawValue()
}

// UInt32 represents a nullable uint32.
type UInt32 struct {
	NullableBase[uint32]
}

// NewUInt32 returns a non-null UInt32 containing value.
func NewUInt32(value uint32) UInt32 {
	return UInt32{
		NewNullableBase(value),
	}
}

// NewNullUInt32 returns a null UInt32.
func NewNullUInt32() UInt32 {
	return UInt32{}
}

// Equal reports whether nv and value are both null or contain the same uint32.
func (nv UInt32) Equal(value UInt32) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.RawValue() == value.RawValue()
}

// Int64 represents a nullable int64.
type Int64 struct {
	NullableBase[int64]
}

// NewInt64 returns a non-null Int64 containing value.
func NewInt64(value int64) Int64 {
	return Int64{
		NewNullableBase(value),
	}
}

// NewNullInt64 returns a null Int64.
func NewNullInt64() Int64 {
	return Int64{}
}

// Equal reports whether nv and value are both null or contain the same int64.
func (nv Int64) Equal(value Int64) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.RawValue() == value.RawValue()
}

// UInt64 represents a nullable uint64.
type UInt64 struct {
	NullableBase[uint64]
}

// NewUInt64 returns a non-null UInt64 containing value.
func NewUInt64(value uint64) UInt64 {
	return UInt64{
		NewNullableBase(value),
	}
}

// NewNullUInt64 returns a null UInt64.
func NewNullUInt64() UInt64 {
	return UInt64{}
}

// Equal reports whether nv and value are both null or contain the same uint64.
func (nv UInt64) Equal(value UInt64) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.RawValue() == value.RawValue()
}

// Byte is an alias for UInt8.
type Byte = UInt8

// Rune is an alias for Int32.
type Rune = Int32

// Float32 represents a nullable float32.
type Float32 struct {
	NullableBase[float32]
}

// NewFloat32 returns a non-null Float32 containing value.
func NewFloat32(value float32) Float32 {
	return Float32{
		NewNullableBase(value),
	}
}

// NewNullFloat32 returns a null Float32.
func NewNullFloat32() Float32 {
	return Float32{}
}

// Equal reports whether nv and value are both null or contain equal float32 values.
func (nv Float32) Equal(value Float32) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.RawValue() == value.RawValue()
}

// EqualBits reports whether nv and value are both null or have the same float32 bit pattern.
func (nv Float32) EqualBits(value Float32) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || math.Float32bits(nv.RawValue()) == math.Float32bits(value.RawValue())
}

// EqualEpsilon reports whether nv and value are both null or differ by at most epsilon.
// It panics if epsilon is negative or NaN.
func (nv Float32) EqualEpsilon(value Float32, epsilon float32) bool {
	if epsilon < 0 || math.IsNaN(float64(epsilon)) {
		panic("nullable: invalid epsilon")
	}

	if nv.IsNull() != value.IsNull() {
		return false
	}

	if nv.IsNull() {
		return true
	}

	diff := nv.RawValue() - value.RawValue()
	if diff < 0 {
		diff = -diff
	}
	return diff <= epsilon
}

// Float64 represents a nullable float64.
type Float64 struct {
	NullableBase[float64]
}

// NewFloat64 returns a non-null Float64 containing value.
func NewFloat64(value float64) Float64 {
	return Float64{
		NewNullableBase(value),
	}
}

// NewNullFloat64 returns a null Float64.
func NewNullFloat64() Float64 {
	return Float64{}
}

// Equal reports whether nv and value are both null or contain equal float64 values.
func (nv Float64) Equal(value Float64) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.RawValue() == value.RawValue()
}

// EqualBits reports whether nv and value are both null or have the same float64 bit pattern.
func (nv Float64) EqualBits(value Float64) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || math.Float64bits(nv.RawValue()) == math.Float64bits(value.RawValue())
}

// EqualEpsilon reports whether nv and value are both null or differ by at most epsilon.
// It panics if epsilon is negative or NaN.
func (nv Float64) EqualEpsilon(value Float64, epsilon float64) bool {
	if epsilon < 0 || math.IsNaN(epsilon) {
		panic("nullable: invalid epsilon")
	}

	if nv.IsNull() != value.IsNull() {
		return false
	}

	if nv.IsNull() {
		return true
	}

	return math.Abs(nv.RawValue()-value.RawValue()) <= epsilon
}

// Bool represents a nullable bool.
type Bool struct {
	NullableBase[bool]
}

// NewBool returns a non-null Bool containing value.
func NewBool(value bool) Bool {
	return Bool{
		NewNullableBase(value),
	}
}

// NewNullBool returns a null Bool.
func NewNullBool() Bool {
	return Bool{}
}

// Equal reports whether nv and value are both null or contain the same bool.
func (nv Bool) Equal(value Bool) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.RawValue() == value.RawValue()
}

// Time represents a nullable time.Time.
type Time struct {
	NullableBase[time.Time]
}

// NewTime returns a non-null Time containing value.
func NewTime(value time.Time) Time {
	return Time{
		NewNullableBase(value),
	}
}

// NewNullTime returns a null Time.
func NewNullTime() Time {
	return Time{}
}

// Equal reports whether nv and value are both null or contain equal time values.
func (nv Time) Equal(value Time) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.RawValue().Equal(value.RawValue())
}
