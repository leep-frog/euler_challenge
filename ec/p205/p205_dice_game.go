package p205

import (
	"math/rand"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/combinatorics"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/fraction"
)

func P205() *ecmodels.Problem {
	return ecmodels.NoInputNode(205, func(o command.Output) {
		peter := &DiceRoller{4, 9}
		colin := &DiceRoller{6, 6}

		pProbs := peter.probs()
		cProbs := colin.probs()

		// Cumulate colin
		for i := 1; i < len(cProbs); i++ {
			cProbs[i] = cProbs[i] + (cProbs[i-1])
		}

		probability := 0.0
		for i := 1; i < len(cProbs); i++ {
			probability = probability + pProbs[i]*cProbs[i-1]
		}
		o.Stdoutf("%.7f\n", probability)
	}, &ecmodels.Execution{
		Want: "0.5731441",
	})
}

type DiceRoller struct {
	sides int
	rolls int
}

func (dr *DiceRoller) roll() int {
	var sum int
	for i := 0; i < dr.rolls; i++ {
		sum += 1 + (rand.Int() % dr.sides)
	}
	return sum
}

func (dr *DiceRoller) probs() []float64 {
	probs := map[int]int{}
	dr.uniqueRolls(1, 0, nil, probs)

	var totalRolls int
	for _, v := range probs {
		totalRolls += v
	}

	maxRoll := dr.sides * dr.rolls
	rollFreqs := make([]float64, maxRoll+1, maxRoll+1)
	for i := range rollFreqs {
		rollFreqs[i] = fraction.New(probs[i], totalRolls).ToFloat()
	}
	return rollFreqs
}

func (dr *DiceRoller) uniqueRolls(start, sum int, rolls []int, probs map[int]int) {
	if len(rolls) >= dr.rolls {
		probs[sum] += combinatorics.PermutationCount(rolls).ToInt()
		return
	}

	for i := start; i <= dr.sides; i++ {
		dr.uniqueRolls(i, sum+i, append(rolls, i), probs)
	}
}
