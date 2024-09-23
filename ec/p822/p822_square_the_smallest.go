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

type squareRepr struct {
	k   int
	pow int
	log float64
}

func (sr *squareRepr) increment() {
	sr.incrementTimes(1)
}

func (sr *squareRepr) incrementTimes(n int) {
	pow := maths.Pow(2, n)
	sr.log *= float64(pow)
	sr.pow *= pow
}

func (sr *squareRepr) incrementUpTo(max *squareRepr) int {
	var increments int
	for 2*sr.log < max.log {
		sr.increment()
		increments++
	}
	return increments
}

func P822() *ecmodels.Problem {
	return ecmodels.IntsInputNode(822, 2, 0, func(o command.Output, ns []int) {

		n, m := ns[0], ns[1]

		var srs []*squareRepr
		for i := 2; i <= n; i++ {
			fmt.Println("adding", i)
			srs = append(srs, &squareRepr{i, 1, math.Log(float64(i))})
		}

		// fmt.Println(brute(srs, m))
		fmt.Println(elegant(srs, m))

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

func elegant(nums []*squareRepr, times int) int {
	var increments int
	for _, n := range nums {
		increments += n.incrementUpTo(nums[len(nums)-1])
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
		return sr1.log < sr2.log
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
		sum = (sum + maths.PowMod(sr.k, sr.pow, mod)) % mod
		fmt.Println("summing", sr.k, sr.pow)
		return true
	})
	return sum
}
