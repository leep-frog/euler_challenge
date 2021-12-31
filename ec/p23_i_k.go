package eulerchallenge

import (
  "github.com/leep-frog/command"
  "github.com/leep-frog/euler_challenge/parse"
)

func P23() *command.Node {
  return command.SerialNodes(
    command.Description("stuff stuff stuff"),
    command.StringNode("FILE", ""),
    command.ExecutorNode(func(o command.Output, d *command.Data) {
      lines := parse.ReadFileLines(d.String("FILE"))
      o.Stdoutln(lines)
    }),
  )
}