package lab13

import (
	"net/http"
	"universitySignalTransformation/pkg/lab10"
	"universitySignalTransformation/pkg/lab11"
	"universitySignalTransformation/pkg/lab12"
	"universitySignalTransformation/pkg/lab6"
	"universitySignalTransformation/pkg/lab8"
	"universitySignalTransformation/pkg/lab9"
	"universitySignalTransformation/pkg/utils"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type BER3DPoint struct {
	Alpha float64
	Beta  float64
	BER   float64
}

func ConnectedTransmissionSystem(configuaration string, inputWord string, coderChoose string, signalModulator string, alpha float64, beta float64) ([]byte, []byte, []byte) {
	var HammingCodeOutput []byte = make([]byte, 0)

	if coderChoose == "Hamming_7_4" {
		HammingCodeOutput = lab10.HammingCode(inputWord)
	} else if coderChoose == "Hamming_15_11" {
		HammingCodeOutput = lab11.HammingCode(inputWord)
	}

	// Modulation
	ModulatorInput := utils.FromBitSliceToIntSlice(HammingCodeOutput)
	ModulatorOutput := lab6.SignalGenerationExercise(ModulatorInput, signalModulator)

	// Connected transmission channels
	var yt []float64

	if configuaration == "1" {
		yt1 := lab12.TransmisionChannel(ModulatorOutput, alpha)
		yt = lab12.TransmisionChannel2(yt1, beta)
	} else if configuaration == "2" {
		yt1 := lab12.TransmisionChannel2(ModulatorOutput, beta)
		yt = lab12.TransmisionChannel(yt1, alpha)
	} else {
		panic("Invalid configuation")
	}

	// Demodulation
	var DemodulatorOutput []int

	if signalModulator == "ASK" || signalModulator == "PSK" {
		_, _, _, _, DemodulatorOutput = lab8.Demodulator(ModulatorInput, signalModulator, yt)
	} else if signalModulator == "FSK" {
		_, _, _, _, _, _, _, DemodulatorOutput = lab9.Demodulator(ModulatorInput, yt)
	}

	// Decoding
	HammingDecodeInput := utils.FromIntSliceToBitSlice(DemodulatorOutput)

	var HammingDecodeOutput []byte = make([]byte, 0)

	if coderChoose == "Hamming_7_4" {
		HammingDecodeOutput = lab10.HammingDecode(HammingDecodeInput)
	} else if coderChoose == "Hamming_15_11" {
		HammingDecodeOutput = lab11.HammingDecode(HammingDecodeInput)
	}

	return HammingDecodeOutput, HammingCodeOutput, HammingDecodeInput
}

func GenerateBER3DDataExercise13Part1(inputWord string, coderChoose string, signalModulator string, configuaration string) []BER3DPoint {
	points := make([]BER3DPoint, 0)

	for i := 0; i < 10; i++ {
		alpha := float64(i) * 30.0 / 9.0

		for j := 0; j < 10; j++ {
			beta := float64(j) * 1.0 / 9.0

			_, inputBits, outputBits := ConnectedTransmissionSystem(configuaration, inputWord, coderChoose, signalModulator, alpha, beta)
			ber := lab12.BER(inputBits, outputBits)

			points = append(points, BER3DPoint{
				Alpha: alpha,
				Beta:  beta,
				BER:   ber,
			})
		}
	}

	return points
}
func FromBER3DPointsToChart3DData(points []BER3DPoint) []opts.Chart3DData {
	data := make([]opts.Chart3DData, 0)

	for _, point := range points {
		data = append(data, opts.Chart3DData{
			Value: []interface{}{
				point.Alpha,
				point.Beta,
				point.BER,
			},
		})
	}

	return data
}

func DrawBER3DChartExercise13Part1(inputWord string, coderChoose string, signalModulator string, configuaration string) *charts.Bar3D {
	points := GenerateBER3DDataExercise13Part1(inputWord, coderChoose, signalModulator, configuaration)

	bar3d := charts.NewBar3D()

	bar3d.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Lab13 - Ćwiczenie 1",
			Subtitle: coderChoose + " - " + signalModulator + " - konfiguracja " + configuaration + " - Słowo: " + inputWord,
		}),
		charts.WithVisualMapOpts(opts.VisualMap{Min: 0, Max: 1, Calculable: opts.Bool(true)}),
		charts.WithXAxis3DOpts(opts.XAxis3D{Name: "Alpha", Type: "value"}),
		charts.WithYAxis3DOpts(opts.YAxis3D{Name: "Beta", Type: "value"}),
		charts.WithZAxis3DOpts(opts.ZAxis3D{Name: "BER", Type: "value"}),
	)

	bar3d.AddSeries("BER", FromBER3DPointsToChart3DData(points))

	return bar3d
}

func DrawExercise13Part1Hamming_7_4_Config1(w http.ResponseWriter, r *http.Request) {
	inputWord := r.URL.Query().Get("word")

	chartASK := DrawBER3DChartExercise13Part1(inputWord, "Hamming_7_4", "ASK", "1")
	chartPSK := DrawBER3DChartExercise13Part1(inputWord, "Hamming_7_4", "PSK", "1")
	chartFSK := DrawBER3DChartExercise13Part1(inputWord, "Hamming_7_4", "FSK", "1")

	page := components.NewPage()
	page.AddCharts(chartASK, chartPSK, chartFSK)
	page.Render(w)
}

func DrawExercise13Part1Hamming_7_4_Config2(w http.ResponseWriter, r *http.Request) {
	inputWord := r.URL.Query().Get("word")

	chartASK := DrawBER3DChartExercise13Part1(inputWord, "Hamming_7_4", "ASK", "2")
	chartPSK := DrawBER3DChartExercise13Part1(inputWord, "Hamming_7_4", "PSK", "2")
	chartFSK := DrawBER3DChartExercise13Part1(inputWord, "Hamming_7_4", "FSK", "2")

	page := components.NewPage()
	page.AddCharts(chartASK, chartPSK, chartFSK)
	page.Render(w)
}

func DrawExercise13Part1Hamming_15_11_Config1(w http.ResponseWriter, r *http.Request) {
	inputWord := r.URL.Query().Get("word")

	chartASK := DrawBER3DChartExercise13Part1(inputWord, "Hamming_15_11", "ASK", "1")
	chartPSK := DrawBER3DChartExercise13Part1(inputWord, "Hamming_15_11", "PSK", "1")
	chartFSK := DrawBER3DChartExercise13Part1(inputWord, "Hamming_15_11", "FSK", "1")

	page := components.NewPage()
	page.AddCharts(chartASK, chartPSK, chartFSK)
	page.Render(w)
}

func DrawExercise13Part1Hamming_15_11_Config2(w http.ResponseWriter, r *http.Request) {
	inputWord := r.URL.Query().Get("word")

	chartASK := DrawBER3DChartExercise13Part1(inputWord, "Hamming_15_11", "ASK", "2")
	chartPSK := DrawBER3DChartExercise13Part1(inputWord, "Hamming_15_11", "PSK", "2")
	chartFSK := DrawBER3DChartExercise13Part1(inputWord, "Hamming_15_11", "FSK", "2")

	page := components.NewPage()
	page.AddCharts(chartASK, chartPSK, chartFSK)
	page.Render(w)
}
