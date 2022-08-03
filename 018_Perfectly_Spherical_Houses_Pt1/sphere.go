package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Coord struct {
	x int
	y int
}

type Visited []Coord

func GoNorth(current Coord) Coord {
	current.y += 1
	return current
}

func GoSouth(current Coord) Coord {
	current.y -= 1
	return current
}

func GoEast(current Coord) Coord {
	current.x += 1
	return current
}

func GoWest(current Coord) Coord {
	current.x -= 1
	return current
}

func FollowDirections(input string) [2]int {
	current := Coord{x: 0, y: 0}
	visited := Visited{current}
	deliveries := map[Coord]int{current: 1}

	for _, char := range input {
		switch char {
		case '^':
			current = GoNorth(current)
		case 'v':
			current = GoSouth(current)
		case '>':
			current = GoEast(current)
		case '<':
			current = GoWest(current)
		}
		visited = append(visited, current)
		if _, ok := deliveries[current]; !ok {
			deliveries[current] = 1
		} else {
			deliveries[current] += 1
		}
	}

	return [2]int{len(visited), len(deliveries)}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	file, err := os.Open("/Users/kyle/Documents/Learning/code-katas/018_Perfectly_Spherical_Houses_Pt1/input_long.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	const maxCapacity int = 64000
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {
		line := scanner.Text()
		result := FollowDirections(line)

		visited := result[0]
		deliveries := result[1]
		fmt.Printf("Santa delivers %v gifts to %v unique houses \n", visited, deliveries)
		fmt.Printf("He has achieved %.2f percent delivery efficiency \n", float64(deliveries)/float64(visited)*100)

	}

	check(scanner.Err())
}
