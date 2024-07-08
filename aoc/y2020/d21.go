package y2020

import (
	"regexp"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/functional"
	"golang.org/x/exp/maps"
)

func Day21() aoc.Day {
	return &day21{}
}

type day21 struct{}

func (d *day21) Solve(lines []string, o command.Output) {
	r := regexp.MustCompile(`^(.*) \(contains (.*)\)$`)
	allergenMap := map[string]map[string]bool{}
	allFoods := map[string]int{}
	for _, line := range lines {
		m := r.FindStringSubmatch(line)
		foods := strings.Split(m[1], " ")
		foodMap := maths.NewSimpleSet(foods...)
		allergens := strings.Split(m[2], ", ")
		for _, f := range foods {
			allFoods[f]++
		}
		for _, a := range allergens {
			if am, ok := allergenMap[a]; ok {
				allergenMap[a] = maths.Intersection(foodMap, am)
			} else {
				allergenMap[a] = map[string]bool{}
				for _, f := range foods {
					allergenMap[a][f] = true
				}
			}
		}
	}

	for _, am := range allergenMap {
		for f := range am {
			delete(allFoods, f)
		}
	}
	var sum int
	for _, v := range allFoods {
		sum += v
	}

	solvedIngredients := map[string]string{}
	for len(allergenMap) > 0 {
		var removeAllergens, removeIngredients []string
		for allergen, ingredients := range allergenMap {
			if len(ingredients) == 1 {
				ingr := maps.Keys(ingredients)[0]
				removeAllergens = append(removeAllergens, allergen)
				removeIngredients = append(removeIngredients, ingr)
				solvedIngredients[ingr] = allergen
				continue
			}
		}

		for _, im := range allergenMap {
			for _, tr := range removeIngredients {
				delete(im, tr)
			}
		}
		for _, r := range removeAllergens {
			delete(allergenMap, r)
		}
	}

	var soln [][]string
	for i, a := range solvedIngredients {
		soln = append(soln, []string{i, a})
	}

	functional.SortFunc(soln, func(this, that []string) bool {
		return this[1] < that[1]
	})
	o.Stdoutln(sum, strings.Join(functional.Map(soln, func(s []string) string { return s[0] }), ","))
}

func (d *day21) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"5 mxmxvkd,sqjhc,fvjkl",
			},
		},
		{
			ExpectedOutput: []string{
				"1685 ntft,nhx,kfxr,xmhsbd,rrjb,xzhxj,chbtp,cqvc",
			},
		},
	}
}
