package generator

import (
	"strconv"

	"github.com/leep-frog/euler_challenge/maths"
)

var (
	cachedPrimeFactors    = map[int]map[int]int{}
	cachedFactors         = map[int][]int{}
	cachedFactorCounts    = map[int]int{}
	coprimeCache          = map[int]map[int]bool{}
	cachedResilienceCount = map[int]int{}
)

func clearCaches() {
	cachedPrimeFactors = map[int]map[int]int{}
	cachedFactors = map[int][]int{}
	cachedFactorCounts = map[int]int{}
	coprimeCache = map[int]map[int]bool{}
	cachedResilienceCount = map[int]int{}
}

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
	return copy(PrimeFactors(n, p))
}

func FactorCount(n int, p *Generator[int]) int {
	return CompositeCacher(n, p, cachedFactorCounts, func(i int) int {
		if i == 0 {
			return 0
		}
		if i == 1 {
			return 1
		}
		// Factors are 1 and the prime number itself
		return 2
	}, func(primeFactor, otherFactor int) int {
		primeCnt := 1
		rem := otherFactor
		for ; rem%primeFactor == 0; rem, primeCnt = rem/primeFactor, primeCnt+1 {
		}

		// prime^primeCnt * rem = n
		fc := FactorCount(rem, p)

		// Now every factor of rem can be used to create primeCnt new factors.
		// If we have one factor, f, then that factor can create:
		// [f, f*pi, f*pi^2, f*pi^3, ..., f*pi^piCnt] (len = piCnt + 1)
		cachedFactorCounts[n] = fc * (primeCnt + 1)
		return fc * (primeCnt + 1)
	})
}

// AKA Relative Prime Count? (I think)?
func ResilienceCount(n int, p *Generator[int]) int {
	return CompositeCacher(n, p, cachedResilienceCount, func(i int) int {
		if i <= 1 {
			panic("IDK")
		}
		// Factors are 1 and the prime number itself
		return i - 1
	}, func(primeFactor, otherFactor int) int {
		r := ResilienceCount(otherFactor, p)
		// If already has one of the prime, then just multiply
		if otherFactor%primeFactor == 0 {
			return r * primeFactor
		}
		return r * (primeFactor - 1)
	})
}

// CompositeCacher evaluates data for a number, n, by combining info already known
// for two of it's factors (primeFactor being the smallest factor which is inherently prime,
// and otherFactor which is the largest factor of n != n). If n is zero, one, or prime,
// then the value generated is created from the provided forZeroOnePrime function.
func CompositeCacher[T any](n int, p *Generator[int], cache map[int]T, forZeroOnePrime func(int) T, forNonPrime func(primeFactor, otherFactor int) T) T {
	if n < 1 {
		return forZeroOnePrime(0)
	}
	if n == 1 {
		return forZeroOnePrime(1)
	}
	if r, ok := cache[n]; ok {
		return r
	}

	if IsPrime(n, p) {
		r := forZeroOnePrime(n)
		cache[n] = r
		return r
	}

	for i := 0; ; i++ {
		pi := int(p.Nth(i))
		if n%pi != 0 {
			continue
		}

		r := forNonPrime(pi, n/pi)
		cache[n] = r
		return r
	}
}

func Factors(n int, p *Generator[int]) []int {
	return CompositeCacher(n, p, cachedFactors, func(i int) []int {
		if n < 1 {
			return nil
		}
		if n == 1 {
			return []int{1}
		}
		return []int{1, n}
	}, func(primeFactor, otherFactor int) []int {
		// primeFactor is guaranteed to be the smallest factor and (otherFactor = n/primeFactor) the largest.
		additional := Factors(otherFactor, p)
		mAdditional := []int{1}
		for _, a := range additional {
			mAdditional = append(mAdditional, a*primeFactor)
		}
		return maths.MergeSort(additional, mAdditional, true)
	})
}

func PrimeFactors(n int, p *Generator[int]) map[int]int {
	// TODO: update this to use composite cache
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
