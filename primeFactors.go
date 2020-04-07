/*
* Prime Factors
*
* Author: Catarina Silva
* Email: c.alexandracorreia@av.it.pt
* Email: c.alexandracorreia@ua.pt
* 
* Version: 1.0
*/

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