package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// The interface in my battleships code
type Generator interface {
	Generate(height, width int) ([]bool, error)
}

// Your custom generator implementation
type CustomGenerator struct {
	// Store any parameters here...
	numOfShips int
}

// Creates a new instance of your generator.
// parameters should be sent to this function and stored in the
// CustomGenerator instance.
func NewCustomGenerator(numOfShips int) *CustomGenerator {
	return &CustomGenerator{
		// Init params here...
		numOfShips: numOfShips,
	}
}

// The grid generator. Don't change the input/output params as they
// need to match the interface. Height and width are the size of the grid.
// The []bool is the generated grid. True means there is a ship and false means
// There is no ship. The grid has been flattened (because its easier to handle in software).
// For example:
//
//   | 1 | 1 | 0 |
//   | 0 | 0 | 1 |   ==   [1, 1, 0, 0, 0, 1, 0, 1, 0]
//   | 0 | 1 | 0 |       (where 0 == false and 1 == true)
//
func (g *CustomGenerator) Generate(height, width int) ([]bool, error) {
	rand.Seed(time.Now().UnixNano())
	// Your generation code goes here...
	var size = height * width
	var grid []bool = make([]bool, size)
	if g.numOfShips > size {
		return nil, errors.New("Too many ships for grid, sadge --says Joe")

	}
	for existingShips := 0; existingShips < g.numOfShips; {
		position := rand.Intn(size)
		if !grid[position] {
			grid[position] = true
			existingShips++
		}
	}

	return grid, nil
}

// Just a simple method to output your generated grid to the terminal
// Feel free to change this if you want
func printGrid(grid []bool, height, width int) {

	var rows []string
	var actualSize = len(grid)
	var expectedSize = height * width

	if actualSize != expectedSize {
		panic(fmt.Sprintf("You fucked up\nExpected grid size = (%d*%d) = %d, Received %d", height, width, expectedSize, actualSize))
	}

	for y := 0; y < height; y++ {
		row := ""
		for x := 0; x < width; x++ {
			i := width*y + x
			if grid[i] {
				row += "S "
			} else {
				row += "- "
			}
		}
		rows = append(rows, strings.TrimSpace(row))
	}

	fmt.Println(strings.Join(rows, "\n"))
}

func main() {

	height, width := 3, 3
	grid, err := NewCustomGenerator(5).Generate(height, width)
	if err != nil {
		panic(err)
	}

	printGrid(grid, height, width)

}
