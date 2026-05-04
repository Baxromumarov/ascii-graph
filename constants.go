package main

const (
	cellWidth          = 4
	defaultChartHeight = 10
	pieMinRadius       = 5
	pieAspectX         = 2.0
	pieOuterPadding    = 0.25
	pieBorderThreshold = 0.45
)

const (
	pointRune           = '●'
	barRune             = '█'
	defaultHorizontal   = '─'
	defaultVertical     = '│'
	defaultArcDownRight = '╭'
	defaultArcDownLeft  = '╮'
	defaultArcUpRight   = '╰'
	defaultArcUpLeft    = '╯'
	defaultEndCap       = '╴'
	defaultStartCap     = '╶'
	defaultUpRight      = '└'
	defaultDownTick     = '┬'
	defaultRightTick    = '├'
	pieEmptyRune        = ' '
	pieBorderRune       = '•'
	pieFallbackSymbol   = '•'
	pieLegendLeftTick   = '┤'
)

const (
	pieErrNegativeValues       = "pie chart requires non-negative values"
	pieErrTotalGreaterThanZero = "pie chart requires a total greater than zero"
	pieErrAtLeastOnePositive   = "pie chart requires at least one positive value"
	pieSymbolPool              = "█▓▒░#@%*+=-~"
	pieTitleFormat             = "%s\n\n"
	pieRowFormat               = "%-*s  %s\n"
	pieTotalLabelFormat        = "Total: %s"
	pieLegendItemFormat        = "%c  %-12s %8s  %6.2f%%"
	pieLegendTitle             = "Pie Summary"
	pieLegendTitlePadding      = " "
	pieLegendRowPrefix         = "│ "
	pieLegendRowSuffix         = " │"
)

// Direction bitmasks for line routing
const (
	dirUp uint8 = 1 << iota
	dirRight
	dirDown
	dirLeft
)

// This is came from here: https://github.com/guptarohit/asciigraph/blob/master/options.go
type CharSet struct {
	Horizontal   string
	Vertical     string
	ArcDownRight string
	ArcDownLeft  string
	ArcUpRight   string
	ArcUpLeft    string
	EndCap       string
	StartCap     string
	UpRight      string
	DownTick     string
	RightTick    string
}

type CharRunes struct {
	Horizontal   rune
	Vertical     rune
	ArcDownRight rune
	ArcDownLeft  rune
	ArcUpRight   rune
	ArcUpLeft    rune
	EndCap       rune
	StartCap     rune
	UpRight      rune
	DownTick     rune
	RightTick    rune
}

var DefaultCharSet = CharSet{
	Horizontal:   "─",
	Vertical:     "│",
	ArcDownRight: "╭",
	ArcDownLeft:  "╮",
	ArcUpRight:   "╰",
	ArcUpLeft:    "╯",
	EndCap:       "╴",
	StartCap:     "╶",
	UpRight:      "└",
	DownTick:     "┬",
	RightTick:    "├",
}

func (c CharSet) getRune(field string, fallback rune) rune {
	if field == "" {
		return fallback
	}

	r := []rune(field)
	if len(r) == 0 {
		return fallback
	}

	return r[0]
}

func (c CharSet) Runes() CharRunes {
	return CharRunes{
		Horizontal:   c.getRune(c.Horizontal, defaultHorizontal),
		Vertical:     c.getRune(c.Vertical, defaultVertical),
		ArcDownRight: c.getRune(c.ArcDownRight, defaultArcDownRight),
		ArcDownLeft:  c.getRune(c.ArcDownLeft, defaultArcDownLeft),
		ArcUpRight:   c.getRune(c.ArcUpRight, defaultArcUpRight),
		ArcUpLeft:    c.getRune(c.ArcUpLeft, defaultArcUpLeft),
		EndCap:       c.getRune(c.EndCap, defaultEndCap),
		StartCap:     c.getRune(c.StartCap, defaultStartCap),
		UpRight:      c.getRune(c.UpRight, defaultUpRight),
		DownTick:     c.getRune(c.DownTick, defaultDownTick),
		RightTick:    c.getRune(c.RightTick, defaultRightTick),
	}
}

type Option string

const (
	BarChart    Option = "BarChart"
	LineGraph   Option = "LineGraph"
	PieChart    Option = "PieChart"
	ScatterPlot Option = "ScatterPlot"
)
