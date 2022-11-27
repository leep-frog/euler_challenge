package eulerchallenge

import (
	"fmt"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

var (
	runLongTests = false
	// to keep import
	one = maths.One()
	// filter out tests
	timeLimit  = 10.0
	testFilter = func(cct *codingChallengeTest) bool {
		return true //&& cct.num == 88
	}
)

type codingChallengeTest struct {
	num      int
	name     string
	args     []string
	want     []string
	estimate float64
	skip     string

	elapsed float64
}

func (cct *codingChallengeTest) shouldSkip() (string, bool) {
	if cct.skip != "" {
		return cct.skip, true
	}
	if timeLimit != 0 && cct.estimate >= timeLimit {
		return fmt.Sprintf("Skipping due to test length (limit=%.2f, estimate=%.2f)", timeLimit, cct.estimate), true
	}
	return "", false
}

func TestAll(t *testing.T) {
	var tests []*codingChallengeTest
	for _, p := range getProblems() {
		for _, ex := range p.executions {
			tests = append(tests, &codingChallengeTest{
				p.num,
				fmt.Sprintf("Problem %d, args %v, estimate %.1f", p.num, ex.args, ex.estimate),
				append([]string{fmt.Sprintf("%d", p.num)}, ex.args...),
				[]string{ex.want},
				ex.estimate,
				ex.skip,
				0.0,
			})
		}
	}

	var totalEst float64
	for _, test := range tests {
		if _, skip := test.shouldSkip(); !skip && testFilter(test) {
			totalEst += test.estimate
		}
	}
	minEst, secEst := int(totalEst)/60, int(totalEst)%60
	t.Logf("Test estimate: %dm:%ds", minEst, secEst)

	for _, test := range tests {
		test.test(t)
	}
	sort.SliceStable(tests, func(i, j int) bool {
		return tests[i].elapsed > tests[j].elapsed
	})
	t.Logf("==================")
	t.Logf("Long tests:")
	for i := 1; i < maths.Min(5, len(tests)) && tests[i].elapsed > 5; i++ {
		test := tests[i]
		t.Logf("Test %q took %5.2f seconds", test.name, test.elapsed)
	}
}

func (ct *codingChallengeTest) test(t *testing.T) {
	if !testFilter(ct) {
		// Don't do t.Skip here because it just crowds the verbose test output.
		return
	}
	t.Run(ct.name, func(t *testing.T) {
		if msg, skip := ct.shouldSkip(); skip {
			t.Skipf("Skipping test: %s", msg)
		}

		start := time.Now()
		etc := &command.ExecuteTestCase{
			Node: command.AsNode(&command.BranchNode{
				Branches: Branches(),
			}),
			Args:          ct.args,
			WantStdout:    fmt.Sprintf("%s\n", strings.Join(ct.want, "\n")),
			SkipDataCheck: true,
		}
		command.ExecuteTest(t, etc)

		estimate := ct.estimate
		if estimate == 0 {
			estimate = 0.1
		} else if ct.estimate <= 0.1 {
			t.Fatalf("redundant estimate (default is 0.1)")
		}

		ct.elapsed = float64(time.Now().Sub(start).Microseconds()) / 1_000_000.0
		if ct.elapsed > 2*estimate {
			t.Logf("(Too long) Test took %5.2f seconds, expected %5.2f", ct.elapsed, estimate)
		}
		if estimate > 0.5 && ct.elapsed < 0.25*estimate {
			t.Logf("(Bad estimate) Test took %5.2f seconds, expected %5.2f", ct.elapsed, estimate)
		}
	})
}
