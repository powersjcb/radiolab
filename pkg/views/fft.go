package views

import (
	"github.com/powersjcb/radiolab/pkg/lib"
	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
	"io"
	"math/cmplx"
)

var gridStyle = chart.Style{
	Hidden: false,
	StrokeColor: drawing.ColorBlack,
	StrokeWidth: 0.5,
}

func FFTPlot(w io.Writer, fft []lib.SpectralPoint) error {
	xValues := make([]float64, len(fft))
	yValues := make([]float64, len(fft))
	for i := 0; i < len(fft); i++ {
		xValues[i] = fft[i].Frequency
		yValues[i] = cmplx.Abs(fft[i].Value)
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name:           "Frequency, Hz",
			GridMajorStyle: gridStyle,
			GridMinorStyle: gridStyle,
		},
		YAxis: chart.YAxis{
			Name:           "",
			GridMajorStyle: gridStyle,
			GridMinorStyle: gridStyle,
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

	return graph.Render(chart.PNG, w)
}
