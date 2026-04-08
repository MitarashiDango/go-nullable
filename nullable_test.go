package nullable_test

import (
	"database/sql"
	"encoding/json"
	"testing"
	"time"

	"github.com/MitarashiDango/go-nullable/v2"
)

func TestNullableBase_Value_InitValue(t *testing.T) {
	var s nullable.String

	if s.Value() != "" {
		t.Fatalf("expected empty string, got %q", s.Value())
	}
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

	if s.Value() != "test" {
		t.Fatalf("expected %q, got %q", "test", s.Value())
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
		if s.Value() != "test" {
			t.Fatalf("Value(): expected %q, got %q", "test", s.Value())
		}
		if s.RawValue() != "test" {
			t.Fatalf("RawValue(): expected %q, got %q", "test", s.RawValue())
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

	if s.Value() != "" {
		t.Fatalf("expected empty string, got %q", s.Value())
	}
	if !s.IsNull() {
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
