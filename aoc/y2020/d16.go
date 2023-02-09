package y2020

import (
	"regexp"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/functional"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day16() aoc.Day {
	return &day16{}
}

type day16 struct{}

type Section struct {
	name string
	rnge *maths.Range
}

func (s *Section) Matches(tickets [][]int, colIndex int) bool {
	return functional.All(tickets, func(ticket []int) bool {
		return s.rnge.Contains(ticket[colIndex])
	})
}

func (d *day16) Solve(lines []string, o command.Output) {

	sectionRegex := regexp.MustCompile("^([a-z ]*): (.*)$")
	var sections []*Section

	var yourTicket []int
	var allTickets [][]int
	var validRange *maths.Range

	// Construct objects
	var inputSection int
	for _, line := range lines {
		if line == "" {
			continue
		}

		switch inputSection {
		case 0:
			// Category rules section
			if line == "your ticket:" {
				inputSection++
				continue
			}
			m := sectionRegex.FindStringSubmatch(line)
			key := m[1]
			ranges := functional.Map(strings.Split(m[2], " or "), func(s string) *maths.Range {
				vs := strings.Split(s, "-")
				return maths.NewRange(parse.Atoi(vs[0]), parse.Atoi(vs[1]))
			})

			for i, r := range ranges {
				// Global range
				if validRange == nil {
					validRange = r
				} else {
					validRange = validRange.Merge(r)
				}

				// Per-section ranges
				if i == 0 {
					sections = append(sections, &Section{key, r})
				} else {
					sections[len(sections)-1].rnge = sections[len(sections)-1].rnge.Merge(r)
				}
			}
		case 1:
			// Your ticket section
			if line == "nearby tickets:" {
				inputSection++
				continue
			}
			yourTicket = parse.AtoiArray(strings.Split(line, ","))
		case 2:
			// Nearby tickets section
			allTickets = append(allTickets, parse.AtoiArray(strings.Split(line, ",")))
		}
	}

	// Part 1
	var errRate int
	validTickets := [][]int{
		yourTicket,
	}
	for _, t := range allTickets {
		for _, v := range t {
			if !validRange.Contains(v) {
				errRate += v
				goto INVALID_TICKET
			}
		}
		validTickets = append(validTickets, t)
	INVALID_TICKET:
	}

	var rcs []int
	for i := 0; i < len(validTickets[0]); i++ {
		rcs = append(rcs, i)
	}

	ms := maths.MatchItems(validTickets, sections, rcs)
	departureProd := 1
	for _, m := range ms {
		if strings.HasPrefix(m.Left.name, "departure") {
			departureProd *= yourTicket[m.Right]
		}
	}
	o.Stdoutln(errRate, departureProd)
}

func (d *day16) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"71 1",
			},
		},
		{
			ExpectedOutput: []string{
				"25059 3253972369789",
			},
		},
	}
}
