package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P100() *problem {
	return noInputNode(100, func(o command.Output) {
		/* // Used to get initial values, which made me realize that ratios are constant
		prevB, prevR := 1.0, 1.0
		for blue := 15; blue < 1_000_000_00; blue++ {
			bf := float64(blue)

			a := 1.0
			b := 2.0*bf - 1.0
			c := bf - bf*bf

			for _, rf := range maths.QuadraticRoots(a, b, c) {
				if rf < 0 {
					continue
				}

				red := int(rf)
				if 2*blue*(blue-1) == (blue+red)*(blue+red-1) {
					fmt.Println("yup", blue, red, int(rf*(rf/prevR)))
					prevB, prevR = bf, rf
					_ = prevB
				}
			}
		}
		/**/

		rs := []int{6, 35}
		for {
			prev := rs[len(rs)-1]
			prevPrev := rs[len(rs)-2]
			nextR := int(float64(prev) * float64(prev) / float64(prevPrev))

			// B - B^2 + 2BR + R^2 - R = 0
			// Initially tried with blue ratio R^2 + (2B - 1) * R + B - B^2
			// - B^2 + (2R + 1)*B + (R^2 - R) = 0
			a := -1.0
			b := float64(2*nextR) + 1.0
			c := float64(nextR)*float64(nextR) - float64(nextR)

			opts := maths.QuadraticRoots(a, b, c)
			if len(opts) == 0 {
				o.Terminatef("red value %d doesn't result in valid quadratic", int(nextR))
			}

			ir := int(nextR)
			ib := int(opts[0])
			if 2*ib*(ib-1) != (ib+ir)*(ib+ir-1) {
				o.Terminatef("red value %d doesn't equal 1/2", ir)
			}
			rs = append(rs, nextR)
			if ib+ir >= 1_000_000_000_000 {
				o.Stdoutln(ib)
				return
			}
		}
	})
}
