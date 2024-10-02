package generator

import (
	"math/rand"
	"strconv"

	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/maths"
)

var (
	cachedPrimeFactors           = map[int]map[int]int{}
	cachedPrimeFactorsFast       = [][][]int{}
	cachedPrimeFactorIndicesFast = []PrimeFactoredNumber{}
	cachedFactors                = map[int][]int{}
	cachedFactorCounts           = map[int]int{}
	coprimeCache                 = map[int]map[int]bool{}
	cachedResilienceCount        = map[int]int{}
)

func ClearCaches() {
	cachedPrimeFactors = map[int]map[int]int{}
	cachedPrimeFactorsFast = [][][]int{}
	cachedPrimeFactorIndicesFast = []PrimeFactoredNumber{}
	cachedFactors = map[int][]int{}
	cachedFactorCounts = map[int]int{}
	coprimeCache = map[int]map[int]bool{}
	cachedResilienceCount = map[int]int{}
}

// TODO: Just make a prime generator that loads the primes we've already evaluated

type Prime struct {
	*Generator[int]
}

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
			maths.Insert(coprimeCache, a, b, false)
			return false
		}
	}
	maths.Insert(coprimeCache, a, b, true)
	return true
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

// CompositeCacherFast is like CompositeCacher, but uses a slice to cache known values rather than a map.
// This works very well for use cases that requires iterative values (all/most values up to n), but less well for sparse
// sets of numbers.
//
// While a slice is faster, this requires memory of size O(n) (since the slice will be of length n)
// hence why this won't work well for caching sparse sets of numbers.
func CompositeCacherFast[T any](p *Prime, n int, pCache *[]T, forZeroOnePrime func(int) T, forNonPrime func(primeFactor, otherFactor int) T) T {

	cache := *pCache
	for len(cache) <= n {
		k := len(cache)
		if k <= 1 {
			cache = append(cache, forZeroOnePrime(k))
			continue
		}

		if p.Contains(k) {
			cache = append(cache, forZeroOnePrime(k))
			continue
		}

		for i := 0; ; i++ {
			pi := int(p.Nth(i))
			if k%pi != 0 {
				continue
			}

			cache = append(cache, forNonPrime(pi, k/pi))
			break
		}
	}
	*pCache = cache
	return cache[n]
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

// TODO: Returning map[int]int might actually be faster.  Try it out
func (p *Prime) PrimeFactorsFast(n int) [][]int {
	return CompositeCacherFast(p, n, &cachedPrimeFactorsFast,
		func(i int) [][]int {
			if i <= 1 {
				return nil
			}
			return [][]int{{i, 1}}
		},
		func(primeFactor, otherFactor int) [][]int {
			pff := p.PrimeFactorsFast(otherFactor)

			var m [][]int
			var added bool
			for _, pf := range pff {
				if pf[0] == primeFactor {
					added = true
					m = append(m, []int{pf[0], pf[1] + 1})
				} else {
					m = append(m, []int{pf[0], pf[1]})
				}
			}

			if added {
				return m
			}
			return append(m, []int{primeFactor, 1})
			// for _, p := range m {
			// 	if p[0] == primeFactor {
			// 		p[1]++
			// 		return m
			// 	}
			// }

			// // m := copy(
			// // m[primeFactor]++
			// return append(m, []int{primeFactor, 1})
		},
	)
}

// Stores primes in descending order
type PrimeIndexCount []int

func (pic PrimeIndexCount) Prime(p *Prime) int {
	return p.Nth(pic[0])
}

func (pic PrimeIndexCount) Count() int {
	return pic[1]
}

// type PrimeIndexNumber []PrimeIndexCount

// func (pin PrimeIndexNumber) Times() int {
// 	bread.MergeSort()
// 	return pic[1]
// }

// Empty array means 1
type PrimeFactoredNumber []PrimeIndexCount

func (pfn PrimeFactoredNumber) ToInt(p *Prime) int {
	prod := 1
	for _, pf := range pfn {
		prod *= maths.Pow(p.Nth(pf[0]), pf[1])
	}
	return prod
}

func (pfn PrimeFactoredNumber) ToBigInt(p *Prime) *maths.Int {
	prod := maths.One()
	for _, pf := range pfn {
		prod = prod.Times(maths.BigPow(p.Nth(pf[0]), pf[1]))
	}
	return prod
}

func (pfn PrimeFactoredNumber) TimesInt(d int) PrimeFactoredNumber {
	return pfn.Times(primeSingleton.PrimeFactoredNumberFast(d))
}

func (pfn PrimeFactoredNumber) Times(that PrimeFactoredNumber) PrimeFactoredNumber {
	var res PrimeFactoredNumber

	for ai, bi := 0, 0; ai < len(pfn) || bi < len(that); {
		if ai == len(pfn) {
			res = append(res, []int{that[bi][0], that[bi][1]})
			bi++
		} else if bi == len(that) || pfn[ai][0] > that[bi][0] {
			res = append(res, []int{pfn[ai][0], pfn[ai][1]})
			ai++
		} else if pfn[ai][0] == that[bi][0] {
			cnt := pfn[ai][1] + that[bi][1]
			if cnt != 0 {
				res = append(res, []int{pfn[ai][0], cnt})
			}
			ai++
			bi++
		} else {
			res = append(res, []int{that[bi][0], that[bi][1]})
			bi++
		}
	}
	return res
}

func (pfn PrimeFactoredNumber) DivInt(d int) PrimeFactoredNumber {
	return pfn.Div(primeSingleton.PrimeFactoredNumberFast(d))
}

func (pfn PrimeFactoredNumber) Div(that PrimeFactoredNumber) PrimeFactoredNumber {
	var res PrimeFactoredNumber

	for ai, bi := 0, 0; ai < len(pfn) || bi < len(that); {
		if ai == len(pfn) {
			res = append(res, []int{that[bi][0], -that[bi][1]})
			bi++
		} else if bi == len(that) || pfn[ai][0] > that[bi][0] {
			res = append(res, []int{pfn[ai][0], pfn[ai][1]})
			ai++
		} else if pfn[ai][0] == that[bi][0] {
			cnt := pfn[ai][1] - that[bi][1]
			if cnt != 0 {
				res = append(res, []int{pfn[ai][0], cnt})
			}
			ai++
			bi++
		} else {
			res = append(res, []int{that[bi][0], -that[bi][1]})
			bi++
		}
	}
	return res
}

func (pfn PrimeFactoredNumber) Pow(k int) PrimeFactoredNumber {
	var res PrimeFactoredNumber
	for _, pic := range pfn {
		res = append(res, []int{pic[0], pic[1] * k})
	}
	return res
}

func (pfn PrimeFactoredNumber) NumFactors(p *Prime, mod int) int {
	res := 1
	for _, pic := range pfn {
		harmonic := maths.PowMod(p.Nth(pic[0]), pic[1]+1, mod)
		harmonic = (harmonic + mod - 1) % mod
		harmonic = (harmonic * maths.PowMod(p.Nth(pic[0])-1, -1, mod)) % mod
		res = (res * harmonic) % mod
	}
	return res
}

// Iterate iterates over the factor values in pfn and that and runs f
// on all values
// TODO: Test this
func (pfn PrimeFactoredNumber) Iterate(that PrimeFactoredNumber, f func(primeFactorIndex, thisCnt, thatCnt int)) {
	for ai, bi := 0, 0; ai < len(pfn) || bi < len(that); {
		if ai == len(pfn) {
			f(that[bi][0], 0, that[bi][1])
			bi++
		} else if bi == len(that) || pfn[ai][0] > that[bi][0] {
			f(pfn[ai][0], pfn[ai][1], 0)
			ai++
		} else if pfn[ai][0] == that[bi][0] {
			f(pfn[ai][0], pfn[ai][1], that[bi][1])
			ai++
			bi++
		} else {
			f(that[bi][0], 0, that[bi][1])
			bi++
		}
	}
}

func (pfn PrimeFactoredNumber) Eq(that PrimeFactoredNumber) bool {
	return pfn.Cmp(that) == 0
}

func (pfn PrimeFactoredNumber) Cmp(that PrimeFactoredNumber) int {
	for ai := 0; ai < len(pfn) || ai < len(that); ai++ {
		if ai == len(pfn) {
			return -1
		} else if ai == len(that) {
			return 1
		} else if pfn[ai][0] == that[ai][0] {
			if pfn[ai][1] < that[ai][1] {
				return -1
			} else if pfn[ai][1] > that[ai][1] {
				return 1
			}
		} else if pfn[ai][0] > that[ai][0] {
			return 1
		} else {
			return -1
		}
	}

	return 0
}

func (p *Prime) PrimeFactoredNumberFast(n int) PrimeFactoredNumber {
	return CompositeCacherFast(p, n, &cachedPrimeFactorIndicesFast,
		func(i int) PrimeFactoredNumber {
			if i <= 1 {
				return nil
			}
			return []PrimeIndexCount{{p.index(i), 1}}
		},
		func(primeFactor, otherFactor int) PrimeFactoredNumber {
			pff := p.PrimeFactoredNumberFast(otherFactor)

			pfi := p.index(primeFactor)

			var m []PrimeIndexCount
			var added bool
			for _, pf := range pff {
				if pf[0] == pfi {
					added = true
					m = append(m, []int{pf[0], pf[1] + 1})
				} else {
					m = append(m, []int{pf[0], pf[1]})
				}
			}

			if added {
				return m
			}
			return append(m, []int{pfi, 1})
		},
	)
}

var (
	factoredFactorialCache = []PrimeFactoredNumber{
		primeSingleton.PrimeFactoredNumberFast(1),
		primeSingleton.PrimeFactoredNumberFast(1),
	}
)

func (p *Prime) PrimeFactoredNumberFactorial(n int) PrimeFactoredNumber {
	for len(factoredFactorialCache) <= n {
		lastIdx := len(factoredFactorialCache) - 1
		factoredFactorialCache = append(factoredFactorialCache, factoredFactorialCache[lastIdx].Times(p.PrimeFactoredNumberFast(lastIdx+1)))
	}

	return factoredFactorialCache[n]
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

func (p *Prime) index(pi int) int {
	if !p.Contains(pi) {
		return -1
	}

	lastIdx := len(p.values) - 1
	for pa := p.values[lastIdx]; pa < pi; pa, lastIdx = p.Nth(lastIdx+1), lastIdx+1 {
	}
	return p.set[strconv.Itoa(pi)]
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
