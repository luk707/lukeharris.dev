package testUtils /* import "lukeharris.dev/testUtils" */

import (
	"reflect"
	"testing"
)

type TestUtils struct {
	t *testing.T
}

// Returns pointer to new TestUtils instance
func Setup(t *testing.T) *TestUtils {
	return &TestUtils{t}
}

// Asserts if strings are equal
func (u *TestUtils) StrEq(got, expected string) {
	if got != expected {
		u.t.Errorf("Got \"%s\", expected \"%s\"", got, expected)
	}
}

// Asserts if booleans are equal
func (u *TestUtils) BoolEq(got, expected bool) {
	if got != expected {
		u.t.Errorf("Got %t, expected %t", got, expected)
	}
}

// Asserts if value is nil
func (u *TestUtils) IsNil(got interface{}) {
	if !reflect.ValueOf(got).IsNil() {
		u.t.Errorf("Got \"%v\", expected nil", got)
	}
}

// Asserts if value is non-nil
func (u *TestUtils) IsNotNil(got interface{}) {
	if reflect.ValueOf(got).IsNil() {
		u.t.Error("Got nil, expected non-nil value", got)
	}
}
