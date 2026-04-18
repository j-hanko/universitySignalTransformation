package utils

import (
	"fmt"
	"math"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

// X axis options
func TimeAxisLabels(Tc, fs float64, step int) []string {
	N := int(math.Round(Tc * fs))
	labels := make([]string, 0, N+1)

	for n := 0; n <= N; n++ {
		if n%step == 0 {
			t := float64(n) / fs
			labels = append(labels, fmt.Sprintf("%.2f", t))
		} else {
			labels = append(labels, "")
		}
	}
	return labels
}

func FrequencyAxisLabels(Tc, fs float64, step int) []string {
	N := int(math.Round(Tc * fs))
	labels := make([]string, 0, N/2)

	for k := 0; k < N/2; k++ {
		if k%step == 0 {
			fk := float64(k) * fs / float64(N)
			labels = append(labels, fmt.Sprintf("%.2f", fk))
		} else {
			labels = append(labels, "")
		}
	}
	return labels
}

func LogSpectrumPoints(Tc, fs float64, spectrumDB []float64) []opts.ScatterData {
	N := int(math.Round(Tc * fs))
	points := make([]opts.ScatterData, 0, len(spectrumDB)-1)

	for k := 1; k < len(spectrumDB); k++ {
		fk := float64(k) * fs / float64(N)
		points = append(points, opts.ScatterData{Value: []interface{}{fk, spectrumDB[k]}})
	}

	return points
}

// Chart description and options
func SetChartOptions(line *charts.Line, title, subtitle, xAxisName string) {
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWonderland}),
		charts.WithTitleOpts(opts.Title{
			Title:    title,
			Subtitle: subtitle,
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name:         xAxisName,
			NameLocation: "middle",
			NameGap:      30,
			AxisLabel: &opts.AxisLabel{
				Show:     opts.Bool(true),
				Interval: "0",
			},
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name:         "Amplituda",
			NameLocation: "middle",
			NameGap:      50,
			Position:     "left",
			AxisLabel: &opts.AxisLabel{
				Show: opts.Bool(true),
			},
		}),
	)
}

func SetSpectrumChartOptions(line *charts.Line, title, subtitle string) {
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWonderland}),
		charts.WithTitleOpts(opts.Title{
			Title:    title,
			Subtitle: subtitle,
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name:         "Częstotliwość [Hz]",
			NameLocation: "middle",
			NameGap:      30,
			Min:          0,
			Max:          150,
			AxisLabel: &opts.AxisLabel{
				Show:     opts.Bool(true),
				Interval: "0",
			},
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name:         "Amplituda [dB]",
			NameLocation: "middle",
			NameGap:      50,
			Position:     "left",
			Min:          -80,
			Max:          30,
			AxisLabel: &opts.AxisLabel{
				Show: opts.Bool(true),
			},
		}),
	)
}
