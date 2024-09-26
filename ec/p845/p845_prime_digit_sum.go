package p845

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/combinatorics"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
	"golang.org/x/exp/slices"
)

var (
	p = generator.Primes()
)

func P845() *ecmodels.Problem {
	return ecmodels.IntInputNode(845, func(o command.Output, n int) {
		o.Stdoutln(search(n))
	}, []*ecmodels.Execution{
		{
			Args: []string{"61"},
			Want: "157",
		},
		{
			Args: []string{"100000000"},
			Want: "403539364",
		},
		{
			Args:     []string{"10000000000000000"},
			Want:     "45009328011709400",
			Estimate: 7,
		},
	})
}

func search(ni int) int {
	n := maths.NewInt(ni)

	left, right := 0, 1
	leftCnt, rightCnt := maths.Zero(), maths.Zero()

	for rightCnt.LT(n) {
		left, right = right, right*2
		leftCnt, rightCnt = rightCnt, dInverse(right)
	}

	if leftCnt.EQ(n) {
		return left
	}
	if rightCnt.EQ(n) {
		return right
	}

	for i := 0; i < 1000; i++ {
		mid := (left + right) / 2
		midCnt := dInverse(mid)

		midCmp := midCnt.Cmp(n)
		if midCmp == 0 {

			for ; !p.Contains(bread.Sum(maths.Digits(mid))); mid-- {
			}
			return mid
		} else if midCmp < 0 {
			left = mid
		} else {
			right = mid
		}
	}
	panic("NOOO")
}

func dInverse(n int) *maths.Int {
	ds := maths.Digits(n)

	sum := maths.Zero()

	for i := 0; i < len(ds); i++ {
		for k := 0; k < ds[i]; k++ {
			curSum := bread.Sum(ds[:i]) + k
			v := count(len(ds)-1-i, 9, curSum, nil)
			sum = sum.Plus(v)
		}
	}

	if p.Contains(bread.Sum(ds)) {
		sum = sum.PlusInt(1)
	}

	return sum
}

var (
	countCache = map[string]*maths.Int{}
)

func count(size, max, curSum int, counts []int) *maths.Int {
	countsCopy := bread.Copy(counts)
	slices.Sort(countsCopy)

	code := fmt.Sprintf("%d-%d-%d-%v", size, max, curSum, countsCopy)
	if v, ok := countCache[code]; ok {
		return v.Copy()
	}

	if size == 0 {
		if !p.Contains(curSum) {
			return maths.Zero()
		}

		r := combinatorics.PermutationFromCount(counts)
		countCache[code] = r
		return r
	}

	sum := maths.Zero()
	for i := max; i >= 0; i-- {
		for cnt := 1; cnt <= size; cnt++ {
			sum = sum.Plus(count(size-cnt, i-1, curSum+i*cnt, append(counts, cnt)))
		}
	}

	countCache[code] = sum
	return sum
}
