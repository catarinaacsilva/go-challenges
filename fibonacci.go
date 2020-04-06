package main

import "fmt"

func main(){
	var p, c, r int = 1,2,2

	for c < 4000000{
		t := c
		c = c + p
		p = t
		if c%2 == 0{
			r += c
		}
	}
	fmt.Printf("Result: %d\n", r)
}