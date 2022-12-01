package eulerchallenge

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

type UlamSequence struct {
	a, b  int
	vs    []int
	diffs []int
	// TODO: make slice?
	contains []bool
	m        map[int]int
	mod      int
	modMap   [][]int
}

func NewUlam(a, b int) *UlamSequence {
	c := make([]bool, b+1, b+1)
	c[a] = true
	c[b] = true
	mod := 1234
	modMap := make([][]int, mod, mod)
	modMap[a%mod] = []int{a}
	modMap[b%mod] = []int{b}
	return &UlamSequence{
		a,
		b,
		[]int{a, b},
		[]int{b - a},
		c,
		map[int]int{a + b: 1},
		mod,
		modMap,
	}
}

func (u *UlamSequence) generateNext() int {
	for i := u.vs[len(u.vs)-1] + 1; ; i++ {
		cnt := 0
		wantMod := i % u.mod
		for modA := 0; modA < u.mod; modA++ {
			modB := (wantMod - modA + u.mod) % u.mod
			if modB < modA {
				continue
			}
			for _, a := range u.modMap[modA] {
				for _, b := range u.modMap[modB] {
					if modA == modB && b < a {
						continue
					}
					if a == b {
						continue
					}
					if a+b == i {
						cnt++
					}
					// bs are sorted in increasing order
					if a+b > i {
						break
					}
					if cnt == 2 {
						goto DONZO
					}
				}
			}
		}
	DONZO:
		if cnt == 1 {
			u.diffs = append(u.diffs, i-u.vs[len(u.vs)-1])
			u.vs = append(u.vs, i)
			u.modMap[wantMod] = append(u.modMap[wantMod], i)
			return i
		}
	}
}

func (u *UlamSequence) midGenerateNext() int {

	for i := u.vs[len(u.vs)-1] + 1; ; i++ {
		cnt := 0
		for j := 0; u.vs[j]*2 < i; j++ {
			if u.contains[i-u.vs[j]] {
				cnt++
				if cnt == 2 {
					break
				}
			}
		}
		if cnt == 1 {
			u.diffs = append(u.diffs, i-u.vs[len(u.vs)-1])
			u.vs = append(u.vs, i)
			return i
		}
	}
}

func (u *UlamSequence) oldgenerateNext() int {
	for i := u.vs[len(u.vs)-1]; ; i++ {
		cnt := u.m[i]
		delete(u.m, i)
		if cnt == 1 {
			for _, v := range u.vs {
				u.m[v+i]++
			}
			u.diffs = append(u.diffs, i-u.vs[len(u.vs)-1])
			u.vs = append(u.vs, i)
			return i
		}
	}
}

func (u *UlamSequence) at(i int) int {
	for i >= len(u.vs) {
		u.generateNext()
	}
	return u.vs[i]
}

func (u *UlamSequence) kth(k int) int {
	evenCount := 0
	var start int
	for evenCount < 2 {
		if u.at(start)%2 == 0 {
			evenCount++
		}
		start++
	}

	for patternLength := 25; ; patternLength++ {
		if patternLength%10_000 == 0 {
			fmt.Println("PL", u.b, patternLength)
		}
		valid := true
		sum := 0
		for i := 0; i < patternLength; i++ {
			u.at(start + patternLength*2)
			sum += u.diffs[start+i]
			if u.diffs[start+i] != u.diffs[start+i+patternLength] {
				valid = false
				break
			}
		}

		if valid {
			value := u.at(start) + ((k-start)/patternLength)*sum
			for i := ((k - start) % patternLength); i > 0; i-- {
				value += u.at(start+i) - u.at(start+i-1)
			}
			return value
		}
	}
}

func P167() *problem {
	return intInputNode(167, func(o command.Output, n int) {

		fmt.Println("START")
		// fmt.Println("U", NewUlam(2, 5).generateNext())
		fmt.Println("U", NewUlam(2, 5).at(50))
		// return

		sum := 0
		for k := 2; k <= 10; k++ {
			v := NewUlam(2, 2*k+1).kth(maths.Pow(10, 11) - 1)
			fmt.Println(k, v)
			sum += v
		}
		fmt.Println(sum)

		return
	}, []*execution{
		{
			args: []string{"1"},
			want: "",
		},
	})
}

/*
START
U 177
2 393749999981
3 484615384605
4 400450450395
5 399877149781
6 399966136001
7 637499999951
PL 17 10000
PL 17 20000
PL 17 30000
PL 17 40000
PL 17 50000
PL 17 60000
PL 17 70000
PL 17 80000
PL 17 90000
PL 17 100000
PL 17 110000
PL 17 120000
8 400001574629
PL 19 10000
PL 19 20000
PL 19 30000
PL 19 40000
PL 19 50000
PL 19 60000
PL 19 70000
PL 19 80000
PL 19 90000
PL 19 100000
PL 19 110000
PL 19 120000
PL 19 130000
PL 19 140000
PL 19 150000
PL 19 160000
PL 19 170000
PL 19 180000
PL 19 190000
PL 19 200000
PL 19 210000
PL 19 220000
PL 19 230000
PL 19 240000
PL 19 250000
PL 19 260000
PL 19 270000
PL 19 280000
PL 19 290000
PL 19 300000
PL 19 310000
PL 19 320000
PL 19 330000
PL 19 340000
PL 19 350000
PL 19 360000
PL 19 370000
PL 19 380000
9 399999473477
PL 21 10000
PL 21 20000
PL 21 30000
PL 21 40000
PL 21 50000
PL 21 60000
PL 21 70000
PL 21 80000
PL 21 90000
PL 21 100000
PL 21 110000
PL 21 120000
PL 21 130000
PL 21 140000
PL 21 150000
PL 21 160000
PL 21 170000
PL 21 180000
PL 21 190000
PL 21 200000
PL 21 210000
PL 21 220000
PL 21 230000
PL 21 240000
PL 21 250000
PL 21 260000
PL 21 270000
PL 21 280000
PL 21 290000
PL 21 300000
PL 21 310000
PL 21 320000
PL 21 330000
PL 21 340000
PL 21 350000
PL 21 360000
PL 21 370000
PL 21 380000
PL 21 390000
PL 21 400000
PL 21 410000
PL 21 420000
PL 21 430000
PL 21 440000
PL 21 450000
PL 21 460000
PL 21 470000
PL 21 480000
PL 21 490000
PL 21 500000
PL 21 510000
PL 21 520000
PL 21 530000
PL 21 540000
PL 21 550000
PL 21 560000
PL 21 570000
PL 21 580000
PL 21 590000
PL 21 600000
PL 21 610000
PL 21 620000
PL 21 630000
PL 21 640000
PL 21 650000
PL 21 660000
PL 21 670000
PL 21 680000
PL 21 690000
PL 21 700000
PL 21 710000
PL 21 720000
PL 21 730000
PL 21 740000
PL 21 750000
PL 21 760000
PL 21 770000
PL 21 780000
PL 21 790000
PL 21 800000
PL 21 810000
PL 21 820000
PL 21 830000
PL 21 840000
PL 21 850000
PL 21 860000
PL 21 870000
PL 21 880000
PL 21 890000
PL 21 900000
PL 21 910000
PL 21 920000
PL 21 930000
PL 21 940000
PL 21 950000
PL 21 960000
PL 21 970000
PL 21 980000
PL 21 990000
PL 21 1000000
PL 21 1010000
PL 21 1020000
PL 21 1030000
PL 21 1040000
PL 21 1050000
PL 21 1060000
PL 21 1070000
PL 21 1080000
PL 21 1090000
PL 21 1100000
PL 21 1110000
PL 21 1120000
PL 21 1130000
PL 21 1140000
PL 21 1150000
PL 21 1160000
PL 21 1170000
PL 21 1180000
PL 21 1190000
PL 21 1200000
PL 21 1210000
PL 21 1220000
PL 21 1230000
PL 21 1240000
PL 21 1250000
PL 21 1260000
PL 21 1270000
PL 21 1280000
PL 21 1290000
PL 21 1300000
PL 21 1310000
PL 21 1320000
PL 21 1330000
PL 21 1340000
PL 21 1350000
PL 21 1360000
PL 21 1370000
PL 21 1380000
PL 21 1390000
PL 21 1400000
PL 21 1410000
PL 21 1420000
PL 21 1430000
PL 21 1440000
PL 21 1450000
PL 21 1460000
PL 21 1470000
PL 21 1480000
PL 21 1490000
PL 21 1500000
PL 21 1510000
PL 21 1520000
PL 21 1530000
PL 21 1540000
PL 21 1550000
PL 21 1560000
PL 21 1570000
PL 21 1580000
PL 21 1590000
PL 21 1600000
PL 21 1610000
PL 21 1620000
PL 21 1630000
PL 21 1640000
PL 21 1650000
PL 21 1660000
PL 21 1670000
PL 21 1680000
PL 21 1690000
PL 21 1700000
PL 21 1710000
PL 21 1720000
PL 21 1730000
PL 21 1740000
PL 21 1750000
PL 21 1760000
PL 21 1770000
PL 21 1780000
PL 21 1790000
PL 21 1800000
PL 21 1810000
PL 21 1820000
PL 21 1830000
PL 21 1840000
PL 21 1850000
PL 21 1860000
PL 21 1870000
PL 21 1880000
PL 21 1890000
PL 21 1900000
PL 21 1910000
PL 21 1920000
PL 21 1930000
PL 21 1940000
PL 21 1950000
PL 21 1960000
PL 21 1970000
PL 21 1980000
PL 21 1990000
PL 21 2000000
PL 21 2010000
PL 21 2020000
PL 21 2030000
PL 21 2040000
PL 21 2050000
PL 21 2060000
PL 21 2070000
PL 21 2080000
PL 21 2090000
10 399999900065
3916160068885
*/
