package linkedList_test

import (
	"testing"

	"lukeharris.dev/linkedList"
	"lukeharris.dev/testUtils"
)

func MakeList(length int) *linkedList.LinkedList {
	var head *linkedList.LinkedList
	for i := length; i > 0; i-- {
		head = &linkedList.LinkedList{i, head}
	}
	return head
}

func TestString(t *testing.T) {
	utils := &testUtils.TestUtils{T: t}

	t1 := MakeList(3)
	t1Expect := "1 -> 2 -> 3"
	t1Got := t1.String()
	utils.StrEq(t1Got, t1Expect)

	t2 := MakeList(3)
	t2.GetTail().Next = t2
	t2Got := t2.String()
	t2Expect := "1 -> 2 -> 3 -> HEAD"
	utils.StrEq(t2Got, t2Expect)
}
