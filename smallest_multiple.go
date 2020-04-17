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

func smallest_multiple() int{
	var cond bool = false
	var i int = 0

	for cond == false{
		i++
		if i%2 == 0{
			cond = true
		}
	}

	return i

}


func main(){
	fmt.Println(smallest_multiple())
}