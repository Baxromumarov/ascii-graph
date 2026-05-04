package main

import "fmt"

type Canvas struct {
	rows    int
	cols    int
	cells   [][]rune
	charSet CharRunes
}

func NewCanvas(rows, cols int, charSet CharRunes) *Canvas {
	cells := make([][]rune, rows)
	for i := range cells {
		cells[i] = make([]rune, cols)
		for j := range cells[i] {
			cells[i][j] = ' '
		}
	}

	return &Canvas{
		rows:    rows,
		cols:    cols,
		cells:   cells,
		charSet: charSet,
	}
}

func (c *Canvas) Set(x, y int, ch rune) {
	if y < 0 ||
		y >= c.rows ||
		x < 0 ||
		x >= c.cols {

		return
	}

	c.cells[y][x] = ch
}

func (c *Canvas) Get(x, y int) rune {
	if y < 0 ||
		y >= c.rows ||
		x < 0 ||
		x >= c.cols {

		return ' '
	}

	return c.cells[y][x]
}

func (c *Canvas) Row(y int) []rune {
	if y < 0 || y >= c.rows {
		return nil
	}

	return c.cells[y]
}

func (c *Canvas) Dimensions() (rows, cols int) {
	return c.rows, c.cols
}

func (c *Canvas) Print(labelFunc func(y int) string) {
	for y := c.rows - 1; y >= 0; y-- {
		label := labelFunc(y)
		fmt.Printf("%8s %s\n", label, string(c.cells[y]))
	}
}

// LabelRow prints a centered label row below the canvas.
func (c *Canvas) LabelRow(labelFunc func(i int) string, centers []int) {
	labelRow := make([]rune, c.cols)
	for i := range labelRow {
		labelRow[i] = ' '
	}

	for i, center := range centers {
		putCentered(labelRow, center, labelFunc(i))
	}

	fmt.Printf("%8s %s\n", "", string(labelRow))
}
