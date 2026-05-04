package main

import (
	"fmt"
	"strings"
)

const cellWidth = 4

var (
	Star           = "*"
	Plot           = "┤"
	Horizontal     = "─"
	VerticalLine   = "│"
	ArcDownRight   = "╭"
	ArcDownLeft    = "╮"
	ArcUpRight     = "╰"
	ArcUpLeft      = "╯"
	EndCap         = "╴"
	StartCap       = "╶"
	UpRight        = "└"
	DownHorizontal = "┬"
	LeftArrow      = ">"
	UpArrow        = "^"
)

type Graph struct {
	vec        []int
	opt        Option
	yMax, yMin int
}

type Point struct {
	X, Y int
}

type Option string

var (
	BarChart    Option = "BarChart"
	LineGraph   Option = "LineGraph"
	PieChart    Option = "PieChart"
	ScatterPlot Option = "ScatterPlot"
)

func NewGraph(input []int, opt Option) *Graph {
	v := make([]int, len(input))
	copy(v, input)

	out := &Graph{
		vec: v,
		opt: opt,
	}

	yMax, yMin := out.minMax()
	out.yMax = yMax
	out.yMin = yMin

	return out
}

func (g *Graph) Draw() error {
	if len(g.vec) == 0 {
		return fmt.Errorf("vec cannot be empty")
	}

	return g.draw()

}

func (g *Graph) draw() error {
	if len(g.vec) == 0 {
		return fmt.Errorf("vec cannot be empty")
	}

	fmt.Println("Input:", g.vec)
	fmt.Println("yMax:", g.yMax, "yMin:", g.yMin)
	fmt.Println()

	points := g.points()

	fmt.Print("Points: ")
	for i, p := range points {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("(%d,%d)", p.X, p.Y)
	}
	fmt.Println()
	fmt.Println()

	return g.matrix(points)
}

// matrix prints stars while preserving the original axis style.
func (g *Graph) matrix(points []Point) error {
	if len(g.vec) == 0 {
		return fmt.Errorf("vec cannot be empty")
	}

	width := len(g.vec) * cellWidth
	bottom := min(1, g.yMin)

	for y := g.yMax; y >= bottom; y-- {
		row := make([]rune, width+1)
		for i := range row {
			row[i] = ' '
		}

		for _, p := range points {
			if p.Y != y {
				continue
			}

			col := p.X*cellWidth + 1
			if col >= 0 && col < len(row) {
				row[col] = []rune(Star)[0]
			}
		}

		fmt.Printf("%v %s%s\n  %s\n", y, Plot, string(row), VerticalLine)
	}

	if err := drawXAxis(len(g.vec)); err != nil {
		return err
	}

	return nil
}

// [Deprecated]
func (g *Graph) drawYAxis() error {
	fmt.Println("yMax:", g.yMax, "yMin:", g.yMin)
	fmt.Println()

	// draw the y-axis
	n := int(g.yMax)
	for i := n; i > 0; i-- {
		fmt.Printf("%v %s\n  %s\n", i, Plot, VerticalLine)
	}

	return nil
}

func drawXAxis(max int) error {
	if max < 0 {
		return fmt.Errorf("max must be non-negative")
	}

	width := max * cellWidth

	axis := make([]string, width+1)
	for i := range axis {
		axis[i] = Horizontal
	}

	for n := 0; n <= max; n++ {
		pos := n * cellWidth
		axis[pos] = DownHorizontal
	}

	labelRow := make([]rune, width+cellWidth+1)
	for i := range labelRow {
		labelRow[i] = ' '
	}

	for n := 0; n <= max; n++ {
		pos := n*cellWidth + 1
		putText(labelRow, pos, fmt.Sprintf("%d", n))
	}

	var b strings.Builder
	b.WriteString("  " + UpRight)
	b.WriteString(strings.Join(axis, ""))
	b.WriteString("\n")
	b.WriteString("  ")

	b.WriteString(string(labelRow))

	fmt.Println(b.String())

	return nil
}

func (g *Graph) points() []Point {
	points := make([]Point, len(g.vec))

	for x, y := range g.vec {
		points[x] = Point{
			X: x,
			Y: y,
		}
	}

	return points
}

func putText(row []rune, pos int, text string) {
	for i, ch := range text {
		if pos+i < len(row) {
			row[pos+i] = ch
		}
	}
}

func (g *Graph) minMax() (yMax int, yMin int) {
	yMax = g.vec[0]
	yMin = g.vec[0]
	n := len(g.vec)

	for i := 0; i < n-1; i++ {
		yMax = max(yMax, g.vec[i+1])
		yMin = min(yMin, g.vec[i+1])
	}

	return yMax, yMin
}

func main() {
	arr := []int{4, 5, 2, 3}
	g := NewGraph(arr, ScatterPlot)

	err := g.Draw()
	if err != nil {
		fmt.Println(err)
	}
}
