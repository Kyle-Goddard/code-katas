package squareorroot

import (
	"fmt"
	"math"
)

func main() {
	numberList := []int{5, 4, 3, 2, 1, 9, 8, 7, 6}

	fmt.Printf("%v\n", squareOrRootList(numberList))
}

func squareOrRootList(inputSlice []int) []int {
	for i, num := range inputSlice {
		inputSlice[i] = getSquareOrRoot(num)
	}
	return inputSlice
}

func getSquareOrRoot(num int) int {
	root := math.Sqrt(float64(num))

	if root == math.Floor(root) || root == math.Ceil(root) {
		return int(root)
	}

	return int(math.Pow(float64(num), 2))
}
