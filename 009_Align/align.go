package align

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func PaddingDots(x int) string {
	var result string
	for x > 0 {
		result += "."
		x--
	}
	return result
}

func AlignText(word string, width int, alignment string) (string, error) {
	if width < len(word) {
		return "", errors.New("max width is less than the length of the input word")
	}

	var result string

	switch alignment {
	case "L":
		result = word + PaddingDots(width-len(word))
		// result = word + strings.Repeat(".", width-len(word))
		// chose not to use strings.Repeat because it panics if count is negative or if the result of (len(s) * count) overflows.

	case "R":
		result = PaddingDots(width-len(word)) + word
		// result = strings.Repeat(".", width-len(word)) + word

	case "C":
		halfPadding := (width - len(word)) / 2
		remainder := (width - len(word)) % 2
		result = PaddingDots(int(remainder+halfPadding)) + word + PaddingDots(int(halfPadding))

	default:
		return "", errors.New("invalid arguement: text alignment should be L, R, or C only")
	}

	return result, nil
}

func main() {
	if len(os.Args) < 4 {
		log.Fatalf("invalid input: please follow the input format: <word> <width> <text-align>")
	}

	word := os.Args[1]
	width, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("invalid input: please provide a valid integer for the width input")
	}
	align := os.Args[3]

	aligned, err := AlignText(word, width, align)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(aligned)
}

// flags package
