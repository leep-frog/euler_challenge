package y2015

import (
	"regexp"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/point"
)

func Day06() aoc.Day {
	return &day06{}
}

type day06 struct{}

type day6part1 struct {
	on map[string]bool
}

func (d1 *day6part1) toggle(code string) {
	if d1.on[code] {
		delete(d1.on, code)
	} else {
		d1.on[code] = true
	}
}

func (d1 *day6part1) turnOn(code string) {
	d1.on[code] = true
}

func (d1 *day6part1) turnOff(code string) {
	delete(d1.on, code)
}

func (d1 *day6part1) count() int {
	return len(d1.on)
}

type day6part2 struct {
	on map[string]int
}

func (d2 *day6part2) toggle(code string) {
	d2.on[code] += 2
}

func (d2 *day6part2) turnOn(code string) {
	d2.on[code]++
}

func (d2 *day6part2) turnOff(code string) {
	d2.on[code] = maths.Max(0, d2.on[code]-1)
}

func (d2 *day6part2) count() int {
	var sum int
	for _, v := range d2.on {
		sum += v
	}
	return sum
}

type lighter interface {
	toggle(string)
	turnOn(string)
	turnOff(string)
	count() int
}

func (d *day06) Solve(lines []string, o command.Output) {
	o.Stdoutln(d.solve(lines, &day6part1{map[string]bool{}}), d.solve(lines, &day6part2{map[string]int{}}))
}

func (d *day06) solve(lines []string, lighter lighter) int {
	r := regexp.MustCompile("^(.*) ([0-9,]*) through ([0-9,]*)$")
	for _, line := range lines {
		m := r.FindStringSubmatch(line)
		leftCoordStr := parse.AtoiArray(strings.Split(m[2], ","))
		rightCoordStr := parse.AtoiArray(strings.Split(m[3], ","))
		leftCoord := point.New(leftCoordStr[0], leftCoordStr[1])
		rightCoord := point.New(rightCoordStr[0], rightCoordStr[1])
		for x := leftCoord.X; x <= rightCoord.X; x++ {
			for y := leftCoord.Y; y <= rightCoord.Y; y++ {
				code := point.New(x, y).String()
				switch m[1] {
				case "toggle":
					lighter.toggle(code)
				case "turn on":
					lighter.turnOn(code)
				case "turn off":
					lighter.turnOff(code)
				default:
					panic("Unknown")
				}
			}
		}
	}
	return lighter.count()
}

func (d *day06) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"998996 1001996",
			},
		},
		{
			ExpectedOutput: []string{
				"377891 14110788",
			},
		},
	}
}
