package lab1

import (
	"fmt"
	"math"
	"universitySignalTransformation/pkg/utils"

	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

const (
	Tc   = 1.7
	f    = 50
	fs   = 8000
	step = 800
	phi  = math.Pi / 2
)

//Numery wybranych funkcji 13,3,13,13

func signalGenerationExerise1(Tc, f, fs, phi float64, formula string) []float64 {
	N := int(math.Round(Tc * fs))
	sliceOfData := make([]float64, 0)
	for n := 0; n <= N; n++ {
		t := float64(n) / fs
		x := 0.9*math.Sin(2*math.Pi*f*t+phi)*math.Cos(21*math.Pi*t) + (t - 0.66*t)
		y := (math.Pow(t, 3) - 1) + math.Cos(4*math.Pow(t, 2)*math.Pi)*t
		z := x / (math.Abs(y*math.Cos(5*t)-x*y) + 3)
		v := (x * 662) / (math.Abs(x-y) + 0.5)

		if formula == "x" {
			sliceOfData = append(sliceOfData, x)
		} else if formula == "y" {
			sliceOfData = append(sliceOfData, y)
		} else if formula == "z" {
			sliceOfData = append(sliceOfData, z)
		} else if formula == "v" {
			sliceOfData = append(sliceOfData, v)
		} else {
			fmt.Println("Wrong formula argument")
			return nil
		}

	}
	return sliceOfData
}

func signalGenerationExerise2(Tc, fs, h float64, formula string) []float64 {
	N := int(math.Round(Tc * fs))
	sliceOfData := make([]float64, 0)
	for n := 0; n <= N; n++ {
		t := float64(n) / fs
		if formula == "u" {
			if t >= 0 && t < 0.3 {
				u := 0.1*math.Cos(36*math.Pi*t) + math.Sin(22*math.Pi*t)
				sliceOfData = append(sliceOfData, u)
			} else if t >= 0.3 && t < 0.8 {
				u := (t - 0.3) * math.Cos(26*math.Pi*t+math.Sin(12*t))
				sliceOfData = append(sliceOfData, u)
			} else if t >= 0.8 && t < 1 {
				u := 0.1 * (math.Log2(t+2)*math.Sin(6*math.Pi*t) + math.Log2(math.Cos(44*math.Pi*t)+2))
				sliceOfData = append(sliceOfData, u)
			}
		} else if formula == "b" {
			b1 := float64(0)
			for i := 0; i < int(h); i++ {
				calc := math.Sin((6*math.Pi+1)*t*math.Pi) * math.Sin(math.Pow(h, 3)) / (6*h + 2)
				b1 = b1 + calc
			}
			sliceOfData = append(sliceOfData, b1)

		} else {
			fmt.Println("Wrong formula argument")
			return nil
		}

	}
	return sliceOfData
}

func DrawExercise1(w http.ResponseWriter, _ *http.Request) {
	x := signalGenerationExerise1(Tc, f, fs, phi, "x")
	y := signalGenerationExerise1(Tc, f, fs, phi, "y")
	z := signalGenerationExerise1(Tc, f, fs, phi, "z")
	v := signalGenerationExerise1(Tc, f, fs, phi, "v")

	chart1 := charts.NewLine()
	utils.SetChartOptions(chart1, "Laboratorium 1", "Zadanie 1", "Czas [s]")
	chart1.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("x(t)", utils.FromSliceToLineData(x)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart2 := charts.NewLine()
	utils.SetChartOptions(chart2, "Laboratorium 1", "Zadanie 1", "Czas [s]")
	chart2.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("y(t)", utils.FromSliceToLineData(y)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart3 := charts.NewLine()
	utils.SetChartOptions(chart3, "Laboratorium 1", "Zadanie 1", "Czas [s]")
	chart3.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("z(t)", utils.FromSliceToLineData(z)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart4 := charts.NewLine()
	utils.SetChartOptions(chart4, "Laboratorium 1", "Zadanie 1", "Czas [s]")
	chart4.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("v(t)", utils.FromSliceToLineData(v)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	page := components.NewPage()
	page.AddCharts(chart1, chart2, chart3, chart4)
	page.Render(w)
}

func DrawExercise2(w http.ResponseWriter, _ *http.Request) {

	U := signalGenerationExerise2(Tc, fs, 2, "u")
	b1 := signalGenerationExerise2(Tc, fs, 2, "b")
	b2 := signalGenerationExerise2(Tc, fs, 4, "b")
	b3 := signalGenerationExerise2(Tc, fs, 8, "b")

	chart1 := charts.NewLine()
	utils.SetChartOptions(chart1, "Laboratorium 1", "Zadanie 2", "Czas [s]")
	chart1.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("x(t)", utils.FromSliceToLineData(U)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart2 := charts.NewLine()
	utils.SetChartOptions(chart2, "Laboratorium 1", "Zadanie 2", "Czas [s]")
	chart2.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("b1(t)", utils.FromSliceToLineData(b1)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart3 := charts.NewLine()
	utils.SetChartOptions(chart3, "Laboratorium 1", "Zadanie 2", "Czas [s]")
	chart3.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("b2(t)", utils.FromSliceToLineData(b2)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart4 := charts.NewLine()
	utils.SetChartOptions(chart4, "Laboratorium 1", "Zadanie 2", "Czas [s]")
	chart4.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("b3(t)", utils.FromSliceToLineData(b3)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	page := components.NewPage()
	page.AddCharts(chart1, chart2, chart3, chart4)
	page.Render(w)
}
