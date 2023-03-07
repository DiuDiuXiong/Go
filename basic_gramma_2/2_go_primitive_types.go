package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/**
* 1. bool
* 2. string
* 3. (u)int, (u)int8, (u)int16, (u)int32, (u)int64
* 4. uintptr: pointer
* 5. byte
* 6. rune: char for golang, 32bit for multiple characters, which compatible with utf-8
* 7. float32, float64, complex64, complex128 | complex number (has i component, i = sqrt(-1))
 */

/**
* Golang only has mandatory format conversion, 强制类型转换
 */

func euler() {
	var c complex128 = 3 + 4i
	fmt.Println(cmplx.Abs(c)) // get the mod 3^2 + 4^2 = 5^2

	// euler formula, e^(i*pi) + 1 = 0 （see euler.jpg）
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)

	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	a, b := 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b))) // we have to explicit convert a*a + b*b to float64 2. we have to explicitly call int(...)
	// has one issue, since float is not accurate, when a,b large enough, may get 4.999999, then int(...) will make it a 4
	fmt.Println(c)
}

/*
func main() {
	euler()
	triangle()
}

*/
