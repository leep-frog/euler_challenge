package eulerchallenge

import (
	"sort"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/parse"
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
		uf := NewUnionFind()
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

// UnionFind implements a union find object for keeping track of distinct groups.
type UnionFind struct {
	// elementMap is a map from element to group
	elementMap map[int]int
	// setMap is a from group number to vertices in that set
	setMap map[int]map[int]bool
	// groupCount is used to keep track of the next group number
	groupCount int
}

// NewUnionFind returns an initialized UnionFind object.
func NewUnionFind() *UnionFind {
	return &UnionFind{
		map[int]int{},
		map[int]map[int]bool{},
		0,
	}
}

// Merge merges the groups for a and b. If a and b are already in the same group
// then nothing happens and false is returned. If neither a nor b is in a group,
// then a new group is created. If one of the elements isn't in a group, then it
// is merged into the other groups. If both are in a group, then the groups
// are merged.
func (uf *UnionFind) Merge(a, b int) bool {
	if uf.Connected(a, b) {
		return false
	}

	ag, aInGroup := uf.elementMap[a]
	bg, bInGroup := uf.elementMap[b]
	if !aInGroup && !bInGroup {
		uf.groupCount++
		uf.elementMap[a] = uf.groupCount
		uf.elementMap[b] = uf.groupCount
		uf.setMap[uf.groupCount] = map[int]bool{
			a: true,
			b: true,
		}
	} else if !aInGroup {
		uf.setMap[bg][a] = true
		uf.elementMap[a] = bg
	} else if !bInGroup {
		uf.setMap[ag][b] = true
		uf.elementMap[b] = ag
	} else {
		for v := range uf.setMap[ag] {
			uf.setMap[bg][v] = true
			uf.elementMap[v] = bg
		}
		delete(uf.setMap, ag)
	}
	return true
}

// Connected returns whether or not a and b are in the same set.
func (uf *UnionFind) Connected(a, b int) bool {
	ag, aInGroup := uf.elementMap[a]
	bg, bInGroup := uf.elementMap[b]
	return aInGroup && bInGroup && ag == bg
}
