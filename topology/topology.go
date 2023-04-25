// Package topology implements algorithms for topological sorting
// and processing.
package topology

import "fmt"

// Graph is a topoligcal graph that calculates and caches
// data at each topological node.
type Graph[V any] struct {
	items  map[string]Node[V]
	Values map[string]V
}

// NewGraph constructs a new topological graph from the provided inputs.
func NewGraph[V any](items []Node[V]) *Graph[V] {
	itemMap := map[string]Node[V]{}
	for _, item := range items {
		itemMap[item.Code()] = item
	}
	return &Graph[V]{itemMap, map[string]V{}}
}

// Get retrieves the value for the item with the provided key.
func (g *Graph[V]) Get(key string) V {
	if v, ok := g.Values[key]; ok {
		return v
	}
	d, ok := g.items[key]
	if !ok {
		panic(fmt.Sprintf("Unknown topological key: %q", key))
	}
	v := d.Process(g)
	g.Values[key] = v
	return v
}

// Graphical is an intermediate interface which just wraps the `Graph` type.
// It is required to avoid an interface cycle in `Graph -> Node -> Graph`.
type Graphical[V any] interface {
	// Get retrieves the value for the item with the provided key.
	Get(string) V
}

// Node is a topological node that depends on zero or more other
// topological nodes.
type Node[V any] interface {
	// Code returns a unique string identifier of the node.
	Code() string
	// Process calculates the value for the node, given the entire
	// topological `Graph` which can be used to get values for dependencies.
	Process(Graphical[V]) V
}

// Topological is an interface used by the `Process` function.
type TopologicalNode[CTX any, ID comparable] interface {
	// A unique identifier of the
	Code(CTX) ID
	Dependencies(CTX) []ID
	Process(CTX)
}

// Process processes the set of nodes provided in topological order.
func Process[CTX any, T TopologicalNode[CTX, ID], ID comparable](ctx CTX, items []T) {
	// Map from item to list of items that depend on that item.
	dependents := map[ID][]ID{}
	// Map from item to list of items that the item is still waiting on.
	dependencies := map[ID]map[ID]bool{}
	itemMap := map[ID]T{}

	for _, item := range items {
		c := item.Code(ctx)
		itemMap[c] = item

		deps := item.Dependencies(ctx)
		dependencies[c] = map[ID]bool{}
		for _, d := range deps {
			dependents[d] = append(dependents[d], c)
			dependencies[c][d] = true
		}
	}

	processed := map[ID]bool{}
	for _, item := range items {
		recursiveProcess(ctx, item, itemMap, dependents, dependencies, processed)
	}
}

// recursiveProcess recursively processes the provided node and any downstream nodes
// that are now unblocked after processing the provided node.
func recursiveProcess[CTX any, T TopologicalNode[CTX, ID], ID comparable](ctx CTX, item T, itemMap map[ID]T, dependents map[ID][]ID, dependencies map[ID]map[ID]bool, processed map[ID]bool) {
	c := item.Code(ctx)
	// Check if still waiting on dependencies.
	if len(dependencies[c]) != 0 {
		return
	}

	// Check if already processed.
	if processed[c] {
		return
	}
	processed[c] = true

	// Process the item.
	item.Process(ctx)

	// Process dependents
	for _, dc := range dependents[c] {
		d := itemMap[dc]
		// d is no longer waiting on c
		deps := dependencies[dc]
		delete(deps, c)
		if len(deps) == 0 {
			recursiveProcess(ctx, d, itemMap, dependents, dependencies, processed)
		}
	}
}
