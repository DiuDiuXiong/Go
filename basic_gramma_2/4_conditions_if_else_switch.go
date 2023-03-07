package main

import (
	"fmt"
	"os"
)

func if_else() {
	const filename = "abc.txt"
	contents, err := os.ReadFile(filename)
	// if ... {} else {}
	if err != nil {
		fmt.Println(err) // con such file or directory
	} else {
		fmt.Printf("%s\n", contents)
	}

	// if error is not empty, then something wrong happened
	if contents2, err2 := os.ReadFile(filename); err != nil {
		fmt.Println(err2)
	} else {
		fmt.Printf("%s\n", contents2)
	}
	//  contents2, err2 are out of scope here
}

/*
*
 1. in go, switch do not need break
 2. panic will raise the error
 3. this also work:
    switch <string>:
    case <str1>:
    ...
    case <str2>:
    ...
    case <...>
    default:
    panic(wrong message)
 4. use fallthrough to stop break and continue on next case
 5. switch can be followed by expression or no expression at all
*/
func grade(score int) string {
	g := ""
	switch {
	case score < 60 && score >= 0:
		g = "F"
	case score < 70:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		panic(fmt.Sprintf("Wrong score: %d", score))
	}
	return g
}

/*
func main() {
	//if_else()
	fmt.Println(grade(0), grade(59), grade(60), grade(100), grade(101))
	//fmt.Println(grade(0), grade(59), grade(60), grade(100), grade(89))
}
*/
