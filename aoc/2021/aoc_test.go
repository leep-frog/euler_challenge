package twentyone

import (
	"strings"
	"testing"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/command/commandertest"
	"github.com/leep-frog/command/commandtest"
)

func TestAll(t *testing.T) {
	for _, test := range []struct {
		name string
		node command.Node
		args []string
		want []string
	}{
		{
			name: "2021 d25",
			node: D25(),
			want: []string{""},
		},
		/*{
			name: "2021 d24",
			node: D24(),
			want: []string{""},
		},*/
	} {
		t.Run(test.name, func(t *testing.T) {
			etc := &commandtest.ExecuteTestCase{
				Node:          test.node,
				Args:          test.args,
				WantStdout:    strings.Join(test.want, "\n"),
				SkipDataCheck: true,
			}
			commandertest.ExecuteTest(t, etc)
		})
	}
}
