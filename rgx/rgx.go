package rgx

import (
	"fmt"
	"regexp"

	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/functional"
)

type Rgx struct {
	r *regexp.Regexp
}

func New(pattern string) *Rgx {
	return &Rgx{regexp.MustCompile(pattern)}
}

func (r *Rgx) MustMatch(input string) []string {
	match, ok := r.Match(input)
	if !ok {
		panic(fmt.Sprintf("input %q did not match pattern %q", input, r.r.String()))
	}
	return match
}

func (r *Rgx) Match(input string) ([]string, bool) {
	match := r.r.FindStringSubmatch(input)
	if len(match) == 0 {
		return nil, false
	}
	return match[1:], true
}

func (r *Rgx) MatchInts(input string) []int {
	return functional.Map(r.MustMatch(input), func(s string) int {
		return parse.Atoi(s)
	})
}

func (r *Rgx) ReplaceAll(input, replaceStr string) string {
	return r.r.ReplaceAllString(input, replaceStr)
}
