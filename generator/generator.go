package generator

import (
	"strconv"
	"fmt"
	"strings"
	"log"

	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/command/cache"
)

const (
	primesName = "primes"
	fibName = "fibonaccis"
	triName = "triangulars"
)

type Generatable[T any] interface {
	LTE(T, T) bool
	String(T) string
	FromString(string) T
}

func newBigGeneratable() Generatable[*maths.Int] {
	return &bigGeneratable{}
}

type bigGeneratable struct {}

func (bg *bigGeneratable) LTE(this, that *maths.Int) bool {
	return this.LTE(that)
}

func (bg *bigGeneratable) String(i *maths.Int) string {
	return i.String()
}

func (bg *bigGeneratable) FromString(s string)*maths.Int {
	return maths.MustIntFromString(s)
}

func newIntGeneratable() Generatable[int] {
	return &intGeneratable{}
}

type intGeneratable struct {}

func (ig *intGeneratable) LTE(this, that int) bool {
	return this <= that
}

func (ig *intGeneratable) String(i int) string {
	return strconv.Itoa(i)
}

func (ig *intGeneratable) FromString(s string) int {
	return parse.Atoi(s)
}

type Generator[T any] struct {
	name string

	values []T
	set    map[string]bool

	g Generatable[T]

	f   func(*Generator[T]) T
	idx int
}

func (g *Generator[T]) Reset() {
	g.idx = 0
}

func (g *Generator[T]) last() T {
	return g.values[len(g.values)-1]
}

func (g *Generator[T]) len() int {
	return len(g.values)
}

func (g *Generator[T]) Nth(i int) T {
	for g.len() <= i {
		g.getNext()
	}
	return g.values[i]
}

func (g *Generator[T]) Next() T {
	g.idx++
	return g.Nth(g.idx - 1)
}

func (g *Generator[T]) getNext() T {
	i := g.f(g)
	g.values = append(g.values, i)
	s := g.g.String(i)
	g.set[s] = true
	//putCache(g.name, g.values, g.g)
	return i
}

var (
	newCache = func() *cache.Cache{
		return cache.NewCache()
	}
)


func (g *Generator[T]) Contains(t T) bool {
	for ; g.len() == 0 || g.g.LTE(g.last(), t); g.getNext() {
	}
	return g.set[g.g.String(t)]
}

func getFromCache(name string) []string {
	c := newCache()
	name = fmt.Sprintf("coding_challenge_%s", name)
	s, err := c.Get(name)
	if err != nil {
		log.Fatalf(fmt.Sprintf("failed to get cache: %v", err))
	}
	sl := strings.Split(s, "\n")
	if len(sl) > 0 && sl[len(sl) - 1] == "" {
		sl = sl[:len(sl)-1]
	}
	return sl
}

func putCache[T any](name string, sl []T, g Generatable[T]) {
	var r []string
	for _, s := range sl {
		r = append(r, g.String(s))
	}
	c := newCache()
	name = fmt.Sprintf("coding_challenge_%s", name)
	if err := c.Put(name, strings.Join(r, "\n")); err != nil {
		panic(fmt.Sprintf("failed to write to struct: %v", err))
	}
}

func NewGenerator[T any](name string, start T, g Generatable[T], f func(*Generator[T]) T) *Generator[T] {
	gt := &Generator[T]{
		name: name,
		g: g,
		f: func(g *Generator[T]) T {
			if len(g.values) == 0 {
				return start
			}
			return f(g)
		},
		set: map[string]bool{},
	}
	/*for _, line := range getFromCache(name) {
		if line == "" {
			continue
		}
		t := g.FromString(line)
		gt.values = append(gt.values, t)
		gt.set[g.String(t)] = true
	}*/
	return gt
}

func PrimeFactors(n int, p *Generator[int]) map[int]int {
	if n <= 1 {
		return nil
	}
	r := map[int]int{}
	for i := 0; ; i++ {
		pi := int(p.Nth(i))
		for n%pi == 0 {
			r[pi]++
			n = n / pi
			if n == 1 {
				return r
			}
		}
	}
}

func Primes() *Generator[int] {
	return NewGenerator(primesName, 2, newIntGeneratable(), func(g *Generator[int]) int {
		for i := g.last() + 1; ; i++ {
			newPrime := true
			for _, p := range g.values {
				// Only need to check up to square root of i.
				if p*p > i {
					break
				}
				if rem := i % p; rem == 0 {
					newPrime = false
					break
				}
			}
			if newPrime {
				return i
			}
		}
	})
}

// t_n  = n(2n−1) >= 2 * n * n
func IsHexagonal(tn int) bool {
	if tn < 1 {
		return false
	}
	
	n := maths.Sqrt((tn ) / 2)
	for ; n * (2 * n - 1)  < tn; n ++ {}
	return n * (2 * n - 1) == tn
}


// t_n  = n(3n−1)/2
func IsPentagonal(tn int) bool {
	if tn < 1 {
		return false
	}
	
	n := maths.Sqrt((2 * tn ) / 3)
	for ; n * (3 * n - 1) / 2 < tn; n ++ {}
	return n * (3 * n - 1) / 2 == tn
}

// t_n = n(n+1)/2
func IsTriangular(tn int) bool {
	if tn < 1 {
		return false
	}
	n2 := tn * 2
	n := maths.Sqrt(n2)
	return n * (n + 1) / 2 == tn
}

func IsPrime(n int, p *Generator[int]) bool {
	if n <= 1 {
		return false
	}
	ogIdx := p.idx
	defer func() {p.idx = ogIdx}()
	p.Reset()
	for pn := p.Next(); pn * pn <= n; pn = p.Next() {
		if n % pn == 0 {
			return false
		}
	}
	return true
}

func BigPrimes() *Generator[*maths.Int] {
	return NewGenerator(primesName, maths.NewInt(2), newBigGeneratable(), func(g *Generator[*maths.Int]) *maths.Int {
		for i := g.last().Plus(maths.One()); ; i.PP() {
			newPrime := true
			for _, p := range g.values {
				if p.Times(p).GT(i) {
					break
				}
				if i.Mod(p).IsZero() {
					newPrime = false
					break
				}
			}
			if newPrime {
				return i
			}
		}
	})
}

func Fibonaccis() *Generator[int] {
	a, b := 1, 1
	return NewGenerator(fibName, 1, newIntGeneratable(), func(g *Generator[int]) int {
		r := b
		b = a + b
		a = r
		return int(a)
	})
}

func BigFibonaccis() *Generator[*maths.Int] {
	a, b := maths.One(), maths.One()
	return NewGenerator(fibName, maths.One(), newBigGeneratable(), func(g *Generator[*maths.Int]) *maths.Int {
		r := b
		b = a.Plus(b)
		a = r
		return a
	})
}

func Triangulars() *Generator[int] {
	i := 1
	return NewGenerator(triName, 1, newIntGeneratable(), func(g *Generator[int]) int {
		i++
		return g.last() + i
	})
}

func Pentagonals() *Generator[int] {
	i := 1
	return NewGenerator(triName, 1, newIntGeneratable(), func(g *Generator[int]) int {
		i += 3
		return g.last() + i
	})
}

func Hexagonals() *Generator[int] {
	i := 1
	return NewGenerator(triName, 1, newIntGeneratable(), func(g *Generator[int]) int {
		i += 4
		return g.last() + i
	})
}