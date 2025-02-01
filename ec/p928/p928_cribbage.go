package p928

import (
	"fmt"
	"time"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

var (
	cards = []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10,
	}

	// maxHandScore = 86
	maxHandScore = 340
)

func P928() *ecmodels.Problem {
	return ecmodels.NoInputNode(928, func(o command.Output) {

		// var max int
		// for i := 1; i <= 13; i++ {
		// 	v := i
		// 	if i > 10 {
		// 		v = 10
		// 	}
		// 	max += 4 * v
		// }
		// fmt.Println(max)

		h := []int{
			0, 2, 1, 1, 1, 1, // Should be 16
			// 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 1,
		}

		// hand := []int{0}
		// for i := 1; i <= 13; i++ {
		// 	for cnt := 0; cnt <= 4; cnt++ {
		// 		hand = append(hand, cnt)

		// 		hand = hand[:i]
		// 	}
		// }

		fmt.Println(h, cribbageScore(h), handScore(h))
		fmt.Println(iterHands([]int{0}, 0, 1, 0, 0, 0, map[int]int{0: 1}, 1))
		fmt.Println()
	}, &ecmodels.Execution{
		Want: "",
	})
}

var (
	handCount  = 0
	equalCount = 0
	best       = 0
)

// 69111 without handPerms
func iterHands(hand []int, runLength, runCount, prevRunScore, pairScore, handScore int, sumsToWaysToMake map[int]int, handPerms int) int {

	if prevRunScore+pairScore+2*sumsToWaysToMake[15] > maxHandScore {
		return 0
	}

	if len(hand) == 14 {

		if handScore == 0 {
			return 0
		}

		if runLength >= 3 {
			prevRunScore += runLength * runCount
		}

		cs := pairScore + prevRunScore + 2*sumsToWaysToMake[15]

		// if pairScore+cribbageScore(hand) == handScore {
		// 	// fmt.Println(hand)
		// 	return 1
		// }

		handCount++
		if handCount%1_000_007 == 0 {
			fmt.Println(handCount, time.Now(), hand)
			fmt.Println("PAIR", pairScore, "RUN SCORE", prevRunScore, "15s", sumsToWaysToMake)
		}

		if cs == handScore {
			// fmt.Println(hand, cs)
			if equalCount%1000 == 0 {
				fmt.Println(hand, handScore, cs, handPerms, sumsToWaysToMake)
			}
			equalCount++

			if handScore > best {
				best = handScore
				fmt.Println("******************************************************* NEW BEST")
				fmt.Println(hand, handScore, cs, handPerms, sumsToWaysToMake)

			}
			return handPerms
		}
		return 0
	}

	var sum int
	for cnt := 0; cnt <= 4; cnt++ {
		v := len(hand)
		if v > 10 {
			v = 10
		}

		rl, rc := runLength+1, runCount*cnt
		prs := prevRunScore
		if cnt == 0 {
			if runLength >= 3 {
				prs += runLength * runCount
			}
			rl = 0
			rc = 1
		}

		var iters [][]int
		for subCnt := 1; subCnt <= cnt; subCnt++ {
			sumIncr := subCnt * v
			for curSum, waysToMake := range sumsToWaysToMake {
				if sumIncr+curSum <= 15 {
					iters = append(iters, []int{sumIncr + curSum, waysToMake * maths.Choose(cnt, subCnt).ToInt()})
				}
			}
		}

		for _, idxIncr := range iters {
			sumsToWaysToMake[idxIncr[0]] += idxIncr[1]
		}

		sum += iterHands(append(hand, cnt), rl, rc, prs, pairScore+cnt*(cnt-1), handScore+v*cnt, sumsToWaysToMake, handPerms*maths.Choose(4, cnt).ToInt())

		for _, idxIncr := range iters {
			sumsToWaysToMake[idxIncr[0]] -= idxIncr[1]
			if sumsToWaysToMake[idxIncr[0]] == 0 {
				delete(sumsToWaysToMake, idxIncr[0])
			}
		}
	}
	return sum
}

func cribbageScore(hand []int) int {
	var sum int

	// pairs
	// for _, cnt := range hand {
	// 	sum += cnt * (cnt - 1)
	// }

	// fmt.Println("PAIR", sum)

	// runs
	// runLength, runCount := 0, 1
	// for _, cnt := range hand {
	// 	if cnt == 0 {
	// 		if runLength >= 3 {
	// 			sum += runLength * runCount
	// 		}

	// 		runLength = 0
	// 		runCount = 1
	// 	} else {
	// 		runLength++
	// 		runCount *= cnt
	// 	}
	// }

	// if runLength >= 3 {
	// 	sum += runLength * runCount
	// }

	// fmt.Println("PAIR AND RUNS", sum)

	// fifteens
	sumCounts := make([]int, 16)
	sumCounts[0] = 1

	// for i, cnt := range hand {

	for i := len(hand) - 1; i >= 0; i-- {
		cnt := hand[i]
		v := i
		if i > 10 {
			v = 10
		}

		var incrs [][]int
		for subCnt := cnt; subCnt >= 1; subCnt-- {
			subSum := v * subCnt
			waysToMake := maths.Choose(cnt, subCnt).ToInt()

			for addToBase := 15 - subSum; addToBase >= 0; addToBase-- {
				if sumCounts[addToBase] > 0 {
					incrs = append(incrs, []int{addToBase + subSum, waysToMake * sumCounts[addToBase]})
				}
			}

		}

		// fmt.Println("INCRS", incrs)

		for _, idxIncr := range incrs {
			sumCounts[idxIncr[0]] += idxIncr[1]
		}
		// fmt.Println("SUM COUNTS AFTER", i, sumCounts)
	}

	return sum + 2*sumCounts[15]
}

func handScore(hand []int) int {
	var sum int
	for i, cnt := range hand {
		if i > 10 {
			sum += 10 * cnt
		} else {
			sum += i * cnt
		}
	}
	return sum
}

// [0 2 3 0 0 0 0 2 0 0 1 0 0 1] 42 42 2304 map[0:1 1:2 2:4 3:6 4:6 5:6 6:4 7:4 8:5 9:8 10:14 11:16 12:20 13:20 14:17 15:16]
// A A 2 2 2 7 7 10 K
// A 2 2 10/K (12)
// 7 7 A (2)
// 7 2 2 2 A (2)
