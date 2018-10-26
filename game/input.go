package game

import (
	"github.com/hajimehoshi/ebiten"
)

func (g *Game) IsCreditKey() bool {
	return ebiten.IsKeyPressed(ebiten.KeyC)
}
func (g *Game) IsBetKey() bool {
	return ebiten.IsKeyPressed(ebiten.KeyB)
}
func (g *Game) IsHeldKey1() bool {
	return ebiten.IsKeyPressed(ebiten.Key1)
}
func (g *Game) IsHeldKey2() bool {
	return ebiten.IsKeyPressed(ebiten.Key2)
}
func (g *Game) IsHeldKey3() bool {
	return ebiten.IsKeyPressed(ebiten.Key3)
}
func (g *Game) IsHeldKey4() bool {
	return ebiten.IsKeyPressed(ebiten.Key4)
}
func (g *Game) IsHeldKey5() bool {
	return ebiten.IsKeyPressed(ebiten.Key5)
}
