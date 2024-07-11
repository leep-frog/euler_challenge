package p33

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

func P33() *ecmodels.Problem {
	return ecmodels.NoInputNode(33, func(o command.Output) {
		topProd := 1
		bottomProd := 1
		for a1 := 1; a1 < 10; a1++ {
			for a2 := 1; a2 < 10; a2++ {
				for b1 := 1; b1 < 10; b1++ {
					for b2 := 1; b2 < 10; b2++ {
						topLeft := a1*10 + a2
						bottomLeft := b1*10 + b2
						if topLeft >= bottomLeft {
							continue
						}
						if (a1 == b1) && checker(topLeft, bottomLeft, a2, b2) {
							topProd *= topLeft
							bottomProd *= bottomLeft
						} else if a1 == b2 && checker(topLeft, bottomLeft, a2, b1) {
							topProd *= topLeft
							bottomProd *= bottomLeft
						} else if a2 == b1 && checker(topLeft, bottomLeft, a1, b2) {
							topProd *= topLeft
							bottomProd *= bottomLeft
						} else if a2 == b2 && checker(topLeft, bottomLeft, a1, b1) {
							topProd *= topLeft
							bottomProd *= bottomLeft
						}
					}
				}
			}
		}
		o.Stdoutln(topProd, bottomProd)
	}, &ecmodels.Execution{
		// Answer is actually 100
		Want: "387296 38729600",
	})
}

func checker(topLeft, bottomLeft, topRight, bottomRight int) bool {
	return topLeft*bottomRight == bottomLeft*topRight
}
