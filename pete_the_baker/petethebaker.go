package petethebaker

import "fmt"

func Cakes(recipe, available map[string]int) int {
	amts := make([]int, 0)

	for key := range recipe {
		amts = append(amts, available[key]/recipe[key])
	}

	if len(amts) == 0 {
		return 0
	}

	val := amts[0]
	for _, a := range amts {
		if a < val {
			val = a
		}
	}
	return val
}

func main() {
	rec := map[string]int{"flour": 500, "sugar": 200, "eggs": 1}
	ing := map[string]int{"flour": 1200, "sugar": 1200, "eggs": 5, "milk": 200}
	fmt.Println(Cakes(rec, ing))
	rec = map[string]int{"apples": 3, "flour": 300, "sugar": 150, "milk": 100, "oil": 100}
	ing = map[string]int{"sugar": 500, "flour": 2000, "milk": 2000}
	fmt.Println(Cakes(rec, ing))
}

// must return 2
// cakes({flour: 500, sugar: 200, eggs: 1}, {flour: 1200, sugar: 1200, eggs: 5, milk: 200});
// must return 0
// cakes({apples: 3, flour: 300, sugar: 150, milk: 100, oil: 100}, {sugar: 500, flour: 2000, milk: 2000});
