package p930

import (
	"fmt"
	"time"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/combinatorics"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/equations"
	"github.com/leep-frog/euler_challenge/fraction"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/profiler"
)

var (
	pf = profiler.New()
)

func P930() *ecmodels.Problem {
	return ecmodels.IntsInputNode(930, 2, 0, func(o command.Output, ns []int) {
		n, m := ns[0], ns[1]

		/*var sum float64
		for a := 2; a <= n; a++ {
			for b := 2; b <= m; b++ {
				sum += solveSimple(a, b)
			}
		}
		fmt.Println("SUM FOR", n, m, sum)
		return*/

		solve(n, m)
		// solveG(n, m)
		// fmt.Println(reduce(fr(123456789012345678, 98765432179876543)))
		fmt.Println("EQ.PROFILER")
		fmt.Println(equations.Profiler)
		return

		/*
			evMap := map[string]*fraction.Rational{}

			countsMap := map[string]int{}
			calcProbs(n, m, m, make([]int, n), countsMap)

			var totalCount int
			for _, v := range countsMap {
				totalCount += v
			}

			probsMap := map[string]*fraction.Rational{}
			for k, v := range countsMap {
				probsMap[k] = fr(v, totalCount)
			}
			fmt.Println(probsMap)

			// func
			// return

			cur := make([]int, n)
			cur[0] = m

			ba := &ballArrangement{cur, n, m, false}
			ba.bruteSpread(0, fr(1, 1), evMap)

			// fmt.Println(evMap)
			// for k,
			summation := zero
			for k := range evMap {
				evMap[k] = evMap[k].Times(probsMap[k])
				summation = summation.Plus(evMap[k])
			}
			fmt.Println(evMap)
			fmt.Println(summation)

			// o.Stdoutln(n)

			return* /

			/*arrs := allArrangements(n, n, m, m, m, nil)
			fmt.Println(arrs)

			for _, arr := range arrs {
				fmt.Println(arr, arr.solve(evMap))
			}* /
			allArrangementsThree(n, m, m, false, cur, evMap)

			for k, v := range evMap {
				fmt.Println("HERE WE GO")
				fmt.Println(k, v)
			}
			fmt.Println(evMap)*/
	}, []*ecmodels.Execution{
		{
			Args: []string{"1"},
			Want: "",
		},
		{
			Args: []string{"2"},
			Want: "",
		},
	})
}

var (
	cache = map[string]bool{}
)

func F(n, m int) {

}

func allArrangements(n, initN, initM, m, max int, cur []int) []*ballArrangement {
	if n == 0 {
		if m == 0 {
			return []*ballArrangement{{
				bread.Copy(cur),
				initN,
				initM,
				cur[0] == initM,
			}}
		}
		return nil
	}

	var r []*ballArrangement
	for v := max; v >= 0; v-- {

		nextMax := max
		if len(cur) == 0 {
			nextMax = v
		}

		r = append(r, allArrangements(n-1, initN, initM, m-v, nextMax, append(cur, v))...)
	}
	return r
}

func allArrangementsTwo(n, m int, cur []int) []*ballArrangement {
	if n == 0 {
		if m == 0 {
			// return []*ballArrangement{{bread.Copy(cur)}}
			return nil
		}
		return nil
	}

	var r []*ballArrangement
	for v := 0; v <= m; v++ {
		r = append(r, allArrangementsTwo(n-1, m-v, append(cur, v))...)
	}
	return r
}

type ballArrangement struct {
	// Ordered by biggest ball first, and adjacent number is second biggest
	balls  []int
	n      int
	m      int
	solved bool
}

var (
	solvedArrangements []string
)

func allArrangementsThree(n, initM, m int, solved bool, cur []int, evMap map[string]*fraction.Rational) {
	if m == 0 {
		ba := &ballArrangement{cur, n, initM, solved}
		fmt.Println("SOLVING", ba)
		// ba.solve("", evMap)
		fmt.Println(ba)
		// ba.alreadySolved()
		return
	}

	for i := 0; i < n; i++ {
		cur[i]++
		allArrangementsThree(n, initM, m-1, solved || cur[i] == initM, cur, evMap)
		cur[i]--
	}
}

func (ba *ballArrangement) String() string {
	// return fmt.Sprintf("%v %v", ba.balls, ba.solved)
	return fmt.Sprintf("%v", ba.balls)
}

var (
	half  = fr(1, 2)
	halfE = vr(1, 2, 0, 1)
	zero  = fr(0, 1)
)

func (ba *ballArrangement) alreadySolved(evMap map[string]*variableRational) (*variableRational, bool) {
	shifted, revShifted := make([]int, ba.n), make([]int, ba.n)
	for shift := 0; shift < ba.n; shift++ {
		for idx := 0; idx < ba.n; idx++ {
			shifted[(idx+shift)%ba.n] = ba.balls[idx]
			revShifted[(ba.n-1-idx+shift)%ba.n] = ba.balls[idx]
		}

		code := fmt.Sprintf("%v", shifted)
		revCode := fmt.Sprintf("%v", revShifted)
		if v, ok := evMap[code]; ok {
			return v, true
		}
		if v, ok := evMap[revCode]; ok {
			return v, true
		}
	}
	return nil, false
}

func (ba *ballArrangement) solve(depth string, evMap map[string]*variableRational) *variableRational {

	if len(depth) > 10 {
		panic("UGH")
	}
	code := ba.String()
	if v, ok := ba.alreadySolved(evMap); ok {
		return v
	}

	if ba.solved {
		evMap[code] = vr(0, 1, 0, 1)
		return evMap[code]
	}

	ev := vr(0, 1, 0, 1)
	for i, v := range ba.balls {
		if v == 0 {
			continue
		}

		oddsOfMove := vr(v, ba.m, 0, 1)

		leftIdx := (i + ba.n - 1) % ba.n
		rightIdx := (i + 1) % ba.n

		// Move to the left
		ba.balls[i]--
		ba.balls[leftIdx]++
		fmt.Println("CHECKING  LEFT", depth, ba)
		if ba.balls[leftIdx] == ba.m {
			ba.solved = true
		}
		leftEV := oddsOfMove.Times(ba.solve(depth+"  ", evMap).Plus(vr(1, 1, 0, 1)))
		ba.solved = false
		ba.balls[i]++
		ba.balls[leftIdx]--

		// Move to the left
		ba.balls[i]--
		ba.balls[rightIdx]++
		fmt.Println("CHECKING RIGHT", depth, ba)
		if ba.balls[rightIdx] == ba.m {
			ba.solved = true
		}
		rightEV := oddsOfMove.Times(ba.solve(depth+"  ", evMap).Plus(vr(1, 1, 0, 1)))
		ba.solved = false
		ba.balls[i]++
		ba.balls[rightIdx]--

		ev = ev.Plus(leftEV.Plus(rightEV).Times(halfE))
	}

	evMap[code] = ev
	return ev
}

func fr(num, den int) *fraction.Rational {
	return fraction.NewRational(num, den)
}

func vr(num, den, enum, eden int) *variableRational {
	return &variableRational{
		fr(num, den),
		fr(enum, eden),
	}
}

type variableRational struct {
	c        *fraction.Rational
	variable *fraction.Rational
}

func (vr *variableRational) Times(that *variableRational) *variableRational {
	if vr.variable.NEQ(zero) && that.variable.NEQ(zero) {
		panic("TWO VAR MULTIPLYING")
	}

	return &variableRational{
		vr.c.Times(that.c),
		vr.c.Times(that.variable).Plus(vr.variable.Times(that.c)),
	}
}

func (vr *variableRational) Plus(that *variableRational) *variableRational {
	return &variableRational{
		vr.c.Plus(that.c),
		vr.variable.Plus(that.variable),
	}
}

var (
	stop = fr(1, maths.Pow(2, 30))
)

func (ba *ballArrangement) bruteSpread(steps int, prob *fraction.Rational, evMap map[string]*fraction.Rational) {
	if prob.LT(stop) {
		return
	}

	code := ba.smartCode()
	curV, ok := evMap[code]
	if !ok {
		curV = fr(0, 1)
	}

	evMap[code] = curV.Plus(prob.Times(fr(steps, 1)))

	for i, v := range ba.balls {
		oddsOfMove := fr(v, ba.m)
		leftIdx := (i + ba.n - 1) % ba.n
		rightIdx := (i + 1) % ba.n

		// Move to the left
		ba.balls[i]--
		ba.balls[leftIdx]++
		if ba.balls[leftIdx] != ba.m {
			ba.bruteSpread(steps+1, prob.Times(oddsOfMove).Times(half), evMap)
		}
		ba.balls[i]++
		ba.balls[leftIdx]--

		// Move to the right
		ba.balls[i]--
		ba.balls[rightIdx]++
		if ba.balls[rightIdx] != ba.m {
			ba.bruteSpread(steps+1, prob.Times(oddsOfMove).Times(half), evMap)
		}
		ba.balls[i]++
		ba.balls[rightIdx]--
	}
}

var (
	smartCodeAliases = map[string]string{}
)

func (ba *ballArrangement) smartCode() string {

	baseCode := fmt.Sprintf("%v", ba.balls)
	if v, ok := smartCodeAliases[baseCode]; ok {
		return v
	}

	var best string
	var codeAliases []string
	shifted, revShifted := make([]int, ba.n), make([]int, ba.n)
	for shift := 0; shift < ba.n; shift++ {
		for idx := 0; idx < ba.n; idx++ {
			shifted[(idx+shift)%ba.n] = ba.balls[idx]
			revShifted[(ba.n-1-idx+shift)%ba.n] = ba.balls[idx]
		}

		code := fmt.Sprintf("%v", shifted)
		revCode := fmt.Sprintf("%v", revShifted)
		codeAliases = append(codeAliases, code, revCode)
		if code > best {
			best = code
		}
		if revCode > best {
			best = revCode
		}
	}

	for _, codeAlias := range codeAliases {
		smartCodeAliases[codeAlias] = best
	}
	return best
}

func createEquation(ba *ballArrangement) *equations.Equation {
	// c = 1

	// fmt.Println("CREATING EQUATION FOR", ba)

	m := map[equations.Variable]*fraction.Rational{
		equations.Variable(ba.smartCode()): fr(-1, 1),
	}

	if ba.solved {
		// fmt.Println("EQ", m)
		return equations.NewEq(m, fr(0, 1))
	}

	for i, v := range ba.balls {
		if v == 0 {
			continue
		}

		oddsOfMove := fr(v, ba.m).Times(half)

		leftIdx := (i + ba.n - 1) % ba.n
		rightIdx := (i + 1) % ba.n

		// Move to the left
		ba.balls[i]--
		ba.balls[leftIdx]++
		leftCode := ba.smartCode()
		ba.balls[i]++
		ba.balls[leftIdx]--

		// Move to the left
		ba.balls[i]--
		ba.balls[rightIdx]++
		rightCode := ba.smartCode()
		ba.balls[i]++
		ba.balls[rightIdx]--

		if v, ok := m[equations.Variable(leftCode)]; ok {
			m[equations.Variable(leftCode)] = v.Plus(oddsOfMove)
		} else {
			m[equations.Variable(leftCode)] = oddsOfMove.Copy()
		}

		if v, ok := m[equations.Variable(rightCode)]; ok {
			m[equations.Variable(rightCode)] = v.Plus(oddsOfMove)
		} else {
			m[equations.Variable(rightCode)] = oddsOfMove.Copy()
		}
	}

	// fmt.Println("EQ", m)

	return equations.NewEq(m, fr(1, 1))
	// vs := equations.NewVs()
}

func calcProbs(n, initM, m int, solved bool, cur []int, countsMap map[string]int, bas *[]*ballArrangement) {
	if m == 0 {
		ba := &ballArrangement{
			balls:  bread.Copy(cur),
			n:      n,
			m:      initM,
			solved: solved,
		}

		if _, ok := countsMap[ba.smartCode()]; !ok {
			*bas = append(*bas, ba)
		}
		countsMap[ba.smartCode()]++
		return
	}

	for spot := 0; spot < n; spot++ {
		cur[spot]++
		calcProbs(n, initM, m-1, solved || cur[spot] == initM, cur, countsMap, bas)
		cur[spot]--
	}
}

func findAllArrangements(ba *ballArrangement, countsMap map[string]*fraction.Rational, bas *[]*ballArrangement) {
	if _, ok := countsMap[ba.smartCode()]; ok {
		return
	}

	var solved bool
	for _, b := range ba.balls {
		if b == ba.m {
			solved = true
		}
	}

	bac := &ballArrangement{
		balls:  bread.Copy(ba.balls),
		n:      ba.n,
		m:      ba.m,
		solved: solved,
	}

	*bas = append(*bas, bac)

	countsMap[ba.smartCode()] = ba.calcInitProb()

	for i, v := range ba.balls {
		if v == 0 {
			continue
		}

		leftIdx := (i + ba.n - 1) % ba.n
		rightIdx := (i + 1) % ba.n

		// Move to the left
		ba.balls[i]--
		ba.balls[leftIdx]++
		findAllArrangements(ba, countsMap, bas)
		ba.balls[i]++
		ba.balls[leftIdx]--

		// Move to the right
		ba.balls[i]--
		ba.balls[rightIdx]++
		findAllArrangements(ba, countsMap, bas)
		ba.balls[i]++
		ba.balls[rightIdx]--
	}
}

func (ba *ballArrangement) calcInitProb() *fraction.Rational {

	oddsForSingleOrder := fraction.NewBigRationalFromInt(maths.One().Int(), maths.BigPow(ba.n, ba.m).Int())

	var counts []int
	for _, cnt := range ba.balls {
		if cnt > 0 {
			counts = append(counts, cnt)
		}
	}

	ballPlacementOrderingCount := fraction.NewBigRationalFromInt(combinatorics.PermutationFromCount(counts).Int(), maths.One().Int())

	uniqueCodes := map[string]bool{}
	shifted, revShifted := make([]int, ba.n), make([]int, ba.n)
	for shift := 0; shift < ba.n; shift++ {
		for idx := 0; idx < ba.n; idx++ {
			shifted[(idx+shift)%ba.n] = ba.balls[idx]
			revShifted[(ba.n-1-idx+shift)%ba.n] = ba.balls[idx]
		}

		code := fmt.Sprintf("%v", shifted)
		revCode := fmt.Sprintf("%v", revShifted)
		uniqueCodes[code] = true
		uniqueCodes[revCode] = true
	}

	return ballPlacementOrderingCount.Times(oddsForSingleOrder).Times(fr(len(uniqueCodes), 1))
}

/*
(3, 2)

E1: [2 0 0] P() = 1/3
	E1 = 0
E2: [1 1 0] P() = 2/3
	E2 = 1 + 1/2 * E2 + 1/2 * E1
	0 = 1 + 1/2 * E1 - 1/2 * E2

*/

func solveG(N, M int) {
	sum := fr(0, 1)
	for n := 2; n <= N; n++ {
		for m := 2; m <= M; m++ {
			sum = sum.Plus(solve(n, m))
		}
	}
	fmt.Println(sum, "<== SUM FRACTION")
	fmt.Println("SUM FLOAT", sum.Float64())
}

var (
	solveCche = map[int]map[int]*fraction.Rational{
		2:  {},
		3:  {},
		4:  {},
		5:  {},
		6:  {},
		7:  {},
		8:  {},
		9:  {},
		10: {},
		11: {},
		12: {},
	}
)

func solve(n, m int) *fraction.Rational {

	// fmt.Println("==================== SOLVING FOR", n, m)
	uniqueArrs := []*ballArrangement{}

	balls := make([]int, n)
	balls[0] = m
	init := &ballArrangement{
		balls: balls,
		n:     n,
		m:     m,
	}
	countsMapTwo := map[string]*fraction.Rational{}

	pf.Start("calcProbs")

	findAllArrangements(init, countsMapTwo, &uniqueArrs)

	fmt.Println("Number of arrangements", len(uniqueArrs))

	pf.Start("CreateEqs")
	fmt.Println("  CREATE EQS", time.Now())
	var eqs []*equations.Equation
	for _, ba := range uniqueArrs {
		eqs = append(eqs, createEquation(ba))
	}

	// equations.SolveLin(eqs)
	/*fmt.Println("\n\nFOR FUN SOLVE")
	oSoln := equations.SolveLin(eqs)
	// var oSum float64
	oSum := fraction.NewRational(0, 1)
	for k, freq := range countsMapTwo {
		// oSum += freq.Float64() * oSoln[equations.Variable(k)]
		oSum = oSum.Plus(freq.Times(fraction.NewRationalFromFloat(oSoln[equations.Variable(k)])))
	}

	return oSum*/
	// fmt.Println("END FOR FUN SOLVE", oSum, "\n\n")

	pf.Start("Solve   ")
	fmt.Println("  Solve", time.Now())
	solns := equations.Solve(eqs, countsMapTwo)

	fmt.Println("  Sum", time.Now())
	pf.Start("Sum     ")
	sum := fr(0, 1)
	for k, freq := range countsMapTwo {
		sum = sum.Plus(freq.Times(solns[equations.Variable(k)]))
	}

	pf.End()

	fmt.Println("  Done", n, m, sum)
	fmt.Println("  Done at", time.Now())

	return sum
}

func simpleSolve() {
	eqs := []*equations.Equation{
		equations.NewEq(map[equations.Variable]*fraction.Rational{
			"E1": fr(1, 1),
		}, fr(0, 1)),
		equations.NewEq(map[equations.Variable]*fraction.Rational{
			"E1": fr(1, 2),
			"E2": fr(-1, 2),
		}, fr(1, 1)),
	}

	equations.Solve(eqs, nil)
}

const (
	precision = 14
)

func reduce(f *fraction.Rational) *fraction.Rational {
	n, d := f.Numer(), f.Denom()
	numerDigs := len(n.String())
	denomDigs := len(d.String())

	if numerDigs > 14 && denomDigs > 14 {
		post := fraction.NewBigRationalFromInt(n.SignificantDigits(precision).Int(), d.SignificantDigits(precision).Int())
		fmt.Println("PREPOST", f, post)
		panic("ARGH")
		return post
	}
	return f
}

/*
4 2 5/2
4 3 63/4
4 4 141/2

10/4
63/4
282/4

        4.127581862769911
				552516.7138726487
				4.127581428274515e+06
8 8 ==> 4.127581862693018e+06
        4.127581862538576e+06
				4.1275818626223775e+06
8 7 ==> 570256.9432509352
				570256.943244465
7 7 ==> 112348.31113028523
				241124.38773524063

*/

func findAllArrangementsTwo(ba *ballArrangement, countsMap map[string]*fraction.Rational) {
	if _, ok := countsMap[ba.smartCode()]; ok {
		return
	}

	countsMap[ba.smartCode()] = ba.calcInitProb()

	for i, v := range ba.balls {
		if v == 0 {
			continue
		}

		leftIdx := (i + ba.n - 1) % ba.n
		rightIdx := (i + 1) % ba.n

		// Move to the left
		ba.balls[i]--
		ba.balls[leftIdx]++
		findAllArrangementsTwo(ba, countsMap)
		ba.balls[i]++
		ba.balls[leftIdx]--

		// Move to the right
		ba.balls[i]--
		ba.balls[rightIdx]++
		findAllArrangementsTwo(ba, countsMap)
		ba.balls[i]++
		ba.balls[rightIdx]--
	}
}

func populateNeighbors(ba *ballArrangement, checked map[string]int, initProbMap map[int]float64, neighborStepMap map[int]map[int]float64) {
	if _, ok := checked[ba.smartCode()]; ok {
		return
	}
	curIdx := len(checked)
	checked[ba.smartCode()] = curIdx
	initProbMap[curIdx] = ba.calcInitProb().Float64()

	curNeighbors := map[int]float64{}
	for i, v := range ba.balls {
		if v == 0 {
			continue
		}

		oddsOfMove := float64(v) / (2.0 * float64(ba.m))

		leftIdx := (i + ba.n - 1) % ba.n
		rightIdx := (i + 1) % ba.n

		// Move to the left
		ba.balls[i]--
		ba.balls[leftIdx]++
		populateNeighbors(ba, checked, initProbMap, neighborStepMap)
		curNeighbors[checked[ba.smartCode()]] += oddsOfMove
		ba.balls[i]++
		ba.balls[leftIdx]--

		// Move to the right
		ba.balls[i]--
		ba.balls[rightIdx]++
		populateNeighbors(ba, checked, initProbMap, neighborStepMap)
		curNeighbors[checked[ba.smartCode()]] += oddsOfMove
		ba.balls[i]++
		ba.balls[rightIdx]--
	}
	neighborStepMap[curIdx] = curNeighbors

	// var neighbors []*neighborStep
	// for k, v := range curNeighbors {
	// 	if k == 0 {
	// 		continue
	// 	}
	// 	neighbors = append(neighbors, &neighborStep{k, v})
	// }
	// (*allNeighbors)[curIdx] = neighbors
}

func solveSimple(n, m int) float64 {

	fmt.Println("==================== SOLVING FOR", n, m)
	// uniqueArrs := []*ballArrangement{}

	balls := make([]int, n)
	balls[0] = m
	init := &ballArrangement{
		balls: balls,
		n:     n,
		m:     m,
	}
	countsMapTwo := map[string]*fraction.Rational{}
	findAllArrangementsTwo(init, countsMapTwo)
	numConfigurations := len(countsMapTwo)
	// fmt.Println("NUM CONFIGS", numConfigurations)

	init = &ballArrangement{
		balls: balls,
		n:     n,
		m:     m,
	}

	// neighborStepProb[a][b] is the probability of going to configuration b
	// from configuration b
	initProbMap := map[int]float64{}
	neighborStepMap := map[int]map[int]float64{}
	populateNeighbors(init, map[string]int{}, initProbMap, neighborStepMap)
	var initProbs []float64
	for i := 0; i < len(initProbMap); i++ {
		initProbs = append(initProbs, initProbMap[i])
	}

	// probArray is the probability that a given configuration
	// is solved after numStep steps
	numStep := 0
	probArray := make([]float64, numConfigurations)
	nextArray := make([]float64, numConfigurations)

	_ = numStep

	// When numStep is 0, then only P(solved configuration) = 1
	probArray[0] = 1

	revNeighborStepMap := map[int]map[int]float64{}
	for aCfg, m := range neighborStepMap {
		for bCfg, prob := range m {

			if _, ok := revNeighborStepMap[bCfg]; !ok {
				revNeighborStepMap[bCfg] = map[int]float64{}
			}

			revNeighborStepMap[bCfg][aCfg] += prob
		}
	}

	// revNeighborProb := make([][]*neighborStep, numConfigurations)
	neighborProb := make([][]*neighborStep, numConfigurations)
	for a := 0; a < numConfigurations; a++ {
		m := neighborStepMap[a]
		for b, prob := range m {
			// revNeighborProb[b] = append(revNeighborProb[b], &neighborStep{a, prob})
			if a > 0 {
				neighborProb[a] = append(neighborProb[a], &neighborStep{b, prob})
			}
		}
	}

	var totalEv float64
	// upTo := 300000
	for numStep = 1; ; numStep++ {
		var evIncr float64
		for i := range probArray {
			var sum float64
			for _, neighborStep := range neighborProb[i] {
				sum += neighborStep.stepProb * probArray[neighborStep.idx]
			}
			nextArray[i] = sum
			evIncr += sum * initProbs[i]
		}
		// fmt.Println("INCR", numStep, evIncr, )
		// if numStep == upTo {
		// 	fmt.Println("INCR", evIncr, evIncr*float64(numStep))
		// }
		curIncr := evIncr * float64(numStep)
		totalEv += curIncr
		if curIncr < 0.000000000001 {
			fmt.Println("BREAKING AT", numStep)
			break
		}
		// if numStep%100 == 0 {
		// 	fmt.Println(totalEv)
		// }

		nextArray, probArray = probArray, nextArray
	}
	fmt.Println(n, m, totalEv)

	return totalEv
}

type neighborStep struct {
	idx      int
	stepProb float64
}

func (ns *neighborStep) String() string {
	return fmt.Sprintf("{%d, %0.6f}", ns.idx, ns.stepProb)
}

/*

7 7 : 0.689
8 8 : 7.54s
9 9 : 508
*/
