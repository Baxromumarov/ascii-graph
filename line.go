package main

type LineRenderer struct{ CartesianBase }

func (l *LineRenderer) Option() Option { return LineGraph }

func (l *LineRenderer) Render(g *Graph) error {
	return l.renderCartesian(
		g,
		func(
			canvas *Canvas,
			points []Point,
			axisY int,
		) {
			l.drawLineGraph(canvas, points)

			for _, p := range points {
				canvas.Set(p.X, p.Y, pointRune)
			}
		})
}

func (l *LineRenderer) drawLineGraph(canvas *Canvas, points []Point) {
	rows, cols := canvas.Dimensions()
	connections := makeConnectionGrid(rows, cols)

	for i := 0; i < len(points)-1; i++ {
		l.drawLineSegment(connections, points[i], points[i+1])
	}

	for y := range rows {
		for x := range cols {
			if connections[y][x] == 0 {
				continue
			}
			canvas.Set(x, y, l.runeForMask(connections[y][x], canvas.charSet))
		}
	}
}

func (l *LineRenderer) drawLineSegment(grid [][]uint8, a, b Point) {
	dx := b.X - a.X
	dy := b.Y - a.Y
	if dx == 0 && dy == 0 {
		return
	}

	sx := sign(dx)
	sy := sign(dy)
	dxAbs := abs(dx)
	dyAbs := abs(dy)

	if dxAbs == 0 {
		x, y := a.X, a.Y
		for y != b.Y {
			nextY := y + sy
			connectCells(grid, x, y, x, nextY)
			y = nextY
		}
		return
	}

	// Bresenham's algorithm [https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm]
	x, y := a.X, a.Y
	err := 0
	for range dxAbs {
		prevX, prevY := x, y
		x += sx
		connectCells(grid, prevX, prevY, x, y)

		err += dyAbs
		for err >= dxAbs && y != b.Y {
			nextY := y + sy
			connectCells(grid, x, y, x, nextY)
			y = nextY
			err -= dxAbs
		}
	}
}

func (l *LineRenderer) runeForMask(mask uint8, ch CharRunes) rune {
	mask &= dirUp | dirRight | dirDown | dirLeft

	switch mask {
	case dirLeft | dirRight:
		return ch.Horizontal
	case dirUp | dirDown:
		return ch.Vertical
	case dirDown | dirRight:
		return ch.ArcDownRight
	case dirDown | dirLeft:
		return ch.ArcDownLeft
	case dirUp | dirRight:
		return ch.ArcUpRight
	case dirUp | dirLeft:
		return ch.ArcUpLeft
	case dirRight:
		return ch.StartCap
	case dirLeft:
		return ch.EndCap
	case dirUp, dirDown:
		return ch.Vertical
	default:
		return ch.RightTick
	}
}
