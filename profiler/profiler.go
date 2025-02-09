package profiler

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"golang.org/x/exp/maps"
)

type Profiler struct {
	time         map[string]int64
	counts       map[string]int64
	current      string
	currentStart time.Time
}

func New() *Profiler {
	return &Profiler{
		time:   map[string]int64{},
		counts: map[string]int64{},
	}
}

func (p *Profiler) Start(code string) {
	if p.current != "" {
		p.End()
	}

	p.current = code
	p.currentStart = time.Now()
}

func (p *Profiler) End() {
	if p.current == "" {
		panic("Ending stopped time")
	}
	p.time[p.current] += (time.Now().UnixMilli() - p.currentStart.UnixMilli())
	p.counts[p.current]++
}

func (p *Profiler) String() string {
	keys := maps.Keys(p.time)
	slices.Sort(keys)

	var r []string
	for _, k := range keys {
		millis := p.time[k]
		counts := p.counts[k]
		r = append(r, fmt.Sprintf("%s: %d.%ds; calls: %d; ixns/s: %0.5f", k, millis/1000, millis%1000, counts, 1000*float64(counts)/float64(millis)))
	}
	return strings.Join(r, "\n")
}

/*
PLUSING: 0.504s; calls: 115604; ixns/s: 229373.01587
RECIP: 0.1s; calls: 196; ixns/s: 196000.00000
TIMES RECIP: 0.22s; calls: 196; ixns/s: 8909.09091
TIMES SOLVE COEF: 0.415s; calls: 3467; ixns/s: 8354.21687

AFTER PLUSING: 2.190s; calls: 2043097; ixns/s: 932921.00457
APPLY IN PLACE: 0.315s; calls: 157080; ixns/s: 498666.66667
DELETE: 0.53s; calls: 44879; ixns/s: 846773.58491
GET FIRST: 0.20s; calls: 561; ixns/s: 28050.00000
GET SOLN: 0.2s; calls: 1121; ixns/s: 560500.00000
KEY ITER: 2.37s; calls: 44318; ixns/s: 21756.50466
PLUS ADD EQ: 0.1s; calls: 22159; ixns/s: 22159000.00000
PLUSING: 9.187s; calls: 2043097; ixns/s: 222390.00762
RECIP: 0.4s; calls: 561; ixns/s: 140250.00000
TIMES RECIP: 0.118s; calls: 561; ixns/s: 4754.23729
TIMES SOLVE COEF: 8.293s; calls: 22159; ixns/s: 2672.01254
VAR CNT MAP: 0.838s; calls: 562; ixns/s: 670.64439
VAR ORDER: 0.21s; calls: 562; ixns/s: 26761.90476
VAR SORT: 0.207s; calls: 562; ixns/s: 2714.97585

PLUSING: 361.32s; calls: 25095006; ixns/s: 69509.09061
TIMES SOLVE COEF: 308.959s; calls: 116207; ixns/s: 376.12434
AFTER PLUSING: 24.66s; calls: 25095006; ixns/s: 1042757.66642
APPLY IN PLACE: 3.608s; calls: 935028; ixns/s: 259154.10200
DELETE: 0.319s; calls: 233782; ixns/s: 732858.93417
GET FIRST: 0.92s; calls: 1368; ixns/s: 14869.56522
GET SOLN: 0.24s; calls: 2735; ixns/s: 113958.33333
KEY ITER: 17.927s; calls: 232414; ixns/s: 12964.46701
PLUS ADD EQ: 0.20s; calls: 116207; ixns/s: 5810350.00000
RECIP: 0.12s; calls: 1368; ixns/s: 114000.00000
TIMES RECIP: 1.876s; calls: 1368; ixns/s: 729.21109
VAR CNT MAP: 6.798s; calls: 1369; ixns/s: 201.38276
VAR ORDER: 0.59s; calls: 1369; ixns/s: 23203.38983
VAR SORT: 1.12s; calls: 1369; ixns/s: 1352.76680
*/
