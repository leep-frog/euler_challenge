package profiler

import "fmt"

func Progress(prefix string, idx, total, increment int) {
	if idx%increment == 0 {
		fmt.Printf("%s: (%d/%d)\r", prefix, idx, total)
	}
}
