package y2022

import (
	"regexp"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day21() aoc.Day {
	return &day21{}
}

type day21 struct{}

type mathMonkey struct {
	left      string
	operation string
	right     string
}

func (d *day21) evaluateMonkeyValue(monkeyID string, monkeyNumbers map[string]int, monkeys map[string]*mathMonkey) int {
	if v, ok := monkeyNumbers[monkeyID]; ok {
		return v
	}

	monkey := monkeys[monkeyID]

	left, right := d.evaluateMonkeyValue(monkey.left, monkeyNumbers, monkeys), d.evaluateMonkeyValue(monkey.right, monkeyNumbers, monkeys)
	switch monkey.operation {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	}
	panic("Unknown monkey operation")
}

func (d *day21) evaluateHumanValue(shouldEqual int, monkeyID string, monkeys map[string]*mathMonkey, hasHuman map[string]bool, monkeyNumbers map[string]int) int {
	if monkeyID == "humn" {
		return shouldEqual
	}

	m := monkeys[monkeyID]

	left, right := m.left, m.right
	if hasHuman[right] == hasHuman[left] {
		panic("Both sides cannot both have the human")
	}

	// Left will be the one that hasHuman the human
	var swapped bool
	if hasHuman[right] {
		swapped = true
		left, right = right, left
	}

	rightValue := d.evaluateMonkeyValue(right, monkeyNumbers, monkeys)

	var newEqual int
	switch monkeys[monkeyID].operation {
	case "+":
		newEqual = shouldEqual - rightValue
	case "-":
		if swapped {
			newEqual = rightValue - shouldEqual
		} else {
			newEqual = shouldEqual + rightValue
		}
	case "*":
		newEqual = shouldEqual / rightValue
	case "/":
		if swapped {
			newEqual = rightValue / shouldEqual
		} else {
			newEqual = shouldEqual * rightValue
		}
	}
	return d.evaluateHumanValue(newEqual, left, monkeys, hasHuman, monkeyNumbers)
}

func (d *day21) populateHasHumanMap(monkeyID string, monkeys map[string]*mathMonkey, hasHuman map[string]bool) bool {
	if v, ok := hasHuman[monkeyID]; ok {
		return v
	}

	if monkeyID == "humn" {
		hasHuman[monkeyID] = true
		return true
	}

	m, ok := monkeys[monkeyID]
	if !ok {
		// Number monkey
		hasHuman[monkeyID] = false
		return false
	}

	if d.populateHasHumanMap(m.left, monkeys, hasHuman) || d.populateHasHumanMap(m.right, monkeys, hasHuman) {
		hasHuman[monkeyID] = true
		return true
	}

	hasHuman[monkeyID] = false
	return false
}

func (d *day21) Solve(lines []string, o command.Output) {
	numberMonkey := regexp.MustCompile("^([a-z]+): ([0-9]+)$")
	opMonkey := regexp.MustCompile(`^([a-z]+): ([a-z]+) ([\+\-\/\*]) ([a-z]+)`)

	// Map for number monkeys
	monkeyNumbers := map[string]int{}
	// Map for math monkeys
	monkeys := map[string]*mathMonkey{}

	// Populate monkey maps
	for _, line := range lines {
		nm := numberMonkey.FindStringSubmatch(line)
		if nm != nil {
			monkeyNumbers[nm[1]] = parse.Atoi(nm[2])
			continue
		}
		om := opMonkey.FindStringSubmatch(line)
		monkeys[om[1]] = &mathMonkey{om[2], om[3], om[4]}
	}

	// Populate hasHuman map. hasHuman[monkeyID] == true iff the monkey's value
	// depends on the human's value.
	hasHuman := map[string]bool{}
	d.populateHasHumanMap("root", monkeys, hasHuman)

	// Solve part 1
	part1 := d.evaluateMonkeyValue("root", monkeyNumbers, monkeys)

	// Solve part 2
	root := monkeys["root"]
	var part2 int
	if hasHuman[root.left] {
		part2 = d.evaluateHumanValue(d.evaluateMonkeyValue(root.right, monkeyNumbers, monkeys), root.left, monkeys, hasHuman, monkeyNumbers)
	} else {
		part2 = d.evaluateHumanValue(d.evaluateMonkeyValue(root.left, monkeyNumbers, monkeys), root.right, monkeys, hasHuman, monkeyNumbers)
	}

	o.Stdoutln(part1, part2)
}

func (d *day21) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"152 301",
			},
		},
		{
			ExpectedOutput: []string{
				"331120084396440 3378273370680",
			},
		},
	}
}
