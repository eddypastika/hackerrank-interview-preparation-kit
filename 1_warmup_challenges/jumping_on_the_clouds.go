package main

import "fmt"

func main() {
	//clouds := []int32{0,0,0,0,1,0}
	//clouds := []int32{0,0,1,0,0,1,0}
	clouds := []int32{0, 0, 0, 1, 0, 0}

	fmt.Println("Input: ", clouds)
	fmt.Println("Result: ", jumpingOnClouds(clouds))
}

func jumpingOnClouds(c []int32) int32 {
	var countStep int32

	idx := 0
	for idx < len(c)-1 {
		if idx == len(c)-2 {
			if c[idx+1] == 0 {
				countStep++
			}
			break
		}

		if c[idx+2] == 0 {
			idx += 2
		} else {
			idx++
		}
		countStep++
	}

	return countStep
}
