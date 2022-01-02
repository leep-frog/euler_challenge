package eulerchallenge

import (
	"fmt"
	"sort"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P29() *command.Node {
	return command.SerialNodes(
		command.Description("Get the number of distinct terms generated by a^b where a and b are between 2 and n"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			n := d.Int(N)
			p := generator.Primes()

			unique := map[string]bool{}
			for a := 2; a <= n; a++ {
				factors := generator.PrimeFactors(a, p)
				for b := 2; b <= n; b++ {
					scaled := map[int]int{}
					for k, v := range factors {
						scaled[k] = v * b
					}
					unique[polyCode(scaled)] = true
				}
			}

			o.Stdoutln(len(unique))
		}),
	)
}

func polyCode(m map[int]int) string {
	var a []int
	for k := range m {
		a = append(a, k)
	}
	sort.Ints(a)

	var r []string
	for _, k := range a {
		r = append(r, fmt.Sprintf("(%d, %d)", k, m[k]))
	}
	return strings.Join(r, ", ")
}
