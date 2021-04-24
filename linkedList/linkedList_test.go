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

	t3 := MakeList(3)
	t3.GetTail().Next = t3.Next
	t3Got := t3.String()
	t3Expect := "1 -> 2 -> 3 -> HEAD~1"
	utils.StrEq(t3Got, t3Expect)
}

func TestIsCyclical(t *testing.T) {
	utils := &testUtils.TestUtils{T: t}

	t1 := MakeList(3)
	t1.GetTail().Next = t1
	t1Expect := true
	t1Got := t1.IsCyclical()
	utils.BoolEq(t1Got, t1Expect)

	t2 := MakeList(3)
	t2.GetTail().Next = t1.Next
	t2Expect := true
	t2Got := t2.IsCyclical()
	utils.BoolEq(t2Got, t2Expect)

	t3 := MakeList(3)
	t3Expect := false
	t3Got := t3.IsCyclical()
	utils.BoolEq(t3Got, t3Expect)
}

func TestGetTail(t *testing.T) {
	utils := &testUtils.TestUtils{T: t}

	t1 := MakeList(3)
	t1.GetTail().Next = t1
	t1Got := t1.GetTail()
	utils.IsNil(t1Got)

	t2 := MakeList(3)
	t2.GetTail().Next = t1.Next
	t2Got := t2.GetTail()
	utils.IsNil(t2Got)

	t3 := MakeList(3)
	t3Got := t3.GetTail()
	utils.IsNotNil(t3Got)
}
