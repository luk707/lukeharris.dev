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

/*
Formats the linked list as a string

Examples:

- Simple linked list without any cycles
  `1 -> 2 -> 3`


- Circular linked list where the tail precedes the head
  `1 -> 2 -> 3 -> HEAD`

- Circular linked list where the tail precedes the nth element after the head
  `1 -> 2 -> 3 -> HEAD~n`

*/
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
	c1 := head
	c2 := head
	for c1.Next != nil {
		c1 = c1.Next
		if c1.Next != nil {
			c1 = c1.Next
			c2 = c2.Next
		}
		if c1 == c2 {
			return nil
		}
	}
	return c1
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
