package main

import (
	"github.com/rcjmurillo/game-of-life/game"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	g := game.NewGame(15, 15)
	g.Seed([]game.Point{{6, 7}, {7, 5}, {7, 6}, {7, 8}, {7, 9}, {8, 7}})
	g.Start()

	done := make(chan bool)
	go func() {
		<-signals
		g.Stop()
		done <- true
	}()
	<-done
}
