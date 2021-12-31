package twentyone

import (
	"testing"

	"github.com/leep-frog/command"
)

func TestAll(t *testing.T) {
	for _, test := range []struct {
		name string
		node *command.Node
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
			etc := &command.ExecuteTestCase{
				Node:          test.node,
				Args:          test.args,
				WantStdout:    test.want,
				SkipDataCheck: true,
			}
			command.ExecuteTest(t, etc)
		})
	}
}
