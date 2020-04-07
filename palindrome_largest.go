package main

import "fmt"

func palindrome(n int)bool{
	var r, sum, tmp int = 0,0,0

	tmp = n
	for n>0{
		r = n%10
		sum = (sum*10)+r
		n = n/10
	}
	if tmp == sum{
		return true
	}
	return false
}

func palindrome_largest(a, b int)int{
	var p, result int = 0,0

	for a<1000{

		for b<1000{
			result = a * b
			
			if palindrome(result) == true && result > p{
				p = result
			}
			b++
		}
		a++	
		b = 1
	}
	return p
}

func main(){
	var palindrom_largest = palindrome_largest(100, 100)
	fmt.Println(palindrom_largest)
}
 