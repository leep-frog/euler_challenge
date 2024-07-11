package p84

import (
	"fmt"
	"sort"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/equilibrium"
)

/*type spaceProb struct {
	space func(int) int
	prob  float64
}

func fixedSpace(i int) func(int) int {
	return func(int) int {
		return i
	}
}

func funcNextRail() func(int) int {
	return func(i int) int {
		return (((i+5)/10)%4)*10 + 5
	}
}

func funcNextUtility() func(int) int {
	return func(i int) int {
		if i > 12 && i < 28 {
			return 28
		}
		return 12
	}
}*/

func funcNextRail() func(int) int {
	return func(i int) int {
		return (((i+5)/10)%4)*10 + 5
	}
}

func funcNextUtility() func(int) int {
	return func(i int) int {
		if i > 12 && i < 28 {
			return 28
		}
		return 12
	}
}

const (
	JAIL    = 10
	TO_JAIL = 30
	GO      = 0
)

var (
	COMMUNITY_CHEST = map[int]bool{
		2:  true,
		17: true,
		33: true,
	}
	CHANCE = map[int]bool{
		7:  true,
		22: true,
		36: true,
	}
)

type monopolySquare struct {
	idx     int
	doubles int
}

type monopolyBoard struct {
	spaces    map[int]*monopolySquare
	diceSides int
}

func (mb *monopolyBoard) weight(i int, m map[int]float64) float64 {
	return m[i] + m[i+doubles[1]] + m[i+doubles[2]]
}

var (
	doubles = []int{0, 1000, 2000}
)

func (ms *monopolySquare) Code(*monopolyBoard) int {
	return ms.idx + 1000*ms.doubles
}

func (ms *monopolySquare) Paths(board *monopolyBoard) []*equilibrium.WeightedPath[*monopolyBoard, *monopolySquare, int] {
	mWs := map[int]float64{}
	for d1 := 0; d1 < board.diceSides; d1++ {
		for d2 := 0; d2 < board.diceSides; d2++ {
			to := (ms.idx + d1 + d2) % 40
			if d1 == d2 {
				if ms.doubles == 2 {
					mWs[JAIL]++
				} else {
					mWs[doubles[ms.doubles+1]+to]++
				}
			} else {
				mWs[to]++
			}
		}
	}
	/*for j := 0; j < board.diceSides - 1; j++ {
		smaller := (ms.idx + 2 + j) % 40
		larger := (ms.idx + 2 * board.diceSides - j) % 40
		mWs[smaller] += float64(j + 1)
		mWs[larger] += float64(j + 1)
	}
	mWs[(ms.idx + board.diceSides + 1) % 40] += float64(board.diceSides)*/

	// Go to jail if relevant
	mWs[JAIL] += mWs[TO_JAIL]
	delete(mWs, TO_JAIL)
	mWs[JAIL] += mWs[TO_JAIL+doubles[1]]
	delete(mWs, TO_JAIL+doubles[1])
	mWs[JAIL] += mWs[TO_JAIL+doubles[2]]
	delete(mWs, TO_JAIL+doubles[2])

	for k, v := range mWs {
		numDoubles := k / 1000
		doubleOffset := numDoubles * 1000
		simpleK := k % 1000

		if COMMUNITY_CHEST[simpleK] {
			mWs[GO+doubleOffset] += v / 16.0
			mWs[JAIL] += v / 16.0
			mWs[k] = v * 14 / 16.0
		}
		if CHANCE[simpleK] {
			mWs[GO+doubleOffset] += v / 16.0
			mWs[JAIL] += v / 16.0
			mWs[11+doubleOffset] += v / 16.0
			mWs[24+doubleOffset] += v / 16.0
			mWs[39+doubleOffset] += v / 16.0
			mWs[5+doubleOffset] += v / 16.0
			mWs[39+doubleOffset] += v / 16.0
			mWs[k-3+doubleOffset] += v / 16.0
			mWs[funcNextRail()(k)+doubleOffset] += v * 2.0 / 16.0
			mWs[funcNextUtility()(k)+doubleOffset] += v * 1.0 / 16.0

			// Remainder
			mWs[k] = v * 6 / 16.0
		}
	}

	var ws []*equilibrium.WeightedPath[*monopolyBoard, *monopolySquare, int]
	for k, v := range mWs {
		ws = append(ws, &equilibrium.WeightedPath[*monopolyBoard, *monopolySquare, int]{board.spaces[k], v})
	}
	return ws
}

func P84() *ecmodels.Problem {
	return ecmodels.IntInputNode(84, func(o command.Output, n int) {
		board := &monopolyBoard{map[int]*monopolySquare{}, n}
		var spaces []*monopolySquare
		for i := 0; i < 40; i++ {
			board.spaces[i] = &monopolySquare{i, 0}
			board.spaces[i+doubles[1]] = &monopolySquare{i, 1}
			board.spaces[i+doubles[2]] = &monopolySquare{i, 2}
			spaces = append(spaces, board.spaces[i])
		}
		ws := equilibrium.Equilibrium(board, spaces, map[int]float64{0: 1})
		for i := 0; i <= 10; i++ {
			fmt.Printf("%5.2f ", 100*board.weight(i, ws))
		}
		fmt.Println()
		for i := 0; i < 9; i++ {
			fmt.Printf("%5.2f%s%5.2f\n", 100*board.weight(40-1-i, ws), strings.Repeat(" ", 55), 100*board.weight(10+1+i, ws))
		}
		for i := 0; i <= 10; i++ {
			fmt.Printf("%5.2f ", 100*board.weight(30-i, ws))
		}
		fmt.Println()

		var arrWs [][]float64
		for i := 0; i < 40; i++ {
			arrWs = append(arrWs, []float64{float64(i), board.weight(i, ws)})
		}
		sort.SliceStable(arrWs, func(i, j int) bool { return arrWs[i][1] > arrWs[j][1] })

		for i := 0; i < 3; i++ {
			fmt.Print(arrWs[i][0])
		}
		/*fn := float64(n)
		odds := 1 / (fn * fn)

		// special spaces:
		GO := 0
		JAIL := 10
		TO_JAIL := 30

		communityChest := maths.Set(2, 17, 33)
		change := maths.Set(7, 22, 36)
		ccProbs := []*spaceProb{
			{fixedSpace(GO), 1.0 / 16.0},
			{fixedSpace(JAIL), 1.0 / 16.0},
		}
		ccFixed := 14.0 / 16.0
		chProbs := []*spaceProb{
			{fixedSpace(GO), 1.0 / 16.0},
			{fixedSpace(JAIL), 1.0 / 16.0},
			{fixedSpace(11), 1.0 / 16.0},
			{fixedSpace(24), 1.0 / 16.0},
			{fixedSpace(39), 1.0 / 16.0},
			{fixedSpace(5), 1.0 / 16.0},
			{funcNextRail(), 2.0 / 16.0},
			{funcNextUtility(), 1.0 / 16.0},
			// Back three spaces
			{func(i int) int { return (i + (40 - 3)) % 40 }, 1.0 / 16.0},
		}
		chFixed := 6.0 / 16.0

		// 40 by
		board := make([][]float64, 40)
		for i := range board {
			// [landing_here_with_no_doubles, one double, two doubles]
			board[i] = make([]float64, 3)
		}

		board[0][0] = 1
		for change := true; change; {
			change = false
			var newBoard [][]float64
			for _ = range board {
				newBoard = append(newBoard, make([]float64, 3))
			}
			for i, vals := range board {
				for dieOne := 1; dieOne <= n; dieOne++ {
					for dieTwo := 1; dieTwo <= n; dieTwo++ {
						space := (i + dieOne + dieTwo) % 40
						if space == TO_JAIL {
							space = JAIL
						}

						spaceProbs := []*spaceProb{
							{fixedSpace(space), 1.0},
						}

						for _, sp := range spaceProbs {
							curSpace := sp.space(space)
							if dieOne != dieTwo {
								newBoard[curSpace][0] += odds * (vals[0] + vals[1] + vals[2])
							} else {
								newBoard[curSpace][1] += odds * vals[0]
								newBoard[curSpace][2] += odds * vals[1]
								//newBoard[space][0] += odds * vals[2]
								newBoard[JAIL][0] += odds * vals[2]
								// TODO: jail +=
							}
						}
					}
				}
			}

			// normalize
			sum := 0.0
			for _, vals := range newBoard {
				sum += vals[0] + vals[1] + vals[2]
			}
			fmt.Println(sum)
			for _, vals := range newBoard {
				vals[0] /= sum
				vals[1] /= sum
				vals[2] /= sum
			}

			// Check diff
			cumBoard := convertBoard(board)
			for i, val := range convertBoard(newBoard) {
				if maths.Abs(val-cumBoard[i]) > 0.0001 {
					change = true
					break
				}
			}
			board = newBoard
		}
		for i, val := range convertBoard(board) {
			o.Stdoutln(i, val)
		}*/
	}, []*ecmodels.Execution{
		{
			// ended up getting correct answer (101524) with 6 dice lol
			// so I guess bug + wrong dice = success!
			Args: []string{"6"},
			Want: "0",
			Skip: "doesn't actually work",
		},
		{
			Args: []string{"6"},
			Want: "0",
			Skip: "doesn't actually work",
		},
	})
}

/*func convertBoard(board [][]float64) []float64 {
	var cum []float64
	for _, vals := range board {
		cum = append(cum, bread.Sum(vals))
	}
	return cum
}
*/
