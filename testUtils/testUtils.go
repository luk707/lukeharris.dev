package testUtils /* import "lukeharris.dev/testUtils" */

import "testing"

type TestUtils struct {
	T *testing.T
}

// Asserts if strings are equal
func (u *TestUtils) StrEq(got, expected string) {
	if got != expected {
		u.T.Errorf("Got \"%s\", expected \"%s\"", got, expected)
	}
}
