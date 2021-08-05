package solution

// you can also use imports, for example:
//import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
const divd = (1000 * 1000 * 1000) + 7

func Solution(A []int) int {
	// write your code in Go 1.4
	cu := 0
	ver := 0
	last := A[len(A)-1]
	for i, v := range A[:len(A)-1] {
		diffr := (last - v)
		cu += diffr
		prod := diffr * (i + 1)
		if prod > ver {
			ver = prod
		}
	}
	return (cu - ver) % divd
}
