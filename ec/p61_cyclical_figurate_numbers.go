package eulerchallenge

import (
  "github.com/leep-frog/command"
)

func P61() *command.Node {
  return command.SerialNodes(
    command.Description("https://projecteuler.net/problem=61"),
    command.IntNode(N, "", command.IntPositive()),
    command.ExecutorNode(func(o command.Output, d *command.Data) {
      n := d.Int(N)
      o.Stdoutln(n)
    }),
  )
}