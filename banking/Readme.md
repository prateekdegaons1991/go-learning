# Banking Project

## Overview
This project is a simple banking application developed to demonstrate basic banking operations such as account creation, deposits, withdrawals, and balance inquiries.

## Features
- Create a new bank account
- Deposit money into an account
- Withdraw money from an account
- Check account balance

## Technologies Used
- Go (Golang)
- Standard Library

## Installation
1. Clone the repository:
    ```sh
    git clone https://github.com/prateekdegaons1991/banking.git
    ```
2. Navigate to the project directory:
    ```sh
    cd banking
    ```
3. Build the project:
    ```sh
    go build
    ```

## Usage
1. Run the application:
    ```sh
    go run bank.go
    ```


    ## Important Go Concepts

    - **Packages and Imports**: The project demonstrates how to organize code into packages and import them as needed.
    - **Structs**: Used to define the structure of a bank account.
    - **Methods**: Attached to structs to perform operations like deposits and withdrawals.
    - **Error Handling**: Proper error handling to manage invalid operations.
    - **Concurrency**: Use of goroutines and channels to handle concurrent transactions.
    - **Interfaces**: Define behaviors that different types can implement.
    - **Unit Testing**: Writing tests to ensure the correctness of the banking operations using Go's testing package.
    - **Standard Library**: Utilization of Go's standard library for various functionalities.


## What are pointer in Go?
Pointers in Go are variables that store the memory address of another variable. They are used to indirectly access and modify the value of the variable they point to. Pointers are particularly useful for:

- **Efficiency**: Passing large structs or arrays by pointer can avoid the overhead of copying data.
- **Mutability**: Functions can modify the value of a variable passed by pointer.
- **Shared State**: Multiple functions can operate on the same data by sharing pointers.

### Syntax
To declare a pointer, use the `*` operator:
```go
var ptr *int
```

To get the address of a variable, use the `&` operator:
```go
var x int = 10
ptr = &x
```

To access the value stored at the pointer, use the `*` operator (dereferencing):
```go
fmt.Println(*ptr) // Output: 10
```

### Example
```go
package main

import "fmt"

func main() {
    var a int = 58
    var ptr *int = &a

    fmt.Println("Value of a:", a)
    fmt.Println("Address of a:", &a)
    fmt.Println("Value of ptr:", ptr)
    fmt.Println("Value pointed to by ptr:", *ptr)

    *ptr = 100
    fmt.Println("New value of a:", a)
}
```
In this example, `ptr` is a pointer to `a`. Changing the value at `ptr` also changes the value of `a`.