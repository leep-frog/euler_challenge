package maths

import (
	"math/big"

	"github.com/leep-frog/euler_challenge/parse"
)

// TODO: Rename to Int after old one is removed
type Bint struct {
	i *big.Int
}

func NewBint(n int) *Bint {
	return nb(big.NewInt(int64(n)))
}

func NewBBint(n *big.Int) *Bint {
	return nb(n)
}

func nb(i *big.Int) *Bint {
	return &Bint{i}
}

func (i *Bint) Int() *big.Int {
	return i.i
}

func (i *Bint) Copy() *Bint {
	b := NewBint(0)
	return b.Plus(i)
}

func (i *Bint) Negative() bool {
	return i.LT(NewBint(0))
}

func (i *Bint) Negation() *Bint {
	return nb(big.NewInt(0).Neg(i.i))
}

func (i *Bint) String() string {
	return i.i.String()
}

func (i *Bint) DigitSum() int {
	// TODO: Add copy function
	var sum int
	for ns, idx := i.String(), 0; idx < len(ns); idx++ {
		c := ns[idx : idx+1]
		if c != "-" {
			sum += parse.Atoi(c)
		}
	}
	return sum
}

func (i *Bint) Plus(j *Bint) *Bint {
	return nb(big.NewInt(1).Add(i.i, j.i))
}

func (i *Bint) Minus(j *Bint) *Bint {
	return nb(big.NewInt(1).Sub(i.i, j.i))
}

func (i *Bint) Times(j *Bint) *Bint {
	return nb(big.NewInt(1).Mul(i.i, j.i))
}

func (i *Bint) Div(j *Bint) *Bint {
	return nb(big.NewInt(1).Div(i.i, j.i))
}

func (i *Bint) Mod(j *Bint) *Bint {
	_, m := big.NewInt(1).DivMod(i.i, j.i, big.NewInt(1))
	return nb(m)
}

func (i *Bint) Divide(j *Bint) (*Bint, *Bint) {
	q, m := big.NewInt(1).DivMod(i.i, j.i, big.NewInt(1))
	return nb(q), nb(m)
}

func (i *Bint) Cmp(j *Bint) int {
	return i.i.Cmp(j.i)
}

func (i *Bint) LT(j *Bint) bool {
	return i.Cmp(j) < 0
}

func (i *Bint) LTE(j *Bint) bool {
	return i.Cmp(j) <= 0
}

func (i *Bint) GT(j *Bint) bool {
	return i.Cmp(j) > 0
}

func (i *Bint) GTE(j *Bint) bool {
	return i.Cmp(j) >= 0
}

func (i *Bint) EQ(j *Bint) bool {
	return i.Cmp(j) == 0
}

func (i *Bint) NEQ(j *Bint) bool {
	return i.Cmp(j) != 0
}
