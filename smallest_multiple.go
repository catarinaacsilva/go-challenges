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

// TODO: problem with n
func smallest_multiple() int{
	var div, result, n int = 1,0,1

	result = n%div

	
	for n%div == 0 && div >= 1 && div <= 10 {
		if n%div < result{
			result = n%div
		}
		div ++
	}
	
	return result
}


func main(){
	fmt.Println(smallest_multiple())
}