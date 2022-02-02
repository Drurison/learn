package main

import (
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
	//var numOfShips int

}

// Creates a new instance of your generator.
// parameters should be sent to this function and stored in the
// CustomGenerator instance.
func NewCustomGenerator() *CustomGenerator {
	return &CustomGenerator{
		// Init params here...
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
	var grid []bool = make([]bool, height*width)
	var size = len(grid)
	for i := 0; i < size; i++ {
		grid[i] = rand.Intn(2) == 1
		
	} 
	return grid, nil
}

// Just a simple method to output your generated grid to the terminal
// Feel free to change this if you want
func printGrid(grid []bool, height, width int) {

	var rows []string
	var size = len(grid)
	var expectedSize = height * width

	if size != expectedSize {
		panic(fmt.Sprintf("You fucked up\nExpected grid size = (%d*%d) = %d, Received %d", height, width, expectedSize, size))
	}

	for y := 0; y < height; y++ {
		row := ""
		for x := 0; x < width; x++ {
			i := width*y + x
			if grid[i] {
				row += "X "
			} else {
				row += "- "
			}
		}
		rows = append(rows, strings.TrimSpace(row))
	}

	fmt.Println(strings.Join(rows, "\n"))
}

func main() {

	height, width := 10, 10

	grid, err := NewCustomGenerator().Generate(height, width)
	if err != nil {
		panic(err)
	}

	printGrid(grid, height, width)
}
