package p458

import (
	"fmt"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

type matrix458 struct {
	values [][]int
}

func (this *matrix458) String() string {
	var r []string
	for _, row := range this.values {
		r = append(r, fmt.Sprintf("%v", row))
	}
	return strings.Join(r, "\n")
}

func (this *matrix458) times(vector []int) []int {
	result := make([]int, len(vector), len(vector))
	for _, row := range this.values {
		for i, c := range row {
			result[i] += c * vector[i]
		}
	}
	return result
}

func (this *matrix458) mult(that *matrix458) *matrix458 {
	var result [][]int
	for iThis := 0; iThis < len(this.values[0]); iThis++ {
		var row []int
		for iThat := 0; iThat < len(that.values[0]); iThat++ {
			var sum int
			for j := 0; j < len(this.values[0]); j++ {
				sum += (this.values[iThis][j] * that.values[j][iThat]) % 1_000_000_000
			}
			row = append(row, sum%1_000_000_000)
		}
		result = append(result, row)
	}
	return &matrix458{result}
}

func finale458(size, n int) int {
	var m [][]int
	m = append(m, make([]int, size+1, size+1))
	for i := 1; i <= size; i++ {
		var row []int
		for j := 0; j <= size; j++ {
			if j == size {
				if i == size {
					row = append(row, size)
				} else {
					row = append(row, 0)
				}
			} else if j >= i {
				row = append(row, 1)
			} else if j == i-1 {
				row = append(row, size-j)
			} else {
				row = append(row, 0)
			}
		}
		m = append(m, row)
	}

	mp := map[int]*matrix458{
		1: {m},
	}
	// TODO: large exponent maths function
	base := mp[1]
	k := 2
	for ; k <= n; k *= 2 {
		base = mp[k/2].mult(mp[k/2])
		mp[k] = base
	}
	k /= 2

	rem := n - k + 1
	for ; k > 0 && rem != 0; k /= 2 {
		if k <= rem {
			base = base.mult(mp[k])
			rem -= k
		}
	}

	return base.values[1][0]
}

// brute458 keeps track of the number of strings that do not
// contain a permutation by tracking qn (see code comments).
// This is basically just a Markov chain (which is matrix multiplication).
func brute458(size, n int) int {

	// qn is size^n - tn (aka elements with the permutation)
	// qn[i] = number of values where the current state is an 'i' letter permutation
	// Once we get to state 'size', however, we stay there.
	qn := make([]int, size+1, size+1)
	// Start with all single letter strings
	qn[1] = size
	for i := 1; i < n; i++ {
		next := make([]int, size+1, size+1)

		for curPermSize := 1; curPermSize < size; curPermSize++ {
			// Some letters send us to the next state
			next[curPermSize+1] = qn[curPermSize] * (size - curPermSize)
			// Remaining letters send us back to all lower sizes
			for j := 1; j <= curPermSize; j++ {
				next[j] += qn[curPermSize]
			}
		}

		for i, v := range next {
			qn[i] = v % 1_000_000_000
		}
	}

	return bread.Sum(qn[:size])
}

func P458() *ecmodels.Problem {
	// This solution is just a matrix multiplication representation of the
	// solution implemented by brute458.
	return ecmodels.IntsInputNode(458, 2, 0, func(o command.Output, ns []int) {
		size := ns[0]
		n := ns[1]

		// Construct the markov matrix
		var m [][]int
		m = append(m, make([]int, size+1, size+1))
		for i := 1; i <= size; i++ {
			var row []int
			for j := 0; j <= size; j++ {
				if j == size {
					if i == size {
						row = append(row, size)
					} else {
						row = append(row, 0)
					}
				} else if j >= i {
					row = append(row, 1)
				} else if j == i-1 {
					row = append(row, size-j)
				} else {
					row = append(row, 0)
				}
			}
			m = append(m, row)
		}

		// Multiply the matrix to the n-th power
		// TODO: large exponent maths function using binary expression as we are here.
		mp := map[int]*matrix458{
			1: {m},
		}

		base := mp[1]
		k := 2
		for ; k <= n; k *= 2 {
			base = mp[k/2].mult(mp[k/2])
			mp[k] = base
		}
		k /= 2

		rem := n - k + 1
		for ; k > 0 && rem != 0; k /= 2 {
			if k <= rem {
				base = base.mult(mp[k])
				rem -= k
			}
		}

		// Think we actually need to do some vector multiplication,
		// but this position in the matrix also does the trick
		// (this required adding +1 to rem initialization).
		o.Stdoutln(base.values[1][0])
	}, []*ecmodels.Execution{
		{
			Args: []string{"3", "2"},
			Want: "9",
		},
		{
			Args: []string{"3", "3"},
			Want: "21",
		},
		{
			Args: []string{"4", "2"},
			Want: "16",
		},
		{
			Args: []string{"4", "3"},
			Want: "64",
		},
		{
			Args: []string{"4", "4"},
			Want: "232",
		},
		{
			Args: []string{"4", "5"},
			Want: "856",
		},
		{
			Args: []string{"4", "6"},
			Want: "3160",
		},
		{
			Args: []string{"7", "7"},
			Want: "818503",
		},
		{
			Args: []string{"7", "1_000_000_000_000"},
			Want: "423341841",
		},
	})
}
