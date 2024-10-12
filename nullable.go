package nullable

import "encoding/json"

func NewNullableBase[T any](value T) NullableBase[T] {
	return NullableBase[T]{
		valid: true,
		value: value,
	}
}

type NullableBase[T any] struct {
	valid bool
	value T
}

func (nv NullableBase[T]) Value() T {
	return nv.value
}

func (nv *NullableBase[T]) SetValue(value T) {
	nv.valid, nv.value = true, value
}

func (nv NullableBase[T]) IsNull() bool {
	return !nv.valid
}

func (nv *NullableBase[T]) SetNull() {
	var zeroValue T
	nv.valid, nv.value = false, zeroValue
}

func (nv NullableBase[T]) MarshalJSON() ([]byte, error) {
	if nv.valid {
		return json.Marshal(nv.value)
	}

	return []byte("null"), nil
}

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
