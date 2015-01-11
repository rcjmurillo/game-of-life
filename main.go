package main

import (
	"github.com/rcjmurillo/game-of-life/game"
	"time"
)

func main() {
	g := game.NewGame(15, 15)
	g.Seed([]game.Point{{6, 7}, {7, 5}, {7, 6}, {7, 8}, {7, 9}, {8, 7}})
	g.Start()
	time.Sleep(40 * time.Second)
	g.Stop()
}
