package main

import (
	"github.com/leep-frog/command"
)

func p12() *command.Node {
	return command.SerialNodes(
		command.Description("Find a triangular number with at least N divisors"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) error {
			ts := Triangulars()
			n := d.Int(N)
			for i := ts.Next(); ; i = ts.Next() {
				var count int
				max := i / 2
				for j := 1; j < max; j++ {
					if i%j == 0 {
						max = (i / j) - 1
						if j*j == i {
							count++
						} else {
							count += 2
						}
						if count > n {
							o.Stdoutf("%d", i)
							return nil
						}
					}
				}
			}
		}),
	)
}
