package p872

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P872() *ecmodels.Problem {
	return ecmodels.NoInputWithExampleNode(872, func(o command.Output, ex bool) {
		t, v := maths.BigPow(10, 17), maths.BigPow(9, 17)
		if ex {
			t, v = maths.NewInt(10), maths.NewInt(3)
		}
		o.Stdoutln(smartSolve(t, v))
	}, []*ecmodels.Execution{
		{
			Args: []string{"-x"},
			Want: "29",
		},
		{
			Want: "2903144925319290239",
		},
	})
}

// All children of the root are v-1, v-2, v-4, v-8, etc.
// All children of leaf nodes, start at v-k, v-2*k, etc. where k is 2 times the previous subtraction
// Ultimately, the solution for f(a, b) is:
// * Get the difference between a and b in binary bit representation.
// * Iterate from smallest bit to biggest bit
// * If bit is a 1, add (a-bitValue) to the sum
func smartSolve(a, b *maths.Int) *maths.Int {
	sum := a.Copy()
	at := a.Copy()
	for diff, pow := a.Minus(b), maths.One(); diff.GT(maths.Zero()); diff, pow = diff.DivInt(2), pow.TimesInt(2) {
		if diff.ModInt(2) == 1 {
			at = at.Minus(pow)
			sum = sum.Plus(at)
		}
	}
	return sum
}

// Deprecated, see smartSolve
/*func smartSolveSmall(a, b int) int {
	sum := a
	at := a
	for diff, pow := a-b, 1; diff > 0; diff, pow = diff/2, pow*2 {

		if diff%2 == 1 {
			at -= pow
			sum += at

		}
	}
	return sum
}

// Deprecated, see smartSolve
func bruteSolve(a, b int) int {
	return tree(a).get(b)
}

func tree(a int) *node {
	root := newNode(1)
	for ; root.v < a; root = root.next() {
	}
	return root
}

type node struct {
	v        int
	children *maths.Heap[*node]
}

func newNode(v int) *node {
	return &node{
		v,
		maths.NewHeap[*node](func(n1, n2 *node) bool {
			return n1.v > n2.v
		}),
	}
}

func (n *node) String() string {
	return strings.Join(n.string(""), "\n")
}

func (n *node) string(indent string) []string {
	r := []string{
		fmt.Sprintf("%s%d", indent, n.v),
	}
	n.children.Iter(func(kid *node) bool {
		r = append(r, kid.string(indent+"  ")...)
		return true
	})
	return r
}

func (n *node) next() *node {
	root := newNode(n.v + 1)

	for cur := n; ; cur = cur.children.Pop() {
		root.children.Push(cur)
		if cur.children.Len() == 0 {
			break
		}
	}

	return root
}

func (n *node) get(target int) int {
	if n.v == target {
		return n.v
	}

	got := -1
	n.children.Iter(func(kid *node) bool {
		got = kid.get(target)
		return got == -1
	})

	if got == -1 {
		return -1
	}
	return got + n.v
}
*/
