package eulerchallenge

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P79() *problem {
	return fileInputNode(79, func(lines []string, o command.Output) {
		codes := map[string]bool{}
		for _, line := range lines {
			for i := 0; i < len(line); i++ {
				codes[line[i:i+1]] = true
			}
		}

		// Now make topology graph
		topology := map[string]map[string]bool{}
		for _, line := range lines {
			before := []string{}
			for i := 0; i < len(line); i++ {
				c := line[i : i+1]
				for _, b := range before {
					maths.Set(topology, b, c, true)
				}
				before = append(before, c)
			}
		}
		r, ok := p79(codes, topology)
		o.Stdoutln(strings.Join(r, ""), ok)
	})
}

func p79(codes map[string]bool, topology map[string]map[string]bool) ([]string, bool) {
	if len(codes) == 0 {
		return nil, true
	}
	for c := range codes {
		if len(topology[c]) == 0 {
			for _, m := range topology {
				delete(m, c)
			}
			delete(codes, c)
			v, ok := p79(codes, topology)
			return append(v, c), ok
		}
	}
	return nil, false
}
