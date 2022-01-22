package twentyone

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/parse"
)

func D24() *problem {
	return command.SerialNodes(
		command.Description(""),
		command.ExecutorNode(func(o command.Output, d *command.Data) error {
			lines := parse.ReadFileLines(filepath.ToSlash(filepath.Join("input", "d24_google.txt")))
			var parsed [][]string
			for _, line := range lines {
				parsed = append(parsed, strings.Split(line, " "))
			}

			ops := parseCode(parsed)

			// Lower than
			//start := NewMN("11841264117189", ops)
			start := NewMN("11841231117189", ops)
			maxV = NewMN("11111111111110", ops)
			gBest = NewMN("11841231117189", ops)
			////////////// 11841264117189
			//gBest = NewMN("12996942829399", ops)

			validLenCheck = -1
			part1 = false
			bfs.ShortestPath(start, ops)
			fmt.Println("BEST", best)
			return nil
		}),
	)
}

type ModelNumber struct {
	digits []int
	value  *int
}

func NewMN(s string, ops []*Op) *ModelNumber {
	var is []int
	for i := 0; i < len(s); i++ {
		is = append(is, parse.Atoi(s[i:i+1]))
	}
	return NewModel(is, ops)
}

func NewModel(is []int, ops []*Op) *ModelNumber {
	return &ModelNumber{
		digits: is,
	}
}

func (mn *ModelNumber) Copy() *ModelNumber {
	r := &ModelNumber{}
	for _, d := range mn.digits {
		r.digits = append(r.digits, d)
	}
	return r
}

func (mn *ModelNumber) Code() string {
	return mn.String()
}

func (mn *ModelNumber) String() string {
	var r []string
	for _, n := range mn.digits {
		r = append(r, strconv.Itoa(n))
	}
	return fmt.Sprintf("[%d] %s", mn.value, strings.Join(r, ""))
}

func altRoute(ops []*Op) {
	zeros := map[string]bool{
		"0": true,
		"w": true,
		"x": true,
		"y": true,
		"z": true,
	}
	for _, op := range ops {
		switch op.kind {
		case InpOp, EqlOp:
			// Input is no longer 0
			delete(zeros, op.a)
		case AddOp:
			if zeros[op.b] {
				op.irrelevant = append(op.irrelevant, "adding 0")
			} else {
				delete(zeros, op.a)
			}
		case MulOp:
			if zeros[op.a] {
				op.irrelevant = append(op.irrelevant, "multiplying 0 by something")
			} else if op.b == "1" {
				op.irrelevant = append(op.irrelevant, "multiplying by one")
			} else if zeros[op.b] {
				zeros[op.a] = true
			}
		case DivOp:
			if zeros[op.a] {
				op.irrelevant = append(op.irrelevant, "dividing 0 by something")
			} else if op.b == "1" {
				op.irrelevant = append(op.irrelevant, "dividing something by one")
			}
		case ModOp:
			if zeros[op.a] {
				op.irrelevant = append(op.irrelevant, "modding 0 by something")
			}
		}
	}
}

var (
	gBest         *ModelNumber
	maxV          *ModelNumber
	minRequired   = 1
	part1         = true
	validLenCheck = 100
)

func (mn *ModelNumber) Done(opsI interface{}) bool {
	if dist := mn.Distance(opsI.([]*Op)); validLenCheck > 0 {
		return dist == 0 && len(valid) > validLenCheck
	} else {
		return dist == 0 && mn.Better(gBest)
	}
}

func (mn *ModelNumber) Better(that *ModelNumber) bool {
	for i := 0; i < len(mn.digits); i++ {
		if mn.digits[i] != that.digits[i] {
			if part1 {
				return mn.digits[i] > that.digits[i]
			}
			return mn.digits[i] < that.digits[i]
		}
	}
	return false
}

var (
	valid   = map[string]bool{}
	best    *ModelNumber
	checked = map[string]bool{}
)

func (mn *ModelNumber) Distance(opsI interface{}) int {
	if mn.value != nil {
		return *mn.value
	}
	vs := &varSet{}
	ints := mn.digits
	for _, o := range opsI.([]*Op) {
		if pop := o.execute(vs, ints); pop != 0 {
			ints = ints[pop:]
		}
	}
	val := vs.get("z")
	mn.value = &val
	return val
}

func (mn *ModelNumber) AdjacentStates(opsI interface{}) []bfs.State {
	ops := opsI.([]*Op)
	var r []bfs.State
	for i := 0; i < len(mn.digits); i++ {
		for j := 0; j <= 7; j++ {
			c := mn.Copy()
			c.digits[i] = ((c.digits[i] + j) % 9) + 1
			if !c.Better(gBest) || !maxV.Better(c) {
				continue
			}
			str := c.String()
			if checked[str] {
				continue
			}
			checked[str] = true
			if c.Distance(ops) == 0 {

				if !valid[str] {
					valid[str] = true
					if best == nil || (c.Better(best)) {
						best = c
					}
				}
			}
			r = append(r, c)
		}
	}
	return r
}

// Starting with the last input index:
// Find the values of w x y and z that work

type varSet map[string]int

func (vs *varSet) set(s string, i int) {
	(*vs)[s] = i
}

func (vs *varSet) get(s string) int {
	if r, ok := (*vs)[s]; ok {
		return r
	}
	if s == "w" || s == "x" || s == "y" || s == "z" {
		return 0
	}
	return parse.Atoi(s)
}

type Op struct {
	kind OpKind
	a    string
	b    string

	irrelevant []string
}

func (o *Op) String() string {
	var suffix string
	if len(o.irrelevant) > 0 {
		suffix = fmt.Sprintf("[%s]", strings.Join(o.irrelevant, " && "))
	}
	return fmt.Sprintf("%s %s %s%s", invOpMap[o.kind], o.a, o.b, suffix)
}

func (o *Op) execute(vs *varSet, code []int) int {
	switch o.kind {
	case InpOp:
		vs.set(o.a, code[0])
		return 1
	case AddOp:
		vs.set(o.a, vs.get(o.a)+vs.get(o.b))
	case MulOp:
		vs.set(o.a, vs.get(o.a)*vs.get(o.b))
	case DivOp:
		vs.set(o.a, vs.get(o.a)/vs.get(o.b))
	case ModOp:
		vs.set(o.a, vs.get(o.a)%vs.get(o.b))
	case EqlOp:
		if vs.get(o.a) == vs.get(o.b) {
			vs.set(o.a, 1)
		} else {
			vs.set(o.a, 0)
		}
	default:
		log.Fatalf("unknown op")
	}
	return 0
}

type OpKind int

const (
	InpOp OpKind = iota
	AddOp
	MulOp
	DivOp
	ModOp
	EqlOp
)

type OpExecutor struct {
	ops []Op
}

var (
	opMap = map[string]OpKind{
		"inp": InpOp,
		"add": AddOp,
		"mul": MulOp,
		"div": DivOp,
		"mod": ModOp,
		"eql": EqlOp,
	}
	invOpMap = map[OpKind]string{
		InpOp: "inp",
		AddOp: "add",
		MulOp: "mul",
		ModOp: "mod",
		DivOp: "div",
		EqlOp: "eql",
	}
)

func parseCode(lines [][]string) []*Op {
	var ops []*Op

	// First get inputs
	for _, line := range lines {
		k, ok := opMap[line[0]]
		if !ok {
			log.Fatalf("unknown kind: %v", line[0])
		}
		o := &Op{
			kind: k,
			a:    line[1],
		}
		if len(line) > 2 {
			o.b = line[2]
		}
		ops = append(ops, o)
	}
	return ops
}
