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

type Direction string

const (
	North Direction = "North"
	South Direction = "South"
	East  Direction = "East"
	West  Direction = "West"
)

func CountDeliveries(input string) int {
	directions := parseInstructions(input)
	housesVisited := map[Coord]int{}
	followDirections(directions, housesVisited)

	return len(housesVisited)
}

func followDirections(directions []Direction, housesVisited map[Coord]int) {
	santaCoord := Coord{x: 0, y: 0}
	roboCoord := Coord{x: 0, y: 0}
	locations := [2]Coord{santaCoord, roboCoord}
	housesVisited[santaCoord] = 1
	whichSanta := 0

	for _, d := range directions {
		switch d {
		case North:
			locations[whichSanta].y += 1
		case South:
			locations[whichSanta].y -= 1
		case East:
			locations[whichSanta].x += 1
		case West:
			locations[whichSanta].x -= 1
		}

		housesVisited[locations[whichSanta]] += 1

		switch whichSanta {
		case 0:
			whichSanta = 1
		case 1:
			whichSanta = 0
		}
	}
}

func parseInstructions(input string) []Direction {
	result := make([]Direction, 0)
	for _, char := range input {
		switch char {
		case '^':
			result = append(result, North)
		case 'v':
			result = append(result, South)
		case '>':
			result = append(result, East)
		case '<':
			result = append(result, West)
		default:
			panic("unknown direction")
		}
	}
	return result
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	file, err := os.Open("/Users/kyle/Documents/Learning/code-katas/019_Perfectly_Spherical_Houses_Pt2/input_long.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	const maxCapacity int = 64000
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {
		line := scanner.Text()
		result := CountDeliveries(line)
		fmt.Println(result)
	}

	check(scanner.Err())
}
