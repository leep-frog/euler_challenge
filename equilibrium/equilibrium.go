package equilibrium

import (
	"github.com/leep-frog/euler_challenge/maths"
)

type Mappable interface {
	~int | ~string
}

type State[M, T any, K Mappable] interface {
	Paths(M) []*WeightedPath[M, T, K]
	Code(M) K
}

type WeightedPath[M, T any, K Mappable] struct {
	State State[M, T, K]
	Weight float64
}

func Equilibrium[M any, T State[M, T, K], K Mappable](globalContext M, allStates []T, initialWeights map[K]float64) map[K]float64 {
	ws := initialWeights
	normalize(ws)

	delta := 0.000001
	
	for i := 0; i < 1000; i++{
		newWs := map[K]float64{}
		for _, s := range allStates {
			c := s.Code(globalContext)
			curWeight := ws[c]
			if curWeight == 0 {
				continue
			}
			for _, path := range s.Paths(globalContext) {
				newWs[path.State.Code(globalContext)] += curWeight * path.Weight
			}
		}
		normalize(newWs)
		oldWs := ws
		ws = newWs
		if len(ws) != len(oldWs) {
			continue
		}

		changed := false
		for k, v := range oldWs {
			nv, ok := ws[k]
			if !ok || maths.Abs(nv - v) >= delta {
				changed = true
				break
			}
		}
		if !changed {
			break
		}
	}
	return ws
}

func normalize[K Mappable](ws map[K]float64) {
	total := 0.0
	for _, v := range ws {
		total += v
	}
	for k := range ws {
		ws[k] /= total
	}
}
