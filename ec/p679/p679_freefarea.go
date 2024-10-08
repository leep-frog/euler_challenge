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
	letters = []string{
		"a",
		"e",
		"f",
		"r",
	}
	wordMap = map[string]bool{
		"free": true,
		"reef": true,
		"area": true,
		"fare": true,
	}
)

func P679() *ecmodels.Problem {
	return ecmodels.IntInputNode(679, func(o command.Output, n int) {
		o.Stdoutln(dp(n, maps.Clone(wordMap), nil))
	}, []*ecmodels.Execution{
		{
			Args: []string{"2"},
			Want: "",
		},
	})
}

var (
	cache = map[string]int{}
)

func dp(remaining int, wordsNeeded map[string]bool, currentWord []string) int {
	if remaining == 0 {
		// fmt.Println(currentWord, wordsNeeded)
		if len(wordsNeeded) == 0 {
			return 1
		}
		return 0
	}

	var start int
	if len(currentWord) > 3 {
		start = len(currentWord) - 3
	}
	codeParts := []string{
		fmt.Sprintf("%d", remaining),
		"PRE",
		strings.Join(currentWord[start:], ""),
		"POST",
	}

	keys := maps.Keys(wordsNeeded)
	slices.Sort(keys)
	codeParts = append(codeParts, keys...)
	code := strings.Join(codeParts, " ")
	if v, ok := cache[code]; ok {
		return v
	}
	// if cac

	var sum int
	for _, letter := range letters {
		nextWord := append(currentWord, letter)

		var removedWord string

		if len(nextWord) >= 4 {
			currentWordString := strings.Join(nextWord[len(nextWord)-4:], "")

			if wordsNeeded[currentWordString] {
				removedWord = currentWordString
				delete(wordsNeeded, currentWordString)
			} else if wordMap[currentWordString] {
				continue
			}

			// if v, ok := wordMap[currentWordString]; ok {

			// }
		}
		sum += dp(remaining-1, wordsNeeded, nextWord)

		if removedWord != "" {
			wordsNeeded[removedWord] = true
		}
	}

	cache[code] = sum
	return sum
}
