package main

import "fmt"

func main() {
	arr := []int32{1, 2, 3, 4, 5}
	d := int32(4)

	// 2 >> 3,4,5,1,2

	fmt.Println("input: ", arr, d)
	fmt.Println("result: ", rotLeft(arr, d))
}

func rotLeft(a []int32, d int32) []int32 {
	res := make([]int32, len(a))

	idx := 0
	for i := d; i < int32(len(a)); i++ {
		res[idx] = a[i]
		idx++
	}


	for i := int32(0); i < d; i++ {
		res[idx] = a[i]
		idx++
	}

	return res
}
