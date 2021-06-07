package main

import (
	"fmt"
)

func main() {
	//arr := []int32{2, 1, 5, 3, 4}
	//arr := []int32{1, 2, 5, 3, 4, 7, 8, 6}
	arr := []int32{1,2,5,3,7,8,6,4}
	fmt.Println(arr)
	minimumBribes(arr)
}

func minimumBribes(q []int32) {
	result := int32(0)
	for i := int32(len(q)) - 1; i >= 0; i-- {
		if q[i]-(i+1) > 2 {
			fmt.Println("Too chaotic")
			return
		}
		for j := max(0, q[i]-2); j < i; j++ {
			if q[j] > q[i] {
				result++
			}
		}
	}
	fmt.Println(result)
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

//func minimumBribes(q []int32) {
//	intArr := make([]int, len(q))
//
//	for i := 0; i < len(q); i++ {
//		intArr[i] = int(q[i])
//	}
//
//	sort.Ints(intArr)
//	fmt.Println(intArr)
//
//	bribe := int32(0)
//	for i := 0; i < len(q); i++ {
//		tmp := q[i] - int32(intArr[i])
//		if tmp > 2 {
//			fmt.Println("Too chaotic")
//			return
//		}
//
//		if tmp > 0 {
//			bribe += tmp
//		}
//	}
//
//	fmt.Println(bribe)
//}


//func minimumBribes(q []int32) {
//	min := q[0]
//	bribeCount := 0
//	tmpMap := make(map[int32]int)
//
//	for i := 1; i < len(q); i++ {
//		if min < q[i] {
//			min = q[i]
//		} else if min > q[i] {
//			_, ok := tmpMap[min]
//			if ok {
//				tmpMap[min]++
//			} else {
//				tmpMap[min] = 1
//			}
//
//			bribeCount++
//		}
//	}
//
//	for _, v := range tmpMap {
//		if v > 2 {
//			fmt.Println("Too chaotic")
//			return
//		}
//	}
//	fmt.Println(bribeCount)
//}
