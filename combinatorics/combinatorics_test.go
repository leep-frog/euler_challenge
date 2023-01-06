package combinatorics

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/leep-frog/euler_challenge/parse"
)

func TestPermutations(t *testing.T) {
	for _, test := range []struct {
		name    string
		parts   []string
		want    []string
		wantRot []string
		wantSet [][]int
	}{
		{
			name:    "small",
			parts:   []string{"0", "1", "2"},
			want:    []string{"012", "021", "102", "120", "201", "210"},
			wantRot: []string{"012", "120", "201"},
			wantSet: [][]int{
				{0},
				{0, 1},
				{0, 1, 2},
				{0, 2},
				{1},
				{1, 2},
				{2},
			},
		},
		{
			name:    "duplicates",
			parts:   []string{"1", "0", "1"},
			want:    []string{"011", "101", "110"},
			wantRot: []string{"101", "011", "110"},
			wantSet: [][]int{
				{1},
				{1, 1},
				{1, 1, 0},
				{1, 0},
				{0},
			},
		},
		{
			name:    "pair of identical",
			parts:   []string{"1", "1"},
			want:    []string{"11"},
			wantRot: []string{"11", "11"},
			wantSet: [][]int{
				{1},
				{1, 1},
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			gots := Permutations(test.parts)
			var got []string
			for _, g := range gots {
				got = append(got, strings.Join(g, ""))
			}
			sort.Strings(got)
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("Permutations(%v) returned incorrect value (-want, +got):\n%s", test.parts, diff)
			}

			if gotCount := PermutationCount(test.parts).ToInt(); gotCount != len(test.want) {
				t.Errorf("PermutationCount(%v) returned %d; want %d", test.parts, gotCount, len(test.want))
			}

			gotRot := Rotations(test.parts)
			if diff := cmp.Diff(test.wantRot, gotRot); diff != "" {
				t.Errorf("Rotations(%v) returned incorrect value (-want, +got):\n%s", test.parts, diff)
			}

			var iParts []int
			for _, p := range test.parts {
				iParts = append(iParts, parse.Atoi(p))
			}

			gotSet := ChooseAllSets(iParts)
			if diff := cmp.Diff(test.wantSet, gotSet); diff != "" {
				t.Errorf("Sets(%v) returned incorrect values (-want, +got):\n%s", test.parts, diff)
			}
		})
	}
}

func TestChooseSets(t *testing.T) {
	ints := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, test := range []struct {
		parts []string
		n     int
		want  [][]string
	}{
		{},
		{
			parts: ints,
		},
		{
			parts: ints,
			n:     1,
			want: [][]string{
				{"0"},
				{"1"},
				{"2"},
				{"3"},
				{"4"},
				{"5"},
				{"6"},
				{"7"},
				{"8"},
				{"9"},
			},
		},
		{
			parts: []string{"1", "2", "3", "4", "5"},
			n:     2,
			want: [][]string{
				{"1", "2"},
				{"1", "3"},
				{"1", "4"},
				{"1", "5"},
				{"2", "3"},
				{"2", "4"},
				{"2", "5"},
				{"3", "4"},
				{"3", "5"},
				{"4", "5"},
			},
		},
	} {
		t.Run(fmt.Sprintf("ChooseSetsOfLength(%v, %d)", test.parts, test.n), func(t *testing.T) {
			got := ChooseSetsOfLength(test.parts, test.n)
			sort.SliceStable(got, func(i, j int) bool { return strings.Join(got[i], "") < strings.Join(got[j], "") })
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("ChooseSets(%v, %d) produced diff (-want, +got):\n%s", test.parts, test.n, diff)
			}
		})
	}
}
