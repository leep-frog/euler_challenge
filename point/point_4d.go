package point

import "fmt"

type Point4D struct {
	W, X, Y, Z int
}

func New4D(w, x, y, z int) *Point4D {
	return &Point4D{w, x, y, z}
}

func (p *Point4D) String() string {
	return fmt.Sprintf("(%d,%d,%d,%d)", p.W, p.X, p.Y, p.Z)
}
