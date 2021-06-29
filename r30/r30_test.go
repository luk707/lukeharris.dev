package r30_test

import (
	"testing"

	"lukeharris.dev/r30"
	"lukeharris.dev/testUtils"
)

func TestStep(t *testing.T) {
	utils := testUtils.Setup(t)

	t1 := []byte{0b00000001, 0b01000000}
	t1Expect := []byte{0b00000011, 0b01100000}
	t1Got := r30.Step(t1)
	utils.BytesEq(t1Got, t1Expect)

	t2 := []byte{0b00000001}
	t2Expect := []byte{0b10000011}
	t2Got := r30.Step(t2)
	utils.BytesEq(t2Got, t2Expect)
}
