# Go Utils

**Features of Go Programming:**

- Support for environment adopting patterns similar to dynamic languages. For example, type inference (x := 0 is valid declaration of a variable x of type int)

- Compilation time is fast.

- Inbuilt concurrency support: lightweight processes (via go routines), channels, select statement.

- Go programs are simple, concise, and safe.

- Support for Interfaces and Type embedding.

- Production of statically linked native binaries without external dependencies.

**Comments**

/* my first program in Go */


## Requirements

- Linux 20.04
- Golang

## Installation

1. `sudo apt install golang-go`

## Run

1. `go build hello.go`

2. `./hello`

## Notes:

1. To create variable: `x int`

2. To assign result of function variable: `x := function()`

3. The `var` statement declares a list of variables

4. Inside a function, the `:=` short assignment statement can be used in place of a `var` declaration with implicit type. 

5. When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side. This is on 9 case at main of hello.go.

6. Constants are declared like variables, but with the `const` keyword.

7. Constants can be character, string, boolean, or numeric values.

8. Constants cannot be declared using the `:=` syntax. 

## Authors

* **Catarina Silva** - [catarinaacsilva](https://github.com/catarinaacsilva)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details