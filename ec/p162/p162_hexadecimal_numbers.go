package p162

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func crt(n int) *maths.Int {
	// Number that contain neither a or f
	// A = set of all hexadecimal strings of length exaclty n with at least one A
	// NA = set of all hexadecimal strings of length n with no As
	// H = set of all hexadecimal strings
	// A = H - !A

	// NAF = set of all hexadecimal strings of length exactly n with at neither A or F
	// AF = set of all hexadecimal strings of length exactly n with at least one A and at least one F
	// AF = -(H - NAF - A - F)

	// 0AF = set of all hexadecimal strings of length exactly n with at least one 0, A, F
	// Triple venn diagram where the middle is what we Want:
	// 0AF = H - N0AF - A - F - 0 + AF + A0 + 0F

	// A = H - NA
	// First one can't be 0
	H := maths.NewInt(15).Times(maths.BigPow(16, n-1))
	NA := maths.NewInt(14).Times(maths.BigPow(15, n-1))  // also NF
	NAF := maths.NewInt(13).Times(maths.BigPow(14, n-1)) // also NF
	NO := maths.NewInt(15).Times(maths.BigPow(15, n-1))
	NOA := maths.NewInt(14).Times(maths.BigPow(14, n-1))
	NOAF := maths.NewInt(13).Times(maths.BigPow(13, n-1))

	A := H.Minus(NA)
	F := A
	O := H.Minus(NO)

	AF := H.Minus(NAF).Minus(A).Minus(A).Negation()
	OA := H.Minus(NOA).Minus(O).Minus(A).Negation()
	OF := OA

	return H.Minus(NOAF).Minus(A).Minus(F).Minus(O).Plus(AF).Plus(OA).Plus(OF)
}

func P162() *ecmodels.Problem {
	return ecmodels.IntInputNode(162, func(o command.Output, n int) {
		sum := maths.Zero()
		for i := 3; i <= n; i++ {
			r := crt(i)
			sum = sum.Plus(r)
		}
		o.Stdoutln(sum.Hex())
	}, []*ecmodels.Execution{
		{
			Args: []string{"3"},
			Want: "4",
		},
		{
			Args: []string{"4"},
			Want: "106",
		},
		{
			Args: []string{"16"},
			Want: "3D58725572C62302",
		},
	})
}
