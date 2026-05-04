package main

import "fmt"

// CartesianBase provides common X/Y axis drawing. Concrete renderers
// embed this and implement renderData().
type CartesianBase struct{}

func (c *CartesianBase) renderCartesian(
	g *Graph,
	renderData func(
		canvas *Canvas,
		points []Point,
		axisY int,
	),
) error {

	height := g.height
	width := max(len(g.yVals)*cellWidth+1, cellWidth+1)

	canvas := NewCanvas(height+1, width, g.charSet.Runes())
	points := g.points(height)
	axisY := g.axisRow(height)

	c.drawAxes(canvas, axisY, points, height)
	renderData(canvas, points, axisY)

	fmt.Printf("%s\n\n", g.opt)
	canvas.Print(func(y int) string {
		return formatNumber(g.valueAtRow(y, height))
	})

	centers := make([]int, len(points))
	for i, p := range points {
		centers[i] = p.X
	}

	canvas.LabelRow(g.labelFor, centers)

	return nil
}

func (c *CartesianBase) drawAxes(
	canvas *Canvas,
	axisY int,
	points []Point,
	height int,
) {
	rows, cols := canvas.Dimensions()
	ch := canvas.charSet

	for y := range rows {
		canvas.Set(0, y, ch.Vertical)
	}

	for x := range cols {
		canvas.Set(x, axisY, ch.Horizontal)
	}

	switch {
	case axisY == 0:
		canvas.Set(0, axisY, ch.UpRight)
	case axisY == height:
		canvas.Set(0, axisY, ch.ArcDownRight)
	default:
		canvas.Set(0, axisY, ch.RightTick)
	}

	for _, p := range points {
		canvas.Set(p.X, axisY, ch.DownTick)
	}
}
