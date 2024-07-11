package p148

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

func P148() *ecmodels.Problem {
	return ecmodels.IntInputNode(148, func(o command.Output, ni int) {
		n := uint64(ni)
		f, start := pascalDivBySeven(n)
		f += brutePascalDivBySeven(start, n)
		o.Stdoutln(n*(n+1)/2 - f)
	}, []*ecmodels.Execution{
		{
			Args:     []string{"1000000000"},
			Want:     "2129970655314432",
			Estimate: 7,
		},
		{
			Args: []string{"100"},
			Want: "2361",
		},
	})
}

// pascalDivBySeven returns the number of values that are
// divisble by 7, but only up to a k*7^x (the second return value)
// for the largest possible x.
func pascalDivBySeven(n uint64) (uint64, uint64) {
	if n < 7 {
		return 0, 0
	}
	seven := uint64(1)
	for ; seven*7 < n; seven *= 7 {
	}

	k := n / seven
	numBig := k * (k - 1) / 2
	sizeBig := seven * (seven - 1) / 2

	numFractured := k * (k + 1) / 2

	lf, _ := pascalDivBySeven(seven)
	return numBig*sizeBig + numFractured*lf, k * seven
}

// brutePascalDivBySeven returns the numbers of elements in rows start through n
// of Pascal's triangle that are divisble by 7.
func brutePascalDivBySeven(start, n uint64) uint64 {
	var sum uint64
	for r := start; r < n; r++ {
		var count uint64
		coef := uint64(1)
		for prevSeven, seven := uint64(1), uint64(7); seven <= r; prevSeven, seven = seven, seven*7 {
			mod := (r % seven) / (prevSeven)
			count += (r / seven) * (6 - mod) * coef
			coef *= mod + 1
		}
		sum += count
	}
	return sum
}
