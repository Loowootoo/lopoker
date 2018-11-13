package ui2d

import (
	"bytes"
	"image"
	"image/color"
	_ "image/png"
	"strconv"

	"github.com/Loowootoo/lopoker/ui2d/assets/pcard"

	"github.com/Loowootoo/lopoker/ui2d/assets/fonts"

	"github.com/Loowootoo/lopoker/game"

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
	tt, err := truetype.Parse(fonts.Water_ttf)
	if err != nil {
		panic(err)
	}
	const dpi = 72
	normalFont := truetype.NewFace(tt, &truetype.Options{
		Size:    20,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	bigFont := truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	return &UI2d{normalFont, bigFont, ebitenImage, bkgImage}
}

func (ui *UI2d) textWidth(str string) int {
	b, _ := font.BoundString(ui.normalFont, str)
	return (b.Max.X - b.Min.X).Ceil()
}

var (
	shadowColor  = color.NRGBA{26, 77, 22, 0x80}
	fontBaseSize = 20
)

func (ui *UI2d) DrawTextWithShadow(rt *ebiten.Image, str string, x, y, scale int, clr color.Color) {
	offsetY := fontBaseSize * scale
	y += offsetY
	text.Draw(rt, str, ui.normalFont, x+1, y+1, shadowColor)
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

func (ui *UI2d) DrawHandCard(screen *ebiten.Image, card [5]game.Card, back [5]bool) {
	op := &ebiten.DrawImageOptions{}
	for i := 0; i < 5; i++ {
		op.GeoM.Reset()
		if back[i] {
			x := 3 * cardWidth
			y := 4 * cardHeight
			srcRect := image.Rect(x, y, x+cardWidth, y+cardHeight)
			op.SourceRect = &srcRect
		} else {
			x := card[i].Number * cardWidth
			y := card[i].Pattern * cardHeight
			srcRect := image.Rect(x, y, x+cardWidth, y+cardHeight)
			op.SourceRect = &srcRect
		}
		op.GeoM.Scale(1, 1)
		op.GeoM.Translate(CardPos[i].X, CardPos[i].Y)
		screen.DrawImage(ui.cardGrp, op)
	}
}

func (ui *UI2d) DrawHandHeld(screen *ebiten.Image, card [5]bool) {
	for i := 0; i < 5; i++ {
		if card[i] {
			ui.DrawTextWithShadowCenter(screen, "HELD", int(CardPos[i].X), int(CardPos[i].Y+165), 1, color.White, 100)
		}
	}
}

func (ui *UI2d) DrawMessage(screen *ebiten.Image, msg string) {
	ui.DrawTextWithShadowCenter(screen, msg, 248, 562, 1, color.White, 308)
}
func (ui *UI2d) DrawCredit(screen *ebiten.Image, credit int) {
	s := strconv.Itoa(credit)
	ui.DrawTextWithShadowCenter(screen, s, 110, 562, 1, color.White, 110)
}
func (ui *UI2d) DrawBet(screen *ebiten.Image, bet int) {
	s := strconv.Itoa(bet)
	ui.DrawTextWithShadowCenter(screen, s, 720, 562, 1, color.White, 60)
}

var odds = [9]int{
	800, 50, 25, 9, 6, 4, 3, 2, 1,
}

func (ui *UI2d) DrawOdds(screen *ebiten.Image, bet int) {
	for i := 0; i < 9; i++ {
		s := strconv.Itoa(bet * odds[i])
		ui.DrawTextWithShadowRight(screen, s, 316, 10+i*34, 1, color.White, 60)
	}
}

func (ui *UI2d) DrawWin(screen *ebiten.Image, win int) {
	s := strconv.Itoa(win)
	ui.DrawTextWithShadowCenter(screen, "YOU WIN", 516, 120, 2, color.White, 200)
	ui.DrawTextWithShadowCenter(screen, s, 516, 200, 2, color.White, 200)
}

func (ui *UI2d) Draw(screen *ebiten.Image, game *game.Game) {
	ui.DrawBackground(screen)
	ui.DrawHandCard(screen, game.Player.Hand, game.Player.BackCard)
	ui.DrawHandHeld(screen, game.Player.Held)
	ui.DrawMessage(screen, game.Message)
	ui.DrawCredit(screen, game.Player.Credit)
	ui.DrawBet(screen, game.Player.Bet)
	ui.DrawOdds(screen, game.Player.Bet)
	if game.GameWin > 0 {
		ui.DrawWin(screen, game.GameWin)
	} else {
		game.SmokeAnim.Draw(screen)
	}
}
