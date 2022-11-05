package generator

import (
	"strconv"

	"github.com/leep-frog/euler_challenge/maths"
)

var (
	cachedPrimeFactors = map[int]map[int]int{}
	cachedFactors      = map[int][]int{}
	coprimeCache       = map[int]map[int]bool{}
)

func Primes() *Generator[int] {
	return newIntGen(&primer{})
}

type primer struct{}

func (p *primer) Next(g *Generator[int]) int {
	if len(g.values) == 0 {
		return 2
	}
	for i := g.Last() + 1; ; i++ {
		newPrime := true
		for _, p := range g.Values() {
			if p*p > i {
				break
			}
			if i%p == 0 {
				newPrime = false
				break
			}
		}
		if newPrime {
			return i
		}
	}
}

func BigPrimes() *Generator[*maths.Int] {
	return newBigGen(&bigPrimer{})
}

type bigPrimer struct{}

func (bp *bigPrimer) Next(g *Generator[*maths.Int]) *maths.Int {
	if len(g.values) == 0 {
		return maths.NewInt(2)
	}
	for i := g.Last().Plus(maths.One()); ; i.PP() {
		newPrime := true
		for _, p := range g.Values() {
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
}

func Coprimes(a, b int, p *Generator[int]) bool {
	if b < a {
		a, b = b, a
	}
	if m, ok := coprimeCache[a]; ok {
		if v, ok2 := m[b]; ok2 {
			return v
		}
	}
	bFactors := PrimeFactors(b, p)
	for k := range PrimeFactors(a, p) {
		if _, ok := bFactors[k]; ok {
			maths.Insert(coprimeCache, a, b, true)
			return true
		}
	}
	maths.Insert(coprimeCache, a, b, false)
	return false
}

func MutablePrimeFactors(n int, p *Generator[int]) map[int]int {
	return copy(primeFactors(n, p))
}

func PrimeFactors(n int, p *Generator[int]) map[int]int {
	return primeFactors(n, p)
}

func Factors(n int, p *Generator[int]) []int {
	if n < 1 {
		return nil
	}
	if n == 1 {
		return []int{1}
	}
	if r, ok := cachedFactors[n]; ok {
		return r
	}

	if IsPrime(n, p) {
		r := []int{1, n}
		cachedFactors[n] = r
		return r
	}

	for i := 0; ; i++ {
		pi := int(p.Nth(i))
		if n%pi != 0 {
			continue
		}
		// pi is guaranteed to be the smallest factor and n/pi the largest
		additional := Factors(n/pi, p)
		var mAdditional []int
		for _, a := range additional {
			mAdditional = append(mAdditional, a*pi)
		}
		// merge sort the two
		merged := []int{1}
		for ai, mi := 0, 0; ai < len(additional) || mi < len(mAdditional); {
			var contender int
			if ai == len(additional) {
				contender = mAdditional[mi]
				mi++
			} else if mi == len(mAdditional) {
				contender = additional[ai]
				ai++
			} else if additional[ai] <= mAdditional[mi] {
				contender = additional[ai]
				ai++
			} else {
				contender = mAdditional[mi]
				mi++
			}
			if contender != merged[len(merged)-1] {
				merged = append(merged, contender)
			}
		}
		cachedFactors[n] = merged
		return merged
	}
}

func primeFactors(n int, p *Generator[int]) map[int]int {
	if n <= 1 {
		return nil
	}
	if r, ok := cachedPrimeFactors[n]; ok {
		return r
	}
	if IsPrime(n, p) {
		r := map[int]int{n: 1}
		cachedPrimeFactors[n] = r
		return r
	}
	ogN := n
	r := map[int]int{}
	for i := 0; ; i++ {
		pi := int(p.Nth(i))
		for n%pi == 0 {
			r[pi]++
			n = n / pi
			if n == 1 {
				cachedPrimeFactors[ogN] = r
				return r
			}
			if extra, ok := cachedPrimeFactors[n]; ok {
				for k, v := range extra {
					r[k] += v
				}
				cachedPrimeFactors[ogN] = r
				return r
			}
		}
	}
}

func IsPrime(n int, p *Generator[int]) bool {
	if n <= 1 {
		return false
	}
	if len(p.values) > 0 && p.Last() >= n {
		if _, has := p.set[strconv.Itoa(n)]; has {
			return true
		}
	}
	for i, pn := 0, p.Nth(0); pn*pn <= n; i, pn = i+1, p.Nth(i+1) {
		if n%pn == 0 {
			return false
		}
	}
	return true
}
