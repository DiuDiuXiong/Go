package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/**
Here, few key points:
1. Use closure to implement fibonacci
2. Functions can implement interfaces, just pass it as receiver
3. An updated version to get rid of the issue that p []byte is too small to read whole integer, need to change intGen to a struct check 7-2 for example
- save the reader
- if err != io.EOF, just return
- if err == io.EOF, generate next number, number -> str, intGen.currentReader = string.NewReader(str)
*/

type intGen func() int

func fibonacci() intGen { // everytime call this function, generate new one
	x1, x2 := 0, 1

	return func() int {
		x1, x2 = x2, x1+x2

		return x1
	}

}

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	// TODO: incorrect if p is too small!
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

/*
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	printFileContents(f)
}
*/
