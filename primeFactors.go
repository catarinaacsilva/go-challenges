package main

import(
	"fmt"
)

func primeFactor(n int){
	for i := 2; i< n; i++{
		for n%i == 0{
			n = n/i
		}
	}
	if (n > 2){
		fmt.Println(n)
	}
}

func main(){
	primeFactor(600851475143)
}