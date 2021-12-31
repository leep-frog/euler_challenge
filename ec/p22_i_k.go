package eulerchallenge

import (
  "github.com/leep-frog/command"
)

func P22() *command.Node {
  return command.SerialNodes(
    command.Description("stuff stuff stuff"),
    command.IntNode(N, "", command.IntPositive()),
    command.ExecutorNode(func(o command.Output, d *command.Data) {
      n := d.Int(N)
      o.Stdoutln(n)
    }),
  )
}