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
	split := strings.Split(line, " ")

	result, intermediate, addWord := "", "", ""

	for i, word := range split {
		fmt.Printf("%v: %s\n", i, word)
		if len(addWord) == 0 {
			addWord = word
		} else {
			addWord = fmt.Sprintf("%s %s", addWord, word)
		}

		if len(addWord) <= width {
			intermediate = addWord
		} else {
			result = fmt.Sprintf("%s\n%s", result, intermediate)
			intermediate = ""
			addWord = ""
		}

	}

	return result
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
		// fmt.Printf("[%s]\n", scanner.Text())

		line := scanner.Text()
		fmt.Printf("[%s]\n", line)
	}

	check(scanner.Err())
}
