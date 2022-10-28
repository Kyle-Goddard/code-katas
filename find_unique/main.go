package main

import "fmt"

func findUniq(arr []float32) float32 {
	var repeated float32
	var uniq float32

	if arr[0] == arr[1] {
		repeated = arr[0]
	} else {
		repeated = arr[2]
	}

	for _, v := range arr {
		if v != repeated {
			uniq = v
			break
		}
	}

	return uniq
}

func main() {
	fmt.Println(findUniq([]float32{1, 1, 1, 2, 1, 1}))
	fmt.Println(findUniq([]float32{0, 0, 0.55, 0, 0}))
}
