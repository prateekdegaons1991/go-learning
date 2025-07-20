# Go Generics Examples

This repository contains examples demonstrating the use of generics in Go.

## What are Generics?

Generics allow you to write flexible and reusable code by enabling functions, types, and data structures to operate on different data types while maintaining type safety.

## Key Concepts

- **Type Parameters:** Functions and types can accept type parameters, specified in square brackets.
- **Constraints:** You can restrict type parameters to certain types or interfaces.
- **Type Inference:** Go can often infer the type parameter from usage.

## Example

```go
// Generic function to find the maximum of two values
func Max[T comparable](a, b T) T {
    if a > b {
        return a
    }
    return b
}

func main() {
    fmt.Println(Max(3, 7))         // Output: 7
    fmt.Println(Max("go", "lang")) // Output: "lang"
}
```

**Explanation:**

- `Max` is a generic function with a type parameter `T` constrained by the `comparable` interface.
- It compares two values of type `T` and returns the greater one.
- The function works with different types (e.g., `int`, `string`) while ensuring type safety.
- Type inference allows you to call `Max` without explicitly specifying the type parameter.