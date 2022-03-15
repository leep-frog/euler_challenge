package eulerchallenge

import (
	"sort"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/unionfind"
)

func P107() *problem {
	return fileInputNode(107, func(lines []string, o command.Output) {
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
		uf := unionfind.New()
		for _, edge := range edges {
			totalWeight += edge.weight
			if uf.Merge(edge.vertexA, edge.vertexB) {
				mstWeight += edge.weight
			}
		}
		o.Stdoutln(totalWeight - mstWeight)
	})
}

type edge107 struct {
	vertexA int
	vertexB int
	weight  int
}
