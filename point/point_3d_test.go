package point

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCrossProduct(t *testing.T) {
	for _, test := range []struct {
		name string
		p    *Point3D
		by   *Point3D
		want *Point3D
	}{
		{
			name: "simple cross product",
			p:    New3D(3, -1, 0),
			by:   New3D(-2, 3, 4),
			want: New3D(-4, -12, 7),
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
		p    *Point3D
		by   *Point3D
		rots []Rotation
		want *Point3D
	}{
		{
			name: "one x rotation",
			p:    New3D(1, 2, 3),
			rots: []Rotation{
				XRot,
			},
			want: New3D(1, 3, -2),
		},
		{
			name: "two x rotations",
			p:    New3D(1, 2, 3),
			rots: []Rotation{
				XRot,
				XRot,
			},
			want: New3D(1, -2, -3),
		},
		{
			name: "three x rotations",
			p:    New3D(1, 2, 3),
			rots: []Rotation{
				XRot,
				XRot,
				XRot,
			},
			want: New3D(1, -3, 2),
		},
		{
			name: "four x rotations",
			p:    New3D(1, 2, 3),
			rots: []Rotation{
				XRot,
				XRot,
				XRot,
				XRot,
			},
			want: New3D(1, 2, 3),
		},
		{
			name: "one y rotation",
			p:    New3D(1, 2, 3),
			rots: []Rotation{
				YRot,
			},
			want: New3D(-3, 2, 1),
		},
		{
			name: "two y rotations",
			p:    New3D(1, 2, 3),
			rots: []Rotation{
				YRot,
				YRot,
			},
			want: New3D(-1, 2, -3),
		},
		{
			name: "three y rotations",
			p:    New3D(1, 2, 3),
			rots: []Rotation{
				YRot,
				YRot,
				YRot,
			},
			want: New3D(3, 2, -1),
		},
		{
			name: "four y rotations",
			p:    New3D(1, 2, 3),
			rots: []Rotation{
				YRot,
				YRot,
				YRot,
				YRot,
			},
			want: New3D(1, 2, 3),
		},
		{
			name: "one z rotation",
			p:    New3D(1, 2, 3),
			rots: []Rotation{
				ZRot,
			},
			want: New3D(2, -1, 3),
		},
		{
			name: "two z rotations",
			p:    New3D(1, 2, 3),
			rots: []Rotation{
				ZRot,
				ZRot,
			},
			want: New3D(-1, -2, 3),
		},
		{
			name: "three z rotations",
			p:    New3D(1, 2, 3),
			rots: []Rotation{
				ZRot,
				ZRot,
				ZRot,
			},
			want: New3D(-2, 1, 3),
		},
		{
			name: "four z rotations",
			p:    New3D(1, 2, 3),
			rots: []Rotation{
				ZRot,
				ZRot,
				ZRot,
			},
			want: New3D(-2, 1, 3),
		},
		// Rotate by a point
		{
			name: "one x rotation by",
			p:    New3D(3, 12, 20),
			by:   New3D(29, 11, 17),
			rots: []Rotation{
				XRot,
			},
			want: New3D(3, 14, 16),
		},
		{
			name: "two x rotations by",
			p:    New3D(3, 12, 20),
			by:   New3D(29, 11, 17),
			rots: []Rotation{
				XRot,
				XRot,
			},
			want: New3D(3, 10, 14),
		},
		{
			name: "three x rotations by",
			p:    New3D(3, 12, 20),
			by:   New3D(29, 11, 17),
			rots: []Rotation{
				XRot,
				XRot,
				XRot,
			},
			want: New3D(3, 8, 18),
		},
		{
			name: "four x rotations by",
			p:    New3D(3, 12, 20),
			by:   New3D(29, 11, 17),
			rots: []Rotation{
				XRot,
				XRot,
				XRot,
				XRot,
			},
			want: New3D(3, 12, 20),
		},
		{
			name: "one y rotation by",
			p:    New3D(3, 12, 20),
			by:   New3D(29, 11, 17),
			rots: []Rotation{
				YRot,
			},
			want: New3D(26, 12, -9),
		},
		{
			name: "two y rotations by",
			p:    New3D(3, 12, 20),
			by:   New3D(29, 11, 17),
			rots: []Rotation{
				YRot,
				YRot,
			},
			want: New3D(55, 12, 14),
		},
		{
			name: "three y rotations by",
			p:    New3D(3, 12, 20),
			by:   New3D(29, 11, 17),
			rots: []Rotation{
				YRot,
				YRot,
				YRot,
			},
			want: New3D(32, 12, 43),
		},
		{
			name: "four y rotations by",
			p:    New3D(3, 12, 20),
			by:   New3D(29, 11, 17),
			rots: []Rotation{
				YRot,
				YRot,
				YRot,
				YRot,
			},
			want: New3D(3, 12, 20),
		},
		{
			name: "one z rotation by",
			p:    New3D(3, 12, 20),
			by:   New3D(29, 11, 17),
			rots: []Rotation{
				ZRot,
			},
			want: New3D(30, 37, 20),
		},
		{
			name: "two z rotations by",
			p:    New3D(3, 12, 20),
			by:   New3D(29, 11, 17),
			rots: []Rotation{
				ZRot,
				ZRot,
			},
			want: New3D(55, 10, 20),
		},
		{
			name: "three z rotations by",
			p:    New3D(3, 12, 20),
			by:   New3D(29, 11, 17),
			rots: []Rotation{
				ZRot,
				ZRot,
				ZRot,
			},
			want: New3D(28, -15, 20),
		},
		{
			name: "four z rotations by",
			p:    New3D(3, 12, 20),
			by:   New3D(29, 11, 17),
			rots: []Rotation{
				ZRot,
				ZRot,
				ZRot,
				ZRot,
			},
			want: New3D(3, 12, 20),
		},
		/* Useful for commenting out tests */
	} {
		t.Run(test.name, func(t *testing.T) {
			if test.by == nil {
				test.by = New3D(0, 0, 0)
			}
			if diff := cmp.Diff(test.want, test.p.Rotate(test.by.X, test.by.Y, test.by.Z, test.rots...)); diff != "" {
				t.Errorf("Point3D operation returned incorrect point (-want, +got):\n%s", diff)
			}
		})
	}
}

func TestRotFuncs(t *testing.T) {
	for _, test := range []struct {
		name string
		p    *Point3D
		by   *Point3D
		want []*Point3D
	}{
		{
			name: "Rotates in all directions",
			p:    New3D(1, 2, 3),
			by:   New3D(0, 0, 0),
			want: []*Point3D{
				// Rotate around X axis
				New3D(1, 2, 3),
				New3D(1, 3, -2),
				New3D(1, -2, -3),
				New3D(1, -3, 2),

				// Rotate around negative X axis
				New3D(-1, 2, -3),
				New3D(-1, -3, -2),
				New3D(-1, -2, 3),
				New3D(-1, 3, 2),

				// Rotate around Y axis
				New3D(2, -1, 3),
				New3D(-3, -1, 2),
				New3D(-2, -1, -3),
				New3D(3, -1, -2),

				// Rotate around negative Y axis
				New3D(-2, 1, 3),
				New3D(-3, 1, -2),
				New3D(2, 1, -3),
				New3D(3, 1, 2),

				// Rotate around Z axis
				New3D(-3, 2, 1),
				New3D(2, 3, 1),
				New3D(3, -2, 1),
				New3D(-2, -3, 1),

				// Rotate around negative Z axis
				New3D(3, 2, -1),
				New3D(2, -3, -1),
				New3D(-3, -2, -1),
				New3D(-2, 3, -1),
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			var got []*Point3D
			for _, f := range RotFuncsByPoint3D(test.by) {
				got = append(got, f(test.p.Copy()))
			}
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("RotFuncsByPoint() returned incorrect points (-want, +got):\n%s", diff)
			}
		})
	}
}
