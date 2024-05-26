package ptr

import "fmt"

func ExampleOf_int() {
	p := Of(42)
	fmt.Println(*p) // Output: 42
}

func ExampleOf_bool() {
	p := Of(true)
	fmt.Println(*p) // Output: true
}

func ExampleOf_string() {
	p := Of("foobar")
	fmt.Println(*p) // Output: foobar
}

func ExampleOf_slice() {
	p := Of([]int{42})
	fmt.Println(*p) // Output: [42]
}

func ExampleOf_struct() {
	type s = struct {
		Foo bool
		Bar int
	}

	p := Of(s{true, 42})
	fmt.Println(*p) // Output: {true 42}
}

func ExampleValue_int() {
	var ptr *int
	fmt.Println(Value(ptr)) // Output: 0
}

func ExampleValue_bool() {
	var ptr *bool
	fmt.Println(Value(ptr)) // Output: false
}

func ExampleValue_string() {
	var ptr *string
	fmt.Printf("%q", Value(ptr)) // Output: ""
}

func ExampleValue_slice() {
	var ptr *[]int
	fmt.Println(Value(ptr)) // Output: []
}

func ExampleValue_struct() {
	type s = struct {
		Foo bool
		Bar int
	}

	var ptr *s
	fmt.Println(Value(ptr)) // Output: {false 0}
}
