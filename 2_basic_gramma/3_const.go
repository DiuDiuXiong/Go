package main

import (
	"fmt"
	"math"
)

func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
		// x error, const must have initial value
	)
	var c int
	c = int(math.Sqrt(a*a + b*b)) // const can do type conversion automatically
	// a = 2 error, cannot change const
	fmt.Println(filename, c)
}

// for enums, use iota for auto increasing
func enums() {
	const (
		cpp = iota
		_
		java
		python
		golang
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

/*
func main() {
	consts()
	enums()
}*/
