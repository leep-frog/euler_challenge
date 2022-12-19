package generator

import "fmt"

// TODO: this was failing for larger values, need to
// more thoroguhly test implementation.
func FinalPrimes(k int) *Generator[int] {
	return newIntGen(&finalPrimer{
		make([][]*finalPair, k, k),
		nil,
		0,
		3,
	})
}

type finalPair struct {
	prime int
	val   int
}

func (fp *finalPair) String() string {
	return fmt.Sprintf("(%d, %d)", fp.prime, fp.val)
}

type finalPrimer struct {
	values    [][]*finalPair
	leftovers []*finalPair
	idx       int
	offset    int
}

func (fp *finalPrimer) insert(pair *finalPair) {
	index := (pair.val - fp.offset) / 2
	if index >= len(fp.values) {
		fp.leftovers = append(fp.leftovers, pair)
	} else {
		fp.values[index] = append(fp.values[index], pair)
	}
}

func (fp *finalPrimer) update() {
	fp.offset += 2 * len(fp.values)
	fp.idx = 0
	oldLeftovers := fp.leftovers
	fp.leftovers = nil
	for _, leftover := range oldLeftovers {
		fp.insert(leftover)
	}
}

// Returns a number and whether the number is prime
func (fp *finalPrimer) check() (int, bool) {
	isPrime := len(fp.values[fp.idx]) == 0
	value := fp.offset + 2*fp.idx

	for _, pair := range fp.values[fp.idx] {
		pair.val += 2 * pair.prime
		fp.insert(pair)
	}
	fp.values[fp.idx] = nil

	fp.idx++
	if fp.idx >= len(fp.values) {
		fp.update()
	}

	if isPrime {
		fp.insert(&finalPair{value, 3 * value})
	}

	return value, isPrime
}

func (fp *finalPrimer) Next(g *Generator[int]) int {
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
