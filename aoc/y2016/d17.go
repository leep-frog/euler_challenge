package y2016

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
)

func Day17() aoc.Day {
	return &day17{}
}

type day17 struct{}

func (d *day17) Solve(lines []string, o command.Output) {
	passcode := lines[0]

	shortest, longest := maths.Smallest[string, int](), maths.Largest[string, int]()
	start := &room{0, 0, passcode}
	start.dfs(0, shortest, longest)

	bestPath := shortest.BestIndex()[len(passcode):]
	o.Stdoutln(bestPath, longest.Best())
}

type room struct {
	x, y int
	path string
}

var (
	validHashes = map[rune]bool{
		'b': true,
		'c': true,
		'd': true,
		'e': true,
		'f': true,
	}
)

func (r *room) dfs(depth int, shortest, longest *maths.Bester[string, int]) {
	// Check if done
	if r.x == 3 && r.y == 3 {
		shortest.IndexCheck(r.path, depth)
		longest.IndexCheck(r.path, depth)
		return
	}

	hash := md5Hash(r.path, 1)[:4]

	// Up
	if r.y > 0 && validHashes[rune(hash[0])] {
		newR := &room{r.x, r.y - 1, r.path + "U"}
		newR.dfs(depth+1, shortest, longest)
	}

	if r.y < 3 && validHashes[rune(hash[1])] {
		newR := &room{r.x, r.y + 1, r.path + "D"}
		newR.dfs(depth+1, shortest, longest)
	}

	if r.x > 0 && validHashes[rune(hash[2])] {
		newR := &room{r.x - 1, r.y, r.path + "L"}
		newR.dfs(depth+1, shortest, longest)
	}

	if r.x < 3 && validHashes[rune(hash[3])] {
		newR := &room{r.x + 1, r.y, r.path + "R"}
		newR.dfs(depth+1, shortest, longest)
	}
}

func (d *day17) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"DDUDRLRRUDRD 492",
			},
		},
		{
			ExpectedOutput: []string{
				"590 RLDUDRDDRR",
			},
		},
	}
}
