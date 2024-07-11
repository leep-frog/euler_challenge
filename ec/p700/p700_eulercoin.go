package p700

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P700() *ecmodels.Problem {
	return ecmodels.NoInputNode(700, func(o command.Output) {
		incr := uint64(1504170715041707)
		mod := uint64(4503599627370517)
		var sum uint64

		best := maths.Smallest[int, uint64]()

		for at := incr; !best.Set() || best.Best() > 1; at = (at + incr) % mod {
			prevBest := best.Best()

			if best.Check(at) {
				sum += at

				// Based on how modulus works, the difference will be the same every
				// k iterations (and will guaranteed not get lower before that)
				// so just automatically jump down by that amount
				if prevBest != 0 {
					diff := prevBest - at
					for at > diff {
						at -= diff
						best.Check(at)
						sum += at
					}
				}

				// We don't need to increment redundantly above the current best (at),
				// so lower the mod by incr
				for mod > at+incr {
					mod = mod - incr
				}

				// Also redundant if the increment is larger than the modulus, so imrpove that as well
				if incr > mod {
					incr = incr % mod
				}
			}
		}

		o.Stdoutln(sum)
	}, &ecmodels.Execution{
		Want: "1517926517777556",
	})
}
