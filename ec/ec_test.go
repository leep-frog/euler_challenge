package eulerchallenge

import (
	"fmt"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/leep-frog/command/commander"
	"github.com/leep-frog/command/commandertest"
	"github.com/leep-frog/command/commandtest"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/functional"
)

var (
	runLongTests = false
	// to keep import
	one = maths.One()
	// filter out tests
	timeLimit  = 100.0
	testFilter = func(cct *codingChallengeTest) bool {
		toCheck := []int{
			// Test numbers to check
			694, // CURRENT_PROBLEM

			// List of problems that use bfs package
			// 18, 60, 61, 81, 82, 83, 88, 96, 108, 109, 118, 119, 122, 127, 151, 152, 233, 243,
		}
		set := maths.NewSimpleSet(toCheck...)
		return true && (len(toCheck) == 0 || set[cct.num])
	}
)

type codingChallengeTest struct {
	num      int
	name     string
	args     []string
	want     []string
	estimate float64
	skip     string

	// The order
	exIdx int

	elapsed float64
}

func (cct *codingChallengeTest) shouldSkip() (string, bool) {
	if cct.skip != "" {
		return cct.skip, true
	}
	if timeLimit != 0 && cct.estimate > timeLimit {
		return fmt.Sprintf("Skipping due to test length (limit=%.2f, estimate=%.2f)", timeLimit, cct.estimate), true
	}
	return "", false
}

func TestAll(t *testing.T) {
	var tests []*codingChallengeTest
	for _, p := range getProblems() {
		for exIdx, ex := range p.Executions {
			tests = append(tests, &codingChallengeTest{
				p.Num,
				fmt.Sprintf("Problem %d, args %v, estimate %.1f", p.Num, ex.Args, ex.Estimate),
				append([]string{fmt.Sprintf("%d", p.Num)}, ex.Args...),
				[]string{ex.Want},
				ex.Estimate,
				ex.Skip,
				exIdx,
				0.0,
			})
		}
	}

	// Get total and per-test estimates
	var totalEst float64
	numToEst := map[int]float64{}
	for _, test := range tests {
		if _, skip := test.shouldSkip(); !skip && testFilter(test) {
			totalEst += test.estimate
			numToEst[test.num] += test.estimate
		}
	}
	minEst, secEst := int(totalEst)/60, int(totalEst)%60
	t.Logf("Test estimate: %dm:%ds", minEst, secEst)

	// Sort the tests by estimate and then by problem
	functional.SortFunc(tests, func(a, b *codingChallengeTest) bool {
		// If the same problem number, maintain order
		if a.num == b.num {
			return a.exIdx < b.exIdx
		}

		// Otherwise return the quicker problem
		aEst, bEst := numToEst[a.num], numToEst[b.num]
		return aEst < bEst
	})

	for _, test := range tests {
		// tmr := profiler.NewTimer()
		// tmr.Start()
		test.test(t)
		// tmr.End()
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
		etc := &commandtest.ExecuteTestCase{
			Node: &commander.BranchNode{
				Branches: Branches(),
			},
			Args:          ct.args,
			WantStdout:    fmt.Sprintf("%s\n", strings.Join(ct.want, "\n")),
			SkipDataCheck: true,
		}
		commandertest.ExecuteTest(t, etc)

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
