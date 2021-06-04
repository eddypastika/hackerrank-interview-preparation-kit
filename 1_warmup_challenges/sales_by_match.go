package main

import "fmt"

func main() {
	ar := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	n := int32(9)

	fmt.Println("Input: ", ar)
	fmt.Println("Result: ", sockMerchant(n, ar))
}

func sockMerchant(n int32, ar []int32) int32 {
	var res int32
	tempMap := make(map[int32]int)

	for i := int32(0); i < n; i++ {
		_, ok := tempMap[ar[i]]
		if ok {
			tempMap[ar[i]]++
		} else {
			tempMap[ar[i]] = 1
		}
	}

	for _, v := range tempMap {
		if v == 2 {
			res++
		}

		if v > 2 {
			res += int32(v / 2)
		}
	}

	return res
}
