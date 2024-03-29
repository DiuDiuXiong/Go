package main

import "fmt"

/**
Here each time we call adder, sum get changed as well, and will affect later call. This is the idea of closure:
Closure:
1. For each function, variable within scope are local variables, variable outside scope are called free variable
2. Each variable that are out of scope of function, but function depends on, will be pointed when returned by compiler(function)
3. Those free variable can be var or struct, while if it is struct, it can further linked to other vars... finally it will finish linking
4. All those stuff:
	- Function itself
	- local vars
	- free vars and linked vars
Are returned together, which forms a closure.
5. See sharedFreeVar() function, which shows pointed free vars can be shared, and not only controlled by function itself
6. For other language check imooc 7-1 11:48
*/

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func sharedFreeVar() (func(), func()) {
	sharedVar := 0
	f1 := func() {
		sharedVar += 1
		fmt.Println(sharedVar)
	}

	f2 := func() {
		sharedVar += 2
		fmt.Println(sharedVar)
	}

	return f1, f2
}

type tricky struct {
	a, b *int
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(a(i))
	}

	f1, f2 := sharedFreeVar()
	f1()
	f2()

	// tricky case, closure of struct
	t := tricky{}
	x1, x2, x3 := 0, 0, 0
	t.a, t.b = &x1, &x2
	f := func() {
		*t.a += 1
		*t.b += 2
		fmt.Println(*t.a, *t.b)
	}

	f() // a = 1, b = 2
	x1 += 1
	f() // a = 3, b = 4
	t.a = &x3
	f() // point to father struct, rather than point to each variable separately
}
