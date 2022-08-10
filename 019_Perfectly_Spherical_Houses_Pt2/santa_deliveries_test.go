package main

import (
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
			got := CountDeliveries(tC.input)
			want := tC.want
			assert.Equal(t, got, want)
		})
	}
}
