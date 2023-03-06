package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello Go")
	fmt.Println(runtime.GOARCH)
}
