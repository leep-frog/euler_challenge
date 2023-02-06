package y2015

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/maths"
	"golang.org/x/exp/slices"
)

func Day19() aoc.Day {
	return &day19{}
}

type day19 struct{}

type molecule struct {
	m string
}

func (m *molecule) Code(mc *moleculeContext, path bfs.Path[*molecule]) string {
	return m.m
}

var (
	dpth = 0
)

func (m *molecule) Done(mc *moleculeContext, path bfs.Path[*molecule]) bool {
	if len(m.m) > 125 && dpth == 0 {
		dpth = 1
	}
	if dpth >= 1 && path.Len() > dpth {
		fmt.Println("DEPTH", dpth, path.Len())
		dpth = path.Len()
	}
	if mc.valid[m.m] {
		if _, ok := mc.got[m.m]; !ok {
			mc.got[m.m] = path.Len()
		}
	}
	return false
}

func (m *molecule) AdjacentStates(mc *moleculeContext, path bfs.Path[*molecule]) []*molecule {
	var r []*molecule
	for t := range mc.transformations(m.m) {
		r = append(r, &molecule{t})
	}
	return r
}

type opt struct {
	m     string
	depth int
}

func (o *opt) String() string {
	return o.m
}

func (d *day19) analyze(opts [][]*opt, depth, idx int, parts []string, mc *moleculeContext) {
	if idx == len(opts) {
		s := strings.Join(parts, "")
		// if len(s) > 125 {
		fmt.Println("CHECKING", s, mc.valid)
		// }
		bfs.ContextDistanceSearch[*moleculeContext, string, bfs.Int](mc, []*molecule2{{s, 0, 0}})
		if len(s) > 125 {
			fmt.Println("Done CHECKING", s)
		}
		return
	}

	var prefix string
	if idx != 0 {
		if idx%2 == 0 {
			prefix = Ar
		} else {
			prefix = Rn

		}
	}
	for _, o := range opts[idx] {
		d.analyze(opts, depth+o.depth, idx+1, append(parts, prefix+o.m), mc)
	}
}

var (
	Rn = "6"
	Ar = "7"
)

func (d *day19) final(dth int, mol string, valid, validMiddles map[string]bool, mc *moleculeContext, middleToLeft map[string]map[string]bool) []*opt {
	pre := strings.Repeat("| ", dth)
	fmt.Println(pre, "------------ FINAL:", mol)
	var parts []string
	var prevIdx, rDepth int
	for i := 0; i < len(mol); i++ {
		if string(mol[i]) == Rn {
			rDepth++
			if rDepth == 1 {
				parts = append(parts, mol[prevIdx:i])
				prevIdx = i + 1
			}
			// } else if mol[i] == 'A' && i < len(mol)-1 && mol[i+1] == 'r' {
		} else if string(mol[i]) == Ar {
			rDepth--
			if rDepth == 0 {
				parts = append(parts, mol[prevIdx:i])
				prevIdx = i + 1
			}
		}
	}
	parts = append(parts, mol[prevIdx:])

	var opts [][]*opt
	for i, p := range parts {
		if i%2 == 0 {
			opts = append(opts, []*opt{{p, 0}})
		} else {
			opts = append(opts, d.final(dth+1, p, validMiddles, validMiddles, mc, middleToLeft))
		}
	}

	// Now determine what the odd ones can be
	/*if len(parts) > 0 {
		for i := 0; i < len(parts)-1; i += 2 {
			leftValids := map[string]bool{}
			for _, o := range opts[i+1] {
				for k := range middleToLeft[o.m] {
					leftValids[k] = true
				}
			}
			opts[i] = d.final(dth+1, parts[i], leftValids, validMiddles, mc, middleToLeft)
		}
	}*/

	nmc := &moleculeContext{
		mc.ops,
		mc.maxLen,
		valid,
		map[string]int{},
		mc.prefixes,
	}

	if len(mol) > 125 {
		fmt.Println(opts)
	}
	d.analyze(opts, 0, 0, nil, nmc)
	fmt.Println(pre, "RES", nmc.got)
	if len(nmc.got) == 0 {
		panic(fmt.Sprintf("NOPE: %s", mol))
	}

	var r []*opt
	for s, d := range nmc.got {
		r = append(r, &opt{s, d})
	}
	return r
}

func (d *day19) donzo(depth int, mol string, transformations [][]string, best *maths.Bester[int, int]) {
	if mol == "e" {
		fmt.Println("YEE", depth)
	}
	if best.Set() && depth >= best.Best() {
		return
	}

	if mol == "e" {
		if best.Check(depth) {
			fmt.Println("BEST", best.Best())
		}
		fmt.Println("BESTT", depth)
		return
	}

	for idx, t := range transformations {
		if depth == 0 {
			fmt.Println("IDX", idx)
		}
		parts := strings.Split(mol, t[0])
		for i := 1; i < len(parts); i++ {
			d.donzo(depth+1, strings.Join(parts[:i], t[0])+t[1]+strings.Join(parts[i:], t[0]), transformations, best)
		}
	}
}

func (d *day19) rec2(oneof map[string]bool, mc *moleculeContext, mol string, depth int, cache map[string]int, best *maths.Bester[int, int]) int {
	// Solve in sections of `Rn[^rRn]*r`
	return 0
}

func (d *day19) rec(mc *moleculeContext, mol string, depth int, cache map[string]int, best *maths.Bester[int, int]) int {
	// if best.Set() && depth >= best.Best() {
	// 	return
	// }

	if v, ok := cache[mol]; ok {
		return v
	}

	if mol == "e" {
		best.Check(depth)
		fmt.Println("HURRAY", depth)
		cache[mol] = 0
		return 0
	}

	// if v, ok := cache[mol]; ok {
	// 	betterOne := maths.Min(v[0], depth)
	// 	best.Check(betterOne + v[1])
	// 	cache[mol] = []int{betterOne, v[1]}
	// 	return
	// }

	// bestAfter := -1
	shortest := maths.Smallest[int, int]()
	for newMol := range mc.transformations(mol) {
		shortest.Check(d.rec(mc, newMol, depth+1, cache, best))
	}
	cache[mol] = shortest.Best() + 1
	return shortest.Best() + 1
	// cache[mol] = []int{}
}

/*func (d *day19) solve2(depth int, mol string, mc *moleculeContext, leftOfRn map[string]bool, rightOfRn map[string]bool, best *maths.Bester[int, int]) {
	parts := strings.SplitN(mol, "Rn", 2)

	if len(parts) == 1 {
		_, dist := bfs.ContextSearch[*moleculeContext, string](mc, []*molecule{&molecule{mol, 0}})
		best.Check(dist)
		return
	}

	leftOps
	for opt, dist := range mc.Explore(mol, leftOfRn) {

	}
}*/

func (d *day19) makeReplacements(line string) string {
	reps := map[string]string{
		"Ti": "1",
		"Ca": "2",
		"Si": "3",
		"Mg": "4",
		"Al": "5",
		"Rn": "6",
		"Ar": "7",
		"Th": "8",
	}
	for k, v := range reps {
		line = strings.ReplaceAll(line, k, v)
	}
	return line
}

func (d *day19) Solve(lines []string, o command.Output) {
	fmt.Println("START")
	pattern := regexp.MustCompile("^(.*)6(.*)7$")
	ops := map[string][]string{}
	revOps := map[string][]string{}
	validMiddles := map[string]bool{}
	prefixes := map[string]bool{}
	// Map from middle value to options for left
	middleToLeft := map[string]map[string]bool{}
	var maxLen, revMaxLen int
	for _, line := range lines[:len(lines)-2] {
		line = d.makeReplacements(line)
		parts := strings.Split(line, " => ")
		ops[parts[0]] = append(ops[parts[0]], parts[1])
		revOps[parts[1]] = append(revOps[parts[1]], parts[0])
		maxLen = maths.Max(maxLen, len(parts[0]))
		revMaxLen = maths.Max(revMaxLen, len(parts[1]))
		m := pattern.FindStringSubmatch(parts[1])
		if len(m) > 0 {
			validMiddles[m[2]] = true
			if middleToLeft[m[2]] == nil {
				middleToLeft[m[2]] = map[string]bool{}
			}
			middleToLeft[m[2]][m[1]] = true
		}
		for i := 1; i <= len(parts[1]); i++ {
			prefixes[parts[1][:i]] = true
		}
	}
	mol := d.makeReplacements(lines[len(lines)-1])
	mc := &moleculeContext{ops, maxLen, nil, nil, prefixes}
	revMC := &moleculeContext{revOps, revMaxLen, map[string]bool{}, map[string]int{}, prefixes}
	_ = revMC
	part1 := len(mc.transformations(mol))
	fmt.Println("UNO", part1)

	fmt.Println()
	fmt.Println("PART 2 ------")
	fmt.Println(prefixes)

	/*	d.rec(revMC, mol, 0, map[string]int{}, maths.Smallest[int, int]())
		return

		// d.donzo(0, mol, ts)

		m4 := &molecule4{mol, 0, 0, strings.Count(mol, Rn)}
		q, w := bfs.ContextDistanceSearch[*moleculeContext, string, bfs.Int](revMC, []*molecule4{m4})
		fmt.Println(q, w)
		return
		// return
		/*path, dist := bfs.ContextDistanceSearch[*moleculeContext, string, bfs.Int, *molecule2](revMC, []*molecule2{{mol, 0, 0}})
		fmt.Println(path, dist)
		fmt.Println(path[len(path)-1].depth)
		return*/

	/*leftOfRnArr := map[string][]string{}
	leftOfRnBool := map[string]bool{}
	for conv := range mc.ops {
		if parts := strings.Split(conv, "Rn"); len(parts) > 1 {
			leftOfRnArr[parts[0]] = append(leftOfRnArr[parts[0]], parts[1])
			leftOfRnBool[parts[0]] = true
		}
	}

	best := maths.Smallest[int, int]()
	d.solve2(0, mol, revMC, leftOfRnBool, best)
	fmt.Println(best)*/

	var ts [][]string
	for from, tos := range revMC.ops {
		for _, to := range tos {
			ts = append(ts, []string{from, to})
		}
	}
	slices.SortFunc(ts, func(this, that []string) bool {
		thisDist := len(this[0]) - len(this[1])
		thatDist := len(that[0]) - len(that[1])
		return thisDist > thatDist
	})
	fmt.Println(middleToLeft)

	d.donzo(0, mol, ts, maths.Smallest[int, int]())
	return
}

type moleculeContext struct {
	ops      map[string][]string
	maxLen   int
	valid    map[string]bool
	got      map[string]int
	prefixes map[string]bool
}

func (mc *moleculeContext) Explore(molecule string, valid map[string]bool) map[string]int {
	r := map[string]int{}
	mc.explore(molecule, valid, r, 0)
	return r
}

func (mc *moleculeContext) explore(molecule string, valid map[string]bool, values map[string]int, depth int) {
	if valid[molecule] {
		if v, ok := values[molecule]; ok {
			values[molecule] = maths.Min(v, depth)
		} else {
			values[molecule] = depth
		}
	}

	for t := range mc.transformations(molecule) {
		mc.explore(t, valid, values, depth+1)
	}
}

func (mc *moleculeContext) transformations(molecule string) map[string]bool {
	transformations := map[string]bool{}
	for i := range molecule {
		for size := 1; size <= mc.maxLen; size++ {
			if i+size > len(molecule) {
				break
			}
			for _, v := range mc.ops[molecule[i:i+size]] {
				newMol := molecule[:i] + v
				if i+size < len(molecule) {
					newMol += molecule[i+size:]
				}
				transformations[newMol] = true
			}
		}
	}
	return transformations
}

func (d *day19) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"",
			},
		},
		{
			ExpectedOutput: []string{
				"",
			},
		},
	}
}

type molecule2 struct {
	m     string
	idx   int
	depth int
}

func (m *molecule2) Code(mc *moleculeContext) string {
	return m.String()
}

func (m *molecule2) String() string {
	return fmt.Sprintf("%s %d", m.m, m.idx)
}

func (m *molecule2) Distance(mc *moleculeContext) bfs.Int {
	return bfs.Int(m.depth)
}

var (
	dtt = 0
)

func (m *molecule2) Done(mc *moleculeContext) bool {
	if len(m.m) > 125 && dpth == 0 {
		dpth = 1
	}
	if dpth >= 1 && m.depth > dpth {
		fmt.Println(time.Now(), "DEPTH", dpth, m.depth)
		dpth = m.depth
	}
	if mc.valid[m.m] {
		if _, ok := mc.got[m.m]; !ok {
			mc.got[m.m] = m.depth
		}
	}
	return false
	/*if m.depth > dtt {
		dtt = m.depth
		fmt.Println(time.Now(), "DEPTH", dtt)
	}
	return m.m == "e"*/
}

func (m *molecule2) AdjacentStates(mc *moleculeContext) []*molecule2 {
	if m.idx > len(m.m) {
		return nil
	}

	r := []*molecule2{
		{m.m, m.idx + 1, m.depth},
	}

	for offset := 1; offset <= mc.maxLen && m.idx-offset >= 0; offset++ {
		before, mid, after := m.m[:m.idx-offset], m.m[m.idx-offset:m.idx], m.m[m.idx:]

		validPrefix := len(before) == 0
		for i := 1; len(before)-i > 0; i++ {
			if mc.prefixes[before[len(before)-i:]] {
				validPrefix = true
				break
			}
		}

		solnPrefixes := map[string]bool{}
		for v := range mc.valid {
			for i := 1; i < len(v); i++ {
				solnPrefixes[v[:i]] = true
			}
		}
		validPrefix = validPrefix || solnPrefixes[before]

		if !validPrefix {
			// fmt.Println("INVALID", before, mid, after)
			continue
		}

		for _, to := range mc.ops[mid] {
			r = append(r, &molecule2{before + to + after, m.idx - offset + len(to), m.depth + 1})
		}
	}
	return r
}

type molecule3 struct {
	m     string
	idx   int
	depth int
}

func (m *molecule3) Code(mc *moleculeContext) string {
	return m.String()
}

func (m *molecule3) String() string {
	return fmt.Sprintf("%s %d", m.m, m.idx)
}

func (m *molecule3) Distance(mc *moleculeContext) bfs.Int {
	return bfs.Int(len(m.m))
	// return bfs.Int(m.depth)
}

var (
	dtt3 = 0
)

func (m *molecule3) Done(mc *moleculeContext) bool {
	if len(m.m) > 125 && dtt3 == 0 {
		dtt3 = 1
	}
	if dtt3 >= 1 && m.depth > dtt3 {
		fmt.Println(time.Now(), "DEPTH", dtt3, m.depth)
		dtt3 = m.depth
	}
	if mc.valid[m.m] {
		if _, ok := mc.got[m.m]; !ok {
			mc.got[m.m] = m.depth
		}
	}
	return m.m == "e"
}

func (m *molecule3) AdjacentStates(mc *moleculeContext) []*molecule3 {
	if m.idx > len(m.m) {
		return nil
	}

	r := []*molecule3{
		{m.m, m.idx + 1, m.depth},
	}

	for offset := 1; offset <= mc.maxLen && m.idx-offset >= 0; offset++ {
		before, mid, after := m.m[:m.idx-offset], m.m[m.idx-offset:m.idx], m.m[m.idx:]

		validPrefix := len(before) == 0
		for i := 1; len(before)-i > 0; i++ {
			if mc.prefixes[before[len(before)-i:]] {
				validPrefix = true
				break
			}
		}

		solnPrefixes := map[string]bool{}
		for v := range mc.valid {
			for i := 1; i < len(v); i++ {
				solnPrefixes[v[:i]] = true
			}
		}
		validPrefix = validPrefix || solnPrefixes[before]

		if !validPrefix {
			// fmt.Println("INVALID", before, mid, after)
			continue
		}

		for _, to := range mc.ops[mid] {
			r = append(r, &molecule3{before + to + after, m.idx - offset + len(to), m.depth + 1})
		}
	}
	return r
}

type molecule4 struct {
	m       string
	idx     int
	depth   int
	rnCount int
}

func (m *molecule4) Code(mc *moleculeContext) string {
	return m.String()
}

func (m *molecule4) String() string {
	return ""
	// return fmt.Sprintf("%s %d", m.m, m.idx)
}

var (
	idk = 0
)

func (m *molecule4) Distance(mc *moleculeContext) bfs.Int {
	// return bfs.Int(m.depth + len(m.m)/2 - (6 * m.rnCount))
	// return bfs.Int(len(m.m))
	idk--
	return bfs.Int(idk)
	// return bfs.Int(m.depth + len(m.m)/2 - (6 * m.rnCount))
}

var (
	dtt4 = 0
)

func (m *molecule4) Done(mc *moleculeContext) bool {
	if len(m.m) > 125 && dtt4 == 0 {
		dtt4 = 1
	}
	if dtt4 >= 1 && m.depth > dtt4 {
		fmt.Println(time.Now(), "DEPTH", dtt4, m.depth)
		dtt4 = m.depth
	}
	if mc.valid[m.m] {
		if _, ok := mc.got[m.m]; !ok {
			mc.got[m.m] = m.depth
		}
	}
	return m.m == "e"
}

func (m *molecule4) AdjacentStates(mc *moleculeContext) []*molecule4 {
	if m.idx > len(m.m) {
		return nil
	}

	r := []*molecule4{
		{m.m, m.idx + 1, m.depth, m.rnCount},
	}

	for offset := 1; offset <= mc.maxLen && m.idx-offset >= 0; offset++ {
		before, mid, after := m.m[:m.idx-offset], m.m[m.idx-offset:m.idx], m.m[m.idx:]

		validPrefix := len(before) == 0
		for i := 1; len(before)-i > 0; i++ {
			if mc.prefixes[before[len(before)-i:]] {
				validPrefix = true
				break
			}
		}

		solnPrefixes := map[string]bool{}
		for v := range mc.valid {
			for i := 1; i < len(v); i++ {
				solnPrefixes[v[:i]] = true
			}
		}
		validPrefix = validPrefix || solnPrefixes[before]

		if !validPrefix {
			// fmt.Println("INVALID", before, mid, after)
			continue
		}

		for _, to := range mc.ops[mid] {
			rnc := m.rnCount
			if strings.Contains(mid, Rn) {
				rnc--
			}
			r = append(r, &molecule4{before + to + after, m.idx - offset + len(to), m.depth + 1, rnc})
		}
	}
	return r
}
