package p303

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P303() *ecmodels.Problem {
	return ecmodels.IntInputNode(303, func(o command.Output, pow int) {

		n := uint64(maths.Pow(10, pow))

		// Used to use a map[uint64]bool for this, but iterating over keys was expensive
		var s []uint64
		for i := uint64(1); i <= n; i++ {
			s = append(s, i)
		}

		o.Stdoutln(dfs(s, &smallDigitIncrementer{}, n))
	}, []*ecmodels.Execution{
		{
			Args: []string{"2"},
			Want: "11363107",
		},
		{
			Args: []string{"4"},
			Want: "1111981904675169",
		},
	})
}

func markCompleted(s []uint64, sum, n, k, v uint64) ([]uint64, uint64) {
	remove := map[uint64]bool{}
	for cur := k; cur <= n; cur *= 10 {
		remove[cur] = true
		// This intentionally uses k because the ratio stays the same:
		// (v / k) == (10v / 10k) == (100v / 100k) == ...
		sum += v
	}

	var r []uint64
	for _, v := range s {
		if !remove[v] {
			r = append(r, v)
		}
	}
	return r, sum
}

func dfs(s []uint64, sdi *smallDigitIncrementer, n uint64) uint64 {

	var sum uint64

	// First, clear all the /9+/ values because they require the largest values
	// but create a very simple pattern (k ones followed by 4k twos)
	for nines, ones, twos := uint64(9), []int{1}, []int{2, 2, 2, 2}; nines <= n; nines, ones, twos = nines*10+9, append(ones, 1), append(twos, 2, 2, 2, 2) {
		s, sum = markCompleted(s, sum, n, nines, maths.IntFromDigits(append(ones, twos...)).DivInt(int(nines)).Int().Uint64())
	}

	for len(s) > 0 {
		v := sdi.next()

		for _, k := range s {
			if v%k == 0 {
				s, sum = markCompleted(s, sum, n, k, v/k)
			}
		}
	}
	return sum
}

type smallDigitIncrementer struct {
	prev uint64
}

func (sdi *smallDigitIncrementer) next() uint64 {
	tenPow := uint64(1)
	for cur := sdi.prev; cur > 0 && cur%10 == 2; cur, tenPow = cur/10, tenPow*10 {
		sdi.prev -= tenPow * 2
	}

	sdi.prev += tenPow
	return sdi.prev
}
