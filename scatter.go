package main

type ScatterRenderer struct{ CartesianBase }

func (s *ScatterRenderer) Option() Option { return ScatterPlot }

func (s *ScatterRenderer) Render(g *Graph) error {
	return s.renderCartesian(g, func(canvas *Canvas, points []Point, axisY int) {
		for _, p := range points {
			canvas.Set(p.X, p.Y, pointRune)
		}
	})
}
