# testUtils
--
    import "lukeharris.dev/testUtils"


## Usage

#### type TestUtils

```go
type TestUtils struct {
	T *testing.T
}
```


#### func (*TestUtils) BoolEq

```go
func (u *TestUtils) BoolEq(got, expected bool)
```
Asserts if booleans are equal

#### func (*TestUtils) IsNil

```go
func (u *TestUtils) IsNil(got interface{})
```
Asserts if value is nil

#### func (*TestUtils) IsNotNil

```go
func (u *TestUtils) IsNotNil(got interface{})
```
Asserts if value is non-nil

#### func (*TestUtils) StrEq

```go
func (u *TestUtils) StrEq(got, expected string)
```
Asserts if strings are equal
