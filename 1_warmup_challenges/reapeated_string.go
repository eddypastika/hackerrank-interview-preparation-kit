package main

import (
	"fmt"
)

func main() {
	//inputString := "gfcaaaecbg" // 10 - 3
	//inputNum := 547602
	inputString := "abcda"
	inputNum := 17

	fmt.Println("Input: ", inputString, inputNum)
	fmt.Println("Result: ", repeatedString2(inputString, int64(inputNum)))
}

func repeatedString2(s string, n int64) int64 {
	var aNumb, res int64
	inputLength := int64(len(s))

	for i := 0; i < len(s); i++ {
		if s[i:i+1] == "a" {
			aNumb++
		}
	}

	if aNumb == 0 {
		return 0
	}

	res = n / int64(len(s)) * aNumb
	reminder := n % inputLength
	for i := int64(0); i < reminder; i++ {
		if s[i]  == 'a' {
			res++
		}
	}

	return res
}
