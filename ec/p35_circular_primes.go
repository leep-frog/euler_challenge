package eulerchallenge

import (
	"strconv"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func P35() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=35"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			n := d.Int(N)

			checked := map[string]bool{}
			unique := map[string]bool{}
			p := generator.Primes()
			for i := 0; p.Nth(i) < n; i++ {
				pn := p.Nth(i)
				spn := strconv.Itoa(pn)
				if checked[spn] {
					continue
				}
				var digits []string
				for j := 0; j < len(spn); j++ {
					digits = append(digits, spn[j:j+1])
				}

				allPrime := true
				rots := maths.Rotations(digits)
				for _, rot := range rots {
					checked[rot] = true
					if !generator.IsPrime(parse.Atoi(rot), p) {
						allPrime = false
					}
				}
				if allPrime {
					for _, rot := range rots {
						unique[rot] = true
					}
				}
			}
			o.Stdoutln(len(unique))
		}),
	)
}
