package main

import (
	"fmt"
	"math"
	"strings"
)

type PieRenderer struct{}

type pieSlice struct {
	origIndex int     // Original index in the input values slice
	start     float64 // Start angle in radians [0, 2π)
	end       float64 // End angle in radians [0, 2π)
	value     float64 // Raw numeric value
	percent   float64 // Percentage of total (0-100)
	symbol    rune    // ASCII/Unicode symbol used to draw this slice
}

func (p *PieRenderer) Option() Option { return PieChart }

func (p *PieRenderer) Render(g *Graph) error {
	total, err := validateAndSum(g.yVals)
	if err != nil {
		return err
	}

	slices := buildPieSlices(g.yVals, total)
	if len(slices) == 0 {
		return fmt.Errorf(pieErrAtLeastOnePositive)
	}

	radius := max(pieMinRadius, g.height)
	rows := radius*2 + 1
	cols := rows*2 + 1

	buffer := renderPieBuffer(slices, rows, cols, float64(radius))

	legend := buildLegend(g, total, slices)
	legendStart := max(0, (rows-len(legend))/2)

	fmt.Printf(pieTitleFormat, PieChart)

	for y := range rows {
		legendText := string(pieEmptyRune)
		if i := y - legendStart; i >= 0 && i < len(legend) {
			legendText = legend[i]
		}

		cleanRow := strings.TrimRight(string(buffer[y]), string(pieEmptyRune))

		fmt.Printf(pieRowFormat, cols, cleanRow, legendText)
	}

	return nil
}

func validateAndSum(values []float64) (float64, error) {
	total := 0.0
	for _, v := range values {
		if v < 0 {
			return 0, fmt.Errorf(pieErrNegativeValues)
		}
		total += v
	}

	if total == 0 {
		return 0, fmt.Errorf(pieErrTotalGreaterThanZero)
	}

	return total, nil
}

func buildPieSlices(vals []float64, total float64) []pieSlice {
	slices := make([]pieSlice, 0, len(vals))
	currentAngle := 0.0
	symbolIndex := 0

	for i, v := range vals {
		if v <= 0 {
			continue
		}

		span := (v / total) * 2.0 * math.Pi

		s := pieSlice{
			origIndex: i,
			start:     currentAngle,
			end:       currentAngle + span,
			value:     v,
			percent:   (v / total) * 100,
			symbol:    pickSymbol(symbolIndex),
		}

		slices = append(slices, s)
		currentAngle = s.end
		symbolIndex++
	}

	if len(slices) > 0 {
		slices[len(slices)-1].end = 2 * math.Pi
	}

	return slices
}

func pickSymbol(index int) rune {
	pool := []rune(pieSymbolPool)
	if index >= 0 && index < len(pool) {
		return pool[index]
	}

	return pieFallbackSymbol
}

func renderPieBuffer(
	slices []pieSlice,
	rows, cols int,
	radius float64,
) [][]rune {
	centerX := float64(cols-1) / 2.0
	centerY := float64(rows-1) / 2.0

	buffer := make([][]rune, rows)
	for y := range rows {
		buffer[y] = make([]rune, cols)
		for x := range cols {
			buffer[y][x] = pieEmptyRune
		}
	}

	for y := range rows {
		for x := range cols {
			dx := (float64(x) - centerX) / pieAspectX
			dy := float64(y) - centerY
			dist := math.Sqrt(dx*dx + dy*dy)

			if dist > radius+pieOuterPadding {
				continue
			}

			if radius-dist <= pieBorderThreshold {
				buffer[y][x] = pieBorderRune
				continue
			}

			angle := angleClockwiseFromTop(dx, dy)
			sliceIdx := findSliceForAngle(slices, angle)
			buffer[y][x] = slices[sliceIdx].symbol
		}
	}

	return buffer
}

func angleClockwiseFromTop(dx, dy float64) float64 {
	phi := math.Atan2(-dy, dx)
	angle := (math.Pi / 2.0) - phi

	if angle < 0 {
		angle += 2 * math.Pi
	}
	if angle >= 2*math.Pi {
		angle -= 2 * math.Pi
	}

	return angle
}

func findSliceForAngle(slices []pieSlice, angle float64) int {
	for i, s := range slices {
		isLast := i == len(slices)-1
		if angle >= s.start && (angle < s.end || isLast) {
			return i
		}
	}

	return len(slices) - 1
}

// From the internet: A legend on a graph is a key, typically a box located within or beside the chart,
// that identifies the meaning of different colors, symbols, patterns, or lines used to represent data sets
func buildLegend(
	g *Graph,
	total float64,
	slices []pieSlice,
) []string {
	items := []string{
		fmt.Sprintf(pieTotalLabelFormat, formatNumber(total)),
	}

	for _, s := range slices {
		items = append(items, fmt.Sprintf(
			pieLegendItemFormat,
			s.symbol,
			g.labelFor(s.origIndex),
			formatNumber(s.value),
			s.percent,
		))
	}

	width := len([]rune(pieLegendTitle))
	for _, item := range items {
		width = max(width, len([]rune(item)))
	}

	lines := []string{
		// Top border: ┌─ Title ──────┐
		buildTopBorder(width),

		// Total value line: │ Total: 123 │
		buildContentLine(items[0], width),

		// Separator: ├──────────────┤
		buildSeparator(width),
	}

	// Each slice entry: │ ▪ Label = 10 (20.0%) │
	for _, item := range items[1:] {
		lines = append(lines, buildContentLine(item, width))
	}

	// Bottom border: └──────────────┘
	lines = append(lines, buildBottomBorder(width))

	return lines
}

// buildTopBorder creates the top of the legend box.
// Format: ┌─ Title ──────┐
func buildTopBorder(contentWidth int) string {
	title := pieLegendTitlePadding + pieLegendTitle + pieLegendTitlePadding
	fillWidth := contentWidth - len([]rune(pieLegendTitle))

	return string(defaultArcDownRight) +
		string(defaultHorizontal) +
		title +
		strings.Repeat(string(defaultHorizontal), fillWidth) +
		string(defaultHorizontal) +
		string(defaultArcDownLeft)
}

// buildContentLine creates a middle row with content.
// Format: │ Content      │
func buildContentLine(text string, width int) string {
	return pieLegendRowPrefix + padRight(text, width) + pieLegendRowSuffix
}

// buildSeparator creates the horizontal divider line.
// Format: ├──────────────┤
func buildSeparator(width int) string {
	return string(defaultRightTick) +
		strings.Repeat(string(defaultHorizontal), width+2) +
		string(pieLegendLeftTick)
}

// buildBottomBorder creates the bottom of the legend box.
// Format: └──────────────┘
func buildBottomBorder(width int) string {
	return string(defaultArcUpRight) +
		strings.Repeat(string(defaultHorizontal), width+2) +
		string(defaultArcUpLeft)
}

// padRight pads a string with spaces to reach the target width (in runes).
func padRight(text string, width int) string {
	padding := width - len([]rune(text))
	if padding <= 0 {
		return text
	}

	return text + strings.Repeat(string(pieEmptyRune), padding)
}
