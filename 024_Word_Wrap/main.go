package main

import (
	"fmt"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	file, err := os.Open("./[024] Word Wrap/input_text.txt")
	check(err)

	fmt.Print(file)
}
