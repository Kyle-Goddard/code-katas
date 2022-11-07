package wordwrap_test

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

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

func TestWrap(t *testing.T) {
	line := "Hello, World. It's a beautiful day"
	wrapWidth := 15

	got := Wrap(line, wrapWidth)
	want := "Hello, World.\nIt's a\nbeautiful day"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

const maxLineWidth int = 25 // char

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func TestRunWrap(t *testing.T) {
	file, err := os.Open("/Users/kyle/Documents/Learning/code-katas/024_Word_Wrap/input_text.txt")
	Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	const maxCapacity int = 64000
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(Wrap(line, maxLineWidth))
	}
	Check(scanner.Err())
}
