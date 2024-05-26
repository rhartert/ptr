// Package ptr provides utilities to reduce the boilerplate of working with
// pointers of primitive types.
package ptr

// Of returns a pointer to the given value.
func Of[T any](v T) *T {
	return &v
}

// Value returns the value referenced by the given pointer. If the pointer is
// nil, it returns the zero value of the pointer's type.
func Value[T any](ptr *T) T {
	if ptr == nil {
		ptr = new(T) // zero value
	}
	return *ptr
}
