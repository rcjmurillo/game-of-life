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

	g := game.NewGame(41, 41)
	g.Seed([]game.Point{
		{17, 21}, {18, 20}, {18, 22}, {19, 21},
		{19, 19}, {19, 23},
		{20, 19}, {20, 23},
		{21, 19}, {21, 23},
		{23, 21}, {22, 20}, {22, 22}, {21, 21},
	})
	g.Start()

	done := make(chan bool)
	go func() {
		<-signals
		g.Stop()
		done <- true
	}()
	<-done
}
