package main

import (
	"fmt"
	"testing"

	"github.com/leep-frog/command"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func TestYears(t *testing.T) {
	keys := maps.Keys(years)
	slices.Sort(keys)
	for _, y := range keys {
		year := years[y]
		for dayNumber, day := range year.Days {
			for _, cse := range day.Cases() {
				args := []string{
					fmt.Sprintf("%d", year),
					fmt.Sprintf("%d", dayNumber+1),
					"--suffix",
					cse.FileSuffix,
				}
				command.ExecuteTest(t, &command.ExecuteTestCase{
					Node:       node(),
					Args:       args,
					WantStdout: cse.ExpectedOutput + "\n",
				})
			}
		}
	}
}
