package main

import "fmt"

/*
*
1. To define, see methods below
  - to define, [length]<type>
  - for :=, initial value must be given
  - use ... to let compiler count
  - for i := range arr will give idx of arr
  - for i, v := range arr will iterate over idx, val_idx

2. array is value type!! not like other language that array is reference type.
	So [10]int and [20]int are different data type, func f(arr [10]int) will copy the array.
	And if you pass by pointer you will be able to change array.
	But this is still not easy to use, so Golang, we usually use slice.
*/

func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
	arr[0] = 100 // do not change over scope, since array is value type
}

func changeArray(arr *[5]int) { // here *(arr)[0] has same effect as arr[0]
	arr[0] = -1
	arr[1] = -2
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	printArray(arr1)
	printArray(arr3) // if length not 5, it will be an error

	// iterate over array
	// 1. for idx iteration
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	// 2. range idx
	for i := range arr3 {
		fmt.Println(arr3[i])
	}
	// 3. get idx, val together
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

}
