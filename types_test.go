package nullable_test

import (
	"encoding/json"
	"testing"

	"github.com/MitarashiDango/go-nullable"
)

func Test_MarshalJSON_ValidString(t *testing.T) {
	type Test struct {
		Value nullable.String `json:"value"`
	}

	input := &Test{
		Value: nullable.NewString("test value"),
	}

	expected := `{"value":"test value"}`

	actual, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	if expected != string(actual) {
		t.Fail()
		t.Logf("expected: %v, actual: %v", expected, string(actual))
	}
}

func Test_MarshalJSON_NullString(t *testing.T) {
	type Test struct {
		Value nullable.String `json:"value"`
	}

	input := &Test{
		Value: nullable.NewNullString(),
	}

	expected := `{"value":null}`

	actual, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	if expected != string(actual) {
		t.Fail()
		t.Logf("expected: %v, actual: %v", expected, string(actual))
	}
}

func Test_MarshalJSON_ValidInt(t *testing.T) {
	type Test struct {
		Value nullable.Int `json:"value"`
	}

	input := &Test{
		Value: nullable.NewInt(-123),
	}

	expected := `{"value":-123}`

	actual, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	if expected != string(actual) {
		t.Fail()
		t.Logf("expected: %v, actual: %v", expected, string(actual))
	}
}

func Test_MarshalJSON_NullInt(t *testing.T) {
	type Test struct {
		Value nullable.Int `json:"value"`
	}

	input := &Test{
		Value: nullable.NewNullInt(),
	}

	expected := `{"value":null}`

	actual, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	if expected != string(actual) {
		t.Fail()
		t.Logf("expected: %v, actual: %v", expected, string(actual))
	}
}

func Test_MarshalJSON_ValidUInt(t *testing.T) {
	type Test struct {
		Value nullable.UInt `json:"value"`
	}

	input := &Test{
		Value: nullable.NewUInt(123),
	}

	expected := `{"value":123}`

	actual, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	if expected != string(actual) {
		t.Fail()
		t.Logf("expected: %v, actual: %v", expected, string(actual))
	}
}

func Test_MarshalJSON_NullUInt(t *testing.T) {
	type Test struct {
		Value nullable.UInt `json:"value"`
	}

	input := &Test{
		Value: nullable.NewNullUInt(),
	}

	expected := `{"value":null}`

	actual, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	if expected != string(actual) {
		t.Fail()
		t.Logf("expected: %v, actual: %v", expected, string(actual))
	}
}

func Test_MarshalJSON_ValidInt8(t *testing.T) {
	type Test struct {
		Value nullable.Int8 `json:"value"`
	}

	input := &Test{
		Value: nullable.NewInt8(-123),
	}

	expected := `{"value":-123}`

	actual, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	if expected != string(actual) {
		t.Fail()
		t.Logf("expected: %v, actual: %v", expected, string(actual))
	}
}

func Test_MarshalJSON_NullInt8(t *testing.T) {
	type Test struct {
		Value nullable.Int8 `json:"value"`
	}

	input := &Test{
		Value: nullable.NewNullInt8(),
	}

	expected := `{"value":null}`

	actual, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	if expected != string(actual) {
		t.Fail()
		t.Logf("expected: %v, actual: %v", expected, string(actual))
	}
}

func Test_MarshalJSON_ValidUInt8(t *testing.T) {
	type Test struct {
		Value nullable.UInt8 `json:"value"`
	}

	input := &Test{
		Value: nullable.NewUInt8(123),
	}

	expected := `{"value":123}`

	actual, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	if expected != string(actual) {
		t.Fail()
		t.Logf("expected: %v, actual: %v", expected, string(actual))
	}
}

func Test_MarshalJSON_NullUInt8(t *testing.T) {
	type Test struct {
		Value nullable.UInt8 `json:"value"`
	}

	input := &Test{
		Value: nullable.NewNullUInt8(),
	}

	expected := `{"value":null}`

	actual, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	if expected != string(actual) {
		t.Fail()
		t.Logf("expected: %v, actual: %v", expected, string(actual))
	}
}

func Test_MarshalJSON_ValidInt16(t *testing.T) {
	type Test struct {
		Value nullable.Int16 `json:"value"`
	}

	input := &Test{
		Value: nullable.NewInt16(-123),
	}

	expected := `{"value":-123}`

	actual, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	if expected != string(actual) {
		t.Fail()
		t.Logf("expected: %v, actual: %v", expected, string(actual))
	}
}

func Test_MarshalJSON_NullInt16(t *testing.T) {
	type Test struct {
		Value nullable.Int16 `json:"value"`
	}

	input := &Test{
		Value: nullable.NewNullInt16(),
	}

	expected := `{"value":null}`

	actual, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	if expected != string(actual) {
		t.Fail()
		t.Logf("expected: %v, actual: %v", expected, string(actual))
	}
}

func Test_MarshalJSON_ValidUInt16(t *testing.T) {
	type Test struct {
		Value nullable.UInt16 `json:"value"`
	}

	input := &Test{
		Value: nullable.NewUInt16(123),
	}

	expected := `{"value":123}`

	actual, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	if expected != string(actual) {
		t.Fail()
		t.Logf("expected: %v, actual: %v", expected, string(actual))
	}
}

func Test_MarshalJSON_NullUInt16(t *testing.T) {
	type Test struct {
		Value nullable.UInt16 `json:"value"`
	}

	input := &Test{
		Value: nullable.NewNullUInt16(),
	}

	expected := `{"value":null}`

	actual, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	if expected != string(actual) {
		t.Fail()
		t.Logf("expected: %v, actual: %v", expected, string(actual))
	}
}

func Test_MarshalJSON_ValidInt32(t *testing.T) {
	type Test struct {
		Value nullable.Int32 `json:"value"`
	}

	input := &Test{
		Value: nullable.NewInt32(-123),
	}

	expected := `{"value":-123}`

	actual, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	if expected != string(actual) {
		t.Fail()
		t.Logf("expected: %v, actual: %v", expected, string(actual))
	}
}

func Test_MarshalJSON_NullInt32(t *testing.T) {
	type Test struct {
		Value nullable.Int32 `json:"value"`
	}

	input := &Test{
		Value: nullable.NewNullInt32(),
	}

	expected := `{"value":null}`

	actual, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	if expected != string(actual) {
		t.Fail()
		t.Logf("expected: %v, actual: %v", expected, string(actual))
	}
}

func Test_MarshalJSON_ValidUInt32(t *testing.T) {
	type Test struct {
		Value nullable.UInt32 `json:"value"`
	}

	input := &Test{
		Value: nullable.NewUInt32(123),
	}

	expected := `{"value":123}`

	actual, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	if expected != string(actual) {
		t.Fail()
		t.Logf("expected: %v, actual: %v", expected, string(actual))
	}
}

func Test_MarshalJSON_NullUInt32(t *testing.T) {
	type Test struct {
		Value nullable.UInt32 `json:"value"`
	}

	input := &Test{
		Value: nullable.NewNullUInt32(),
	}

	expected := `{"value":null}`

	actual, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	if expected != string(actual) {
		t.Fail()
		t.Logf("expected: %v, actual: %v", expected, string(actual))
	}
}

func Test_MarshalJSON_ValidInt64(t *testing.T) {
	type Test struct {
		Value nullable.Int64 `json:"value"`
	}

	input := &Test{
		Value: nullable.NewInt64(-123),
	}

	expected := `{"value":-123}`

	actual, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	if expected != string(actual) {
		t.Fail()
		t.Logf("expected: %v, actual: %v", expected, string(actual))
	}
}

func Test_MarshalJSON_NullInt64(t *testing.T) {
	type Test struct {
		Value nullable.Int64 `json:"value"`
	}

	input := &Test{
		Value: nullable.NewNullInt64(),
	}

	expected := `{"value":null}`

	actual, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	if expected != string(actual) {
		t.Fail()
		t.Logf("expected: %v, actual: %v", expected, string(actual))
	}
}

func Test_MarshalJSON_ValidUInt64(t *testing.T) {
	type Test struct {
		Value nullable.UInt64 `json:"value"`
	}

	input := &Test{
		Value: nullable.NewUInt64(123),
	}

	expected := `{"value":123}`

	actual, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	if expected != string(actual) {
		t.Fail()
		t.Logf("expected: %v, actual: %v", expected, string(actual))
	}
}

func Test_MarshalJSON_NullUInt64(t *testing.T) {
	type Test struct {
		Value nullable.UInt64 `json:"value"`
	}

	input := &Test{
		Value: nullable.NewNullUInt64(),
	}

	expected := `{"value":null}`

	actual, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	if expected != string(actual) {
		t.Fail()
		t.Logf("expected: %v, actual: %v", expected, string(actual))
	}
}

func Test_MarshalJSON_ValidFloat32(t *testing.T) {
	type Test struct {
		Value nullable.Float32 `json:"value"`
	}

	input := &Test{
		Value: nullable.NewFloat32(-123.456),
	}

	expected := `{"value":-123.456}`

	actual, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	if expected != string(actual) {
		t.Fail()
		t.Logf("expected: %v, actual: %v", expected, string(actual))
	}
}

func Test_MarshalJSON_NullFloat32(t *testing.T) {
	type Test struct {
		Value nullable.Float32 `json:"value"`
	}

	input := &Test{
		Value: nullable.NewNullFloat32(),
	}

	expected := `{"value":null}`

	actual, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	if expected != string(actual) {
		t.Fail()
		t.Logf("expected: %v, actual: %v", expected, string(actual))
	}
}

func Test_MarshalJSON_ValidFloat64(t *testing.T) {
	type Test struct {
		Value nullable.Float64 `json:"value"`
	}

	input := &Test{
		Value: nullable.NewFloat64(-123.456),
	}

	expected := `{"value":-123.456}`

	actual, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	if expected != string(actual) {
		t.Fail()
		t.Logf("expected: %v, actual: %v", expected, string(actual))
	}
}

func Test_MarshalJSON_NullFloat64(t *testing.T) {
	type Test struct {
		Value nullable.Float64 `json:"value"`
	}

	input := &Test{
		Value: nullable.NewNullFloat64(),
	}

	expected := `{"value":null}`

	actual, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	if expected != string(actual) {
		t.Fail()
		t.Logf("expected: %v, actual: %v", expected, string(actual))
	}
}

func Test_MarshalJSON_ValidBool_True(t *testing.T) {
	type Test struct {
		Value nullable.Bool `json:"value"`
	}

	input := &Test{
		Value: nullable.NewBool(true),
	}

	expected := `{"value":true}`

	actual, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	if expected != string(actual) {
		t.Fail()
		t.Logf("expected: %v, actual: %v", expected, string(actual))
	}
}

func Test_MarshalJSON_ValidBool_False(t *testing.T) {
	type Test struct {
		Value nullable.Bool `json:"value"`
	}

	input := &Test{
		Value: nullable.NewBool(false),
	}

	expected := `{"value":false}`

	actual, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	if expected != string(actual) {
		t.Fail()
		t.Logf("expected: %v, actual: %v", expected, string(actual))
	}
}

func Test_MarshalJSON_NullBool(t *testing.T) {
	type Test struct {
		Value nullable.Bool `json:"value"`
	}

	input := &Test{
		Value: nullable.NewNullBool(),
	}

	expected := `{"value":null}`

	actual, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	if expected != string(actual) {
		t.Fail()
		t.Logf("expected: %v, actual: %v", expected, string(actual))
	}
}
