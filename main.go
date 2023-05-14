package main

import (
	"fmt"

	"github.com/anaseto/gruid"
)

func main() {
	// Create a new 20x20 grid.
	gd := gruid.NewGrid(20, 20)
	// Fill the whole grid with dots.
	gd.Fill(gruid.Cell{Rune: '.'})
	// Define a range (5,5)-(15,15).
	rg := gruid.NewRange(5, 5, 15, 15)
	// Define a slice of the grid using the range.
	rectangle := gd.Slice(rg)
	// Fill the rectangle with #.
	rectangle.Fill(gruid.Cell{Rune: '#'})
	// Print the grid using a non-styled string representation.
	fmt.Print(gd)
}
