package p107

import (
	"sort"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/unionfind"
)

func P107() *ecmodels.Problem {
	return ecmodels.FileInputNode(107, func(lines []string, o command.Output) {
		var edges []*edge107
		for a, line := range lines {
			for b, s := range strings.Split(line, ",") {
				if b >= a {
					break
				}
				if s == "-" {
					continue
				}
				edges = append(edges, &edge107{a, b, parse.Atoi(s)})
			}
		}
		sort.SliceStable(edges, func(i, j int) bool { return edges[i].weight < edges[j].weight })

		var totalWeight, mstWeight int
		uf := unionfind.New[int]()
		for _, edge := range edges {
			totalWeight += edge.weight
			if uf.Merge(edge.vertexA, edge.vertexB) {
				mstWeight += edge.weight
			}
		}
		o.Stdoutln(totalWeight - mstWeight)
	}, []*ecmodels.Execution{
		{
			Args: []string{"p107_network.txt"},
			Want: "259679",
		},
		{
			Args: []string{"p107_example.txt"},
			Want: "150",
		},
	})
}

type edge107 struct {
	vertexA int
	vertexB int
	weight  int
}
