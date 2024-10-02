package p860

/*import (
	"fmt"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/combinatorics"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

const (
	mod = 989898989
)

func P860() *ecmodels.Problem {
	return ecmodels.IntInputNode(860, func(o command.Output, n int) {
		// o.Stdoutln(f(0, n, &stacks{}))
		// if !true {
		// 	g := &stacks{
		// 		goldGold:     0,
		// 		silverGold:   1,
		// 		goldSilver:   21,
		// 		silverSilver: 5,
		// 	}
		// 	o.Stdoutln("FAIR", g.isFair())
		// 	return
		// }

		// numer :=

		// o.Stdoutln(dp(0, n, 0, make([]int, len(vals)), maths.Factorial(n).ModInt(mod)))
		o.Stdoutln(clever(n))
		// o.Stdoutln(wut(n))
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

func wut(n int) int {
	var cnt int
	combinatorics.EvaluateCombos(&combinatorics.Combinatorics[int]{
		Parts:            []int{-4, -1, 1, 4},
		AllowReplacement: true,
		MinLength:        n,
		MaxLength:        n,
	}, func(t []int) {
		if bread.Sum(t) == 0 {
			cnt++
			// fmt.Println(t)
		}
	})
	return cnt
}

var (
	factoMod = []int{
		1, 1,
	}
)

func modFactorial(n int) int {
	for n >= len(factoMod) {
		lastIdx := len(factoMod) - 1
		factoMod = append(factoMod, (factoMod[lastIdx]*(lastIdx+1))%mod)
	}
	return factoMod[n]
}

func clever(n int) int {

	p := generator.Primes()
	_ = p

	numerator0 := p.PrimeFactoredNumberFactorial(n)
	numerator := modFactorial(n)

	var sum int
	for smallOffset, bigOffset := 0, 0; smallOffset+bigOffset <= n; smallOffset, bigOffset = smallOffset+2, bigOffset+8 {

		// Initial values are (0, 0,
		// (from combinatorics package), initial values are

		// Values can all be derived from initial a value:
		// d = a + smallOffset
		// We can then solve the following system of equations:
		// a + b + c + d = n
		// c = b + bigOffset
		//
		// a + b + (b + bigOffset) + (a + smallOffset) = n
		// 2*a + 2*b + bigOffset + smallOffset = n
		// b = (n - 2*a - bigOffset - smallOffest) / 2
		fmt.Println("OFFSET", smallOffset)

		// Initial values will be the above when a=0:
		// (0, (n - 2*0 - bigOffset - smallOffset)/2, b + bigOffset, 0 + smallOffset)
		// (0, n/2 - bigOffset/2 - smallOffset/2, n/2 - bigOffset/2 - smallOffset/2 + bigOffset, smallOffset)
		// (0, n/2 - bigOffset/2 - smallOffset/2, n/2 + bigOffset/2 - smallOffset/2, smallOffset)
		initialB := (n - bigOffset - smallOffset) / 2

		curCount0 := numerator0.TimesInt(1)
		curCount := numerator
		for _, v := range []int{initialB, initialB + bigOffset, smallOffset} {
			if v <= 1 {
				continue
			}

			divFact := modFactorial(v)
			curCount = (curCount * maths.PowMod(divFact, -1, mod)) % mod
			// curCount = (curCount * maths.PowMod( curCount.Div(maths.Factorial(v))
		}

		// count

		// Need b to be positive:
		// b = (n - 2*a - bigOffset - smallOffest) / 2 >= 0
		//     n >= 2*a + bigOffset + smallOffest

		// Smallest values after setting a are (a, 0, bigOffset, a+smallOffset)
		for a := 0; a+bigOffset+a+smallOffset <= n; a++ {

			b := (n - 2*a - bigOffset - smallOffset) / 2
			c := b + bigOffset
			d := a + smallOffset

			// (logic copied from combinatorics package)

			// v := combinatorics.PermutationFromCount([]int{a, b, c, d}).ModInt(mod)
			// fmt.Println(v, curCount.ToInt(p))

			v := curCount
			if smallOffset > 0 {
				// v = (v * 2) % mod
				sum = (sum + v) % mod
			}
			sum = (sum + v) % mod

			// a is incremented by one
			curCount = (curCount * maths.PowMod(a+1, -1, mod)) % mod
			// d is incremented by one
			curCount = (curCount * maths.PowMod(d+1, -1, mod)) % mod
			// b is decremented by one
			curCount = (curCount * b) % mod
			// c is decremented by one
			curCount = (curCount * c) % mod
		}
	}
	return sum
}

func f(idx, remaining int, game *stacks) int {
	if remaining == 0 {

		if game.isFair() {
			if !game.symmetric() {
				fmt.Println("===========")
				// fmt.Println(game.prettyString())
				fmt.Println(game)
				fmt.Println("FAIR ^")
			}
			return game.count()
		}
		return 0
	}

	var incr, decr func()

	switch idx {
	case 0:
		incr, decr = func() { game.goldGold++ }, func() { game.goldGold-- }
	case 1:
		incr, decr = func() { game.goldSilver++ }, func() { game.goldSilver-- }
	case 2:
		incr, decr = func() { game.silverGold++ }, func() { game.silverGold-- }
	case 3:
		incr, decr = func() { game.silverSilver++ }, func() { game.silverSilver-- }
	// case 4:
	// incr, decr = func() { game.gold++ }, func() { game.gold-- }
	// case 5:
	// incr, decr = func() { game.silver++ }, func() { game.silver-- }
	default:
		return 0
	}

	var cnt int
	for i := 0; i <= remaining; i++ {
		cnt += f(idx+1, remaining-i, game)
		incr()
	}

	for i := 0; i <= remaining; i++ {
		decr()
	}

	return cnt
}

type stacks struct {
	// bottomTop
	goldGold, goldSilver, silverGold, silverSilver, gold, silver int
}

func (s *stacks) symmetric() bool {
	return s.goldGold == s.silverSilver && s.goldSilver == s.silverGold
}

func (s *stacks) String() string {
	return fmt.Sprintf("%d %d %d %d", s.goldGold, s.silverGold, s.goldSilver, s.silverSilver)
}

func (s *stacks) prettyString() string {
	var top, bottom []string

	for i := 0; i < s.goldGold; i++ {
		bottom = append(bottom, "G")
		top = append(top, "G")
	}

	for i := 0; i < s.goldSilver; i++ {
		bottom = append(bottom, "G")
		top = append(top, "S")
	}

	for i := 0; i < s.silverGold; i++ {
		bottom = append(bottom, "S")
		top = append(top, "G")
	}

	for i := 0; i < s.silverSilver; i++ {
		bottom = append(bottom, "S")
		top = append(top, "S")
	}

	for i := 0; i < s.gold; i++ {
		bottom = append(bottom, "G")
		top = append(top, " ")
	}

	for i := 0; i < s.silver; i++ {
		bottom = append(bottom, "S")
		top = append(top, " ")
	}

	return fmt.Sprintf("%s\n%s", strings.Join(top, " "), strings.Join(bottom, " "))
}

func (s *stacks) isFair() bool {
	return !s.playerOneWins(true) && s.playerOneWins(false)
}

func (s *stacks) count() int {
	return combinatorics.PermutationFromCount([]int{
		s.goldGold, s.goldSilver, s.silverGold, s.silverSilver,
	}).ToInt()
}

func (s *stacks) playerOneWins(player1 bool) bool {
	if player1 { // gold
		var wins bool

		if s.goldGold > 0 {
			// Take both golds
			s.goldGold--
			wins = wins || s.playerOneWins(!player1)

			// Take one gold
			s.gold++
			wins = wins || s.playerOneWins(!player1)
			s.gold--
			s.goldGold++

			if wins {
				return true
			}
		}

		if s.goldSilver > 0 {
			s.goldSilver--
			wins = wins || s.playerOneWins(!player1)
			s.goldSilver++

			if wins {
				return true
			}
		}

		if s.silverGold > 0 {
			s.silverGold--
			s.silver++
			wins = wins || s.playerOneWins(!player1)
			s.silver--
			s.silverGold++

			if wins {
				return true
			}
		}

		if s.gold > 0 {
			s.gold--
			wins = wins || s.playerOneWins(!player1)
			s.gold++

			if wins {
				return true
			}
		}

		return wins
	}

	// player 2 (silver)
	var wins bool

	if s.silverSilver > 0 {
		// Take both silvers
		s.silverSilver--
		wins = wins || !s.playerOneWins(!player1)

		// Take one silver
		s.silver++
		wins = wins || !s.playerOneWins(!player1)
		s.silver--
		s.silverSilver++

		if wins {
			return false
		}
	}

	if s.silverGold > 0 {
		s.silverGold--
		wins = wins || !s.playerOneWins(!player1)
		s.silverGold++

		if wins {
			return false
		}
	}

	if s.goldSilver > 0 {
		s.goldSilver--
		s.gold++
		wins = wins || !s.playerOneWins(!player1)
		s.gold--
		s.goldSilver++

		if wins {
			return false
		}
	}

	if s.silver > 0 {
		s.silver--
		wins = wins || !s.playerOneWins(!player1)
		s.silver++

		if wins {
			return false
		}
	}

	return !wins
}
*/
