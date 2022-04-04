package eulerchallenge

import (
  "github.com/leep-frog/command"
)

func P148() *problem {
  return intInputNode(148, func(o command.Output, n int) {
    o.Stdoutln(n)
  })
}