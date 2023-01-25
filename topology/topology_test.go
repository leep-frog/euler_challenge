package topology

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type operation int

const (
	plus operation = iota
	minus
	times
	divide
)

type value struct {
	key string
	v   int
}

func (v *value) Code() string {
	return v.key
}

func (v *value) Process(g Graphical[int]) int {
	return v.v
}

type expression struct {
	key string
	a   string
	op  operation
	b   string
}

func (e *expression) Code() string {
	return e.key
}

func (e *expression) Process(g Graphical[int]) int {
	a, b := g.Get(e.a), g.Get(e.b)
	switch e.op {
	case plus:
		return a + b
	case minus:
		return a - b
	case times:
		return a * b
	case divide:
		return a / b
	}
	panic("Unknown operation")
}

func TestProcess(t *testing.T) {
	for _, test := range []struct {
		name       string
		items      []Node[int]
		key        string
		want       int
		wantValues map[string]int
	}{
		{
			name: "Returns single value",
			items: []Node[int]{
				&value{"a", 7},
			},
			key:  "a",
			want: 7,
			wantValues: map[string]int{
				"a": 7,
			},
		},
		{
			name: "Ignores unused values",
			items: []Node[int]{
				&value{"a", 7},
				&value{"b", 8},
			},
			key:  "a",
			want: 7,
			wantValues: map[string]int{
				"a": 7,
			},
		},
		{
			name: "Process nested values",
			items: []Node[int]{
				&value{"a", 7},
				&value{"b", 8},
				&expression{"c", "a", plus, "b"},
			},
			key:  "c",
			want: 15,
			wantValues: map[string]int{
				"a": 7,
				"b": 8,
				"c": 15,
			},
		},
		{
			name: "Process deeper nested values",
			items: []Node[int]{
				&value{"a", 7},
				&value{"b", 8},
				&expression{"c", "a", times, "b"},
				&expression{"d", "b", plus, "a"},
				&expression{"e", "d", minus, "c"},
				&expression{"f", "e", divide, "b"},
			},
			key:  "f",
			want: -5,
			wantValues: map[string]int{
				"a": 7,
				"b": 8,
				"c": 56,
				"d": 15,
				"e": -41,
				"f": -5,
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			g := NewGraph(test.items)
			if got := g.Get(test.key); got != test.want {
				t.Errorf("NewGraph(%v).Get(%q) returned %d; want %d", test.items, test.key, got, test.want)
			}

			if diff := cmp.Diff(test.wantValues, g.Values); diff != "" {
				t.Errorf("NewGraph(%v).Get(%q) resulted in incorrect graph (-want, +got):\n%s", test.items, test.key, diff)
			}
		})
	}
}
