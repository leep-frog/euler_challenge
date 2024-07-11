package p191

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P191() *ecmodels.Problem {
	return ecmodels.IntInputNode(191, func(o command.Output, n int) {

		prev := &prizeStringStatus{
			zeroAbsentZeroLate: 1,
		}

		for i := 0; i < n; i++ {
			next := &prizeStringStatus{}

			// First, all of the invalids can put anything next
			next.invalids = prev.invalids * 3

			// zeroAbsentZeroLate
			next.zeroAbsentZeroLate += prev.zeroAbsentZeroLate // On time
			next.zeroAbsentOneLate += prev.zeroAbsentZeroLate  // Late
			next.oneAbsentZeroLate += prev.zeroAbsentZeroLate  // Absent

			// oneAbsentZeroLate
			next.zeroAbsentZeroLate += prev.oneAbsentZeroLate // On time
			next.zeroAbsentOneLate += prev.oneAbsentZeroLate  // Late
			next.twoAbsentZeroLate += prev.oneAbsentZeroLate  // Absent

			// twoAbsentZeroLate
			next.zeroAbsentZeroLate += prev.twoAbsentZeroLate // On time
			next.zeroAbsentOneLate += prev.twoAbsentZeroLate  // Late
			next.invalids += prev.twoAbsentZeroLate           // Absent

			// zeroAbsentOneLate
			next.zeroAbsentOneLate += prev.zeroAbsentOneLate // On time
			next.invalids += prev.zeroAbsentOneLate          // Late
			next.oneAbsentOneLate += prev.zeroAbsentOneLate  // Absent

			// oneAbsentOneLate
			next.zeroAbsentOneLate += prev.oneAbsentOneLate // On time
			next.invalids += prev.oneAbsentOneLate          // Late
			next.twoAbsentOneLate += prev.oneAbsentOneLate  // Absent

			// twoAbsentOneLate
			next.zeroAbsentOneLate += prev.twoAbsentOneLate // On time
			next.invalids += prev.twoAbsentOneLate          // Late
			next.invalids += prev.twoAbsentOneLate          // Absent

			prev = next
		}

		o.Stdoutln(maths.Pow(3, n) - prev.invalids)
	}, []*ecmodels.Execution{
		{
			Args: []string{"4"},
			Want: "43",
		},
		{
			Args: []string{"30"},
			Want: "1918080160",
		},
	})
}

type prizeStringStatus struct {
	zeroAbsentZeroLate int
	oneAbsentZeroLate  int
	twoAbsentZeroLate  int

	zeroAbsentOneLate int
	oneAbsentOneLate  int
	twoAbsentOneLate  int

	invalids int
}
