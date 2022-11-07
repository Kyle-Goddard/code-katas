package reversewords

import (
	"fmt"
	"unicode"
)

func main() {
	str := "The quick brown fox jumps over the lazy  dog."

	fmt.Printf("%s \n", reverseStrBest(str))
}

func reverseStrBest(str string) string {
	var rev string
	var word string

	for _, i := range str {
		if i == ' ' {
			rev = rev + word + " " // Adds word and space to result
			word = ""              // Empties word variable
		} else {
			word = string(i) + word // Adds letter to temporary word variable
		}
	}

	return rev + word // reverse those words
}

func reverseStr(str string) string {
	words := []string{""}
	i := 0

	for _, char := range str {
		if unicode.IsSpace(char) {
			words = append(words, "")
			i++
		} else {
			words[i] = fmt.Sprintf("%s%s", words[i], string(char))
		}
	}

	res := reverseWord(words[0])
	for i, w := range words {
		if i > 0 {
			res = fmt.Sprintf("%s %s", res, reverseWord(w))
		}
	}

	return res
}

func reverseWord(word string) string {
	res := ""
	for i := range word {
		res = fmt.Sprintf("%s%s", res, string(word[len(word)-i-1]))
	}

	return res
}
