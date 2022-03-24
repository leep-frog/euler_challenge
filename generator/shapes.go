package generator

import "github.com/leep-frog/euler_challenge/maths"

type shapeNumberGenerator struct {
	shape int
	jump  int
}

func (sng *shapeNumberGenerator) Next(g *Generator[int]) int {
	if len(g.values) == 0 {
		sng.jump = 1
		return 1
	}
	sng.jump += sng.shape - 2
	return g.Last() + sng.jump
}

func ShapeNumberGenerator(n int) *Generator[int] {
	return newIntGen(&shapeNumberGenerator{n, 0})
}

func Triangulars() *Generator[int] {
	return ShapeNumberGenerator(3)
}

// t_n = n(n+1)/2
func IsTriangular(tn int) bool {
	if tn < 1 {
		return false
	}
	n2 := tn * 2
	n := maths.Sqrt(n2)
	return n*(n+1)/2 == tn
}

func Pentagonals() *Generator[int] {
	return ShapeNumberGenerator(5)
}

// t_n  = n(3n−1)/2
func IsPentagonal(tn int) bool {
	if tn < 1 {
		return false
	}

	n := maths.Sqrt((2 * tn) / 3)
	for ; n*(3*n-1)/2 < tn; n++ {
	}
	return n*(3*n-1)/2 == tn
}

func Hexagonals() *Generator[int] {
	return ShapeNumberGenerator(6)
}

// t_n  = n(2n−1) >= 2 * n * n
func IsHexagonal(tn int) bool {
	if tn < 1 {
		return false
	}

	n := maths.Sqrt((tn) / 2)
	for ; n*(2*n-1) < tn; n++ {
	}
	return n*(2*n-1) == tn
}
