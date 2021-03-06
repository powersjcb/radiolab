package views

import (
	"github.com/wcharczuk/go-chart"
	"os"
)

// buf contains pairs of 8bit signed IQ values
func IQPlot(points []complex128) error {
	xValues := make([]float64, len(points))
	yValues := make([]float64, len(points))
	for i := 0; i < len(points); i++ {
		xValues[i] = float64(i)
		yValues[i] = real(points[i])
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name:           "time",
		},
		YAxis: chart.YAxis{
			Name:           "I",
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					StrokeWidth: chart.Disabled,
					DotWidth: 1,
				},
				XValues:         xValues,
				YValues:         yValues,

			},
		},
	}

	f, _ := os.Create("iq.png")
	defer f.Close()
	return graph.Render(chart.PNG, f)
}
