package eulerchallenge

import (
	"fmt"
	"strconv"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/bfs"
)

func P127() *problem {
	return intInputNode(127, func(o command.Output, n int) {
		g := generator.Primes()
		var sum int
		for c := 1; c < n; c++ {
			if c % 1000 == 0 {
				fmt.Println(c)
			}
			cRad := calcRadical(c, g)

			smallesPrimes := []int{}
			for i := 0; len(smallesPrimes) < 2; i++ {
				if c % g.Nth(i) != 0 {
					smallesPrimes = append(smallesPrimes, g.Nth(i))
				}
			}
			// a and b will have a radical of at least 6 (smallest two primes are 2*3=6)
			if cRad*smallesPrimes[0]*smallesPrimes[1] >= c {
				continue
			}
			for a := 1; a < c - a; a++ {
				b := c - a
				// rad (abc) = rad(a) + rad(b) + rad(c) since GCD(a, b) = GCD(a, c) = GCD(b, c) = 1
				aRad, bRad := calcRadical(a, g), calcRadical(b, g) 
				if cRad + aRad + bRad >= c {
					continue
				}
				if generator.Coprimes(aRad, bRad, g) || generator.Coprimes(aRad, cRad, g) || generator.Coprimes(bRad, cRad, g) {
					continue
				}				
				//fmt.Println(a, b, c)
				//_ = b
				//calcRadical(a, g)
				/*pnc = true
				_ = radicalCache[a]
				//
				pnc = false*/
				/*if cRad - 10 - 4 > c {
					continue
				}*/
				/*aRad, bRad := calcRadical(a, g), calcRadical(b, g)
				if aRad + bRad + cRad > c {
					continue
				}*/
				
				/*generator.PrimeFactors(a, g)
				generator.PrimeFactors(b, g)*/
			}
		}
		/*for a := 1; a < n/2; a++ {
			ra := calcRadical(a, g)
			for b := maths.Max(2*ra, a+1); a+b < n; b++ {
				c := a + b
				if generator.Coprimes(a, b, g) || generator.Coprimes(a, c, g) || generator.Coprimes(b, c, g) {
					continue
				}
				r := calcRadical(a*b*c, g)
				//r := newRadical(a*b*c, g)
				if r >= c {
					continue
				}
				fmt.Println(a, b, c)
				sum += c
			}
		}*/
		/*for c := 1; c < n; c++ {
			if calcRadical(c, g) == c {
				continue
			}
			for b := c - 1; b > c-b; b-- {
				a := c - b
				if calcRadical(a*b*c, g) >= c {
					continue
				}
				if generator.Coprimes(a, b, g) || generator.Coprimes(a, c, g) || generator.Coprimes(b, c, g) {
					continue
				}
				fmt.Println(a, b, c)
				sum += c
			}
		}*/
		// c has to
		fmt.Println(sum, sum/2)
	})
}

// search for a
type node127 struct {
	primeFactors map[int]int
	a            int
}

type context127 struct {
	g *generator.Generator[int]
	c int
	cFactors map[int]int
}

func (n *node127) Code(*context127, bfs.DFSPath[*node127]) string {
	return strconv.Itoa(n.a)
}

func (n *node127) Done(*context127, bfs.DFSPath[*node127]) bool {
	return false
}

/*func (n *node127) AdjacentStates(ctx *context127, path bfs.DFSPath[*node127]) int {
	var r []*node127
	for i := 0;; i++ {
		pi := ctx.g.Nth(i)
		if n.primeFactors[pi] {

		} else {

		}
	}
}*/
