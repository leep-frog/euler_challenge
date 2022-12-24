package eulerchallenge

import (
	"fmt"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func P701() *problem {
	return intInputNode(701, func(o command.Output, n int) {

		initState := &p701state{n, 0, make([]int, n, n), map[int]int{}, 0, ""}

		sc := make([][]map[string][]int, n*n, n*n)
		for i := range sc {
			for j := 0; j < n*n; j++ {
				sc[i] = append(sc[i], map[string][]int{})
			}
		}

		var sci stateCache701 = sc
		fmt.Println("START")
		v := rec701(initState, &sci)
		o.Stdoutf("%.9f\n", float64(v[1])/float64(v[0]))
	}, []*execution{
		{
			args: []string{"6"},
			want: "10.742542512147338",
		},
		{
			args: []string{"7"},
			want: "13.510998363327158",
		},
	})
}

/**************/
// map from index to max area size to string(squares + setSizes) to (multiplier, areaSum, count)
type stateCache701 [][]map[string][]int

// Have code as an input, so we compute it only once
func (sc *stateCache701) put(state *p701state, code string, count, areaSum int) {
	(*sc)[state.index][state.maxArea][code] = []int{count, areaSum}
}

func (sc *stateCache701) check(state *p701state, code string) ([]int, bool) {
	m := (*sc)[state.index][state.maxArea]
	if v, ok := m[code]; ok {
		return v, true
	}
	return nil, false
}

func (state *p701state) code() string {
	keys := maps.Keys(state.setSizes)
	slices.Sort(keys)

	var kvs []string
	for _, k := range keys {
		kvs = append(kvs, fmt.Sprintf("%d:%d", k, state.setSizes[k]))
	}
	return fmt.Sprintf("%v %s", state.squares, strings.Join(kvs, " "))
}

// Return the areaSum and number of squares
func rec701(state *p701state, sc *stateCache701) []int {
	if state.index == state.size*state.size {
		return []int{1, state.maxArea}
	}

	code := state.code()
	if v, ok := sc.check(state, code); ok {
		return v
	}

	v := rec701(state.next(false), sc)
	u := rec701(state.next(true), sc)

	r := []int{
		v[0] + u[0],
		v[1] + u[1],
	}
	sc.put(state, code, r[0], r[1])
	return r
}

/*************/

// TODO: Change to p701state
type p701state struct {
	size     int
	index    int
	squares  []int
	setSizes map[int]int
	maxArea  int
	strRep   string
}

func (state *p701state) copy() *p701state {
	return &p701state{state.size, state.index, maths.CopySlice(state.squares), maths.CopyMap(state.setSizes), state.maxArea, ""}
}

func (state *p701state) readjust() {
	// Now simplify the ordering
	numberMap := map[int]int{}
	for i, s := range state.squares {
		if _, ok := numberMap[s]; !ok && s != 0 {
			numberMap[s] = len(numberMap) + 1
		}
		state.squares[i] = numberMap[s]
	}

	// Now update setSizes
	newSetSizes := map[int]int{}
	for k, v := range state.setSizes {
		idx, ok := numberMap[k]
		if ok {
			newSetSizes[idx] = v
		}
		if ok && numberMap[k] == 0 {
			fmt.Println(k, state.squares, numberMap, state.setSizes)
			panic("ARGH")
		}
	}

	state.setSizes = newSetSizes
}

func (state *p701state) next(filled bool) *p701state {
	cp := state.copy()

	mod := cp.index % cp.size

	if !filled {
		cp.squares[mod] = 0
		cp.readjust()
		cp.index++
		return cp
	}

	var left, up int

	upFilled := cp.index >= cp.size && cp.squares[mod] != 0
	if upFilled {
		up = cp.squares[mod]
	}
	leftFilled := mod != 0 && cp.squares[mod-1] != 0
	if leftFilled {
		left = cp.squares[mod-1]
	}

	var newSquare int
	if leftFilled && upFilled {
		newSquare = up
		if left != up {
			// combine left and up
			for i, s := range cp.squares {
				if s == left {
					cp.squares[i] = up
				}
			}
			cp.setSizes[up] += cp.setSizes[left]
			delete(cp.setSizes, left)
		}
	} else if leftFilled && !upFilled {
		newSquare = left
	} else if !leftFilled && upFilled {
		newSquare = up
	} else { // !leftFilled && !upFilled
		newSquare = cp.size + 1
	}
	if newSquare == 0 {
		panic("WUT")
	}
	cp.setSizes[newSquare]++
	cp.squares[mod] = newSquare
	cp.maxArea = maths.Max(cp.maxArea, cp.setSizes[newSquare])

	cp.readjust()
	cp.index++
	return cp
}

/*
func (state *p701state) draw() string {
	r := []string{fmt.Sprintln("size:", state.size, "index:", state.index, "squares:", state.squares, "setSizes:", state.setSizes, "maxArea:", state.maxArea)}

	// Print empty rows
	for i := 0; i < (state.index/state.size)-1; i++ {
		r = append(r, strings.Repeat("X ", state.size)+"\n")
	}

	if state.index >= state.size {
		// Now print remaining xs in row
		r = append(r, strings.Repeat("X ", state.index%state.size))

		// Now print squares in second to last row
		for i := (state.index % state.size); i < state.size; i++ {
			r = append(r, fmt.Sprintf("%d ", state.squares[i]))
		}
		r = append(r, "\n")
	}

	for i := 0; i < state.index%state.size; i++ {
		r = append(r, fmt.Sprintf("%d ", state.squares[i]))
	}
	r = append(r, "\n")
	return strings.Join(r, "")
}

func (state *p701state) cmp(that *p701state) int {
	if state.maxArea != that.maxArea {
		if state.maxArea < that.maxArea {
			return -1
		}
		return 1
	}
	for i, sq := range state.squares {
		if sq != that.squares[i] {
			if sq < that.squares[i] {
				return -1
			}
			return 1
		}
	}
	// setSizes guaranteed to be the same
	if len(state.setSizes) != len(that.setSizes) {
		// fmt.Println(state.squares, that.squares)
		// fmt.Println(state.setSizes, that.setSizes)
		panic("Unexpected")
	}
	for k, v := range state.setSizes {
		if v != that.setSizes[k] {
			if v < that.setSizes[k] {
				return -1
			}
			return 1
		}
	}

	//
	return 0
}*/
