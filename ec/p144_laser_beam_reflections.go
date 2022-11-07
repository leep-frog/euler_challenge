package eulerchallenge

import (
	"fmt"
	"log"
	"math"
	"sort"

	"github.com/leep-frog/command"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

type customPlot struct {
	minX, maxX, minY, maxY float64

	plotters []plot.Plotter
}

func (cp *customPlot) addEllipse(xCoeff, b float64) {
	xBound := math.Sqrt(b / xCoeff)
	var ellipseTop, ellipseBottom plotter.XYs
	for x := -xBound; x <= xBound; x += 0.25 {
		y := math.Sqrt(b - xCoeff*x*x)
		ellipseTop = append(ellipseTop, plotter.XY{x, y})
		ellipseBottom = append(ellipseBottom, plotter.XY{x, -y})
	}
	lineTop, err := plotter.NewLine(ellipseTop)
	if err != nil {
		log.Fatalf("nope top: %v", err)
	}
	lineBottom, err := plotter.NewLine(ellipseBottom)
	if err != nil {
		log.Fatalf("nope bottom: %v", err)
	}
	cp.plotters = append(cp.plotters, lineTop, lineBottom)
}

func (cp *customPlot) addPoint(x, y float64) {
	sc, err := plotter.NewScatter(plotter.XYs{
		plotter.XY{x, y},
	})
	if err != nil {
		log.Fatalf("nope sc: %v", err)
	}
	cp.plotters = append(cp.plotters, sc)
}

func (cp *customPlot) addLineSegment(fromX, fromY, toX, toY float64) {
	ls, err := plotter.NewLine(plotter.XYs{
		plotter.XY{fromX, fromY},
		plotter.XY{toX, toY},
	})
	if err != nil {
		log.Fatalf("nope laser1: %v", err)
	}
	cp.plotters = append(cp.plotters, ls)
}

func (cp *customPlot) addBorder() {
	ls, err := plotter.NewLine(plotter.XYs{
		plotter.XY{cp.minX, cp.minY},
		plotter.XY{cp.minX, cp.maxY},
		plotter.XY{cp.maxX, cp.maxY},
		plotter.XY{cp.maxX, cp.minY},
		plotter.XY{cp.minX, cp.minY},
	})
	if err != nil {
		log.Fatalf("nope laser1: %v", err)
	}
	cp.plotters = append(cp.plotters, ls)
}

func (cp *customPlot) addLineThrough(x, y, m float64) {
	b := y - m*x
	leftX := cp.minX
	leftY := m*leftX + b
	if leftY < cp.minY {
		leftY = cp.minY
		leftX = (leftY - b) / m
	}
	if leftY > cp.maxY {
		leftY = cp.maxY
		leftX = (leftY - b) / m
	}

	rightX := cp.maxX
	rightY := m*rightX + b
	if rightY < cp.minY {
		rightY = cp.minY
		rightX = (rightY - b) / m
	}
	if rightY > cp.maxY {
		rightY = cp.maxY
		rightX = (rightY - b) / m
	}
	ls, err := plotter.NewLine(plotter.XYs{
		plotter.XY{leftX, leftY},
		plotter.XY{rightX, rightY},
	})
	if err != nil {
		log.Fatalf("nope laser1: %v", err)
	}
	cp.plotters = append(cp.plotters, ls)
}

func plot144(old, new []float64, mr, mn, mt float64, count int) {
	p := plot.New()
	cp := &customPlot{-7, 7, -12, 12, nil}
	cp.addEllipse(4, 100)
	cp.addPoint(old[0], old[1])
	cp.addPoint(new[0], new[1])
	cp.addLineSegment(old[0], old[1], new[0], new[1])
	cp.addLineThrough(new[0], new[1], mt)
	cp.addLineThrough(new[0], new[1], mn)
	cp.addLineThrough(new[0], new[1], mr)
	cp.addBorder()
	p.Add(cp.plotters...)
	p.Save(100, 200, fmt.Sprintf("plot-%d.png", count))
}

func P144() *problem {
	return noInputNode(144, func(o command.Output) {

		var points plotter.XYs

		first := []float64{0, 10.1}
		cur := []float64{1.4, -9.6}
		m := (cur[1] - first[1]) / (cur[0] - first[0])
		b := first[1] - m*first[0]
		cur = first
		points = append(points, plotter.XY{cur[0], cur[1]})
		mbs := [][]float64{{m, b}}
		for count := 0; count < 10000000; count++ {
			// Get next point
			xs := solveQuadratic(m*m+4, 2*m*b, b*b-100)
			sort.SliceStable(xs, func(i, j int) bool {
				return math.Abs(xs[i]-cur[0]) > math.Abs(xs[j]-cur[0])
			})
			cur = []float64{
				xs[0],
				m*xs[0] + b,
			}
			points = append(points, plotter.XY{cur[0], cur[1]})

			// Tangent slope = -4x/y
			mt := (-4 * cur[0] / cur[1])
			// Normal slope: // mn := -1 / mt
			// mt/1 = rise/run = y/x -> (x, y) = (1, mt)
			// n =
			// d - 2(d * n)n
			// d = (x=1, y=m); n = (x=-mt, y=1)

			d := []float64{-m, 1}
			n := []float64{-mt, 1}
			nMag := math.Sqrt(mt*mt + 1)
			n[0], n[1] = n[0]/nMag, n[1]/nMag

			dot := d[0]*n[0] + d[1]*n[1]

			r := []float64{
				d[0] - 2*dot*n[0],
				d[1] - 2*dot*n[1],
			}
			m = -r[0] / r[1]
			b = cur[1] - m*cur[0]
			mbs = append(mbs, []float64{m, b})
			//plot144(old, cur, m, mn, mt, count)
			if -0.01 < cur[0] && cur[0] < 0.01 && cur[1] > 0 {
				o.Stdoutln(count)
				return
			}
		}
	}, &execution{
		want: "354",
	})
}

func solveQuadratic(a, b, c float64) []float64 {
	det := b*b - 4*a*c
	if det < 0 {
		return nil
	}
	return []float64{
		(-b - math.Sqrt(det)) / (2 * a),
		(-b + math.Sqrt(det)) / (2 * a),
	}
}
