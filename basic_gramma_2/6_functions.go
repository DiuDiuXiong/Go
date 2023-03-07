package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

/*
*
* 1. func <func_name>(<param name> <type>) <return type> {...}
* 2. function can return two/many values: func <func_name>(...) (<return type 1>, <return type 2>)
* 3. check div2(...) (<return param1>, <return param2> <type 1>), then we don't have to specify what to return, just return is enough
* 4. check apply(op func(int, int) int, a, b int) int {...}, functions can be used as input to other functions
  - func <func_name> (<function_name> func(<input types>...) <output types...>, <other params>)

* 5. there are no lambda, override, overload, default param number, operator overload, optional parameter
* 6. func_name func(name... type) <return type> can provide uncertain amount of parameters, check sum(number ...int) int {} function, this is the only special function aspect allowed
*/
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

func div2(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func threeReturns(a, b int) (int, int, int) {
	sum := a + b
	diff := a - b
	multi := a * b
	return sum, diff, multi
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+"(%d %d)", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(number ...int) int {
	s := 0
	for i := range number {
		s += number[i]
	}
	return s
}

/*
func main() {
	fmt.Println(div(13, 3))
	var q, r = div(13, 3)
	fmt.Println(q, r)
	var _, r2 = div(13, 3)
	fmt.Println(r2)
	if result, err := eval(1, 2, ":)"); err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println(result)
	}

	if result, err := eval(1, 2, "*"); err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println(result)
	}

	fmt.Println(apply(pow, 3, 4))
	// anonymous function
	fmt.Println(apply(
		func(a int, b int) int { return a + b },
		3,
		4))
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5))

}
*/
