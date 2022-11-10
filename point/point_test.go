package point

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestConvexHull(t *testing.T) {
	permutatedCH := &ConvexHull2D[int]{
		[]*Point2D[int]{
			New2D(-2, 4),
			New2D(5, -7),
			New2D(1, 2),
		},
	}

	squareCH := &ConvexHull2D[int]{
		[]*Point2D[int]{
			New2D(-2, 2),
			New2D(-2, -2),
			New2D(2, -2),
			New2D(2, 2),
		},
	}

	for _, test := range []struct {
		name   string
		points []*Point2D[int]
		want   *ConvexHull2D[int]
	}{
		{
			name: "Permutation 1",
			points: []*Point2D[int]{
				New2D(1, 2),
				New2D(5, -7),
				New2D(-2, 4),
			},
			want: permutatedCH,
		},
		{
			name: "Permutation 2",
			points: []*Point2D[int]{
				New2D(1, 2),
				New2D(-2, 4),
				New2D(5, -7),
			},
			want: permutatedCH,
		},
		{
			name: "Permutation 3",
			points: []*Point2D[int]{
				New2D(5, -7),
				New2D(1, 2),
				New2D(-2, 4),
			},
			want: permutatedCH,
		},
		{
			name: "Permutation 4",
			points: []*Point2D[int]{
				New2D(5, -7),
				New2D(-2, 4),
				New2D(1, 2),
			},
			want: permutatedCH,
		},
		{
			name: "Permutation 5",
			points: []*Point2D[int]{
				New2D(-2, 4),
				New2D(1, 2),
				New2D(5, -7),
			},
			want: permutatedCH,
		},
		{
			name: "Permutation 1",
			points: []*Point2D[int]{
				New2D(-2, 4),
				New2D(5, -7),
				New2D(1, 2),
			},
			want: permutatedCH,
		},
		{
			name: "Duplicate points",
			points: []*Point2D[int]{
				New2D(-2, 4),
				New2D(-2, 4),
				New2D(-2, 4),
				New2D(5, -7),
				New2D(1, 2),
				New2D(1, 2),
			},
			want: permutatedCH,
		},
		{
			name: "Square",
			points: []*Point2D[int]{
				New2D(-2, 2),
				New2D(-2, -2),
				New2D(2, 2),
				New2D(2, -2),
			},
			want: squareCH,
		},
		{
			name: "Square with points on lines",
			points: []*Point2D[int]{
				New2D(-2, 2),
				New2D(-2, -2),
				New2D(2, 2),
				New2D(2, -2),
				// Points on lines
				New2D(2, 0),
				New2D(-2, 0),
				New2D(0, 2),
				New2D(0, -2),
			},
			want: squareCH,
		},
		{
			name: "Square with points in the middle",
			points: []*Point2D[int]{
				New2D(-2, 2),
				New2D(-2, -2),
				New2D(2, 2),
				New2D(2, -2),
				// Points on lines
				New2D(2, -1),
				New2D(2, 0),
				New2D(2, 1),
				New2D(-2, 0),
				New2D(0, 2),
				New2D(0, -2),
				// Points in the middle
				New2D(0, 0),
				New2D(1, 1),
				New2D(1, -1),
				New2D(-1, 1),
				New2D(-1, -1),
			},
			want: squareCH,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			ch := ConvexHull2DFromPoints(test.points...)
			fmt.Println(test.name, ch.Points)
			if diff := cmp.Diff(test.want, ch); diff != "" {
				t.Errorf("ConvexHull2DFromPoints(%v) produced incorrect convex hull (-want, +got):\n%s", test.points, diff)
			}
		})
	}
}

func TestBetween(t *testing.T) {
	for _, test := range []struct {
		name string
		p    *Point2D[int]
		p2   *Point2D[int]
		q    *Point2D[int]
		want bool
	}{
		{
			name: "all the same point",
			p:    New2D(1, 2),
			q:    New2D(1, 2),
			p2:   New2D(1, 2),
			want: true,
		},
		{
			name: "p and p2 the same point",
			p:    New2D(1, 2),
			q:    New2D(3, 4),
			p2:   New2D(1, 2),
		},
		{
			name: "p and q the same point",
			p:    New2D(1, 2),
			q:    New2D(1, 2),
			p2:   New2D(3, 4),
			want: true,
		},
		{
			name: "p2 and q the same point",
			p:    New2D(1, 2),
			q:    New2D(3, 4),
			p2:   New2D(3, 4),
			want: true,
		},
		{
			name: "q is betwen p and p2",
			p:    New2D(1, 2),
			q:    New2D(2, 3),
			p2:   New2D(3, 4),
			want: true,
		},
		{
			name: "q is not betwen p and p2",
			p:    New2D(1, 2),
			q:    New2D(3, 4),
			p2:   New2D(3, 3),
		},
		{
			name: "q is the origin and is betwen p and p2",
			p:    New2D(-7, 5),
			q:    New2D(0, 0),
			p2:   New2D(7, -5),
			want: true,
		},
		{
			name: "q is the origin and is betwen p and p2",
			p:    New2D(-7, 5),
			q:    New2D(0, 0),
			p2:   New2D(7, -6),
		},
		/* Useful for commenting out tests. */
	} {
		t.Run(test.name, func(t *testing.T) {
			if got := test.p.Between(test.q, test.p2); got != test.want {
				t.Errorf("%v.Between(%v, %v) returned %v; want %v", test.p, test.q, test.p2, got, test.want)
			}
		})
	}
}
