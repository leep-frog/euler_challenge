package point

import (
	"fmt"

	"github.com/leep-frog/euler_challenge/maths"
)

// TODO: let this be generic type
type Point struct {
	X, Y, Z int
}

func Origin() *Point {
	return New(0, 0, 0)
}

func New(x, y, z int) *Point {
	return &Point{x, y, z}
}

func rotateFunc(x, y, z int, rs ...Rotation) func(*Point) *Point {
	return func(p *Point) *Point {
		return p.Rotate(x, y, z, rs...)
	}
}

func (p *Point) Code() string {
	return p.String()
}

func (p *Point) Minus(s *Point) *Point {
	return New(p.X-s.X, p.Y-s.Y, p.Z-s.Z)
}

func (p *Point) Cross(s *Point) *Point {
	iMatrix := [][]int{
		{p.Y, p.Z},
		{s.Y, s.Z},
	}
	jMatrix := [][]int{
		{p.X, p.Z},
		{s.X, s.Z},
	}
	kMatrix := [][]int{
		{p.X, p.Y},
		{s.X, s.Y},
	}

	return &Point{
		int(maths.Determinant(maths.BiggifyIntMatrix(iMatrix)).Num().Int64()),
		-int(maths.Determinant(maths.BiggifyIntMatrix(jMatrix)).Num().Int64()),
		int(maths.Determinant(maths.BiggifyIntMatrix(kMatrix)).Num().Int64()),
	}
}

func RotFuncsByPoint(p *Point) []func(*Point) *Point {
	return RotFuncs(p.X, p.Y, p.Z)
}

func RotFuncs(x, y, z int) []func(*Point) *Point {
	return []func(*Point) *Point{
		// Rotate around X axis
		rotateFunc(x, y, z),
		rotateFunc(x, y, z, XRot),
		rotateFunc(x, y, z, XRot, XRot),
		rotateFunc(x, y, z, XRot, XRot, XRot),

		// Rotate around negative X axis
		rotateFunc(x, y, z, YRot, YRot),
		rotateFunc(x, y, z, YRot, YRot, XRot),
		rotateFunc(x, y, z, YRot, YRot, XRot, XRot),
		rotateFunc(x, y, z, YRot, YRot, XRot, XRot, XRot),

		// Face Y axis and rotate around
		rotateFunc(x, y, z, ZRot),
		rotateFunc(x, y, z, ZRot, YRot),
		rotateFunc(x, y, z, ZRot, YRot, YRot),
		rotateFunc(x, y, z, ZRot, YRot, YRot, YRot),

		// Face negative Y axis and rotate around
		rotateFunc(x, y, z, ZRot, ZRot, ZRot),
		rotateFunc(x, y, z, ZRot, ZRot, ZRot, YRot),
		rotateFunc(x, y, z, ZRot, ZRot, ZRot, YRot, YRot),
		rotateFunc(x, y, z, ZRot, ZRot, ZRot, YRot, YRot, YRot),

		// Face Z axis and rotate around
		rotateFunc(x, y, z, YRot),
		rotateFunc(x, y, z, YRot, ZRot),
		rotateFunc(x, y, z, YRot, ZRot, ZRot),
		rotateFunc(x, y, z, YRot, ZRot, ZRot, ZRot),

		// Face negative Y axis and rotate around
		rotateFunc(x, y, z, YRot, YRot, YRot),
		rotateFunc(x, y, z, YRot, YRot, YRot, ZRot),
		rotateFunc(x, y, z, YRot, YRot, YRot, ZRot, ZRot),
		rotateFunc(x, y, z, YRot, YRot, YRot, ZRot, ZRot, ZRot),
	}
}

func (p *Point) Offset(x, y, z int) *Point {
	p.X += x
	p.Y += y
	p.Z += z
	return p
}

func (p *Point) String() string {
	return fmt.Sprintf("%d,%d,%d", p.X, p.Y, p.Z)
}

func NewPoint(x, y, z int) *Point {
	return &Point{x, y, z}
}

func (p *Point) RotateX(x, y, z int) *Point {
	return NewPoint(p.X, y+(p.Z-z), z+(y-p.Y))
}

func (p *Point) RotateY(x, y, z int) *Point {
	return NewPoint(x+(z-p.Z), p.Y, z+(p.X-x))
}

func (p *Point) RotateZ(x, y, z int) *Point {
	return NewPoint(x+(p.Y-y), y+(x-p.X), p.Z)
}

func (p *Point) Copy() *Point {
	return &Point{p.X, p.Y, p.Z}
}

func (p *Point) Rotate(x, y, z int, rs ...Rotation) *Point {
	if len(rs) == 0 {
		return p.Copy()
	}

	ret := p
	for _, r := range rs {
		switch r {
		case XRot:
			ret = ret.RotateX(x, y, z)
		case YRot:
			ret = ret.RotateY(x, y, z)
		case ZRot:
			ret = ret.RotateZ(x, y, z)
		}
	}
	return ret
}

type Rotation int

const (
	XRot Rotation = iota
	YRot
	ZRot
)
