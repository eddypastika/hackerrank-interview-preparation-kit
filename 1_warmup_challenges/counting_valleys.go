package main

import "fmt"

func main() {
	steps := 8
	path := "UDDDUDUU"

	fmt.Println("Input: ", steps, path)
	fmt.Println("Result: ", countingValleys(int32(steps), path))
}

func countingValleys(steps int32, path string) int32 {
	var countValley int32

	seaLevel := 0
	for i := 0; i < int(steps); i++ {
		if path[i:i+1] == "D" {
			seaLevel--
		} else {
			seaLevel++
		}

		if seaLevel == 0 && path[i:i+1] == "U" {
			countValley++
		}
	}

	return countValley
}
