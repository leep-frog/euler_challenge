package equations

import (
	"fmt"
	"strings"

	"github.com/leep-frog/euler_challenge/fraction"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/profiler"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"gonum.org/v1/gonum/mat"
)

/*
- [8-10x]    (8-10x): Use big.Float instead of big.Rat (possibly adding larger square)
- [160-200x] (20x) Remove all map references-
- Better sorting
*/

var (
	pf = Profiler
)

type Variable string

// VarSet is a collection of (variable, coefficient) pairs
type VariableSet struct {
	m map[Variable]*fraction.Rational
	c *fraction.Rational
}

func (vs *VariableSet) String() string {
	vars := maps.Keys(vs.m)
	slices.Sort(vars)
	var r []string
	for _, v := range vars {
		r = append(r, fmt.Sprintf("%v*%s", vs.m[v], v))
	}
	r = append(r, fmt.Sprintf("%v", vs.c))

	return strings.Join(r, " + ")
}

var (
	zero = f(0, 1)
)

func (vs *VariableSet) Times(fr *fraction.Rational) *VariableSet {
	if fr.EQ(zero) {
		return &VariableSet{
			map[Variable]*fraction.Rational{},
			zero.Copy(),
		}
	}

	res := &VariableSet{
		map[Variable]*fraction.Rational{},
		reduce(vs.c.Times(fr)),
	}

	for variable, value := range vs.m {
		res.m[variable] = reduce(value.Times(fr))

	}
	return res
}

func (vs *VariableSet) Plus(that *VariableSet) *VariableSet {
	keys := map[Variable]bool{}
	Profiler.Start("KEY ITER")
	for k := range vs.m {
		keys[k] = true
	}
	for k := range that.m {
		keys[k] = true
	}
	Profiler.End()

	res := &VariableSet{
		map[Variable]*fraction.Rational{},
		reduce(vs.c.Plus(that.c)),
	}

	for k := range keys {
		a, b := vs.m[k], that.m[k]
		if a == nil {
			res.m[k] = b.Copy()
		} else if b == nil {
			res.m[k] = a.Copy()
		} else {
			Profiler.Start("PLUSING")
			v := reduce(a.Plus(b))
			Profiler.Start("AFTER PLUSING")
			if v.NEQ(zero) {
				res.m[k] = v
			}
		}
	}

	return res
}

func (vs *VariableSet) ApplyInPlace(soln *VariableSolution) {
	coef, ok := vs.m[soln.variable]
	if !ok {
		return
	}
	delete(vs.m, soln.variable)

	cOffset := reduce(coef.Times(soln.value))
	vs.c = reduce(vs.c.Plus(cOffset))
}

func (vs *VariableSet) GetSolution() *VariableSolution {
	if len(vs.m) != 1 {
		// panic("Can only get a solution if only one constant is defined")
		return nil
	}

	variable := maps.Keys(vs.m)[0]
	coef := vs.m[variable].Reciprocal().Negate()
	value := reduce(vs.c.Times(coef))
	return &VariableSolution{
		variable: variable,
		value:    value,
	}
}

func (vs *VariableSet) HasVar(v Variable) bool {
	_, ok := vs.m[v]
	return ok
}

// Equation is an equation of the form a1*x1 + a2*x2 + ... + an*xn + c = eq
type Equation struct {
	vs *VariableSet
	// eq *fraction.Rational
}

func NewEq(variables map[Variable]*fraction.Rational, c *fraction.Rational) *Equation {
	return &Equation{NewVs(variables, c)}
}

func NewVs(variables map[Variable]*fraction.Rational, c *fraction.Rational) *VariableSet {
	if c == nil {
		c = zero.Copy()
	}
	return &VariableSet{variables, c}
}

func (e *Equation) String() string {
	return fmt.Sprintf("%v = %v", e.vs, "e.eq")
}

func (e *Equation) HasVar(v Variable) bool {
	return e.vs.HasVar(v)
}

var (
	Profiler = profiler.New()
)

func SolveLin(eqs []*Equation) map[Variable]float64 {
	// gonum
	// var b linsolve.BiCG
	// _ = b
	// linsolve.Iterative()

	// data := make([]float64, 36)
	// for i := range data {
	// 	data[i] = 0
	// }
	// a := mat.NewDense(6, 6, data)
	// _ = a

	variableToIndex := map[Variable]int{}
	indexToVariable := []Variable{}
	for _, eq := range eqs {
		for k := range eq.vs.m {
			if _, ok := variableToIndex[k]; !ok {
				variableToIndex[k] = len(variableToIndex)
				indexToVariable = append(indexToVariable, k)
			}
		}
	}

	n := len(variableToIndex)

	// fmt.Println("NUM VARS + EQS", len(variableToIndex), len(eqs), variableToIndex)

	var data, bData []float64
	// matrix := sparse.NewDOK(n, n)
	for _, eq := range eqs {
		eqData := make([]float64, n)
		fr := make([]*fraction.Rational, n)
		for k, v := range eq.vs.m {
			eqData[variableToIndex[k]] = v.Float64()
			fr[variableToIndex[k]] = v
			// matrix.Set(eqi, variableToIndex[k], v.Float64())
		}
		// fmt.Println(eq)
		// fmt.Println(eqData)
		data = append(data, eqData...)
		// fmt.Println(strings.ReplaceAll(fmt.Sprintf("%v", fr), "<nil>", "0"))
		bData = append(bData, -eq.vs.c.Float64())
	}

	// sparse.

	a := mat.NewDense(n, n, data)

	// fmt.Println("A", data)
	// a.
	b := mat.NewVecDense(n, bData)
	// _ = a
	// fmt.Println("B", b)
	// a := mat.NewBandDense(n, n, 1, 1, data)
	// fmt.Println("PRE", b)
	b.SolveVec(a, b)
	// fmt.Println("SOLS", b.At(0, 0), b.At(1, 0))

	r := map[Variable]float64{}
	for idx := 0; idx < b.Len(); idx++ {
		r[indexToVariable[idx]] = b.At(idx, 0)
	}
	return r

	// res, err := linsolve.Iterative(a, b, &linsolve.BiCG{}, nil)
	// fmt.Println("RES", res)
	// fmt.Println("ERR", err)
	// fmt.Println("POST", b)

	// soln, err := linsolve.Iterative(a, nil, nil, nil)
	// linsolve.BiCG
}

func Solve(eqs []*Equation, countsMap map[string]*fraction.Rational) map[Variable]*fraction.Rational {
	/*varCntMap := map[Variable]int{}

	for _, eq := range eqs {
		for variable := range eq.vs.m {
			varCntMap[variable]++
		}
	}

	var vars []Variable
	for variable := range varCntMap {
		vars = append(vars, variable)
	}

	// TODO: try a different sort order?
	slices.SortFunc(vars, func(a, b Variable) int {
		if varCntMap[a] < varCntMap[b] {
			return -1
		}
		if varCntMap[a] > varCntMap[b] {
			return 1
		}

		if a < b {
			return -1
		}
		if a > b {
			return 1
		}

		return 0
	})*/

	// fmt.Println(time.Now(), "RECUR SOLVE", len(eqs), len(vars))

	fmt.Println("NUM EQS", len(eqs))
	solns := solve(eqs, nil)
	slices.SortFunc(solns, func(a, b *VariableSolution) int {
		if a.value.LT(b.value) {
			return -1
		}
		return 1
	})

	for i, s := range solns {
		if i > 0 {
			fmt.Printf("(+ %v =) ", s.value.Minus(solns[i-1].value))
		}
		// fmt.Println(s.variable, s.value) //, fmt.Sprintf("with prob: (%v)", countsMap[string(s.variable)]), s.value.Times(countsMap[string(s.variable)]))
		fmt.Println(s.variable, s.value)
		/*k := functional.Map(strings.Split(strings.ReplaceAll(strings.ReplaceAll(string(s.variable), "]", ""), "[", ""), " "), parse.Atoi)
		dists := make([]int, 4)
		for ai, a := range k {
			if a == 0 {
				continue
			}
			for abDist, b := range k[ai+1:] {
				if b == 0 {
					continue
				}
				dists[maths.Min(1+abDist, len(k)-(abDist+1))] += a * b
			}
		}
		fmt.Println(s.variable, s.value, dists) //, fmt.Sprintf("with prob: (%v)", countsMap[string(s.variable)]), s.value.Times(countsMap[string(s.variable)]))*/
		// fmt.Println(k, dists)
	}

	// fmt.Println(time.Now(), "RECUR SOLVED")

	r := map[Variable]*fraction.Rational{}
	for _, soln := range solns {
		r[soln.variable] = soln.value
	}
	return r
}

type VariableSolution struct {
	variable Variable
	value    *fraction.Rational
}

func (vs *VariableSolution) String() string {
	return fmt.Sprintf("%v === %v", vs.variable, vs.value)
}

func solve(eqs []*Equation, varOrderOld []Variable) []*VariableSolution {
	fmt.Println("EQS LEFT", len(eqs))

	varCntMap := map[Variable]int{}

	pf.Start("VAR CNT MAP")
	// O(n by n)
	// Need this to be n or n lg(n)
	// Perhaps choose a heap?
	for _, eq := range eqs {
		for variable := range eq.vs.m {
			varCntMap[variable]++
		}
	}

	pf.Start("VAR ORDER")
	// O(n)
	var varOrder []Variable
	for variable := range varCntMap {
		varOrder = append(varOrder, variable)
	}

	pf.Start("VAR SORT")
	// TODO: try a different sort order?
	// O(n log(n))
	slices.SortFunc(varOrder, func(a, b Variable) int {
		if varCntMap[a] < varCntMap[b] {
			return -1
		}
		if varCntMap[a] > varCntMap[b] {
			return 1
		}

		if a < b {
			return -1
		}
		if a > b {
			return 1
		}

		return 0
	})

	if len(varOrder) == 0 {
		return nil
	}

	solveVar := varOrder[0]
	var nextEqs []*Equation

	pf.Start("GET FIRST")
	// Get the first equation that has the variable
	// O(n)
	var solveEq *Equation
	for _, eq := range eqs {
		if solveEq == nil && eq.HasVar(solveVar) {
			solveEq = eq
		} else {
			nextEqs = append(nextEqs, eq)
		}
	}

	// Get the best equation that has the variable
	/*var solveEq *Equation
	var bestCnt, bestIdx int
	for i, eq := range eqs {
		if eq.HasVar(solveVar) {
			if solveEq == nil || len(eq.vs.m) < bestCnt {
				bestCnt = len(eq.vs.m)
				bestIdx = i
				solveEq = eq
			}
			// if len(eq.
		}
	}
	nextEqs = append(eqs[:bestIdx], eqs[bestIdx+1:]...)*/

	if solveEq == nil {
		panic(fmt.Sprintf("No equation has variable %s", solveVar))
	}

	// Let x_s, a_s be solveVar and its coefficient in solveEq
	// a1 * x1 + a2 * x2 + ... + a_s * x_s + ... an * xn + c = 0
	// a1 * x1 + a2 * x2 + ...       +       ... an * xn + c = -a_s * x_s
	// (a1 * x1 + a2 * x2 + ...       +       ... an * xn + c)/(-a_s) = x_s
	// (-a1/a_s * x1 - a2/a_s * x2 + ...       +       ... - an/(a_s) * xn - c/a_s = x_s

	Profiler.Start("RECIP")
	aSReciprocal := solveEq.vs.m[solveVar].Reciprocal().Negate()

	Profiler.Start("TIMES RECIP")
	invertedVS := solveEq.vs.Times(aSReciprocal)

	Profiler.Start("DELETE")
	delete(invertedVS.m, solveVar)

	// O(n x n)
	for _, eq := range nextEqs {
		if !eq.HasVar(solveVar) {
			continue
		}

		solveCoef := eq.vs.m[solveVar]

		Profiler.Start("TIMES SOLVE COEF")
		addEq := invertedVS.Times(solveCoef)

		Profiler.Start("PLUS ADD EQ")
		eq.vs = eq.vs.Plus(addEq)

		Profiler.Start("DELETE")
		delete(eq.vs.m, solveVar)
		Profiler.End()
	}

	solns := solve(nextEqs, varOrder[1:])

	for _, soln := range solns {
		Profiler.Start("APPLY IN PLACE")
		solveEq.vs.ApplyInPlace(soln)
	}

	// if len(eq.

	Profiler.Start("GET SOLN")
	if soln := solveEq.vs.GetSolution(); soln != nil {
		solns = append(solns, soln)
	} else {
		panic("Panic I think?")
	}

	Profiler.End()

	return solns

}

func f(num, den int) *fraction.Rational {
	return fraction.NewRational(num, den)
}

var (
	threshold, _ = maths.IntFromString("1000000000000000000")
	precision    = 14
)

func reduce(f *fraction.Rational) *fraction.Rational {
	return f
	n, d := f.Numer(), f.Denom()
	/**/
	for n.Abs().GT(threshold) && d.Abs().GT(threshold) {
		n, d = n.DivInt(10), d.DivInt(10)
	}
	/**/

	r := fraction.NewBigRationalFromInt(n.Int(), d.Int())
	c := fmt.Sprintf("%v", r)
	if len(c) > 100 {
		panic("NOO")
	}
	return r
}
