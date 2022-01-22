package twentyone

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/parse"
)

const (
	downCuc  = 1
	rightCuc = 2
)

func D25() *problem {
	return command.SerialNodes(
		command.Description(""),
		command.ExecutorNode(func(o command.Output, d *command.Data) error {
			var grid [][]int
			for _, line := range parse.ReadFileLines(filepath.ToSlash(filepath.Join("input", "d25_google.txt"))) {
				var row []int
				for i := 0; i < len(line); i++ {
					if c := line[i : i+1]; c == "v" {
						row = append(row, downCuc)
					} else if c == ">" {
						row = append(row, rightCuc)
					} else {
						row = append(row, 0)
					}
				}
				grid = append(grid, row)
			}

			fmt.Println(grid)

			height := len(grid)
			width := len(grid[0])

			printCucGrid(grid)
			//return nil

			gridChange := true
			count := 0
			moveSet := map[int]map[int]int{}
			for ; gridChange; count++ {
				gridChange = false

				// Right cucumbers
				for i, row := range grid {
					for j, cell := range row {
						if cell == rightCuc {
							if newJ := (j + 1) % width; grid[i][newJ] == 0 {
								addMS(moveSet, i, j, 0)
								addMS(moveSet, i, newJ, rightCuc)
								gridChange = true
							}
						}
					}
				}

				for i, m := range moveSet {
					for j, v := range m {
						grid[i][j] = v
					}
				}
				moveSet = map[int]map[int]int{}

				// Down cucumbers
				for i, row := range grid {
					for j, cell := range row {
						if cell == downCuc {
							if newI := (i + 1) % height; grid[newI][j] == 0 {
								addMS(moveSet, i, j, 0)
								addMS(moveSet, newI, j, downCuc)
								gridChange = true
							}
						}
					}
				}

				for i, m := range moveSet {
					for j, v := range m {
						grid[i][j] = v
					}
				}
				moveSet = map[int]map[int]int{}

			}

			fmt.Println("================")
			printCucGrid(grid)
			fmt.Println(count)
			return nil
		}),
	)
}

func addMS(ms map[int]map[int]int, i, j, v int) {
	if ms[i] == nil {
		ms[i] = map[int]int{}
	}
	ms[i][j] = v
}

func printCucGrid(grid [][]int) {
	for _, row := range grid {
		for _, cell := range row {
			if cell == rightCuc {
				fmt.Printf(">")
			} else if cell == downCuc {
				fmt.Printf("v")
			} else if cell == 0 {
				fmt.Printf(".")
			} else {
				log.Fatalf("wut: %d", cell)
			}
		}
		fmt.Println()
	}
}
