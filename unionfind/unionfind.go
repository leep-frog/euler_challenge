package unionfind

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
func New() *UnionFind {
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
