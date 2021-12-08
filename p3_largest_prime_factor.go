package main

import (
	"github.com/leep-frog/command"
)

// TODO: move all of these to helper directory
func getPrimeFactors(n int, p *Generator) map[int]int {
	r := map[int]int{}
	for i := 0; ; i++ {
		pi := p.Nth(i)
		for n%pi == 0 {
			r[pi]++
			n = n / pi
			if n == 1 {
				return r
			}
		}
	}
}

func p3() *command.Node {
	return command.SerialNodes(
		command.Description("Find the largest prime factor of N"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) error {
			factors := getPrimeFactors(d.Int(N), Primer())

			max := 0
			for f := range factors {
				if f > max {
					max = f
				}
			}

			o.Stdoutf("%d", max)
			return nil
		}),
	)
}
