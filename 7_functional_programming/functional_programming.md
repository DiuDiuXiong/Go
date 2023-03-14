# Functional Programming
1. Means functions can be used as variable, parameter, return value
2. For example in [function example](../2_basic_gramma/6_functions.go).
`func apply(op func(int, int) int, a, b int) int` Can use function as input.
3. Closure [check adder](./adder/adder.go)
4. For most most classic functional programming (not for golang)
   - Non-change (only const & programming)
   - Function can only have one parameter
5. Check [01_fibonacci](01_fibonacci.go). 
   - Use closure to store fibonacci
   - Functions can also implement interface, just pass it as receiver same as struct.
6. Check [02_count_tree](02_count_tree.go)
   - Closure allow recursion without passing variable to next, just define a shared variable out of function scope.