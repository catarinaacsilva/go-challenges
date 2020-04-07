/*
* Multiple of 3 and 5
*
* Author: Catarina Silva
* Email: c.alexandracorreia@av.it.pt
* Email: c.alexandracorreia@ua.pt
* 
* Version: 1.0
*/

package main

import "fmt"

func mult3(num int) bool{
	if num%3 == 0{
		return true
	}
	return false
}

func mult5(num int) bool{
	if num%5 == 0{
		return true
	}
	return false
}

func main(){
	var sum int
	for i:=0; i<1000; i++ {
		if mult3(i) == true{
			sum += + i
		} else if mult5(i) == true {
			sum += i
		} else{
			sum = sum
		}
	}
	fmt.Println(sum)
}