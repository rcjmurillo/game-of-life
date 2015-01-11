package game

type cell bool

type point [2]int

// Game holds the game's logic and the grid of cells where the game takes place
type Game struct {
	grid [][]cell
}

// Seed set the living cells on the grid according to points
func (g *Game) Seed(points []point) {
	for _, p := range points {
		g.grid[p[0]][p[1]] = true
	}
}

// neighbours returns the count of alive and dead neighbours from the specified point
func (g *Game) neighbours(p point) (alive, dead int) {
	deltas := []point{{-1, -1}, {-1, 0}, {1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, -1}}
	for _, d := range deltas {
		x, y := p[0]+d[0], p[1]+d[1]
		if x >= 0 && x < len(g.grid) && y >= 0 && y < len(g.grid) { // Only count for valid indexes
			if g.grid[x][y] {
				alive++
			} else {
				dead++
			}
		}
	}
	return
}

// tick evaluates the rules of the game for every cell on the grid and creates a new
// grid holding the next generation of cells.
func (g *Game) tick() {

}

// newGrid creates a new slice of cells of the specified size
func newGrid(size int) [][]cell {
	grid := make([][]cell, size)
	for i := range grid {
		grid[i] = make([]cell, len(grid))
	}
	return grid
}

// NewGame creates and returns a new Game struct with a grid of the specified size
func NewGame(size int) *Game {
	g := &Game{newGrid(size)}
	return g
}
