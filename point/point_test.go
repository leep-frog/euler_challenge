package point

import (
	"fmt"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/leep-frog/euler_challenge/fraction"
)

func TestQuadrant(t *testing.T) {
	for _, test := range []struct {
		name string
		p    *Point[int]
		want int
	}{
		{
			"Origin",
			New(0, 0),
			-1,
		},
		{
			"Left axis",
			New(-1, 0),
			0,
		},
		{
			"First quadrant",
			New(-1, 1),
			0,
		},
		{
			"Top axis",
			New(0, 1),
			1,
		},
		{
			"Second quadrant",
			New(1, 1),
			1,
		},
		{
			"Right axis",
			New(1, 0),
			2,
		},
		{
			"Third quadrant",
			New(1, -1),
			2,
		},
		{
			"Bottom axis",
			New(0, -1),
			3,
		},
		{
			"Fourth quadrant",
			New(-1, -1),
			3,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			if diff := cmp.Diff(test.want, test.p.Quadrant()); diff != "" {
				t.Errorf("(%v).Quadrant() returned incorrect value (-want, +got):\n%s", test.p, diff)
			}
		})
	}
}

func TestLineSegmentPointFunctions(t *testing.T) {
	for _, test := range []struct {
		name             string
		ls               *RationalLineSegment
		p                *RationalPoint
		wantInBoxExc     bool
		wantInBoxInc     bool
		wantOnSegmentExc bool
		wantOnSegmentInc bool
		wantHalfPlane    bool
		wantHasVertex    bool
	}{
		{
			name:             "Point is p",
			ls:               NewRationalLineSegment(NewRationalPointI(2, 5), NewRationalPointI(6, 11)),
			p:                NewRationalPointI(2, 5),
			wantInBoxInc:     true,
			wantOnSegmentInc: true,
			wantHasVertex:    true,
		},
		{
			name:             "Point is q",
			ls:               NewRationalLineSegment(NewRationalPointI(2, 5), NewRationalPointI(6, 11)),
			p:                NewRationalPointI(6, 11),
			wantInBoxInc:     true,
			wantOnSegmentInc: true,
			wantHasVertex:    true,
		},
		{
			name:             "Point is midpoint",
			ls:               NewRationalLineSegment(NewRationalPointI(2, 5), NewRationalPointI(6, 11)),
			p:                NewRationalPointI(4, 8),
			wantInBoxInc:     true,
			wantInBoxExc:     true,
			wantOnSegmentInc: true,
			wantOnSegmentExc: true,
		},
		{
			name:         "Point is above line, but not on boundary",
			ls:           NewRationalLineSegment(NewRationalPointI(2, 5), NewRationalPointI(6, 11)),
			p:            NewRationalPointI(4, 9),
			wantInBoxInc: true,
			wantInBoxExc: true,
		},
		{
			name:         "Point is above line and on boundary",
			ls:           NewRationalLineSegment(NewRationalPointI(2, 5), NewRationalPointI(6, 11)),
			p:            NewRationalPointI(4, 11),
			wantInBoxInc: true,
		},
		{
			name:          "Point is below line, but not on boundary",
			ls:            NewRationalLineSegment(NewRationalPointI(2, 5), NewRationalPointI(6, 11)),
			p:             NewRationalPointI(4, 7),
			wantInBoxInc:  true,
			wantInBoxExc:  true,
			wantHalfPlane: true,
		},
		{
			name:          "Point is below line and on boundary",
			ls:            NewRationalLineSegment(NewRationalPointI(2, 5), NewRationalPointI(6, 11)),
			p:             NewRationalPointI(4, 5),
			wantInBoxInc:  true,
			wantHalfPlane: true,
		},
		{
			name:          "Point is to the right of line, but not on boundary",
			ls:            NewRationalLineSegment(NewRationalPointI(2, 5), NewRationalPointI(6, 11)),
			p:             NewRationalPointI(5, 8),
			wantInBoxInc:  true,
			wantInBoxExc:  true,
			wantHalfPlane: true,
		},
		{
			name:          "Point is to the right of line and on boundary",
			ls:            NewRationalLineSegment(NewRationalPointI(2, 5), NewRationalPointI(6, 11)),
			p:             NewRationalPointI(6, 8),
			wantInBoxInc:  true,
			wantHalfPlane: true,
		},
		{
			name:         "Point is to the left of line, but not on boundary",
			ls:           NewRationalLineSegment(NewRationalPointI(2, 5), NewRationalPointI(6, 11)),
			p:            NewRationalPointI(3, 8),
			wantInBoxInc: true,
			wantInBoxExc: true,
		},
		{
			name:         "Point is to the left of line and on boundary",
			ls:           NewRationalLineSegment(NewRationalPointI(2, 5), NewRationalPointI(6, 11)),
			p:            NewRationalPointI(2, 8),
			wantInBoxInc: true,
		},
		{
			name:             "Vertical line: Point is p",
			ls:               NewRationalLineSegment(NewRationalPointI(2, 5), NewRationalPointI(2, 11)),
			p:                NewRationalPointI(2, 5),
			wantInBoxInc:     true,
			wantOnSegmentInc: true,
			wantHasVertex:    true,
		},
		{
			name:             "Vertical line: Point is q",
			ls:               NewRationalLineSegment(NewRationalPointI(2, 5), NewRationalPointI(2, 11)),
			p:                NewRationalPointI(2, 11),
			wantInBoxInc:     true,
			wantOnSegmentInc: true,
			wantHasVertex:    true,
		},
		{
			name:             "Vertical line: Point is in the middle",
			ls:               NewRationalLineSegment(NewRationalPointI(2, 5), NewRationalPointI(2, 11)),
			p:                NewRationalPointI(2, 6),
			wantInBoxInc:     true,
			wantOnSegmentInc: true,
			wantOnSegmentExc: true,
		},
		{
			name: "Vertical line: Point is not in between",
			ls:   NewRationalLineSegment(NewRationalPointI(2, 5), NewRationalPointI(2, 11)),
			p:    NewRationalPointI(2, 4),
		},
		{
			name:          "Vertical line: Point is to the right",
			ls:            NewRationalLineSegment(NewRationalPointI(2, 5), NewRationalPointI(2, 11)),
			p:             NewRationalPointI(3, 8),
			wantHalfPlane: true,
		},
		{
			name: "Vertical line: Point is to the left",
			ls:   NewRationalLineSegment(NewRationalPointI(2, 5), NewRationalPointI(2, 11)),
			p:    NewRationalPointI(1, 8),
		},
		// Vertical line on y axis
		{
			name:             "Y-Axis: Point is p",
			ls:               NewRationalLineSegment(NewRationalPointI(0, -3), NewRationalPointI(0, 7)),
			p:                NewRationalPointI(0, -3),
			wantInBoxInc:     true,
			wantOnSegmentInc: true,
			wantHasVertex:    true,
		},
		{
			name:             "Y-Axis: Point is q",
			ls:               NewRationalLineSegment(NewRationalPointI(0, -3), NewRationalPointI(0, 7)),
			p:                NewRationalPointI(0, 7),
			wantInBoxInc:     true,
			wantOnSegmentInc: true,
			wantHasVertex:    true,
		},
		{
			name:             "Y-Axis: Point is in the middle positive",
			ls:               NewRationalLineSegment(NewRationalPointI(0, -3), NewRationalPointI(0, 7)),
			p:                NewRationalPointI(0, 2),
			wantInBoxInc:     true,
			wantOnSegmentInc: true,
			wantOnSegmentExc: true,
		},
		{
			name:             "Y-Axis: Point is in the middle negative",
			ls:               NewRationalLineSegment(NewRationalPointI(0, -3), NewRationalPointI(0, 7)),
			p:                NewRationalPointI(0, -2),
			wantInBoxInc:     true,
			wantOnSegmentInc: true,
			wantOnSegmentExc: true,
		},
		{
			name:             "Y-Axis: Point is the origin",
			ls:               NewRationalLineSegment(NewRationalPointI(0, -3), NewRationalPointI(0, 7)),
			p:                NewRationalPointI(0, 0),
			wantInBoxInc:     true,
			wantOnSegmentInc: true,
			wantOnSegmentExc: true,
		},
		{
			name: "Y-Axis: Point is not in between",
			ls:   NewRationalLineSegment(NewRationalPointI(0, -3), NewRationalPointI(0, 7)),
			p:    NewRationalPointI(0, -4),
		},
		{
			name:          "Y-Axis: Point is to the right",
			ls:            NewRationalLineSegment(NewRationalPointI(0, -3), NewRationalPointI(0, 7)),
			p:             NewRationalPointI(1, 2),
			wantHalfPlane: true,
		},
		{
			name: "Y-Axis: Point is to the left",
			ls:   NewRationalLineSegment(NewRationalPointI(0, -3), NewRationalPointI(0, 7)),
			p:    NewRationalPointI(-1, 2),
		},
		// Horizontal line on x axis
		{
			name:             "X-Axis: Point is p",
			ls:               NewRationalLineSegment(NewRationalPointI(-3, 0), NewRationalPointI(7, 0)),
			p:                NewRationalPointI(-3, 0),
			wantInBoxInc:     true,
			wantOnSegmentInc: true,
			wantHasVertex:    true,
		},
		{
			name:             "X-Axis: Point is q",
			ls:               NewRationalLineSegment(NewRationalPointI(-3, 0), NewRationalPointI(7, 0)),
			p:                NewRationalPointI(7, 0),
			wantInBoxInc:     true,
			wantOnSegmentInc: true,
			wantHasVertex:    true,
		},
		{
			name:             "X-Axis: Point is in the middle positive",
			ls:               NewRationalLineSegment(NewRationalPointI(-3, 0), NewRationalPointI(7, 0)),
			p:                NewRationalPointI(2, 0),
			wantInBoxInc:     true,
			wantOnSegmentInc: true,
			wantOnSegmentExc: true,
		},
		{
			name:             "X-Axis: Point is in the middle negative",
			ls:               NewRationalLineSegment(NewRationalPointI(-3, 0), NewRationalPointI(7, 0)),
			p:                NewRationalPointI(-2, 0),
			wantInBoxInc:     true,
			wantOnSegmentInc: true,
			wantOnSegmentExc: true,
		},
		{
			name:             "X-Axis: Point is the origin",
			ls:               NewRationalLineSegment(NewRationalPointI(-3, 0), NewRationalPointI(7, 0)),
			p:                NewRationalPointI(0, 0),
			wantInBoxInc:     true,
			wantOnSegmentInc: true,
			wantOnSegmentExc: true,
		},
		{
			name: "X-Axis: Point is not in between",
			ls:   NewRationalLineSegment(NewRationalPointI(-3, 0), NewRationalPointI(7, 0)),
			p:    NewRationalPointI(8, 0),
		},
		{
			name: "X-Axis: Point is above",
			ls:   NewRationalLineSegment(NewRationalPointI(-3, 0), NewRationalPointI(7, 0)),
			p:    NewRationalPointI(1, 1),
		},
		{
			name:          "X-Axis: Point is below",
			ls:            NewRationalLineSegment(NewRationalPointI(-3, 0), NewRationalPointI(7, 0)),
			p:             NewRationalPointI(1, -1),
			wantHalfPlane: true,
		},
		{
			name: "Another horizontal line",
			ls:   NewRationalLineSegment(NewRationalPointI(3, -2), NewRationalPointI(3, 9)),
			p:    NewRationalPointI(3, 17),
		},
		{
			name:             "Another vertical line",
			ls:               NewRationalLineSegment(NewRationalPointI(1, 17), NewRationalPointI(4, 17)),
			p:                NewRationalPointI(3, 17),
			wantInBoxInc:     true,
			wantOnSegmentExc: true,
			wantOnSegmentInc: true,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			// InBox
			if diff := cmp.Diff(test.wantInBoxExc, test.ls.InBoxExclusive(test.p)); diff != "" {
				t.Errorf("(%v).InBoxExclusive(%v) returned incorrect result (-want, +got):\n%s", test.ls, test.p, diff)
			}
			if diff := cmp.Diff(test.wantInBoxInc, test.ls.InBoxInclusive(test.p)); diff != "" {
				t.Errorf("(%v).InBoxInclusive(%v) returned incorrect result (-want, +got):\n%s", test.ls, test.p, diff)
			}

			// OnSegment
			if diff := cmp.Diff(test.wantOnSegmentExc, test.ls.OnSegmentExclusive(test.p)); diff != "" {
				t.Errorf("(%v).OnSegmentExclusive(%v) returned incorrect result (-want, +got):\n%s", test.ls, test.p, diff)
			}
			if diff := cmp.Diff(test.wantOnSegmentInc, test.ls.OnSegmentInclusive(test.p)); diff != "" {
				t.Errorf("(%v).OnSegmentInclusive(%v) returned incorrect result (-want, +got):\n%s", test.ls, test.p, diff)
			}

			// HalfPlane
			if diff := cmp.Diff(test.wantHalfPlane, test.ls.HalfPlane(test.p)); diff != "" {
				t.Errorf("(%v).HalfPlane(%v) returned incorrect result (-want, +got):\n%s", test.ls, test.p, diff)
			}
			// HasVertex
			if diff := cmp.Diff(test.wantHasVertex, test.ls.HasVertex(test.p)); diff != "" {
				t.Errorf("(%v).HasVertex(%v) returned incorrect result (-want, +got):\n%s", test.ls, test.p, diff)
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	for _, test := range []struct {
		name string
		ls1  *RationalLineSegment
		ls2  *RationalLineSegment
		want *RationalPoint
	}{
		{
			"Intersect at origin",
			NewRationalLineSegment(NewRationalPointI(-1, 0), NewRationalPointI(1, 0)),
			NewRationalLineSegment(NewRationalPointI(0, -1), NewRationalPointI(0, 1)),
			NewRationalPoint(fraction.NewRational(0, 1), fraction.NewRational(0, 1)),
		},
		{
			"Asymmetric intersect at origin",
			NewRationalLineSegment(NewRationalPointI(-2, 0), NewRationalPointI(3, 0)),
			NewRationalLineSegment(NewRationalPointI(0, -7), NewRationalPointI(0, 5)),
			NewRationalPoint(fraction.NewRational(0, 1), fraction.NewRational(0, 1)),
		},
		{
			"Diagonal at origin",
			NewRationalLineSegment(NewRationalPointI(-1, -2), NewRationalPointI(1, 2)),
			NewRationalLineSegment(NewRationalPointI(-5, 7), NewRationalPointI(5, -7)),
			NewRationalPoint(fraction.NewRational(0, 1), fraction.NewRational(0, 1)),
		},
		{
			"Vertical horizontal intersection that's not the origin",
			NewRationalLineSegment(NewRationalPointI(3, -2), NewRationalPointI(3, 9)),
			NewRationalLineSegment(NewRationalPointI(1, 7), NewRationalPointI(4, 7)),
			NewRationalPoint(fraction.NewRational(3, 1), fraction.NewRational(7, 1)),
		},
		{
			"Doesn't intersect",
			NewRationalLineSegment(NewRationalPointI(3, -2), NewRationalPointI(3, 9)),
			NewRationalLineSegment(NewRationalPointI(1, 17), NewRationalPointI(4, 17)),
			nil,
		},
		{
			"One vertical",
			NewRationalLineSegment(NewRationalPointI(3, -2), NewRationalPointI(3, 9)),
			NewRationalLineSegment(NewRationalPointI(0, 0), NewRationalPointI(4, 7)),
			NewRationalPoint(fraction.NewRational(3, 1), fraction.NewRational(21, 4)),
		},
		{
			"One horizontal",
			NewRationalLineSegment(NewRationalPointI(2, -1), NewRationalPointI(3, 9)),
			NewRationalLineSegment(NewRationalPointI(1, 7), NewRationalPointI(4, 7)),
			NewRationalPoint(fraction.NewRational(14, 5), fraction.NewRational(7, 1)),
		},
		{
			"Big numbers",
			NewRationalLineSegment(NewRationalPointI(46, 53), NewRationalPointI(117, 462)),
			NewRationalLineSegment(NewRationalPointI(34, 332), NewRationalPointI(287, 117)),
			NewRationalPoint(fraction.NewRational(10290629, 118742), fraction.NewRational(34108189, 118742)),
		},
		{
			"Intersect at endpoint",
			NewRationalLineSegment(NewRationalPointI(3, -2), NewRationalPointI(3, 5)),
			NewRationalLineSegment(NewRationalPointI(1, 3), NewRationalPointI(5, 7)),
			nil,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			fmt.Println(test.name, "===============")
			p := test.ls1.Intersect(test.ls2)
			if diff := cmp.Diff(test.want, p, fraction.CmpOpts()...); diff != "" {
				t.Errorf("(%v).Intersect(%v) returned incorrect result (-want, +got):\n%s", test.ls1, test.ls2, diff)
			}

			// Reverse order
			q := test.ls2.Intersect(test.ls1)
			if diff := cmp.Diff(test.want, q, fraction.CmpOpts()...); diff != "" {
				t.Errorf("(%v).Intersect(%v) returned incorrect result (-want, +got):\n%s", test.ls2, test.ls1, diff)
			}
		})
	}
}

func TestLineSegment(t *testing.T) {
	for _, test := range []struct {
		name    string
		a       *Point[int]
		b       *Point[int]
		want    *LineSegment[int]
		wantInt *LineSegmentInt
		wantM   *fraction.Rational
		wantB   *fraction.Rational
	}{
		{
			name:    "Same point",
			a:       New(2, 3),
			b:       New(2, 3),
			want:    &LineSegment[int]{New(2, 3), New(2, 3)},
			wantInt: &LineSegmentInt{&LineSegment[int]{New(2, 3), New(2, 3)}},
			wantM:   fraction.NewRational(0, 0),
			wantB:   fraction.NewRational(0, 0),
		},
		{
			name:    "Different points",
			a:       New(2, 3),
			b:       New(4, 5),
			want:    &LineSegment[int]{New(2, 3), New(4, 5)},
			wantInt: &LineSegmentInt{&LineSegment[int]{New(2, 3), New(4, 5)}},
			wantM:   fraction.NewRational(1, 1),
			wantB:   fraction.NewRational(1, 1),
		},
		{
			name:    "Orders points",
			a:       New(4, 5),
			b:       New(2, 3),
			want:    &LineSegment[int]{New(2, 3), New(4, 5)},
			wantInt: &LineSegmentInt{&LineSegment[int]{New(2, 3), New(4, 5)}},
			wantM:   fraction.NewRational(1, 1),
			wantB:   fraction.NewRational(1, 1),
		},
		{
			name:    "Orders points by x",
			a:       New(4, 5),
			b:       New(2, 5),
			want:    &LineSegment[int]{New(2, 5), New(4, 5)},
			wantInt: &LineSegmentInt{&LineSegment[int]{New(2, 5), New(4, 5)}},
			wantM:   fraction.NewRational(0, 1),
			wantB:   fraction.NewRational(5, 1),
		},
		{
			name:    "Orders points by x",
			a:       New(2, 5),
			b:       New(4, 5),
			want:    &LineSegment[int]{New(2, 5), New(4, 5)},
			wantInt: &LineSegmentInt{&LineSegment[int]{New(2, 5), New(4, 5)}},
			wantM:   fraction.NewRational(0, 1),
			wantB:   fraction.NewRational(5, 1),
		},
		{
			name:    "Orders points by y",
			a:       New(2, 3),
			b:       New(2, 5),
			want:    &LineSegment[int]{New(2, 3), New(2, 5)},
			wantInt: &LineSegmentInt{&LineSegment[int]{New(2, 3), New(2, 5)}},
			wantM:   fraction.NewRational(1, 0),
			wantB:   fraction.NewRational(1, 0),
		},
		{
			name:    "Orders points by y",
			a:       New(2, 5),
			b:       New(2, 3),
			want:    &LineSegment[int]{New(2, 3), New(2, 5)},
			wantInt: &LineSegmentInt{&LineSegment[int]{New(2, 3), New(2, 5)}},
			wantM:   fraction.NewRational(1, 0),
			wantB:   fraction.NewRational(1, 0),
		},
		{
			name:    "Vertical Line",
			a:       New(0, 5),
			b:       New(0, -3),
			want:    &LineSegment[int]{New(0, -3), New(0, 5)},
			wantInt: &LineSegmentInt{&LineSegment[int]{New(0, -3), New(0, 5)}},
			wantM:   fraction.NewRational(1, 0),
			wantB:   fraction.NewRational(0, 0),
		},
		{
			name:    "Horizontal Line",
			a:       New(5, 0),
			b:       New(-3, 0),
			want:    &LineSegment[int]{New(-3, 0), New(5, 0)},
			wantInt: &LineSegmentInt{&LineSegment[int]{New(-3, 0), New(5, 0)}},
			wantM:   fraction.NewRational(0, 1),
			wantB:   fraction.NewRational(0, 1),
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			if diff := cmp.Diff(test.want, NewLineSegment(test.a, test.b)); diff != "" {
				t.Errorf("NewLineSegment(%v, %v) returned incorrect value (-want, +got):\n%s", test.a, test.b, diff)
			}

			if diff := cmp.Diff(test.wantInt, NewLineSegmentInt(test.a, test.b)); diff != "" {
				t.Errorf("NewLineSegment(%v, %v) returned incorrect value (-want, +got):\n%s", test.a, test.b, diff)
			}

			ls := NewRationalLineSegment(NewRationalPointI(test.a.X, test.a.Y), NewRationalPointI(test.b.X, test.b.Y))
			m, b := ls.EquationMB()
			if diff := cmp.Diff(test.wantM, m, fraction.CmpOpts()...); diff != "" {
				t.Errorf("(%v).Slope() returned incorrect M value: %v", ls, diff)
			}
			if diff := cmp.Diff(test.wantB, b, fraction.CmpOpts()...); diff != "" {
				t.Errorf("(%v).Slope() returned incorrect B value: %v", ls, diff)
			}
		})
	}
}

func TestDist(t *testing.T) {
	for _, test := range []struct {
		name string
		a    *Point[int]
		b    *Point[int]
		want float64
	}{
		{
			name: "same point",
			a:    New(4, 5),
			b:    New(4, 5),
			want: 0,
		},
		{
			name: "point and the origin",
			a:    New(3, 4),
			b:    Origin[int](),
			want: 5.0,
		},
		{
			name: "offset points (linear offset of previous test)",
			a:    New(3, 4),
			b:    New(6, 8),
			want: 5.0,
		},
		{
			name: "same",
			a:    New(11, 11),
			b:    Origin[int](),
			// If 11 is on the outside, then precision is slightly off
			want: math.Sqrt(11 * 11 * 2.0),
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			if got := test.a.Dist(test.b); got != test.want {
				t.Errorf("%v.Dist(%v) returned %.2f; want %.2f", test.a, test.b, got, test.want)
			}

			ab, ba := test.a.Dist(test.b), test.b.Dist(test.a)
			if ab != ba {
				t.Errorf("%v.Dist(%v) returned %.2f, but %v.Dist(%v) returned %.2f", test.a, test.b, ab, test.b, test.a, ba)
			}
		})
	}
}

func TestArea(t *testing.T) {
	for _, test := range []struct {
		name string
		t    *Triangle[int]
		want float64
	}{
		{
			name: "same point",
			t:    NewTriangle(New(0, 0), New(3, 0), New(0, 4)),
			want: 6,
		},
		{
			name: "offset from previous example",
			t:    NewTriangle(New(6, 7), New(9, 7), New(6, 11)),
			want: 6,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			if got := test.t.Area(); got != test.want {
				t.Errorf("%v.Area() returned %.2f; want %.2f", test.t, got, test.want)
			}
		})
	}
}

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
		name    string
		p       *Point[int]
		p2      *Point[int]
		q       *Point[int]
		wantInc bool
		wantExc bool
	}{
		{
			name:    "all the same point",
			p:       New(1, 2),
			q:       New(1, 2),
			p2:      New(1, 2),
			wantInc: true,
		},
		{
			name: "p and p2 the same point",
			p:    New(1, 2),
			q:    New(3, 4),
			p2:   New(1, 2),
		},
		{
			name:    "p and q the same point",
			p:       New(1, 2),
			q:       New(1, 2),
			p2:      New(3, 4),
			wantInc: true,
		},
		{
			name:    "p2 and q the same point",
			p:       New(1, 2),
			q:       New(3, 4),
			p2:      New(3, 4),
			wantInc: true,
		},
		{
			name:    "q is betwen p and p2",
			p:       New(1, 2),
			q:       New(2, 3),
			p2:      New(3, 4),
			wantInc: true,
			wantExc: true,
		},
		{
			name: "q is not betwen p and p2",
			p:    New(1, 2),
			q:    New(3, 4),
			p2:   New(3, 3),
		},
		{
			name:    "q is the origin and is betwen p and p2",
			p:       New(-7, 5),
			q:       New(0, 0),
			p2:      New(7, -5),
			wantInc: true,
			wantExc: true,
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
			if diff := cmp.Diff(test.wantInc, test.p.BetweenInclusive(test.q, test.p2)); diff != "" {
				t.Errorf("(%v).BetweenInclusive(%v, %v) returned incorrect result (-want, +got):\n%s", test.p, test.q, test.p2, diff)
			}
			if diff := cmp.Diff(test.wantExc, test.p.BetweenExclusive(test.q, test.p2)); diff != "" {
				t.Errorf("(%v).BetweenExclusive(%v, %v) returned incorrect result (-want, +got):\n%s", test.p, test.q, test.p2, diff)
			}
		})
	}
}
