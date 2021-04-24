package linkedList /* import "lukeharris.dev/linkedList" */

import (
	"fmt"
	"strings"
)

type void struct{}

var member void

// Represents node of a linked list
type LinkedList struct {
	Value interface{}
	Next  *LinkedList
}

// Formats the linked list as a string
func (head *LinkedList) String() string {
	var sb strings.Builder
	seen := make(map[*LinkedList]int)
	current := head
	depth := 0
	for current.Next != nil {
		sb.WriteString(fmt.Sprintf("(%v) -> ", current.Value))
		if d, ok := seen[current.Next]; ok {
			if d == 0 {
				sb.WriteString("HEAD")
				return sb.String()
			}
			sb.WriteString(fmt.Sprintf("HEAD~%d", d))
			return sb.String()
		}
		seen[current] = depth
		current = current.Next
		depth++
	}
	sb.WriteString(fmt.Sprintf("(%v)", current.Value))
	return sb.String()
}

// Returns tail node of linked list or returns nil if it is cyclical
func (head *LinkedList) GetTail() *LinkedList {
	seen := make(map[*LinkedList]void)
	current := head
	for current.Next != nil {
		if _, ok := seen[current.Next]; ok {
			return nil
		}
		seen[current] = member
		current = current.Next
	}
	return current
}

// Returns true if list is cyclical
func (head *LinkedList) IsCyclical() bool {
	return head.GetTail() == nil
}

// Returns head of linked list containing values of a given array
func New(arr []interface{}) *LinkedList {
	var head *LinkedList

	for i := len(arr) - 1; i >= 0; i-- {
		head = &LinkedList{arr[i], head}
	}
	return head
}
