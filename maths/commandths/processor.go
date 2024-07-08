package commandths

import (
	"fmt"

	"github.com/leep-frog/functional"
)

// Evaluate everything (assume parens are evaled)
// Then we are left with:
// # op # op # op # op

type numericalTerm struct {
	value      int
	prev, next *operationTerm
}

type operationTerm struct {
	op         Operation[int]
	prev, next *numericalTerm
	position   int
}

func (ne *numericalTerm) String() string {
	return ne.string(map[int]bool{})
}

func (ne *numericalTerm) string(m map[int]bool) string {
	if m[ne.value] {
		return fmt.Sprintf("NE(%d) CYCLE", ne.value)
	}
	m[ne.value] = true
	if ne.next == nil {
		return fmt.Sprintf("NE(%d)", ne.value)
	}
	return fmt.Sprintf("NE(%d) %v", ne.value, ne.next.String())
}

func (ot *operationTerm) String() string {
	return ot.string(map[int]bool{})
}

func (ot *operationTerm) string(m map[int]bool) string {
	if m[ot.position] {
		return fmt.Sprintf("NE(%d) CYCLE", ot.position)
	}
	m[ot.position] = true
	if ot.next == nil {
		return fmt.Sprintf("OT(%d)", ot.position)
	}
	return fmt.Sprintf("NE(%d) %v", ot.position, ot.next.String())
}

func (ne *numericalTerm) evaluate() int {
	var ops []*operationTerm
	got := map[int]bool{}
	for cur := ne; cur.next != nil && cur.next.next != nil; cur = cur.next.next {
		if got[cur.next.position] {
			panic("Infinite loop")
		}
		got[cur.next.position] = true
		ops = append(ops, cur.next)
	}
	functional.SortFunc(ops, func(this, that *operationTerm) bool {
		if this.op.PemdasPriority() != that.op.PemdasPriority() {
			return this.op.PemdasPriority() < that.op.PemdasPriority()
		}
		return this.position < that.position
	})

	finalValue := ne
	for _, op := range ops {
		finalValue = op.evaluate()
	}
	return finalValue.value
}

func (ot *operationTerm) evaluate() *numericalTerm {
	newTerm := &numericalTerm{
		ot.op.Evaluate(ot.prev.value, ot.next.value),
		nil, nil,
	}
	if ot.prev.prev != nil {
		ot.prev.prev.next = newTerm
		newTerm.prev = ot.prev.prev
	}
	if ot.next.next != nil {
		ot.next.next.prev = newTerm
		newTerm.next = ot.next.next
	}
	return newTerm
}
