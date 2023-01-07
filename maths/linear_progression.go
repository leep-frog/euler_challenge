package maths

import "fmt"

func NewLinearProgression(start, offset int) *LinearProgression {
	return NewBigLinearProgression(NewInt(start), NewInt(offset))
}

func NewBigLinearProgression(start, offset *Int) *LinearProgression {
	return &LinearProgression{start, offset}
}

type LinearProgression struct {
	start  *Int
	offset *Int
}

func (lp *LinearProgression) Start() *Int {
	return lp.start
}

func (lp *LinearProgression) String() string {
	return fmt.Sprintf("[%v + k*%v]", lp.start, lp.offset)
}

func (lp *LinearProgression) Merge(that *LinearProgression) *LinearProgression {
	a, b := lp, that
	// Make a lower
	if a.start.GT(b.start) {
		a, b = b, a
	}

	bStart := b.start.Minus(a.start)
	for bStart.Mod(a.offset).NEQ(Zero()) {
		bStart = bStart.Plus(b.offset)
	}

	return &LinearProgression{
		bStart.Plus(a.start),
		// TODO: LowestCommonFactor, above as well?
		a.offset.Times(b.offset),
	}
}
