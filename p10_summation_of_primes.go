package main

import (
	"github.com/leep-frog/command"
)

func p10() *command.Node {
	return command.SerialNodes(
		command.Description("Find the sum of all primes lower than N"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) error {
			p := &Primer{}

			for i := 0; p.Next() < d.Int(N); i++ {
			}

			sum := 0
			for i := 0; i < len(p.Primes)-1; i++ {
				sum += p.Primes[i]
			}
			o.Stdoutf("%d", sum)
			return nil
		}),
	)
}
