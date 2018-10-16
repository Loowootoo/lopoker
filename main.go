package main

import (
	"Loowootoo/lopoker/game"
	"Loowootoo/lopoker/ui2d"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const winWidth, winHeight int = 800, 600

var newGame *game.Game
var ui *ui2d.UI2d

func init() {
	rand.Seed(time.Now().UnixNano())
}

func update(screen *ebiten.Image) error {
	newGame.Run()
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	for i := 0; i < 5; i++ {
		//		ui.DrawCard(screen, ui2d.CardPos[i], newGame.Player.Hand[i].GetVal())
		ui.DrawCard(screen, ui2d.CardPos[i], newGame.Player.HandSort[i].GetVal())
	}
	return nil
}

func main() {
	newGame = game.NewGame(1000)
	ui = ui2d.NewUI2d()
	err := ebiten.Run(update, winWidth, winHeight, 1, "LoPoker !!!")
	if err != nil {
		panic(err)
	}
}
