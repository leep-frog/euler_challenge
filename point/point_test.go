package point

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestConvexHull(t *testing.T) {
	permutatedCH := &ConvexHull[int]{
		[]*Point[int]{
			New(-2, 4),
			New(5, -7),
			New(1, 2),
		},
	}

	squareCH := &ConvexHull[int]{
		[]*Point[int]{
			New(-2, 2),
			New(-2, -2),
			New(2, -2),
			New(2, 2),
		},
	}

	for _, test := range []struct {
		name   string
		points []*Point[int]
		want   *ConvexHull[int]
	}{
		{
			name: "Permutation 1",
			points: []*Point[int]{
				New(1, 2),
				New(5, -7),
				New(-2, 4),
			},
			want: permutatedCH,
		},
		{
			name: "Permutation 2",
			points: []*Point[int]{
				New(1, 2),
				New(-2, 4),
				New(5, -7),
			},
			want: permutatedCH,
		},
		{
			name: "Permutation 3",
			points: []*Point[int]{
				New(5, -7),
				New(1, 2),
				New(-2, 4),
			},
			want: permutatedCH,
		},
		{
			name: "Permutation 4",
			points: []*Point[int]{
				New(5, -7),
				New(-2, 4),
				New(1, 2),
			},
			want: permutatedCH,
		},
		{
			name: "Permutation 5",
			points: []*Point[int]{
				New(-2, 4),
				New(1, 2),
				New(5, -7),
			},
			want: permutatedCH,
		},
		{
			name: "Permutation 1",
			points: []*Point[int]{
				New(-2, 4),
				New(5, -7),
				New(1, 2),
			},
			want: permutatedCH,
		},
		{
			name: "Duplicate points",
			points: []*Point[int]{
				New(-2, 4),
				New(-2, 4),
				New(-2, 4),
				New(5, -7),
				New(1, 2),
				New(1, 2),
			},
			want: permutatedCH,
		},
		{
			name: "Square",
			points: []*Point[int]{
				New(-2, 2),
				New(-2, -2),
				New(2, 2),
				New(2, -2),
			},
			want: squareCH,
		},
		{
			name: "Square with points on lines",
			points: []*Point[int]{
				New(-2, 2),
				New(-2, -2),
				New(2, 2),
				New(2, -2),
				// Points on lines
				New(2, 0),
				New(-2, 0),
				New(0, 2),
				New(0, -2),
			},
			want: squareCH,
		},
		{
			name: "Square with points in the middle",
			points: []*Point[int]{
				New(-2, 2),
				New(-2, -2),
				New(2, 2),
				New(2, -2),
				// Points on lines
				New(2, -1),
				New(2, 0),
				New(2, 1),
				New(-2, 0),
				New(0, 2),
				New(0, -2),
				// Points in the middle
				New(0, 0),
				New(1, 1),
				New(1, -1),
				New(-1, 1),
				New(-1, -1),
			},
			want: squareCH,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			ch := ConvexHullFromPoints(test.points...)
			fmt.Println(test.name, ch.Points)
			if diff := cmp.Diff(test.want, ch); diff != "" {
				t.Errorf("ConvexHullFromPoints(%v) produced incorrect convex hull (-want, +got):\n%s", test.points, diff)
			}
		})
	}
}

func TestBetween(t *testing.T) {
	for _, test := range []struct {
		name string
		p    *Point[int]
		p2   *Point[int]
		q    *Point[int]
		want bool
	}{
		{
			name: "all the same point",
			p:    New(1, 2),
			q:    New(1, 2),
			p2:   New(1, 2),
			want: true,
		},
		{
			name: "p and p2 the same point",
			p:    New(1, 2),
			q:    New(3, 4),
			p2:   New(1, 2),
		},
		{
			name: "p and q the same point",
			p:    New(1, 2),
			q:    New(1, 2),
			p2:   New(3, 4),
			want: true,
		},
		{
			name: "p2 and q the same point",
			p:    New(1, 2),
			q:    New(3, 4),
			p2:   New(3, 4),
			want: true,
		},
		{
			name: "q is betwen p and p2",
			p:    New(1, 2),
			q:    New(2, 3),
			p2:   New(3, 4),
			want: true,
		},
		{
			name: "q is not betwen p and p2",
			p:    New(1, 2),
			q:    New(3, 4),
			p2:   New(3, 3),
		},
		{
			name: "q is the origin and is betwen p and p2",
			p:    New(-7, 5),
			q:    New(0, 0),
			p2:   New(7, -5),
			want: true,
		},
		{
			name: "q is the origin and is betwen p and p2",
			p:    New(-7, 5),
			q:    New(0, 0),
			p2:   New(7, -6),
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
