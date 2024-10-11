package nullable_test

import (
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

func Test_ComparableNullable_Equals_Valid_001(t *testing.T) {
	var s1 nullable.String
	var s2 nullable.String
	s1.SetValue("test")
	s2.SetValue("test")

	if !s1.Equals(s2) {
		t.FailNow()
	}
}

func Test_ComparableNullable_Equals_Valid_002(t *testing.T) {
	var s1 nullable.String
	var s2 nullable.String

	if !s1.Equals(s2) {
		t.FailNow()
	}
}

func Test_ComparableNullable_Equals_Invalid_001(t *testing.T) {
	var s1 nullable.String
	var s2 nullable.String
	s1.SetValue("test1")
	s2.SetValue("test2")

	if s1.Equals(s2) {
		t.FailNow()
	}
}

func Test_ComparableNullable_Equals_Invalid_002(t *testing.T) {
	var s1 nullable.String
	var s2 nullable.String
	s1.SetValue("test1")

	if s1.Equals(s2) {
		t.FailNow()
	}
}

func Test_ComparableNullable_Equals_Invalid_003(t *testing.T) {
	var s1 nullable.String
	var s2 nullable.String
	s2.SetValue("test2")

	if s1.Equals(s2) {
		t.FailNow()
	}
}
