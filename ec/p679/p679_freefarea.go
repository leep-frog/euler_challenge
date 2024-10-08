package p679

import (
	"fmt"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

var (
	letters = map[string]int{
		"a": 1,
		"e": 2,
		"f": 3,
		"r": 4,
	}
	words = map[string]bool{
		"free": true,
		"reef": true,
		"area": true,
		"fare": true,
	}
	wordMap = map[int]bool{}
)

func P679() *ecmodels.Problem {
	return ecmodels.IntInputNode(679, func(o command.Output, n int) {

		for word := range words {
			var wordNumber int
			for _, l := range strings.Split(word, "") {
				wordNumber = 10*wordNumber + letters[l]
			}
			wordMap[wordNumber] = true
		}

		fmt.Println(wordMap)

		o.Stdoutln(dp(n, maps.Clone(wordMap), 0))
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

func dp(remaining int, wordsNeeded map[int]bool, currentWord int) int {
	if remaining == 0 {
		if len(wordsNeeded) == 0 {
			return 1
		}
		return 0
	}

	codeParts := []string{
		fmt.Sprintf("%d %d", remaining, currentWord%1000),
	}

	keys := maps.Keys(wordsNeeded)
	slices.Sort(keys)
	for _, k := range keys {
		codeParts = append(codeParts, fmt.Sprintf("%d", k))
	}

	code := strings.Join(codeParts, " ")
	if v, ok := cache[code]; ok {
		return v
	}

	var sum int
	for _, letter := range letters {
		nextWord := 10*currentWord + letter

		var removedWord bool

		if nextWord >= 1000 {
			// currentWordString := strings.Join(nextWord[len(nextWord)-4:], "")

			if wordsNeeded[nextWord] {
				removedWord = true
				delete(wordsNeeded, nextWord)
			} else if wordMap[nextWord] {
				continue
			}
		}
		sum += dp(remaining-1, wordsNeeded, nextWord%1000)

		if removedWord {
			wordsNeeded[nextWord] = true
		}
	}

	cache[code] = sum
	return sum
}
