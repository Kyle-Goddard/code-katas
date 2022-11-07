package mergelists

import "fmt"

func main() {
	a, b := []int{1, 2, 4, 6, 9}, []int{2, 3, 5, 7, 9, 13, 15, 17}

	fmt.Printf("%v, %v --> %v \n", a, b, mergeLists(a, b))
}

func mergeLists(a []int, b []int) []int {
	outputList := make([]int, len(a)+len(b))
	pa, pb := 0, 0

	for i := range outputList {
		if pa < len(a) && pb < len(b) {
			if a[pa] < b[pb] {
				outputList[i] = a[pa]
				pa++
			} else {
				outputList[i] = b[pb]
				pb++
			}
		} else {
			if pa < len(a) {
				outputList[i] = a[pa]
				pa++
			}

			if pb < len(b) {
				outputList[i] = b[pb]
				pb++
			}
		}
	}

	return outputList
}
