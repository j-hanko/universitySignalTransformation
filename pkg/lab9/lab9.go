package lab9

import (
	"math"
	"net/http"
	"universitySignalTransformation/pkg/lab6"
	"universitySignalTransformation/pkg/utils"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

const (
	step = 5000
	W    = 2
	Tc   = 1
)

func Demodulator(bits []int) ([]float64, []float64, []float64, []float64, []float64, []float64, []float64, []int) {
	B := len(bits)
	Tb := Tc / float64(B)
	fn := W / Tb
	fn1 := (W + 1) / Tb
	fn2 := (W + 2) / Tb
	fs := int(1000 * fn)

	slice_x1 := make([]float64, 0)
	slice_x2 := make([]float64, 0)
	slice_p1 := make([]float64, 0)
	slice_p2 := make([]float64, 0)
	slice_p := make([]float64, 0)
	slice_c := make([]float64, 0)
	decodedBits := make([]int, 0)

	z := lab6.SignalGenerationExercise(bits, "FSK")
	for n := 0; n < len(z); n++ {
		t := float64(n) / float64(fs)
		x1 := z[n] * math.Sin(2.0*math.Pi*fn1*t)
		slice_x1 = append(slice_x1, x1)
		x2 := z[n] * math.Sin(2.0*math.Pi*fn2*t)
		slice_x2 = append(slice_x2, x2)
	}

	samplesPerBit := len(z) / B
	sum1 := 0.0
	sum2 := 0.0

	for i := 0; i < len(slice_x1); i++ {
		if i%samplesPerBit == 0 {
			sum1 = 0.0
		}
		sum1 += slice_x1[i]
		slice_p1 = append(slice_p1, sum1)
	}

	for i := 0; i < len(slice_x2); i++ {
		if i%samplesPerBit == 0 {
			sum2 = 0.0
		}
		sum2 += slice_x2[i]
		slice_p2 = append(slice_p2, sum2)
	}

	for i := 0; i < len(slice_p1); i++ {
		slice_p = append(slice_p, slice_p2[i]-slice_p1[i])
	}

	slice_c = make([]float64, len(slice_p))

	for bit := 0; bit < B; bit++ {
		start := bit * samplesPerBit
		end := (bit + 1) * samplesPerBit

		if end > len(slice_p) {
			end = len(slice_p)
		}

		mean := 0.0
		amount := 0
		for i := start; i < end; i++ {
			mean += slice_p[i]
			amount += 1
		}
		mean = mean / float64(amount)

		var decision float64
		var bitValue int

		if mean > 0 {
			decision = 1.0
			bitValue = 1
		} else {
			decision = 0.0
			bitValue = 0
		}

		decodedBits = append(decodedBits, bitValue)

		for i := start; i < end; i++ {
			slice_c[i] = decision
		}
	}

	return z, slice_x1, slice_x2, slice_p1, slice_p2, slice_p, slice_c, decodedBits
}

func DrawDemodulator(w http.ResponseWriter, _ *http.Request) {
	ASCII_word := "bot"
	bits := utils.ASCII_to_bit(ASCII_word)

	img1, img2, img3, img4, img5, img6, img7, _ := Demodulator(bits)

	B := len(bits)
	Tb := Tc / float64(B)
	fn := W / Tb
	fs := 1000 * fn
	realTc := float64(len(img1)) / float64(fs)

	chart1 := charts.NewLine()
	utils.SetChartOptions(chart1, "Laboratorium 9", "Zadanie 1", "Czas [s]")
	chart1.SetXAxis(utils.TimeAxisLabels(realTc, fs, step)).
		AddSeries("z(t)", utils.FromSliceToLineData(img1)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{
			Smooth:     opts.Bool(false),
			ShowSymbol: opts.Bool(false),
		}))

	chart2 := charts.NewLine()
	utils.SetChartOptions(chart2, "Laboratorium 9", "Zadanie 1", "Czas [s]")
	chart2.SetXAxis(utils.TimeAxisLabels(realTc, fs, step)).
		AddSeries("x1(t)", utils.FromSliceToLineData(img2)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{
			Smooth:     opts.Bool(false),
			ShowSymbol: opts.Bool(false),
		}))

	chart3 := charts.NewLine()
	utils.SetChartOptions(chart3, "Laboratorium 9", "Zadanie 1", "Czas [s]")
	chart3.SetXAxis(utils.TimeAxisLabels(realTc, fs, step)).
		AddSeries("x2(t)", utils.FromSliceToLineData(img3)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{
			Smooth:     opts.Bool(false),
			ShowSymbol: opts.Bool(false),
		}))

	chart4 := charts.NewLine()
	utils.SetChartOptions(chart4, "Laboratorium 9", "Zadanie 1", "Czas [s]")
	chart4.SetXAxis(utils.TimeAxisLabels(realTc, fs, step)).
		AddSeries("p1(t)", utils.FromSliceToLineData(img4)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{
			Smooth:     opts.Bool(false),
			ShowSymbol: opts.Bool(false),
		}))

	chart5 := charts.NewLine()
	utils.SetChartOptions(chart5, "Laboratorium 9", "Zadanie 1", "Czas [s]")
	chart5.SetXAxis(utils.TimeAxisLabels(realTc, fs, step)).
		AddSeries("p2(t)", utils.FromSliceToLineData(img5)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{
			Smooth:     opts.Bool(false),
			ShowSymbol: opts.Bool(false),
		}))

	chart6 := charts.NewLine()
	utils.SetChartOptions(chart6, "Laboratorium 9", "Zadanie 1", "Czas [s]")
	chart6.SetXAxis(utils.TimeAxisLabels(realTc, fs, step)).
		AddSeries("p(t)", utils.FromSliceToLineData(img6)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{
			Smooth:     opts.Bool(false),
			ShowSymbol: opts.Bool(false),
		}))

	chart7 := charts.NewLine()
	utils.SetChartOptions(chart7, "Laboratorium 9", "Zadanie 1", "Czas [s]")
	chart7.SetXAxis(utils.TimeAxisLabels(realTc, fs, step)).
		AddSeries("c(t)", utils.FromSliceToLineData(img7)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{
			Smooth:     opts.Bool(false),
			ShowSymbol: opts.Bool(false),
		}))

	page := components.NewPage()
	page.AddCharts(chart1, chart2, chart3, chart4, chart5, chart6, chart7)
	page.Render(w)
}
