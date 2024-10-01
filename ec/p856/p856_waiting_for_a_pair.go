package p856

import (
	"fmt"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/fraction"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func P856() *ecmodels.Problem {
	return ecmodels.NoInputNode(856, func(o command.Output) {
		suits := 4
		values := 13

		o.Stdoutf("%.8f\n", expectedValue(suits-1, map[int]int{suits: values - 1}, suits*values-1, suits*values).Float64())
	}, &ecmodels.Execution{
		Want: "17.09661501",
	})
}

var (
	cache = map[string]*fraction.Rational{}
)

func expectedValue(prevCnt int, counts map[int]int, cardsLeft, totalCards int) *fraction.Rational {
	if cardsLeft == 0 {
		return fraction.NewRational(totalCards, 1)
	}

	// Construct the code for the cache
	keys := maps.Keys(counts)
	slices.Sort(keys)
	codeParts := []string{fmt.Sprintf("%d", prevCnt)}
	for _, k := range keys {
		codeParts = append(codeParts, fmt.Sprintf("%d:%d", k, counts[k]))
	}
	code := strings.Join(codeParts, " ")

	// Check the cache
	if v, ok := cache[code]; ok {
		return v
	}

	// Start with the EV that we draw the pair
	ev := fraction.NewRational((totalCards-cardsLeft+1)*prevCnt, cardsLeft)

	// Iterate over all possible cards to draw
	for numberOfCards, freq := range counts {
		// Probability a value with `numberOfCards` remaining was picked
		curP := fraction.NewRational(numberOfCards*freq, cardsLeft)
		newCounts := maps.Clone(counts)
		if prevCnt > 0 {
			newCounts[prevCnt]++
		}
		newCounts[numberOfCards]--
		if newCounts[numberOfCards] == 0 {
			delete(newCounts, numberOfCards)
		}

		// Add the expected value of this situation
		ev = ev.Plus(curP.Times(expectedValue(numberOfCards-1, newCounts, cardsLeft-1, totalCards)))
	}

	cache[code] = ev
	return ev
}
