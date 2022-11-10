package point

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCrossProduct(t *testing.T) {
	for _, test := range []struct {
		name string
		p    *Point
		by   *Point
		want *Point
	}{
		{
			name: "simple cross product",
			p:    NewPoint(3, -1, 0),
			by:   NewPoint(-2, 3, 4),
			want: NewPoint(-4, -12, 7),
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			if diff := cmp.Diff(test.want, test.p.Cross(test.by)); diff != "" {
				t.Errorf("(%v).Cross(%v) returned diff (-want, +got):\n%s", test.p, test.by, diff)
			}
		})
	}
}

func TestPoint(t *testing.T) {
	for _, test := range []struct {
		name string
		p    *Point
		by   *Point
		rots []Rotation
		want *Point
	}{
		{
			name: "one x rotation",
			p:    NewPoint(1, 2, 3),
			rots: []Rotation{
				XRot,
			},
			want: NewPoint(1, 3, -2),
		},
		{
			name: "two x rotations",
			p:    NewPoint(1, 2, 3),
			rots: []Rotation{
				XRot,
				XRot,
			},
			want: NewPoint(1, -2, -3),
		},
		{
			name: "three x rotations",
			p:    NewPoint(1, 2, 3),
			rots: []Rotation{
				XRot,
				XRot,
				XRot,
			},
			want: NewPoint(1, -3, 2),
		},
		{
			name: "four x rotations",
			p:    NewPoint(1, 2, 3),
			rots: []Rotation{
				XRot,
				XRot,
				XRot,
				XRot,
			},
			want: NewPoint(1, 2, 3),
		},
		{
			name: "one y rotation",
			p:    NewPoint(1, 2, 3),
			rots: []Rotation{
				YRot,
			},
			want: NewPoint(-3, 2, 1),
		},
		{
			name: "two y rotations",
			p:    NewPoint(1, 2, 3),
			rots: []Rotation{
				YRot,
				YRot,
			},
			want: NewPoint(-1, 2, -3),
		},
		{
			name: "three y rotations",
			p:    NewPoint(1, 2, 3),
			rots: []Rotation{
				YRot,
				YRot,
				YRot,
			},
			want: NewPoint(3, 2, -1),
		},
		{
			name: "four y rotations",
			p:    NewPoint(1, 2, 3),
			rots: []Rotation{
				YRot,
				YRot,
				YRot,
				YRot,
			},
			want: NewPoint(1, 2, 3),
		},
		{
			name: "one z rotation",
			p:    NewPoint(1, 2, 3),
			rots: []Rotation{
				ZRot,
			},
			want: NewPoint(2, -1, 3),
		},
		{
			name: "two z rotations",
			p:    NewPoint(1, 2, 3),
			rots: []Rotation{
				ZRot,
				ZRot,
			},
			want: NewPoint(-1, -2, 3),
		},
		{
			name: "three z rotations",
			p:    NewPoint(1, 2, 3),
			rots: []Rotation{
				ZRot,
				ZRot,
				ZRot,
			},
			want: NewPoint(-2, 1, 3),
		},
		{
			name: "four z rotations",
			p:    NewPoint(1, 2, 3),
			rots: []Rotation{
				ZRot,
				ZRot,
				ZRot,
			},
			want: NewPoint(-2, 1, 3),
		},
		// Rotate by a point
		{
			name: "one x rotation by",
			p:    NewPoint(3, 12, 20),
			by:   NewPoint(29, 11, 17),
			rots: []Rotation{
				XRot,
			},
			want: NewPoint(3, 14, 16),
		},
		{
			name: "two x rotations by",
			p:    NewPoint(3, 12, 20),
			by:   NewPoint(29, 11, 17),
			rots: []Rotation{
				XRot,
				XRot,
			},
			want: NewPoint(3, 10, 14),
		},
		{
			name: "three x rotations by",
			p:    NewPoint(3, 12, 20),
			by:   NewPoint(29, 11, 17),
			rots: []Rotation{
				XRot,
				XRot,
				XRot,
			},
			want: NewPoint(3, 8, 18),
		},
		{
			name: "four x rotations by",
			p:    NewPoint(3, 12, 20),
			by:   NewPoint(29, 11, 17),
			rots: []Rotation{
				XRot,
				XRot,
				XRot,
				XRot,
			},
			want: NewPoint(3, 12, 20),
		},
		{
			name: "one y rotation by",
			p:    NewPoint(3, 12, 20),
			by:   NewPoint(29, 11, 17),
			rots: []Rotation{
				YRot,
			},
			want: NewPoint(26, 12, -9),
		},
		{
			name: "two y rotations by",
			p:    NewPoint(3, 12, 20),
			by:   NewPoint(29, 11, 17),
			rots: []Rotation{
				YRot,
				YRot,
			},
			want: NewPoint(55, 12, 14),
		},
		{
			name: "three y rotations by",
			p:    NewPoint(3, 12, 20),
			by:   NewPoint(29, 11, 17),
			rots: []Rotation{
				YRot,
				YRot,
				YRot,
			},
			want: NewPoint(32, 12, 43),
		},
		{
			name: "four y rotations by",
			p:    NewPoint(3, 12, 20),
			by:   NewPoint(29, 11, 17),
			rots: []Rotation{
				YRot,
				YRot,
				YRot,
				YRot,
			},
			want: NewPoint(3, 12, 20),
		},
		{
			name: "one z rotation by",
			p:    NewPoint(3, 12, 20),
			by:   NewPoint(29, 11, 17),
			rots: []Rotation{
				ZRot,
			},
			want: NewPoint(30, 37, 20),
		},
		{
			name: "two z rotations by",
			p:    NewPoint(3, 12, 20),
			by:   NewPoint(29, 11, 17),
			rots: []Rotation{
				ZRot,
				ZRot,
			},
			want: NewPoint(55, 10, 20),
		},
		{
			name: "three z rotations by",
			p:    NewPoint(3, 12, 20),
			by:   NewPoint(29, 11, 17),
			rots: []Rotation{
				ZRot,
				ZRot,
				ZRot,
			},
			want: NewPoint(28, -15, 20),
		},
		{
			name: "four z rotations by",
			p:    NewPoint(3, 12, 20),
			by:   NewPoint(29, 11, 17),
			rots: []Rotation{
				ZRot,
				ZRot,
				ZRot,
				ZRot,
			},
			want: NewPoint(3, 12, 20),
		},
		/* Useful for commenting out tests */
	} {
		t.Run(test.name, func(t *testing.T) {
			if test.by == nil {
				test.by = NewPoint(0, 0, 0)
			}
			if diff := cmp.Diff(test.want, test.p.Rotate(test.by.X, test.by.Y, test.by.Z, test.rots...)); diff != "" {
				t.Errorf("Point operation returned incorrect point (-want, +got):\n%s", diff)
			}
		})
	}
}

func TestRotFuncs(t *testing.T) {
	for _, test := range []struct {
		name string
		p    *Point
		by   *Point
		want []*Point
	}{
		{
			name: "Rotates in all directions",
			p:    NewPoint(1, 2, 3),
			by:   NewPoint(0, 0, 0),
			want: []*Point{
				// Rotate around X axis
				NewPoint(1, 2, 3),
				NewPoint(1, 3, -2),
				NewPoint(1, -2, -3),
				NewPoint(1, -3, 2),

				// Rotate around negative X axis
				NewPoint(-1, 2, -3),
				NewPoint(-1, -3, -2),
				NewPoint(-1, -2, 3),
				NewPoint(-1, 3, 2),

				// Rotate around Y axis
				NewPoint(2, -1, 3),
				NewPoint(-3, -1, 2),
				NewPoint(-2, -1, -3),
				NewPoint(3, -1, -2),

				// Rotate around negative Y axis
				NewPoint(-2, 1, 3),
				NewPoint(-3, 1, -2),
				NewPoint(2, 1, -3),
				NewPoint(3, 1, 2),

				// Rotate around Z axis
				NewPoint(-3, 2, 1),
				NewPoint(2, 3, 1),
				NewPoint(3, -2, 1),
				NewPoint(-2, -3, 1),

				// Rotate around negative Z axis
				NewPoint(3, 2, -1),
				NewPoint(2, -3, -1),
				NewPoint(-3, -2, -1),
				NewPoint(-2, 3, -1),
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			var got []*Point
			for _, f := range RotFuncsByPoint(test.by) {
				got = append(got, f(test.p.Copy()))
			}
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("RotFuncsByPoint() returned incorrect points (-want, +got):\n%s", diff)
			}
		})
	}
}
