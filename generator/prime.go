package generator

import (
	"math/rand"
	"strconv"

	"github.com/leep-frog/euler_challenge/bread"
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

// TODO: Just make a prime generator that loads the primes we've already evaluated

type Prime struct {
	*Generator[int]
}

var (
	rppCache = map[string]int{}
)

func (p *Prime) recPrimePi(rem, minIdx, maxV, sign int, start int) int {
	if rem <= 0 {
		return 0
	}
	// code := fmt.Sprintf("%d %d %d", rem, minIdx, maxV)
	// if v, ok := rppCache[code]; ok {
	// 	return v * sign
	// }
	sum := rem * sign

	/*for iter, prime := g.Start(minIdx); prime <= rem; prime = iter.Next() {
		sum += recPrimePi(rem/prime, iter.Idx+1, -sign, g)
	}*/
	for i := minIdx; ; i++ {
		prime := p.Nth(i)
		if prime > rem || prime > maxV {
			break
		}
		sum += p.recPrimePi(rem/prime, i+1, maxV, -sign, start+1)
	}

	// rppCache[code] = sum * sign
	return sum
}

func brutePrimePi(x int) int {
	iter, prime := Primes().Start(0)
	for prime <= x {
		prime = iter.Next()
	}
	return iter.Idx - 1
}

var (
	primePiCache = map[int]int{}
)

func (p *Prime) PrimePi(x int) int {
	if v, ok := primePiCache[x]; ok {
		return v
	}
	if x <= 1 {
		return brutePrimePi(x)
	}
	summation := p.recPrimePi(x, 0, maths.Sqrt(x), 1, 0)
	r := summation + p.PrimePi(maths.Sqrt(x)) - 1
	primePiCache[x] = r
	return r
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

func (p *Prime) Coprimes(a, b int) bool {
	if b < a {
		a, b = b, a
	}
	if m, ok := coprimeCache[a]; ok {
		if v, ok2 := m[b]; ok2 {
			return v
		}
	}
	bFactors := p.PrimeFactors(b)
	for k := range p.PrimeFactors(a) {
		if _, ok := bFactors[k]; ok {
			maths.Insert(coprimeCache, a, b, true)
			return true
		}
	}
	maths.Insert(coprimeCache, a, b, false)
	return false
}

func (p *Prime) MutablePrimeFactors(n int) map[int]int {
	return copy(p.PrimeFactors(n))
}

func (p *Prime) FactorCount(n int) int {
	return CompositeCacher(p, n, cachedFactorCounts, func(i int) int {
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
		fc := p.FactorCount(rem)

		// Now every factor of rem can be used to create primeCnt new factors.
		// If we have one factor, f, then that factor can create:
		// [f, f*pi, f*pi^2, f*pi^3, ..., f*pi^piCnt] (len = piCnt + 1)
		cachedFactorCounts[n] = fc * (primeCnt + 1)
		return fc * (primeCnt + 1)
	})
}

// AKA Relative Prime Count? (I think)?
func (p *Prime) ResilienceCount(n int) int {
	return CompositeCacher(p, n, cachedResilienceCount, func(i int) int {
		if i <= 1 {
			panic("IDK")
		}
		// Factors are 1 and the prime number itself
		return i - 1
	}, func(primeFactor, otherFactor int) int {
		r := p.ResilienceCount(otherFactor)
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
func CompositeCacher[T any](p *Prime, n int, cache map[int]T, forZeroOnePrime func(int) T, forNonPrime func(primeFactor, otherFactor int) T) T {
	if n < 1 {
		return forZeroOnePrime(0)
	}
	if n == 1 {
		return forZeroOnePrime(1)
	}
	if r, ok := cache[n]; ok {
		return r
	}

	if p.Contains(n) {
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

func (p *Prime) Factors(n int) []int {
	return CompositeCacher(p, n, cachedFactors, func(i int) []int {
		if n < 1 {
			return nil
		}
		if n == 1 {
			return []int{1}
		}
		return []int{1, n}
	}, func(primeFactor, otherFactor int) []int {
		// primeFactor is guaranteed to be the smallest factor and (otherFactor = n/primeFactor) the largest.
		additional := p.Factors(otherFactor)
		mAdditional := []int{1}
		for _, a := range additional {
			// TODO: Check this logic if multiple factors (2*2*2 * ...) and primeFactor is 2
			mAdditional = append(mAdditional, a*primeFactor)
		}
		return bread.MergeSort(additional, mAdditional, true)
	})
}

func (p *Prime) PrimeFactors(n int) map[int]int {
	return CompositeCacher(p, n, cachedPrimeFactors,
		func(i int) map[int]int {
			if i <= 1 {
				return nil
			}
			return map[int]int{i: 1}
		},
		func(primeFactor, otherFactor int) map[int]int {
			m := copy(p.PrimeFactors(otherFactor))
			m[primeFactor]++
			return m
		},
	)
}

// func (p *Prime) PrimeFactors(n int) map[int]int {
// 	// TODO: update this to use composite cache
// 	if n <= 1 {
// 		return nil
// 	}
// 	if r, ok := cachedPrimeFactors[n]; ok {
// 		return r
// 	}
// 	if p.Contains(n) {
// 		r := map[int]int{n: 1}
// 		cachedPrimeFactors[n] = r
// 		return r
// 	}
// 	ogN := n
// 	for i := 0; ; i++ {
// 		pi := int(p.Nth(i))
// 		if n%pi == 0 {
// 			fs := copy(p.PrimeFactors(n / pi))
// 			fs[pi]++
// 			cachedPrimeFactors[ogN] = fs
// 			return fs
// 		}
// 	}
// 	panic("Should not reach here")
// }

// Overrides Generator.Contains
func (p *Prime) Contains(n int) bool {
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

func (p *Prime) FermatContains(n, checks int) bool {

	if n < 100 {
		return p.Contains(n)
	}

	for i := 0; i <= checks; i++ {
		a := (rand.Int() % (n - 4)) + 2
		if maths.PowMod(a, n-1, n) != 1 {
			return false
		}
	}
	return true
}
