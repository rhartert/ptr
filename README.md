# Go Pointer

[![Go Reference](https://pkg.go.dev/badge/github.com/rhartert/ptr.svg)](https://pkg.go.dev/github.com/rhartert/ptr)
[![Go Report Card](https://goreportcard.com/badge/github.com/rhartert/ptr)](https://goreportcard.com/report/github.com/rhartert/ptr)

This package offers utility functions to reduce the boilerplate of working with
pointers of primitive types. 

Package `ptr` is meant as replacement for the type-specific functions (e.g.
`func String(v string) *string`) that are ubiquitous in many Go packages.

```go
type Response struct {
    IntValue    *int
    BoolValue   *bool
    StringValue *string
}

func respond() Response {
    return Response{
        IntValue:    ptr.Of(42),
        BoolValue:   ptr.Of(true),
        StringValue: ptr.Of("foo"),
    }
}
```

The `ptr.Value` function provides a convenient way to deal with `nil`
pointers by returning the zero value of the pointer's type

```go
var ptr *int
fmt.Println(Value(ptr)) // Output: 0
```

