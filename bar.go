package main

type BarRenderer struct{ CartesianBase }

func (b *BarRenderer) Option() Option { return BarChart }

func (b *BarRenderer) Render(g *Graph) error {
	return b.renderCartesian(g, func(
		canvas *Canvas,
		points []Point,
		axisY int,
	) {
		for _, p := range points {
			b.drawBar(canvas, axisY, p)
		}
	},
	)
}

func (b *BarRenderer) drawBar(canvas *Canvas, axisY int, p Point) {
	start := min(axisY, p.Y)
	end := max(axisY, p.Y)

	for y := start; y <= end; y++ {
		canvas.Set(p.X, y, barRune)
	}
}
