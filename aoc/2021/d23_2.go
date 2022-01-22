package twentyone

import (
	"fmt"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/bfs"
)

var (
	amphipodMap = map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"D": 3,
		".": -1,
	}
	moveCost = []int{1, 10, 100, 1000}
)

func D23() *problem {
	return command.SerialNodes(
		command.Description(""),
		command.ExecutorNode(func(o command.Output, d *command.Data) error {
			state := &State{
				/*positions: []int{
					15, 17, 20, 23,
					8, 12, 13, 18,
					10, 11, 21, 22,
					9, 14, 16, 19,
				},
				/* */
				/*positions: []int{
					15, 19, 20, 21,
					13, 14, 17, 18,
					11, 12, 22, 23,
					8, 9, 10, 16,
				},
				/* */
				positions: []int{
					8, 12, 15, 20,
					13, 14, 18, 23,
					11, 17, 19, 22,
					9, 10, 16, 21,
				},
				/* */
				/*positions: []int{
					8, 9, 16, 17,
					10, 11, 18, 19,
					12, 13, 20, 21,
					14, 15, 22, 23,
				},
				/* */
			}

			_, dist := bfs.ShortestOffsetPath(state, nil)
			o.Stdoutln(dist)
			return nil
		}),
	)
}

type State struct {
	// 8-d array
	positions []int
}

func (s *State) copy() *State {
	var newP []int
	for _, p := range s.positions {
		newP = append(newP, p)
	}
	return &State{newP}
}

func (s *State) String() string {
	invert := map[int]int{}
	for i := 0; i < len(s.positions); i++ {
		invert[s.pos(i)] = i
	}

	return strings.Join([]string{
		fmt.Sprintf("#############"),
		fmt.Sprintf("#%s%s.%s.%s.%s.%s%s#", get(invert, 1), get(invert, 2), get(invert, 3), get(invert, 4), get(invert, 5), get(invert, 6), get(invert, 7)),
		fmt.Sprintf("###%s#%s#%s#%s###", get(invert, 8), get(invert, 10), get(invert, 12), get(invert, 14)),
		fmt.Sprintf("  #%s#%s#%s#%s#", get(invert, 9), get(invert, 11), get(invert, 13), get(invert, 15)),
		fmt.Sprintf("  #%s#%s#%s#%s#", get(invert, 16), get(invert, 18), get(invert, 20), get(invert, 22)),
		fmt.Sprintf("  #%s#%s#%s#%s#", get(invert, 17), get(invert, 19), get(invert, 21), get(invert, 23)),
		"  #########",
	}, "\n")
}

func get(m map[int]int, i int) string {
	if r, ok := m[i]; ok {
		if r < 4 {
			return "A"
		}
		if r < 8 {
			return "B"
		}
		if r < 12 {
			return "C"
		}
		return "D"
	}
	return "."
}

func (s *State) Done(interface{}) bool {
	for i := 0; i < len(s.positions); i++ {
		if !s.finalPosition(i) {
			return false
		}
	}
	return true
}

func (s *State) Code() string {
	var r []string
	for _, p := range s.positions {
		r = append(r, fmt.Sprintf("%d", p))
	}
	return strings.Join(r, ",")
}

func (s *State) finalPosition(i int) bool {
	ep := endPoses(i)

	if s.pos(i) == ep[3] {
		return true
	}

	inPoses := map[int]bool{}
	for idx := range partnerIdx(i) {
		inPoses[s.pos(idx)] = true
	}

	if s.pos(i) == ep[2] && inPoses[ep[3]] {
		return true
	}

	if s.pos(i) == ep[1] && inPoses[ep[2]] && inPoses[ep[3]] {
		return true
	}

	if s.pos(i) == ep[0] && inPoses[ep[1]] && inPoses[ep[2]] && inPoses[ep[3]] {
		return true
	}

	return false
}

func (s *State) pos(i int) int {
	return s.positions[i]
}

func partnerIdx(i int) map[int]bool {
	m := map[int]bool{}
	start := i - (i % 4)
	for j := start; j < start+4; j++ {
		m[j] = true
	}
	delete(m, i)
	return m
}

func endPoses(i int) []int {
	if i < 4 {
		return []int{8, 9, 16, 17}
	}
	if i < 8 {
		return []int{10, 11, 18, 19}
	}
	if i < 12 {
		return []int{12, 13, 20, 21}
	}
	return []int{14, 15, 22, 23}
}

func (s *State) AdjacentStates(interface{}) []*bfs.AdjacentState {
	occupied := map[int]bool{}
	for j := 0; j < len(s.positions); j++ {
		occupied[s.pos(j)] = true
	}

	var r []*bfs.AdjacentState
	for i := 0; i < len(s.positions); i++ {
		if s.finalPosition(i) {
			continue
		}

		r = append(r, s.moveI(occupied, i)...)
	}
	return r
}

func (s *State) moveI(occupied map[int]bool, i int) []*bfs.AdjacentState {
	var nexts [][]int
	for next, dist := range nMap[s.pos(i)] {
		nexts = append(nexts, []int{next, dist})
	}
	var newSs []*bfs.AdjacentState
	checked := map[int]bool{}
	for len(nexts) > 0 {
		next, dist := nexts[0][0], nexts[0][1]
		nexts = nexts[1:]
		if occupied[next] || checked[next] {
			continue
		}
		checked[next] = true

		for newNext, newDist := range nMap[next] {
			nexts = append(nexts, []int{newNext, dist + newDist})
		}

		// Move must be from room to hallway or vice versa
		if HallwaySpace(s.pos(i)) == HallwaySpace(next) {
			continue
		}

		// If moving from hallway, must be to final resting place.
		og := s.pos(i)
		s.positions[i] = next
		if HallwaySpace(og) && !s.finalPosition(i) {
			s.positions[i] = og
			continue
		}
		s.positions[i] = og

		newS := s.copy()
		newS.positions[i] = next
		newSs = append(newSs, &bfs.AdjacentState{
			State:  newS,
			Offset: energyCoeffs[i] * dist,
		})
	}

	return newSs
}

func HallwaySpace(pos int) bool {
	return pos <= 7
}

var (
	energyCoeffs = []int{
		1,
		1,
		1,
		1,
		10,
		10,
		10,
		10,
		100,
		100,
		100,
		100,
		1000,
		1000,
		1000,
		1000,
	}
	nMap = map[int]map[int]int{
		1: {
			2: 1,
		},
		2: {
			1: 1,
			3: 2,
			8: 2,
		},
		3: {
			2:  2,
			4:  2,
			8:  2,
			10: 2,
		},
		4: {
			3:  2,
			5:  2,
			10: 2,
			12: 2,
		},
		5: {
			4:  2,
			6:  2,
			12: 2,
			14: 2,
		},
		6: {
			5:  2,
			7:  1,
			14: 2,
		},
		7: {
			6: 1,
		},
		8: {
			2: 2,
			3: 2,
			9: 1,
		},
		9: {
			8:  1,
			16: 1,
		},
		10: {
			3:  2,
			4:  2,
			11: 1,
		},
		11: {
			10: 1,
			18: 1,
		},
		12: {
			4:  2,
			5:  2,
			13: 1,
		},
		13: {
			12: 1,
			20: 1,
		},
		14: {
			5:  2,
			6:  2,
			15: 1,
		},
		15: {
			14: 1,
			22: 1,
		},
		16: {
			9:  1,
			17: 1,
		},
		17: {
			16: 1,
		},
		18: {
			11: 1,
			19: 1,
		},
		19: {
			18: 1,
		},
		20: {
			13: 1,
			21: 1,
		},
		21: {
			20: 1,
		},
		22: {
			15: 1,
			23: 1,
		},
		23: {
			22: 1,
		},
	}
)
