package ui2d

import (
	"Loowootoo/lopoker/ui2d/assets/fonts"
	"Loowootoo/lopoker/ui2d/assets/pcard"
	"bytes"
	"image"
	"image/color"
	_ "image/png"
	"strings"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

type UI2d struct {
	normalFont font.Face
	bigFont    font.Face
	cardGrp    *ebiten.Image
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
	return &UI2d{normalFont, bigFont, ebitenImage}
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
	for _, line := range strings.Split(str, "\n") {
		y += offsetY
		text.Draw(rt, line, ui.normalFont, x+2, y+2, shadowColor)
		text.Draw(rt, line, ui.normalFont, x, y, clr)
	}
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
