package main

import (
	"errors"
	"fmt"
)

func tryRecover() {
	defer func() {
		// defer must be a function call
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(r)
		}

	}()
	panic(errors.New("This is an error. ")) // this one is fine, it is panic error
	//panic(123) // this will go to else branch, since panic can pass in any type of parameter

}

func main() {
	tryRecover()
}
