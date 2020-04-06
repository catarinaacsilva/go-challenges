package main

// import "fmt"

import(
	"fmt"
	"math/cmplx"
	"math"
)

var c, python, java bool

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)


func add(x int, y int) int {
	return x + y
}

/*
 * Split value
 * Input: int that represents the sum
 * Output: x and y (both int)
 * Note: (x int, y int) = (x,y int)
 */
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	//fmt.Println(x)
	//fmt.Println(y)
	return
}

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Printf("hello, world\n")

	result := add(2, 3)
	fmt.Println(result)

	fmt.Println(split(17))

	var i int
	python = true // just test
	fmt.Println(i, c, python, java)

	var j, k int = 1, 2
	fmt.Printf("j=%d \nk=%d\n", j, k)

	var l, m int = 1, 2
	n := 3
	c, python, java := true, false, "no!"
	fmt.Println(l, m, n, c, python, java)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	v := 42
	p := 42.2
	fmt.Printf("v is of type %T and n is of type %T\n", v, p)

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}