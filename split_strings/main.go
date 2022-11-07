package main

import (
	"fmt"
	"regexp"
)

func Solution(str string) []string {
	if len(str)%2 != 0 {
		str = fmt.Sprintf("%s_", str)
	}
	split := make([]string, len(str)/2)

	re := regexp.MustCompile(`[[:alpha:]_]{2}`)
	for i, match := range re.FindAllString(str, -1) {
		split[i] = match
	}

	return split
}

func main() {
	fmt.Printf("%v \n", Solution("abcdef"))
}
