package main

import (
	"fmt"
	"math"
)

func CalcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func LengthOfNonRepeatingSubStr(s string) int {
	lastOccured := make(map[rune]int)

	//lastOccured := make([]int, 0xffff) //65535, 65k for storage
	//for i := range lastOccured {
	//	lastOccured[i] = -1
	//}
	start := 0

	maxLength := 0
	for i, ch := range []rune(s) { // []rune(s) so all transferred to type that can consume many characters
		// avoid default 0 value, check exist
		if lastI, ok := lastOccured[ch]; ok && lastI >= start {
			//if lastI := lastOccured[ch]; lastI != -1 && lastI >= start {
			start = lastOccured[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccured[ch] = i

	}
	return maxLength
}

func notCovered() {
	fmt.Println("This function is not covered")
}

func main() {

}
