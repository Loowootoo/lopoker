package main

import (
	"Loowootoo/lopoker/game"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const winWidth, winHeight int = 1280, 720

func init() {
	rand.Seed(time.Now().UnixNano())
}

func update(screen *ebiten.Image) error {

	if ebiten.IsDrawingSkipped() {
		return nil
	}
	return nil
}

func main() {
	game := game.NewGame(1000)
	game.Run()
	err := ebiten.Run(update, winWidth, winHeight, 1, "LoPoker !!!")
	if err != nil {
		panic(err)
	}
}
