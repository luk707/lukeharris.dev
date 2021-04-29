# linkedList
--
    import "lukeharris.dev/linkedList"


## Usage

#### func  Push

```go
func Push(head **LinkedList, data interface{})
```
Creates new head node with value for a given linked list

#### type LinkedList

```go
type LinkedList struct {
	Value interface{}
	Next  *LinkedList
}
```

Represents node of a linked list

#### func  New

```go
func New(arr []interface{}) *LinkedList
```
Returns head of linked list containing values of a given array

#### func (*LinkedList) GetTail

```go
func (head *LinkedList) GetTail() *LinkedList
```
Returns tail node of linked list or returns nil if it is cyclical

#### func (*LinkedList) Insert

```go
func (head *LinkedList) Insert(data interface{})
```
Creates a new node after a given node

#### func (*LinkedList) IsCyclical

```go
func (head *LinkedList) IsCyclical() bool
```
Returns true if list is cyclical

#### func (*LinkedList) String

```go
func (head *LinkedList) String() string
```
Formats the linked list as a string

Examples:

- Simple linked list without any cycles

    `(1) -> (2) -> (3)`

- Circular linked list where the tail precedes the head

    `(1) -> (2) -> (3) -> HEAD`

- Circular linked list where the tail precedes the nth element after the head

    `(1) -> (2) -> (3) -> HEAD~n`
