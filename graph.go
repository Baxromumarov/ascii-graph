package main

import (
	"fmt"
	"math"
)

type Number interface {
	~int |
		~int8 |
		~int16 |
		~int32 |
		~int64 |
		~uint |
		~uint8 |
		~uint16 |
		~uint32 |
		~uint64 |
		~uintptr |
		~float32 |
		~float64
}

type Graph struct {
	yVals   []float64
	xVals   []string
	opt     Option
	yMax    float64
	yMin    float64
	height  int
	charSet CharSet
}

type Point struct {
	X int
	Y int
}

func NewGraph[T Number](yVals []T, opt Option) *Graph {
	vals := make([]float64, len(yVals))
	for i, v := range yVals {
		vals[i] = float64(v)
	}

	out := &Graph{
		yVals:   vals,
		opt:     opt,
		height:  defaultChartHeight,
		charSet: DefaultCharSet,
	}

	out.yMax, out.yMin = out.minMax()

	return out
}

func (g *Graph) Draw() error {
	if len(g.yVals) == 0 {
		return fmt.Errorf("input yVals cannot be empty")
	}

	renderer, err := GetRenderer(g.opt)
	if err != nil {
		return err
	}

	return renderer.Render(g)
}

func (g *Graph) WithXVals(vals []string) *Graph {
	g.xVals = make([]string, len(vals))
	copy(g.xVals, vals)
	return g
}

func (g *Graph) WithHeight(height int) *Graph {
	g.height = max(0, height)

	return g
}

func (g *Graph) WithCharSet(charSet CharSet) *Graph {
	g.charSet = charSet

	return g
}

// graph helpers
func (g *Graph) points(height int) []Point {
	points := make([]Point, len(g.yVals))
	for i, v := range g.yVals {
		points[i] = Point{
			X: i*cellWidth + 1,
			Y: g.scaleValue(v, height),
		}
	}

	return points
}

func (g *Graph) scaleValue(val float64, height int) int {
	if g.yMax == g.yMin {
		return height / 2
	}

	normalized := (val - g.yMin) / (g.yMax - g.yMin)

	y := int(math.Round(normalized * float64(height)))

	return clamp(y, 0, height)
}

func (g *Graph) valueAtRow(y, height int) float64 {
	if height == 0 {
		return g.yMin
	}

	ratio := float64(y) / float64(height)

	return g.yMin + ratio*(g.yMax-g.yMin)
}

func (g *Graph) axisRow(height int) int {
	if g.yMin <= 0 && g.yMax >= 0 {
		return g.scaleValue(0, height)
	}

	if g.yMin > 0 {
		return 0
	}

	return height
}

func (g *Graph) labelFor(index int) string {
	if index >= 0 &&
		index < len(g.xVals) &&
		g.xVals[index] != "" {

		return g.xVals[index]
	}

	return fmt.Sprintf("%d", index)
}

func (g *Graph) minMax() (yMax float64, yMin float64) {
	if len(g.yVals) == 0 {
		return 0, 0
	}

	yMax, yMin = g.yVals[0], g.yVals[0]

	for _, v := range g.yVals[1:] {
		yMax = max(v, yMax)
		yMin = min(v, yMin)
	}

	return yMax, yMin
}
