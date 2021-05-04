# linkedList
--
    import "lukeharris.dev/linkedList"


## Usage

#### func  Push

```go
func Push(head **Node, data interface{})
```
Creates new head node with value for a given linked list

#### type Node

```go
type Node struct {
	Data interface{}
	Next *Node
}
```

Represents node of a linked list

#### func  New

```go
func New(arr []interface{}) *Node
```
Returns head of linked list containing values of a given array

#### func (*Node) GetTail

```go
func (head *Node) GetTail() *Node
```
Returns tail node of linked list or returns nil if it is cyclical

#### func (*Node) Insert

```go
func (head *Node) Insert(data interface{})
```
Creates a new node after a given node

#### func (*Node) IsCyclical

```go
func (head *Node) IsCyclical() bool
```
Returns true if list is cyclical

#### func (*Node) String

```go
func (head *Node) String() string
```
Formats the linked list as a string

Examples:

- Simple linked list without any cycles

    `(1) -> (2) -> (3)`

- Circular linked list where the tail precedes the head

    `(1) -> (2) -> (3) -> HEAD`

- Circular linked list where the tail precedes the nth element after the head

    `(1) -> (2) -> (3) -> HEAD~n`
