package aoccmd

import (
	"fmt"
	"strings"
	"testing"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/functional"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func TestYears(t *testing.T) {
	n := CLI().Node()

	keys := functional.Filter(maps.Keys(years), func(y int) bool { return y > 2000 })
	slices.Sort(keys)

	for _, y := range keys {
		year := years[y]
		if year.Number != 2016 {
			continue
		}
		for dayNumber, day := range year.Days {
			dayNumber++
			for _, cse := range day.Cases() {
				t.Run(fmt.Sprintf("%d.%d %s", year.Number, dayNumber, cse.FileSuffix), func(t *testing.T) {
					var wantOutput string
					if strings.Join(cse.ExpectedOutput, "") != "" {
						wantOutput = strings.Join(cse.ExpectedOutput, "\n") + "\n"
					}
					args := []string{
						fmt.Sprintf("%d", year.Number),
						fmt.Sprintf("%d", dayNumber),
						"--suffix",
						cse.FileSuffix,
					}
					command.ExecuteTest(t, &command.ExecuteTestCase{
						Node:       n,
						Args:       args,
						WantStdout: wantOutput,
						WantData: &command.Data{Values: map[string]interface{}{
							"YEAR":   year,
							"DAY":    dayNumber,
							"suffix": cse.FileSuffix,
						}},
					})
				})
			}
		}
	}
}
