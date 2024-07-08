package y2016

import (
	"crypto/md5"
	"fmt"
	"io"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func Day14() aoc.Day {
	return &day14{}
}

type day14 struct{}

type possibleKey struct {
	index int
	key   string
}

func md5Hash(s string, times int) string {
	for i := 0; i < times; i++ {
		h := md5.New()
		if _, err := io.WriteString(h, s); err != nil {
			panic("ARGH")
		}
		s = fmt.Sprintf("%x", h.Sum(nil))
	}
	return s
}

func (d *day14) Solve(lines []string, o command.Output) {
	o.Stdoutln(d.solve(lines, 1), d.solve(lines, 2017))
}

func (d *day14) solve(lines []string, times int) int {
	salt := lines[0]
	keys := map[int]bool{}
	possibleKeys := map[rune][]*possibleKey{}
	var untilCount int

	needKeys := 64
	delay := 1000

	for i := 0; untilCount < delay; i++ {
		if len(keys) >= needKeys {
			untilCount++
		}

		hashed := md5Hash(fmt.Sprintf("%s%d", salt, i), times)

		var prevLetter rune
		var letterCount int
		firstTriplet := true
		// Need the separate map so a hash can't be verified by itself. So we don't
		// check a triple and quintuple in the same key 'abcddddd123'
		toAdd := map[rune][]*possibleKey{}
		for ci, c := range hashed {
			if c == prevLetter {
				letterCount++
			} else {
				prevLetter = c
				letterCount = 1
			}

			if letterCount == 3 && firstTriplet {
				firstTriplet = false

				toAdd[prevLetter] = append(toAdd[prevLetter], &possibleKey{i, hashed})
			}

			if letterCount == 5 && (ci >= len(hashed) || rune(hashed[ci+1]) != prevLetter) {
				for _, pk := range possibleKeys[prevLetter] {
					if i-pk.index < delay {
						keys[pk.index] = true
					}
				}
				possibleKeys[prevLetter] = nil
			}
		}

		for k, v := range toAdd {
			possibleKeys[k] = append(possibleKeys[k], v...)
		}
	}
	keyArr := maps.Keys(keys)
	slices.Sort(keyArr)
	return keyArr[needKeys-1]
}

func (d *day14) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"22728 27039",
			},
		},
		{
			ExpectedOutput: []string{
				"23890 22696",
			},
		},
	}
}
