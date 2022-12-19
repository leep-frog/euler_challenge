package generator

import (
	"fmt"

	"github.com/leep-frog/euler_challenge/maths"
)

var (
	MaxSieveBatchSize = 100_000_000
	MaxSieveSize      = 180_000_000
	primeSingleton    = BatchedSievedPrimes()
)

// Primes returns the best, generic way to compute primes.
func Primes() *Prime {
	return primeSingleton
}

// BatchedSievedPrimes sieves primes in batches of size k.
func BatchedSievedPrimes() *Prime {
	return BatchedSievedPrimesWithSize(MaxSieveBatchSize)
}

// BatchedSievedPrimesWithSize sieves primes in batches of size k.
func BatchedSievedPrimesWithSize(k int) *Prime {
	k = maths.Min(k, MaxSieveBatchSize)
	return &Prime{newIntGen(&batchedSievedPrimer{
		make([][]*sievePair, k, k),
		nil,
		0,
		3,
	})}
}

type sievePair struct {
	prime int
	val   int
}

func (sp *sievePair) String() string {
	return fmt.Sprintf("(%d, %d)", sp.prime, sp.val)
}

type batchedSievedPrimer struct {
	values    [][]*sievePair
	leftovers []*sievePair
	idx       int
	offset    int
}

func (bsp *batchedSievedPrimer) insert(pair *sievePair) {
	index := (pair.val - bsp.offset) / 2
	if index >= len(bsp.values) {
		bsp.leftovers = append(bsp.leftovers, pair)
	} else {
		bsp.values[index] = append(bsp.values[index], pair)
	}
}

func (bsp *batchedSievedPrimer) update() {
	bsp.offset += 2 * len(bsp.values)
	bsp.idx = 0
	oldLeftovers := bsp.leftovers
	bsp.leftovers = nil
	for _, leftover := range oldLeftovers {
		bsp.insert(leftover)
	}
}

// Returns a number and whether the number is prime
func (bsp *batchedSievedPrimer) check() (int, bool) {
	isPrime := len(bsp.values[bsp.idx]) == 0
	value := bsp.offset + 2*bsp.idx

	for _, pair := range bsp.values[bsp.idx] {
		pair.val += 2 * pair.prime
		bsp.insert(pair)
	}
	bsp.values[bsp.idx] = nil

	bsp.idx++
	if bsp.idx >= len(bsp.values) {
		bsp.update()
	}

	if isPrime {
		bsp.insert(&sievePair{value, 3 * value})
	}

	return value, isPrime
}

func (fp *batchedSievedPrimer) Next(g *Generator[int]) int {
	if len(g.values) == 0 {
		return 2
	}

	for {
		v, ok := fp.check()
		for ; !ok; v, ok = fp.check() {
		}

		return v
	}
}

// SievedPrimesUpTo returns a generator that sieves primes up to k
func SievedPrimesUpTo(k int) *Prime {
	return &Prime{newIntGen(&fixedValuePrimer{nil, -1, maths.Min(k, MaxSieveSize)})}
}

// SievedPrimes returns a generator that sieves primes up to a fixed maximum
func SievedPrimes() *Prime {
	return SievedPrimesUpTo(MaxSieveSize)
}

type fixedValuePrimer struct {
	values []bool
	idx    int
	size   int
}

func (bp *fixedValuePrimer) Next(g *Generator[int]) int {
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

// BasicPrimes returns the simplest method to generate primes
func BasicPrimes() *Prime {
	return &Prime{newIntGen(&primer{})}
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
