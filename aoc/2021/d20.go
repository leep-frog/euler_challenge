package twentyone

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/point"
)

func D19() *problem {
	return command.SerialNodes(
		command.Description(""),
		command.ExecutorNode(func(o command.Output, d *command.Data) error {

			// Load scanners
			var scanners []*Scanner
			curScanner := &Scanner{
				loc: point.NewPoint(0, 0, 0),
			}
			//lines := eulerchallenge.ReadFileLines(filepath.ToSlash(filepath.Join("..", "aoc", "2021", "input", "d19_example.txt")))
			lines := parse.ReadFileLines(filepath.ToSlash(filepath.Join("..", "aoc", "2021", "input", "d19.txt")))
			//lines := eulerchallenge.ReadFileLines(filepath.ToSlash(filepath.Join("..", "aoc", "2021", "input", "d19_2.txt")))
			for i := 1; i < len(lines); i++ {
				line := lines[i]
				if line == "" {
					i++
					scanners = append(scanners, curScanner)
					curScanner = &Scanner{
						loc: point.NewPoint(0, 0, 0),
					}
					continue
				}

				numStrs := strings.Split(line, ",")
				var nums []int
				for _, ns := range numStrs {
					nums = append(nums, parse.Atoi(ns))
				}
				curScanner.Append(nums)
			}

			ique := &Scanner{}
			for _, scanner := range scanners {
				for _, p := range scanner.points {
					ique.AppendPoint(p)
				}
			}
			fmt.Println(len(ique.points))

			absoluteScanners := []int{
				0,
			}
			inAs := map[int]bool{
				0: true,
			}
			checked := map[int]map[int]bool{}
			// Iterate through pairs of scanners
			// N
			for i := 0; i < len(scanners); i++ {
				ai := absoluteScanners[i]
				// N * N
				for j := 0; j < len(scanners); j++ {
					if j >= len(scanners) {
						break
					}
					if checked[ai][j] || checked[j][ai] {
						continue
					}
					if ai == j {
						continue
					}
					if checked[ai] == nil {
						checked[ai] = map[int]bool{}
					}
					if checked[j] == nil {
						checked[j] = map[int]bool{}
					}
					checked[ai][j] = true
					checked[j][ai] = true
					//fmt.Println("checking", i, ai, j)
					if checkIntersect(scanners, ai, j) && !inAs[j] {
						absoluteScanners = append(absoluteScanners, j)
						inAs[j] = true
					}
				}
			}

			unique := &Scanner{}
			for _, scanner := range scanners {
				for _, p := range scanner.points {
					unique.AppendPoint(p)
				}
			}

			fmt.Println("DONE", len(scanners), len(absoluteScanners))
			/*scanners[0].Sort()
			fmt.Println(scanners[0])
			fmt.Println(len(scanners[0].points))*/
			unique.Sort()
			sort.Ints(absoluteScanners)
			fmt.Println(absoluteScanners)
			//fmt.Println(unique)
			fmt.Println(len(unique.points))

			fmt.Println("PART 2==========")
			var max int
			for i := 0; i < len(scanners); i++ {
				iloc := scanners[i].loc
				if iloc == nil {
					iloc = point.NewPoint(0, 0, 0)
				}
				fmt.Println(iloc)
				for j := i + 1; j < len(scanners); j++ {
					jloc := scanners[j].loc
					md := abs(iloc.X-jloc.X) + abs(iloc.Y-jloc.Y) + abs(iloc.Z-jloc.Z)
					if md > max {
						max = md
					}
				}
			}
			fmt.Println(max)
			return nil
		}),
	)
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func checkIntersect(scanners []*Scanner, i, j int) bool {
	scannerI := scanners[i]
	scannerJ := scanners[j]
	// N * N * B
	for _, ip := range scannerI.points {
		// N * N * B * B
		for _, jp := range scannerJ.points {
			offsetX := ip.X - jp.X
			offsetY := ip.Y - jp.Y
			offsetZ := ip.Z - jp.Z

			// N * N * B * B * 24
			for _, rotFunc := range point.RotFuncsByPoint(ip) {
				// N * N * B * B * 24 * B = N^2 * B^3 * 24
				var intersectCount int
				for _, jp2 := range scannerJ.points {
					relativeP := rotFunc(jp2.Copy().Offset(offsetX, offsetY, offsetZ))
					if scannerI.Has(relativeP) {
						intersectCount++
					}
					if intersectCount >= 12 {
						fmt.Println("INTERSECTED: ", i, j)
						scannerJ.Shift(offsetX, offsetY, offsetZ, rotFunc)
						return true
					}
				}
			}
		}
	}
	return false
}

func printScanners(scanners []*Scanner) {
	for idx, s := range scanners {
		fmt.Println(idx, "==============")
		fmt.Println(s)
	}
}

type Scanner struct {
	loc    *point.Point
	points []*point.Point
	hasMap map[int]map[int]map[int]bool
}

func (s *Scanner) Has(p *point.Point) bool {
	return s.hasMap[p.X][p.Y][p.Z]
}

func (s *Scanner) Shift(offsetX, offsetY, offsetZ int, rotFunc func(*point.Point) *point.Point) {
	newS := &Scanner{}
	for _, p := range s.points {
		newS.AppendPoint(rotFunc(p.Copy().Offset(offsetX, offsetY, offsetZ)))
	}
	s.loc = rotFunc(s.loc.Copy().Offset(offsetX, offsetY, offsetZ))
	s.points = newS.points
	s.hasMap = newS.hasMap
}

func (s *Scanner) Append(r []int) {
	s.AppendPoint(point.NewPoint(r[0], r[1], r[2]))
}

func (s *Scanner) AppendPoint(p *point.Point) {
	if s.hasMap[p.X][p.Y][p.Z] {
		return
	}
	s.points = append(s.points, point.NewPoint(p.X, p.Y, p.Z))
	if s.hasMap == nil {
		s.hasMap = map[int]map[int]map[int]bool{}
	}
	if s.hasMap[p.X] == nil {
		s.hasMap[p.X] = map[int]map[int]bool{}
	}
	if s.hasMap[p.X][p.Y] == nil {
		s.hasMap[p.X][p.Y] = map[int]bool{}
	}
	s.hasMap[p.X][p.Y][p.Z] = true
}

func (s *Scanner) Sort() {
	sort.SliceStable(s.points, func(i, j int) bool { return s.points[i].X < s.points[j].X })
}

func (s *Scanner) String() string {
	var ret []string
	for _, r := range s.points {
		ret = append(ret, r.String())
	}
	return strings.Join(ret, "\n")
}
