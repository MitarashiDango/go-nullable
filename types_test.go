package nullable_test

import (
	"math"
	"testing"
	"time"

	"github.com/MitarashiDango/go-nullable/v2"
)

func TestString_Equal(t *testing.T) {
	tests := []struct {
		name     string
		v1       nullable.String
		v2       nullable.String
		expected bool
	}{
		{"BothValid_SameValue", nullable.NewString("test"), nullable.NewString("test"), true},
		{"BothNull", nullable.NewNullString(), nullable.NewNullString(), true},
		{"BothValid_DifferentValue", nullable.NewString("test1"), nullable.NewString("test2"), false},
		{"ValidAndNull", nullable.NewString("test1"), nullable.NewNullString(), false},
		{"NullAndValid", nullable.NewNullString(), nullable.NewString("test2"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v1.Equal(tt.v2); got != tt.expected {
				t.Fatalf("Equal() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestInt_Equal(t *testing.T) {
	tests := []struct {
		name     string
		v1       nullable.Int
		v2       nullable.Int
		expected bool
	}{
		{"BothValid_SameValue", nullable.NewInt(123), nullable.NewInt(123), true},
		{"BothNull", nullable.NewNullInt(), nullable.NewNullInt(), true},
		{"BothValid_DifferentValue", nullable.NewInt(123), nullable.NewInt(456), false},
		{"ValidAndNull", nullable.NewInt(123), nullable.NewNullInt(), false},
		{"NullAndValid", nullable.NewNullInt(), nullable.NewInt(123), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v1.Equal(tt.v2); got != tt.expected {
				t.Fatalf("Equal() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestUInt_Equal(t *testing.T) {
	tests := []struct {
		name     string
		v1       nullable.UInt
		v2       nullable.UInt
		expected bool
	}{
		{"BothValid_SameValue", nullable.NewUInt(123), nullable.NewUInt(123), true},
		{"BothNull", nullable.NewNullUInt(), nullable.NewNullUInt(), true},
		{"BothValid_DifferentValue", nullable.NewUInt(123), nullable.NewUInt(456), false},
		{"ValidAndNull", nullable.NewUInt(123), nullable.NewNullUInt(), false},
		{"NullAndValid", nullable.NewNullUInt(), nullable.NewUInt(123), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v1.Equal(tt.v2); got != tt.expected {
				t.Fatalf("Equal() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestInt8_Equal(t *testing.T) {
	tests := []struct {
		name     string
		v1       nullable.Int8
		v2       nullable.Int8
		expected bool
	}{
		{"BothValid_SameValue", nullable.NewInt8(12), nullable.NewInt8(12), true},
		{"BothNull", nullable.NewNullInt8(), nullable.NewNullInt8(), true},
		{"BothValid_DifferentValue", nullable.NewInt8(12), nullable.NewInt8(34), false},
		{"ValidAndNull", nullable.NewInt8(12), nullable.NewNullInt8(), false},
		{"NullAndValid", nullable.NewNullInt8(), nullable.NewInt8(12), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v1.Equal(tt.v2); got != tt.expected {
				t.Fatalf("Equal() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestUInt8_Equal(t *testing.T) {
	tests := []struct {
		name     string
		v1       nullable.UInt8
		v2       nullable.UInt8
		expected bool
	}{
		{"BothValid_SameValue", nullable.NewUInt8(12), nullable.NewUInt8(12), true},
		{"BothNull", nullable.NewNullUInt8(), nullable.NewNullUInt8(), true},
		{"BothValid_DifferentValue", nullable.NewUInt8(12), nullable.NewUInt8(34), false},
		{"ValidAndNull", nullable.NewUInt8(12), nullable.NewNullUInt8(), false},
		{"NullAndValid", nullable.NewNullUInt8(), nullable.NewUInt8(12), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v1.Equal(tt.v2); got != tt.expected {
				t.Fatalf("Equal() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestInt16_Equal(t *testing.T) {
	tests := []struct {
		name     string
		v1       nullable.Int16
		v2       nullable.Int16
		expected bool
	}{
		{"BothValid_SameValue", nullable.NewInt16(123), nullable.NewInt16(123), true},
		{"BothNull", nullable.NewNullInt16(), nullable.NewNullInt16(), true},
		{"BothValid_DifferentValue", nullable.NewInt16(123), nullable.NewInt16(456), false},
		{"ValidAndNull", nullable.NewInt16(123), nullable.NewNullInt16(), false},
		{"NullAndValid", nullable.NewNullInt16(), nullable.NewInt16(123), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v1.Equal(tt.v2); got != tt.expected {
				t.Fatalf("Equal() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestUInt16_Equal(t *testing.T) {
	tests := []struct {
		name     string
		v1       nullable.UInt16
		v2       nullable.UInt16
		expected bool
	}{
		{"BothValid_SameValue", nullable.NewUInt16(123), nullable.NewUInt16(123), true},
		{"BothNull", nullable.NewNullUInt16(), nullable.NewNullUInt16(), true},
		{"BothValid_DifferentValue", nullable.NewUInt16(123), nullable.NewUInt16(456), false},
		{"ValidAndNull", nullable.NewUInt16(123), nullable.NewNullUInt16(), false},
		{"NullAndValid", nullable.NewNullUInt16(), nullable.NewUInt16(123), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v1.Equal(tt.v2); got != tt.expected {
				t.Fatalf("Equal() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestInt32_Equal(t *testing.T) {
	tests := []struct {
		name     string
		v1       nullable.Int32
		v2       nullable.Int32
		expected bool
	}{
		{"BothValid_SameValue", nullable.NewInt32(123), nullable.NewInt32(123), true},
		{"BothNull", nullable.NewNullInt32(), nullable.NewNullInt32(), true},
		{"BothValid_DifferentValue", nullable.NewInt32(123), nullable.NewInt32(456), false},
		{"ValidAndNull", nullable.NewInt32(123), nullable.NewNullInt32(), false},
		{"NullAndValid", nullable.NewNullInt32(), nullable.NewInt32(123), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v1.Equal(tt.v2); got != tt.expected {
				t.Fatalf("Equal() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestUInt32_Equal(t *testing.T) {
	tests := []struct {
		name     string
		v1       nullable.UInt32
		v2       nullable.UInt32
		expected bool
	}{
		{"BothValid_SameValue", nullable.NewUInt32(123), nullable.NewUInt32(123), true},
		{"BothNull", nullable.NewNullUInt32(), nullable.NewNullUInt32(), true},
		{"BothValid_DifferentValue", nullable.NewUInt32(123), nullable.NewUInt32(456), false},
		{"ValidAndNull", nullable.NewUInt32(123), nullable.NewNullUInt32(), false},
		{"NullAndValid", nullable.NewNullUInt32(), nullable.NewUInt32(123), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v1.Equal(tt.v2); got != tt.expected {
				t.Fatalf("Equal() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestInt64_Equal(t *testing.T) {
	tests := []struct {
		name     string
		v1       nullable.Int64
		v2       nullable.Int64
		expected bool
	}{
		{"BothValid_SameValue", nullable.NewInt64(123), nullable.NewInt64(123), true},
		{"BothNull", nullable.NewNullInt64(), nullable.NewNullInt64(), true},
		{"BothValid_DifferentValue", nullable.NewInt64(123), nullable.NewInt64(456), false},
		{"ValidAndNull", nullable.NewInt64(123), nullable.NewNullInt64(), false},
		{"NullAndValid", nullable.NewNullInt64(), nullable.NewInt64(123), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v1.Equal(tt.v2); got != tt.expected {
				t.Fatalf("Equal() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestUInt64_Equal(t *testing.T) {
	tests := []struct {
		name     string
		v1       nullable.UInt64
		v2       nullable.UInt64
		expected bool
	}{
		{"BothValid_SameValue", nullable.NewUInt64(123), nullable.NewUInt64(123), true},
		{"BothNull", nullable.NewNullUInt64(), nullable.NewNullUInt64(), true},
		{"BothValid_DifferentValue", nullable.NewUInt64(123), nullable.NewUInt64(456), false},
		{"ValidAndNull", nullable.NewUInt64(123), nullable.NewNullUInt64(), false},
		{"NullAndValid", nullable.NewNullUInt64(), nullable.NewUInt64(123), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v1.Equal(tt.v2); got != tt.expected {
				t.Fatalf("Equal() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestFloat32_Equal(t *testing.T) {
	tests := []struct {
		name     string
		v1       nullable.Float32
		v2       nullable.Float32
		expected bool
	}{
		{"BothValid_SameValue", nullable.NewFloat32(1.23), nullable.NewFloat32(1.23), true},
		{"BothNull", nullable.NewNullFloat32(), nullable.NewNullFloat32(), true},
		{"BothValid_DifferentValue", nullable.NewFloat32(1.23), nullable.NewFloat32(4.56), false},
		{"ValidAndNull", nullable.NewFloat32(1.23), nullable.NewNullFloat32(), false},
		{"NullAndValid", nullable.NewNullFloat32(), nullable.NewFloat32(1.23), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v1.Equal(tt.v2); got != tt.expected {
				t.Fatalf("Equal() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestFloat64_Equal(t *testing.T) {
	tests := []struct {
		name     string
		v1       nullable.Float64
		v2       nullable.Float64
		expected bool
	}{
		{"BothValid_SameValue", nullable.NewFloat64(1.23), nullable.NewFloat64(1.23), true},
		{"BothNull", nullable.NewNullFloat64(), nullable.NewNullFloat64(), true},
		{"BothValid_DifferentValue", nullable.NewFloat64(1.23), nullable.NewFloat64(4.56), false},
		{"ValidAndNull", nullable.NewFloat64(1.23), nullable.NewNullFloat64(), false},
		{"NullAndValid", nullable.NewNullFloat64(), nullable.NewFloat64(1.23), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v1.Equal(tt.v2); got != tt.expected {
				t.Fatalf("Equal() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestBool_Equal(t *testing.T) {
	tests := []struct {
		name     string
		v1       nullable.Bool
		v2       nullable.Bool
		expected bool
	}{
		{"BothValid_SameValue", nullable.NewBool(true), nullable.NewBool(true), true},
		{"BothNull", nullable.NewNullBool(), nullable.NewNullBool(), true},
		{"BothValid_DifferentValue", nullable.NewBool(true), nullable.NewBool(false), false},
		{"ValidAndNull", nullable.NewBool(true), nullable.NewNullBool(), false},
		{"NullAndValid", nullable.NewNullBool(), nullable.NewBool(true), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v1.Equal(tt.v2); got != tt.expected {
				t.Fatalf("Equal() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestTime_Equal(t *testing.T) {
	t1 := time.Date(2025, 6, 1, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2025, 6, 2, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		v1       nullable.Time
		v2       nullable.Time
		expected bool
	}{
		{"BothValid_SameValue", nullable.NewTime(t1), nullable.NewTime(t1), true},
		{"BothNull", nullable.NewNullTime(), nullable.NewNullTime(), true},
		{"BothValid_DifferentValue", nullable.NewTime(t1), nullable.NewTime(t2), false},
		{"ValidAndNull", nullable.NewTime(t1), nullable.NewNullTime(), false},
		{"NullAndValid", nullable.NewNullTime(), nullable.NewTime(t1), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v1.Equal(tt.v2); got != tt.expected {
				t.Fatalf("Equal() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestFloat32_EqualBits(t *testing.T) {
	tests := []struct {
		name     string
		v1       nullable.Float32
		v2       nullable.Float32
		expected bool
	}{
		{"BothValid_SameValue", nullable.NewFloat32(1.23), nullable.NewFloat32(1.23), true},
		{"BothValid_DifferentValue", nullable.NewFloat32(1.23), nullable.NewFloat32(4.56), false},
		{"BothNull", nullable.NewNullFloat32(), nullable.NewNullFloat32(), true},
		{"OneNull", nullable.NewFloat32(1.23), nullable.NewNullFloat32(), false},
		{"NaN", nullable.NewFloat32(float32(math.NaN())), nullable.NewFloat32(float32(math.NaN())), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v1.EqualBits(tt.v2); got != tt.expected {
				t.Fatalf("EqualBits() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestFloat64_EqualBits(t *testing.T) {
	tests := []struct {
		name     string
		v1       nullable.Float64
		v2       nullable.Float64
		expected bool
	}{
		{"BothValid_SameValue", nullable.NewFloat64(1.23), nullable.NewFloat64(1.23), true},
		{"BothValid_DifferentValue", nullable.NewFloat64(1.23), nullable.NewFloat64(4.56), false},
		{"BothNull", nullable.NewNullFloat64(), nullable.NewNullFloat64(), true},
		{"OneNull", nullable.NewFloat64(1.23), nullable.NewNullFloat64(), false},
		{"NaN", nullable.NewFloat64(math.NaN()), nullable.NewFloat64(math.NaN()), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v1.EqualBits(tt.v2); got != tt.expected {
				t.Fatalf("EqualBits() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestFloat32_EqualEpsilon(t *testing.T) {
	tests := []struct {
		name     string
		v1       nullable.Float32
		v2       nullable.Float32
		epsilon  float32
		expected bool
	}{
		{"WithinEpsilon", nullable.NewFloat32(1.0), nullable.NewFloat32(1.0001), 0.001, true},
		{"OutsideEpsilon", nullable.NewFloat32(1.0), nullable.NewFloat32(1.01), 0.001, false},
		{"ExactBoundary", nullable.NewFloat32(1.0), nullable.NewFloat32(1.5), 0.5, true},
		{"BothNull", nullable.NewNullFloat32(), nullable.NewNullFloat32(), 0.001, true},
		{"OneNull", nullable.NewFloat32(1.0), nullable.NewNullFloat32(), 0.001, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v1.EqualEpsilon(tt.v2, tt.epsilon); got != tt.expected {
				t.Fatalf("EqualEpsilon() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestFloat64_EqualEpsilon(t *testing.T) {
	tests := []struct {
		name     string
		v1       nullable.Float64
		v2       nullable.Float64
		epsilon  float64
		expected bool
	}{
		{"WithinEpsilon", nullable.NewFloat64(1.0), nullable.NewFloat64(1.0000001), 0.001, true},
		{"OutsideEpsilon", nullable.NewFloat64(1.0), nullable.NewFloat64(1.01), 0.001, false},
		{"ExactBoundary", nullable.NewFloat64(1.0), nullable.NewFloat64(1.5), 0.5, true},
		{"BothNull", nullable.NewNullFloat64(), nullable.NewNullFloat64(), 0.001, true},
		{"OneNull", nullable.NewFloat64(1.0), nullable.NewNullFloat64(), 0.001, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v1.EqualEpsilon(tt.v2, tt.epsilon); got != tt.expected {
				t.Fatalf("EqualEpsilon() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
