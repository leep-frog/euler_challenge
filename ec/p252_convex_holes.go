package eulerchallenge

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/point"
)

func P252() *problem {
	return intInputNode(252, func(o command.Output, n int) {
		pts := generatePoints252(n)
		point.CreatePlot(fmt.Sprintf("252-%d.png", n), 800, 800, pts)
		o.Stdoutln(n)
	}, []*execution{
		{
			args: []string{"20"},
		},
	})
}

func generatePoints252(n int) point.Points[int] {
	s := []int{290797}
	var t []int

	for i := 0; i <= 2*n; i++ {
		s = append(s, (s[i]*s[i])%50515093)
		t = append(t, (s[i]%2000)-1000)
	}

	var ps []*point.Point[int]
	for k := 1; k <= n; k++ {
		ps = append(ps, point.New(t[2*k-1], t[2*k]))
	}
	return point.Points[int](ps)
}
