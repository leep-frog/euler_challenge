package twentyone

import (
	"github.com/leep-frog/command/command"
)

func D21() *problem {
	return command.SerialNodes(
		command.IntNode("ONE_POS", "Player one's starting position"),
		command.IntNode("TWO_POS", "Player two's starting position"),
		command.ExecutorNode(func(o command.Output, d *command.Data) error {
			starts := []int{d.Int("ONE_POS") - 1, d.Int("TWO_POS") - 1}
			nextDice := 1
			counts := make([]int, len(starts))
			var numRolls int
			for turn := 0; ; turn = (turn + 1) % len(starts) {
				for i := 0; i < 3; i++ {
					starts[turn] = (starts[turn] + nextDice) % 10
					nextDice = (nextDice + 1) % 100
					numRolls++
				}
				counts[turn] += starts[turn] + 1
				if counts[turn] >= 1000 {
					o.Stdoutln(numRolls * counts[1-turn])
					return nil
				}
			}
		}),
	)
}

var (
	cache = map[int]map[int]map[int]map[int]map[bool][]int{}
)

func addCache(oneScore, twoScore, onePos, twoPos int, onesTurn bool, n, d int) {
	if cache[oneScore] == nil {
		cache[oneScore] = map[int]map[int]map[int]map[bool][]int{}
	}

	if cache[oneScore][twoScore] == nil {
		cache[oneScore][twoScore] = map[int]map[int]map[bool][]int{}
	}

	if cache[oneScore][twoScore][onePos] == nil {
		cache[oneScore][twoScore][onePos] = map[int]map[bool][]int{}
	}

	if cache[oneScore][twoScore][onePos][twoPos] == nil {
		cache[oneScore][twoScore][onePos][twoPos] = map[bool][]int{}
	}

	cache[oneScore][twoScore][onePos][twoPos][onesTurn] = []int{n, d}
}

func oneWins(oneScore, twoScore, onePos, twoPos int, onesTurn bool) (int, int) {
	win := 21
	if oneScore >= win {
		return 1, 1
	}
	if twoScore >= win {
		return 0, 1
	}

	if r, ok := cache[oneScore][twoScore][onePos][twoPos][onesTurn]; ok {
		return r[0], r[1]
	}

	var num, den int
	if onesTurn {
		for rollOne := 1; rollOne <= 3; rollOne++ {
			for rollTwo := 1; rollTwo <= 3; rollTwo++ {
				for rollThree := 1; rollThree <= 3; rollThree++ {
					newOnePos := (onePos + rollOne + rollTwo + rollThree) % 10
					newOneScore := oneScore + newOnePos + 1
					//fmt.Println(onePos, rollOne, rollTwo, rollThree, newOnePos, newOneScore)
					n, d := oneWins(newOneScore, twoScore, newOnePos, twoPos, false)
					num += n
					den += d
				}
			}
		}
		addCache(oneScore, twoScore, onePos, twoPos, onesTurn, num, den)
		return num, den
	}

	for rollOne := 1; rollOne <= 3; rollOne++ {
		for rollTwo := 1; rollTwo <= 3; rollTwo++ {
			for rollThree := 1; rollThree <= 3; rollThree++ {
				newTwoPos := (twoPos + rollOne + rollTwo + rollThree) % 10
				newTwoScore := twoScore + newTwoPos + 1
				n, d := oneWins(oneScore, newTwoScore, onePos, newTwoPos, true)
				num += n
				den += d
			}
		}
	}
	addCache(oneScore, twoScore, onePos, twoPos, onesTurn, num, den)
	return num, den

}

func D21_2() *problem {
	return command.SerialNodes(
		command.IntNode("ONE_POS", "Player one's starting position"),
		command.IntNode("TWO_POS", "Player two's starting position"),
		command.ExecutorNode(func(o command.Output, d *command.Data) error {
			num, den := oneWins(0, 0, d.Int("ONE_POS")-1, d.Int("TWO_POS")-1, true)
			if num > den-num {
				o.Stdoutln(num)
			} else {
				o.Stdoutln(den - num)
			}

			return nil
		}),
	)
}
