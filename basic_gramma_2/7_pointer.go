package main

import "fmt"

/**
* pretty much like C
* 1. Pointer in golang cannot do calculation, can only assign/re-assign
* 2. Golang is pass by value!! Only pass by value. To keep efficiency and avoid copy params, use pointer
* 3.
- *<type>: means pointer to that type
- &<param>: means address of param
- *<ptr>: means content ptr is pointing to
* 4. for passing 'Object' to functions, 'Object' should contain only pointers to avoid copy efficiency waste
*/

func tryPointer() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a, *pa)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	tryPointer()
	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)
}
