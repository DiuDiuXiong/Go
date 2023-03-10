package main

import "fmt"

var (
	aa = 3 // these are only module variable, not global variable
	ss = "lll"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d, %q\n", a, s) // a => 0, s=> empty string
}

func variableInitialValue() {
	var a, b int = 3, 4 // all variables those are defined are required to be use
	var s string = "abc"
	fmt.Println(a, s, b)
}

func variableTypeDeduction() {
	var a, b, c, d = 3, 4, true, "abc" // don't need to specify the type
	fmt.Println(a, b, c, d)
}

func variableShorter() {
	a, b, c, d := 3, 4, true, "abc" // ... := ... is same as var ... = ...
	b = 5                           // b := 5 is false since this is re-declare type
	fmt.Println(a, b, c, d)
}

/*
func main() {
	fmt.Println("Hello world!")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)
}*/
