package commandths

import "github.com/leep-frog/euler_challenge/maths"

type Operable interface {
	~int
}

var (
	Operations = []Operation[int]{
		&Plus[int]{},
		&Minus[int]{},
		&Times[int]{},
		&Divide[int]{},
		&Exponent[int]{},
		&Modulo[int]{},
	}

	OperationMap = func() map[string]Operation[int] {
		m := map[string]Operation[int]{}
		for _, o := range Operations {
			for _, sym := range o.Symbols() {
				m[sym] = o
			}
		}
		return m
	}()
)

// Operation
type Operation[T maths.Mathable] interface {
	Symbols() []string
	Evaluate(a, b T) T
	PemdasPriority() pemdasPriority
}

type pemdasPriority int

const (
	parenthesesPriority pemdasPriority = iota
	exponentPriority
	moduloPriority
	mulDivPriority
	plusMinPriority
)

type Plus[T maths.Mathable] struct{}

func (*Plus[T]) Symbols() []string              { return []string{"+", "p"} }
func (*Plus[T]) Evaluate(a, b T) T              { return a + b }
func (*Plus[T]) PemdasPriority() pemdasPriority { return plusMinPriority }

type Minus[T maths.Mathable] struct{}

func (*Minus[T]) Symbols() []string              { return []string{"-", "m"} }
func (*Minus[T]) Evaluate(a, b T) T              { return a - b }
func (*Minus[T]) PemdasPriority() pemdasPriority { return plusMinPriority }

type Times[T maths.Mathable] struct{}

func (*Times[T]) Symbols() []string              { return []string{"*", "t"} }
func (*Times[T]) Evaluate(a, b T) T              { return a * b }
func (*Times[T]) PemdasPriority() pemdasPriority { return mulDivPriority }

type Divide[T maths.Mathable] struct{}

func (*Divide[T]) Symbols() []string              { return []string{"/", "d"} }
func (*Divide[T]) Evaluate(a, b T) T              { return a / b }
func (*Divide[T]) PemdasPriority() pemdasPriority { return mulDivPriority }

type Exponent[T maths.Mathable] struct{}

func (*Exponent[T]) Symbols() []string              { return []string{"^", "e"} }
func (*Exponent[T]) Evaluate(a, b T) T              { return maths.Pow(a, b) }
func (*Exponent[T]) PemdasPriority() pemdasPriority { return exponentPriority }

type Modulo[T Operable] struct{}

func (*Modulo[T]) Symbols() []string              { return []string{"%", "o"} }
func (*Modulo[T]) Evaluate(a, b T) T              { return a % b }
func (*Modulo[T]) PemdasPriority() pemdasPriority { return moduloPriority }

// TODO: Choose function (Choose(a, b))
