package twentyone

import (
	"fmt"
	"log"
	"path/filepath"
	"regexp"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/parse"
)

func D22_JIC() *problem {
	return command.SerialNodes(
		command.Description(""),
		command.ExecutorNode(func(o command.Output, d *command.Data) error {
			// on x=-20..26,y=-36..17,z=-47..7

			/*oc := &OnCubesJIC{}
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
			region := &CubeJIC{
				x1: -50,
				x2: 50,
				y1: -50,
				y2: 50,
				z1: -50,
				z2: 50,
			}
			r := regexp.MustCompile("^([a-z]+) x=(-?[0-9]+)..(-?[0-9]+),y=(-?[0-9]+)..(-?[0-9]+),z=(-?[0-9]+)..(-?[0-9]+)")
			lines := parse.ReadFileLines(filepath.ToSlash(filepath.Join("..", "aoc", "2021", "input", "d22.txt")))
			//lines := eulerchallenge.ReadFileLines(filepath.ToSlash(filepath.Join("..", "aoc", "2021", "input", "d22_real.txt")))
			cubes := []*CubeJIC{}
			for idx, line := range lines {
				if idx == len(lines)-1 && line == "" {
					break
				}
				m := r.FindStringSubmatch(line)
				if m == nil {
					log.Fatalf("bad news bears: %q", line)
				}
				x1, x2, y1, y2, z1, z2 := parse.Atoi(m[2]), parse.Atoi(m[3]), parse.Atoi(m[4]), parse.Atoi(m[5]), parse.Atoi(m[6]), parse.Atoi(m[7])
				depth := 1
				if m[1] == "off" {
					depth = 0
				}
				c := &CubeJIC{
					x1:    x1,
					x2:    x2,
					y1:    y1,
					y2:    y2,
					z1:    z1,
					z2:    z2,
					depth: depth,
				}
				if c.Intersect(region) != nil {
					cubes = append(cubes, c)
				}
			}

			oc := &OnCubesJIC{}
			for _, c := range cubes {
				for i := c.x1; i <= c.x2; i++ {
					for j := c.y1; j <= c.y2; j++ {
						for k := c.z1; k <= c.z2; k++ {
							if c.depth == 1 {
								oc.On(i, j, k)
							} else {
								oc.Off(i, j, k)
							}
						}
					}
				}
			}

			fmt.Println(oc.count)

			initLen := len(cubes)
			for i := 0; i < initLen; i++ {

				iCube := cubes[i]
				fmt.Println(iCube)
				newCs := []*CubeJIC{}
				for j := i + 1; j < len(cubes); j++ {

					jCube := cubes[j]
					if newC := iCube.Intersect(jCube); newC != nil {
						newCs = append(newCs, newC)
					}
				}
				cubes = append(cubes, newCs...)
			}

			var size int
			for _, c := range cubes {
				size += c.Size()
			}
			fmt.Println(size)
			return nil
		}),
	)
}

type CubeJIC struct {
	x1, x2, y1, y2, z1, z2 int
	depth                  int
}

func (c *CubeJIC) String() string {
	return fmt.Sprintf("[%d] %d,%d  %d,%d  %d,%d", c.depth, c.x1, c.x2, c.y1, c.y2, c.z1, c.z2)
}

func (c *CubeJIC) Size() int {
	return c.depth * ((c.x2 - c.x1) * (c.y2 - c.y1) * (c.z2 - c.z1))
}

func (c *CubeJIC) Intersect(that *CubeJIC) *CubeJIC {
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

	return &CubeJIC{
		x1:    x1,
		x2:    x2,
		y1:    y1,
		y2:    y2,
		z1:    z1,
		z2:    z2,
		depth: newDepth,
	}
}

type OnCubesJIC struct {
	m     map[int]map[int]map[int]bool
	count int
}

func (oc *OnCubesJIC) Off(x, y, z int) {
	if oc.m[x] == nil || oc.m[x][y] == nil || !oc.m[x][y][z] {
		return
	}
	oc.count--
	delete(oc.m[x][y], z)
}

func (oc *OnCubesJIC) On(x, y, z int) {
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
