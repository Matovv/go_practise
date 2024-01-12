package main

import (
	"fmt"
	"testing"
)

func TestSquare(t *testing.T) {
	fmt.Println("//////", "Testing square are calculation ...", "//////")
	testSquare := Square{
		Width:  10,
		Height: 10,
	}
	area := Info(testSquare)

	got := area
	want := 100.0

	if got != want {
		t.Errorf("Incorrect calculation of the square area, got: %f, want: %f!", got, want)
	}

}
