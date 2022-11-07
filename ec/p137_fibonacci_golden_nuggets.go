package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P137() *problem {
	return intInputNode(137, func(o command.Output, n int) {
		// From recursiveness of fibonacci function:
		// f(x) = x +   x^2 + 2x^3 + 3x^4 + 5x^5 + ...
		//      = x + | x^2 + x^3  + 2x^4 + 3x^5 + ...
		//            |       x^3  + x^4  + 2x^5 + ...
		//      = x + x*f(x) + x^2*f(x)
		// f(x) = x + xf(x) + x^2f(x)
		// 0 = x^2f(x) + (f(X)x + 1)x - f(x)
		// We know f(x), so rational if quadratic determinant is rational:
		// a = f(x), b = f(x) + 1 c = - f(x)
		//
		// After finding the first several, noticed that each solution is
		// f(2*x-1)*f(2*x) for x = 1, 2, 3, ...
		// Series: 2, 15, 104, 714
		f := generator.Fibonaccis()
		o.Stdoutln(f.Nth(2*n-1) * f.Nth(2*n))
	}, []*execution{
		{
			args: []string{"15"},
			want: "1120149658760",
		},
		{
			args: []string{"10"},
			want: "74049690",
		},
	})
}
