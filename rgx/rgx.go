package rgx

import (
	"fmt"
	"regexp"
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
