package p153

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P153() *ecmodels.Problem {
	return ecmodels.IntInputNode(153, func(o command.Output, n int) {

		// Create an array of the sum of factors for each number
		counts := make([]int, n/2+1)
		for i := 1; i <= n/2; i++ {
			for j := i; j <= n/2; j += i {
				counts[j] += i
			}
		}

		// Cumulate the array
		for i := 0; i < len(counts)-1; i++ {
			counts[i+1] += counts[i]
		}

		var actualTotal int
		for i := 1; i <= n/2; i++ {
			// Number of integers with i as a factor = n/i
			actualTotal += (n / i) * i
		}
		// Numbers greater than n/2 only occur once
		start := (n / 2) + 1
		end := n
		// start + (start+1) + (start+2) + ... + (end-2) + (end-1) + end
		// = start*(end-start+1) + 1 + 2 + 3 + ... + (end - start)
		// = start*(end-start+1) + (end - start)*(end-start+1)/2
		actualTotal += start*(end-start+1) + (end-start)*(end-start+1)/2

		var iSums int
		for i := 1; i*i <= n; i++ {
			for j := i; j*j+i*i <= n; j++ {
				if !maths.Coprime(i, j) && (i != 1) && (j != 1) {
					continue
				}
				val := 2 * (i + j)
				upTo := n / (j*j + i*i)
				offset := counts[upTo] * val
				if i == j {
					offset /= 2
				}
				iSums += offset
			}
		}
		o.Stdoutln(iSums + actualTotal)
	}, []*ecmodels.Execution{
		{
			Args: []string{"5"},
			Want: "35",
		},
		{
			Args: []string{"100000"},
			Want: "17924657155",
		},
		{
			Args:     []string{"100000000"},
			Want:     "17971254122360635",
			Estimate: 10,
		},
	})
}
