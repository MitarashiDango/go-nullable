package nullable

import "encoding/json"

func NewNullable[T any](value T) Nullable[T] {
	return Nullable[T]{
		valid: true,
		value: value,
	}
}

type Nullable[T any] struct {
	valid bool
	value T
}

func (nv Nullable[T]) Value() T {
	return nv.value
}

func (nv *Nullable[T]) SetValue(value T) {
	nv.valid, nv.value = true, value
}

func (nv Nullable[T]) IsNull() bool {
	return !nv.valid
}

func (nv *Nullable[T]) SetNull() {
	var zeroValue T
	nv.valid, nv.value = false, zeroValue
}

func (nv Nullable[T]) MarshalJSON() ([]byte, error) {
	if nv.valid {
		return json.Marshal(nv.value)
	}

	return []byte("null"), nil
}

func (nv *Nullable[T]) UnmarshalJSON(data []byte) error {
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

func NewComparableNullable[T comparable](value T) ComparableNullable[T] {
	return ComparableNullable[T]{
		Nullable[T]{
			valid: true,
			value: value,
		},
	}
}

type ComparableNullable[T comparable] struct {
	Nullable[T]
}

func (nv ComparableNullable[T]) Equals(value ComparableNullable[T]) bool {
	if nv.valid != value.valid {
		return false
	}

	if !value.valid {
		return true
	}

	return nv.value == value.value
}
