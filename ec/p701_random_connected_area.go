package eulerchallenge

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/unionfind"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type connectedArea struct {
	id   int
	size int
}

func (ca *connectedArea) String() string {
	return fmt.Sprintf("(%d,%d)", ca.id, ca.size)
}

func (ca *connectedArea) Copy() *connectedArea {
	return &connectedArea{ca.id, ca.size}
}

type connectedAreaRow struct {
	squares []*connectedArea
	maxSize int
	count   int
}

func (car *connectedAreaRow) String() string {
	return fmt.Sprintf("%d %v", car.maxSize, car.squares)
}

func (car *connectedAreaRow) combine(row []*connectedArea) *connectedAreaRow {
	// TODO: copy row
	row = maths.CopySlice(row)
	added := map[int]bool{}
	for i, v := range row {
		if v == nil {
			continue
		}
		if added[v.id] {
			row[i] = row[i-1]
		} else {
			added[v.id] = true
			row[i] = v.Copy()
		}
	}

	// Calculate all the ones that should be connected
	topDownMap := map[int]map[int]bool{}
	for i, ca := range car.squares {
		if row[i] != nil {
			if topDownMap[ca.id] == nil {
				topDownMap[ca.id] = map[int]bool{}
			}
			topDownMap[ca.id][row[i].id] = true
		}
	}

	// Create a union find of groups
	uf := unionfind.New()
	for _, group := range topDownMap {
		for g := range group {
			for g2 := range group {
				uf.Merge(g, g2)
			}
		}
	}

	// Map from ID to connectedArea
	caMap := map[int]*connectedArea{}
	for _, ca := range row {
		caMap[ca.id] = ca
	}

	// Map from ca.id, to what the id should be
	updateMap := map[int]int{}
	for _, set := range uf.Sets() {
		for a := range set {
			for b := range set {
				updateMap[a] = maths.Min(updateMap[a], -b)
			}
		}
	}
	for _, k := range maps.Keys(updateMap) {
		updateMap[k] = -updateMap[k]
	}

	// Update each connectedArea to point to the connectedArea in the same group
	for i, ca := range row {
		if ca == nil {
			continue
		}
		row[i] = caMap[updateMap[ca.id]]
	}

	checked := map[int]bool{}
	max := car.maxSize
	for i, ca := range row {
		if ca == nil || car.squares[i] == nil || checked[car.squares[i].id] {
			continue
		}
		checked[car.squares[i].id] = true

		ca.size += car.squares[i].size
		max = maths.Max(max, ca.size)
	}

	return &connectedAreaRow{
		squares: row,
		maxSize: max,
		count:   car.count,
	}

	/*map1 := map[int][]*connectedArea{}
	map2 := map[int][]*connectedArea{}

	for i, ca := range row {

		if ca == nil || car.squares[i] == nil {
			continue
		}

		map1[ca.id] = append(map1[ca.id], car.squares[i])
		map2[car.squares[i].id] = append(map2[car.squares[i].id], ca)
	}

	// First, see which new rows should be combined
	for id, connected := range map2 {

	}

	// See if rows should be combined
	idChecked := map[int]bool{}
	for _, ca := range row {
		if idChecked[ca.id] {
			continue
		}
		idChecked[ca.id] = true

		for _, connected := range map1[ca.id] {
			ca.size += connected.size
		}
	}*/

	// var newSquares []*connectedArea
	// var maxArea int
	// startID := len(car.squares)
	// combined := map[int][]int{}
	// for i, ca := range row {
	// 	if ca == nil {
	// 		newSquares = append(newSquares, nil)
	// 		continue
	// 	}

	// 	leftFilled := i > 0 && row[i-1] != nil
	// 	var left *connectedArea
	// 	if leftFilled {
	// 		left = newSquares[i-1]
	// 	}
	// 	upFilled := car.squares[i] != nil
	// 	up := car.squares[i]

	// 	// Now create new cell depending on which neighbors are filled
	// 	if leftFilled && upFilled {
	// 		if left.id == up.id {
	// 			newSquares = append(newSquares, left)
	// 		} else {

	// 		}
	// 	} else if leftFilled && !upFilled {
	// 		newSquares = append(newSquares, left)
	// 	} else if !leftFilled && upFilled {

	// 	} else { // !leftFilled && !upFilled
	// 		newSquares = append(newSquares, &connectedArea{startID, 1})
	// 		startID++
	// 	}

	// }

	// return &connectedAreaRow{newSquares, maxArea, car.count}
}

var (
	cache = map[string]int{}
)

func uniqueRows(n int) [][]bool {
	return maths.GenerateCombos(&maths.Combinatorics[bool]{
		Parts:            []bool{true, false},
		MinLength:        n,
		MaxLength:        n,
		AllowReplacement: true,
		OrderMatters:     true,
	})
}

/*func rec701(n, rem, count, maxArea int, rows map[string][]*connectedArea) int {
	if rem == 0 {
		return count * maxArea
	}

	var nextRows [][]*connectedArea

	for _, boolRow := range uniqueRows(n) {

	}

	return rec701(n, rem-1, count, maxArea, nextRow)
}*/

func P701() *problem {
	return intInputNode(701, func(o command.Output, n int) {

		// For each unique row
		/*for _, boolRow := range uniqueRows(n) {
			// Convert to row of connectedAreas
			var row []*connectedArea
			id := 1
			for i, filled := range boolRow {
				if !filled {
					row = append(row, nil)
					continue
				}

				if i > 0 && boolRow[i-1] {
					row[i-1].size++
					row = append(row, row[i-1])
				} else {
					row = append(row, &connectedArea{id, 1})
					id++
				}
			}
			fmt.Println(boolRow, row)
		}*/
		// o.Stdoutln(len(uniqueRows))

		/*things := []*p701ctx{
			{n, 3, []int{1, 2}, map[int]int{1: 1, 2: 1}, 1, 2},
		}

		for _, t := range things {
			fmt.Println("+++++++++++++++++++")
			fmt.Println(t.draw())
			fmt.Println("-------------------")
			fmt.Println(t.next(true).draw())
			fmt.Println("-------------------")
			fmt.Println(t.next(false).draw())
			fmt.Println("-------------------")
			fmt.Println(t.draw())
		}

		return*/

		/*states := map[string]*p701ctx{}
		start := &p701ctx{n, 0, make([]int, n, n), map[int]int{}, 0, 1}
		states[start.String()] = start*/
		states := []*p701ctx{
			{n, 0, make([]int, n, n), map[int]int{}, 0, 1, ""},
		}
		var nextStates []*p701ctx

		for i := 0; i < n*n; i++ {
			fmt.Println("I", i, len(states))
			// nextStates := map[string]*p701ctx{}

			fmt.Println("Generating", time.Now())
			for _, state := range states {
				nextStates = append(nextStates, state.next(true), state.next(false))
				// for _, filled := range []bool{true, false} {
				// 	empty := state.next(filled)
				// 	emptyCode := empty.String()
				// 	if nextStates[emptyCode] == nil {
				// 		nextStates[emptyCode] = empty
				// 	} else {
				// 		nextStates[emptyCode].count += empty.count
				// 	}
				// }
			}

			fmt.Println("Sorting", len(nextStates), time.Now())
			slices.SortFunc(nextStates, func(this, that *p701ctx) bool {
				//return this.String() < that.String()
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

// func (ctx *p701ctx) LT(that *p701ctx) bool {
// 	return ctx.cmp(that) <
// }

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

func (ctx *p701ctx) String() string {
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
}

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
