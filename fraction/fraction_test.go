package fraction

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/leep-frog/euler_challenge/generator"
)

func TestIsTriangular(t *testing.T) {
	p := generator.Primes()
	for _, test := range []struct {
		f    *Fraction
		want *Fraction
	}{
		{New(6, 4), New(3, 2)},
		{New(1488, 66), New(248, 11)},
		{New(7*5*3*3*3, 3*3*7*7*7*5*8), New(3, 7*7*8)},
	} {
		t.Run(fmt.Sprintf("(%v).Simplify", test.f), func(t *testing.T) {
			if diff := cmp.Diff(test.want, test.f.Copy().Simplify(p)); diff != "" {
				t.Errorf("(%v).Simplify() returned incorrect result (-want, +got):\n%s", test.f, diff)
			}
		})
	}
}
