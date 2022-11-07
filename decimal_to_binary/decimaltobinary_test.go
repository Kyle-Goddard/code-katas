package decimaltobinary_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func DecimalToBinary(num int) []string {
	var maxLen int
	if num < 1 {
		maxLen = 1
	} else {
		maxLen = int(math.Floor(math.Log2(float64(num))) + 1)
	}
	result := make([]string, maxLen)

	sum := 0
	var power int

	for i := 0; i < maxLen; i++ {
		power = int(math.Pow(2, float64(maxLen-i-1)))
		if sum+power <= num {
			result[i] = "1"
			sum += power
		} else {
			result[i] = "0"
		}
	}

	return result
}

func TestDecimalToBinary(t *testing.T) {
	testCases := []struct {
		desc int
		want []string
	}{
		{
			0,
			[]string{"0"},
		},
		{
			1,
			[]string{"1"},
		},
		{
			2,
			[]string{"1", "0"},
		},
		{
			5,
			[]string{"1", "0", "1"},
		},
		{
			21,
			[]string{"1", "0", "1", "0", "1"},
		},
	}
	for _, tC := range testCases {
		t.Run("should work for "+fmt.Sprint(tC.desc), func(t *testing.T) {
			got := DecimalToBinary(tC.desc)

			assert.Equal(t, tC.want, got)
		})
	}
}
