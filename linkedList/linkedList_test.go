package linkedList_test

import (
	"testing"

	"go.lukeharris.dev/linkedList"
	"go.lukeharris.dev/testUtils"
)

func MakeList(length int) *linkedList.Node {
	var head *linkedList.Node
	for i := length; i > 0; i-- {
		head = &linkedList.Node{i, head}
	}
	return head
}

func TestString(t *testing.T) {
	utils := testUtils.Setup(t)

	t1 := MakeList(3)
	t1Expect := "(1) -> (2) -> (3)"
	t1Got := t1.String()
	utils.StrEq(t1Got, t1Expect)

	t2 := MakeList(3)
	t2.GetTail().Next = t2
	t2Got := t2.String()
	t2Expect := "(1) -> (2) -> (3) -> HEAD"
	utils.StrEq(t2Got, t2Expect)

	t3 := MakeList(3)
	t3.GetTail().Next = t3.Next
	t3Got := t3.String()
	t3Expect := "(1) -> (2) -> (3) -> HEAD~1"
	utils.StrEq(t3Got, t3Expect)
}

func TestGetTail(t *testing.T) {
	utils := testUtils.Setup(t)

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

func TestIsCyclical(t *testing.T) {
	utils := testUtils.Setup(t)

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

func TestNew(t *testing.T) {
	utils := testUtils.Setup(t)

	t1 := linkedList.New([]interface{}{1, 2, 3})
	t1Expect := "(1) -> (2) -> (3)"
	t1Got := t1.String()
	utils.StrEq(t1Got, t1Expect)

	t2 := linkedList.New([]interface{}{"a", "b", "c"})
	t2Expect := "(a) -> (b) -> (c)"
	t2Got := t2.String()
	utils.StrEq(t2Got, t2Expect)

	t3 := linkedList.New([]interface{}{t1, t2})
	t3Expect := "((1) -> (2) -> (3)) -> ((a) -> (b) -> (c))"
	t3Got := t3.String()
	utils.StrEq(t3Got, t3Expect)
}

func TestPush(t *testing.T) {
	utils := testUtils.Setup(t)

	t1 := linkedList.New([]interface{}{2, 3})
	linkedList.Push(&t1, 1)
	t1Expect := "(1) -> (2) -> (3)"
	t1Got := t1.String()
	utils.StrEq(t1Got, t1Expect)

	t2 := linkedList.New([]interface{}{2, 3, 4})
	t2.GetTail().Next = t2
	linkedList.Push(&t2, 1)
	t2Expect := "(1) -> (2) -> (3) -> (4) -> HEAD~1"
	t2Got := t2.String()
	utils.StrEq(t2Got, t2Expect)
}

func TestInsert(t *testing.T) {
	utils := testUtils.Setup(t)

	t1 := linkedList.New([]interface{}{1, 3})
	t1.Insert(2)
	t1Expect := "(1) -> (2) -> (3)"
	t1Got := t1.String()
	utils.StrEq(t1Got, t1Expect)

	t2 := linkedList.New([]interface{}{1, 3, 4})
	t2.GetTail().Next = t2
	t2.Insert(2)
	t2Expect := "(1) -> (2) -> (3) -> (4) -> HEAD"
	t2Got := t2.String()
	utils.StrEq(t2Got, t2Expect)

	t3 := linkedList.New([]interface{}{1, 3, 4})
	t3.GetTail().Next = t3.Next
	t3.Insert(2)
	t3Expect := "(1) -> (2) -> (3) -> (4) -> HEAD~2"
	t3Got := t3.String()
	utils.StrEq(t3Got, t3Expect)
}
