package maths

import (
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/parse"
	"golang.org/x/exp/slices"
)

// TODO: Rename to Int after old one is removed
type Int struct {
	i *big.Int
}

func NewInt(n int) *Int {
	return nb(big.NewInt(int64(n)))
}

func NewInt64(n int64) *Int {
	return nb(big.NewInt(n))
}

func FromBigInt(n *big.Int) *Int {
	return nb(n)
}

func nb(i *big.Int) *Int {
	return &Int{i}
}

func (i *Int) Int() *big.Int {
	return i.i
}

func (i *Int) Copy() *Int {
	return Zero().Plus(i)
}

func (i *Int) IsZero() bool {
	return i.EQ(zero)
}

func Zero() *Int {
	return NewInt(0)
}

func One() *Int {
	return NewInt(1)
}

//func (i *Int) ToInt() int

func (i *Int) Negative() bool {
	return i.LT(zero)
}

func (i *Int) Negation() *Int {
	return nb(big.NewInt(0).Neg(i.i))
}

func (i *Int) String() string {
	return i.i.String()
}

func (i *Int) Digits() []int {
	var r []int
	for ns, idx := i.String(), 0; idx < len(ns); idx++ {
		c := ns[idx : idx+1]
		if c != "-" {
			r = append(r, parse.Atoi(c))
		}
	}
	return r
}

func (i *Int) DigitSum() int {
	return bread.Sum(i.Digits())
}

func (i *Int) Plus(j *Int) *Int {
	return nb(big.NewInt(1).Add(i.i, j.i))
}

func (i *Int) PlusInt(j int) *Int {
	return i.Plus(NewInt(j))
}

func (i *Int) MinusInt(j int) *Int {
	return i.Minus(NewInt(j))
}

func (i *Int) TimesInt(j int) *Int {
	return i.Times(NewInt(j))
}

func (i *Int) Minus(j *Int) *Int {
	return nb(big.NewInt(1).Sub(i.i, j.i))
}

func (i *Int) Times(j *Int) *Int {
	return nb(big.NewInt(1).Mul(i.i, j.i))
}

func (i *Int) DivInt(j int) *Int {
	return i.Div(NewInt(j))
}

// PowMod returns (a^b % mod)
// This can be used to execute division in modulo by providing a negative exponent
func PowMod(a, b, m int) int {
	result := new(big.Int).Exp(
		big.NewInt(int64(a)),
		big.NewInt(int64(b)),
		big.NewInt(int64(m)),
	)
	return int(result.Int64())
}

func (i *Int) ModInt(j int) int {
	return i.Mod(NewInt(j)).ToInt()
}

func (i *Int) ModIntBig(j int) *Int {
	return i.Mod(NewInt(j))
}

func (i *Int) Div(j *Int) *Int {
	return nb(big.NewInt(1).Div(i.i, j.i))
}

func (i *Int) Mod(j *Int) *Int {
	_, m := big.NewInt(1).DivMod(i.i, j.i, big.NewInt(1))
	return nb(m)
}

func (i *Int) Divide(j *Int) (*Int, *Int) {
	q, m := big.NewInt(1).DivMod(i.i, j.i, big.NewInt(1))
	return nb(q), nb(m)
}

func (i *Int) Cmp(j *Int) int {
	return i.i.Cmp(j.i)
}

func (i *Int) LT(j *Int) bool {
	return i.Cmp(j) < 0
}

func (i *Int) LTE(j *Int) bool {
	return i.Cmp(j) <= 0
}

func (i *Int) GT(j *Int) bool {
	return i.Cmp(j) > 0
}

func (i *Int) GTE(j *Int) bool {
	return i.Cmp(j) >= 0
}

func (i *Int) EQ(j *Int) bool {
	return i.Cmp(j) == 0
}

func (i *Int) NEQ(j *Int) bool {
	return i.Cmp(j) != 0
}

func (i *Int) Abs() *Int {
	if i.Negative() {
		return i.Negation()
	}
	return i.Copy()
}

func (i *Int) MagLT(j *Int) bool {
	return i.Abs().Cmp(j.Abs()) < 0
}

func (i *Int) MagLTE(j *Int) bool {
	return i.Abs().Cmp(j.Abs()) <= 0
}

func (i *Int) MagGT(j *Int) bool {
	return i.Abs().Cmp(j.Abs()) > 0
}

func (i *Int) MagGTE(j *Int) bool {
	return i.Abs().Cmp(j.Abs()) >= 0
}

func (i *Int) MagEQ(j *Int) bool {
	return i.Abs().Cmp(j.Abs()) == 0
}

func (i *Int) MagNEQ(j *Int) bool {
	return i.Abs().Cmp(j.Abs()) != 0
}

func (i *Int) Hex() string {
	var hex []string
	for !i.IsZero() {
		q, r := i.Divide(NewInt(16))
		hex = append(hex, hexLetters[r.ToInt()])
		i = q
	}
	return strings.Join(bread.Reverse(hex), "")
}

func FromHex(h string) *Int {
	h = strings.ToUpper(h)
	sum := Zero()
	coef := One()
	for i := len(h) - 1; i >= 0; i-- {
		v, ok := intToHex[h[i:i+1]]
		if !ok {
			log.Fatalf("Unknown hex value: %q", h[i:i+1])
		}
		sum = sum.Plus(coef.TimesInt(v))
		coef = coef.TimesInt(16)
	}
	return sum
}

func (i *Int) ToInt() int {
	return int(i.i.Int64())
}

func (i *Int) ToBinary() *Binary {
	b := &Binary{}
	for ; i.GT(Zero()); i = i.DivInt(2) {
		b.digits = append(b.digits, i.ModIntBig(2).EQ(One()))
	}
	return b
}

func Sort(is []*Int) {
	slices.SortFunc(is, func(this, that *Int) int { return this.Cmp(that) })
}

// TODO
func IntFromDigits(digits []int) *Int {
	var r []string
	for _, d := range digits {
		r = append(r, fmt.Sprintf("%d", d))
	}
	return MustIntFromString(strings.Join(r, ""))
}

func Biggify(is []int) []*Int {
	var r []*Int
	for _, i := range is {
		r = append(r, NewInt(i))
	}
	return r
}

// TODO: make separate files for things (like combinatorics, sets, etc.)
func Choose(n, r int) *Int {
	return Factorial(n).Div(Factorial(r).Times(Factorial(n - r)))
}

var (
	factorialCache = []*Int{
		One(),
		One(),
	}
)

func Factorial(n int) *Int {
	for len(factorialCache) <= n {
		v := len(factorialCache) - 1
		factorialCache = append(factorialCache, factorialCache[v].TimesInt(v+1))
	}
	return factorialCache[n].Copy()
}

func FactorialI(n int) int {
	r := 1
	for i := 2; i <= n; i++ {
		r *= i
	}
	return r
}

func MustIntFromString(s string) *Int {
	i, ok := IntFromString(s)
	if !ok {
		log.Fatalf("Failed to convert string %q to Int", s)
	}
	return i
}

func IntFromString(s string) (*Int, bool) {
	i, ok := Zero().i.SetString(s, 10)
	if !ok {
		return nil, false
	}
	return nb(i), true
}

func (i *Int) TrimDigits(n int) *Int {
	iStr := i.String()
	n = Min(n, len(iStr))
	return MustIntFromString(iStr[len(iStr)-n:])
}
