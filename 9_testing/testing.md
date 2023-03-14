# Testing
## Traditional testing & table-driven test (TDT)

Traditional testing Disadvantage:
- Test data & test logic mix together
- Error message not explicit
- Since due to `assert` if there are one fail, the rest cannot run.

TDT in order to solve issues above:
- Define test data as `[]struct{<param_name> <param_type>} {{..},...,{..}}`
- Use a for loop to iterate over struct above, and test each one. If some of those failed, continue.
- Whenever a file named as: `<...>_test.go` it will be a test file.
- Check [01_triangle_test.go](01_repeating_test.go)/[01_triangle_test.go](01_repeating_test.go) for example

## Test Cover range
Run the two commands below in a row to see which function is tested and which is not.
1. Generate cover profile & run test `go test -coverprofile=c.out`
2. Get cover profile (see who covered, who is not) `go tool cover -html=c.out`

## Benchmark
Check [BenchmarkSubstr](01_repeating_test.go) for example. 
```
func <func_name>(b *testing.B) {
    ...
    for i:= 0; i < b.N; i++ {
        ... run test cases
    }
}
```
Then run the function:
1. `b.N` will be determined by system algorithm
2. it will output average time for each function call (thus give benchmark)
```
BenchmarkSubstr-10    	  287188	      4143 ns/op
```
Means 287188 times with average function calling time be 4143 ns/op.
3. Same thing be performed on terminal `go test -bench .`
4. Some time to prepare stage can take much longer, `b.ResetTimer()` before start counting time. [BenchmarkSubStr2](./01_repeating_test.go)
5. To see which cpu command cause most time (detail check)
    - `go test -bench . -cpuprofile cpu.out` while cpu.out is a binary file
    - `go tool pprof cpu.out` && `web` to get a view of which part cause more time
6. [As web graph shown](pprof001.svg)
   - The larger the graph, the longer time it takes
   - We see it take a long time to decode `string` to `[]rune` for [`lengthOfNonRepeatingSubStr`](01_testing_example.go)
   - Another long is during `map` operation (hash, space, refactor)
   - So the updated version, change to `make([]int, 0xffff)`, which are faster (4161, 5444399875) ==> (49347,1773945209)
7. So code modification: `-cpuprofile` get cpu data => `go tool pprof` to see actual data ==> see where code are slow ==> modify code