package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P128() *problem {
	return intInputNode(128, func(o command.Output, n int) {
		seqLen := 1
		g := generator.Primes()
		// After checking all integers, noticed that only works
		// at first corner and last spot
		for i, oneJump, jump := 2, false, 5; seqLen < n; {
			layer := hexIntToLayer(i)
			layerStart := hexLayerToStart(layer)
			corner := hexCorner(layer, layerStart, i)
			side := hexSideNumber(layer, layerStart, i)
			var neighbors []int
			if corner {
				neighbors = hexCornerNeighbors(layer, i-layerStart, side)
			} else {
				neighbors = hexNonCornerNeighbors(layer, i-layerStart, side)
			}
			var numPrimes int
			for _, neighbor := range neighbors {
				if g.Contains(maths.Abs(i - neighbor)) {
					numPrimes++
				}
			}
			if numPrimes == 3 {
				seqLen++
			}
			if seqLen >= n {
				o.Stdoutln(i)
				return
			}
			if oneJump {
				i++
			} else {
				i += jump
				jump += 6
			}
			oneJump = !oneJump
		}
	})
}

func hexCornerNeighbors(layer, index, side int) []int {
	nextLen := hexLayerLen(layer + 1)
	nextSideIndex := side * (nextLen / 6)
	return []int{
		// next and previous in same layer
		hexGetValueFromIndex(layer, index+1),
		hexGetValueFromIndex(layer, index-1),
		hexGetCorner(layer-1, side),
		hexGetValueFromIndex(layer+1, nextSideIndex+1),
		hexGetValueFromIndex(layer+1, nextSideIndex),
		hexGetValueFromIndex(layer+1, nextSideIndex-1),
	}
}

func hexNonCornerNeighbors(layer, index, side int) []int {
	//nextLen := hexLayerLen(layer + 1)
	//nextSideIndex := side * (nextLen / 6)
	return []int{
		// next and previous in same layer
		hexGetValueFromIndex(layer, index+1),
		hexGetValueFromIndex(layer, index-1),
		hexGetValueFromIndex(layer+1, index+side),
		hexGetValueFromIndex(layer+1, index+side+1),
		hexGetValueFromIndex(layer-1, index-side),
		hexGetValueFromIndex(layer-1, index-side-1),
	}
}

func hexGetCorner(layer, side int) int {
	start := hexLayerToStart(layer)
	layerLen := hexLayerLen(layer)
	return start + (layerLen/6)*side
}

func hexGetValueFromIndex(layer, index int) int {
	start := hexLayerToStart(layer)
	layerLen := hexLayerLen(layer)
	return start + ((index + layerLen) % layerLen)
}

func hexGetValueFromSideIndex(layer, side, index int) int {
	start := hexLayerToStart(layer)
	sideLen := hexLayerLen(layer) / 6
	return start + sideLen*(side%6) + ((index + sideLen) % sideLen)
}

func hexIntToLayer(k int) int {
	if k == 1 {
		return 0
	}
	roots := maths.QuadraticRoots(1, 1, 1-(2*float64((k+4)/6)))
	return int(roots[1]) + 1
}

func hexLayerToStart(layer int) int {
	if layer == 0 {
		return 1
	}
	return (layer*(layer-1)/2)*6 + 2
}

func hexLayerLen(layer int) int {
	if layer == 0 {
		return 0
	}
	return 6 * layer
}

func hexCorner(layer, layerStart, k int) bool {
	return (k-layerStart)%layer == 0
}

// return 0 through 5
func hexSideNumber(layer, layerStart, k int) int {
	return (k - layerStart) / layer
}
