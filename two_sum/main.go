package main

import "fmt"

func TwoSum(numbers []int, target int) [2]int {
	if len(numbers) < 2 {
		return [2]int{}
	}
	for i := 0; i < len(numbers)-1; i++ {
		for j := 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{}
}

func main() {
	arr := []int{1, 2, 3}
	target := 4
	fmt.Printf("%v \n", TwoSum(arr, target))
}
