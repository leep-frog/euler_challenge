package point

import "fmt"

type Point struct {
	X, Y, Z int
}

func rotateFunc(x, y, z int, rs ...Rotation) func(*Point) *Point {
	return func(p *Point) *Point {
		return p.Rotate(x, y, z, rs...)
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
