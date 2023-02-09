package y2020

import (
	"regexp"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"golang.org/x/exp/maps"
)

func Day14() aoc.Day {
	return &day14{}
}

type day14 struct{}

type bitMask struct {
	parts  [][]int
	xCount int
}

func (bm *bitMask) rec(start int, options []int) []int {
	r := []int{start}
	for i, o := range options {
		r = append(r, bm.rec(start+o, options[i+1:])...)
	}
	return r
}

func (bm *bitMask) options(k int) []int {
	var opts []int
	for _, part := range bm.parts {
		pow := maths.Pow(2, part[0])
		isSet := k%(pow*2) >= pow

		if part[1] == 1 {
			if !isSet {
				k += pow
			}
		}

		if part[1] == -1 {
			if isSet {
				k -= pow
			}
			opts = append(opts, pow)
		}
	}
	return bm.rec(k, opts)
}

func (bm *bitMask) apply(k int) int {
	for _, part := range bm.parts {
		if part[1] == -1 {
			continue
		}
		pow, shouldSet := maths.Pow(2, part[0]), part[1] == 1

		isSet := k%(pow*2) >= pow
		if isSet && !shouldSet {
			k -= pow
		}
		if !isSet && shouldSet {
			k += pow
		}
	}
	return k
}

func (d *day14) makeBitMask(k string) *bitMask {
	var parts [][]int
	var xCount int
	for pow, v := range bread.Reverse(strings.Split(k, "")) {
		switch v {
		case "1":
			parts = append(parts, []int{pow, 1})
		case "0":
			parts = append(parts, []int{pow, 0})
		case "X":
			parts = append(parts, []int{pow, -1})
			xCount++
		}
	}
	return &bitMask{parts, xCount}
}

func (d *day14) Solve(lines []string, o command.Output) {
	var bm *bitMask
	maskRegex := regexp.MustCompile("^mask += +([X10]+)$")
	memRegex := regexp.MustCompile(`^mem\[([0-9]+)\] += +([0-9]+)$`)
	mem1, mem2 := map[int]int{}, map[int]int{}
	for _, line := range lines {
		m := maskRegex.FindStringSubmatch(line)
		if m != nil {
			bm = d.makeBitMask(m[1])
			continue
		}
		m = memRegex.FindStringSubmatch(line)
		address, value := parse.Atoi(m[1]), parse.Atoi(m[2])
		mem1[address] = bm.apply(value)
		for _, address := range bm.options(parse.Atoi(m[1])) {
			mem2[address] = parse.Atoi(m[2])
		}
	}
	o.Stdoutln(bread.Sum(maps.Values(mem1)), bread.Sum(maps.Values(mem2)))
}

func (d *day14) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"51 208",
			},
		},
		{
			ExpectedOutput: []string{
				"14722016054794 3618217244644",
			},
		},
	}
}
