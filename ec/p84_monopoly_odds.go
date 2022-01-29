package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

type spaceProb struct {
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
}

func P84() *problem {
	return intInputNode(84, func(o command.Output, n int) {
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
	})
}

func convertBoard(board [][]float64) []float64 {
	var cum []float64
	for _, vals := range board {
		cum = append(cum, maths.SumSys(vals...))
	}
	return cum
}
