package nullable

import "time"

type String struct {
	NullableBase[string]
}

func NewString(value string) String {
	return String{
		NewNullableBase(value),
	}
}

func NewNullString() String {
	return String{}
}

func (nv String) Equal(value String) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.Value() == value.Value()
}

type Int struct {
	NullableBase[int]
}

func NewInt(value int) Int {
	return Int{
		NewNullableBase(value),
	}
}

func NewNullInt() Int {
	return Int{}
}

func (nv Int) Equal(value Int) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.Value() == value.Value()
}

type UInt struct {
	NullableBase[uint]
}

func NewUInt(value uint) UInt {
	return UInt{
		NewNullableBase(value),
	}
}

func NewNullUInt() UInt {
	return UInt{}
}

func (nv UInt) Equal(value UInt) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.Value() == value.Value()
}

type Int8 struct {
	NullableBase[int8]
}

func NewInt8(value int8) Int8 {
	return Int8{
		NewNullableBase(value),
	}
}

func NewNullInt8() Int8 {
	return Int8{}
}

func (nv Int8) Equal(value Int8) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.Value() == value.Value()
}

type UInt8 struct {
	NullableBase[uint8]
}

func NewUInt8(value uint8) UInt8 {
	return UInt8{
		NewNullableBase(value),
	}
}

func NewNullUInt8() UInt8 {
	return UInt8{}
}

func (nv UInt8) Equal(value UInt8) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.Value() == value.Value()
}

type Int16 struct {
	NullableBase[int16]
}

func NewInt16(value int16) Int16 {
	return Int16{
		NewNullableBase(value),
	}
}

func NewNullInt16() Int16 {
	return Int16{}
}

func (nv Int16) Equal(value Int16) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.Value() == value.Value()
}

type UInt16 struct {
	NullableBase[uint16]
}

func NewUInt16(value uint16) UInt16 {
	return UInt16{
		NewNullableBase(value),
	}
}

func NewNullUInt16() UInt16 {
	return UInt16{}
}

func (nv UInt16) Equal(value UInt16) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.Value() == value.Value()
}

type Int32 struct {
	NullableBase[int32]
}

func NewInt32(value int32) Int32 {
	return Int32{
		NewNullableBase(value),
	}
}

func NewNullInt32() Int32 {
	return Int32{}
}

func (nv Int32) Equal(value Int32) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.Value() == value.Value()
}

type UInt32 struct {
	NullableBase[uint32]
}

func NewUInt32(value uint32) UInt32 {
	return UInt32{
		NewNullableBase(value),
	}
}

func NewNullUInt32() UInt32 {
	return UInt32{}
}

func (nv UInt32) Equal(value UInt32) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.Value() == value.Value()
}

type Int64 struct {
	NullableBase[int64]
}

func NewInt64(value int64) Int64 {
	return Int64{
		NewNullableBase(value),
	}
}

func NewNullInt64() Int64 {
	return Int64{}
}

func (nv Int64) Equal(value Int64) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.Value() == value.Value()
}

type UInt64 struct {
	NullableBase[uint64]
}

func NewUInt64(value uint64) UInt64 {
	return UInt64{
		NewNullableBase(value),
	}
}

func NewNullUInt64() UInt64 {
	return UInt64{}
}

func (nv UInt64) Equal(value UInt64) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.Value() == value.Value()
}

type Byte = UInt8
type Rune = Int32

type Float32 struct {
	NullableBase[float32]
}

func NewFloat32(value float32) Float32 {
	return Float32{
		NewNullableBase(value),
	}
}

func NewNullFloat32() Float32 {
	return Float32{}
}

func (nv Float32) Equal(value Float32) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.Value() == value.Value()
}

type Float64 struct {
	NullableBase[float64]
}

func NewFloat64(value float64) Float64 {
	return Float64{
		NewNullableBase(value),
	}
}

func NewNullFloat64() Float64 {
	return Float64{}
}

func (nv Float64) Equal(value Float64) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.Value() == value.Value()
}

type Bool struct {
	NullableBase[bool]
}

func NewBool(value bool) Bool {
	return Bool{
		NewNullableBase(value),
	}
}

func NewNullBool() Bool {
	return Bool{}
}

func (nv Bool) Equal(value Bool) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.Value() == value.Value()
}

type Time struct {
	NullableBase[time.Time]
}

func NewTime(value time.Time) Time {
	return Time{
		NewNullableBase(value),
	}
}

func NewNullTime() Time {
	return Time{}
}

func (nv Time) Equal(value Time) bool {
	if nv.IsNull() != value.IsNull() {
		return false
	}

	return nv.IsNull() || nv.Value().Equal(value.Value())
}
