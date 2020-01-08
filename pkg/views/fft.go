package views

import (
	"github.com/powersjcb/radiolab/pkg/lib"
	"github.com/wcharczuk/go-chart"
	"os"
)

func FFTPlot(fft []lib.SpectralPoint) error {
	xValues := make([]float64, len(fft))
	yValues := make([]float64, len(fft))
	for i := 0; i < len(fft); i++ {
		xValues[i] = fft[i].Frequency
		yValues[i] = real(fft[i].Value)
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name:           "Frequency, Hz",
		},
		YAxis: chart.YAxis{
			Name:           "",
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

	f, _ := os.Create("fft.png")
	defer f.Close()
	return graph.Render(chart.PNG, f)
}
