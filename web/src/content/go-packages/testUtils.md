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


#### func (*TestUtils) StrEq

```go
func (u *TestUtils) StrEq(got, expected string)
```
Asserts if strings are equal
