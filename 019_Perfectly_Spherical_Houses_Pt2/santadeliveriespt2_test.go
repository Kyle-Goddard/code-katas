package santadeliveriespt2_test

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountDeliveries(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  int
	}{
		{
			desc:  "One step right > ",
			input: ">",
			want:  2,
		},
		{
			desc:  "4 houses in a square",
			input: "^>v<",
			want:  4,
		},
		{
			desc:  "Very Lucky children",
			input: "^v^v^v^v^v",
			want:  2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := CountDeliveries(tC.input)
			want := tC.want
			assert.Equal(t, want, got)
		})
	}
}

func TestCountDeliveriesWithRobot(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  int
	}{
		{
			desc:  "test 1: Simplest test",
			input: "^v",
			want:  3,
		},
		{
			desc:  "test 2: Each go back to starting point",
			input: "^>v<",
			want:  3,
		},
		{
			desc:  "test 3: Each keep moving in opposite directions",
			input: "^v^v^v^v^v",
			want:  11,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := CountDeliveriesWithRobot(tC.input)
			want := tC.want
			assert.Equal(t, want, got)
		})
	}
}

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

func CountDeliveriesWithRobot(input string) int {
	directions := parseInstructions(input)
	housesVisited := map[Coord]int{}
	followDirectionsWithRobot(directions, housesVisited)

	return len(housesVisited)
}

func followDirections(directions []Direction, housesVisited map[Coord]int) {
	santaCoord := Coord{x: 0, y: 0}
	housesVisited[santaCoord] = 1

	for _, d := range directions {
		switch d {
		case North:
			santaCoord.y += 1
		case South:
			santaCoord.y -= 1
		case East:
			santaCoord.x += 1
		case West:
			santaCoord.x -= 1
		}

		housesVisited[santaCoord] += 1
	}
}

func followDirectionsWithRobot(directions []Direction, housesVisited map[Coord]int) {
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

func TestRunDeliveries(t *testing.T) {
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
		resultWithRobot := CountDeliveriesWithRobot(line)
		fmt.Printf("santa: %v,  santa+robot: %v \n", result, resultWithRobot)
	}

	check(scanner.Err())
}
