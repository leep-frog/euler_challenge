package bfs

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBinarySearch(t *testing.T) {
	for _, test := range []struct {
		start, target, want int
		wantOk              bool
	}{
		{0, 0, 0, true},
		{0, 1, 1, false},
		{0, 2, 1, false},
		{0, 3, 1, true},
		{0, 4, 2, false},
		{0, 5, 2, false},
		{0, 6, 2, true},
		{0, 7, 3, false},
		{0, 8, 3, false},
		{0, 9, 3, true},
		{0, 10, 4, false},
		{0, 11, 4, false},
		{0, 12, 4, true},
		{0, 13, 5, false},
		{0, 14, 5, false},
		{0, 15, 5, true},
		{0, 16, 6, false},
		{0, 17, 6, false},
		{0, 18, 6, true},
		{0, 19, 7, false},
	} {
		t.Run(fmt.Sprintf("BinarySearch(start=%d, target=%d)", test.start, test.target), func(t *testing.T) {
			got, gotOk := BinarySearch[int](test.start, test.target, func(i int) int { return i * 3 })

			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("returned incorrect int result (-want, +got):\n%s", diff)
			}
			if diff := cmp.Diff(test.wantOk, gotOk); diff != "" {
				t.Errorf("returned incorrect bool result (-want, +got):\n%s", diff)
			}

		})
	}
}

func TestBinarySearch_DifferentStarts(t *testing.T) {

	starts := []int{
		0, 1, 2, 3, 4, 10,
	}

	for _, start := range starts {
		for _, test := range []struct {
			target, want int
			wantOk       bool
		}{
			{30, 10, true},
			{31, 11, false},
			{32, 11, false},
			{33, 11, true},
			{34, 12, false},
			{35, 12, false},
			{36, 12, true},
			{37, 13, false},
			{38, 13, false},
			{39, 13, true},
			{40, 14, false},

			{59, 20, false},
			{60, 20, true},
			{61, 21, false},

			{89, 30, false},
			{90, 30, true},
			{91, 31, false},

			{119, 40, false},
			{120, 40, true},
			{121, 41, false},
		} {
			t.Run(fmt.Sprintf("BinarySearch(start=%d, target=%d)", start, test.target), func(t *testing.T) {
				got, gotOk := BinarySearch[int](start, test.target, func(i int) int { return i * 3 })

				if diff := cmp.Diff(test.want, got); diff != "" {
					t.Errorf("returned incorrect int result (-want, +got):\n%s", diff)
				}
				if diff := cmp.Diff(test.wantOk, gotOk); diff != "" {
					t.Errorf("returned incorrect bool result (-want, +got):\n%s", diff)
				}
			})
		}
	}
}

func TestBinarySearch_Panic(t *testing.T) {
	for _, test := range []struct {
		start, target int
		wantPanic     string
	}{
		{-1, 0, "invalid start=-1"},
		{1, 0, "invalid start=1; startValue=3; target=0"},
		{10, 29, "invalid start=10; startValue=30; target=29"},
	} {
		t.Run(fmt.Sprintf("BinarySearch(start=%d, target=%d)", test.start, test.target), func(t *testing.T) {

			var gotRecover string
			f := func() {
				defer func() {
					gotRecover = fmt.Sprintf("%v", recover())
				}()

				BinarySearch[int](test.start, test.target, func(i int) int { return i * 3 })
			}

			f()

			if diff := cmp.Diff(test.wantPanic, gotRecover); diff != "" {
				t.Errorf("resulted in incorrect panic (-want, +got):\n%s", diff)
			}
		})
	}
}
