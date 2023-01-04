package y2022

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day20() aoc.Day {
	return &day20{}
}

/*type trackedNode[T any] struct {
	node *sqrtNode[T]
	pos  int
}

type sqrtNode[T any] struct {
	next, prev  *sqrtNode[T]
	value       T
	trackedNode *trackedNode[T]
}

type sqrtLinkedList[T any] struct {
	tNodes []*trackedNode[T]
}

func (sll *sqrtLinkedList[T]) elements() []*sqrtNode[T] {
	ts := []*sqrtNode[T]{sll.tNodes[0].node}

	for n := sll.tNodes[1].node; n.trackedNode != nil && n.trackedNode.pos != 0; n = n.next {
		ts = append(ts, n)
	}
}

func newSqrtLinkedList[T any](ts []T) *sqrtLinkedList[T] {
	// lg := int(math.Log2(float64(len(elements))))
	lg := maths.Sqrt(len(ts))

	elements := parse.Map(ts, func(t T) *sqrtNode[T] {
		return &sqrtNode[T]{nil, nil, t}
	})

	for i, e := range elements {
		e.next = elements[(i+1)%len(elements)]
		e.prev = elements[(i+len(elements)-1)%len(elements)]
	}

	var trackedNodes []*trackedNode[T]
	for i := 0; i < len(elements); i += lg {
		trackedNodes = append(trackedNodes, &trackedNode[T]{elements[i], i})
	}

	return &sqrtLinkedList[T]{trackedNodes}
}

/*func (d *day20) sort(trackedNodes []*trackedNode) {
	slices.SortFunc(trackedNodes, func(this, that *trackedNode) bool {
		return this.pos < that.pos
	})
}*/

type day20 struct{}

type simpleNode struct {
	prev, next *simpleNode
	value      int
	// did        bool
}

func (sn *simpleNode) String() string {
	return fmt.Sprintf("%d", sn.value)
}

func printOrder(nodes []*simpleNode) []int {
	start := nodes[0]
	var order []int
	for i := 0; i < len(nodes); i++ {
		order = append(order, start.value)
		start = start.next
	}

	fmt.Println(order)
	return order
}

func (d *day20) brute(elements []int) {
	nodes := parse.Map(elements, func(k int) *simpleNode {
		return &simpleNode{nil, nil, k}
	})

	for i, n := range nodes {
		n.next = nodes[(i+1)%len(nodes)]
		n.prev = nodes[(i+len(nodes)-1)%len(nodes)]
	}

	for _, n := range nodes {
		// moves := n.value % (len(elements) - 1)
		// fmt.Println(n.value, moves)
		// if moves < 0 {
		// 	moves = (len(elements) - 1) + moves + (len(elements) - 1) - 1
		// }

		// moves = moves % (len(elements) - 1)

		// if moves == 0 {
		// 	continue
		// }

		// fmt.Println("MOVING", n.value, moves)
		moves := n.value
		if moves > 0 {
			for i := 0; i < moves; i++ {
				n.prev.next, n.prev, n.next, n.next.prev, n.next.next, n.next.next.prev = n.next, n.next, n.next.next, n.prev, n, n
			}
		} else if moves < 0 {
			for i := 0; i < -moves; i++ {
				n.next.prev, n.next, n.prev, n.prev.next, n.prev.prev, n.prev.prev.next = n.prev, n.prev, n.prev.prev, n.next, n, n
			}
		}
		// printOrder(nodes)
	}

	ordered := printOrder(nodes)

	var zeroIdx int
	for i, n := range ordered {
		if n == 0 {
			zeroIdx = i
			fmt.Println("HEYO", i)
			break
		}
	}

	a := ordered[(zeroIdx+1000)%len(ordered)]
	b := ordered[(zeroIdx+2000)%len(ordered)]
	c := ordered[(zeroIdx+3000)%len(ordered)]
	fmt.Println(a, b, c, a+b+c, len(ordered))

	// for {
	// 	i := 0
	// 	for ; i < len(nodes) && nodes[i].did; i++ {
	// 	}
	// 	if i == len(nodes) {
	// 		break
	// 	}
	// 	n := nodes[i]
	// 	n.did = true

	// 	moveTo := i + n.value
	// 	if moveTo < 0 {
	// 		moveTo--
	// 	}
	// 	fmt.Println("MT", i, n.value, moveTo)
	// 	for moveTo < 0 {
	// 		moveTo += len(elements)
	// 	}
	// 	moveTo = moveTo % len(elements)

	// 	direction := 1
	// 	if moveTo < i {
	// 		direction = -1
	// 	}

	// 	if moveTo == i {
	// 		continue
	// 	}

	// 	fmt.Println("MOVING", n.value, direction, moveTo)
	// 	for start := i; start != moveTo; start += direction {
	// 		nodes[start], nodes[start+direction] = nodes[start+direction], nodes[start]
	// 	}
	// 	fmt.Println(nodes)
	// }

	// for i, n := range nodes {
	// 	n.prev = nodes[(i+len(nodes)-1)%len(nodes)]
	// 	n.next = nodes[(i+1)%len(nodes)]
	// }

	// for _, n := range nodes
}

func (d *day20) Solve(lines []string, o command.Output) {
	elements := parse.Map(lines, func(s string) int {
		return parse.Atoi(s)
	})

	d.brute(elements)

	// for _, e := range elements {
	// 	// First, get the tracked node to the left of it
	// 	var offset int
	// 	for ; e.tracked != nil; e = e.prev {
	// 		offset++
	// 	}
	// 	fromTracked := e.tracked

	// 	// Where are we jumping to
	// 	toPos := e.value % len(elements)
	// 	var i int
	// 	for i+1 < len(trackedNodes) && trackedNodes[i+1].pos >= toPos {
	// 	}
	// 	toTracked := trackedNodes[i]
	// }
}

func (d *day20) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"",
			},
		},
		{
			ExpectedOutput: []string{
				"",
			},
		},
	}
}
