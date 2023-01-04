package point

import (
	"fmt"

	"github.com/leep-frog/euler_challenge/maths"
)

// TODO: let this be generic type
type Point3D struct {
	X, Y, Z int
}

func Origin3D() *Point3D {
	return New3D(0, 0, 0)
}

func New3D(x, y, z int) *Point3D {
	return &Point3D{x, y, z}
}

func rotateFunc(x, y, z int, rs ...Rotation) func(*Point3D) *Point3D {
	return func(p *Point3D) *Point3D {
		return p.Rotate(x, y, z, rs...)
	}
}

func (p *Point3D) Code() string {
	return p.String()
}

func (p *Point3D) Minus(s *Point3D) *Point3D {
	return New3D(p.X-s.X, p.Y-s.Y, p.Z-s.Z)
}

func (p *Point3D) Plus(s *Point3D) *Point3D {
	return New3D(p.X+s.X, p.Y+s.Y, p.Z+s.Z)
}

func (p *Point3D) Cross(s *Point3D) *Point3D {
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

	return &Point3D{
		int(maths.Determinant(maths.BiggifyIntMatrix(iMatrix)).Num().Int64()),
		-int(maths.Determinant(maths.BiggifyIntMatrix(jMatrix)).Num().Int64()),
		int(maths.Determinant(maths.BiggifyIntMatrix(kMatrix)).Num().Int64()),
	}
}

func RotFuncsByPoint3D(p *Point3D) []func(*Point3D) *Point3D {
	return RotFuncs3D(p.X, p.Y, p.Z)
}

func RotFuncs3D(x, y, z int) []func(*Point3D) *Point3D {
	return []func(*Point3D) *Point3D{
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

func (p *Point3D) Offset(x, y, z int) *Point3D {
	p.X += x
	p.Y += y
	p.Z += z
	return p
}

func (p *Point3D) String() string {
	return fmt.Sprintf("(%d,%d,%d)", p.X, p.Y, p.Z)
}

func (p *Point3D) RotateX(x, y, z int) *Point3D {
	return New3D(p.X, y+(p.Z-z), z+(y-p.Y))
}

func (p *Point3D) RotateY(x, y, z int) *Point3D {
	return New3D(x+(z-p.Z), p.Y, z+(p.X-x))
}

func (p *Point3D) RotateZ(x, y, z int) *Point3D {
	return New3D(x+(p.Y-y), y+(x-p.X), p.Z)
}

func (p *Point3D) Copy() *Point3D {
	return &Point3D{p.X, p.Y, p.Z}
}

func (p *Point3D) Rotate(x, y, z int, rs ...Rotation) *Point3D {
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
