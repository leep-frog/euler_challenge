package eulerchallenge

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func P701() *problem {
	return intInputNode(701, func(o command.Output, n int) {

		states := []*p701ctx{
			{n, 0, make([]int, n, n), map[int]int{}, 0, 1, ""},
		}

		sc := make([][]map[string][]int, n*n, n*n)
		for i := range sc {
			for j := 0; j < n*n; j++ {
				sc[i] = append(sc[i], map[string][]int{})
			}
		}

		var sci stateCache701 = sc
		fmt.Println("START")
		v := rec701(states[0], &sci)
		o.Stdoutln(float64(v[1]) / float64(v[0]))
		return
		var nextStates []*p701ctx

		for i := 0; i < n*n; i++ {
			fmt.Println("I", i, len(states))

			fmt.Println("Generating", time.Now())
			for _, state := range states {
				nextStates = append(nextStates, state.next(true), state.next(false))
			}

			fmt.Println("Sorting", len(nextStates), time.Now())
			slices.SortFunc(nextStates, func(this, that *p701ctx) bool {
				return this.cmp(that) < 0
			})

			fmt.Println("Organizing", time.Now())
			states = states[:0]
			for _, ctx := range nextStates {
				if len(states) > 0 && states[len(states)-1].cmp(ctx) == 0 {
					states[len(states)-1].count += ctx.count
				} else {
					states = append(states, ctx)
				}
			}
			nextStates = nextStates[:0]
		}

		var sum int
		for _, v := range states {
			// fmt.Println("=============")
			// fmt.Println(v.draw())
			sum += v.maxArea * v.count
		}
		o.Stdoutln(sum, maths.Pow(2, n*n), float64(sum)/math.Pow(2, float64(n*n)))
		// fmt.Println(states)

	}, []*execution{
		{
			args: []string{"1"},
			want: "",
		},
	})
}

/**************/
// map from index to max area size to string(squares + setSizes) to (multiplier, areaSum, count)
type stateCache701 [][]map[string][]int

func (sc *stateCache701) put(state *p701ctx, count, areaSum int) {
	m := (*sc)[state.index][state.maxArea]
	code := state.code2()
	m[code] = []int{count, areaSum}
}

func (sc *stateCache701) check(state *p701ctx) ([]int, bool) {
	m := (*sc)[state.index][state.maxArea]
	code := state.code2()
	if v, ok := m[code]; ok {
		return v, true
	}
	return nil, false
}

func (ctx *p701ctx) code2() string {
	if ctx.strRep == "" {
		keys := maps.Keys(ctx.setSizes)
		slices.Sort(keys)

		var kvs []string
		for _, k := range keys {
			kvs = append(kvs, fmt.Sprintf("%d:%d", k, ctx.setSizes[k]))
		}
		ctx.strRep = fmt.Sprintf("%v %s", ctx.squares, strings.Join(kvs, " "))
		// ctx.strRep = fmt.Sprintf("[Index:%d Squares:%v SetSizes:{%v} MaxArea:%d]", ctx.index, ctx.squares, strings.Join(kvs, ", "), ctx.maxArea)
	}
	return ctx.strRep
}

// Return the areaSum and number of squares
func rec701(state *p701ctx, sc *stateCache701) []int {
	if state.index == state.size*state.size {
		return []int{1, state.maxArea}
	}

	if v, ok := sc.check(state); ok {
		return v
	}

	v := rec701(state.next(false), sc)
	u := rec701(state.next(true), sc)

	r := []int{
		v[0] + u[0],
		v[1] + u[1],
	}
	sc.put(state, r[0], r[1])
	return r
}

/*************/

// TODO: Change to p701state
type p701ctx struct {
	size     int
	index    int
	squares  []int
	setSizes map[int]int
	maxArea  int
	count    int
	strRep   string
}

func (ctx *p701ctx) draw() string {
	r := []string{fmt.Sprintln("size:", ctx.size, "index:", ctx.index, "squares:", ctx.squares, "setSizes:", ctx.setSizes, "maxArea:", ctx.maxArea, "count:", ctx.count)}

	// Print empty rows
	for i := 0; i < (ctx.index/ctx.size)-1; i++ {
		r = append(r, strings.Repeat("X ", ctx.size)+"\n")
	}

	if ctx.index >= ctx.size {
		// Now print remaining xs in row
		r = append(r, strings.Repeat("X ", ctx.index%ctx.size))

		// Now print squares in second to last row
		for i := (ctx.index % ctx.size); i < ctx.size; i++ {
			r = append(r, fmt.Sprintf("%d ", ctx.squares[i]))
		}
		r = append(r, "\n")
	}

	for i := 0; i < ctx.index%ctx.size; i++ {
		r = append(r, fmt.Sprintf("%d ", ctx.squares[i]))
	}
	r = append(r, "\n")
	return strings.Join(r, "")
}

func (ctx *p701ctx) copy() *p701ctx {
	return &p701ctx{ctx.size, ctx.index, maths.CopySlice(ctx.squares), maths.CopyMap(ctx.setSizes), ctx.maxArea, ctx.count, ""}
}

func (ctx *p701ctx) cmp(that *p701ctx) int {
	if ctx.maxArea != that.maxArea {
		if ctx.maxArea < that.maxArea {
			return -1
		}
		return 1
	}
	for i, sq := range ctx.squares {
		if sq != that.squares[i] {
			if sq < that.squares[i] {
				return -1
			}
			return 1
		}
	}
	// setSizes guaranteed to be the same
	if len(ctx.setSizes) != len(that.setSizes) {
		// fmt.Println(ctx.squares, that.squares)
		// fmt.Println(ctx.setSizes, that.setSizes)
		panic("Unexpected")
	}
	for k, v := range ctx.setSizes {
		if v != that.setSizes[k] {
			if v < that.setSizes[k] {
				return -1
			}
			return 1
		}
	}

	// 10.742542512147338
	return 0
}

/*func (ctx *p701ctx) String() string {
	if ctx.strRep == "" {
		keys := maps.Keys(ctx.setSizes)
		slices.Sort(keys)

		var kvs []string
		for _, k := range keys {
			kvs = append(kvs, fmt.Sprintf("%d:%d", k, ctx.setSizes[k]))
		}
		ctx.strRep = fmt.Sprintf("[Index:%d Squares:%v SetSizes:{%v} MaxArea:%d]", ctx.index, ctx.squares, strings.Join(kvs, ", "), ctx.maxArea)
	}
	return ctx.strRep
}*/

func (ctx *p701ctx) readjust() {
	// Now simplify the ordering
	numberMap := map[int]int{}
	for i, s := range ctx.squares {
		if _, ok := numberMap[s]; !ok && s != 0 {
			numberMap[s] = len(numberMap) + 1
		}
		ctx.squares[i] = numberMap[s]
	}

	// Now update setSizes
	newSetSizes := map[int]int{}
	for k, v := range ctx.setSizes {
		idx, ok := numberMap[k]
		if ok {
			newSetSizes[idx] = v
		}
		if ok && numberMap[k] == 0 {
			fmt.Println(k, ctx.squares, numberMap, ctx.setSizes)
			panic("ARGH")
		}
	}

	ctx.setSizes = newSetSizes
}

func (ctx *p701ctx) next(filled bool) *p701ctx {
	cp := ctx.copy()

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
