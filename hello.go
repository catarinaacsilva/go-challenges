package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Printf("hello, world\n")
	result := add(2, 3)
	fmt.Println(result)
}