package ui2d

import (
	"Loowootoo/lopoker/game"
	"Loowootoo/lopoker/sprlib"
	"Loowootoo/lopoker/ui2d/assets/fonts"
	"Loowootoo/lopoker/ui2d/assets/pcard"
	"bytes"
	"image"
	"image/color"
	_ "image/png"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

type UI2d struct {
	normalFont font.Face
	bigFont    font.Face
	cardGrp    *ebiten.Image
	bkground   *ebiten.Image
	smokeAnim  *sprlib.Sprite
}

type Pos struct {
	X, Y float64
}

var CardPos = [5]Pos{
	{90, 330},
	{220, 330},
	{350, 330},
	{480, 330},
	{610, 330},
}

func NewUI2d() *UI2d {
	img, _, err := image.Decode(bytes.NewReader(pcard.PCardPNG))
	if err != nil {
		panic(err)
	}
	ebitenImage, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	img, _, err = image.Decode(bytes.NewReader(pcard.BkgPNG))
	if err != nil {
		panic(err)
	}

	bkgImage, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	//
	smokeAnim := sprlib.NewSprite()
	smokeAnim.AddAnimFrameFromBytes("default", pcard.SmokePNG, 15000, 15, ebiten.FilterDefault)
	smokeAnim.CenterCoordonnates = false
	smokeAnim.Animated = true
	smokeAnim.Pos = sprlib.Vector{448, 25, 0}
	smokeAnim.Speed = 0
	smokeAnim.Direction = sprlib.Vector{0, 0, 0}
	//
	tt, err := truetype.Parse(fonts.Water_ttf)
	if err != nil {
		panic(err)
	}
	const dpi = 72
	normalFont := truetype.NewFace(tt, &truetype.Options{
		Size:    16,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	bigFont := truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	return &UI2d{normalFont, bigFont, ebitenImage, bkgImage, smokeAnim}
}

func (ui *UI2d) textWidth(str string) int {
	b, _ := font.BoundString(ui.normalFont, str)
	return (b.Max.X - b.Min.X).Ceil()
}

var (
	shadowColor  = color.NRGBA{0, 0, 0, 0x80}
	fontBaseSize = 16
)

func (ui *UI2d) DrawTextWithShadow(rt *ebiten.Image, str string, x, y, scale int, clr color.Color) {
	offsetY := fontBaseSize * scale
	y += offsetY
	text.Draw(rt, str, ui.normalFont, x+2, y+2, shadowColor)
	text.Draw(rt, str, ui.normalFont, x, y, clr)
}

func (ui *UI2d) DrawTextWithShadowCenter(rt *ebiten.Image, str string, x, y, scale int, clr color.Color, width int) {
	w := ui.textWidth(str) * scale
	x += (width - w) / 2
	ui.DrawTextWithShadow(rt, str, x, y, scale, clr)
}

func (ui *UI2d) DrawTextWithShadowRight(rt *ebiten.Image, str string, x, y, scale int, clr color.Color, width int) {
	w := ui.textWidth(str) * scale
	x += width - w
	ui.DrawTextWithShadow(rt, str, x, y, scale, clr)
}

func (ui *UI2d) DrawBackground(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	op.GeoM.Translate(0, 0)
	screen.DrawImage(ui.bkground, op)
}

const cardWidth = 100
const cardHeight = 145

func (ui *UI2d) DrawCard(screen *ebiten.Image, pos Pos, num int) {
	op := &ebiten.DrawImageOptions{}
	x := (num % 13) * cardWidth
	y := (num / 13) * cardHeight
	srcRect := image.Rect(x, y, x+cardWidth, y+cardHeight)
	op.SourceRect = &srcRect
	op.GeoM.Scale(1, 1)
	op.GeoM.Translate(pos.X, pos.Y)
	screen.DrawImage(ui.cardGrp, op)
}

func (ui *UI2d) DrawHandCard(screen *ebiten.Image, card [5]game.Card) {
	op := &ebiten.DrawImageOptions{}
	for i := 0; i < 5; i++ {
		op.GeoM.Reset()
		x := card[i].Number * cardWidth
		y := card[i].Pattern * cardHeight
		srcRect := image.Rect(x, y, x+cardWidth, y+cardHeight)
		op.SourceRect = &srcRect
		op.GeoM.Scale(1, 1)
		op.GeoM.Translate(CardPos[i].X, CardPos[i].Y)
		screen.DrawImage(ui.cardGrp, op)
	}
}

func (ui *UI2d) DrawHandHeld(screen *ebiten.Image, card [5]bool) {
	for i := 0; i < 5; i++ {
		if card[i] {
			ui.DrawTextWithShadowCenter(screen, "HELD", int(CardPos[i].X), int(CardPos[i].Y+150), 1, color.White, 100)
		}
	}
}

func (ui *UI2d) DrawMessage(screen *ebiten.Image, msg string) {
	ui.DrawTextWithShadowCenter(screen, msg, 0, 520, 1, color.White, 600)
}

func (ui *UI2d) Draw(screen *ebiten.Image, game *game.Game) {
	ui.DrawBackground(screen)
	ui.DrawHandCard(screen, game.Player.Hand)
	ui.DrawHandHeld(screen, game.Player.Held)
	ui.DrawMessage(screen, game.Message)
	ui.smokeAnim.Draw(screen)
}
