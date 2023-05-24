package topology

type NodeWithDependencies[ID comparable] interface {
	Code() ID
	Dependencies() []ID
}

type DependencyGraph[ID comparable, N NodeWithDependencies[ID]] struct {
	// Dependencies is a map from node to list of nodes that the node depends on.
	Dependencies map[ID][]ID
	// InvertedDependencies is a map from node to list of nodes that depend on that node.
	InvertedDependencies map[ID][]ID
	// Roots is the list of nodes which have no dependencies.
	Roots []N
	// InvertedRoots is the list of nodes which have nothing depending on them.
	InvertedRoots []N
	Nodes         map[ID]N
}

func GetDependencyGraph[ID comparable, N NodeWithDependencies[ID]](nodes []N) *DependencyGraph[ID, N] {
	dg := &DependencyGraph[ID, N]{
		map[ID][]ID{},
		map[ID][]ID{},
		nil, nil,
		map[ID]N{},
	}

	for _, node := range nodes {
		c := node.Code()
		dg.Nodes[c] = node

		for _, d := range node.Dependencies() {
			dg.InvertedDependencies[d] = append(dg.InvertedDependencies[d], c)
			dg.Dependencies[c] = append(dg.Dependencies[c], d)
		}
	}

	for c, node := range dg.Nodes {
		if len(dg.Dependencies[c]) == 0 {
			dg.Roots = append(dg.Roots, node)
		}
		if len(dg.InvertedDependencies[c]) == 0 {
			dg.InvertedRoots = append(dg.InvertedRoots, node)
		}
	}
	return dg
}
