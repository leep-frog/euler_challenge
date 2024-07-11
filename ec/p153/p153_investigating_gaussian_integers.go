package p153

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
)

func P153() *ecmodels.Problem {
	return ecmodels.IntInputNode(153, func(o command.Output, n int) {
		p := generator.Primes()
		var ss int
		counts := make([]int, n/2+1, n/2+1)
		for i := 1; i <= n/2; i++ {
			if i%1_000_000 == 0 {
				fmt.Println(i)
			}
			for _, k := range p.Factors(i) {
				ss += k
			}
			counts[i] = ss
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
				if p.Coprimes(i, j) && (i != 1) && (j != 1) {
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

		return
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
			Estimate: 300,
		},
	})
}
