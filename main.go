package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/aaronbbrown/battleship/pkg/game"
)

func main() {
	size := 10

	g := game.NewGame(size)
	g.Board1.Generate(g.Ships)

	fmt.Println("Player View")
	pv := game.NewPlayerView(g.Board1)
	fmt.Println(pv.String())

	fmt.Println("Opponent View")
	ov := game.NewOpponentView(g.Board1)
	fmt.Println(ov.String())

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		x := r.Intn(size)
		y := r.Intn(size)
		g.Board1.Fire(x, y)
	}

	fmt.Println("Player View")
	fmt.Println(pv.String())

	fmt.Println("Opponent View")
	fmt.Println(ov.String())

}
