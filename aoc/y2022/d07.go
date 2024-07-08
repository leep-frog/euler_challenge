package y2022

import (
	"fmt"
	"math"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func Day07() aoc.Day {
	return &day07{}
}

type day07 struct{}

type directory struct {
	name           string
	size           int
	subDirectories map[string]*directory
}

func (d *directory) String() string {
	return d.depthString(0) + "\n"
}

func (d *directory) depthString(depth int) string {
	r := []string{fmt.Sprintf("%s%s: %d", strings.Repeat(" ", depth), d.name, d.size)}
	keys := maps.Keys(d.subDirectories)
	slices.Sort(keys)
	for _, k := range keys {
		r = append(r, d.subDirectories[k].depthString(depth+2))
	}
	return strings.Join(r, "\n")
}

func (d *directory) addDirSizes() {
	for _, sub := range d.subDirectories {
		sub.addDirSizes()
		d.size += sub.size
	}
}

// aka part1
func (d *directory) smallDirSizeSum() int {
	var count int
	if d.size <= 100_000 {
		count += d.size
	}
	for _, sub := range d.subDirectories {
		count += sub.smallDirSizeSum()
	}
	return count
}

// return the smallest dir that is at least as big as atLeast
func (d *directory) smallestDir(atleast int) int {
	best := math.MaxInt
	if d.size >= atleast {
		best = d.size
	}
	for _, sub := range d.subDirectories {
		best = maths.Min(best, sub.smallestDir(atleast))
	}
	return best
}

func (d *day07) Solve(lines []string, o command.Output) {

	root := &directory{"/", 0, map[string]*directory{}}
	var dirStack []*directory
	var prevDir *directory
	for i := 0; i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")
		if parts[0] == "$" {
			if parts[1] == "cd" {
				switch parts[2] {
				case "/":
					prevDir = root
					dirStack = []*directory{root}
				case "..":
					dirStack = dirStack[:len(dirStack)-1]
					prevDir = dirStack[len(dirStack)-1]
				default:
					if d, ok := prevDir.subDirectories[parts[2]]; ok {
						prevDir = d
						dirStack = append(dirStack, d)
					} else {
						d := &directory{parts[2], 0, map[string]*directory{}}
						prevDir.subDirectories[parts[2]] = d
						prevDir = d
						dirStack = append(dirStack, d)
					}
				}
			} else {
				// ls
				// do nothing
			}
		} else if parts[0] == "dir" {
		} else {
			// Otherwise, we are getting a file size
			prevDir.size += parse.Atoi(parts[0])
		}
	}
	root.addDirSizes()

	need := 30_000_000
	totalUnused := 70_000_000 - root.size

	// part1
	o.Stdoutln(root.smallDirSizeSum())
	// part2
	o.Stdoutln(root.smallestDir(need - totalUnused))
}

func (d *day07) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"95437",
				"24933642",
			},
		},
		{
			ExpectedOutput: []string{
				"1243729",
				"4443914",
			},
		},
	}
}
