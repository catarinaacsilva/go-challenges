/*
* Smallest Multiple
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

func multiple(value int) int{
	var result int = 0

	for i:=1; i<=10; i++{
		if value%i >= 0{
			result = value
		} else{
			break
		}
	}
	return result
}

func smallest_multiple() int{
	var value, result_mult int = 1, 0

	for multiple(value) != 0{
		result_mult = multiple(value)
		value++
	}

	return result_mult
	

}


func main(){
	fmt.Println(smallest_multiple())
}