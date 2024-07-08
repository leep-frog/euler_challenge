package twentyone

import (
	"fmt"
	"log"
	"path/filepath"
	"regexp"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/parse"
)

func D22() *problem {
	return command.SerialNodes(
		command.Description(""),
		command.ExecutorNode(func(o command.Output, d *command.Data) error {
			// on x=-20..26,y=-36..17,z=-47..7

			/*oc := &OnCubes{}
			r := regexp.MustCompile("^([a-z]+) x=(-?[0-9]+)..(-?[0-9]+),y=(-?[0-9]+)..(-?[0-9]+),z=(-?[0-9]+)..(-?[0-9]+)")
			lines := eulerchallenge.ReadFileLines(filepath.ToSlash(filepath.Join("..", "aoc", "2021", "input", "d22.txt")))
			for idx, line := range lines {
				fmt.Println(idx, len(lines))
				m := r.FindStringSubmatch(line)
				if m == nil {
					log.Fatalf("bad news bears: %q", line)
				}
				on := m[1] == "on"
				x1, x2, y1, y2, z1, z2 := eulerchallenge.Atoi(m[2]), eulerchallenge.Atoi(m[3]), eulerchallenge.Atoi(m[4]), eulerchallenge.Atoi(m[5]), eulerchallenge.Atoi(m[6]), eulerchallenge.Atoi(m[7])
				for i := x1; i <= x2; i++ {
					for j := y1; j <= y2; j++ {
						for k := z1; k <= z2; k++ {
							if on {
								oc.On(i, j, k)
							} else {
								oc.Off(i, j, k)
							}
						}
					}
				}
			}*/
			region := &Cube{
				x1: -50,
				x2: 50,
				y1: -50,
				y2: 50,
				z1: -50,
				z2: 50,
			}
			r := regexp.MustCompile("^([a-z]+) x=(-?[0-9]+)..(-?[0-9]+),y=(-?[0-9]+)..(-?[0-9]+),z=(-?[0-9]+)..(-?[0-9]+)")
			//lines := eulerchallenge.ReadFileLines(filepath.ToSlash(filepath.Join("..", "aoc", "2021", "input", "d22.txt")))
			//lines := eulerchallenge.ReadFileLines(filepath.ToSlash(filepath.Join("..", "aoc", "2021", "input", "d22_small.txt")))
			//lines := eulerchallenge.ReadFileLines(filepath.ToSlash(filepath.Join("..", "aoc", "2021", "input", "d22_big.txt")))
			//lines := eulerchallenge.ReadFileLines(filepath.ToSlash(filepath.Join("..", "aoc", "2021", "input", "d22_real.txt")))
			lines := parse.ReadFileLines(filepath.ToSlash(filepath.Join("..", "aoc", "2021", "input", "d22_google.txt")))
			cubes := []*Cube{}
			for idx, line := range lines {
				if idx == len(lines)-1 && line == "" {
					break
				}
				m := r.FindStringSubmatch(line)
				if m == nil {
					log.Fatalf("bad news bears: %q", line)
				}
				x1, x2, y1, y2, z1, z2 := parse.Atoi(m[2]), parse.Atoi(m[3]), parse.Atoi(m[4]), parse.Atoi(m[5]), parse.Atoi(m[6]), parse.Atoi(m[7])
				c := &Cube{
					x1: x1,
					x2: x2,
					y1: y1,
					y2: y2,
					z1: z1,
					z2: z2,
					on: m[1] == "on",
				}
				if c.Intersect(region) != nil || true {
					cubes = append(cubes, c)
				}
			}

			initLen := len(cubes)
			var relevantCubes []*Cube
			for i := 0; i < initLen; i++ {
				iCube := cubes[i]
				//fmt.Printf("STEP %d =========== %v\n", i, iCube)
				initJen := len(relevantCubes)
				for j := 0; j < initJen; j++ {
					jCube := relevantCubes[j]
					if newC := iCube.Intersect(jCube); newC != nil {
						relevantCubes = append(relevantCubes, newC)
					}
				}
				if iCube.on {
					relevantCubes = append(relevantCubes, iCube)
				}
				/*for _, c := range relevantCubes {
					//fmt.Println(c)
				}*/
				//fmt.Println(cubeSize(relevantCubes))
			}

			fmt.Println(cubeSize(relevantCubes))
			return nil
		}),
	)
}

func cubeSize(cubes []*Cube) int {
	var size int
	for _, c := range cubes {
		size += c.Size()
	}
	return size
}

type Cube struct {
	x1, x2, y1, y2, z1, z2 int
	on                     bool
}

func (c *Cube) String() string {
	return fmt.Sprintf("[%d %v] %d,%d  %d,%d  %d,%d", c.Size(), c.on, c.x1, c.x2, c.y1, c.y2, c.z1, c.z2)
}

func (c *Cube) Size() int {
	r := (c.x2 - c.x1 + 1) * (c.y2 - c.y1 + 1) * (c.z2 - c.z1 + 1)
	if c.on {
		return r
	}
	return -r
}

/*func (c *Cube) Intersect(that *Cube) *Cube {
	x1, x2 := max(c.x1, that.x1), min(c.x2, that.x2)
	y1, y2 := max(c.y1, that.y1), min(c.y2, that.y2)
	z1, z2 := max(c.z1, that.z1), min(c.z2, that.z2)

	if x1 > x2 || y1 > y2 || z1 > z2 {
		return nil
	}

	newDepth := -(that.depth + 1)
	if that.depth < 0 {
		newDepth = (-that.depth) + 1
	}

	return &Cube{
		x1:    x1,
		x2:    x2,
		y1:    y1,
		y2:    y2,
		z1:    z1,
		z2:    z2,
		depth: newDepth,
	}
}*/

func (c *Cube) Intersect(that *Cube) *Cube {
	x1, x2 := max(c.x1, that.x1), min(c.x2, that.x2)
	y1, y2 := max(c.y1, that.y1), min(c.y2, that.y2)
	z1, z2 := max(c.z1, that.z1), min(c.z2, that.z2)

	if x1 > x2 || y1 > y2 || z1 > z2 {
		return nil
	}

	on := c.on == that.on
	if c.on {
		on = !on
	}
	return &Cube{
		x1: x1,
		x2: x2,
		y1: y1,
		y2: y2,
		z1: z1,
		z2: z2,
		on: on,
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type OnCubes struct {
	m     map[int]map[int]map[int]bool
	count int
}

func (oc *OnCubes) Off(x, y, z int) {
	if oc.m[x] == nil || oc.m[x][y] == nil || !oc.m[x][y][z] {
		return
	}
	oc.count--
	delete(oc.m[x][y], z)
}

func (oc *OnCubes) On(x, y, z int) {
	if oc.m == nil {
		oc.m = map[int]map[int]map[int]bool{}
	}
	if oc.m[x] == nil {
		oc.m[x] = map[int]map[int]bool{}
	}
	if oc.m[x][y] == nil {
		oc.m[x][y] = map[int]bool{}
	}
	if oc.m[x][y][z] {
		return
	}
	oc.count++
	oc.m[x][y][z] = true
}
