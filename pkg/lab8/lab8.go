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

func Demodulator(ASCII_word, signal_choose string) ([]float64, []float64, []float64, []float64, []int) {
	bits := utils.ASCII_to_bit(ASCII_word)
	B := len(bits)
	Tb := Tc / float64(B)
	fn := W / Tb
	fs := int(1000 * fn)

	slice_x := make([]float64, 0)
	slice_p := make([]float64, 0)
	slice_c := make([]float64, 0)

	phase := 0.0
	h := 750.0
	if signal_choose == "PSK" {
		phase = math.Pi
		h = 0
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

	for i := 0; i < len(slice_p); i++ {
		if slice_p[i] > h {
			slice_c = append(slice_c, 1)
		} else {
			slice_c = append(slice_c, 0)
		}
	}

	decodedBits := utils.SignalToBits(slice_c, samplesPerBit)

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
	img1, img2, img3, img4, _ := Demodulator(ASCII_word, signalChoose)

	bits := utils.ASCII_to_bit(ASCII_word)
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
