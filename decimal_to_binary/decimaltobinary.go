package decimaltobinary

import (
	"fmt"
	"reflect"
	"testing"
)

func DecimalToBinary(num int) []string {
	return []string{"1"}
}

func main() {
	target := 5

	fmt.Println(DecimalToBinary(target))
}

func TestDecimalToBinary(t *testing.T) {
	t.Run("should work for 0", func(t *testing.T) {
		got := DecimalToBinary(0)
		want := []string{"0"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %s \n want: %s \n", got, want)
		}
	})
}
