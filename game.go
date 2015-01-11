package main

type cell bool

type grid [][]cell

type game struct {
	Grid grid
}

func newGame() game {
	grid := make(grid, 20)
	for i := range grid {
		grid[i] = make([]cell, len(grid))
	}
	g := game{grid}
	return g
}

func init() {
}

func main() {
}
