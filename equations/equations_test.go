package equations

import (
	"testing"

	"github.com/leep-frog/euler_challenge/fraction"
)

func TestSolve(t *testing.T) {
	for _, test := range []struct {
		name string
		eqs  []*Equation
	}{
		{
			name: "Simple equations",
			eqs: []*Equation{
				{
					&VariableSet{
						map[Variable]*fraction.Rational{
							"x": f(2, 1),
							"y": f(1, 1),
						},
						f(-3, 1),
					},
				},
				{
					&VariableSet{
						map[Variable]*fraction.Rational{
							"x": f(1, 1),
							"y": f(-3, 1),
						},
						f(-5, 1),
					},
				},
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {

			Solve(test.eqs)

			t.Fatalf("OOPS")
		})
	}
}
