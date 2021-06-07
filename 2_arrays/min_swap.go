package main

import (
	"fmt"
	"sort"
)

func main() {
	//arr := []int32{4,3,1,2}
	arr := []int32{1, 3, 5, 2, 4, 6, 7}

	fmt.Println("input: ", arr)
	fmt.Println("result: ", minimumSwaps(arr))
}

// Arr ...
type Arr struct {
	value int32
	index int
}

// ByValue ...
type ByValue []Arr

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i].value < a[j].value }

func minimumSwaps(input []int32) int32 {
	arr := make([]Arr, 0, len(input))

	for i, v := range input {
		arr = append(arr, Arr{v, i})
	}

	fmt.Println(arr)

	sort.Sort(ByValue(arr))

	fmt.Println(arr)

	idx := make([]int, 0, len(input))
	for _, ar := range arr {
		idx = append(idx, ar.index)
	}

	var result int32
	for i := 0; i < len(input); i++ {
		if i == idx[i] {
			continue
		}

		input[i], input[idx[i]] = input[idx[i]], input[i]
		idx[i], idx[input[idx[i]]-1] = i, idx[i]
		result++
	}
	return result
}
