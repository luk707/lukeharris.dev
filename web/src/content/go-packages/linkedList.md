# linkedList
--
    import "lukeharris.dev/linkedList"


## Usage

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
