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
	current      string
	currentStart time.Time
}

func New() *Profiler {
	return &Profiler{
		time: map[string]int64{},
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
}

func (p *Profiler) String() string {
	keys := maps.Keys(p.time)
	slices.Sort(keys)

	var r []string
	for _, k := range keys {
		millis := p.time[k]
		r = append(r, fmt.Sprintf("%s: %d.%ds", k, millis/1000, millis%1000))
	}
	return strings.Join(r, "\n")
}
