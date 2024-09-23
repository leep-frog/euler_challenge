package p822

import (
	"fmt"
	"math"

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

func (sr *squareRepr) lt(that *squareRepr) bool {
	left := float64(sr.twoPow)*log2 + sr.loglogk
	right := float64(that.twoPow)*log2 + that.loglogk
	return left < right
}

func P822() *ecmodels.Problem {
	return ecmodels.IntsInputNode(822, 2, 0, func(o command.Output, ns []int) {

		n, m := ns[0], ns[1]

		var srs []*squareRepr
		for i := 2; i <= n; i++ {
			fmt.Println("adding", i)
			srs = append(srs, &squareRepr{i, 0, math.Log(float64(i))})
		}

		last := srs[len(srs)-1]
		sqrt := maths.Sqrt(last.k) - 1
		maxRoot := &squareRepr{sqrt, 0, math.Log(float64(sqrt))}
		_ = maxRoot

		fmt.Println(brute(srs, m))
		// fmt.Println(elegant(srs, m, maxRoot))

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

	incrAll := (times - increments) / len(nums)
	for _, n := range nums {
		fmt.Println("DOING", n, incrAll)
		n.incrementTimes(incrAll)
		// for i := 0; i < incrAll; i++ {
		// n.increment()
		// }
	}

	return brute(nums, (times-increments)%len(nums))
}

func brute(nums []*squareRepr, times int) int {
	heap := maths.NewHeap(func(sr1, sr2 *squareRepr) bool {
		return sr1.lt(sr2)
	})

	for _, n := range nums {
		heap.Push(n)
	}

	for i := 0; i < times; i++ {
		sr := heap.Pop()
		sr.increment()
		fmt.Println("incrementing", sr.k)
		heap.Push(sr)
	}

	var sum int
	heap.Iter(func(sr *squareRepr) bool {
		sum = (sum + maths.PowMod(sr.k, maths.Pow(2, sr.twoPow), mod)) % mod
		fmt.Println("summing", sr.k, sr.twoPow)
		return true
	})
	return sum
}
