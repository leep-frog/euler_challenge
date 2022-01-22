package twentyone

/*var (
	amphipodMap = map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"D": 3,
		".": -1,
	}
	moveCost = []int{1, 10, 100, 1000}
)

type Amphipod struct {
	kind         int
	space        int
	hallwayMoved bool
}

type Space struct {
	// -1 if a hallway
	room      int
	neighbors []int
	occupied  bool
}

func D23() *problem {
	return command.SerialNodes(
		command.Description(""),
		command.ExecutorNode(func(o command.Output, d *command.Data) error {
			states := &States{{
				//positions: []int{9, 15, 8, 12, 10, 13, 11, 14}, // 12521
				//positions: []int{8, 12, 14, 15, 9, 11, 10, 13},
				positions: []int{11, 13, 9, 14, 12, 15, 8, 10},
			}}

			c := &checker{
				m: map[int]*checker{},
			}

			fmt.Println((*states)[0])
			// Don't need checked because of max three move situation
			for states.Len() > 0 {
				state := heap.Pop(states).(*State)
				if _, ok := c.check(state.positions); ok {
					continue
				}
				c.add(state.positions, state.energy)
				if state.done() {
					fmt.Println("FINAL STATE")
					fmt.Println(state)
					return nil
				}

				for _, newState := range state.possibleMoves() {
					heap.Push(states, newState)
				}
			}
			return nil
		}),
	)
}

type checker struct {
	m map[int]*checker

	energy int
}

func (c *checker) check(positions []int) (int, bool) {
	if len(positions) == 0 {
		return c.energy, true
	}
	m, ok := c.m[positions[0]]
	if !ok {
		return 0, false
	}
	return m.check(positions[1:])
}

func (c *checker) add(positions []int, energy int) {
	if len(positions) == 0 {
		c.energy = energy
		return
	}

	if _, ok := c.m[positions[0]]; !ok {
		c.m[positions[0]] = &checker{
			m: map[int]*checker{},
		}
	}
	c.m[positions[0]].add(positions[1:], energy)
}

type State struct {
	// 8-d array
	positions []int

	energy int
}

func (s *State) copy() *State {
	var newP []int
	for _, p := range s.positions {
		newP = append(newP, p)
	}
	return &State{newP, s.energy}
}

func (s *State) String() string {
	invert := map[int]int{}
	for i := 0; i < len(s.positions); i++ {
		invert[s.pos(i)] = i
	}

	return strings.Join([]string{
		fmt.Sprintf("############# %d", s.energy),
		fmt.Sprintf("#%s%s.%s.%s.%s.%s%s#", get(invert, 1), get(invert, 2), get(invert, 3), get(invert, 4), get(invert, 5), get(invert, 6), get(invert, 7)),
		fmt.Sprintf("###%s#%s#%s#%s###", get(invert, 8), get(invert, 10), get(invert, 12), get(invert, 14)),
		fmt.Sprintf("  #%s#%s#%s#%s#", get(invert, 9), get(invert, 11), get(invert, 13), get(invert, 15)),
		"  #########",
	}, "\n")
}

func get(m map[int]int, i int) string {
	if r, ok := m[i]; ok {
		if r < 2 {
			return "A"
		}
		if r < 4 {
			return "B"
		}
		if r < 6 {
			return "C"
		}
		return "D"
	}
	return "."
}

func (s *State) done() bool {
	for i := 0; i < len(s.positions); i++ {
		if !s.finalPosition(i) {
			return false
		}
	}
	return true
}

#############
#...........#
###B#C#B#D###
  #D#C#B#A#
  #D#B#A#C#
  #A#D#C#A#
  #########

#############
#12.3. 4. 5. 67#
###8 #10#12#14###
  #9 #11#13#15#
  #16#18#20#22#
  #17#19#21#23#
  #########

func (s *State) finalPosition(i int) bool {
	ep := endPoses(i)
	return s.pos(i) == ep[1] || (s.pos(partnerPos(i)) == ep[1] && s.pos(i) == ep[0])
}

func (s *State) pos(i int) int {
	return s.positions[i]
}

func partnerPos(i int) int {
	if i%2 == 0 {
		return i + 1
	}
	return i - 1
}

func endPoses(i int) []int {
	if i%2 == 1 {
		i--
	}
	return []int{i + 8, i + 9}
}

func (s *State) possibleMoves() []*State {
	occupied := map[int]bool{}
	for j := 0; j < len(s.positions); j++ {
		occupied[s.pos(j)] = true
	}

	var r []*State
	for i := 0; i < len(s.positions); i++ {
		if s.finalPosition(i) {
			continue
		}

		r = append(r, s.moveI(occupied, i)...)
	}
	return r
}

func (s *State) moveI(occupied map[int]bool, i int) []*State {
	var nexts [][]int
	for next, dist := range nMap[s.pos(i)] {
		nexts = append(nexts, []int{next, dist})
	}
	var newSs []*State
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
		newS.energy += energyCoeffs[i] * dist
		newSs = append(newSs, newS)
	}

	return newSs
}

func HallwaySpace(pos int) bool {
	return pos <= 7
}

var (
	energyCoeffs = map[int]int{
		0: 1,
		1: 1,
		2: 10,
		3: 10,
		4: 100,
		5: 100,
		6: 1000,
		7: 1000,
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
			8: 1,
		},
		10: {
			3:  2,
			4:  2,
			11: 1,
		},
		11: {
			10: 1,
		},
		12: {
			4:  2,
			5:  2,
			13: 1,
		},
		13: {
			12: 1,
		},
		14: {
			5:  2,
			6:  2,
			15: 1,
		},
		15: {
			14: 1,
		},
	}
)

type States []*State

func (ss *States) Len() int {
	return len(*ss)
}

func (ss *States) Less(i, j int) bool {
	return (*ss)[i].energy < (*ss)[j].energy
}

func (ss *States) Push(x interface{}) {
	*ss = append(*ss, x.(*State))
}

func (ss *States) Pop() interface{} {
	r := (*ss)[len(*ss)-1]
	*ss = (*ss)[:len(*ss)-1]
	return r
}

func (ss *States) Swap(i, j int) {
	tmp := (*ss)[i]
	(*ss)[i] = (*ss)[j]
	(*ss)[j] = tmp
}

// returns (pod, toSpace) pairs
/*func getMoves(spaces []*Space, pods []*Amphipod) [][]int {
	occupied := map[int]bool{}
	for _, pod := range pods {
		occupied[pod.space] = true
	}

	var moves [][]int
	for _, pod := range pods {
		for _, ns := range spaces[pod.space].neighbors {
			// If both hallway spaces
			if HallwaySpace(pod.space) == HallwaySpace(ns) {
				// Only move into rooms
			} else {

			}
		}
	}
}

func HallwaySpace(i int) bool {
	return i <= 7
}

func RoomSpace(i int) bool {
	return !HallwaySpace(i)
}

// Returns space, distance pairs
func neighbors(i int, spaces []*Space) [][]int {
	var r [][]int
	// Pairs of spaces and distance
	explore := [][]int{{i, 0}}

	for len(explore) > 0 {
		next := explore[0]
		space := spaces[next[0]]
		dist := next[1]

		for _, neighbor := range space.neighbors {
			if spaces[neighbor].occupied {
				continue
			}

			// If both hallways or both rooms
			if HallwaySpace(neighbor) == HallwaySpace(i) {
				continue
			}
		}
	}
	return r
}

func NewHallway() ([]*Space, []*Amphipod) {
	spaces := []*Space{
		// Nil space
		{},
		// Hallway spaces
		// 1
		{
			room: -1,
		},
		// 2
		{
			room: -1,
		},
		// 3
		{
			room: -1,
		},
		// 4
		{
			room: -1,
		},
		// 5
		{
			room: -1,
		},
		// 6
		{
			room: -1,
		},
		// 7
		{
			room: -1,
		},
		// Room spaces
		// 8
		{
			room:     1,
			occupied: true,
		},
		// 9
		{
			room:     1,
			occupied: true,
		},

		// 10
		{
			room:     2,
			occupied: true,
		},
		// 11
		{
			room:     2,
			occupied: true,
		},

		// 12
		{
			room:     3,
			occupied: true,
		},
		// 13
		{
			room:     3,
			occupied: true,
		},

		// 14
		{
			room:     5,
			occupied: true,
		},
		// 15
		{
			room:     5,
			occupied: true,
		},
	}

	neighbors := map[int][]int{
		1:  {2},
		2:  {3, 8},
		3:  {4, 8, 10},
		4:  {5, 10, 12},
		5:  {6, 12, 14},
		6:  {7, 14},
		8:  {9},
		10: {11},
		12: {13},
		14: {15},
	}

	for n1, ns := range neighbors {
		for _, n2 := range ns {
			spaces[n1].neighbors = append(spaces[n1].neighbors, n2)
			spaces[n2].neighbors = append(spaces[n2].neighbors, n1)
		}
	}

	amphipods := []*Amphipod{
		// Room 1
		{
			kind:  1,
			space: 8,
		},
		{
			kind:  3,
			space: 9,
		},
		// Room 2
		{
			kind:  4,
			space: 10,
		},
		{
			kind:  3,
			space: 11,
		},
		// Room 3
		{
			kind:  1,
			space: 12,
		},
		{
			kind:  4,
			space: 13,
		},
		// Room 4
		{
			kind:  2,
			space: 14,
		},
		{
			kind:  2,
			space: 15,
		},
	}

	return spaces, amphipods
}
*/

/*var (
	moves = map[int]map[int]int{
		1: {
			8:  3,
			9:  4,
			10: 5,
			11: 6,
			12: 7,
			13: 8,
			14: 9,
			15: 10,
		},
		2: {
			8:  2,
			9:  3,
			10: 4,
			11: 5,
			12: 6,
			13: 7,
			14: 8,
			15: 9,
		},
		3: {
			8:  2,
			9:  3,
			10: 2,
			11: 3,
			12: 4,
			13: 5,
			14: 6,
			15: 7,
		},
		4: {
			8:  4,
			9:  5,
			10: 2,
			11: 3,
			12: 2,
			13: 3,
			14: 4,
			15: 5,
		},
		5: {
			8:  6,
			9:  7,
			10: 4,
			11: 5,
			12: 2,
			13: 3,
			14: 2,
			15: 3,
		},
		6: {
			8:  8,
			9:  9,
			10: 6,
			11: 7,
			12: 4,
			13: 5,
			14: 2,
			15: 3,
		},
		7: {
			8:  9,
			9:  10,
			10: 7,
			11: 8,
			12: 5,
			13: 6,
			14: 3,
			15: 4,
		},
	}
)*/
