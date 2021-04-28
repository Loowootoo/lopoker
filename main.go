package main

import (
	"math/rand"
	"time"

	"github.com/Loowootoo/lopoker/game"
	"github.com/Loowootoo/lopoker/ui2d"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

const winWidth, winHeight int = 800, 600

var newGame *game.Game
var ui *ui2d.UI2d

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (g *Game) Update() error {
	newGame.GameLoop()
	return nil
}

func (q *Game) Draw(screen *ebiten.Image) {
	ui.Draw(screen, newGame)
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return winWidth, winHeight
}

func main() {
	newGame = game.NewGame(0)
	newGame.Shuffle()
	ui = ui2d.NewUI2d()
	ebiten.SetWindowSize(winWidth, winHeight)
	ebiten.SetWindowTitle("LoPoker !")

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
