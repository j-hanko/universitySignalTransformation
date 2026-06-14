package lab12

import (
	"math/rand/v2"
	"net/http"
	"universitySignalTransformation/pkg/lab10"
	"universitySignalTransformation/pkg/lab11"
	"universitySignalTransformation/pkg/lab6"
	"universitySignalTransformation/pkg/lab8"
	"universitySignalTransformation/pkg/lab9"
	"universitySignalTransformation/pkg/utils"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func TransmisionChannel(input []float64, alpha float64) (output []float64) {
	output = make([]float64, len(input))

	for i := 0; i < len(input); i++ {
		noise := rand.Float64()*2 - 1
		output[i] = input[i] + alpha*noise
	}

	return output
}

func TransmisionChannel2(input []float64, beta float64) (output []float64) {
	output = make([]float64, len(input))

	if len(input) == 0 {
		return output
	}

	if beta == 0 {
		for i := 0; i < len(input); i++ {
			output[i] = input[i]
		}
		return output
	}

	for i := 0; i < len(input); i++ {
		t := float64(i) / float64(len(input))

		g := 0.0

		if beta < 1.0 && t < 1.0-beta {
			tmp := 1.0 - (t / (1.0 - beta))
			g = tmp * tmp
		} else {
			g = 0.0
		}

		output[i] = input[i] * g
	}

	return output
}

func SimpleTransmissionSystem(inputWord string, coderChoose string, signalModulator string, parameter float64, transmisionChannelChoose string) ([]byte, []byte, []byte) {
	//Coder
	var HammingCodeOutput []byte = make([]byte, 0)
	if coderChoose == "Hamming_7_4" {
		HammingCodeOutput = lab10.HammingCode(inputWord)
	} else if coderChoose == "Hamming_15_11" {
		HammingCodeOutput = lab11.HammingCode(inputWord)
	}

	//Modulation
	ModulatorInput := utils.FromBitSliceToIntSlice(HammingCodeOutput)
	ModulatorOutput := lab6.SignalGenerationExercise(ModulatorInput, signalModulator)

	//Transmision channel
	var yt []float64

	if transmisionChannelChoose == "Exercise1" {
		yt = TransmisionChannel(ModulatorOutput, parameter)
	} else if transmisionChannelChoose == "Exercise2" {
		yt = TransmisionChannel2(ModulatorOutput, parameter)
	}

	//Demodulation
	var DemodulatorOutput []int

	if signalModulator == "ASK" || signalModulator == "PSK" {
		_, _, _, _, DemodulatorOutput = lab8.Demodulator(ModulatorInput, signalModulator, yt)
	} else if signalModulator == "FSK" {
		_, _, _, _, _, _, _, DemodulatorOutput = lab9.Demodulator(ModulatorInput, yt)
	}

	//Decoding
	HammingDecodeInput := utils.FromIntSliceToBitSlice(DemodulatorOutput)

	var HammingDecodeOutput []byte = make([]byte, 0)
	if coderChoose == "Hamming_7_4" {
		HammingDecodeOutput = lab10.HammingDecode(HammingDecodeInput)
	} else if coderChoose == "Hamming_15_11" {
		HammingDecodeOutput = lab11.HammingDecode(HammingDecodeInput)
	}

	return HammingDecodeOutput, HammingCodeOutput, HammingDecodeInput
}

func BER(input []byte, output []byte) float64 {
	if len(input) == 0 {
		return 0
	}

	minLength := len(input)
	if len(output) < minLength {
		minLength = len(output)
	}

	errors := 0

	for i := 0; i < minLength; i++ {
		if input[i] != output[i] {
			errors++
		}
	}

	if len(input) > len(output) {
		errors += len(input) - len(output)
	} else if len(output) > len(input) {
		errors += len(output) - len(input)
	}

	return float64(errors) / float64(len(input))
}

func GenerateBERDataExercise1(inputWord string, coderChoose string, signalModulator string) ([]float64, []float64) {
	alphaValues := make([]float64, 0)
	berValues := make([]float64, 0)

	for i := 0; i < 10; i++ {
		alpha := float64(i) * 30.0 / 9.0

		_, inputBits, outputBits := SimpleTransmissionSystem(inputWord, coderChoose, signalModulator, alpha, "Exercise1")
		ber := BER(inputBits, outputBits)

		alphaValues = append(alphaValues, alpha)
		berValues = append(berValues, ber)
	}

	return alphaValues, berValues
}

func GenerateBERDataExercise2(inputWord string, coderChoose string, signalModulator string) ([]float64, []float64) {
	betaValues := make([]float64, 0)
	berValues := make([]float64, 0)

	for i := 0; i < 10; i++ {
		beta := float64(i) * 1.0 / 9.0

		_, inputBits, outputBits := SimpleTransmissionSystem(inputWord, coderChoose, signalModulator, beta, "Exercise2")
		ber := BER(inputBits, outputBits)

		betaValues = append(betaValues, beta)
		berValues = append(berValues, ber)
	}

	return betaValues, berValues
}

func DrawBERChartExercise1(inputWord string, coderChoose string, signalModulator string) *charts.Line {
	alphaValues, berValues := GenerateBERDataExercise1(inputWord, coderChoose, signalModulator)

	chart := charts.NewLine()

	utils.SetBERChartOptions(chart, "Lab12 - Ćwiczenie 1", coderChoose+" - "+signalModulator+" - Słowo informacyjne: "+inputWord, "Alpha")

	chart.SetXAxis(utils.AlphaAxisLabels(alphaValues)).
		AddSeries("BER", utils.FromSliceToLineData(berValues)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{
			Smooth:     opts.Bool(false),
			ShowSymbol: opts.Bool(true),
		}))

	return chart
}

func DrawBERChartExercise2(inputWord string, coderChoose string) *charts.Line {
	betaValuesASK, berValuesASK := GenerateBERDataExercise2(inputWord, coderChoose, "ASK")
	_, berValuesPSK := GenerateBERDataExercise2(inputWord, coderChoose, "PSK")
	_, berValuesFSK := GenerateBERDataExercise2(inputWord, coderChoose, "FSK")

	chart := charts.NewLine()

	utils.SetBERChartOptions(chart, "Lab12 - Ćwiczenie 2", coderChoose+" - ASK/PSK/FSK - Słowo informacyjne: "+inputWord, "Beta")

	chart.SetXAxis(utils.AlphaAxisLabels(betaValuesASK)).
		AddSeries("ASK", utils.FromSliceToLineData(berValuesASK)).
		AddSeries("PSK", utils.FromSliceToLineData(berValuesPSK)).
		AddSeries("FSK", utils.FromSliceToLineData(berValuesFSK)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{
			Smooth:     opts.Bool(false),
			ShowSymbol: opts.Bool(true),
		}))

	return chart
}

func DrawExercise1Hamming_7_4(w http.ResponseWriter, r *http.Request) {
	inputWord := r.URL.Query().Get("word")

	chartASK := DrawBERChartExercise1(inputWord, "Hamming_7_4", "ASK")
	chartPSK := DrawBERChartExercise1(inputWord, "Hamming_7_4", "PSK")
	chartFSK := DrawBERChartExercise1(inputWord, "Hamming_7_4", "FSK")

	page := components.NewPage()
	page.AddCharts(chartASK, chartPSK, chartFSK)
	page.Render(w)
}

func DrawExercise1Hamming_15_11(w http.ResponseWriter, r *http.Request) {
	inputWord := r.URL.Query().Get("word")

	chartASK := DrawBERChartExercise1(inputWord, "Hamming_15_11", "ASK")
	chartPSK := DrawBERChartExercise1(inputWord, "Hamming_15_11", "PSK")
	chartFSK := DrawBERChartExercise1(inputWord, "Hamming_15_11", "FSK")

	page := components.NewPage()
	page.AddCharts(chartASK, chartPSK, chartFSK)
	page.Render(w)
}

func DrawExercise2Hamming_7_4(w http.ResponseWriter, r *http.Request) {
	inputWord := r.URL.Query().Get("word")

	chart := DrawBERChartExercise2(inputWord, "Hamming_7_4")

	page := components.NewPage()
	page.AddCharts(chart)
	page.Render(w)
}

func DrawExercise2Hamming_15_11(w http.ResponseWriter, r *http.Request) {
	inputWord := r.URL.Query().Get("word")

	chart := DrawBERChartExercise2(inputWord, "Hamming_15_11")

	page := components.NewPage()
	page.AddCharts(chart)
	page.Render(w)
}
