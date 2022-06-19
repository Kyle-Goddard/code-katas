package main

import "fmt"

func main() {
	numberList := []int{5, 4, 3, 2, 1, 9, 8, 7, 6}
	rotate := 0

	for rotate <= 10 {
		fmt.Printf("%v rotated by %v --> %v \n", numberList, rotate, rotateList(numberList, rotate))
		rotate++
	}
}

func rotateList(inputSlice []int, rotateBy int) []int {
	outputSlice := make([]int, len(inputSlice))

	for i := range inputSlice {
		outputSlice[i] = inputSlice[(i+rotateBy)%len(inputSlice)]
	}

	return outputSlice
}
