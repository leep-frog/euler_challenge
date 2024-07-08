package eulerchallenge

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P158() *problem {
	return intInputNode(158, func(o command.Output, letters int) {

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
	}, []*execution{
		{
			args: []string{"3"},
			want: "10400",
		},
		{
			args: []string{"4"},
			want: "164450",
		},
		{
			args: []string{"5"},
			want: "1710280",
		},
		{
			args: []string{"26"},
			want: "409511334375",
		},
	})
}
