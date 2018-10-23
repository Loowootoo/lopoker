package sprlib

import (
	"fmt"
	"math"
)

type Vector struct {
	X, Y, Z float64
}

//向量常數值
var (
	Zero     = Vector{0, 0, 0}
	Up       = Vector{0, 1, 0}
	Down     = Vector{0, -1, 0}
	Left     = Vector{-1, 0, 0}
	Right    = Vector{1, 0, 0}
	Forward  = Vector{0, 0, 1}
	Backward = Vector{0, 0, -1}
	One      = Vector{1, 1, 1}
	MinusOne = Vector{-1, -1, -1}
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

func (v *Vector) String() string {
	return fmt.Sprintf("(%f,%f,%f)", v.X, v.Y, v.Z)
}

func NewVector2(x, y float64) Vector {
	return Vector{x, y, 1}
}

func NewVector3(x, y, z float64) Vector {
	return Vector{x, y, z}
}

func (v *Vector) Add(vect Vector) Vector {
	return Vector{v.X + vect.X, v.Y + vect.Y, v.Z + vect.Z}
}

func (v *Vector) Sub(vect Vector) Vector {
	return Vector{v.X - vect.X, v.Y - vect.Y, v.Z - vect.Z}
}

func (v *Vector) Mul(vect Vector) Vector {
	return Vector{v.X * vect.X, v.Y * vect.Y, v.Z * vect.Z}
}

func (v *Vector) Mul2(vect float64) Vector {
	return Vector{v.X * vect, v.Y * vect, v.Z * vect}
}

func (v *Vector) Distance(vect Vector) float64 {
	x := v.X - vect.X
	y := v.Y - vect.Y
	return math.Sqrt(float64(x*x + y*y))
}

func (v *Vector) Div(vect Vector) Vector {
	return Vector{v.X / vect.X, v.Y / vect.Y, v.Z / vect.Z}
}

func (v *Vector) fixAngle() {
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

func (v *Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Lerp(from, to Vector, t float64) Vector {
	return NewVector2(from.X+((to.X-from.X)*t), from.Y+((to.Y-from.Y)*t))
}

func (v *Vector) Normalize() {
	l := v.Length()
	v.X /= l
	v.Y /= l
	v.Z /= l
}

func (v *Vector) Normalized() Vector {
	l := v.Length()
	if l == 0 {
		return NewVector3(0, 0, 0)
	}
	return NewVector3(v.X/l, v.Y/l, v.Z/l)
}
