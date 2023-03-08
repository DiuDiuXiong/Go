package main

import "fmt"

/**
arr := [...]int{0,1,2,3,4,5,6,7}
s := arr[2:6]  // s is [2,3,4,5]
*/

/**
1. []<type> if do not include length, then it is a slice
2. slice is not an object, it is view to the original array, so when we pass slice to functions, the original array can be changed
3. after slice being created, it can be re-sliced: <slice_param> = <slice_param>[?:?]
4. check out of scope slice
- if sliceA = sliceB[x:y], where x and/or y exceed scope of sliceB, but not scope of original array, it will return extended scope of original array, check sLong, sShort
- if exceed scope of original array, out of scope error
5. Phenomenon above is because slice is composed of:
	- ptr: pointer to head of array
	- len: length of slice (index exceed here will give error)
	- cap: as long as do not exceed cap, can continue to extend (head to tail length)
6. Slice can only extend backward, not forward
7. When use append(<slice>, <num>)
- if slice not exceed original array, change the number next to slice
- if slice exceed original array, Golang will create a new array, copy the old one and return slice to that
*/

func updateSlice(s []int) {
	s[0] = 100
}

func update2DSlice(s [][]int) {
	s[0][0] = 100
	s[1][1] = 200
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	s1 := arr[2:]
	fmt.Println("arr[2:] = ", arr[2:])
	s2 := arr[:]
	fmt.Println("arr[:] = ", arr[:])

	updateSlice(s1)
	fmt.Println("After updateSlice(s1 arr[2:])", s1)
	updateSlice(s2)
	fmt.Println("After updateSlice(s2 arr[:])", s2)

	// re-slice
	s2 = s2[:5] // first 5
	fmt.Println(s2)
	s2 = s2[2:] // chunk first 2
	fmt.Println(s2)

	// 2D - slice
	arr2 := [...][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("Before update arr2", arr2)
	twoDSlice := make([][]int, 2)
	twoDSlice[0] = arr2[0][:]
	twoDSlice[1] = arr2[1][:]
	update2DSlice(twoDSlice)
	fmt.Println("After update arr2", arr2)

	// out of scope slice
	arrTest := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	sLong := arrTest[2:6] // [2,3,4,5]
	fmt.Println(sLong)
	sShort := sLong[3:5] // [5,6] !! extending here
	fmt.Println(sShort)
	//sFull := arrTest[:]
	//sFullExtend := sFull[8:9] here is out of scope error
	//fmt.Println(sFullExtend)
	//fmt.Println(sLong[4]) error, exceed boundary
	fmt.Printf("sLong=%v, len(sLong)=%d, cap(sLong)=%d\n", sLong, len(sLong), cap(sLong))
	fmt.Printf("sShort=%v, len(sShort)=%d, cap(sShort)=%d\n", sShort, len(sShort), cap(sShort))

	// append
	s2 = sLong
	fmt.Println("s2 =", s2)
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12) // out of scope here, Golang will create a new array and copy the old one
	// s5 is viewing different array
	fmt.Println("s3, s4, s5 =", s3, s4, s5)
	fmt.Println("s2 =", s2)
	fmt.Println("arrTest =", arrTest)

	// s5 is viewing different array
	s4[0] = -1
	s5[0] = 999
	fmt.Println(arrTest)
}
