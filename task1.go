package main

import "fmt"

func main() {
	//	var input [4]int = [4]int{5, 2, 3, 1}
	input := make([]int, 4, 4)
	var min, max int
	var flagmin, flagmax, flag bool = true, true, true

	for i := 0; i < 4; i++ {
		fmt.Scan(&input[i])
	}

	min = input[0]
	max = input[0]

	for i := 0; i < 4; i++ {
		if min > input[i] {
			flagmin = false
		}
		if max < input[i] {
			flagmax = false
		}
	}

	if flagmin {
		for i := 1; i < 4; i++ {
			if input[i-1] > input[i] {
				flag = false
			}
		}
	} else if flagmax {
		for i := 1; i < 4; i++ {
			if input[i-1] < input[i] {
				flag = false
			}
		}
	}

	if flag {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
