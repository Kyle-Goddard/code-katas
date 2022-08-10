package main

import (
	"testing"
)

var current = Coord{
	x: 0,
	y: 0,
}

func TestGoNorth(t *testing.T) {
	assertCorrectOutput := func(t testing.TB, got, want Coord) {
		t.Helper()
		if got != want {
			t.Errorf("got %+v but expected %+v", got, want)
		}
	}

	t.Run("move north", func(t *testing.T) {
		got := GoNorth(current)
		want := Coord{
			x: 0,
			y: 1,
		}

		assertCorrectOutput(t, got, want)
	})
}

func TestGoSouth(t *testing.T) {
	assertCorrectOutput := func(t testing.TB, got, want Coord) {
		t.Helper()
		if got != want {
			t.Errorf("got %+v but expected %+v", got, want)
		}
	}

	t.Run("move south", func(t *testing.T) {
		got := GoSouth(current)
		want := Coord{
			x: 0,
			y: -1,
		}

		assertCorrectOutput(t, got, want)
	})
}

func TestGoEast(t *testing.T) {
	assertCorrectOutput := func(t testing.TB, got, want Coord) {
		t.Helper()
		if got != want {
			t.Errorf("got %+v but expected %+v", got, want)
		}
	}

	t.Run("move east", func(t *testing.T) {
		got := GoEast(current)
		want := Coord{
			x: 1,
			y: 0,
		}

		assertCorrectOutput(t, got, want)
	})
}

func TestGoWest(t *testing.T) {
	assertCorrectOutput := func(t testing.TB, got, want Coord) {
		t.Helper()
		if got != want {
			t.Errorf("got %+v but expected %+v", got, want)
		}
	}

	t.Run("move west", func(t *testing.T) {
		got := GoWest(current)
		want := Coord{
			x: -1,
			y: 0,
		}

		assertCorrectOutput(t, got, want)
	})
}

func TestFollowDirections(t *testing.T) {
	assertCorrectOutput := func(t testing.TB, got, want [2]int) {
		t.Helper()
		if got != want {
			t.Errorf("got %+v but expected %+v", got, want)
		}
	}

	testCases := []struct {
		desc         string
		input        string
		numVisited   int
		numDelivered int
	}{
		{
			desc:         "One step right > ",
			input:        ">",
			numVisited:   2,
			numDelivered: 2,
		},
		{
			desc:         "4 houses in a square",
			input:        "^>v<",
			numVisited:   5,
			numDelivered: 4,
		},
		{
			desc:         "Very Lucky children",
			input:        "^v^v^v^v^v",
			numVisited:   11,
			numDelivered: 2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := FollowDirections(tC.input)
			want := [2]int{tC.numVisited, tC.numDelivered}

			assertCorrectOutput(t, got, want)
		})
	}
}
