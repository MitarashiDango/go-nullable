package nullable_test

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/MitarashiDango/go-nullable/v2"
)

var (
	_ sql.Scanner      = (*nullable.NullableBase[string])(nil)
	_ driver.Valuer    = nullable.NullableBase[string]{}
	_ json.Marshaler   = nullable.NullableBase[string]{}
	_ json.Unmarshaler = (*nullable.NullableBase[string])(nil)
	_ sql.Scanner      = (*nullable.String)(nil)
	_ driver.Valuer    = nullable.String{}
	_ json.Marshaler   = nullable.String{}
	_ json.Unmarshaler = (*nullable.String)(nil)
)

type sqlNullable interface {
	sql.Scanner
	driver.Valuer
}

func pointerTo[T any](value T) *T {
	return &value
}

var errMarshalJSON = errors.New("marshal JSON error")

type marshalJSONErrorValue struct{}

func (marshalJSONErrorValue) MarshalJSON() ([]byte, error) {
	return nil, errMarshalJSON
}

func TestNullableBase_IsNull_InitValue(t *testing.T) {
	var s nullable.String

	if !s.IsNull() {
		t.Fatal("expected IsNull() to be true")
	}
}

func TestNullableBase_SetValue(t *testing.T) {
	var s nullable.String
	s.SetValue("test")

	if s.RawValue() != "test" {
		t.Fatalf("expected %q, got %q", "test", s.RawValue())
	}
	if s.IsNull() {
		t.Fatal("expected IsNull() to be false")
	}
}

func TestNullableBase_ValueOrZero_Valid(t *testing.T) {
	var s nullable.String
	s.SetValue("test")

	if s.ValueOrZero() != "test" {
		t.Fatalf("expected %q, got %q", "test", s.ValueOrZero())
	}
	if s.IsNull() {
		t.Fatal("expected IsNull() to be false")
	}
}

func TestNullableBase_ValueOrZero_Null(t *testing.T) {
	t.Run("AfterSetNull", func(t *testing.T) {
		var s nullable.String
		s.SetNull()

		if s.ValueOrZero() != "" {
			t.Fatalf("expected empty string, got %q", s.ValueOrZero())
		}
		if !s.IsNull() {
			t.Fatal("expected IsNull() to be true")
		}
	})

	t.Run("AfterSetSqlNull", func(t *testing.T) {
		var s nullable.String
		s.SetSqlNull(sql.Null[string]{
			V:     "test",
			Valid: false,
		})

		if s.ValueOrZero() != "" {
			t.Fatalf("ValueOrZero(): expected empty string, got %q", s.ValueOrZero())
		}
		if s.RawValue() != "" {
			t.Fatalf("RawValue(): expected empty string, got %q", s.RawValue())
		}
		if !s.IsNull() {
			t.Fatal("expected IsNull() to be true")
		}
	})
}

func TestNullableBase_RawValue_Valid(t *testing.T) {
	var s nullable.String
	s.SetValue("test")

	if s.RawValue() != "test" {
		t.Fatalf("expected %q, got %q", "test", s.RawValue())
	}
}

func TestNullableBase_RawValue_Null(t *testing.T) {
	var s nullable.String
	s.SetNull()

	if s.RawValue() != "" {
		t.Fatalf("expected empty string, got %q", s.RawValue())
	}
}

func TestNullableBase_SetNull(t *testing.T) {
	var s nullable.String
	s.SetValue("test")
	s.SetNull()

	if s.RawValue() != "" {
		t.Fatalf("expected empty string, got %q", s.RawValue())
	}
	if !s.IsNull() {
		t.Fatal("expected IsNull() to be true")
	}
}

func TestNullableBase_SetNull_ClearsReferenceValue(t *testing.T) {
	value := nullable.NewNullableBase([]byte("test"))

	value.SetNull()

	if value.RawValue() != nil {
		t.Fatalf("expected the stored slice to be nil, got %v", value.RawValue())
	}
	if !value.IsNull() {
		t.Fatal("expected IsNull() to be true")
	}
}

func TestNullableBase_SqlNull_Valid(t *testing.T) {
	var value nullable.String
	value.SetValue("test value")

	actual := value.SqlNull()

	if actual.V != "test value" {
		t.Fatalf("expected V = %q, got %q", "test value", actual.V)
	}
	if !actual.Valid {
		t.Fatal("expected Valid to be true")
	}
}

func TestNullableBase_SqlNull_Null(t *testing.T) {
	var value nullable.String
	value.SetNull()

	actual := value.SqlNull()

	if actual.V != "" {
		t.Fatalf("expected V = %q, got %q", "", actual.V)
	}
	if actual.Valid {
		t.Fatal("expected Valid to be false")
	}
}

func TestNullableBase_SetSqlNull_Valid(t *testing.T) {
	var value nullable.String
	value.SetSqlNull(sql.Null[string]{
		V:     "test value",
		Valid: true,
	})

	expected := nullable.NewString("test value")
	if !value.Equal(expected) {
		t.Fatalf("expected %v, got %v", expected, value)
	}
}

func TestNullableBase_SetSqlNull_Null(t *testing.T) {
	var value nullable.String
	value.SetSqlNull(sql.Null[string]{
		V:     "",
		Valid: false,
	})

	expected := nullable.NewNullString()
	if !value.Equal(expected) {
		t.Fatalf("expected %v, got %v", expected, value)
	}
}

func TestNullableBase_Value_Valid(t *testing.T) {
	tv := time.Date(2025, 6, 1, 12, 30, 45, 0, time.UTC)
	tests := []struct {
		name     string
		value    driver.Valuer
		expected driver.Value
	}{
		{"String", nullable.NewString("test"), "test"},
		{"Int", nullable.NewInt(-123), int64(-123)},
		{"UInt", nullable.NewUInt(123), int64(123)},
		{"Int8", nullable.NewInt8(-123), int64(-123)},
		{"UInt8", nullable.NewUInt8(123), int64(123)},
		{"Int16", nullable.NewInt16(-123), int64(-123)},
		{"UInt16", nullable.NewUInt16(123), int64(123)},
		{"Int32", nullable.NewInt32(-123), int64(-123)},
		{"UInt32", nullable.NewUInt32(123), int64(123)},
		{"Int64", nullable.NewInt64(-123), int64(-123)},
		{"UInt64", nullable.NewUInt64(123), int64(123)},
		{"Float32", nullable.NewFloat32(-123.456), float64(float32(-123.456))},
		{"Float64", nullable.NewFloat64(-123.456), float64(-123.456)},
		{"Bool", nullable.NewBool(true), true},
		{"Time", nullable.NewTime(tv), tv},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := tt.value.Value()
			if err != nil {
				t.Fatalf("Value() returned error: %v", err)
			}
			if actual != tt.expected {
				t.Fatalf("expected %v (%T), got %v (%T)", tt.expected, tt.expected, actual, actual)
			}
		})
	}
}

func TestNullableBase_Value_Null(t *testing.T) {
	s := nullable.NewNullString()

	v, err := s.Value()
	if err != nil {
		t.Fatalf("Value() returned error: %v", err)
	}
	if v != nil {
		t.Fatalf("expected nil, got %v", v)
	}
}

func TestNullableBase_Value_Error(t *testing.T) {
	tests := []struct {
		name  string
		value driver.Valuer
	}{
		{"UInt64HighBitSet", nullable.NewUInt64(uint64(1) << 63)},
		{"UnsupportedGenericType", nullable.NewNullableBase(struct{}{})},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := tt.value.Value()
			if err == nil {
				t.Fatal("expected Value() to return an error")
			}
			if actual != nil {
				t.Fatalf("expected nil value, got %v", actual)
			}
		})
	}
}

func TestNullableBase_Scan_Valid(t *testing.T) {
	tv := time.Date(2025, 6, 1, 12, 30, 45, 0, time.UTC)
	tests := []struct {
		name          string
		target        sqlNullable
		source        driver.Value
		expected      driver.Value
		expectedState any
	}{
		{"StringFromString", &nullable.String{}, "test", "test", pointerTo(nullable.NewString("test"))},
		{"StringFromBytes", &nullable.String{}, []byte("test"), "test", pointerTo(nullable.NewString("test"))},
		{"IntFromInt64", &nullable.Int{}, int64(-123), int64(-123), pointerTo(nullable.NewInt(-123))},
		{"UIntFromInt64", &nullable.UInt{}, int64(123), int64(123), pointerTo(nullable.NewUInt(123))},
		{
			"Float32FromFloat64", &nullable.Float32{}, float64(1.25), float64(1.25),
			pointerTo(nullable.NewFloat32(1.25)),
		},
		{
			"Float64FromFloat64", &nullable.Float64{}, float64(1.25), float64(1.25),
			pointerTo(nullable.NewFloat64(1.25)),
		},
		{"BoolFromBool", &nullable.Bool{}, true, true, pointerTo(nullable.NewBool(true))},
		{"TimeFromTime", &nullable.Time{}, tv, tv, pointerTo(nullable.NewTime(tv))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.target.Scan(tt.source); err != nil {
				t.Fatalf("Scan() returned error: %v", err)
			}
			if !reflect.DeepEqual(tt.target, tt.expectedState) {
				t.Fatalf("expected stored state %v, got %v", tt.expectedState, tt.target)
			}

			actual, err := tt.target.Value()
			if err != nil {
				t.Fatalf("Value() returned error: %v", err)
			}
			if actual != tt.expected {
				t.Fatalf("expected %v (%T), got %v (%T)", tt.expected, tt.expected, actual, actual)
			}
		})
	}
}

func TestNullableBase_Scan_Null(t *testing.T) {
	var s nullable.String
	s.SetValue("test")

	if err := s.Scan(nil); err != nil {
		t.Fatalf("Scan() returned error: %v", err)
	}

	if !s.IsNull() {
		t.Fatal("expected IsNull() to be true")
	}
	if s.RawValue() != "" {
		t.Fatalf("expected the stored value to be cleared, got %q", s.RawValue())
	}
}

func TestNullableBase_Scan_ErrorPreservesState(t *testing.T) {
	t.Run("UnsupportedType", func(t *testing.T) {
		value := nullable.NewString("before")
		if err := value.Scan(struct{}{}); err == nil {
			t.Fatal("expected Scan() to return an error")
		}
		if !value.Equal(nullable.NewString("before")) {
			t.Fatalf("expected the original value to be preserved, got %v", value)
		}
	})

	t.Run("IntegerOverflow", func(t *testing.T) {
		value := nullable.NewInt8(12)
		if err := value.Scan(int64(128)); err == nil {
			t.Fatal("expected Scan() to return an error")
		}
		if !value.Equal(nullable.NewInt8(12)) {
			t.Fatalf("expected the original value to be preserved, got %v", value)
		}
	})

	t.Run("NegativeUnsignedInteger", func(t *testing.T) {
		value := nullable.NewUInt(12)
		if err := value.Scan(int64(-1)); err == nil {
			t.Fatal("expected Scan() to return an error")
		}
		if !value.Equal(nullable.NewUInt(12)) {
			t.Fatalf("expected the original value to be preserved, got %v", value)
		}
	})
}

func TestNullableBase_MarshalJSON_ValidValue(t *testing.T) {
	tv := time.Date(2025, 6, 1, 12, 30, 45, 0, time.UTC)

	tests := []struct {
		name     string
		input    any
		expected string
	}{
		{"String", struct {
			Value nullable.String `json:"value"`
		}{nullable.NewString("test value")}, `{"value":"test value"}`},
		{"Int", struct {
			Value nullable.Int `json:"value"`
		}{nullable.NewInt(-123)}, `{"value":-123}`},
		{"UInt", struct {
			Value nullable.UInt `json:"value"`
		}{nullable.NewUInt(123)}, `{"value":123}`},
		{"Int8", struct {
			Value nullable.Int8 `json:"value"`
		}{nullable.NewInt8(-123)}, `{"value":-123}`},
		{"UInt8", struct {
			Value nullable.UInt8 `json:"value"`
		}{nullable.NewUInt8(123)}, `{"value":123}`},
		{"Int16", struct {
			Value nullable.Int16 `json:"value"`
		}{nullable.NewInt16(-123)}, `{"value":-123}`},
		{"UInt16", struct {
			Value nullable.UInt16 `json:"value"`
		}{nullable.NewUInt16(123)}, `{"value":123}`},
		{"Int32", struct {
			Value nullable.Int32 `json:"value"`
		}{nullable.NewInt32(-123)}, `{"value":-123}`},
		{"UInt32", struct {
			Value nullable.UInt32 `json:"value"`
		}{nullable.NewUInt32(123)}, `{"value":123}`},
		{"Int64", struct {
			Value nullable.Int64 `json:"value"`
		}{nullable.NewInt64(-123)}, `{"value":-123}`},
		{"UInt64", struct {
			Value nullable.UInt64 `json:"value"`
		}{nullable.NewUInt64(123)}, `{"value":123}`},
		{"Float32", struct {
			Value nullable.Float32 `json:"value"`
		}{nullable.NewFloat32(-123.456)}, `{"value":-123.456}`},
		{"Float64", struct {
			Value nullable.Float64 `json:"value"`
		}{nullable.NewFloat64(-123.456)}, `{"value":-123.456}`},
		{"Bool_True", struct {
			Value nullable.Bool `json:"value"`
		}{nullable.NewBool(true)}, `{"value":true}`},
		{"Bool_False", struct {
			Value nullable.Bool `json:"value"`
		}{nullable.NewBool(false)}, `{"value":false}`},
		{"Time", struct {
			Value nullable.Time `json:"value"`
		}{nullable.NewTime(tv)}, `{"value":"2025-06-01T12:30:45Z"}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := json.Marshal(tt.input)
			if err != nil {
				t.Fatalf("json.Marshal failed: %v", err)
			}
			if tt.expected != string(actual) {
				t.Fatalf("expected: %v, actual: %v", tt.expected, string(actual))
			}
		})
	}
}

func TestNullableBase_MarshalJSON_Null(t *testing.T) {
	tests := []struct {
		name  string
		input any
	}{
		{"String", struct {
			Value nullable.String `json:"value"`
		}{nullable.NewNullString()}},
		{"Int", struct {
			Value nullable.Int `json:"value"`
		}{nullable.NewNullInt()}},
		{"UInt", struct {
			Value nullable.UInt `json:"value"`
		}{nullable.NewNullUInt()}},
		{"Int8", struct {
			Value nullable.Int8 `json:"value"`
		}{nullable.NewNullInt8()}},
		{"UInt8", struct {
			Value nullable.UInt8 `json:"value"`
		}{nullable.NewNullUInt8()}},
		{"Int16", struct {
			Value nullable.Int16 `json:"value"`
		}{nullable.NewNullInt16()}},
		{"UInt16", struct {
			Value nullable.UInt16 `json:"value"`
		}{nullable.NewNullUInt16()}},
		{"Int32", struct {
			Value nullable.Int32 `json:"value"`
		}{nullable.NewNullInt32()}},
		{"UInt32", struct {
			Value nullable.UInt32 `json:"value"`
		}{nullable.NewNullUInt32()}},
		{"Int64", struct {
			Value nullable.Int64 `json:"value"`
		}{nullable.NewNullInt64()}},
		{"UInt64", struct {
			Value nullable.UInt64 `json:"value"`
		}{nullable.NewNullUInt64()}},
		{"Float32", struct {
			Value nullable.Float32 `json:"value"`
		}{nullable.NewNullFloat32()}},
		{"Float64", struct {
			Value nullable.Float64 `json:"value"`
		}{nullable.NewNullFloat64()}},
		{"Bool", struct {
			Value nullable.Bool `json:"value"`
		}{nullable.NewNullBool()}},
		{"Time", struct {
			Value nullable.Time `json:"value"`
		}{nullable.NewNullTime()}},
	}

	expected := `{"value":null}`

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := json.Marshal(tt.input)
			if err != nil {
				t.Fatalf("json.Marshal failed: %v", err)
			}
			if expected != string(actual) {
				t.Fatalf("expected: %v, actual: %v", expected, string(actual))
			}
		})
	}
}

func TestNullableBase_MarshalJSON_Error(t *testing.T) {
	value := nullable.NewNullableBase(marshalJSONErrorValue{})

	_, err := value.MarshalJSON()
	if !errors.Is(err, errMarshalJSON) {
		t.Fatalf("expected %v, got %v", errMarshalJSON, err)
	}
}

func TestNullableBase_UnmarshalJSON_ValidValue(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		newTarget func() any
	}{
		{"String", `{"value":"test value"}`, func() any {
			return &struct {
				Value nullable.String `json:"value"`
			}{}
		}},
		{"Int", `{"value":-123}`, func() any {
			return &struct {
				Value nullable.Int `json:"value"`
			}{}
		}},
		{"UInt", `{"value":123}`, func() any {
			return &struct {
				Value nullable.UInt `json:"value"`
			}{}
		}},
		{"Int8", `{"value":-123}`, func() any {
			return &struct {
				Value nullable.Int8 `json:"value"`
			}{}
		}},
		{"UInt8", `{"value":123}`, func() any {
			return &struct {
				Value nullable.UInt8 `json:"value"`
			}{}
		}},
		{"Int16", `{"value":-123}`, func() any {
			return &struct {
				Value nullable.Int16 `json:"value"`
			}{}
		}},
		{"UInt16", `{"value":123}`, func() any {
			return &struct {
				Value nullable.UInt16 `json:"value"`
			}{}
		}},
		{"Int32", `{"value":-123}`, func() any {
			return &struct {
				Value nullable.Int32 `json:"value"`
			}{}
		}},
		{"UInt32", `{"value":123}`, func() any {
			return &struct {
				Value nullable.UInt32 `json:"value"`
			}{}
		}},
		{"Int64", `{"value":-123}`, func() any {
			return &struct {
				Value nullable.Int64 `json:"value"`
			}{}
		}},
		{"UInt64", `{"value":123}`, func() any {
			return &struct {
				Value nullable.UInt64 `json:"value"`
			}{}
		}},
		{"Float32", `{"value":-123.456}`, func() any {
			return &struct {
				Value nullable.Float32 `json:"value"`
			}{}
		}},
		{"Float64", `{"value":-123.456}`, func() any {
			return &struct {
				Value nullable.Float64 `json:"value"`
			}{}
		}},
		{"Bool_True", `{"value":true}`, func() any {
			return &struct {
				Value nullable.Bool `json:"value"`
			}{}
		}},
		{"Bool_False", `{"value":false}`, func() any {
			return &struct {
				Value nullable.Bool `json:"value"`
			}{}
		}},
		{"Time", `{"value":"2025-06-01T12:30:45Z"}`, func() any {
			return &struct {
				Value nullable.Time `json:"value"`
			}{}
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			target := tt.newTarget()
			if err := json.Unmarshal([]byte(tt.input), target); err != nil {
				t.Fatalf("json.Unmarshal failed: %v", err)
			}
			actual, err := json.Marshal(target)
			if err != nil {
				t.Fatalf("json.Marshal failed: %v", err)
			}
			if tt.input != string(actual) {
				t.Fatalf("expected: %v, actual: %v", tt.input, string(actual))
			}
		})
	}
}

func TestNullableBase_UnmarshalJSON_Null(t *testing.T) {
	tests := []struct {
		name      string
		newTarget func() any
	}{
		{"String", func() any {
			return &struct {
				Value nullable.String `json:"value"`
			}{}
		}},
		{"Int", func() any {
			return &struct {
				Value nullable.Int `json:"value"`
			}{}
		}},
		{"UInt", func() any {
			return &struct {
				Value nullable.UInt `json:"value"`
			}{}
		}},
		{"Int8", func() any {
			return &struct {
				Value nullable.Int8 `json:"value"`
			}{}
		}},
		{"UInt8", func() any {
			return &struct {
				Value nullable.UInt8 `json:"value"`
			}{}
		}},
		{"Int16", func() any {
			return &struct {
				Value nullable.Int16 `json:"value"`
			}{}
		}},
		{"UInt16", func() any {
			return &struct {
				Value nullable.UInt16 `json:"value"`
			}{}
		}},
		{"Int32", func() any {
			return &struct {
				Value nullable.Int32 `json:"value"`
			}{}
		}},
		{"UInt32", func() any {
			return &struct {
				Value nullable.UInt32 `json:"value"`
			}{}
		}},
		{"Int64", func() any {
			return &struct {
				Value nullable.Int64 `json:"value"`
			}{}
		}},
		{"UInt64", func() any {
			return &struct {
				Value nullable.UInt64 `json:"value"`
			}{}
		}},
		{"Float32", func() any {
			return &struct {
				Value nullable.Float32 `json:"value"`
			}{}
		}},
		{"Float64", func() any {
			return &struct {
				Value nullable.Float64 `json:"value"`
			}{}
		}},
		{"Bool", func() any {
			return &struct {
				Value nullable.Bool `json:"value"`
			}{}
		}},
		{"Time", func() any {
			return &struct {
				Value nullable.Time `json:"value"`
			}{}
		}},
	}

	input := `{"value":null}`
	expected := `{"value":null}`

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			target := tt.newTarget()
			if err := json.Unmarshal([]byte(input), target); err != nil {
				t.Fatalf("json.Unmarshal failed: %v", err)
			}
			actual, err := json.Marshal(target)
			if err != nil {
				t.Fatalf("json.Marshal failed: %v", err)
			}
			if expected != string(actual) {
				t.Fatalf("expected: %v, actual: %v", expected, string(actual))
			}
		})
	}
}

func TestNullableBase_UnmarshalJSON_NullClearsValue(t *testing.T) {
	value := nullable.NewString("before")

	if err := json.Unmarshal([]byte("null"), &value); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}
	if !value.IsNull() {
		t.Fatal("expected IsNull() to be true")
	}
	if value.RawValue() != "" {
		t.Fatalf("expected the stored value to be cleared, got %q", value.RawValue())
	}
}

func TestNullableBase_UnmarshalJSON_ErrorPreservesState(t *testing.T) {
	t.Run("TypeMismatch", func(t *testing.T) {
		value := nullable.NewString("before")
		if err := json.Unmarshal([]byte("123"), &value); err == nil {
			t.Fatal("expected json.Unmarshal to return an error")
		}
		if !value.Equal(nullable.NewString("before")) {
			t.Fatalf("expected the original value to be preserved, got %v", value)
		}
	})

	t.Run("IntegerOverflow", func(t *testing.T) {
		value := nullable.NewInt8(12)
		if err := json.Unmarshal([]byte("128"), &value); err == nil {
			t.Fatal("expected json.Unmarshal to return an error")
		}
		if !value.Equal(nullable.NewInt8(12)) {
			t.Fatalf("expected the original value to be preserved, got %v", value)
		}
	})
}
