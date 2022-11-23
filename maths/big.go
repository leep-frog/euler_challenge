package maths

import (
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/leep-frog/euler_challenge/parse"
	"golang.org/x/exp/slices"
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
	return Zero().Plus(i)
}

func (i *Bint) IsZero() bool {
	return i.EQ(zero)
}

func Zero() *Bint {
	return NewBint(0)
}

func One() *Bint {
	return NewBint(1)
}

//func (i *Bint) ToInt() int

func (i *Bint) Negative() bool {
	return i.LT(zero)
}

func (i *Bint) Negation() *Bint {
	return nb(big.NewInt(0).Neg(i.i))
}

func (i *Bint) String() string {
	return i.i.String()
}

func (i *Bint) Digits() []int {
	var r []int
	for ns, idx := i.String(), 0; idx < len(ns); idx++ {
		c := ns[idx : idx+1]
		if c != "-" {
			r = append(r, parse.Atoi(c))
		}
	}
	return r
}

func (i *Bint) DigitSum() int {
	return SumSys(i.Digits()...)
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

func (i *Bint) DivInt(j int) *Bint {
	return i.Div(NewBint(j))
}

func (i *Bint) ModInt(j int) int {
	return i.Mod(NewBint(j)).ToInt()
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

func (i *Bint) Abs() *Bint {
	if i.Negative() {
		return i.Negation()
	}
	return i.Copy()
}

func (i *Bint) MagLT(j *Bint) bool {
	return i.Abs().Cmp(j.Abs()) < 0
}

func (i *Bint) MagLTE(j *Bint) bool {
	return i.Abs().Cmp(j.Abs()) <= 0
}

func (i *Bint) MagGT(j *Bint) bool {
	return i.Abs().Cmp(j.Abs()) > 0
}

func (i *Bint) MagGTE(j *Bint) bool {
	return i.Abs().Cmp(j.Abs()) >= 0
}

func (i *Bint) MagEQ(j *Bint) bool {
	return i.Abs().Cmp(j.Abs()) == 0
}

func (i *Bint) MagNEQ(j *Bint) bool {
	return i.Abs().Cmp(j.Abs()) != 0
}

func (i *Bint) Hex() string {
	var hex []string
	for !i.IsZero() {
		q, r := i.Divide(NewBint(16))
		hex = append(hex, hexLetters[r.ToInt()])
		i = q
	}
	return strings.Join(Reverse(hex), "")
}

func (i *Bint) ToInt() int {
	return int(i.i.Int64())
}

func Sort(is []*Bint) {
	slices.SortFunc(is, func(this, that *Bint) bool { return this.LT(that) })
}

// TODO
func IntFromDigits(digits []int) *Bint {
	var r []string
	for _, d := range digits {
		r = append(r, fmt.Sprintf("%d", d))
	}
	return MustIntFromString(strings.Join(r, ""))
}

func Biggify(is []int) []*Bint {
	var r []*Bint
	for _, i := range is {
		r = append(r, NewBint(i))
	}
	return r
}

// TODO: make separate files for things (like combinatorics, sets, etc.)
func Choose(n, r int) *Bint {
	return Factorial(n).Div(Factorial(r).Times(Factorial(n - r)))
}

func Factorial(n int) *Bint {
	r := One()
	for i := 1; i <= n; i++ {
		r = r.Times(NewBint(i))
	}
	return r
}

func FactorialI(n int) int {
	r := 1
	for i := 2; i <= n; i++ {
		r *= i
	}
	return r
}

func MustIntFromString(s string) *Bint {
	i, ok := IntFromString(s)
	if !ok {
		log.Fatalf("Failed to convert string %q to Int", s)
	}
	return i
}

func IntFromString(s string) (*Bint, bool) {
	i, ok := Zero().i.SetString(s, 10)
	if !ok {
		return nil, false
	}
	return nb(i), true
}

func (i *Bint) TrimDigits(n int) *Bint {
	iStr := i.String()
	n = Min(n, len(iStr))
	return MustIntFromString(iStr[len(iStr)-n:])
}
