package y2015

import (
	"fmt"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day21() aoc.Day {
	return &day21{}
}

type day21 struct{}

func (d *day21) Solve(lines []string, o command.Output) {
	weapons := []*item{
		{8, 4, 0},
		{10, 5, 0},
		{25, 6, 0},
		{40, 7, 0},
		{74, 8, 0},
	}

	armor := []*item{
		{0, 0, 0},
		{13, 0, 1},
		{31, 0, 2},
		{53, 0, 3},
		{75, 0, 4},
		{102, 0, 5},
	}

	rings := []*item{
		{0, 0, 0},
		{0, 0, 0},
		{25, 1, 0},
		{50, 2, 0},
		{100, 3, 0},
		{20, 0, 1},
		{40, 0, 2},
		{80, 0, 3},
	}

	boss := &player{
		parse.Atoi(strings.Split(lines[0], ": ")[1]),
		parse.Atoi(strings.Split(lines[1], ": ")[1]),
		parse.Atoi(strings.Split(lines[2], ": ")[1]),
	}

	part1 := maths.Smallest[int, int]()
	part2 := maths.Largest[int, int]()
	for _, w := range weapons {
		for _, a := range armor {
			for i, r1 := range rings {
				for _, r2 := range rings[i+1:] {
					cost := w.cost + a.cost + r1.cost + r2.cost
					p := &player{
						100,
						w.weapon + a.weapon + r1.weapon + r2.weapon,
						w.armor + a.armor + r1.armor + r2.armor,
					}

					if p.fight(boss) {
						part1.Check(cost)
					} else {
						part2.Check(cost)
					}
				}
			}
		}
	}
	o.Stdoutln(part1.Best(), part2.Best())
}

type item struct {
	cost   int
	weapon int
	armor  int
}

type player struct {
	hp     int
	damage int
	armor  int
}

func (p *player) String() string {
	return fmt.Sprintf("{hp:%d, damage:%d, armor:%d}", p.hp, p.damage, p.armor)
}

func (p *player) copy() *player {
	return &player{p.hp, p.damage, p.armor}
}

func (p *player) fight(boss *player) bool {
	pDamage := maths.Max(1, p.damage-boss.armor)
	bossDamage := maths.Max(1, boss.damage-p.armor)

	playerHits := boss.hp / pDamage
	if boss.hp%pDamage != 0 {
		playerHits++
	}

	bossHits := p.hp / bossDamage
	if p.hp%bossDamage != 0 {
		bossHits++
	}

	return playerHits <= bossHits
}

func (d *day21) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			ExpectedOutput: []string{
				"91 158",
			},
		},
	}
}
