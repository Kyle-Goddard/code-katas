package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v \n", gameOfThree(10))
	fmt.Printf("%v \n", gameOfThree(100))
	fmt.Printf("%v \n", gameOfThree(1000))
	fmt.Printf("%v \n", gameOfThree(10000))
}

func gameOfThree(x int) []int {
	ans := []int{x}
	ansLast := x

	for ansLast > 1 {
		ans = append(ans, nextNum(ansLast))
		ansLast = ans[len(ans)-1]
	}

	return ans
}

func nextNum(x int) int {
	switch x % 3 {
	case 0:
		return x / 3
	case 1:
		return x - 1
	case 2:
		return x + 1
	}
	return 0
}
