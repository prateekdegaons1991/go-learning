package main

func MyPointer() {
	// Pointers
	// A pointer is a variable that stores the memory address of another variable.
	// Pointers are used to pass variables by reference, allowing functions to modify the original variable.
	// In Go, pointers are created using the & operator and dereferenced using the * operator.

	var x int = 42
	p := &x

	*p = 100 // modify x through the pointer

	println(x) // prints 100
}

// In Go, pointers are not as commonly used as in languages like C or C++ because Go has garbage collection and does not require manual memory management.
// However, pointers are still useful in certain situations, such as when passing large structs to functions or when working with interfaces.
// Pointers are also used in Go to implement linked lists, trees, and other data structures that require dynamic memory allocation.
// Pointers are a powerful feature of Go, but they should be used with care to avoid memory leaks and other issues.
// In Go, the zero value of a pointer is nil, which means it does not point to any memory location.
