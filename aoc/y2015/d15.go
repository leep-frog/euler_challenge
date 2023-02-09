package y2015

import (
	"regexp"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"golang.org/x/exp/maps"
)

func Day15() aoc.Day {
	return &day15{}
}

type day15 struct{}

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calroies   int
}

func (ig *ingredient) add(that *ingredient, volume int) *ingredient {
	return &ingredient{
		"",
		ig.capacity + (that.capacity * volume),
		ig.durability + (that.durability * volume),
		ig.flavor + (that.flavor * volume),
		ig.texture + (that.texture * volume),
		ig.calroies + (that.calroies * volume),
	}
}

func (d *day15) rec(ingredients map[string]*ingredient, rem int, amounts *ingredient, best1, best2 *maths.Bester[int, int]) {
	if len(ingredients) == 0 {
		if amounts.capacity > 0 && amounts.durability > 0 && amounts.flavor > 0 && amounts.texture > 0 {
			score := amounts.capacity * amounts.durability * amounts.flavor * amounts.texture
			best1.Check(score)
			if amounts.calroies == 500 {
				best2.Check(score)
			}
		}
		return
	}

	for _, name := range maps.Keys(ingredients) {
		ingr := ingredients[name]
		delete(ingredients, name)

		volumeStart := 0
		volumeEnd := rem
		if len(ingredients) == 0 {
			volumeStart = rem
		}
		for volume := volumeStart; volume <= volumeEnd; volume++ {
			d.rec(ingredients, rem-volume, amounts.add(ingr, volume), best1, best2)
		}
		ingredients[name] = ingr
	}
}

func (d *day15) Solve(lines []string, o command.Output) {
	r := regexp.MustCompile("^(.*): capacity (-?[0-9]+), durability (-?[0-9]+), flavor (-?[0-9]+), texture (-?[0-9]+), calories (-?[0-9]+)$")
	ingredients := map[string]*ingredient{}
	for _, line := range lines {
		m := r.FindStringSubmatch(line)
		ingredients[m[1]] = &ingredient{
			m[1],
			parse.Atoi(m[2]),
			parse.Atoi(m[3]),
			parse.Atoi(m[4]),
			parse.Atoi(m[5]),
			parse.Atoi(m[6]),
		}
	}

	best1 := maths.Largest[int, int]()
	best2 := maths.Largest[int, int]()
	d.rec(ingredients, 100, &ingredient{}, best1, best2)
	o.Stdoutln(best1.Best(), best2.Best())
}

func (d *day15) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"62842880 57600000",
			},
		},
		{
			ExpectedOutput: []string{
				"13882464 11171160",
			},
		},
	}
}
