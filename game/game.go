package game

import (
	"fmt"
	"time"
)

// Point represents a coordinate on the grid of the game
type Point struct {
	X, Y int
}

type cell bool

func (c cell) String() string {
	if c {
		return "[0]"
	}
	return "[ ]"
}

// Game holds the game's logic and the grid of cells where the game takes place
type Game struct {
	grid          [][]cell // 2D slice of cells, true means alive and false means dead
	rows, columns int
	stop          bool
}

// Seed set the living cells on the grid according to points
func (g *Game) Seed(points []Point) {
	for _, p := range points {
		if g.validPoint(p) {
			g.grid[p.X][p.Y] = true
		}
	}
}

func (g *Game) validPoint(p Point) bool {
	return p.X >= 0 && p.X < g.rows && p.Y >= 0 && p.Y < g.columns
}

// neighbours returns the count of alive and dead neighbours from the specified point
func (g *Game) neighbours(p Point) (alive int) {
	deltas := []Point{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for _, d := range deltas {
		x, y := p.X+d.X, p.Y+d.Y
		if g.validPoint(Point{x, y}) { // Only count for valid indexes
			if g.grid[x][y] {
				alive++
			}
		}
	}
	return
}

// tick applies the rules of the game simultaneously for every cell on the grid and creates a new
// grid holding the next generation of cells.
func (g *Game) tick() {
	nextGrid := newGrid(g.rows, g.columns)
	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.columns; j++ {
			c := bool(g.grid[i][j])
			alive := g.neighbours(Point{i, j})
			// Applying rules
			if c && alive < 2 { // 1. Any live cell with fewer than two live neighbours dies
				nextGrid[i][j] = cell(false)
			} else if c && (alive == 2 || alive == 3) { // 2. Any live cell with two or three live neighbours lives on the next generation
				nextGrid[i][j] = cell(true)
			} else if c && alive > 3 { // 3. Any live cell with more than three live neighbours dies
				nextGrid[i][j] = cell(false)
			} else if !c && alive == 3 { // 4. Any dead cell with exactly three live neighbours become a live cell
				nextGrid[i][j] = cell(true)
			}
		}
	}
	g.grid = nextGrid // Next generation
}

func (g *Game) String() string {
	var s string
	for _, row := range g.grid {
		for _, cell := range row {
			s += fmt.Sprintf("%1s", cell)
		}
		s += fmt.Sprint("\n")
	}
	s += fmt.Sprint("\n")
	return s
}

// Start the game ticking sleeping 1 second between each tick
func (g *Game) Start() {
	go func() {
		for !g.stop {
			fmt.Print(g)
			g.tick()
			time.Sleep(500 * time.Millisecond)
		}
	}()
}

// Stop stops the game ticking
func (g *Game) Stop() {
	g.stop = true
}

// newGrid creates a new slice of cells of the specified size
func newGrid(rows, columns int) [][]cell {
	grid := make([][]cell, rows)
	for i := range grid {
		grid[i] = make([]cell, columns)
	}
	return grid
}

// NewGame creates and returns a new Game struct pointer with a grid of the specified size
func NewGame(rows, columns int) *Game {
	g := &Game{newGrid(rows, columns), rows, columns, false}
	return g
}
