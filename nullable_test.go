package nullable_test

import (
	"database/sql"
	"testing"

	"github.com/MitarashiDango/go-nullable"
)

func Test_Nullable_Value_InitValue(t *testing.T) {
	var s nullable.String

	if s.Value() != "" {
		t.FailNow()
	}
}

func Test_Nullable_IsNull_InitValue(t *testing.T) {
	var s nullable.String

	if !s.IsNull() {
		t.FailNow()
	}
}

func Test_Nullable_SetValue(t *testing.T) {
	var s nullable.String
	s.SetValue("test")

	if s.Value() != "test" {
		t.FailNow()
	}

	if s.IsNull() {
		t.FailNow()
	}
}

func Test_Nullable_SetNull(t *testing.T) {
	var s nullable.String
	s.SetValue("test")
	s.SetNull()

	if s.Value() != "" {
		t.FailNow()
	}

	if !s.IsNull() {
		t.FailNow()
	}
}

func Test_ComparableNullable_Equal_Valid_001(t *testing.T) {
	var s1 nullable.String
	var s2 nullable.String
	s1.SetValue("test")
	s2.SetValue("test")

	if !s1.Equal(s2) {
		t.FailNow()
	}
}

func Test_ComparableNullable_Equal_Valid_002(t *testing.T) {
	var s1 nullable.String
	var s2 nullable.String

	if !s1.Equal(s2) {
		t.FailNow()
	}
}

func Test_ComparableNullable_Equal_Invalid_001(t *testing.T) {
	var s1 nullable.String
	var s2 nullable.String
	s1.SetValue("test1")
	s2.SetValue("test2")

	if s1.Equal(s2) {
		t.FailNow()
	}
}

func Test_ComparableNullable_Equal_Invalid_002(t *testing.T) {
	var s1 nullable.String
	var s2 nullable.String
	s1.SetValue("test1")

	if s1.Equal(s2) {
		t.FailNow()
	}
}

func Test_ComparableNullable_Equal_Invalid_003(t *testing.T) {
	var s1 nullable.String
	var s2 nullable.String
	s2.SetValue("test2")

	if s1.Equal(s2) {
		t.FailNow()
	}
}

func Test_NullableBase_SqlNull_Valid(t *testing.T) {
	var value nullable.String
	value.SetValue("test value")

	expected := sql.Null[string]{
		V:     "test value",
		Valid: true,
	}

	actual := value.SqlNull()

	if actual.V != expected.V {
		t.FailNow()
	}

	if actual.Valid != expected.Valid {
		t.FailNow()
	}
}

func Test_NullableBase_SqlNull_Invalid(t *testing.T) {
	var value nullable.String
	value.SetNull()

	expected := sql.Null[string]{
		V:     "",
		Valid: false,
	}

	actual := value.SqlNull()

	if actual.V != expected.V {
		t.FailNow()
	}

	if actual.Valid != expected.Valid {
		t.FailNow()
	}
}

func Test_NullableBase_SetSqlNull_Valid(t *testing.T) {
	var value nullable.String
	value2 := sql.Null[string]{
		V:     "test value",
		Valid: true,
	}

	expected := nullable.NewString("test value")

	value.SetSqlNull(value2)

	if !value.Equal(expected) {
		t.FailNow()
	}
}

func Test_NullableBase_SetSqlNull_Invalid(t *testing.T) {
	var value nullable.String
	value2 := sql.Null[string]{
		V:     "",
		Valid: false,
	}

	expected := nullable.NewNullString()

	value.SetSqlNull(value2)

	if !value.Equal(expected) {
		t.FailNow()
	}
}
