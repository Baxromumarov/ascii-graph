package main

import (
	"fmt"
	"math"
	"strings"
)

func clamp(value, minValue, maxValue int) int {
	if value < minValue {
		return minValue
	}

	if value > maxValue {
		return maxValue
	}

	return value
}

func abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}

func sign(v int) int {
	switch {
	case v > 0:
		return 1
	case v < 0:
		return -1
	default:
		return 0
	}
}

func formatNumber(v float64) string {
	if math.Abs(v-math.Round(v)) < 1e-9 {
		return fmt.Sprintf("%d", int(math.Round(v)))
	}

	t := fmt.Sprintf("%.2f", v)
	t = strings.TrimRight(t, "0")
	t = strings.TrimRight(t, ".")

	return t
}

func putCentered(row []rune, center int, text string) {
	start := center - len([]rune(text))/2
	for i, ch := range text {
		pos := start + i
		if pos >= 0 && pos < len(row) {
			row[pos] = ch
		}
	}
}

func makeConnectionGrid(rows, cols int) [][]uint8 {
	m := make([][]uint8, rows)
	for i := range m {
		m[i] = make([]uint8, cols)
	}

	return m
}

func connectCells(grid [][]uint8, x1, y1, x2, y2 int) {
	if !inBounds(grid, x1, y1) || !inBounds(grid, x2, y2) {
		return
	}

	dx := x2 - x1
	dy := y2 - y1

	switch {
	case dx == 1 && dy == 0:
		grid[y1][x1] |= dirRight
		grid[y2][x2] |= dirLeft
	case dx == -1 && dy == 0:
		grid[y1][x1] |= dirLeft
		grid[y2][x2] |= dirRight
	case dx == 0 && dy == 1:
		grid[y1][x1] |= dirUp
		grid[y2][x2] |= dirDown
	case dx == 0 && dy == -1:
		grid[y1][x1] |= dirDown
		grid[y2][x2] |= dirUp
	}
}

func inBounds(grid [][]uint8, x, y int) bool {
	if y < 0 || y >= len(grid) {
		return false
	}

	if x < 0 || x >= len(grid[y]) {
		return false
	}
	return true
}
