package fraction

import (
	"fmt"
	"math/big"

	"github.com/leep-frog/euler_challenge/maths"
)

type Rational struct {
	// if r is nil, then undefined
	r *big.Rat
}

func UndefinedRational() *Rational {
	return &Rational{}
}

func NewRational(n, d int) *Rational {
	if d == 0 {
		return &Rational{}
	}
	return nr(n, d)
}

func NewBigRational(r *big.Rat) *Rational {
	return nrb(r)
}

func nr(n, d int) *Rational {
	if d == 0 {
		return UndefinedRational()
	}
	return nrb(big.NewRat(int64(n), int64(d)))
}

func nri(n, d *big.Int) *Rational {
	if d.Cmp(big.NewInt(0)) == 0 {
		return UndefinedRational()
	}
	return nrb(big.NewRat(1, 1).SetFrac(n, d))
}

func nrb(r *big.Rat) *Rational {
	return &Rational{r}
}

func (r *Rational) Float64() float64 {
	f, _ := r.r.Float64()
	return f
}

func (r *Rational) Plus(j *Rational) *Rational {
	if r.Undefined() || j.Undefined() {
		return UndefinedRational()
	}
	return nrb(big.NewRat(1, 1).Add(r.r, j.r))
}

func (r *Rational) Minus(j *Rational) *Rational {
	if r.Undefined() || j.Undefined() {
		return UndefinedRational()
	}
	return nrb(big.NewRat(1, 1).Sub(r.r, j.r))
}

func (r *Rational) Times(j *Rational) *Rational {
	if r.Undefined() || j.Undefined() {
		return UndefinedRational()
	}
	return nrb(big.NewRat(1, 1).Mul(r.r, j.r))
}

func (r *Rational) Numer() *maths.Int {
	if r.Undefined() {
		return maths.NewInt(1)
	}
	return maths.FromBigInt(r.r.Num())
}

func (r *Rational) Denom() *maths.Int {
	if r.Undefined() {
		return maths.NewInt(0)
	}
	return maths.FromBigInt(r.r.Denom())
}

func (r *Rational) Div(j *Rational) *Rational {
	if r.Undefined() || j.Undefined() {
		return UndefinedRational()
	}
	n := r.Numer().Times(j.Denom())
	d := r.Denom().Times(j.Numer())
	return nri(n.Int(), d.Int())
}

func (r *Rational) Cmp(j *Rational) int {
	if r.r == nil {
		if j.r == nil {
			return 0
		}
		return -1
	}
	if j.r == nil {
		return 1
	}
	return r.r.Cmp(j.r)
}

func (r *Rational) LT(j *Rational) bool {
	return r.Cmp(j) < 0
}

func (r *Rational) LTE(j *Rational) bool {
	return r.Cmp(j) <= 0
}

func (r *Rational) GT(j *Rational) bool {
	return r.Cmp(j) > 0
}

func (r *Rational) GTE(j *Rational) bool {
	return r.Cmp(j) >= 0
}

func (r *Rational) EQ(j *Rational) bool {
	return r.Cmp(j) == 0
}

func (r *Rational) NEQ(j *Rational) bool {
	return r.Cmp(j) != 0
}

func (r *Rational) Reciprocal() *Rational {
	return nri(r.r.Denom(), r.r.Num())
}

func (r *Rational) Code() string {
	return r.String()
}

func (r *Rational) String() string {
	return fmt.Sprintf("%v", r.r)
}

func (r *Rational) Negate() *Rational {
	return nrb(big.NewRat(1, 1).Neg(r.r))
}

func (r *Rational) Copy() *Rational {
	return nrb(big.NewRat(0, 1).Set(r.r))
}

func (r *Rational) Undefined() bool {
	return r.r == nil
}
