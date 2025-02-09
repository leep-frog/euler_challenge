package p98

import (
	"sort"
	"strconv"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P98() *ecmodels.Problem {
	return ecmodels.FileInputNode(98, func(lines []string, o command.Output) {
		words := strings.Split(strings.ReplaceAll(lines[0], "\"", ""), ",")

		best := maths.Largest[int, int]()

		var maxLen int
		wordsByLength := map[int][]string{}
		wordSet := map[string]bool{}
		for _, word := range words {
			wordSet[word] = true
			wordsByLength[len(word)] = append(wordsByLength[len(word)], word)
			maxLen = maths.Max(maxLen, len(word))
		}

		var atLeastOne bool
		for length := maxLen; length >= 0; length-- {
			anagramSets := map[string]map[string]bool{}
			for _, word := range wordsByLength[length] {
				chars := []string{}
				for i := 0; i < len(word); i++ {
					chars = append(chars, word[i:i+1])
				}
				sort.Strings(chars)
				sorted := strings.Join(chars, "")
				if anagramSets[sorted] == nil {
					anagramSets[sorted] = map[string]bool{}
				}
				anagramSets[sorted][word] = true
			}

			// Remove redundant anagram sets
			remove := []string{}
			for key, set := range anagramSets {
				if len(set) == 1 {
					remove = append(remove, key)
				}
			}
			for _, key := range remove {
				delete(anagramSets, key)
			}

			if len(anagramSets) > 0 {
				patternMap, anagramMap := createAnagramMap(length)

				for _, set := range anagramSets {
					for word := range set {
						pattern := stringAnagramPattern(word)
						// Iterate over each square that matches the word's pattern
						for _, square := range patternMap[pattern] {
							// Create the mapping
							mapping := map[int]string{}
							for i, d := range maths.Digits(square) {
								mapping[d] = word[i : i+1]
							}

							// Iterate over anagrams and check if they make words
							for a := range anagramMap[anagramStringSet(square)] {
								if a == square {
									continue
								}
								var newWord []string
								for _, d := range maths.Digits(a) {
									newWord = append(newWord, mapping[d])
								}

								if wordSet[strings.Join(newWord, "")] {
									best.Check(square)
									best.Check(a)
									atLeastOne = true
								}
							}
						}
					}
				}
			}

			if atLeastOne {
				o.Stdoutln(best.Best())
				return
			}
		}
	}, []*ecmodels.Execution{
		{
			Want: "18769",
		},
	})
}

func anagramPattern(n int) string {
	return stringAnagramPattern(strconv.Itoa(n))
}

func stringAnagramPattern(nStr string) string {
	letterIdx := 0
	checked := map[string]string{}
	var r []string
	for i := 0; i < len(nStr); i++ {
		c := nStr[i : i+1]
		if _, ok := checked[c]; !ok {
			checked[c] = ecmodels.Letters[letterIdx : letterIdx+1]
			letterIdx++
		}
		r = append(r, checked[c])
	}
	return strings.Join(r, "")
}

func createAnagramMap(length int) (map[string][]int, map[string]map[int]bool) {
	maxSquare := 1
	for i := 0; i < length; i++ {
		maxSquare *= 10
	}

	start := maths.Sqrt(maxSquare/10) - 1
	end := maths.Sqrt(maxSquare) + 1

	// Map from pattern (22171278 (value) becomes AABCBACD (key)) to list of squares with that pattern
	patternMap := map[string][]int{}
	for i := start; i <= end; i++ {
		square := i * i
		pattern := anagramPattern(square) // AABCBACD
		patternMap[pattern] = append(patternMap[pattern], square)
	}

	// Map from sorted digits to squares containing those digits
	anagramMap := map[string]map[int]bool{}
	for i := start; i <= end; i++ {
		square := i * i
		sorted := anagramStringSet(square)
		if anagramMap[sorted] == nil {
			anagramMap[sorted] = map[int]bool{}
		}
		anagramMap[sorted][square] = true
	}

	return patternMap, anagramMap
}

func anagramStringSet(n int) string {
	var digits []string
	for _, d := range maths.Digits(n) {
		digits = append(digits, strconv.Itoa(d))
	}
	sort.Strings(digits)
	return strings.Join(digits, "")
}
