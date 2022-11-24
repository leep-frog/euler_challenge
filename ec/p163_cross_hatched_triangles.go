package eulerchallenge

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func search163(v *vertex163, vertices, path []*vertex163) int {
	if len(path) == 3 {
		// Make sure they are not collinear points (aka on the same type of line)
		if _, ok := path[0].m[v.id]; !ok {
			return 0
		}

		if path[0].m[path[1].id] == path[0].m[path[2].id] {
			return 0
		}
		return 1
	}

	cnt := 0
	for vid := range v.m {
		if vid < v.id {
			continue
		}
		path = append(path, vertices[vid])
		cnt += search163(vertices[vid], vertices, path)
		path = path[:len(path)-1]
	}
	return cnt
}

type vertex163 struct {
	id int
	m  map[int]string
}

func newVertex163(id int) *vertex163 {
	return &vertex163{id, map[int]string{}}
}

func connectVertices(vs []*vertex163, line string) {
	for i, v := range vs {
		for j := i + 1; j < len(vs); j++ {
			v.connect(vs[j], line)
		}
	}
}

func (v *vertex163) String() string {
	return fmt.Sprintf("V%d", v.id)
}

func (v *vertex163) connect(that *vertex163, line string) {
	v.m[that.id] = line
	that.m[v.id] = line
}

func P163() *problem {
	return intInputNode(163, func(o command.Output, n int) {

		var vid int
		var vertices []*vertex163
		var horzRows, zigRows [][]*vertex163
		for i := 0; i <= n; i++ {
			var hRow, zRow []*vertex163
			for j := 0; j < 2*i+1; j++ {
				hRow = append(hRow, newVertex163(vid))
				vid++
			}
			horzRows = append(horzRows, hRow)
			vertices = append(vertices, hRow...)

			if i == n {
				continue
			}
			for j := 0; j < 4*i+3; j++ {
				zRow = append(zRow, newVertex163(vid))
				vid++
			}
			zigRows = append(zigRows, zRow)
			vertices = append(vertices, zRow...)
		}

		// Now go through types of lines

		// horizontal lines
		for _, hRow := range horzRows[1:] {
			connectVertices(hRow, "horizontal")
		}

		// (0, 0)-->(1, 2) line
		for i := 0; i < n; i++ {
			var downLeftRow, downRightRow []*vertex163
			for j := i; j <= n; j++ {
				row := horzRows[j]
				downLeftRow = append(downLeftRow, row[i*2])
				downRightRow = append(downRightRow, row[len(row)-1-i*2])
			}
			for j := i; j < n; j++ {
				row := zigRows[j]
				downLeftRow = append(downLeftRow, row[i*4])
				downRightRow = append(downRightRow, row[len(row)-1-i*4])
			}
			connectVertices(downLeftRow, "downLeft")
			connectVertices(downRightRow, "downRight")
		}

		// vertical lines
		for i := 0; i < n; i++ {
			var vertRowLeft, vertRowRight []*vertex163
			for j := i; j <= n; j++ {
				row := horzRows[j]
				vertRowLeft = append(vertRowLeft, row[j-i])
				vertRowRight = append(vertRowRight, row[len(row)-1-(j-i)])
			}
			for j := i; j < n; j++ {
				row := zigRows[j]
				vertRowLeft = append(vertRowLeft, row[2*(j-i)+1])
				vertRowRight = append(vertRowRight, row[len(row)-1-(2*(j-i)+1)])
			}
			connectVertices(vertRowLeft, "vertical")
			if i != 0 {
				connectVertices(vertRowRight, "vertical")
			}
		}

		// (0, 0)-->(2, 1) line starting at sides of upright triangles
		for i := 0; i < n; i++ {
			toRight := []*vertex163{
				zigRows[i][0],
				zigRows[i][1],
				horzRows[i+1][2],
			}
			toLeft := []*vertex163{
				zigRows[i][len(zigRows[i])-1],
				zigRows[i][len(zigRows[i])-2],
				horzRows[i+1][len(horzRows[i+1])-3],
			}
			for j := 0; j < maths.Min(i, n-i-1); j++ {
				zRow := zigRows[i+j+1]
				hRow := horzRows[i+j+2]
				toRight = append(toRight,
					zRow[8*(j+1)-1],
					zRow[8*(j+1)],
					zRow[8*(j+1)+1],
					hRow[4*(j+1)+2],
				)
				toLeft = append(toLeft,
					zRow[len(zRow)-1-(8*(j+1)-1)],
					zRow[len(zRow)-1-(8*(j+1))],
					zRow[len(zRow)-1-(8*(j+1)+1)],
					hRow[len(hRow)-1-(4*(j+1)+2)],
				)
			}
			connectVertices(toRight, "toRight")
			connectVertices(toLeft, "toLeft")
		}

		// (0, 0)-->(2, 1) line starting at sides of upright triangles
		for i := 1; i < n; i++ {
			toRight := []*vertex163{
				horzRows[i][0],
			}
			toLeft := []*vertex163{
				horzRows[i][len(horzRows[i])-1],
			}

			for j := 0; j < maths.Min(i, n-i); j++ {
				zRow := zigRows[i+j]
				hRow := horzRows[i+j+1]
				toRight = append(toRight,
					zRow[8*j+3],
					zRow[8*j+4],
					zRow[8*j+5],
					hRow[4*(j+1)],
				)
				toLeft = append(toLeft,
					zRow[len(zRow)-1-(8*j+3)],
					zRow[len(zRow)-1-(8*j+4)],
					zRow[len(zRow)-1-(8*j+5)],
					hRow[len(hRow)-1-(4*(j+1))],
				)
			}
			connectVertices(toRight, "toRight")
			connectVertices(toLeft, "toLeft")
		}

		var sum int
		for _, v := range vertices {
			sum += search163(v, vertices, []*vertex163{v})
		}
		o.Stdoutln(sum)
	}, []*execution{
		{
			args: []string{"1"},
			want: "16",
		},
		{
			args: []string{"2"},
			want: "104",
		},
		{
			args:     []string{"36"},
			want:     "343047",
			estimate: 1.5,
		},
	})
}
