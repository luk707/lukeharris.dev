package testUtils /* import "lukeharris.dev/testUtils" */

import (
	"reflect"
	"testing"
)

type TestUtils struct {
	T *testing.T
}

// Asserts if strings are equal
func (u *TestUtils) StrEq(got, expected string) {
	if got != expected {
		u.T.Errorf("Got \"%s\", expected \"%s\"", got, expected)
	}
}

// Asserts if booleans are equal
func (u *TestUtils) BoolEq(got, expected bool) {
	if got != expected {
		u.T.Errorf("Got %t, expected %t", got, expected)
	}
}

// Asserts if value is nil
func (u *TestUtils) IsNil(got interface{}) {
	if !reflect.ValueOf(got).IsNil() {
		u.T.Errorf("Got \"%v\", expected nil", got)
	}
}

// Asserts if value is non-nil
func (u *TestUtils) IsNotNil(got interface{}) {
	if reflect.ValueOf(got).IsNil() {
		u.T.Error("Got nil, expected non-nil value", got)
	}
}
