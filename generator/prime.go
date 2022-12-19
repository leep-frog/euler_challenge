package generator

import (
	"fmt"
	"strconv"

	"github.com/leep-frog/euler_challenge/maths"
)

var (
	cachedPrimeFactors    = map[int]map[int]int{}
	cachedFactors         = map[int][]int{}
	cachedFactorCounts    = map[int]int{}
	coprimeCache          = map[int]map[int]bool{}
	cachedResilienceCount = map[int]int{}
	//primeGenerator        = &Prime{newIntGen(&primer{})}
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

func Primes() *Prime {
	return &Prime{newIntGen(&primer{})}
	//return &Prime{newIntGen(&primer{})}
	// Prettry sure we can use the same instance like below. Just verify
	// test speed doesn't change (or if it does, that it gets better).
	//return primeGenerator
}

func BetterPrimes() *Generator[int] {
	return newIntGen(&betterPrimer{})
}

func PrimesUpTo(k int) *Generator[int] {
	return newIntGen(&bestPrimer{nil, -1, k})
}

// Pair of prime and current multiple
type primePair struct {
	prime int
	val   int
}

func (pp *primePair) String() string {
	return fmt.Sprintf("(%d, %d)", pp.prime, pp.val)
}

type betterPrimer struct {
	heap *maths.Heap[*primePair]
	// slice containing
	rem []int
	idx int
}

func (bp *betterPrimer) Next(g *Generator[int]) int {
	if len(g.values) == 0 {
		return 2
	}
	if len(g.values) == 1 {
		bp.heap = maths.NewHeap(func(pp1, pp2 *primePair) bool {
			return pp1.val < pp2.val
		})
		bp.heap.Push(&primePair{3, 9})
		bp.idx = 5
		return 3
	}

	for ; ; bp.idx += 2 {

		valid := true
		pp := bp.heap.Pop()
		for ; pp.val <= bp.idx; pp = bp.heap.Pop() {
			valid = valid && (pp.val != bp.idx)
			pp.val += pp.prime
			bp.heap.Push(pp)
		}
		// The last one just needs to be pushed unchanged
		bp.heap.Push(pp)

		if valid {
			i := bp.idx
			bp.idx += 2
			bp.heap.Push(&primePair{i, 2 * i})
			return i
		}
	}
	/*for ; ; bp.idx += 2 {
		pp := bp.heap.Peek()

		// bp.idx is prime, in which case add the other thing back on the stack
		if pp.val != bp.idx {
			i := bp.idx
			bp.idx += 2
			return i
		}

		for ; pp.val
		// bp.idx is not prime, so increment its value
		pp.val += pp.prime
		bp.heap.Push(pp)
	}*/
}

type bestPrimer struct {
	values []bool
	idx    int
	size   int
}

func (bp *bestPrimer) Next(g *Generator[int]) int {
	if len(g.values) == 0 {
		return 2
	}
	if len(g.values) == 1 {
		// TODO: change size to half
		bp.values = make([]bool, bp.size, bp.size)
		for i := 3; i < len(bp.values); i += 3 {
			bp.values[i] = true
		}
		bp.values = bp.values[5:]
		bp.idx = 5
		return 3
	}

	offset := 0
	for ; offset < len(bp.values) && bp.values[offset]; offset += 2 {
	}
	if offset+2 > len(bp.values) {
		panic(fmt.Sprintf("above maximum for prime generator; nth=%d, prime=%d", len(g.values), g.Last()))
	}

	prime := offset + bp.idx
	for i := offset; i < len(bp.values); i += prime {
		bp.values[i] = true
	}
	bp.idx += offset + 2
	bp.values = bp.values[offset+2:]
	return prime
}

var (
	rppCache = map[string]int{}
)

func recPrimePi(rem, minIdx, maxV, sign int, g *Generator[int], start int) int {
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
		prime := g.Nth(i)
		if prime > rem || prime > maxV {
			break
		}
		sum += recPrimePi(rem/prime, i+1, maxV, -sign, g, start+1)
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

func PrimePi(x int, primes *Generator[int]) int {
	if v, ok := primePiCache[x]; ok {
		return v
	}
	if x <= 1 {
		return brutePrimePi(x)
	}
	summation := recPrimePi(x, 0, maths.Sqrt(x), 1, primes, 0)
	r := summation + PrimePi(maths.Sqrt(x), primes) - 1
	primePiCache[x] = r
	return r
}

// TODO: Cached primer

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
			mAdditional = append(mAdditional, a*primeFactor)
		}
		return maths.MergeSort(additional, mAdditional, true)
	})
}

func (p *Prime) PrimeFactors(n int) map[int]int {
	// TODO: update this to use composite cache
	if n <= 1 {
		return nil
	}
	if r, ok := cachedPrimeFactors[n]; ok {
		return r
	}
	if p.Contains(n) {
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
