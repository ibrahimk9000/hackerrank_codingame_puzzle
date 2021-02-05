package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	a := [][]int{}
	cube := 0
	for i := 0; i < len(A); i++ {
		d := A[i]
		temp := []int{}
		for s := 0; s < d; s++ {
			if s+i < len(A) {
				if A[s+i] >= d {
					temp = append(temp, A[s+i])

				} else {

					d = A[s+i]
					if len(temp) < d {
						temp = append(temp, A[s+i])

					}

				}
			}

		}
		a = append(a, temp)
	}
	for _, v := range a {
		if cube >= len(v) {
			continue
		} else {
			cube = len(v)
		}

	}
	return cube

}

// write your code in Go 1.4
