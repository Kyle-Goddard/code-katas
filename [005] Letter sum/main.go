package main

import "fmt"

func main() {
	line := "excellent"

	fmt.Printf("%s --> sum: %v \n", line, letterSum(line))
}

func letterSum(line string) int {
	sum := 0

	for _, char := range line {
		letterNum := int(char) - 96
		if letterNum > 0 && letterNum <= 26 { // any runes outside range a-z (lowercase) will not be added
			sum += letterNum
		}
	}
	return sum
}
