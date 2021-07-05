package testUtils /* import "go.lukeharris.dev/testUtils" */

import (
	"bytes"
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

// Asserts if bytes are equal
func (u *TestUtils) BytesEq(got, expected []byte) {
	if !bytes.Equal(got, expected) {
		u.t.Errorf("Got %b, expected %b", got, expected)
	}
}
