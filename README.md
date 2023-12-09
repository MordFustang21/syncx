# syncx
A generic wrapper around `sync.Map`. This package is made to avoid having to do
typecasts when retrieving values out of the `sync.Map`

Example
```go
type MyStruct struct{}

m := syncx.NewMap[comparable, MyStruct]()

out, found := m.Load("key")
// out is of type MyStruct not any
```
