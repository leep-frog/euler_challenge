package p158

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P158() *ecmodels.Problem {
	return ecmodels.IntInputNode(158, func(o command.Output, letters int) {

		best := maths.LargestT[int, *maths.Int]()

		for n := 3; n <= letters; n++ {
			// First get a fixed set of letters. There are (26 choose i) of these.
			// Each set is generically identical to any other set because they all
			// have a single fixed ordering.
			coef := maths.Choose(26, n)

			// Now iterate over the spot where the lexicographically adjacent letters are.
			sum := maths.Zero()
			for spot := 0; spot < n; spot++ {
				// Now choose the letters that are lexicographically decreasing.
				// Subtract one because the only invalid set is the 'spot' highest
				// letters (e.g. 'zyx' in the spot=3 case).
				left := maths.Choose(n, spot+1).Minus(maths.One())

				// Now, assuming we have a valid set of left half letters,
				// there is only one option for the right half since they all must
				// be in reverse lexicographical order. So we simply
				// add the number to our sum
				sum = sum.Plus(left)
			}

			// Finally multiply together
			best.IndexCheck(n, coef.Times(sum))
		}

		o.Stdoutln(best.Best())
	}, []*ecmodels.Execution{
		{
			Args: []string{"3"},
			Want: "10400",
		},
		{
			Args: []string{"4"},
			Want: "164450",
		},
		{
			Args: []string{"5"},
			Want: "1710280",
		},
		{
			Args: []string{"26"},
			Want: "409511334375",
		},
	})
}
