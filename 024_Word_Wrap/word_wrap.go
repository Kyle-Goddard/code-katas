package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const maxLineWidth int = 25 // char

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Wrap(line string, width int) string {
	words := strings.Fields(line)

	if len(words) == 0 {
		return ""
	}

	result := ""

	newLine := words[0]

	for _, word := range words[1:] {
		if len(newLine)+len(word) < width {
			newLine = fmt.Sprintf("%s %s", newLine, word)
		} else {
			result = fmt.Sprintf("%s%s\n", result, newLine)
			newLine = word
		}
	}
	result = fmt.Sprintf("%s%s", result, newLine)

	return strings.TrimSuffix(result, "\n")
}

func main() {
	file, err := os.Open("/Users/kyle/Documents/Learning/code-katas/024_Word_Wrap/input_text.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	const maxCapacity int = 64000
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(Wrap(line, maxLineWidth))
	}
	check(scanner.Err())
}
