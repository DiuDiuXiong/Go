package main

import (
	"bufio"
	"fmt"
	"os"
)

func tryDefer() { // defer is stack like
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	return
	//panic("Error occured")
	fmt.Println(4)
	defer fmt.Println(5)
}

func fibonacci() func() int { // everytime call this function, generate new one
	x1, x2 := 0, 1

	return func() int {
		x1, x2 = x2, x1+x2

		return x1
	}

}

func forDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func forDefer1() {
	var i = 0
	ptrI := &i
	for ; *ptrI < 100; *ptrI++ {
		defer printF(ptrI)
		if *ptrI == 30 {
			panic("printed too many")
		}
	}
}

func printF(prtI *int) {
	fmt.Println(*prtI)
}

func writeFile(filename string) { // here will first flush content then close file
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file) // write to buffer to allow writing to file at once to save time
	defer writer.Flush()            // write to file
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f()) // write to buffer
	}
}

func writeFileError(filename string) { // here will first flush content then close file

	file, err := os.OpenFile(
		filename, os.O_EXCL|os.O_CREATE, 0666) // here set flag O_EXCL, will not be able to open if file already exist
	defer file.Close()
	//err = errors.New("this is a custom error")
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok { // as OpenFile documentation said: If there is an error, it will be of type *PathError.
			// know it is an error interface, but use this line to check if it is of type we expected.
			panic(err) // something unexpected happened
		} else {
			fmt.Println(pathError.Op, "\n", pathError.Path, "\n", pathError.Err)
		}
		return
	}

	writer := bufio.NewWriter(file) // write to buffer to allow writing to file at once to save time
	defer writer.Flush()            // write to file
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f()) // write to buffer
	}
}

func main() {
	// tryDefer()
	writeFile("./8_error_handling/01_defer/fib.txt")

	//forDefer()
	//forDefer1()
	writeFileError("./8_error_handling/01_defer/fib.txt")
}
