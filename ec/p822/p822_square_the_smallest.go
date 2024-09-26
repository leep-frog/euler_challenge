package p822

import (
	"math"
	"math/big"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

var (
	mod            = 1234567891
	PRECISION uint = 100
	log2           = newFloat(math.Log(2))
)

// squareRepr is a specific representation of a number k^(2^twoPow).
// Every time the number is squared, we simply need to increment twoPow
type squareRepr struct {
	k      int
	twoPow int

	// The number will be k^(2^twoPow), we take the double log of that:
	// log(k^2^twoPow) = 2 ^ twoPow * log(k)
	// log(log(k^2^twoPow)) = log(2 ^ twoPow * log(k))
	//                      = log(2^twoPow) + log(log(k))
	//                      = twoPow * log(2) + log(log(k))
	loglogk *big.Float
}

func (sr *squareRepr) incrementTimes(n int) {
	sr.twoPow += n
}

func (sr *squareRepr) incrementUpTo(rootMax *squareRepr) int {
	var increments int
	for sr.lt(rootMax) {
		sr.incrementTimes(1)
		increments++
	}
	return increments
}

func newFloat(f float64) *big.Float {
	r := new(big.Float)
	r.SetPrec(PRECISION)
	return r.SetFloat64(f)
}

func newFlint(i int) *big.Float {
	r := new(big.Float)
	r.SetPrec(PRECISION)
	return r.SetInt64(int64(i))
}

func (sr *squareRepr) lt(that *squareRepr) bool {
	ltp := newFlint(sr.twoPow)
	ltp = ltp.Mul(ltp, log2)
	ltp = ltp.Add(ltp, sr.loglogk)

	rtp := newFlint(that.twoPow)
	rtp = rtp.Mul(rtp, log2)
	rtp = rtp.Add(rtp, that.loglogk)

	return ltp.Cmp(rtp) <= 0
}

func P822() *ecmodels.Problem {
	return ecmodels.IntsInputNode(822, 2, 0, func(o command.Output, ns []int) {

		n, m := ns[0], ns[1]

		var srs []*squareRepr
		for i := 2; i <= n; i++ {
			srs = append(srs, &squareRepr{i, 0, newFloat(math.Log(math.Log(float64(i))))})
		}

		sqrtIdx := (maths.Sqrt(n)) - 1
		sqrtSq := srs[sqrtIdx]
		maxRoot := &squareRepr{sqrtSq.k, 0, sqrtSq.loglogk}

		o.Stdoutln(elegant(srs, m, maxRoot))
	}, []*ecmodels.Execution{
		{
			Args: []string{"5", "3"},
			Want: "34",
		},
		{
			Args: []string{"10", "100"},
			Want: "845339386",
		},
		{
			Args: []string{"10000", "10000000000000000"},
			Want: "950591530",
		},
	})
}

// elegant uses the fact that all numbers will be squared a similar number of
// times to reduce the number of operations required.
func elegant(nums []*squareRepr, times int, maxRoot *squareRepr) int {

	// Increment each number to be bigger than the squareRoot of the largest number
	var increments int
	for _, n := range nums {
		increments += n.incrementUpTo(maxRoot)
	}

	// Now, each number will be squared at least that many times
	incrAll := ((times - increments) / len(nums))
	for _, n := range nums {
		n.incrementTimes(incrAll)
	}

	// Finally, simply run the brute force algorithm for the remaining times
	return brute(nums, (times-increments)%len(nums))
}

// brute simply adds all number representations to a heap and iterates over the
// smallest element n times, squaring it.
func brute(nums []*squareRepr, n int) int {

	// Create the heap
	heap := maths.NewHeap(func(sr1, sr2 *squareRepr) bool {
		return sr1.lt(sr2)
	})
	for _, num := range nums {
		heap.Push(num)
	}

	// Square the smallest number in the heap, n times
	for i := 0; i < n; i++ {
		sr := heap.Pop()
		sr.incrementTimes(1)
		heap.Push(sr)
	}

	// Calculate the mod sum
	var sum int
	heap.Iter(func(sr *squareRepr) bool {

		// Consider 2^k % 3:
		// 2^1 % 3 = 2 % 3 = 2
		// 2^2 % 3 = 4 % 3 = 1
		// 2^3 % 3 = 8 % 3 = 2
		// 2^4 % 3 = 16 % 3 = 1
		// So while the modulo is 3, the possibilities are (mod-1), hence why we
		// take the PowMod with (mod-1)
		moddedTwoPow := maths.PowMod(2, sr.twoPow, mod-1)

		// Increment the sum by k^moddedTwoPos
		sum = (sum + maths.PowMod(sr.k, moddedTwoPow, mod)) % mod
		return true
	})
	return sum
}
