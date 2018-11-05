package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// repeatingKeyPressed return true when key is pressed considering the repeat state.
func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 100
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}

func (g *Game) IsStartKey() bool {
	return repeatingKeyPressed(ebiten.KeyS)
}
func (g *Game) IsCreditKey() bool {
	return repeatingKeyPressed(ebiten.KeyC)
}
func (g *Game) IsBetKey() bool {
	return repeatingKeyPressed(ebiten.KeyB)
}
func (g *Game) IsHeldKey1() bool {
	return repeatingKeyPressed(ebiten.Key1)
}
func (g *Game) IsHeldKey2() bool {
	return repeatingKeyPressed(ebiten.Key2)
}
func (g *Game) IsHeldKey3() bool {
	return repeatingKeyPressed(ebiten.Key3)
}
func (g *Game) IsHeldKey4() bool {
	return repeatingKeyPressed(ebiten.Key4)
}
func (g *Game) IsHeldKey5() bool {
	return repeatingKeyPressed(ebiten.Key5)
}
