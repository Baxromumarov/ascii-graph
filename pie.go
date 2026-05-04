package main

import (
	"fmt"
	"math"
	"strings"
)

type PieRenderer struct{}

func (p *PieRenderer) Option() Option { return PieChart }

func (p *PieRenderer) Render(g *Graph) error {
	total := 0.0
	for _, v := range g.yVals {
		if v < 0 {
			return fmt.Errorf("pie chart requires non-negative values")
		}

		total += v
	}

	if total == 0 {
		return fmt.Errorf("pie chart requires a total greater than zero")
	}

	fmt.Printf("%s\n\n", PieChart)
	fmt.Printf("Total: %s\n\n", formatNumber(total))

	for i, v := range g.yVals {
		percent := (v / total) * 100
		barLen := int(math.Round((percent / 100) * pieBarWidth))
		if barLen == 0 && v > 0 {
			barLen = 1
		}

		barText := strings.Repeat(string(barRune), barLen)

		fmt.Printf(
			"%-12s %8s (%6.2f%%) %s\n",
			g.labelFor(i),
			formatNumber(v),
			percent,
			barText,
		)
	}
	return nil
}
