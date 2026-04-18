package lab2

import (
	"math"
	"net/http"
	"universitySignalTransformation/pkg/utils"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

const (
	Tc   = 1.0
	f    = float64(50)
	fs   = float64(1000)
	step = 50
	A    = float64(2)

	Tc1   = 1.0
	fOld  = 50.0
	fs1   = 1000.0
	step1 = 50
	A1    = 2.0

	Tc2   = 1.0
	fs2   = 2000.0
	step2 = 50
	A2    = 2.0
)

func signalGenerationExerise1(Tc, f, fs, A float64) []float64 {
	N := int(math.Round(Tc * fs))
	sliceOfData := make([]float64, 0, N)

	for n := 0; n < N; n++ {
		t := float64(n) / fs
		x := A * math.Sin(2*math.Pi*f*t)
		sliceOfData = append(sliceOfData, x)
	}
	return sliceOfData
}

func signalGenerationExerise1Spectrum(x_re, x_im []float64) []float64 {
	N := len(x_re)
	M := make([]float64, 0, N/2)

	for k := 0; k < N/2; k++ {
		M = append(M, math.Sqrt(x_re[k]*x_re[k]+x_im[k]*x_im[k]))
	}

	return M
}

func signalGenerationExerise2Spectrum_dB(m_k []float64) []float64 {
	sliceOfData := make([]float64, 0, len(m_k))
	for i := 0; i < len(m_k); i++ {
		m_k_amp := float64(10) * math.Log10(m_k[i])
		sliceOfData = append(sliceOfData, m_k_amp)
	}
	return sliceOfData
}

func signalGenerationExerise2Signal(Tc, fs, A float64) []float64 {
	N := int(math.Round(Tc * fs))
	sliceOfData := make([]float64, 0, N)

	f1 := float64(10)
	f2 := fs/float64(2) - f1
	f3 := f1 / 2

	for n := 0; n < N; n++ {
		t := float64(n) / fs
		x := A*math.Sin(2*math.Pi*f1*t) + A*math.Sin(2*math.Pi*f2*t) + A*math.Sin(2*math.Pi*f3*t)
		sliceOfData = append(sliceOfData, x)
	}

	return sliceOfData
}

func DrawExercise1(w http.ResponseWriter, _ *http.Request) {
	x := signalGenerationExerise1(Tc, f, fs, A)
	xRe, xIm := utils.DFT(x)
	M := signalGenerationExerise1Spectrum(xRe, xIm)

	line := charts.NewLine()
	utils.SetChartOptions(line, "Laboratorium 2", "Zadanie 1", "Czas [s]")
	line.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("X", utils.FromSliceToLineData(x)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	spectrum := charts.NewLine()
	utils.SetChartOptions(spectrum, "Laboratorium 2", "Zadanie 1", "Częstotliwość [Hz]")
	spectrum.SetXAxis(utils.FrequencyAxisLabels(Tc, fs, step)).
		AddSeries("M", utils.FromSliceToLineData(M)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	page := components.NewPage()
	page.AddCharts(line, spectrum)
	page.Render(w)
}

func DrawExercise2(w http.ResponseWriter, _ *http.Request) {
	x1 := signalGenerationExerise1(Tc1, fOld, fs1, A1)
	x1Re, x1Im := utils.DFT(x1)
	M1 := signalGenerationExerise1Spectrum(x1Re, x1Im)
	M1dB := signalGenerationExerise2Spectrum_dB(M1)

	chart1 := charts.NewLine()
	utils.SetChartOptions(chart1, "Laboratorium 2", "Zadanie 2 dla widma z zadania 1", "Częstotliwość [Hz]")
	chart1.SetXAxis(utils.FrequencyAxisLabels(Tc1, fs1, step1)).
		AddSeries("M1", utils.FromSliceToLineData(M1dB)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	x2 := signalGenerationExerise2Signal(Tc2, fs2, A2)
	x2Re, x2Im := utils.DFT(x2)
	M2 := signalGenerationExerise1Spectrum(x2Re, x2Im)
	M2dB := signalGenerationExerise2Spectrum_dB(M2)

	chart2 := charts.NewLine()
	utils.SetChartOptions(chart2, "Laboratorium 2", "Zadanie 2 dla trzech tonów - skala liniowa", "Częstotliwość [Hz]")
	chart2.SetXAxis(utils.FrequencyAxisLabels(Tc2, fs2, step2)).
		AddSeries("M2", utils.FromSliceToLineData(M2dB)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart3 := charts.NewScatter()
	chart3.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWonderland}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Laboratorium 2",
			Subtitle: "Zadanie 2 dla trzech tonów - skala logarytmiczna",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name:         "Częstotliwość [Hz]",
			Type:         "log",
			NameLocation: "middle",
			NameGap:      30,
			AxisLabel: &opts.AxisLabel{
				Show: opts.Bool(true),
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
		}))

	chart3.AddSeries("M3", utils.LogSpectrumPoints(Tc2, fs2, M2dB)).
		SetSeriesOptions(charts.WithScatterChartOpts(opts.ScatterChart{SymbolSize: 3, Symbol: "circle"}))

	page := components.NewPage()
	page.AddCharts(chart1)
	page.Render(w)
}
