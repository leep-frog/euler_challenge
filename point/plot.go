package point

import (
	"fmt"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/font"
)

type Plot struct {
	P *plot.Plot
}

func CreatePlot(name string, width, height font.Length, plottables ...Plottable) error {
	p := NewPlot()
	if err := p.Add(plottables...); err != nil {
		return err
	}
	return p.Save(width, height, name)
}

func NewPlot() *Plot {
	return &Plot{plot.New()}
}

func (p *Plot) Save(width, height font.Length, name string) error {
	return p.P.Save(width, height, name)
}

func (p *Plot) Add(plottables ...Plottable) error {
	// Don't use 'range' function since we append to plottable in this loop.
	for i := 0; i < len(plottables); i++ {
		plottable := plottables[i]
		subPlottables, err := plottable.Plot(p)
		if err != nil {
			return fmt.Errorf("failed to plot plottable: %v", err)
		}
		plottables = append(plottables, subPlottables...)
	}
	return nil
}

// Plottable is the interface required to plot an object.
type Plottable interface {
	// Plot adds relevant data to the plot and returns any other objects that
	// need to be plotted.
	Plot(*Plot) ([]Plottable, error)
}
