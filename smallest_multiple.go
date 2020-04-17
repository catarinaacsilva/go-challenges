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
	"math"
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

func smallest_multiple_v2() int{
	var k, n, i, limit int = 20, 1, 1, 0
	var check bool = true
	var size_p, size_a int // necessário inicializar
	var p = make([]int,size_p)
	var a = make([]int,size_a)
	
	limit = int(math.Sqrt(float64(k)))
	
	for p[i] <= k{
		a[i] = 1
		if check{
			if p[i] <= limit{
				a[i] = int(math.Floor(math.Log(float64(k))/math.Log(float64(p[i]))))	
			}else{
				check = false
			}
		}
		
		n = n * int(math.Pow(float64(p[i]), float64(a[i])))
		i++
	}
	
	return n
}


func main(){
	//fmt.Println(smallest_multiple_v1())
	fmt.Println(smallest_multiple_v2())
}