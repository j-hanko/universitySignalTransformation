package lab8

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
	step = 2500
	A    = 1
	W    = 2
	Tc   = 1
)

func Demodulator(bits []int, signal_choose string) ([]float64, []float64, []float64, []float64, []int) {

	if signal_choose != "ASK" && signal_choose != "PSK" {
		panic("You can choose only ASK or PSK signal")
	}

	B := len(bits)
	Tb := Tc / float64(B)
	fn := W / Tb
	fs := int(1000 * fn)

	slice_x := make([]float64, 0)
	slice_p := make([]float64, 0)
	slice_c := make([]float64, 0)
	decodedBits := make([]int, 0)

	phase := 0.0
	h := 0.0
	if signal_choose == "PSK" {
		phase = math.Pi
	}

	z := lab6.SignalGenerationExercise(bits, signal_choose)
	for n := 0; n < len(z); n++ {
		t := float64(n) / float64(fs)
		x := z[n] * (float64(A) * math.Sin(2*math.Pi*fn*t+phase))
		slice_x = append(slice_x, x)
	}

	samplesPerBit := int(float64(fs) * Tb)
	sum := 0.0
	for i := 0; i < len(slice_x); i++ {
		if i%samplesPerBit == 0 {
			sum = 0.0
		}
		sum += slice_x[i]
		slice_p = append(slice_p, sum)
	}

	if signal_choose == "ASK" {
		h = 500
	}

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
		if mean >= h {
			decision = 1.0
			bitValue = 1
		} else {
			decision = 0.0
			bitValue = 0
		}

		decodedBits = append(decodedBits, bitValue)

		for i := start; i < end; i++ {
			slice_c = append(slice_c, decision)
		}
	}

	return z, slice_x, slice_p, slice_c, decodedBits
}

func DrawDemodulator(w http.ResponseWriter, r *http.Request) {
	ASCII_word := "bot"

	signalChoose := r.URL.Query().Get("signal")

	if signalChoose == "" {
		signalChoose = "ASK"
	}
	if signalChoose != "ASK" && signalChoose != "PSK" {
		http.Error(w, "Niepoprawny signal. Użyj ?signal=ASK albo ?signal=PSK", http.StatusBadRequest)
		return
	}
	bits := utils.ASCII_to_bit(ASCII_word)
	img1, img2, img3, img4, _ := Demodulator(bits, signalChoose)

	B := len(bits)
	Tb := Tc / float64(B)
	fn := W / Tb
	fs := 1000 * fn

	chart1 := charts.NewLine()
	utils.SetChartOptions(chart1, "Laboratorium 8", "Zadanie 1", "Czas [s]")
	chart1.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("z(t)", utils.FromSliceToLineData(img1)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{
			Smooth:     opts.Bool(false),
			ShowSymbol: opts.Bool(false),
		}))

	chart2 := charts.NewLine()
	utils.SetChartOptions(chart2, "Laboratorium 8", "Zadanie 1", "Czas [s]")
	chart2.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("x(t)", utils.FromSliceToLineData(img2)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{
			Smooth:     opts.Bool(false),
			ShowSymbol: opts.Bool(false),
		}))

	chart3 := charts.NewLine()
	utils.SetChartOptions(chart3, "Laboratorium 8", "Zadanie 1", "Czas [s]")
	chart3.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("p(t)", utils.FromSliceToLineData(img3)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{
			Smooth:     opts.Bool(false),
			ShowSymbol: opts.Bool(false),
		}))

	chart4 := charts.NewLine()
	utils.SetChartOptions(chart4, "Laboratorium 8", "Zadanie 1", "Czas [s]")
	chart4.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("c(t)", utils.FromSliceToLineData(img4)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{
			Smooth:     opts.Bool(false),
			ShowSymbol: opts.Bool(false),
		}))

	page := components.NewPage()
	page.AddCharts(chart1, chart2, chart3, chart4)
	page.Render(w)
}
