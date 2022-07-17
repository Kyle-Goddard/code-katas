package main

import "fmt"

func main() {
	str := "(())((()()))"
	fmt.Printf("%v \n", validParenthesis(str))
}

func validParenthesis(parens string) bool {
	stack := []string{}

	for _, b := range parens {
		if string(b) == "(" {
			stack = append(stack, "(")

		} else if string(b) == ")" {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
