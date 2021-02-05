package main

import "fmt"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	var speed int
	fmt.Scan(&speed)

	var lightCount int
	fmt.Scan(&lightCount)

	res := make([]int, lightCount)

	for i := 0; i < lightCount; i++ {
		var distance, duration int
		fmt.Scan(&distance, &duration)

		dis := (distance * 3600) / (duration * 1000)

		res = append(res, dis)

	}
	for i := speed; i >= 1; i-- {
		sec := false

		for _, v := range res {

			if (v/i)%2 != 0 {
				sec = true
				break
			}
		}

		if !sec {
			fmt.Println(i) // Write answer to stdout
			break
		}
	}
}
