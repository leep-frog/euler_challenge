package p679

import (
	"fmt"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
)

var (
	letters = map[string]int{
		"a": 1,
		"e": 3,
		"f": 7,
		"r": 9,
	}
	numbers = []int{}
	words   = map[string]bool{
		"free": true,
		"reef": true,
		"area": true,
		"fare": true,
	}
	wordNumberMap = map[int]bool{}
)

func P679() *ecmodels.Problem {
	return ecmodels.IntInputNode(679, func(o command.Output, n int) {

		// Populate the numbers slice
		for _, letterValue := range letters {
			numbers = append(numbers, letterValue)
		}

		// Populate the wordNumberMap
		p := generator.Primes()
		wordProduct := 1
		for word := range words {
			var wordNumber int
			for _, l := range strings.Split(word, "") {
				wordNumber = 10*wordNumber + letters[l]
			}
			wordNumberMap[wordNumber] = true
			if !p.Contains(wordNumber) {
				o.Stderrf("All word numbers must be prime, but %d has factors %v\n", wordNumber, p.PrimeFactors(wordNumber))
				return
			}
			wordProduct *= wordNumber
		}

		o.Stdoutln(dp(n, wordProduct, 0))
	}, []*ecmodels.Execution{
		{
			Args: []string{"9"},
			Want: "1",
		},
		{
			Args: []string{"15"},
			Want: "72863",
		},
		{
			Args: []string{"30"},
			Want: "644997092988678",
		},
	})
}

var (
	cache = map[string]int{}
)

// Rather than use strings, we use integers for faster processing.
// wordsNeeded is really a set of integers, but we mimic that by making it the
// product of all the wordNumbers (note this requires that all wordNumbers are
// primes which is guaratneed above)
func dp(remaining int, wordsNeeded int, currentWord int) int {
	if remaining == 0 {
		if wordsNeeded == 1 {
			return 1
		}
		return 0
	}

	code := fmt.Sprintf("%d %d %d", remaining, wordsNeeded, currentWord)
	if v, ok := cache[code]; ok {
		return v
	}

	var sum int
	for _, number := range numbers {
		nextWord := 10*currentWord + number

		var removedWord bool

		if wordsNeeded%nextWord == 0 {
			removedWord = true
			wordsNeeded /= nextWord
		} else if wordNumberMap[nextWord] {
			continue
		}

		sum += dp(remaining-1, wordsNeeded, nextWord%1000)

		if removedWord {
			wordsNeeded *= nextWord
		}
	}

	cache[code] = sum
	return sum
}
