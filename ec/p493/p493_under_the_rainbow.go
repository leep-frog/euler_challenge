package p493

import (
	"log"
	"math/big"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/fraction"
	"github.com/leep-frog/euler_challenge/maths"
)

func P493() *ecmodels.Problem {
	return ecmodels.NoInputNode(493, func(o command.Output) {

		numPicks := 4
		numColors := 3
		numMarbles := 2

		if true {
			numPicks = 20
			numColors = 7
			numMarbles = 10
		}

		o.Stdoutln(dp(numColors, numMarbles, numPicks))
	}, &ecmodels.Execution{
		Want: "6.818741802",
	})
}

func dp(colors, marbles, picks int) string {
	counts := map[int]*fraction.Rational{}
	dpRec(1, colors, 1, marbles, picks, marbles, []int{1}, counts)
	return ratCountsToFloat(counts)
}

func dpRec(color, maxColor, ofColor, maxOfColor, numPicks, atMost int, picks []int, counts map[int]*fraction.Rational) {
	if len(picks) == numPicks {

		// Number of ways this set of balls can be arranged
		total := maths.Factorial(numPicks)

		// For a given color, if two balls are used, then those can either be balls (1, 2), (1, 3), etc.
		// So for each color, multiply by (numberOfMarblesOfColor choose numberOfMarblesChosenOfColor)
		colorToCount := map[int]int{}
		for _, p := range picks {
			colorToCount[p]++
		}
		for _, v := range colorToCount {
			total = total.Times(maths.Choose(maxOfColor, v))
		}

		// Generalize the number of color groups we can use (RGB, RBG, GYB, GYO, etc.)
		total = total.Times(maths.Factorial(maxColor).Div(maths.Factorial(maxColor - color)))

		// Colors with the same numbers will result in duplicates (e.g. [1 1 2 3] will result in duplicates for
		// [R R G B] and [R R B G]
		// So any colors with the same amount of numbers need to be removed
		countToNumber := map[int]int{}
		for _, v := range colorToCount {
			countToNumber[v]++
		}
		for _, v := range countToNumber {
			total = total.Div(maths.Factorial(v))
		}

		// Finally, increment the count
		ratTotal, ok := (new(big.Rat)).SetString(total.String())
		if !ok {
			log.Fatalf("Unnable to get string")
		}
		fratTotal := fraction.NewBigRational(ratTotal)
		if v, ok := counts[color]; ok {
			counts[color] = v.Plus(fratTotal)
		} else {
			counts[color] = fratTotal
		}

		return
	}

	// Pick the same ball
	if ofColor < maxOfColor && ofColor < atMost {
		dpRec(color, maxColor, ofColor+1, maxOfColor, numPicks, atMost, append(picks, color), counts)
	}

	// Pick a different ball if have enough of this color and have more colors to choose from
	if color < maxColor {
		dpRec(color+1, maxColor, 1, maxOfColor, numPicks, ofColor, append(picks, color+1), counts)
	}
}

/*type marble struct {
	color int
}

func (m *marble) String() string {
	return fmt.Sprintf("(%d)", m.color)
}

func brute(colors, marbles, picks int) *big.Float {
	ms := map[int]*marble{}
	for i := 0; i < colors*marbles; i++ {
		ms[i] = &marble{i % colors}
	}

	counts := map[int]int{}
	bruteRec(ms, nil, picks, counts)
	fmt.Println("BR COUNTS:", counts)
	return countsToFloat(counts)
}

func bruteRec(marbles map[int]*marble, picks []*marble, rem int, counts map[int]int) {
	if rem == 0 {
		uniqueColors := map[int]bool{}
		for _, p := range picks {
			uniqueColors[p.color] = true
		}
		counts[len(uniqueColors)]++
		return
	}

	keys := maps.Keys(marbles)
	for _, key := range keys {
		m := marbles[key]
		delete(marbles, key)
		bruteRec(marbles, append(picks, m), rem-1, counts)
		marbles[key] = m
	}
}

func countsToFloat(counts map[int]int) *big.Float {
	total := 0
	for _, v := range counts {
		total += v
	}

	ev := big.NewFloat(0.0)
	ev.SetPrec(100)
	for k, v := range counts {
		prod := big.NewFloat(0.0)
		prod.SetPrec(100)
		prod.SetFloat64(float64(k * v))
		ev.Add(ev, prod)
	}
	return ev.Quo(ev, big.NewFloat(float64(total)))
}*/

func ratCountsToFloat(counts map[int]*fraction.Rational) string {
	total := fraction.NewRational(0, 1)
	for _, v := range counts {
		total = total.Plus(v)
	}

	ev := fraction.NewRational(0, 1)
	for k, v := range counts {
		prod := v.Times(fraction.NewRational(k, 1))
		ev = ev.Plus(prod)
	}
	return ev.Div(total).Rat().FloatString(9)
}

/*func monteCarlo(colors, marbles, picks, total int) *big.Float {
	var ms []*marble
	for i := 0; i < colors*marbles; i++ {
		ms = append(ms, &marble{i % colors})
	}

	counts := map[int]int{}
	for i := 0; i < total; i++ {
		counts[simulate(ms, picks)]++
	}

	return countsToFloat(counts)
}

func simulate(marbles []*marble, picks int) int {
	bread.Shuffle(marbles)

	uniqueColors := map[int]bool{}
	for i := 0; i < picks; i++ {
		uniqueColors[marbles[i].color] = true
	}
	return len(uniqueColors)
}
*/
