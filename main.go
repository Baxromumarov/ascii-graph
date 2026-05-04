package main

import "fmt"

func main() {

	if err := demoScatterFloat(); err != nil {
		fmt.Println("scatter error:", err)
	}
	fmt.Println()

	if err := demoLineInt(); err != nil {
		fmt.Println("line error:", err)
	}
	fmt.Println()

	if err := demoBarMixedSign(); err != nil {
		fmt.Println("bar error:", err)
	}

	fmt.Println()

	if err := demoPieChart(); err != nil {
		fmt.Println("pie error:", err)
	}
}

func demoScatterFloat() error {
	values := []float64{1.2, 4.8, 3.4, 6.1, 2.9, -2}
	labels := []string{"A", "B", "C", "D", "E", "F"}

	graph := NewGraph(values, ScatterPlot).
		WithXVals(labels).
		WithHeight(8)

	return graph.Draw()
}

func demoLineInt() error {
	values := []int{2, 5, 3, 8, 7, 9}
	labels := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

	graph := NewGraph(values, LineGraph).
		WithXVals(labels).
		WithHeight(10)

	return graph.Draw()
}

func demoBarMixedSign() error {
	values := []int{-3, 2, 5, -1, 4}
	labels := []string{"Q1", "Q2", "Q3", "Q4", "Q5"}

	graph := NewGraph(values, BarChart).
		WithXVals(labels).
		WithHeight(10)

	return graph.Draw()
}

func demoPieChart() error {
	values := []float32{20, 35, 15, 30}
	labels := []string{"Backend", "Frontend", "Infra", "QA"}

	graph := NewGraph(values, PieChart).
		WithXVals(labels)

	return graph.Draw()
}
