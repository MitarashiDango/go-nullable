// Package nullable provides nullable value types with SQL and JSON support.
package nullable

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

// NewNullableBase returns a non-null NullableBase containing value.
func NewNullableBase[T any](value T) NullableBase[T] {
	return NullableBase[T]{
		valid: true,
		value: value,
	}
}

// NullableBase stores a nullable value and implements SQL and JSON interfaces.
type NullableBase[T any] struct {
	valid bool
	value T
}

// Value returns nil for null values, otherwise the converted driver value.
func (nv NullableBase[T]) Value() (driver.Value, error) {
	if !nv.valid {
		return nil, nil
	}
	return driver.DefaultParameterConverter.ConvertValue(nv.value)
}

// RawValue returns the stored value without checking whether it is null.
func (nv NullableBase[T]) RawValue() T {
	return nv.value
}

// ValueOrZero returns the stored value, or the zero value when it is null.
func (nv NullableBase[T]) ValueOrZero() T {
	if nv.valid {
		return nv.value
	}

	var zeroValue T
	return zeroValue
}

// SetValue stores value and marks the receiver as non-null.
func (nv *NullableBase[T]) SetValue(value T) {
	nv.valid, nv.value = true, value
}

// IsNull reports whether the receiver is null.
func (nv NullableBase[T]) IsNull() bool {
	return !nv.valid
}

// SetNull marks the receiver as null and clears the stored value.
func (nv *NullableBase[T]) SetNull() {
	var zeroValue T
	nv.valid, nv.value = false, zeroValue
}

// SqlNull returns a sql.Null with the same value and null state.
func (nv NullableBase[T]) SqlNull() sql.Null[T] {
	return sql.Null[T]{
		V:     nv.value,
		Valid: nv.valid,
	}
}

// SetSqlNull stores the value and null state from sql.Null.
func (nv *NullableBase[T]) SetSqlNull(value sql.Null[T]) {
	if !value.Valid {
		nv.SetNull()
		return
	}

	nv.SetValue(value.V)
}

// Scan stores a database value using sql.Null scanning rules.
func (nv *NullableBase[T]) Scan(src any) error {
	var s sql.Null[T]
	if err := s.Scan(src); err != nil {
		return err
	}
	nv.SetSqlNull(s)
	return nil
}

// MarshalJSON encodes null when the receiver is null.
func (nv NullableBase[T]) MarshalJSON() ([]byte, error) {
	if nv.valid {
		return json.Marshal(nv.value)
	}

	return []byte("null"), nil
}

// UnmarshalJSON stores null for JSON null, otherwise stores the decoded value.
func (nv *NullableBase[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		var zeroValue T
		nv.valid, nv.value = false, zeroValue
		return nil
	}

	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	nv.valid, nv.value = true, v
	return nil
}
