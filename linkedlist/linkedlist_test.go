package linkedlist

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEndAndString(t *testing.T) {
	cycle := NewCircularList(4, 5, 6, 7, 8)
	singleToCycle := NewList(1, 2, 3)
	End(singleToCycle).Next = cycle

	for _, test := range []struct {
		name    string
		list    *Node[int]
		want    int
		wantNil bool
		wantStr string
		wantLen int
	}{
		{
			name:    "nil node",
			wantNil: true,
		},
		{
			name:    "single node",
			list:    NewList(3),
			want:    3,
			wantStr: "3",
			wantLen: 1,
		},
		{
			name:    "single circular node",
			list:    NewCircularList(3),
			want:    3,
			wantStr: "3 -> (3) -> ...",
			wantLen: 1,
		},
		{
			name:    "multiple nodes",
			list:    NewList(1, 2, 3, 5, 8, 13, 21),
			want:    21,
			wantStr: "1 -> 2 -> 3 -> 5 -> 8 -> 13 -> 21",
			wantLen: 7,
		},
		{
			name:    "multiple circular nodes",
			list:    NewCircularList(1, 2, 3, 5, 8, 13, 21),
			want:    21,
			wantStr: "1 -> 2 -> 3 -> 5 -> 8 -> 13 -> 21 -> (1) -> ...",
			wantLen: 7,
		},
		{
			name:    "single to cycle",
			list:    singleToCycle,
			want:    8,
			wantStr: "1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> (4) -> ...",
			wantLen: 8,
		},
		{
			name:    "Numbered with 0",
			list:    Numbered(0),
			wantNil: true,
		},
		{
			name:    "CircularNumbered with 0",
			list:    CircularNumbered(0),
			wantNil: true,
		},
		{
			name:    "Numbered with 1",
			list:    Numbered(1),
			want:    0,
			wantStr: "0",
			wantLen: 1,
		},
		{
			name:    "CircularNumbered with 1",
			list:    CircularNumbered(1),
			want:    0,
			wantStr: "0 -> (0) -> ...",
			wantLen: 1,
		},
		{
			name:    "Numbered with multiple",
			list:    Numbered(4),
			want:    3,
			wantStr: "0 -> 1 -> 2 -> 3",
			wantLen: 4,
		},
		{
			name:    "Numbered with multiple",
			list:    CircularNumbered(4),
			want:    3,
			wantStr: "0 -> 1 -> 2 -> 3 -> (0) -> ...",
			wantLen: 4,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			got := End(test.list)
			if (got == nil) != test.wantNil {
				t.Errorf("End(%v) expected result nilness to be %v; got %v", test.list, test.wantNil, got)
			}

			if !test.wantNil {
				if got.Value != test.want {
					t.Errorf("End(%v) returned %d; want %d", test.list, got.Value, test.want)
				}
			}

			if diff := cmp.Diff(test.wantStr, test.list.CircularRepresentation()); diff != "" {
				t.Errorf("CircularRepresentation(%v) returned incorrect value (-want, +got):\n%s", test.list, diff)
			}

			gotLen := Len(test.list)
			if test.wantLen != gotLen {
				t.Errorf("Len(%v) returned %d; want %d", test.list, gotLen, test.wantLen)
			}
		})
	}
}
