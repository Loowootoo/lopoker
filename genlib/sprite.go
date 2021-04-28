package genlib

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type AnimFrame struct {
	Image        *ebiten.Image
	MaxFrames    int
	CurrFrame    int
	FrameWidth   int
	FrameHeight  int
	FrameCounter *TimeCounter
	RunOnce      bool
}

/*
func newAnimFrameFromFile(fileName string, duration int, frames int, filter ebiten.Filter) *AnimFrame {
	var err error
	animFrame := new(AnimFrame)
	animFrame.Image, _, err = ebitenutil.NewImageFromFile(fileName, filter)
	if err != nil {
		panic(err)
	}
	animFrame.MaxFrames = frames

	width, height := animFrame.Image.Size()
	animFrame.FrameWidth = width / animFrame.MaxFrames
	animFrame.FrameHeight = height

	animFrame.FrameCounter = NewCounter(duration)
	animFrame.RunOnce = false
	return animFrame
}
*/
func newAnimFrameFromBytes(data []byte, duration int, frames int) *AnimFrame {
	animFrame := new(AnimFrame)
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	animFrame.Image = ebiten.NewImageFromImage(img)

	animFrame.MaxFrames = frames

	width, height := animFrame.Image.Size()
	animFrame.FrameWidth = width / animFrame.MaxFrames
	animFrame.FrameHeight = height

	animFrame.FrameCounter = NewCounter(duration)
	animFrame.RunOnce = false
	return animFrame
}

func (animFrame *AnimFrame) SetFrameDuration(duration int) {
	animFrame.FrameCounter.ResetCounter(duration)
}

type Sprite struct {
	// Animation label currently displayed
	CurrAnimFrame string
	// Array of animations
	AnimFrames         map[string]*AnimFrame
	Pos                Vec2
	Direction          Vec2
	Speed              float64
	Alpha              float64
	Scale              float64
	Visible            bool
	Animated           bool
	CenterCoordonnates bool
}

func NewSprite() *Sprite {
	sprite := new(Sprite)
	sprite.CurrAnimFrame = "default"
	sprite.AnimFrames = make(map[string]*AnimFrame)
	sprite.Alpha = 1
	sprite.Animated = false
	sprite.CenterCoordonnates = true
	sprite.Direction = Vec2{0, 0, 0}
	sprite.Pos = Vec2{0, 0, 0}
	sprite.Speed = 1
	sprite.Visible = true
	return sprite
}

/*
func (sprite *Sprite) AddAnimFrameFromFile(label string, path string, duration int, steps int, filter ebiten.Filter) {
	sprite.AnimFrames[label] = newAnimFrameFromFile(path, duration, steps, filter)
}
*/
func (sprite *Sprite) AddAnimFrameFromBytes(label string, data []byte, duration int, steps int) {
	sprite.AnimFrames[label] = newAnimFrameFromBytes(data, duration, steps)
}
func (sprite *Sprite) GetScale() float64 {
	return (sprite.Pos.Z/100 + 1) / 2
}

func (sprite *Sprite) Update() {
	sprite.Pos.X += sprite.Speed * sprite.Direction.X
	sprite.Pos.Y += sprite.Speed * sprite.Direction.Y
	if sprite.Pos.Z <= 0 {
		sprite.Direction.Z = 1
	} else if sprite.Pos.Z >= 100 {
		sprite.Direction.Z = -1
	}
	sprite.Pos.Z += sprite.Speed * sprite.Direction.Z
	sprite.Scale = sprite.GetScale()
}
func (sprite *Sprite) nextFrame() {
	currAnimFrame := sprite.AnimFrames[sprite.CurrAnimFrame]
	if sprite.Animated {
		if currAnimFrame.FrameCounter.TimeUp() {
			currAnimFrame.CurrFrame++
			if currAnimFrame.CurrFrame+1 > currAnimFrame.MaxFrames {
				currAnimFrame.RunOnce = true
				currAnimFrame.CurrFrame = 0
			}
		}
	}
}

//Draw calculates new coordonnates and draw the sprite on the screen, after drawing, go to the next step of animation
func (sprite *Sprite) Draw(surface *ebiten.Image) {
	if sprite.Visible {
		currAnimFrame := sprite.AnimFrames[sprite.CurrAnimFrame]
		options := &ebiten.DrawImageOptions{}

		if sprite.CenterCoordonnates {
			options.GeoM.Translate(-float64(currAnimFrame.FrameWidth)/2, -float64(currAnimFrame.FrameHeight)/2)
		}
		options.GeoM.Scale(sprite.Scale, sprite.Scale)
		options.GeoM.Translate(sprite.Pos.X, sprite.Pos.Y)
		options.ColorM.Scale(1, 1, 1, sprite.Alpha)
		x0 := currAnimFrame.CurrFrame * currAnimFrame.FrameWidth
		x1 := x0 + currAnimFrame.FrameWidth
		r := image.Rect(x0, 0, x1, currAnimFrame.FrameHeight)
		surface.DrawImage(currAnimFrame.Image.SubImage(r).(*ebiten.Image), options)
		sprite.nextFrame()
	}
}

func (sprite *Sprite) Stop() {
	sprite.Animated = false
	sprite.Visible = false
}

func (sprite *Sprite) Reset() {
	sprite.Animated = false
	sprite.Visible = false
	sprite.AnimFrames[sprite.CurrAnimFrame].CurrFrame = 0
}

func (sprite *Sprite) Start() {
	sprite.Animated = true
	sprite.Visible = true
}
