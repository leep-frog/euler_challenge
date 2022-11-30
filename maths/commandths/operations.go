package commandths

import "github.com/leep-frog/euler_challenge/maths"

var (
	Operations = []Operation[int]{
		&Plus[int]{},
		&Minus[int]{},
		&Times[int]{},
		&Divide[int]{},
		&Exponent[int]{},
	}

	OperationMap = func() map[string]Operation[int] {
		m := map[string]Operation[int]{}
		for _, o := range Operations {
			m[o.Symbol()] = o
		}
		return m
	}()
)

// Operation
type Operation[T maths.Mathable] interface {
	Symbol() string
	Evaluate(a, b T) T
	PemdasPriority() pemdasPriority
}

type pemdasPriority int

const (
	parenthesesPriority pemdasPriority = iota
	exponentPriority
	mulDivPriority
	plusMinPriority
)

type Plus[T maths.Mathable] struct{}

func (*Plus[T]) Symbol() string                 { return "+" }
func (*Plus[T]) Evaluate(a, b T) T              { return a + b }
func (*Plus[T]) PemdasPriority() pemdasPriority { return plusMinPriority }

type Minus[T maths.Mathable] struct{}

func (*Minus[T]) Symbol() string                 { return "-" }
func (*Minus[T]) Evaluate(a, b T) T              { return a - b }
func (*Minus[T]) PemdasPriority() pemdasPriority { return plusMinPriority }

type Times[T maths.Mathable] struct{}

func (*Times[T]) Symbol() string                 { return "*" }
func (*Times[T]) Evaluate(a, b T) T              { return a * b }
func (*Times[T]) PemdasPriority() pemdasPriority { return mulDivPriority }

type Divide[T maths.Mathable] struct{}

func (*Divide[T]) Symbol() string                 { return "/" }
func (*Divide[T]) Evaluate(a, b T) T              { return a / b }
func (*Divide[T]) PemdasPriority() pemdasPriority { return mulDivPriority }

type Exponent[T maths.Mathable] struct{}

func (*Exponent[T]) Symbol() string                 { return "^" }
func (*Exponent[T]) Evaluate(a, b T) T              { return maths.Pow(a, b) }
func (*Exponent[T]) PemdasPriority() pemdasPriority { return exponentPriority }
