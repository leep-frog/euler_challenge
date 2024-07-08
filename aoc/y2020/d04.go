package y2020

import (
	"regexp"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day04() aoc.Day {
	return &day04{}
}

type day04 struct{}

func (d *day04) required() map[string]func(string) bool {
	hairColor := regexp.MustCompile("^[0-9a-f]+$")
	numbers := regexp.MustCompile("^[0-9]+$")
	height := regexp.MustCompile("^([0-9]+)([a-z]+)$")
	eyeColor := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
	return map[string]func(string) bool{
		"byr": func(s string) bool {
			y := parse.Atoi(s)
			return y >= 1920 && y <= 2002
		},
		"iyr": func(s string) bool {
			y := parse.Atoi(s)
			return y >= 2010 && y <= 2020
		},
		"eyr": func(s string) bool {
			y := parse.Atoi(s)
			return y >= 2020 && y <= 2030
		},
		"hgt": func(s string) bool {
			m := height.FindStringSubmatch(s)
			if len(m) != 3 {
				return false
			}

			h := parse.Atoi(m[1])
			if m[2] == "cm" {
				return h >= 150 && h <= 193
			}
			if m[2] == "in" {
				return h >= 59 && h <= 76
			}
			return false
		},
		"hcl": func(s string) bool {
			if s[0] != '#' {
				return false
			}
			return len(s) == 7 && hairColor.MatchString(s[1:])
		},
		"ecl": func(s string) bool {
			return eyeColor[s]
		},
		"pid": func(s string) bool {
			return len(s) == 9 && numbers.MatchString(s)
		},
		// "cid": true,
	}
}

func (d *day04) Solve(lines []string, o command.Output) {
	required1, required2 := d.required(), d.required()
	var sum1, sum2 int
	for _, line := range append(lines, "") {
		if line == "" {
			if len(required1) == 0 {
				sum1++
			}
			if len(required2) == 0 {
				sum2++
			}
			required1, required2 = d.required(), d.required()
		}

		for _, part := range strings.Split(line, " ") {
			split := strings.Split(part, ":")
			code := split[0]
			delete(required1, code)
			if f, ok := required2[code]; ok {
				if f(split[1]) {
					delete(required2, code)
				}
			}
		}
	}
	o.Stdoutln(sum1, sum2)
}

func (d *day04) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"10 6",
			},
		},
		{
			ExpectedOutput: []string{
				"204 179",
			},
		},
	}
}
