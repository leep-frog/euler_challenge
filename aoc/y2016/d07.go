package y2016

import (
	"fmt"
	"regexp"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
)

func Day07() aoc.Day {
	return &day07{}
}

type day07 struct{}

func (d *day07) valid1(parts []string) bool {
	for idx := 0; idx < len(parts); idx += 2 {
		txt := parts[idx]
		for i := 0; i < len(txt)-3; i++ {
			if txt[i] != txt[i+1] && txt[i:i+2] == (txt[i+3:i+4]+txt[i+2:i+3]) {
				return true
			}
		}
	}
	return false
}

func (d *day07) abas(parts []string) map[string]bool {
	r := map[string]bool{}
	for idx := 0; idx < len(parts); idx += 2 {
		txt := parts[idx]
		for i := 0; i < len(txt)-2; i++ {
			if txt[i] != txt[i+1] && txt[i] == txt[i+2] {
				r[txt[i:i+3]] = true
			}
		}
	}
	return r
}

func (d *day07) babs(parts []string, abas map[string]bool) bool {
	for idx := 0; idx < len(parts); idx += 2 {
		txt := parts[idx]
		for i := 0; i < len(txt)-2; i++ {
			if txt[i] != txt[i+1] && txt[i] == txt[i+2] {
				// fmt.Println("CHECKING", string(txt[i]), string(txt[i+1]), string(txt[i+2]), txt[i:i+3], txt[i+1:i+3]+txt[i+1:i+1])
				if abas[txt[i+1:i+3]+txt[i+1:i+2]] {
					return true
				}
			}
		}
	}
	return false
}

func (d *day07) Solve(lines []string, o command.Output) {
	var cnt, cnt2 int
	r := regexp.MustCompile(`[\[\]]`)
	for _, line := range lines {
		parts := r.Split(line, -1)
		if d.valid1(parts) && !d.valid1(parts[1:]) {
			fmt.Println("VALID", line)
			cnt++
		}

		abas := d.abas(parts)
		fmt.Println(line, abas)
		if d.babs(parts[1:], abas) {
			cnt2++
		}
	}
	fmt.Println(cnt, cnt2)
}

func (d *day07) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"",
			},
		},
		{
			ExpectedOutput: []string{
				"",
			},
		},
	}
}
