package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/leep-frog/command"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func TestYears(t *testing.T) {
	keys := maps.Keys(years)
	slices.Sort(keys)
	for _, y := range keys {
		if y != 2015 {
			continue
		}
		year := years[y]
		for dayNumber, day := range year.Days {
			for _, cse := range day.Cases() {
				t.Run(fmt.Sprintf("%d.%d %s", year.Number, dayNumber+1, cse.FileSuffix), func(t *testing.T) {
					var wantOutput string
					if strings.Join(cse.ExpectedOutput, "") != "" {
						wantOutput = strings.Join(cse.ExpectedOutput, "\n") + "\n"
					}
					args := []string{
						fmt.Sprintf("%d", year.Number),
						fmt.Sprintf("%d", dayNumber+1),
						"--suffix",
						cse.FileSuffix,
					}
					command.ExecuteTest(t, &command.ExecuteTestCase{
						Node:       node(),
						Args:       args,
						WantStdout: wantOutput,
						WantData: &command.Data{Values: map[string]interface{}{
							yearArg.Name():    year,
							dayArg.Name():     dayNumber + 1,
							suffixFlag.Name(): cse.FileSuffix,
						}},
					})
				})
			}
		}
	}
}
