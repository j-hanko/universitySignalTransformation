/*
Liczba prążków w zakresie (0, fs/2):
f = 10 fs = 1024 H = 3
Mx: 3
My: 3
Mz: 3

Liczba prążków w zakresie (0, fs/2):
f = 10 fs = 100 H = 5
Mx: 4
My: 2
Mz: 2

Liczba prążków w zakresie (0, fs/2):
f = 50 fs = 900 H = 15
Mx: 8
My: 4
Mz: 4
*/
package lab3

import (
	"fmt"
	"math"
	"net/http"
	"universitySignalTransformation/pkg/utils"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

const (
	Tc    = 1
	fs    = 1024
	step  = 64
	f     = 10
	H     = 3
	f1    = 10
	f2    = 40
	alpha = 2.0
	beta  = 3.0
)

func signalGenerationExerise1(Tc, f, fs float64, H int, formula string) []float64 {
	N := int(math.Round(Tc * fs))
	sliceOfData := make([]float64, 0)
	for n := 0; n < N; n++ {
		t := float64(n) / fs
		if formula == "x" {
			x := float64(0)
			for k := 1; k <= H; k++ {
				calc := math.Pow(-1, float64(k+1)) * math.Sin(2.0*math.Pi*float64(k)*f*t) / float64(k)
				x += calc
			}
			x = 2.0 / math.Pi * x
			sliceOfData = append(sliceOfData, x)
		} else if formula == "y" {
			y := float64(0)
			for k := 1; k <= H; k++ {
				calc := math.Pow(-1, float64(k+1)) * math.Sin(2*math.Pi*float64(2*k-1)*f*t) / math.Pow(float64(2*k-1), 2)
				y += calc
			}
			y = 8.0 / math.Pow(math.Pi, 2) * y
			sliceOfData = append(sliceOfData, y)
		} else if formula == "z" {
			z := float64(0)
			for k := 1; k <= H; k++ {
				calc := math.Sin(2*math.Pi*float64(2*k-1)*f*t) / float64(2*k-1)
				z += calc
			}
			z = 4.0 / math.Pi * z
			sliceOfData = append(sliceOfData, z)
		} else {
			fmt.Println("Wrong formula argument")
			return nil
		}

	}
	return sliceOfData
}

func signalGenerationExerise2(Tc, fs, f1, f2, alpha, beta float64, formula string) []float64 {
	N := int(math.Round(Tc * fs))
	sliceOfData := make([]float64, 0, N)

	for n := 0; n < N; n++ {
		t := float64(n) / fs
		x := 0.5 * math.Sin(2*math.Pi*f1*t)
		y := math.Sin(2*math.Pi*f2*t) + 0.7*math.Sin(2*math.Pi*f1*t)
		z := alpha*x + beta*y

		if formula == "x" {
			sliceOfData = append(sliceOfData, x)
		} else if formula == "y" {
			sliceOfData = append(sliceOfData, y)
		} else if formula == "z" {
			sliceOfData = append(sliceOfData, z)
		} else {
			fmt.Println("Wrong formula argument")
			return nil
		}
	}
	return sliceOfData
}

func stripes(f, fs float64, H int, formula string) int {
	count := 0

	for k := 1; k <= H; k++ {
		var freq float64

		switch formula {
		case "x":
			freq = float64(k) * f
		case "y", "z":
			freq = float64(2*k-1) * f
		default:
			fmt.Println("Wrong formula argument")
			return 0
		}

		if freq > 0 && freq < fs/2 {
			count++
		}
	}

	return count
}

func Exercise1CountStripes() {
	fmt.Println("Liczba prążków w zakresie (0, fs/2):")
	fmt.Println("f =", f, "fs =", fs, "H =", H)
	fmt.Println("Mx:", stripes(f, fs, H, "x"))
	fmt.Println("My:", stripes(f, fs, H, "y"))
	fmt.Println("Mz:", stripes(f, fs, H, "z"))
}

func DrawExercise1(w http.ResponseWriter, _ *http.Request) {
	x := signalGenerationExerise1(Tc, f, fs, H, "x")
	y := signalGenerationExerise1(Tc, f, fs, H, "y")
	z := signalGenerationExerise1(Tc, f, fs, H, "z")

	xRe, xIm := utils.FFT(x)
	yRe, yIm := utils.FFT(y)
	zRe, zIm := utils.FFT(z)

	Mx := utils.Spectrum(xRe, xIm)
	My := utils.Spectrum(yRe, yIm)
	Mz := utils.Spectrum(zRe, zIm)

	chart1 := charts.NewLine()
	utils.SetChartOptions(chart1, "Laboratorium 3", "Zadanie 1 - x(t)", "Czas [s]")
	chart1.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("X", utils.FromSliceToLineData(x)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart2 := charts.NewLine()
	utils.SetChartOptions(chart2, "Laboratorium 3", "Zadanie 1 - y(t)", "Czas [s]")
	chart2.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("Y", utils.FromSliceToLineData(y)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart3 := charts.NewLine()
	utils.SetChartOptions(chart3, "Laboratorium 3", "Zadanie 1 - z(t)", "Czas [s]")
	chart3.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("Z", utils.FromSliceToLineData(z)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart4 := charts.NewLine()
	utils.SetChartOptions(chart4, "Laboratorium 3", "Zadanie 1 - Mx(f)", "Częstotliwość [Hz]")
	chart4.SetXAxis(utils.FrequencyAxisLabels(Tc, fs, step)).
		AddSeries("Mx", utils.FromSliceToLineData(Mx)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart5 := charts.NewLine()
	utils.SetChartOptions(chart5, "Laboratorium 3", "Zadanie 1 - My(f)", "Częstotliwość [Hz]")
	chart5.SetXAxis(utils.FrequencyAxisLabels(Tc, fs, step)).
		AddSeries("My", utils.FromSliceToLineData(My)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart6 := charts.NewLine()
	utils.SetChartOptions(chart6, "Laboratorium 3", "Zadanie 1 - Mz(f)", "Częstotliwość [Hz]")
	chart6.SetXAxis(utils.FrequencyAxisLabels(Tc, fs, step)).
		AddSeries("Mz", utils.FromSliceToLineData(Mz)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	page := components.NewPage()
	page.AddCharts(chart1, chart2, chart3, chart4, chart5, chart6)
	page.Render(w)
}

func DrawExercise2(w http.ResponseWriter, _ *http.Request) {
	x := signalGenerationExerise2(Tc, fs, f1, f2, alpha, beta, "x")
	y := signalGenerationExerise2(Tc, fs, f1, f2, alpha, beta, "y")
	z := signalGenerationExerise2(Tc, fs, f1, f2, alpha, beta, "z")

	xRe, xIm := utils.FFT(x)
	yRe, yIm := utils.FFT(y)
	zRe, zIm := utils.FFT(z)

	Mx := utils.Spectrum(xRe, xIm)
	My := utils.Spectrum(yRe, yIm)
	Mz := utils.Spectrum(zRe, zIm)

	MzHat := utils.MzHatSpectrum(alpha, beta, Mx, My)

	chart1 := charts.NewLine()
	utils.SetChartOptions(chart1, "Laboratorium 3", "Zadanie 2 - x(t)", "Czas [s]")
	chart1.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("x(t)", utils.FromSliceToLineData(x)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart2 := charts.NewLine()
	utils.SetChartOptions(chart2, "Laboratorium 3", "Zadanie 2 - y(t)", "Czas [s]")
	chart2.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("y(t)", utils.FromSliceToLineData(y)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart3 := charts.NewLine()
	utils.SetChartOptions(chart3, "Laboratorium 3", "Zadanie 2 - z(t)", "Czas [s]")
	chart3.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("z(t)", utils.FromSliceToLineData(z)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart4 := charts.NewLine()
	utils.SetChartOptions(chart4, "Laboratorium 3", "Zadanie 2 - Mx(f)", "Częstotliwość [Hz]")
	chart4.SetXAxis(utils.FrequencyAxisLabels(Tc, fs, step)).
		AddSeries("Mx(f)", utils.FromSliceToLineData(Mx)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart5 := charts.NewLine()
	utils.SetChartOptions(chart5, "Laboratorium 3", "Zadanie 2 - My(f)", "Częstotliwość [Hz]")
	chart5.SetXAxis(utils.FrequencyAxisLabels(Tc, fs, step)).
		AddSeries("My(f)", utils.FromSliceToLineData(My)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart6 := charts.NewLine()
	utils.SetChartOptions(chart6, "Laboratorium 3", "Zadanie 2 - Mz(f)", "Częstotliwość [Hz]")
	chart6.SetXAxis(utils.FrequencyAxisLabels(Tc, fs, step)).
		AddSeries("Mz(f)", utils.FromSliceToLineData(Mz)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart7 := charts.NewLine()
	utils.SetChartOptions(chart7, "Laboratorium 3", "Zadanie 2 - MzHat(f)", "Częstotliwość [Hz]")
	chart7.SetXAxis(utils.FrequencyAxisLabels(Tc, fs, step)).
		AddSeries("MzHat(f)", utils.FromSliceToLineData(MzHat)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	page := components.NewPage()
	page.AddCharts(chart1, chart2, chart3, chart4, chart5, chart6, chart7)
	page.Render(w)
}
