package main

import (
	"testing"
)

func TestPaddingDots(t *testing.T) {
	assertCorrectOutput := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q but expected %q", got, want)
		}
	}

	t.Run("one dot", func(t *testing.T) {
		got := paddingDots(1)
		want := "."

		assertCorrectOutput(t, got, want)
	})

	t.Run("three dots", func(t *testing.T) {
		got := paddingDots(3)
		want := "..."

		assertCorrectOutput(t, got, want)
	})

	t.Run("zero dots", func(t *testing.T) {
		got := paddingDots(0)
		want := ""

		assertCorrectOutput(t, got, want)
	})

	t.Run("negative dots", func(t *testing.T) {
		got := paddingDots(-1)
		want := ""

		assertCorrectOutput(t, got, want)
	})
}

func TestAlignText(t *testing.T) {
	assertCorrectOutput := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q but expected %q", got, want)
		}
	}

	t.Run("left aligned text example", func(t *testing.T) {
		got, _ := AlignText("hello", 7, "L")
		want := "hello.."

		assertCorrectOutput(t, got, want)
	})
	t.Run("left aligned text example", func(t *testing.T) {
		got, _ := AlignText("hello", 11, "L")
		want := "hello......"

		assertCorrectOutput(t, got, want)
	})

	t.Run("right aligned text example", func(t *testing.T) {
		got, _ := AlignText("hello", 7, "R")
		want := "..hello"

		assertCorrectOutput(t, got, want)
	})
	t.Run("right aligned text example", func(t *testing.T) {
		got, _ := AlignText("hello", 11, "R")
		want := "......hello"

		assertCorrectOutput(t, got, want)
	})

	t.Run("center aligned text example", func(t *testing.T) {
		got, _ := AlignText("hello", 7, "C")
		want := ".hello."

		assertCorrectOutput(t, got, want)
	})

	t.Run("center aligned text example", func(t *testing.T) {
		got, _ := AlignText("hello", 11, "C")
		want := "...hello..."

		assertCorrectOutput(t, got, want)
	})

	t.Run("center aligned text example", func(t *testing.T) {
		got, _ := AlignText("hi", 7, "C")
		want := "...hi.."

		assertCorrectOutput(t, got, want)
	})

	// t.Run("Testing if TextAlign fails when max width is shorter than the length of the input word", func(t *testing.T) {
	// 	got, _ := AlignText("hello", 7, "L")
	// 	want := "hello.."

	// 	if got != want {
	// 		t.Errorf("got %q but expected %q", got, want)
	// 	}
	// })
}

// // examples
// func ExampleAlignText() {
// 	aligned, err := AlignText("hello", 11, "L")
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	fmt.Println(aligned)
// 	// Output: hello......

// 	aligned, err = AlignText("hello", 11, "R")
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	fmt.Println(aligned)
// 	// Output: ......hello
// }
