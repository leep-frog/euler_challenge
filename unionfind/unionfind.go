package unionfind

import (
	"github.com/leep-frog/euler_challenge/maths"
	"golang.org/x/exp/maps"
)

// UnionFind implements a union find object for keeping track of distinct groups.
type UnionFind[T comparable] struct {
	// elementMap is a map from element to group
	elementMap map[T]int
	// setMap is a map from group number to vertices in that set
	setMap map[int]map[T]bool
	// groupCount is used to keep track of the next group number
	groupCount int
}

// NewUnionFind returns an initialized UnionFind object.
func New[T comparable]() *UnionFind[T] {
	return &UnionFind[T]{
		map[T]int{},
		map[int]map[T]bool{},
		0,
	}
}

func (uf *UnionFind[T]) Sets() []map[T]bool {
	var r []map[T]bool
	for _, k := range uf.setMap {
		r = append(r, maths.CopyMap(k))
	}
	return r
}

func (uf *UnionFind[T]) NumberOfSets() int {
	return len(uf.setMap)
}

// Insert inserts a single element as a set. If it is already in a set
// then nothing happens
func (uf *UnionFind[T]) Insert(a T) {
	if _, ok := uf.elementMap[a]; ok {
		return
	}

	uf.groupCount++
	uf.elementMap[a] = uf.groupCount
	uf.setMap[uf.groupCount] = map[T]bool{
		a: true,
	}
}

// Merge merges the groups for a and b. If a and b are already in the same group
// then nothing happens and false is returned. If neither a nor b is in a group,
// then a new group is created. If one of the elements isn't in a group, then it
// is merged into the other groups. If both are in a group, then the groups
// are merged.
func (uf *UnionFind[T]) Merge(a, b T) bool {
	if uf.Connected(a, b) {
		return false
	}

	ag, aInGroup := uf.elementMap[a]
	bg, bInGroup := uf.elementMap[b]
	if !aInGroup && !bInGroup {
		uf.groupCount++
		uf.elementMap[a] = uf.groupCount
		uf.elementMap[b] = uf.groupCount
		uf.setMap[uf.groupCount] = map[T]bool{
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
func (uf *UnionFind[T]) Connected(a, b T) bool {
	ag, aInGroup := uf.elementMap[a]
	bg, bInGroup := uf.elementMap[b]
	return aInGroup && bInGroup && ag == bg
}

// Elements returns all of the elements that have been connected
// at least once.
func (uf *UnionFind[T]) Elements() []T {
	return maps.Keys(uf.elementMap)
}
