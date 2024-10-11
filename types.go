package nullable

type String = ComparableNullable[string]

type Int = ComparableNullable[int]
type UInt = ComparableNullable[uint]
type Int8 = ComparableNullable[int8]
type UInt8 = ComparableNullable[uint8]
type Int16 = ComparableNullable[int16]
type UInt16 = ComparableNullable[uint16]
type Int32 = ComparableNullable[int32]
type UInt32 = ComparableNullable[uint32]
type Int64 = ComparableNullable[int64]
type UInt64 = ComparableNullable[uint64]
type Byte = UInt8
type Rune = Int32

type Float32 = ComparableNullable[float32]
type Float64 = ComparableNullable[float64]

type Bool = ComparableNullable[bool]
