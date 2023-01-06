package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/functional"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
	"golang.org/x/exp/maps"
)

func getSums(v, n int, p *generator.Prime) *maths.Int {
	sm := maths.Zero()
	for i := 1; i*v <= n; i++ {
		mod4To1Factor := func(factor int) bool { return factor%4 == 1 }
		if functional.None(maps.Keys(p.PrimeFactors(i)), mod4To1Factor) {
			sm = sm.PlusInt(i * v)
		}
	}
	return sm
}

func elegant233(n int, p *generator.Prime) *maths.Int {
	n = maths.Pow(10, n)
	upTo := (n / (5 * 5 * 5 * 13 * 13)) + 1

	var valid []int
	for i := 0; p.Nth(i) <= upTo; i++ {
		prime := p.Nth(i)
		if prime%4 == 1 {
			valid = append(valid, prime)
		}
	}

	// Power solutions:
	// [1, 2, 3] --> (a^1 * b^2 * c^3)
	// [7, 3]
	// [10, 2]

	// [1, 2, 3]
	sum := maths.Zero()
	for _, cv := range valid {
		cube := cv * cv * cv

		if cube > (n/125)+1 {
			break
		}
		for _, sv := range valid {
			if cv == sv {
				continue
			}
			square := sv * sv

			if cube*square > (n/5)+1 {
				break
			}
			for _, dot := range valid {
				if cv == dot || sv == dot {
					continue
				}
				v := cube * square * dot
				if v > n {
					break
				}
				sum = sum.Plus(getSums(v, n, p))
			}
		}
	}

	// [7, 3]
	for _, p7 := range valid {
		seven := maths.Pow(p7, 7)
		if seven > n {
			break
		}
		for _, p3 := range valid {
			if p3 == p7 {
				continue
			}
			v := seven * maths.Pow(p3, 3)
			if v > n {
				break
			}
			sum = sum.Plus(getSums(v, n, p))
		}
	}

	// [10, 2]
	for _, p10 := range valid {
		seven := maths.Pow(p10, 10)
		if seven > n {
			break
		}
		for _, p2 := range valid {
			if p2 == p10 {
				continue
			}
			v := seven * maths.Pow(p2, 2)
			if v > n {
				break
			}
			sum = sum.Plus(getSums(v, n, p))
		}
	}
	return sum
}

func P233() *problem {
	return intInputNode(233, func(o command.Output, n int) {
		o.Stdoutln(elegant233(n, generator.Primes()))
	}, []*execution{
		{
			args:     []string{"11"},
			want:     "271204031455541309",
			estimate: 10,
		},
	})
}

/* Invalid attempts below

var (
	cash    = map[int]int{}
	triCash = map[int]int{}
)

// BFS by largest thing
type node233 struct {
	k int
	f int
}

type ctx233 struct {
	target int
	primes *generator.Generator[int]
	max    int
	m      map[int]bool
}

func (n *node233) Code(*ctx233) string {
	return n.String()
}

func (n *node233) String() string {
	return fmt.Sprintf("%d", n.k)
}

func (n *node233) Done(ctx *ctx233) bool {
	if n.f == ctx.target {
		fmt.Println("BULLSEYE", n, n.f)
	}
	return false
}

func (n *node233) Distance(ctx *ctx233) bfs.Int {
	return bfs.Int(-n.f)
}

func (n *node233) AdjacentStates(ctx *ctx233) []*node233 {
	var r []*node233
	for g, p := ctx.primes.Start(0); n.k*p < ctx.max; p = g.Next() {

		v := n.k * p
		a := &node233{v, bruteAllValues(v, ctx.primes)}
		if a.f <= ctx.target {
			r = append(r, a)
		}
	}
	return r
}

// Calculate all right triangles
func bruteAllValues(n int, p *generator.Generator[int]) int {
	if v, ok := cash[n]; ok {
		return v
	}
	// 45 degree right triangle with side n, n, sqrt(2)n
	// middle is n/2

	if n%2 == 0 {
		var cnt int
		for a := 0; a < n/2; a++ {
			// a^2 + b^2 = n^2/2
			// b^2 = n^2/2 - a^2
			if maths.IsSquare(n*n/2 - a*a) {
				// x, y := (n/2 - a), n/2-maths.Sqrt(n*n/2-a*a)
				// fmt.Println(x, y)
				cnt++
			}
		}
		cash[n] = 8*cnt + 4
		return 8*cnt + 4
	}

	var cnt int
	m := fraction.New(n, 2)
	for a := fraction.New(1, 2); a.LT(m); a = a.Plus(fraction.New(1, 1)) {
		// a^2 + b^2 = n^2/2
		// b^2 = n^2/2 - a^2

		// Let A = 2a and B = 2b
		// (B/2)^2 = n^2/2 - (A/2)^2
		// B^2/4 = n^2/2 - A^2/4
		// B^2/4 = 2*n^2/4 - A^2/4
		// B^2 = 2*n^2 - A^2
		// B and A both need to be odd integers
		if maths.IsSquare(2*n*n - a.N*a.N) {
			cnt++
		}
	}
	cash[n] = 8*cnt + 4
	return 8*cnt + 4
	/*for a := 0; a < n/2; a++ {
	// We want a = X.5 and b = Y.5
	// Therefore (2a and 2b need to be odd integers)
	// Let A = 2a, B = 2b
	// A^2 + B^2 = (2n^2)/2
	// A^2 + B^2 = 4n^2/2
	// A^2 + B^2 = 2n^2
	// B^2 = 2n^2 - A^2
	// B^2 = maths.Sqrt(2n^2 - A^2)

	// 2*b^2 = 2*n^2/2 - 2*a^2
	// 2*b^2 = n^2 - 2*a^2
	// 2*b = maths.Sqrt(n^2 - 2*a^2)
	if maths.IsSquare(n*n - 2*a*a) {
		// x, y := (n/2 - a), n/2-maths.Sqrt(n*n/2-a*a)
		// fmt.Println(x, y)
		cnt++
	}
	return 8*cnt + 4 */ /*
}

// TODO: maths.Cache function
// Calculate all right triangles
func rightTris(n int, p *generator.Generator[int], ignoreCoprimes bool) int {
	if v, ok := triCash[n]; ok {
		return v
	}

	// N is even
	if n%2 == 0 {
		var cnt int
		for a := 1; a < n/2; a++ {
			// a^2 + b^2 = n^2/2
			// b^2 = n^2/2 - a^2
			if maths.IsSquare(n*n/2 - a*a) {
				b := maths.Sqrt(n*n/2 - a*a)
				cp := !generator.Coprimes(a, b, p)
				if cp || ignoreCoprimes {
					// fmt.Println("EVEN", n, a, b, cp)
					cnt++
				}
			}
		}
		cash[n] = cnt
		return cnt
	}

	// N is odd
	var cnt int
	m := fraction.New(n, 2)
	for a := fraction.New(1, 2); a.LT(m); a = a.Plus(fraction.New(1, 1)) {
		// a^2 + b^2 = n^2/2
		// b^2 = n^2/2 - a^2

		// Let A = 2a and B = 2b
		// (B/2)^2 = n^2/2 - (A/2)^2
		// B^2/4 = n^2/2 - A^2/4
		// B^2/4 = 2*n^2/4 - A^2/4
		// B^2 = 2*n^2 - A^2
		// B and A both need to be odd integers
		if maths.IsSquare(2*n*n - a.N*a.N) {
			b := maths.Sqrt(2*n*n - a.N*a.N)
			cp := !generator.Coprimes(a.N, b, p)
			if cp || ignoreCoprimes {
				//fmt.Println("ODD", n, a.N, b, cp)
				cnt++
			}
		}
	}
	cash[n] = cnt
	return cnt
}

// Only for odd n
func bruteNumValuesOdd(n int, primes *generator.Generator[int]) {
	for a := 1; a < n; a++ {
		if maths.IsSquare(2*n*n - a*a) {
			b := maths.Sqrt(2*n*n - a*a)
			if !generator.Coprimes(a, b, primes) {
				fmt.Println("YUP", n, a, maths.Sqrt(2*n*n-a*a))
			}
		}
	}
} */
