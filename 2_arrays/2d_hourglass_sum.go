package main

import "fmt"

func main() {
	arr := [][]int32{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
	}

	fmt.Println("input: ", arr)
	fmt.Println("result: ", hourglassSum(arr))
}

func hourglassSum(arr [][]int32) int32 {
	var max int32

	for i := 0; i < len(arr)-2; i++ {
		sum1 := int32(0)
		for j := 0; j < len(arr)-2; j++ {
			sum1 = arr[i][j] + arr[i][j+1] + arr[i][j+2] +
				arr[i+1][j+1] +
				arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]

			if i == 0 && j == 0 {
				max = sum1
			}
			
			if sum1 > max {
				max = sum1
			}
		}
	}

	return max
}
