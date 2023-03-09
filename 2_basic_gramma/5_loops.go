package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
*
* 1. for/while same like if statement, condition check cannot have brackets
* 2. initial condition; end condition; update condition all can be changed
* 3. there is no: <while (condition) {...}>, instead use <go (condition) {...}>
* 4.
  - sumOneToHundred(): have all
  - convertToBin(n int) string: ignore initial condition
  - printFile(filename string): only exit condition, like while
  - forever(): ignore all stuff, infinite loop
*/
func sumOneToOneHundred() int {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return sum
}

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { //  this is while loop, the only thing here is ending condition, function is returning bool
		fmt.Println(scanner.Text()) // scanner.Text() is read line
	}
}

func forever() {
	for {
		fmt.Println("abc.txt")
	}
}

func main() {
	sumOneToOneHundred()
	fmt.Println(convertToBin(5),
		convertToBin(13),
		convertToBin(389242123193818938),
		convertToBin(0))
	printFile("abc.txt")
	// forever() infinite loop
}
