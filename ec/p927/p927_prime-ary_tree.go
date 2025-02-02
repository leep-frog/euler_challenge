package p927

import (
	"fmt"
	"time"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

var (
	p = generator.Primes()
)

func P927() *ecmodels.Problem {
	return ecmodels.IntInputNode(927, func(o command.Output, input int) {

		rem := []int{
			2, 5, 149, 293, 1601,
			//41897,
			45197, 57977, 58337,
			// 61553,
			65357, 65537, 65789, 99173, 105269, 132857,
			175853, 200297, 287933, 313133, 318629, 319469,
			404837, 457277, 474017, 495413, 613577, 633833,
			657653, 720569, 870329, 967529, 1008353, 1015277,
			1028309, 1261769, 1265129, 1500833, 1623473, 2045909,
			2118629, 2141417, 2423777, 2481317, 3711353, 3828389,
			4145033, 4502117, 4685393, 5063453, 5099033, 5485589,
			5941877, 6022889, 6052373, 6403193, 6441389, 6483077,
			6580097, 6990317, 7152389, 8135189, 8203409, 9469709,
		}

		eligiblePrimes := map[int]bool{}
		// for i := 0; p.Nth(i) <= n; i++ {
		// 	eligiblePrimes[p.Nth(i)] = true
		// }

		fmt.Println(len(rem))
		var rel []int
		for _, v := range rem {
			if v <= input {
				rel = append(rel, v)
				eligiblePrimes[v] = true
			}
		}
		fmt.Println(eligiblePrimes)
		fmt.Println(calcSum(rel, input))
		loop(input, eligiblePrimes)

		// elegant2(input)
		// loop(input)
		return

		m := map[int]bool{
			2: true, 5: true, 149: true, 293: true,
		}

		for i := 0; p.Nth(i) < 300; i++ {
			if m[p.Nth(i)] {
				fmt.Println(i, p.Nth(i))
			}
		}

		otherSearch(input)
		return

		// opt3(input)
		// return

		numPrimes := 3

		notFounds := map[int]map[int]bool{}
		var primes []int
		var nodes []*node
		for i := 0; i < numPrimes; i++ {
			prime := p.Nth(i)
			primes = append(primes, prime)
			notFounds[prime] = map[int]bool{}

			for j := 1; j <= input; j++ {
				notFounds[prime][j] = true
			}

			nodes = append(nodes, &node{prime, 1, maths.One()})
		}

		foundCounts := map[int]int{}
		// for j := 1; j <= input; j++ {
		// 	foundCounts[j] = 0
		// }
		ctx := &context{
			notFounds,
			foundCounts,
			nil,
			numPrimes,
			input,
		}

		bfs.ContextDistanceSearch(ctx, nodes)

		/*candidates := map[int]bool{}
		for i := 1; i <= input; i++ {
			candidates[i] = true
		}

		for pi := 0; pi < 10; pi++ {
			k := p.Nth(pi)
			// calc(p, k)
			// tk(p, k)
			apply(p, k, candidates)
			// els := []int{
			// 	0,
			// 	1,
			// }
			// counts := map[int]int{
			// 	k: 1,
			// }
			// for n := 1; n <= 10; n++ {
			// 	ncs := map[int]int{}
			// 	for num, cnt := range counts {

			// 	}
			// }
		}
		// o.Stdoutln(n)*/
	}, []*ecmodels.Execution{
		{
			Args: []string{"1"},
			Want: "",
		},
		{
			Args: []string{"1000000"},
			Want: "207282955",
		},
	})
}

func tk(p *generator.Prime, k int) {
	fmt.Println("======================")
	fmt.Println("TK:", k)
	v := 1
	for iter := 0; iter <= 15/k; iter++ {
		next := maths.Pow(v, k) + 1

		var fs map[int]int

		if next <= 210_066_388_901 {
			fs = p.PrimeFactors(next)
		}
		fmt.Println("ITER", iter, next, fs)
		v = next
	}
}

func pow(v *maths.Int, pow int) *maths.Int {
	k := maths.One()
	for i := 0; i < pow; i++ {
		k = k.Times(v)
	}
	return k
}

func apply(p *generator.Prime, k int, candidates map[int]bool) {
	// fmt.Println("======================")
	// fmt.Println("TK:", k)

	notFound := map[int]bool{}
	for k := range candidates {
		notFound[k] = true
	}

	v := maths.One()
	for iter := 0; iter <= 15/k; iter++ {
		next := pow(v, k).PlusInt(1)

		var rm []int

		for nf := range notFound {
			if next.ModInt(nf) == 0 {
				rm = append(rm, nf)
			}
		}

		for _, nf := range rm {
			delete(notFound, nf)
		}

		v = next
	}

	for nf := range notFound {
		fmt.Printf("K (%d) removed %d\n", k, nf)
		delete(candidates, nf)
	}
}

func calc(p *generator.Prime, k int) {
	sizeToCounts := map[int]int{
		1: 1,
	}

	fmt.Println("CALCING", k)

	cumSum := 1

	for iter := 0; iter <= 15/k; iter++ {

		next := map[int]int{}

		var sum int

		for topLevelLeaves, cnt := range sizeToCounts {
			// Currently considering configurations that have 'topLevelLeaves' top-level leaves

			// The next tier can be achieved by expanding 1, 2, ..., topLevelLeaves leaves
			for numExpanded := 1; numExpanded <= topLevelLeaves; numExpanded++ {

				// There are permCnt ways to expand exactly numExpanded leaves
				permCnt := cnt * maths.Choose(topLevelLeaves, numExpanded).ToInt()

				// This results in a new
				nextTopLevelLeaves := numExpanded * k

				next[nextTopLevelLeaves] += permCnt
				sum += permCnt
				// next[]

				// next[]
			}
		}

		cumSum += sum

		sizeToCounts = next
		fmt.Println(iter)
		fmt.Println(sum, p.PrimeFactors(sum), p.Factors(sum))
		fmt.Println(cumSum, p.PrimeFactors(cumSum), p.Factors(cumSum))
	}
}

type context struct {
	notFounds   map[int]map[int]bool
	foundCounts map[int]int
	foundInAll  []int
	numPrimes   int
	upTo        int
}

type node struct {
	pow  int
	iter int
	k    *maths.Int
}

func (n *node) String() string {
	return fmt.Sprintf("{%d-%v}", n.pow, n.k)
}

func (n *node) Code(*context) string {
	return n.String()
}

func (n *node) Distance(ctx *context) bfs.Int {
	return bfs.Int(10000*n.iter + n.pow)
}

var (
	counter = 0
)

func (n *node) Done(ctx *context) bool {

	// fmt.Println("CHECKING", n.pow, n.iter)
	var rm []int
	for v := range ctx.notFounds[n.pow] {
		if n.k.ModInt(v) == 0 {
			rm = append(rm, v)
			fmt.Println("FOUND IN", n.pow, v)
			fmt.Println(ctx)
		}
	}

	for _, v := range rm {
		delete(ctx.notFounds[n.pow], v)
		ctx.foundCounts[v]++
		if ctx.foundCounts[v] == ctx.numPrimes {
			ctx.foundInAll = append(ctx.foundInAll, v)
			fmt.Println("FOUND IN ALL", v, ctx.foundInAll)

			for k, v := range ctx.notFounds {
				nfs := maps.Keys(v)
				slices.Sort(nfs)
				fmt.Println("NFS", k, nfs)
			}
		}
	}

	counter++
	if counter%1_000_000 == 0 {
		for k, v := range ctx.notFounds {
			nfs := maps.Keys(v)
			slices.Sort(nfs)
			fmt.Println("NFS", k, nfs)
		}
	}

	return false
}

func (n *node) AdjacentStates(ctx *context) []*node {
	return []*node{{
		n.pow,
		n.iter + 1,
		pow(n.k, n.pow).PlusInt(1),
	}}
}

func opt(k int) {
	var sum int
	for i, incr := 1, 1; i <= k; i, incr = i+incr, incr+2 {
		sum += i
		fmt.Println(sum, i)
	}
}

func opt2(k int) {
	var sum int
	for n := 0; n < k; n++ {
		sum += n*n + 1
		fmt.Println("OPT 2", sum)
	}
	// fmt.Println("OPT2", sum)
}

func opt3(k int) {
	var sum int
	for n := 1; n < k; {
		sum += n
		if n%2 == 1 {
			n *= 2
		} else {
			n = n*2 + 1
		}
		fmt.Println("3 OPT", sum, n)
	}
	// fmt.Println("OPT2", sum)
}

var (
	seqs = [][]int{
		{1, 2, 5, 10, 13, 17, 25, 26, 29, 34, 37, 41, 50, 53, 58, 61, 65, 73, 74, 82, 85, 89, 97, 101, 106, 109, 113, 122, 125, 130, 137, 145, 146, 149, 157, 169, 170, 173, 178, 181, 185, 193, 194, 197, 202, 205, 218, 221, 226, 229, 233, 241, 250, 257, 265, 269, 274, 277, 281, 289},
	}
)

type node2 struct {
	pow   int
	prime int
	iter  int
	value int
}

func (n *node2) String() string {
	return fmt.Sprintf("{%d-%d-%d-%d}", n.pow, n.prime, n.iter, n.value)
}

func (n *node2) Code(*context) string {
	return n.String()
}

func (n *node2) Done(ctx *context) bool {
	return false
}

func (n *node2) AdjacentStates(ctx *context) []*node2 {

	coef := n.value
	newValue := n.value
	for i := 1; i < n.pow; i++ {
		newValue = (newValue * coef) % n.prime
	}
	newValue = (newValue + 1) % n.prime

	if newValue == 0 {
		delete(ctx.notFounds[n.pow], n.prime)
		ctx.foundCounts[n.prime]++
		if ctx.foundCounts[n.prime] == ctx.numPrimes {
			ctx.foundInAll = append(ctx.foundInAll, n.prime)
			slices.Sort(ctx.foundInAll)
			fmt.Println("FOUND IN ALL", n.prime, ctx.foundInAll, calcSum(ctx.foundInAll, ctx.upTo))
			ctx.printNfs()
		}

		counter++
		if counter%1 == 0 {
			ctx.printNfs()
		}

		// fmt.Println("FOUND ONE", n.pow, n.prime, ctx.foundInAll)
		// fmt.Println(ctx.foundCounts, ctx.notFounds)
		return nil
	}

	return []*node2{{
		n.pow,
		n.prime,
		n.iter + 1,
		newValue,
	}}
}

func calcSum(primes []int, upTo int) int {
	// c := &combinatorics.Combinatorics{}
	// var r []int
	// for _, p := range combinatorics.ChooseAllSets(primes) {
	// 	prod := bread.Product(p)
	// 	if prod <= upTo {
	// 		r = append(r, prod)
	// 	}
	// }

	// return append(r, bread.Sum(r)+1)
	var sum int
	getProds(primes, 1, upTo, func(i int) { sum += i })
	return sum
}

func getProds(cur []int, prod int, upTo int, f func(int)) {
	if prod > upTo {
		return
	}

	if len(cur) == 0 {
		f(prod)
		return
	}

	getProds(cur[1:], prod*cur[0], upTo, f)
	getProds(cur[1:], prod, upTo, f)
}

func otherSearch(upTo int) {
	p := generator.Primes()
	_ = p

	// opt3(input)
	// return

	numPrimes := 14
	// searchPrimes := 168

	notFounds := map[int]map[int]bool{}
	var nodes []*node2
	for i := 0; i < numPrimes; i++ {
		prime := p.Nth(i)
		notFounds[prime] = map[int]bool{}

		for j := 0; p.Nth(j) <= upTo; j++ {
			notFounds[prime][p.Nth(j)] = true

			nodes = append(nodes, &node2{prime, p.Nth(j), 1, 1})
		}

	}

	foundCounts := map[int]int{}
	// for j := 1; j <= input; j++ {
	// 	foundCounts[j] = 0
	// }
	ctx := &context{
		notFounds,
		foundCounts,
		nil,
		numPrimes,
		upTo,
	}

	// fmt.Println("START", ctx)

	bfs.ContextSearch(ctx, nodes)
}

func (ctx *context) printNfs() {
	return

	fmt.Println("PRNFS ----------")
	keys := maps.Keys(ctx.notFounds)
	slices.Sort(keys)
	for _, k := range keys {
		nfs := maps.Keys(ctx.notFounds[k])
		slices.Sort(nfs)

		max := maths.Max(nfs...)
		var hmm []int
		for i := 0; p.Nth(i) <= max; i++ {
			if (p.Nth(i)-1)%k == 0 {
				hmm = append(hmm, p.Nth(i))
			}
		}

		if !slices.Equal(nfs, hmm) {

			var diff1 []int
			for _, v := range nfs {
				if !slices.Contains(hmm, v) {
					diff1 = append(diff1, v)
					// fmt.Println("WTF")
					// panic("AHAHAHA")
				}
			}
			var diff2 []int
			for _, v := range hmm {
				if !slices.Contains(nfs, v) {
					diff2 = append(diff2, v)
				}
			}
			fmt.Println("NFS", diff1, diff2, k, nfs, hmm)
			for _, d := range diff2 {
				fmt.Printf("{%d: %v}\n", d-1, p.PrimeFactors(d-1))
			}
			fmt.Println()
		}

	}
	fmt.Println()
}

func elegant(n int) {
	foundByAll := make([]bool, n+1)

	for i := range foundByAll {
		foundByAll[i] = true
	}

	for i := 1; p.Nth(i) <= n; i++ {
		prime := p.Nth(i)

		for k := prime; k < n; k += prime {
			fmt.Println("NO MORE", prime, k+1)
			foundByAll[k+1] = false
		}
	}

	relevantPrimes := []int{}
	for i, v := range foundByAll {
		if v && p.Contains(i) {
			relevantPrimes = append(relevantPrimes, i)
		}
	}
	fmt.Println(relevantPrimes)
}

func elegant2(n int) {
	notFounds := make([]int, n+1)

	for i := 1; p.Nth(i) <= n; i++ {
		fmt.Println("DOING FOR", p.Nth(i))
		prime := p.Nth(i)

		for k := prime; k < n; k += prime {
			fmt.Println("NO MORE", prime, k+1)
			notFounds[k+1]++
		}
	}

	var pairs [][]int

	for nf, cnt := range notFounds {
		pairs = append(pairs, []int{nf, cnt})
	}
	slices.SortFunc(pairs, func(a, b []int) int {
		if a[1] < b[1] {
			return -1
		}
		if a[1] > b[1] {
			return 1
		}
		if a[0] < b[0] {
			return -1
		}
		return 1
	})

	// relevantPrimes := []int{}
	// for i, v := range foundByAll {
	// 	if v && p.Contains(i) {
	// 		relevantPrimes = append(relevantPrimes, i)
	// 	}
	// }
	fmt.Println(pairs)
}

/*
PRNFS ----------
NFS [] [5 13 41 137 149 229 293 397 509 661 677 709 761 809 877 881] 2
[3   7 11    17 19 23 29 31 37    43 47 53 59 61 67 71 73 79 83 89 97 101 103 107 109 113 127 131     139     151 157 163 167 173 179 181 191 193 197 199 211 223 227 233 239 241 251 257 263 269 271 277 281 283 307 311 313 317 331 337 347 349 353 359 367 373 379 383 389 401 409 419 421 431 433 439 443 449 457 461 463 467 479 487 491 499 503 521 523 541 547 557 563 569 571 577 587 593 599 601 607 613 617 619 631 641 643 647 653 659 673 683 691 701 719 727 733 739 743 751 757 769 773 787 797 811 821 823 827 829 839 853 857 859 863 883 887 907 911 919 929 937 941 947 953 967 971 977 983 991 997]
[3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97 101 103 107 109 113 127 131 137 139 149 151 157 163 167 173 179 181 191 193 197 199 211 223 227 229 233 239 241 251 257 263 269 271 277 281 283 293 307 311 313 317 331 337 347 349 353 359 367 373 379 383 389 397 401 409 419 421 431 433 439 443 449 457 461 463 467 479 487 491 499 503 509 521 523 541 547 557 563 569 571 577 587 593 599 601 607 613 617 619 631 641 643 647 653 659 661 673 677 683 691 701 709 719 727 733 739 743 751 757 761 769 773 787 797 809 811 821 823 827 829 839 853 857 859 863 877 881 883 887 907 911 919 929 937 941 947 953 967 971 977 983 991 997]
{4: map[2:2]}
{12: map[2:2 3:1]}
{40: map[2:3 5:1]}
{136: map[2:3 17:1]}
{148: map[2:2 37:1]}
{228: map[2:2 3:1 19:1]}
{292: map[2:2 73:1]}
{396: map[2:2 3:2 11:1]}
{508: map[2:2 127:1]}
{660: map[2:2 3:1 5:1 11:1]}
{676: map[2:2 13:2]}
{708: map[2:2 3:1 59:1]}
{760: map[2:3 5:1 19:1]}
{808: map[2:3 101:1]}
{876: map[2:2 3:1 73:1]}
{880: map[2:4 5:1 11:1]}

NFS [] [19 37 43 73 127 379 433 439 487 757] 3
[7 13 31 61 67 79 97 103 109 139 151 157 163 181 193 199 211 223 229 241 271 277 283 307 313 331 337 349 367 373 397 409 421 457 463 499 523 541 547 571 577 601 607 613 619 631 643 661 673 691 709 727 733 739 751 769 787 811 823 829 853 859 877 883 907 919 937 967 991 997]
[7 13 19 31 37 43 61 67 73 79 97 103 109 127 139 151 157 163 181 193 199 211 223 229 241 271 277 283 307 313 331 337 349 367 373 379 397 409 421 433 439 457 463 487 499 523 541 547 571 577 601 607 613 619 631 643 661 673 691 709 727 733 739 751 757 769 787 811 823 829 853 859 877 883 907 919 937 967 991 997]
{18: map[2:1 3:2]}
{36: map[2:2 3:2]}
{42: map[2:1 3:1 7:1]}
{72: map[2:3 3:2]}
{126: map[2:1 3:2 7:1]}
{378: map[2:1 3:3 7:1]}
{432: map[2:4 3:3]}
{438: map[2:1 3:1 73:1]}
{486: map[2:1 3:5]}
{756: map[2:2 3:3 7:1]}

NFS [] [11 61 571] 5 [31 41 71 101 131 151 181 191 211 241 251 271 281 311 331 401 421 431 461 491 521 541 601 631 641 661 691 701 751 761 811 821 881 911 941 971 991] [11 31 41 61 71 101 131 151 181 191 211 241 251 271 281 311 331 401 421 431 461 491 521 541 571 601 631 641 661 691 701 751 761 811 821 881 911 941 971 991]
{10: map[2:1 5:1]}
{60: map[2:2 3:1 5:1]}
{570: map[2:1 3:1 5:1 19:1]}

NFS [] [29 43 71 281 421 631 659 953] 7 [113 127 197 211 239 337 379 449 463 491 547 617 673 701 743 757 827 883 911 967] [29 43 71 113 127 197 211 239 281 337 379 421 449 463 491 547 617 631 659 673 701 743 757 827 883 911 953 967]
{28: map[2:2 7:1]}
{42: map[2:1 3:1 7:1]}
{70: map[2:1 5:1 7:1]}
{280: map[2:3 5:1 7:1]}
{420: map[2:2 3:1 5:1 7:1]}
{630: map[2:1 3:2 5:1 7:1]}
{658: map[2:1 7:1 47:1]}
{952: map[2:3 7:1 17:1]}

NFS [] [331 397 683] 11 [23 67 89 199 353 419 463 617 661 727 859 881 947] [23 67 89 199 331 353 397 419 463 617 661 683 727 859 881 947]
{330: map[2:1 3:1 5:1 11:1]}
{396: map[2:2 3:2 11:1]}
{682: map[2:1 11:1 31:1]}

NFS [] [79 131 443] 13 [53 157 313 521 547 599 677 859 911 937] [53 79 131 157 313 443 521 547 599 677 859 911 937]
{78: map[2:1 3:1 13:1]}
{130: map[2:1 5:1 13:1]}
{442: map[2:1 13:1 17:1]}

NFS [] [409 613] 17 [103 137 239 307 443 647 919 953] [103 137 239 307 409 443 613 647 919 953]
{408: map[2:3 3:1 17:1]}
{612: map[2:2 3:2 17:1]}

NFS [] [191 229 457] 19 [419 571 647 761] [191 229 419 457 571 647 761]
{190: map[2:1 5:1 19:1]}
{228: map[2:2 3:1 19:1]}
{456: map[2:3 3:1 19:1]}

NFS [] [59] 29 [233 349] [59 233 349]
{58: map[2:1 29:1]}

NFS [] [149] 37 [223 593] [149 223 593]
{148: map[2:2 37:1]}

NFS [] [83] 41 [739 821] [83 739 821]
{82: map[2:1 41:1]}

*/

func loop(n int, eligiblePrimes map[int]bool) {
	// pow := 2
	// k := 1

	// eligiblePrimes := calcedEligible

	// startI := 0
	// for ; p.Nth(startI+1) <= n; startI++ {
	// }
	// fmt.Println(startI, p.Nth(startI), p.Nth(startI+1))

	keys := maps.Keys(eligiblePrimes)
	slices.Sort(keys)

	var del []int

	for _, k := range keys {
		var deleted bool

		fmt.Println("TRYING", k)

		// This may need to loop over n
		// for i := 0; p.Nth(i) <= n; i++ {
		for i := 0; p.Nth(i) <= k; i++ {

			prime := p.Nth(i)

			// Only try primes that are k*n + 1 (TODO: remove this after narrowing down to a few primes)
			if (k-1)%prime != 0 {
				continue
			}

			if !works(prime, k) {
				fmt.Println("DELETING", k, "because of", prime)
				deleted = true
				del = append(del, k)
				break
			}
		}

		if deleted {
			// Remove from eligiblePrimes
			// 	fmt.Println("\n\nREMAINING", prime, time.Now(), eligiblePrimes, len(eligiblePrimes))
			// } else {
			// 	// fmt.Println("DONE WITH", prime)
		}
	}

	for _, d := range del {
		delete(eligiblePrimes, d)
	}

	fmt.Println(eligiblePrimes, calcSum(maps.Keys(eligiblePrimes), n))

	return

	// Below does prime loop, then eligiblePrime loop
	for i := 0; p.Nth(i) <= n; i++ {
		// for i := startI; i >= 0; i-- {
		// for i := 100; i >= 2; i-- {
		prime := p.Nth(i)
		// var r []int
		// for i := 0; p.Nth(i) < 881; i++ {
		// 	prime := p.Nth(i)
		// 	if !works(k, prime) {
		// 		r = append(r, prime)
		// 	}
		// }
		// fmt.Println(k, r)

		var rms []int

		var deleted bool

		for k := range eligiblePrimes {

			// Only try primes that are k*n + 1 (TODO: remove this after narrowing down to a few primes)
			// if (k+1)%prime == 0 {
			// continue
			// }

			if !works(prime, k) {
				rms = append(rms, k)
				fmt.Println("DELETING", k, "because of", prime)
				deleted = true
			}
		}

		for _, k := range rms {
			delete(eligiblePrimes, k)
			fmt.Println("DEL", k)
			// panic("HUZZAH")
		}

		if deleted {
			fmt.Println("\n\nREMAINING", prime, time.Now(), eligiblePrimes, len(eligiblePrimes))
		} else {
			// fmt.Println("DONE WITH", prime)
		}
	}

	fmt.Println(eligiblePrimes)
	// fmt.Println(calcSum(maps.Keys(eligiblePrimes), n))

}

var (
// visited = make([]bool, 10_000_000)
// visitTracker = make([]int, 0, 10_000_000)
)

func works(pow, mod int) bool {
	v := 1
	visited := map[int]bool{}

	for {
		if visited[v] {
			return false
		}

		visited[v] = true

		v = (maths.PowMod(v, pow, mod) + 1) % mod
		if v == 0 {
			return true
		}
	}
}

var (
	calcedEligible = map[int]bool{
		2:       true,
		5:       true,
		137:     true,
		149:     true,
		293:     true,
		509:     true,
		677:     true,
		809:     true,
		1217:    true,
		1277:    true,
		1601:    true,
		2053:    true,
		2633:    true,
		7517:    true,
		8009:    true,
		8117:    true,
		14897:   true,
		20369:   true,
		22349:   true,
		24977:   true,
		26153:   true,
		28433:   true,
		29429:   true,
		30689:   true,
		32189:   true,
		41897:   true,
		45197:   true,
		45833:   true,
		56813:   true,
		57977:   true,
		58337:   true,
		61553:   true,
		64157:   true,
		64793:   true,
		65357:   true,
		65537:   true,
		65789:   true,
		69149:   true,
		91493:   true,
		94613:   true,
		99173:   true,
		105269:  true,
		105977:  true,
		117053:  true,
		127913:  true,
		131297:  true,
		132857:  true,
		175853:  true,
		182333:  true,
		200297:  true,
		202877:  true,
		212297:  true,
		234869:  true,
		244637:  true,
		262193:  true,
		263849:  true,
		287933:  true,
		291173:  true,
		313133:  true,
		318629:  true,
		319469:  true,
		322613:  true,
		324869:  true,
		358349:  true,
		404309:  true,
		404837:  true,
		409817:  true,
		413129:  true,
		436217:  true,
		457277:  true,
		460829:  true,
		468113:  true,
		474017:  true,
		483209:  true,
		495413:  true,
		518717:  true,
		523937:  true,
		547529:  true,
		565289:  true,
		608429:  true,
		613577:  true,
		621113:  true,
		631889:  true,
		633833:  true,
		637709:  true,
		640529:  true,
		657653:  true,
		720569:  true,
		811757:  true,
		824489:  true,
		870329:  true,
		884669:  true,
		959333:  true,
		967529:  true,
		1008353: true,
		1015277: true,
		1028309: true,
		1088933: true,
		1126457: true,
		1200509: true,
		1253729: true,
		1261769: true,
		1265129: true,
		1436957: true,
		1439909: true,
		1461797: true,
		1500833: true,
		1536893: true,
		1563389: true,
		1567169: true,
		1575557: true,
		1588757: true,
		1608473: true,
		1610429: true,
		1623473: true,
		1693577: true,
		1733873: true,
		1784693: true,
		1788653: true,
		1832969: true,
		1836413: true,
		1851809: true,
		1857929: true,
		1858217: true,
		1951253: true,
		1961957: true,
		1962953: true,
		2045909: true,
		2081897: true,
		2118629: true,
		2141417: true,
		2166509: true,
		2184893: true,
		2227109: true,
		2386469: true,
		2417309: true,
		2423777: true,
		2454209: true,
		2470553: true,
		2481317: true,
		2531369: true,
		2812577: true,
		2981057: true,
		2996909: true,
		3076757: true,
		3339113: true,
		3403613: true,
		3494453: true,
		3553469: true,
		3608393: true,
		3705113: true,
		3711353: true,
		3828389: true,
		3854729: true,
		3860033: true,
		3917609: true,
		3928049: true,
		3990449: true,
		3993089: true,
		4012193: true,
		4092593: true,
		4124609: true,
		4145033: true,
		4158893: true,
		4196657: true,
		4222793: true,
		4236437: true,
		4355177: true,
		4377677: true,
		4412729: true,
		4502117: true,
		4541813: true,
		4581389: true,
		4685393: true,
		4951649: true,
		5063453: true,
		5099033: true,
		5485589: true,
		5572949: true,
		5597633: true,
		5644889: true,
		5691689: true,
		5743877: true,
		5889893: true,
		5908337: true,
		5941877: true,
		6022889: true,
		6052373: true,
		6403193: true,
		6416477: true,
		6441389: true,
		6483077: true,
		6580097: true,
		6990317: true,
		7046729: true,
		7152389: true,
		7347533: true,
		7411253: true,
		7431929: true,
		7475609: true,
		7612193: true,
		7923257: true,
		7957253: true,
		8104097: true,
		8135189: true,
		8166737: true,
		8175449: true,
		8203409: true,
		8411009: true,
		8452733: true,
		8496749: true,
		8718953: true,
		8748209: true,
		8833229: true,
		8921453: true,
		9049493: true,
		9157949: true,
		9428033: true,
		9469709: true,
		9627857: true,
		9958973: true,
		9975389: true,
	}
	/*calcedEligible = map[int]bool{
		2:       true,
		5:       true,
		41:      true,
		137:     true,
		149:     true,
		293:     true,
		509:     true,
		677:     true,
		761:     true,
		809:     true,
		881:     true,
		1217:    true,
		1277:    true,
		1601:    true,
		2053:    true,
		2633:    true,
		3701:    true,
		4481:    true,
		5861:    true,
		7121:    true,
		7517:    true,
		8009:    true,
		8117:    true,
		14897:   true,
		20369:   true,
		20441:   true,
		22349:   true,
		24977:   true,
		26153:   true,
		28433:   true,
		29429:   true,
		30689:   true,
		32189:   true,
		34961:   true,
		41897:   true,
		45197:   true,
		45821:   true,
		45833:   true,
		47741:   true,
		56813:   true,
		57977:   true,
		58337:   true,
		61553:   true,
		64157:   true,
		64793:   true,
		65357:   true,
		65537:   true,
		65789:   true,
		69149:   true,
		75641:   true,
		82781:   true,
		91493:   true,
		93941:   true,
		94613:   true,
		99173:   true,
		105269:  true,
		105977:  true,
		110921:  true,
		117053:  true,
		127913:  true,
		131297:  true,
		132857:  true,
		134681:  true,
		148361:  true,
		175853:  true,
		176321:  true,
		182333:  true,
		200297:  true,
		202877:  true,
		212141:  true,
		212297:  true,
		232961:  true,
		234869:  true,
		244637:  true,
		262193:  true,
		263849:  true,
		287933:  true,
		291173:  true,
		313133:  true,
		318629:  true,
		319469:  true,
		322613:  true,
		324869:  true,
		328901:  true,
		358349:  true,
		404309:  true,
		404837:  true,
		409817:  true,
		413129:  true,
		424121:  true,
		424481:  true,
		436217:  true,
		457277:  true,
		460829:  true,
		468113:  true,
		474017:  true,
		483209:  true,
		495413:  true,
		502001:  true,
		502181:  true,
		518717:  true,
		523937:  true,
		547529:  true,
		557261:  true,
		565289:  true,
		576581:  true,
		587621:  true,
		608429:  true,
		613577:  true,
		621113:  true,
		631889:  true,
		633833:  true,
		637709:  true,
		640529:  true,
		657653:  true,
		704861:  true,
		710081:  true,
		720569:  true,
		725321:  true,
		811757:  true,
		824489:  true,
		870329:  true,
		884669:  true,
		959333:  true,
		967529:  true,
		1008353: true,
		1014821: true,
		1015277: true,
		1028309: true,
		1088933: true,
		1126457: true,
		1186181: true,
		1200509: true,
		1253729: true,
		1261769: true,
		1262441: true,
		1265129: true,
		1436957: true,
		1439909: true,
		1461797: true,
		1500833: true,
		1501901: true,
		1532681: true,
		1536893: true,
		1553081: true,
		1563389: true,
		1567169: true,
		1575557: true,
		1588757: true,
		1608473: true,
		1610429: true,
		1623473: true,
		1693361: true,
		1693577: true,
		1733873: true,
		1784693: true,
		1788653: true,
		1832969: true,
		1836413: true,
		1851809: true,
		1857929: true,
		1858217: true,
		1951253: true,
		1961957: true,
		1962953: true,
		2045909: true,
		2081897: true,
		2118629: true,
		2141417: true,
		2166509: true,
		2184893: true,
		2227109: true,
		2386469: true,
		2417309: true,
		2423777: true,
		2454209: true,
		2470553: true,
		2481317: true,
		2531369: true,
		2556161: true,
		2590361: true,
		2755121: true,
		2812577: true,
		2856461: true,
		2908721: true,
		2981057: true,
		2996909: true,
		3076757: true,
		3226001: true,
		3339113: true,
		3403613: true,
		3494453: true,
		3553469: true,
		3608393: true,
		3705113: true,
		3711353: true,
		3828389: true,
		3854729: true,
		3860033: true,
		3910661: true,
		3917609: true,
		3928049: true,
		3990449: true,
		3993089: true,
		4012193: true,
		4028261: true,
		4092593: true,
		4124609: true,
		4145033: true,
		4158893: true,
		4196657: true,
		4222793: true,
		4236437: true,
		4288121: true,
		4355177: true,
		4377677: true,
		4412729: true,
		4438541: true,
		4502117: true,
		4541813: true,
		4581389: true,
		4685393: true,
		4951649: true,
		5063453: true,
		5099033: true,
		5146481: true,
		5341361: true,
		5485589: true,
		5572949: true,
		5582021: true,
		5597633: true,
		5644889: true,
		5691689: true,
		5694341: true,
		5743877: true,
		5889893: true,
		5908337: true,
		5941877: true,
		6022889: true,
		6042521: true,
		6052373: true,
		6403193: true,
		6416477: true,
		6441389: true,
		6483077: true,
		6580097: true,
		6775541: true,
		6839801: true,
		6990317: true,
		7046729: true,
		7098041: true,
		7152389: true,
		7347533: true,
		7411253: true,
		7431929: true,
		7459121: true,
		7475609: true,
		7612193: true,
		7923257: true,
		7957253: true,
		8104097: true,
		8135189: true,
		8166737: true,
		8175449: true,
		8203409: true,
		8411009: true,
		8452733: true,
		8475521: true,
		8496749: true,
		8718953: true,
		8748209: true,
		8833229: true,
		8921453: true,
		9049493: true,
		9074081: true,
		9154961: true,
		9157949: true,
		9287501: true,
		9428033: true,
		9469709: true,
		9622961: true,
		9627857: true,
		9958973: true,
		9975389: true,
	}

	calcedEligible = map[int]bool{
		2:       true,
		5:       true,
		13:      true,
		41:      true,
		137:     true,
		149:     true,
		229:     true,
		293:     true,
		397:     true,
		509:     true,
		661:     true,
		677:     true,
		709:     true,
		761:     true,
		809:     true,
		877:     true,
		881:     true,
		1217:    true,
		1249:    true,
		1277:    true,
		1601:    true,
		2053:    true,
		2633:    true,
		3637:    true,
		3701:    true,
		4481:    true,
		4729:    true,
		5101:    true,
		5449:    true,
		5749:    true,
		5861:    true,
		7121:    true,
		7237:    true,
		7517:    true,
		8009:    true,
		8089:    true,
		8117:    true,
		8377:    true,
		9661:    true,
		14869:   true,
		14897:   true,
		18229:   true,
		19609:   true,
		20369:   true,
		20441:   true,
		21493:   true,
		22349:   true,
		23917:   true,
		24781:   true,
		24977:   true,
		25717:   true,
		26153:   true,
		28433:   true,
		29429:   true,
		30553:   true,
		30689:   true,
		32189:   true,
		34897:   true,
		34961:   true,
		36277:   true,
		40237:   true,
		40693:   true,
		41077:   true,
		41897:   true,
		42709:   true,
		42829:   true,
		42841:   true,
		45197:   true,
		45821:   true,
		45833:   true,
		46261:   true,
		47713:   true,
		47741:   true,
		53593:   true,
		54973:   true,
		56813:   true,
		57601:   true,
		57977:   true,
		58337:   true,
		61553:   true,
		62617:   true,
		64157:   true,
		64793:   true,
		65357:   true,
		65537:   true,
		65789:   true,
		69149:   true,
		75641:   true,
		82009:   true,
		82781:   true,
		86257:   true,
		88069:   true,
		91493:   true,
		93941:   true,
		94613:   true,
		95917:   true,
		99173:   true,
		99241:   true,
		99877:   true,
		105269:  true,
		105977:  true,
		106753:  true,
		110921:  true,
		117053:  true,
		120889:  true,
		123817:  true,
		127913:  true,
		131297:  true,
		132857:  true,
		134681:  true,
		136537:  true,
		141397:  true,
		148361:  true,
		152617:  true,
		175853:  true,
		176321:  true,
		182333:  true,
		200297:  true,
		202877:  true,
		203653:  true,
		212141:  true,
		212297:  true,
		214033:  true,
		220681:  true,
		230221:  true,
		232961:  true,
		234869:  true,
		244637:  true,
		255757:  true,
		262193:  true,
		263849:  true,
		282097:  true,
		285757:  true,
		287933:  true,
		291173:  true,
		296017:  true,
		306949:  true,
		313133:  true,
		318629:  true,
		319469:  true,
		320833:  true,
		322613:  true,
		324869:  true,
		328901:  true,
		341281:  true,
		346201:  true,
		350437:  true,
		358349:  true,
		365689:  true,
		368197:  true,
		369013:  true,
		374929:  true,
		391057:  true,
		401113:  true,
		404309:  true,
		404837:  true,
		409817:  true,
		413129:  true,
		413533:  true,
		424121:  true,
		424481:  true,
		436217:  true,
		443869:  true,
		457277:  true,
		460829:  true,
		468113:  true,
		469717:  true,
		474017:  true,
		476929:  true,
		477409:  true,
		478573:  true,
		483209:  true,
		495413:  true,
		499549:  true,
		502001:  true,
		502181:  true,
		506797:  true,
		518717:  true,
		520381:  true,
		521377:  true,
		523937:  true,
		528097:  true,
		547529:  true,
		551281:  true,
		557261:  true,
		565289:  true,
		572653:  true,
		576581:  true,
		587621:  true,
		608429:  true,
		613577:  true,
		621113:  true,
		622249:  true,
		625969:  true,
		631573:  true,
		631889:  true,
		633833:  true,
		637709:  true,
		640529:  true,
		657653:  true,
		704861:  true,
		710081:  true,
		720569:  true,
		725321:  true,
		760297:  true,
		760657:  true,
		771961:  true,
		789589:  true,
		808177:  true,
		811757:  true,
		824489:  true,
		836761:  true,
		845893:  true,
		864817:  true,
		870329:  true,
		873541:  true,
		884669:  true,
		888109:  true,
		903949:  true,
		919081:  true,
		950233:  true,
		958369:  true,
		959333:  true,
		967529:  true,
		1000849: true,
		1008353: true,
		1014821: true,
		1015277: true,
		1028309: true,
		1029409: true,
		1030741: true,
		1043521: true,
		1046113: true,
		1072381: true,
		1088933: true,
		1119241: true,
		1126457: true,
		1136221: true,
		1141597: true,
		1162981: true,
		1169269: true,
		1179421: true,
		1186181: true,
		1200509: true,
		1208413: true,
		1218829: true,
		1231309: true,
		1253729: true,
		1261769: true,
		1262441: true,
		1265129: true,
		1272961: true,
		1277461: true,
		1285789: true,
		1297477: true,
		1301857: true,
		1368337: true,
		1384921: true,
		1400821: true,
		1426693: true,
		1429969: true,
		1436957: true,
		1437229: true,
		1439909: true,
		1455613: true,
		1461797: true,
		1500833: true,
		1501901: true,
		1532681: true,
		1533901: true,
		1536893: true,
		1540753: true,
		1550413: true,
		1553081: true,
		1563389: true,
		1567169: true,
		1575557: true,
		1583653: true,
		1587109: true,
		1588757: true,
		1608473: true,
		1610429: true,
		1612609: true,
		1623473: true,
		1631341: true,
		1693361: true,
		1693577: true,
		1733873: true,
		1777093: true,
		1784693: true,
		1788653: true,
		1832969: true,
		1836413: true,
		1851809: true,
		1855621: true,
		1857929: true,
		1858217: true,
		1893289: true,
		1936633: true,
		1948861: true,
		1951253: true,
		1961957: true,
		1962953: true,
		2005981: true,
		2039929: true,
		2045909: true,
		2081897: true,
		2118629: true,
		2141417: true,
		2166509: true,
		2184893: true,
		2226769: true,
		2227109: true,
		2300317: true,
		2300869: true,
		2336629: true,
		2386469: true,
		2387533: true,
		2412721: true,
		2417309: true,
		2417341: true,
		2423777: true,
		2437741: true,
		2438629: true,
		2442889: true,
		2445253: true,
		2454209: true,
		2455909: true,
		2470553: true,
		2481317: true,
		2493229: true,
		2531369: true,
		2546317: true,
		2556161: true,
		2590361: true,
		2601817: true,
		2638381: true,
		2650537: true,
		2675833: true,
		2739013: true,
		2755121: true,
		2799409: true,
		2812577: true,
		2821501: true,
		2837209: true,
		2856461: true,
		2908721: true,
		2943937: true,
		2981057: true,
		2996909: true,
		3070213: true,
		3073813: true,
		3076757: true,
		3142717: true,
		3198109: true,
		3208837: true,
		3226001: true,
		3281533: true,
		3339113: true,
		3384049: true,
		3403613: true,
		3486169: true,
		3494453: true,
		3553469: true,
		3595117: true,
		3608393: true,
		3705113: true,
		3711353: true,
		3716701: true,
		3815521: true,
		3828389: true,
		3854729: true,
		3860033: true,
		3880861: true,
		3910661: true,
		3917609: true,
		3919621: true,
		3928049: true,
		3940777: true,
		3990449: true,
		3993089: true,
		4012193: true,
		4021249: true,
		4028261: true,
		4041229: true,
		4092593: true,
		4124609: true,
		4145033: true,
		4158893: true,
		4185673: true,
		4196657: true,
		4222793: true,
		4236437: true,
		4244017: true,
		4262173: true,
		4266289: true,
		4288121: true,
		4296673: true,
		4302829: true,
		4341229: true,
		4355177: true,
		4377677: true,
		4408429: true,
		4412729: true,
		4438541: true,
		4444549: true,
		4476517: true,
		4502117: true,
		4520101: true,
		4527361: true,
		4541813: true,
		4581389: true,
		4592857: true,
		4594657: true,
		4685393: true,
		4716289: true,
		4720393: true,
		4729369: true,
		4747357: true,
		4798033: true,
		4825813: true,
		4827037: true,
		4829497: true,
		4859629: true,
		4951649: true,
		4993837: true,
		5041513: true,
		5058577: true,
		5063453: true,
		5099033: true,
		5146481: true,
		5191861: true,
		5222653: true,
		5341361: true,
		5377081: true,
		5405077: true,
		5410441: true,
		5439613: true,
		5485589: true,
		5502397: true,
		5538289: true,
		5572949: true,
		5582021: true,
		5597633: true,
		5644889: true,
		5691689: true,
		5694341: true,
		5725849: true,
		5728621: true,
		5731681: true,
		5734909: true,
		5743877: true,
		5768701: true,
		5778229: true,
		5809633: true,
		5889893: true,
		5908337: true,
		5941877: true,
		5963653: true,
		6022889: true,
		6042409: true,
		6042521: true,
		6052373: true,
		6130429: true,
		6306829: true,
		6326641: true,
		6348037: true,
		6387853: true,
		6402973: true,
		6403193: true,
		6413857: true,
		6416477: true,
		6430201: true,
		6441389: true,
		6469357: true,
		6483077: true,
		6486973: true,
		6515881: true,
		6540409: true,
		6580097: true,
		6596257: true,
		6775541: true,
		6811117: true,
		6839801: true,
		6867001: true,
		6928513: true,
		6984793: true,
		6985789: true,
		6990317: true,
		7046729: true,
		7097689: true,
		7098041: true,
		7152389: true,
		7333741: true,
		7347533: true,
		7389637: true,
		7411253: true,
		7431929: true,
		7435873: true,
		7459121: true,
		7475609: true,
		7505341: true,
		7506241: true,
		7612193: true,
		7680181: true,
		7754437: true,
		7760089: true,
		7760497: true,
		7820149: true,
		7827493: true,
		7923257: true,
		7957253: true,
		7993201: true,
		8022877: true,
		8104097: true,
		8110681: true,
		8135189: true,
		8166737: true,
		8175449: true,
		8203409: true,
		8387089: true,
		8411009: true,
		8414869: true,
		8452733: true,
		8457217: true,
		8475521: true,
		8496749: true,
		8594557: true,
		8718953: true,
		8748209: true,
		8833229: true,
		8921453: true,
		9000961: true,
		9049493: true,
		9074081: true,
		9152557: true,
		9154961: true,
		9157949: true,
		9216853: true,
		9217837: true,
		9287501: true,
		9316129: true,
		9317089: true,
		9369229: true,
		9428033: true,
		9431797: true,
		9469709: true,
		9561961: true,
		9622961: true,
		9627857: true,
		9672973: true,
		9678073: true,
		9797329: true,
		9878353: true,
		9958973: true,
		9970273: true,
		9975389: true,
	}*/
)

/*
REMAINING 3373 2025-02-02 06:20:43.1950346 -0500 EST m=+23992.234665401 map[2, 5, 149, 293, 1601, 41897, 45197, 57977, 58337, 61553, 65357, 65537, 65789, 99173, 105269, 132857, 175853, 200297, 287933, 313133, 318629, 319469, 404837, 457277, 474017, 495413, 613577, 633833, 657653, 720569, 870329, 967529, 1008353, 1015277, 1028309, 1261769, 1265129, 1500833, 1623473, 2045909, 2118629, 2141417, 2423777, 2481317, 3711353, 3828389, 4145033, 4502117, 4685393, 5063453, 5099033, 5485589, 5941877, 6022889, 6052373, 6403193, 6441389, 6483077, 6580097, 6990317, 7152389, 8135189, 8203409, 9469709,]
*/

// Finished vanilla search up to 571
