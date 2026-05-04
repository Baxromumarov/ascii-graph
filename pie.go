package main

import (
	"fmt"
	"math"
	"strings"
)

type PieRenderer struct{}

type pieSlice struct {
	origIndex int
	start     float64
	end       float64
	value     float64
	percent   float64
	symbol    rune
}

func (p *PieRenderer) Option() Option { return PieChart }

func (p *PieRenderer) Render(g *Graph) error {
	total := 0.0
	for _, v := range g.yVals {
		if v < 0 {
			return fmt.Errorf(pieErrNegativeValues)
		}

		total += v
	}

	if total == 0 {
		return fmt.Errorf(pieErrTotalGreaterThanZero)
	}

	slices := buildPieSlices(g.yVals, total)
	if len(slices) == 0 {
		return fmt.Errorf(pieErrAtLeastOnePositive)
	}

	radius := max(pieMinRadius, g.height)
	rows := radius*2 + 1
	cols := rows*2 + 1
	buffer := renderPieBuffer(slices, rows, cols, float64(radius))
	legend := pieLegendLines(g, total, slices)
	legendStart := max(0, (rows-len(legend))/2)

	fmt.Printf(pieTitleFormat, PieChart)
	for y := range rows {
		legendText := string(pieEmptyRune)
		if i := y - legendStart; i >= 0 && i < len(legend) {
			legendText = legend[i]
		}

		fmt.Printf(pieRowFormat, cols, strings.TrimRight(string(buffer[y]), string(pieEmptyRune)), legendText)
	}

	return nil
}

func buildPieSlices(vals []float64, total float64) []pieSlice {
	slices := make([]pieSlice, 0, len(vals))
	start := 0.0
	symbolIndex := 0

	for i, v := range vals {
		if v <= 0 {
			continue
		}

		span := (v / total) * 2.0 * math.Pi
		s := pieSlice{
			origIndex: i,
			start:     start,
			end:       start + span,
			value:     v,
			percent:   (v / total) * 100,
			symbol:    pieSymbol(symbolIndex),
		}

		slices = append(slices, s)
		start = s.end
		symbolIndex++
	}

	if len(slices) > 0 {
		slices[len(slices)-1].end = 2 * math.Pi
	}

	return slices
}

func pieSymbol(i int) rune {
	pool := []rune(pieSymbolPool)
	if i >= 0 && i < len(pool) {
		return pool[i]
	}

	return pieFallbackSymbol
}

func renderPieBuffer(
	slices []pieSlice,
	rows,
	cols int,
	radius float64,
) [][]rune {
	aspectX := pieAspectX
	cx := float64(cols-1) / 2.0
	cy := float64(rows-1) / 2.0
	buffer := make([][]rune, rows)

	for y := range rows {
		buffer[y] = make([]rune, cols)
		for x := range cols {
			buffer[y][x] = pieEmptyRune
		}
	}

	for y := range rows {
		for x := range cols {
			dx := (float64(x) - cx) / aspectX
			dy := float64(y) - cy
			dist := math.Sqrt(dx*dx + dy*dy)

			if dist > radius+pieOuterPadding {
				continue
			}

			if radius-dist <= pieBorderThreshold {
				buffer[y][x] = pieBorderRune
				continue
			}

			angle := pieAngleClockwiseFromTop(dx, dy)
			sliceIdx := sliceIndexForAngle(slices, angle)
			buffer[y][x] = slices[sliceIdx].symbol
		}
	}

	return buffer
}

func pieLegendLines(
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
		string(defaultArcDownRight) +
			string(defaultHorizontal) +
			pieLegendTitlePadding +
			pieLegendTitle +
			pieLegendTitlePadding +
			strings.Repeat(
				string(defaultHorizontal),
				width-len([]rune(pieLegendTitle)),
			) +
			string(defaultHorizontal) +
			string(defaultArcDownLeft),
		pieLegendRowPrefix +
			padRight(items[0], width) +
			pieLegendRowSuffix,
		string(defaultRightTick) +
			strings.Repeat(
				string(defaultHorizontal),
				width+2,
			) +
			string(pieLegendLeftTick),
	}

	for _, item := range items[1:] {
		lines = append(lines, pieLegendRowPrefix+padRight(item, width)+pieLegendRowSuffix)
	}

	lines = append(lines, string(defaultArcUpRight)+strings.Repeat(string(defaultHorizontal), width+2)+string(defaultArcUpLeft))

	return lines
}

func padRight(text string, width int) string {
	padding := width - len([]rune(text))
	if padding <= 0 {
		return text
	}

	return text + strings.Repeat(string(pieEmptyRune), padding)
}

func pieAngleClockwiseFromTop(dx, dy float64) float64 {
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

func sliceIndexForAngle(slices []pieSlice, angle float64) int {
	for i, s := range slices {
		isLast := i == len(slices)-1
		if angle >= s.start && (angle < s.end || isLast) {
			return i
		}
	}

	return len(slices) - 1
}
