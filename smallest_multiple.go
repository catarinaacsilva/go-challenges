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
	//"math"
)

// first version
func smallest_multiple_v1() int{
	var cond bool = false
	var i int = 0

	for cond == false{
		i++
		if i%2 == 0 && i%3 == 0 && i%4 == 0 && i%5 == 0 && i%6 == 0 && i%7 == 0 && i%8 == 0 && i%9 == 0 && i%10 == 0 && i%11 == 0 && i%12==0 && i%13==0 && i%14==0 && i%14 == 0 && i%15==0 && i%16==0 && i%17 == 0 && i%18 == 0 && i%19 == 0 && i%20==0{
			cond = true
		}
	}

	return i

}

// Second version

func aux_smallest_multiple_v2(a,b int) int{
	var i, sm int = 1, a*b
		
	
	for true{
		sm = a * i
		i++
		if sm%b == 0{
			break
		}
	}
    
    return sm;
}

func smallest_multiple_v2() int{
	var sm int = 2
	for i:=1; i<=20; i++{
		sm = aux_smallest_multiple_v2(i, sm)
	}
	return sm
}

func main(){
	fmt.Println("Result with first approach:")
	fmt.Println(smallest_multiple_v1())
	fmt.Println("Result with second approach:")
	fmt.Println(smallest_multiple_v2())
}