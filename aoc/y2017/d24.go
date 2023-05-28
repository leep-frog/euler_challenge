package y2017

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day24() aoc.Day {
	return &day24{}
}

type day24 struct{}

type port struct {
	id, in, out int
}

type portBuilder struct {
	cur     int
	portMap map[int][]*port
	inUse   map[int]bool
	score   int
}

func (p *port) score() int {
	return p.in + p.out
}

func (p *port) notValue(k int) int {
	if p.in == k {
		return p.out
	}
	return p.in
}

type bridge struct {
	size, score int
}

func (b *bridge) LT(d *bridge) bool {
	if b.size != d.size {
		return b.size < d.size
	}
	return b.score < d.score
}

func (pb *portBuilder) solve(bridgeSize int, best *maths.Bester[int, int], bridgeBest *maths.Bester[int, *bridge]) {
	for _, p := range pb.portMap[pb.cur] {
		if !pb.inUse[p.id] {
			old := pb.cur
			pb.cur = p.notValue(pb.cur)
			pb.inUse[p.id] = true
			pb.score += p.score()

			best.Check(pb.score)
			bridgeBest.Check(&bridge{bridgeSize, pb.score})
			pb.solve(bridgeSize+1, best, bridgeBest)

			pb.cur = old
			delete(pb.inUse, p.id)
			pb.score -= p.score()
		}
	}
}

func (d *day24) Solve(lines []string, o command.Output) {
	var ports []*port
	for i, parts := range parse.Split(lines, "/") {
		p := &port{i, parse.Atoi(parts[0]), parse.Atoi(parts[1])}
		ports = append(ports, p)
	}

	portMap := map[int][]*port{}
	for _, p := range ports {
		portMap[p.in] = append(portMap[p.in], p)
		portMap[p.out] = append(portMap[p.out], p)
	}
	pb := &portBuilder{0, portMap, map[int]bool{}, 0}

	largest := maths.Largest[int, int]()
	largestBridge := maths.LargestT[int, *bridge]()
	pb.solve(0, largest, largestBridge)
	o.Stdoutln(largest.Best(), largestBridge.Best().score)
}

func (d *day24) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"31 19",
			},
		},
		{
			ExpectedOutput: []string{
				"1940 1928",
			},
		},
	}
}
