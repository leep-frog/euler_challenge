package y2015

import (
	"fmt"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day22() aoc.Day {
	return &day22{}
}

type day22 struct{}

func (d *day22) Solve(lines []string, o command.Output) {
	m := map[string]int{}
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		m[parts[0]] = parse.Atoi(parts[1])
	}

	boss := &player{
		hp:     m["Boss Hit Points"],
		damage: m["Boss Damage"],
	}

	w := &wizard{
		m["Wizard Hit Points"],
		m["Wizard Mana"],
		0,
		0,
		nil,
	}

	var parts []string
	for _, part2 := range []bool{false, true} {
		_, b := bfs.DistanceSearch[string, bfs.Int]([]*turn{
			{true, w.copy(), boss.copy(), map[string]*effect{}, "start", part2},
		}, bfs.CumulativeDistanceFunction())
		parts = append(parts, parse.Itos(int(b)))
	}
	o.Stdoutln(strings.Join(parts, " "))
}

type effect struct {
	code string
	turn int
	f    func(*wizard, *player, int)
}

func (e *effect) String() string {
	return fmt.Sprintf("%s:%d", e.code, e.turn)
}

func (e *effect) copy() *effect {
	return &effect{e.code, e.turn, e.f}
}

type wizard struct {
	hp        int
	mana      int
	manaSpent int
	armor     int
	effects   []effect
}

func (w *wizard) String() string {
	return fmt.Sprintf("{hp:%d, mana%d, ms:%d, armor:%d}", w.hp, w.mana, w.manaSpent, w.armor)
}

func (w *wizard) copy() *wizard {
	return &wizard{w.hp, w.mana, w.manaSpent, w.armor, w.effects}
}

func (w *wizard) spendMana(mana int) bool {
	if mana > w.mana {
		return false
	}
	w.mana -= mana
	w.manaSpent += mana
	return true
}

type turn struct {
	wizardTurn bool
	w          *wizard
	boss       *player
	effects    map[string]*effect
	move       string
	part2      bool
}

func (t *turn) copy() *turn {
	effs := map[string]*effect{}
	for k, v := range t.effects {
		effs[k] = v.copy()
	}
	return &turn{t.wizardTurn, t.w.copy(), t.boss.copy(), effs, t.move, t.part2}
}

func (t *turn) endTurn() {
	var toDel []string
	for code, eff := range t.effects {
		eff.turn--
		eff.f(t.w, t.boss, eff.turn)
		if eff.turn == 0 {
			toDel = append(toDel, code)
		}
	}
	for _, td := range toDel {
		delete(t.effects, td)
	}
	if t.wizardTurn && t.part2 {
		t.w.hp--
	}
	t.wizardTurn = !t.wizardTurn
}

func (t *turn) Code() string {
	return fmt.Sprintf("%v %v %v %v", t.wizardTurn, t.w, t.boss, t.effects)
}

func (t *turn) String() string {
	return fmt.Sprintf("MOVE:%s ; %s", t.move, t.Code())
}

func (t *turn) Done() bool {
	return t.boss.hp <= 0 && t.w.hp > 0 && t.w.mana > 0
}

func (t *turn) Distance() bfs.Int {
	return bfs.Int(t.w.manaSpent)
}

func (t *turn) AdjacentStates() []*turn {
	if t.boss.hp <= 0 || t.w.hp <= 0 {
		return nil
	}

	if !t.wizardTurn {
		return []*turn{t.copy().bossMove()}
	}

	ts := []*turn{
		t.copy().magicMissile(),
		t.copy().drain(),
		t.copy().shield(),
		t.copy().poison(),
		t.copy().recharge(),
	}

	var newTs []*turn
	for _, nt := range ts {
		if nt == nil || nt.w.mana < 0 {
			continue
		}
		nt.endTurn()
		newTs = append(newTs, nt)
	}

	return newTs
}

func (t *turn) bossMove() *turn {
	t.w.hp -= maths.Max(1, t.boss.damage-t.w.armor)
	t.move = "boss"
	t.endTurn()
	return t
}

func (t *turn) magicMissile() *turn {
	// Magic Missile costs 53 mana. It instantly does 4 damage.
	if !t.w.spendMana(53) {
		return nil
	}
	t.boss.hp -= 4
	t.move = "magic missile"
	return t
}

func (t *turn) drain() *turn {
	// Drain costs 73 mana. It instantly does 2 damage and heals you for 2 hit points.
	if !t.w.spendMana(73) {
		return nil
	}
	t.boss.hp -= 2
	t.w.hp += 2
	t.move = "drain"
	return t
}

var (
	shieldMove   = "shield"
	poisonMove   = "poison"
	rechargeMove = "recharge"
)

func (t *turn) shield() *turn {
	// Shield costs 113 mana. It starts an effect that lasts for 6 turns. While it is active, your armor is increased by 7.
	if _, ok := t.effects[shieldMove]; ok {
		return nil
	}

	if !t.w.spendMana(113) {
		return nil
	}
	t.w.armor += 7
	t.effects[shieldMove] = &effect{shieldMove, 6, func(w *wizard, p *player, turn int) {
		if turn == 0 {
			w.armor -= 7
		}
	}}
	t.move = "shield"
	return t
}

func (t *turn) poison() *turn {
	// Poison costs 173 mana. It starts an effect that lasts for 6 turns. At the start of each turn while it is active, it deals the boss 3 damage.
	if _, ok := t.effects[poisonMove]; ok {
		return nil
	}

	if !t.w.spendMana(173) {
		return nil
	}
	t.effects[poisonMove] = &effect{poisonMove, 6, func(w *wizard, boss *player, turn int) {
		boss.hp -= 3
	}}
	t.move = "poison"
	return t
}

func (t *turn) recharge() *turn {
	// Recharge costs 229 mana. It starts an effect that lasts for 5 turns. At the start of each turn while it is active, it gives you 101 new mana.
	if _, ok := t.effects[rechargeMove]; ok {
		return nil
	}

	if !t.w.spendMana(229) {
		return nil
	}
	t.effects[rechargeMove] = &effect{rechargeMove, 5, func(w *wizard, boss *player, turn int) {
		w.mana += 101
	}}
	t.move = "recharge"
	return t
}

func (d *day22) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"641 0",
			},
		},
		{
			ExpectedOutput: []string{
				"900 1216",
			},
		},
	}
}
