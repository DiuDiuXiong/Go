package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(s), cap(s), s)
}

/*
1. every time extension triggerred, cap will double
2. to create slice:
- s := []<type>{vals, ...}
- s := make([]<type>, len)
- s := male([]<type>, len, cap)
3. slice initial value is nil
4. copy(slice dst, slice src), copy until capacity
5. to delete an element indexed at x: <slice> = append(<slice>[:x], <slice>[x+1:]...), won't change cap, will cause len-1
- it will also change the original array, for the deleted length due to slice delete element, it will copy from original array
- arr := []int{0,1,2,3,4,5}; s := arr[:]
- s = s.append(s[:2], s[4:])
- original: [0, 1, 2, 3, 4, 5]
- new     : [0, 1, 4, 5, 4, 5] (last 4, 5 copy from original array) (the two originally pointed([3,4]), replaced by [4,5])
6. popping:
- <slice> = <slice>[x:] will cause cap,len -= x
- <slice> = <slice>[:x] will cause len -= x
*/
/*
func main() {
	// Create
	var s []int // Zero value for slice is nil
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)
	printSlice(s2)

	s3 := make([]int, 16, 32)
	printSlice(s3)

	// Copy
	fmt.Println("Copying slice")
	copy(s2, s1) // dest, source, copy until capacity
	copy(s1, s2)
	printSlice(s2)
	printSlice(s1)

	// Delete
	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	// Pop
	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	printSlice(s2)

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(front, tail)
	printSlice(s2)

	// Test Delete
	testArr := [...]int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(testArr)
	sx := testArr[:]

	sx = append(sx[:3], sx[7:]...)
	fmt.Println(testArr)
}*/
