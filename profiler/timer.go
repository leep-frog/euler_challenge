package profiler

import (
	"fmt"
	"time"
)

type Timer struct {
	c     chan bool
	start time.Time
}

func NewTimer() *Timer {
	return &Timer{nil, time.Now()}
}

func (t *Timer) End() {
	t.c <- true
}

func (t *Timer) Start() {
	t.c = make(chan bool)
	t.start = time.Now()
	go func() {
		tr := time.NewTimer(time.Second)
		for {
			select {
			case <-tr.C:
				d := time.Now().Sub(t.start)
				fmt.Printf("%dm:%02ds\r", int(d.Minutes()), int(d.Seconds())%60)
				tr.Reset(time.Second)
			case <-t.c:
				goto END_TIMER
			}
		}

	END_TIMER:
		if !tr.Stop() {
			<-tr.C
		}
		close(t.c)
	}()
}
