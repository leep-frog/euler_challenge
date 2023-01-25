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
// It is required to avoid an interface cycle in `Topological`.
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

/*func recursiveProcess[T Topological[V], V any](item T, itemMap map[string]T, dependents map[string][]string, dependencies map[string]map[string]bool, values map[string]V) {
	c := item.Code()
	// Check if still waiting on dependencies.
	if len(dependencies[c]) != 0 {
		return
	}

	// Check if already processed.
	if _, ok := values[c]; ok {
		return
	}

	// Process the item.
	values[c] = item.Process(values)

	// Process dependents
	for _, dc := range dependents[c] {
		d := itemMap[dc]
		// d is no longer waiting on c
		deps := dependencies[dc]
		delete(deps, c)
		if len(deps) == 0 {
			recursiveProcess(d, itemMap, dependents, dependencies, values)
		}
	}
}

func Process[T Topological[V], V any](items []T) map[string]V {
	// TODO: Can convert string codes to ints so these can be int slices instead of maps

	// Map from item to list of items that depend on that item.
	dependents := map[string][]string{}
	// Map from item to list of items that the item is still waiting on.
	dependencies := map[string]map[string]bool{}
	values := map[string]V{}
	itemMap := map[string]T{}

	for _, item := range items {
		c := item.Code()
		itemMap[c] = item

		deps := item.Dependencies()
		dependencies[c] = map[string]bool{}
		for _, d := range deps {
			dependents[d] = append(dependents[d], c)
			dependencies[c][d] = true
		}
	}

	for _, item := range items {
		recursiveProcess(item, itemMap, dependents, dependencies, values)
	}

	return values
}
*/
