package lab6

import (
	"math"
	"net/http"
	"universitySignalTransformation/pkg/utils"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

const (
	A1   = 0.5
	A2   = 1.0
	step = 1000
	W    = 2
	Tc   = 1
)

func SignalGenerationExerise(bits []int, formula string) []float64 {
	B := len(bits)
	Tb := Tc / float64(B)

	fn := W / Tb
	fn1 := (W + 1) / Tb
	fn2 := (W + 2) / Tb
	fs := 1000 * fn

	N := int(math.Round(Tc * fs))
	sliceOfData := make([]float64, 0)

	for n := 0; n <= N; n++ {
		t := float64(n) / fs
		bitIndex := int(math.Floor(t / Tb))
		if bitIndex >= B {
			bitIndex = B - 1
		}
		bit := bits[bitIndex]
		var value float64
		if formula == "ASK" {
			if bit == 0 {
				value = A1 * math.Sin(2*math.Pi*fn*t)
			} else if bit == 1 {
				value = A2 * math.Sin(2*math.Pi*fn*t)
			}
		} else if formula == "FSK" {
			if bit == 0 {
				value = math.Sin(2 * math.Pi * fn1 * t)
			} else if bit == 1 {
				value = math.Sin(2 * math.Pi * fn2 * t)
			}
		} else if formula == "PSK" {
			if bit == 0 {
				value = math.Sin(2 * math.Pi * fn * t)
			} else if bit == 1 {
				value = math.Sin(2*math.Pi*fn*t + math.Pi)
			}
		}
		sliceOfData = append(sliceOfData, value)
	}
	return sliceOfData
}

func DrawExercise(w http.ResponseWriter, _ *http.Request) {
	bits := utils.ASCII_to_bit("KOSIARKA")

	//Limit bytes to 10 chars
	bits = bits[:10]

	ASK := SignalGenerationExerise(bits, "ASK")
	FSK := SignalGenerationExerise(bits, "FSK")
	PSK := SignalGenerationExerise(bits, "PSK")

	B := len(bits)
	Tb := Tc / float64(B)
	fn := W / Tb
	fs := 1000 * fn

	chart1 := charts.NewLine()

	utils.SetChartOptions(chart1, "Laboratorium 6", "Część 1 - ASK, zA(t)", "Czas [s]")
	chart1.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("ASK - zA(t)", utils.FromSliceToLineData(ASK)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{
			Smooth:     opts.Bool(false),
			ShowSymbol: opts.Bool(false),
		}))
	chart2 := charts.NewLine()
	utils.SetChartOptions(chart2, "Laboratorium 6", "Część 1 - PSK, zP(t)", "Czas [s]")
	chart2.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("PSK - zP(t)", utils.FromSliceToLineData(PSK)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{
			Smooth:     opts.Bool(false),
			ShowSymbol: opts.Bool(false),
		}))
	chart3 := charts.NewLine()
	utils.SetChartOptions(chart3, "Laboratorium 6", "Część 1 - FSK, zF(t)", "Czas [s]")
	chart3.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("FSK - zF(t)", utils.FromSliceToLineData(FSK)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{
			Smooth:     opts.Bool(false),
			ShowSymbol: opts.Bool(false),
		}))
	page := components.NewPage()
	page.AddCharts(chart1, chart2, chart3)
	page.Render(w)
}
