package eulerchallenge

import (
	"github.com/leep-frog/command"
)

func P155() *problem {
	return intInputNode(155, func(o command.Output, n int) {
		o.Stdoutln(n)
	})
}

func d155(n int) {

}

type capacitorGenerator struct {
	remaining int
}

func (c *capacitorGenerator) Code() string {
	return ""
}

func (c *capacitorGenerator) OnPush() {
}

func (c *capacitorGenerator) OnPop() {
}

func (c *capacitorGenerator) Done() bool {
	return false
}

func (c *capacitorGenerator) AdjacentStates() []*capacitorGenerator {
	if c.remaining <= 0 {
		return nil
	}

	var cgs []*capacitorGenerator

	return cgs
}

/*type capacitorGenerator struct {
	remaining int
}

func (c *capacitorGenerator) Code() string {
	return ""
}

func (c *capacitorGenerator) Done() bool {
	return false
}

func (c *capacitorGenerator) AdjacentStates() []*capacitorGenerator {
	if c.remaining <= 0 {
		return nil
	}

	var cgs []*capacitorGenerator

	return cgs
}

type Capacitor interface {
	copy() Capacitor
	capacitance() float64
	mutations(serial bool) []Capacitor
}

type capacitor struct {
	value float64
}

func (c *capacitor) copy() *capacitor {
	return &capacitor{c.value}
}

func (c *capacitor) capacitance() float64 {
	return c.value
}

func (c *capacitor) mutations(serial bool) float64 {

}

func (pc *parallelCircuit) capacitance() float64 {
	var sum float64
	for _, m := range pc.modules {
		sum += m.capacitance()
	}
	return sum
}

func (pc *parallelCircuit) mutations() []Capacitor {
	var pcs []Capacitor
	for _, m := range pc.modules {
		pcs = append(pcs, m.mutations()...)
	}
	pcs = append(pcs, &serialCircuit{
		[]Cap
	})
	return pcs
}

type parallelCircuit struct {
	modules []Capacitor
}

func (sc *serialCircuit) capacitance() float64 {
	var inverseSum float64
	for _, m := range sc.modules {
		inverseSum += 1.0 / m.capacitance()
	}
	return 1.0 / inverseSum
}

func (pc *parallelCircuit) copy() *parallelCircuit {
	var newMs []Capacitor
	for _, m := range pc.modules {
		newMs = append(newMs, m.copy())
	}
	return &parallelCircuit{newMs}
}

type serialCircuit struct {
	modules []Capacitor
}

func (sc *serialCircuit) copy() *serialCircuit {
	var newMs []Capacitor
	for _, m := range sc.modules {
		newMs = append(newMs, m.copy())
	}
	return &serialCircuit{newMs}
}
*/
