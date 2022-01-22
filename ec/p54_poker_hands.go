package eulerchallenge

import (
	"log"
	"sort"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P54() *problem {
	return fileInputNode(54, func(lines []string, o command.Output) {
		var count int
		for _, line := range lines {
			cards := strings.Split(line, " ")
			this, that := newHand(cards[:5]), newHand(cards[5:])
			if this.beats(that) {
				count++
			}
		}
		o.Stdoutln(count)
	})
}

var (
	suitMap = map[string]int{
		"C": 0,
		"D": 1,
		"H": 2,
		"S": 3,
	}

	valueMap = map[string]int{
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"J": 11,
		"Q": 12,
		"K": 13,
		"A": 14,
	}
)

type card struct {
	value int
	suit  int
}

func loadCard(s string) *card {
	v, ok := valueMap[s[0:1]]
	if !ok {
		log.Fatalf("Invalid card value: %d : %v", v, s)
	}
	suit, ok := suitMap[s[1:2]]
	if !ok {
		log.Fatalf("Invalid card suit: %d", suit)
	}
	return &card{v, suit}
}

type hand struct {
	cards []*card
	str   string
}

func (h *hand) String() string {
	return h.str
}

func newHand(sl []string) *hand {
	h := &hand{
		str: strings.Join(sl, " "),
	}
	for _, s := range sl {
		h.cards = append(h.cards, loadCard(s))
	}
	return h
}

func (h *hand) flushScore() int {
	suit := h.cards[0].suit
	best := maths.Largest()
	for _, c := range h.cards {
		if c.suit != suit {
			return 0
		}
		best.Check(c.value)
	}
	return best.Best()
}

func (h *hand) straightScore(flush bool) int {
	// Map from value to suit to bool
	hasValues := map[int]map[int]bool{}
	for _, c := range h.cards {
		if hasValues[c.value] == nil {
			hasValues[c.value] = map[int]bool{}
		}
		hasValues[c.value][c.suit] = true
		if c.value == 14 {
			if hasValues[1] == nil {
				hasValues[1] = map[int]bool{}
			}
			hasValues[1][c.suit] = true
		}
	}

	best := maths.Largest()
	var count int
	suitCounts := make([]int, len(suitMap))
	for i := 14; i >= 1; i-- {
		hasSuits, ok := hasValues[i]

		// Check regular straight
		if !flush {
			if ok {
				count++
			} else {
				count = 0
			}
			if count >= 5 {
				best.Check(i + 4)
			}
		} else {
			for i := range suitCounts {
				if hasSuits[i] {
					suitCounts[i]++
				} else {
					suitCounts[i] = 0
				}

				if suitCounts[i] >= 5 {
					best.Check(100 * (i + 4))
				}
			}
		}
	}
	return best.Best()
}

func (h *hand) counts() map[int][]int {
	vToCount := map[int]int{}
	for _, c := range h.cards {
		vToCount[c.value]++
	}

	countToV := map[int][]int{}
	for k, v := range vToCount {
		countToV[v] = append(countToV[v], k)
	}

	for k := range countToV {
		sort.Sort(sort.Reverse(sort.IntSlice(countToV[k])))
	}
	return countToV
}

func (h *hand) fourOfAKind() int {
	c := h.counts()
	if len(c[4]) > 0 {
		return c[4][0]
	}
	return 0
}

func (h *hand) fullHouse() int {
	c := h.counts()
	if len(c[3]) > 0 {
		if len(c[2]) > 0 {
			return 100*c[3][0] + c[2][0]
		}
	}
	return 0
}

func (h *hand) pairs() int {
	c := h.counts()
	if len(c[2]) > 1 {
		return 100*c[2][0] + c[2][1]
	} else if len(c[2]) > 0 {
		return c[2][0]
	}
	return 0
}

func (h *hand) highCard() int {
	best := maths.Largest()
	for _, c := range h.cards {
		best.Check(c.value)
	}
	return best.Best()
}

func (h *hand) threeOfAKind() int {
	c := h.counts()
	if len(c[3]) > 0 {
		return c[3][0]
	}
	return 0
}

func (h *hand) betterValue(that *hand) bool {
	var hs, ts []int
	for _, c := range h.cards {
		hs = append(hs, c.value)
	}
	for _, c := range that.cards {
		ts = append(ts, c.value)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(hs)))
	sort.Sort(sort.Reverse(sort.IntSlice(ts)))
	for i := 0; i < 5; i++ {
		if hs[i] != ts[i] {
			return hs[i] > ts[i]
		}
	}
	log.Fatalf("Shouldn't end in tie")
	return false
}

func checkHand(bv bool, a, b int) (bool, bool) {
	if a <= 0 && b <= 0 {
		return false, false
	}
	if a == b {
		return bv, true
	}
	return a > b, true
}

type cardCmp struct {
	name string
	f    func(*hand) int
}

func (h *hand) beats(thatH *hand) bool {
	bv := h.betterValue(thatH)

	cmps := []*cardCmp{
		// Straight Flush
		{"StFl", func(d *hand) int { return d.straightScore(true) }},
		// Four of a kind
		{"Four", func(d *hand) int { return d.fourOfAKind() }},
		// Full house
		{"FlHs", func(d *hand) int { return d.fullHouse() }},
		// Flush
		{"Flus", func(d *hand) int { return d.flushScore() }},
		// Straight
		{"Stra", func(d *hand) int { return d.straightScore(false) }},
		// Three of a kind
		{"ToaK", func(d *hand) int { return d.threeOfAKind() }},
		// Pairs
		{"Pair", func(d *hand) int { return d.pairs() }},
	}

	for _, cmp := range cmps {
		better, ok := checkHand(bv, cmp.f(h), cmp.f(thatH))
		if ok {
			return better
		}
	}

	// High card
	return bv
}
