package main

import "fmt"

type ChartRenderer interface {
	Render(g *Graph) error
	Option() Option
}

var Registry = map[Option]ChartRenderer{
	BarChart:    &BarRenderer{},
	LineGraph:   &LineRenderer{},
	ScatterPlot: &ScatterRenderer{},
	PieChart:    &PieRenderer{},
}

func GetRenderer(opt Option) (ChartRenderer, error) {
	r, ok := Registry[opt]
	if !ok {
		return nil, fmt.Errorf("unsupported chart option: %q", opt)
	}

	return r, nil
}
