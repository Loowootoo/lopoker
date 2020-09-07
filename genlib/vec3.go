package genlib

import (
	"fmt"
	"math"
)

type Vec2 struct {
	X, Y, Z float64
}

//向量常數值
var (
	Zero     = Vec2{0, 0, 0}
	Up       = Vec2{0, 1, 0}
	Down     = Vec2{0, -1, 0}
	Left     = Vec2{-1, 0, 0}
	Right    = Vec2{1, 0, 0}
	Forward  = Vec2{0, 0, 1}
	Backward = Vec2{0, 0, -1}
	One      = Vec2{1, 1, 1}
	MinusOne = Vec2{-1, -1, -1}
)

func Roundf(val float64, places int) float64 {
	if places < 0 {
		panic("places should be >= 0")
	}

	factor := math.Pow10(places)
	val = val * factor
	tmp := float64(int(val))
	return tmp / factor
}

func Lerpf(from, to float64, t float64) float64 {
	return from + ((to - from) * t)
}

func LerpAngle(from, to float64, t float64) float64 {
	for to-from > 180 {
		from += 360
	}
	for from-to > 180 {
		to += 360
	}
	return from + ((to - from) * t)
}

func (v *Vec2) String() string {
	return fmt.Sprintf("(%f,%f,%f)", v.X, v.Y, v.Z)
}

func NewVec22(x, y float64) Vec2 {
	return Vec2{x, y, 1}
}

func NewVec23(x, y, z float64) Vec2 {
	return Vec2{x, y, z}
}

func (v *Vec2) Add(vect Vec2) Vec2 {
	return Vec2{v.X + vect.X, v.Y + vect.Y, v.Z + vect.Z}
}

func (v *Vec2) Sub(vect Vec2) Vec2 {
	return Vec2{v.X - vect.X, v.Y - vect.Y, v.Z - vect.Z}
}

func (v *Vec2) Mul(vect Vec2) Vec2 {
	return Vec2{v.X * vect.X, v.Y * vect.Y, v.Z * vect.Z}
}

func (v *Vec2) Mul2(vect float64) Vec2 {
	return Vec2{v.X * vect, v.Y * vect, v.Z * vect}
}

func (v *Vec2) Distance(vect Vec2) float64 {
	x := v.X - vect.X
	y := v.Y - vect.Y
	return math.Sqrt(float64(x*x + y*y))
}

func (v *Vec2) Div(vect Vec2) Vec2 {
	return Vec2{v.X / vect.X, v.Y / vect.Y, v.Z / vect.Z}
}

func (v *Vec2) fixAngle() {
	for v.X >= 360 {
		v.X -= 360
	}
	for v.X <= -360 {
		v.X += 360
	}

	for v.Y >= 360 {
		v.Y -= 360
	}
	for v.Y <= -360 {
		v.Y += 360
	}

	for v.Z >= 360 {
		v.Z -= 360
	}
	for v.Z <= -360 {
		v.Z += 360
	}
}

func (v *Vec2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Lerp(from, to Vec2, t float64) Vec2 {
	return NewVec22(from.X+((to.X-from.X)*t), from.Y+((to.Y-from.Y)*t))
}

func (v *Vec2) Normalize() {
	l := v.Length()
	v.X /= l
	v.Y /= l
	v.Z /= l
}

func (v *Vec2) Normalized() Vec2 {
	l := v.Length()
	if l == 0 {
		return NewVec23(0, 0, 0)
	}
	return NewVec23(v.X/l, v.Y/l, v.Z/l)
}
