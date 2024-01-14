// Testing & benchmarking, test coverage

// package fortest is for testing and exampling purposes
package fortest


// FuncSum takes an unlimited amount of integers and returns their sum.
func FuncSum(i ...int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}
