package y2022

import (
	"regexp"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/point"
	"golang.org/x/exp/slices"
)

func Day15() aoc.Day {
	return &day15{}
}

type day15 struct{}

type sensor struct {
	location *point.Point[int]
	mhDist   int
}

// Returns the first
func (d *day15) boundedRangeLength(ranges [][]int, bound int) (int, bool) {
	for _, r := range ranges {
		r[0], r[1] = maths.Max(0, r[0]), maths.Max(0, r[1])
		r[0], r[1] = maths.Min(bound, r[0]), maths.Min(bound, r[1])
	}
	nRanges := d.simplifyRange(ranges)
	if len(nRanges) == 2 {
		return nRanges[0][1] + 1, true
	}
	return 0, false
}

func (d *day15) simplifyRange(ranges [][]int) [][]int {
	slices.SortFunc(ranges, func(this, that []int) bool {
		return this[0] < that[0]
	})

	if len(ranges) == 0 {
		return nil
	}

	nRanges := [][]int{ranges[0]}
	prev := ranges[0]
	for i := 1; i < len(ranges); i++ {
		r := ranges[i]
		if r[0] <= prev[1]+1 {
			prev[1] = maths.Max(r[1], prev[1])
		} else {
			nRanges = append(nRanges, r)
			prev = r
		}
	}

	return nRanges
}

func (d *day15) rangeLength(ranges [][]int) int {
	nRanges := d.simplifyRange(ranges)

	var count int
	for _, r := range nRanges {
		count += r[1] - r[0] + 1
	}
	return count
}

func (d *day15) Solve(lines []string, o command.Output) {
	r := regexp.MustCompile("Sensor at x=(-?[0-9]+), y=(-?[0-9]+): closest beacon is at x=(-?[0-9]+), y=(-?[0-9]+)")

	// bound := 20
	bound := parse.Atoi(lines[0]) * 2

	// Map from y to x to
	var sensors []*sensor
	beaconCounts := map[int]map[int]bool{}
	for _, line := range lines[1:] {
		match := r.FindStringSubmatch(line)

		s := point.New(parse.Atoi(match[1]), parse.Atoi(match[2]))
		beacon := point.New(parse.Atoi(match[3]), parse.Atoi(match[4]))
		maths.Insert(beaconCounts, beacon.Y, beacon.X, true)

		sensors = append(sensors, &sensor{s, s.ManhattanDistance(beacon)})
	}

	var part1, part2 int
	for y := 0; y <= bound; y++ {
		var rnge [][]int
		for _, sensor := range sensors {
			vertDist := sensor.mhDist - maths.Abs(sensor.location.Y-y)

			if vertDist < 0 {
				continue
			}

			left, right := sensor.location.X-vertDist, sensor.location.X+vertDist
			rnge = append(rnge, []int{left, right})
		}

		if y == bound/2 {
			part1 = d.rangeLength(rnge) - len(beaconCounts[y])
			if part2 != 0 {
				break
			}
		}
		if x, ok := d.boundedRangeLength(rnge, bound); ok {
			part2 = x*4000000 + y
			if part1 != 0 {
				break
			}
		}
	}
	o.Stdoutln(part1, part2)
}

func (d *day15) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"26 56000011",
			},
		},
		{
			ExpectedOutput: []string{
				"5508234 10457634860779",
			},
		},
	}
}
