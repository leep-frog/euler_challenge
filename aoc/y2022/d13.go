package y2022

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/functional"
)

func Day13() aoc.Day {
	return &day13{}
}

type day13 struct{}

type packet struct {
	number      bool
	list        []*packet
	numberValue int
}

var (
	numberRegex = regexp.MustCompile("^[0-9]*$")
)

func (p *packet) toList() *packet {
	if !p.number {
		panic("INVALID")
	}
	return &packet{false, []*packet{p}, 0}
}

func (p *packet) cmp(q *packet) int {
	if p.number != q.number {
		if p.number {
			return p.toList().cmp(q)
		}
		return p.cmp(q.toList())
	}

	// p and q are both numbers, or are both lists
	if p.number {
		if p.numberValue < q.numberValue {
			return -1
		} else if p.numberValue > q.numberValue {
			return 1
		}
		return 0
	}

	// p and q are both lists
	for i, pv := range p.list {
		if i >= len(q.list) {
			return 1
		}
		if cmp := pv.cmp(q.list[i]); cmp != 0 {
			return cmp
		}
	}

	// Check if q has more elements
	if len(q.list) > len(p.list) {
		return -1
	}

	return 0
}

func (p *packet) depthString(depth int, output *[]string) {
	if p.number {
		*output = append(*output, fmt.Sprintf("%s%d", strings.Repeat("  ", depth), p.numberValue))
	} else {
		*output = append(*output, fmt.Sprintf("%s[", strings.Repeat("  ", depth)))
		for _, sub := range p.list {
			sub.depthString(depth+1, output)
		}
		*output = append(*output, fmt.Sprintf("%s]", strings.Repeat("  ", depth)))
	}
}

func (p *packet) String() string {
	var output []string
	p.depthString(0, &output)
	return strings.Join(output, "\n")
}

func (d *day13) parsePacket(q *maths.Queue[string]) *packet {
	c := q.Pop()
	if numberRegex.MatchString(c) {
		fullNumber := []string{c}
		for numberRegex.MatchString(q.Peek()) {
			fullNumber = append(fullNumber, q.Pop())
		}
		return &packet{true, nil, parse.Atoi(strings.Join(fullNumber, ""))}
	}

	// Otherwise, we must be starting an array
	if c != "[" {
		panic("AHAHAH")
	}

	var list []*packet

	for {
		switch q.Peek() {
		case ",":
			q.Pop()
		case "]":
			q.Pop()
			return &packet{false, list, 0}
			// This takes care of the '\[' and '[0-9]' cases
		default:
			list = append(list, d.parsePacket(q))
		}
	}
}

func (d *day13) parsePacketFromString(line string) *packet {
	return d.parsePacket(maths.NewQueue(strings.Split(line, "")...))
}

func (d *day13) Solve(lines []string, o command.Output) {
	// var pairs [][]*packet
	var sum int
	var packets []*packet
	for i := 0; i < len(lines); i += 3 {
		p1 := d.parsePacketFromString(lines[i])
		p2 := d.parsePacketFromString(lines[i+1])

		if p1.cmp(p2) <= 0 {
			sum += (i / 3) + 1
		}
		packets = append(packets, p1, p2)
	}

	base1 := d.parsePacketFromString("[[2]]")
	base2 := d.parsePacketFromString("[[6]]")
	packets = append(packets, base1, base2)
	functional.SortFunc(packets, func(this, that *packet) bool { return this.cmp(that) <= 0 })
	mult := 1
	for i, p := range packets {
		if p.cmp(base1) == 0 || p.cmp(base2) == 0 {
			mult *= i + 1
		}
	}
	o.Stdoutln(sum, mult)
}

func (d *day13) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"13 140",
			},
		},
		{
			ExpectedOutput: []string{
				"6395 24921",
			},
		},
	}
}
