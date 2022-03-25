package generator

import "fmt"

type RightTriangle struct {
	A, B, C, M, N int
}

func (rt *RightTriangle) Perimeter() int {
	return rt.A + rt.B + rt.C
}

func (rt *RightTriangle) String() string {
	return fmt.Sprintf("(A=%d, B=%d, C=%d)", rt.A, rt.B, rt.C)
}

func (rt *RightTriangle) Less(that *RightTriangle) bool {
	if rt.M != that.M {
		return rt.M < that.M
	}
	return rt.N < that.N
}

func RightTriangleGenerator() *Generator[*RightTriangle] {
	return newCustomGen[*RightTriangle](&triangleGenerator{1, 1, Primes()})
}

type triangleGenerator struct {
	m int
	n int
	g *Generator[int]
}

func (tg *triangleGenerator) Next(g *Generator[*RightTriangle]) *RightTriangle {
	// https://en.wikipedia.org/wiki/Pythagorean_triple
	// a = m^2 - n^2
	// b = 2mn
	// c = m^2 + n^2
	// L = 2m^2 + 2mn
	for ; ; tg.m++ {
		for ; tg.n < tg.m; tg.n++ {
			if tg.n%2 == 1 && tg.m%2 == 1 {
				continue
			}
			if tg.n > 1 && Coprimes(tg.m, tg.n, tg.g) {
				continue
			}
			r := &RightTriangle{
				tg.m*tg.m - tg.n*tg.n,
				2 * tg.m * tg.n,
				tg.m*tg.m + tg.n*tg.n,
				tg.m,
				tg.n,
			}
			tg.n++
			return r
		}
		tg.n = 1
	}
}
