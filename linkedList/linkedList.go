package linkedList /* import "lukeharris.dev/linkedList" */

import (
	"fmt"
	"strings"
)

type LinkedList struct {
	Value interface{}
	Next  *LinkedList
}

func (head *LinkedList) String() string {
	var sb strings.Builder
	current := head
	for current.Next != nil {
		sb.WriteString(fmt.Sprintf("%v -> ", current.Value))
		current = current.Next
		if current == head {
			sb.WriteString("HEAD")
			return sb.String()
		}
	}
	sb.WriteString(fmt.Sprintf("%v", current.Value))
	return sb.String()
}

func (head *LinkedList) GetTail() *LinkedList {
	current := head
	for current.Next != nil {
		if current.Next == head {
			return nil
		}
		current = current.Next
	}
	return current
}

func (head *LinkedList) IsCyclical() bool {
	return head.GetTail() == nil
}

func New(arr []interface{}) *LinkedList {
	var head *LinkedList

	for i := len(arr) - 1; i >= 0; i-- {
		head = &LinkedList{arr[i], head}
	}
	return head
}
