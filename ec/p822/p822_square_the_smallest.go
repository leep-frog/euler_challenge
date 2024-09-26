package p822

import (
	"fmt"
	"math"
	"math/big"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

const (
	mod = 1234567891
)

var (
	log2 = math.Log(2)
)

type squareRepr struct {
	k      int
	twoPow int

	// The number will be k^(2^twoPow), we take the double log of that:
	// log(k^2^twoPow) = 2 ^ twoPow * log(k)
	// log(log(k^2^twoPow)) = log(2 ^ twoPow * log(k))
	//                      = log(2^twoPow) + log(log(k))
	//                      = twoPow * log(2) + log(log(k))
	loglogk float64
}

func (sr *squareRepr) increment() {
	sr.incrementTimes(1)
}

func (sr *squareRepr) incrementTimes(n int) {
	sr.twoPow += n
}

func (sr *squareRepr) incrementUpTo(rootMax *squareRepr) int {
	var increments int
	for sr.lt(rootMax) {
		sr.increment()
		increments++
	}
	return increments
}

const PRECISION = 100

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
	ltp = ltp.Mul(ltp, newFloat(log2))
	ltp = ltp.Add(ltp, newFloat(sr.loglogk))

	rtp := newFlint(that.twoPow)
	rtp = rtp.Mul(rtp, newFloat(log2))
	rtp = rtp.Add(rtp, newFloat(that.loglogk))

	return ltp.Cmp(rtp) <= 0

	lf := float64(sr.twoPow)*log2 + sr.loglogk
	left := big.NewFloat(lf)
	rf := float64(that.twoPow)*log2 + that.loglogk
	right := big.NewFloat(rf)

	fmt.Println("LEFT", ltp, left, "RIGHT", rtp, right)

	return left.Cmp(right) < 0
}

func P822() *ecmodels.Problem {
	return ecmodels.IntsInputNode(822, 2, 0, func(o command.Output, ns []int) {

		n, m := ns[0], ns[1]

		var srs []*squareRepr
		for i := 2; i <= n; i++ {
			fmt.Println("adding", i)
			srs = append(srs, &squareRepr{i, 0, math.Log(math.Log(float64(i)))})
		}

		sqrtIdx := (maths.Sqrt(n)) - 1
		sqrtSq := srs[sqrtIdx]
		maxRoot := &squareRepr{sqrtSq.k, 0, sqrtSq.loglogk}
		_ = maxRoot

		// fmt.Println(brute(srs, m))
		fmt.Println(elegant(srs, m, maxRoot))

		// First, get all values to be greater than the
		// o.Stdoutln(n)
	}, []*ecmodels.Execution{
		{
			Args: []string{"1"},
			Want: "",
		},
		{
			Args: []string{"2"},
			Want: "",
		},
	})
}

func elegant(nums []*squareRepr, times int, maxRoot *squareRepr) int {
	var increments int
	for _, n := range nums {
		increments += n.incrementUpTo(maxRoot)
	}

	incrAll := ((times - increments) / len(nums)) - 1
	for _, n := range nums {
		fmt.Println("DOING", n, incrAll)
		n.incrementTimes(incrAll)
		// for i := 0; i < incrAll; i++ {
		// n.increment()
		// }
	}

	return brute(nums, ((times-increments)%len(nums))+len(nums))
}

func brute(nums []*squareRepr, times int) int {
	heap := maths.NewHeap(func(sr1, sr2 *squareRepr) bool {
		return sr1.lt(sr2)
	})

	for _, n := range nums {
		heap.Push(n)
	}

	for i := 0; i < times; i++ {
		// fmt.Println("start incrementing")
		sr := heap.Pop()
		sr.increment()
		// fmt.Println("incrementing", sr.k)
		heap.Push(sr)
		// fmt.Println("done incrementing")
	}

	var sum, cnt int
	heap.Iter(func(sr *squareRepr) bool {
		fmt.Println("summing", sr.k, sr.twoPow)
		// sum = (sum + maths.PowMod(sr.k, maths.PowMod(2, sr.twoPow, mod), mod)) % mod

		// Num twos is sr.twoPos
		// numTwos := sr.twoPow % (mod - 1)

		sum = (sum + maths.PowMod(sr.k, maths.PowMod(2, sr.twoPow, mod-1), mod)) % mod
		cnt += sr.twoPow
		return true
	})
	fmt.Println("COUNT", cnt)
	return sum
}
