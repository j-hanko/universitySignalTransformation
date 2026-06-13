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

func transmisionChannel(input []float64, alpha float64) (output []float64) {
	output = make([]float64, len(input))

	for i := 0; i < len(input); i++ {
		noise := rand.Float64()*2 - 1
		output[i] = input[i] + alpha*noise
	}

	return output
}

func SimpleTransmissionSystem(inputWord string, coderChoose string, signalModulator string, alpha float64) []byte {
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
	yt := transmisionChannel(ModulatorOutput, alpha)

	//Demodulation
	ytInt := utils.FromFloatSliceToIntSlice(yt)
	var DemodulatorOutput []int = make([]int, len(ytInt))

	if signalModulator == "ASK" || signalModulator == "PSK" {
		_, _, _, _, DemodulatorOutput = lab8.Demodulator(ytInt, signalModulator)
	} else if signalModulator == "FSK" {
		_, _, _, _, _, _, _, DemodulatorOutput = lab9.Demodulator(ytInt)
	}

	//Decoding
	HammingDecodeInput := utils.FromIntSliceToBitSlice(DemodulatorOutput)
	var HammingDecodeOutput []byte = make([]byte, 0)
	if coderChoose == "Hamming_7_4" {
		HammingDecodeOutput = lab10.HammingDecode(HammingDecodeInput)
	} else if coderChoose == "Hamming_15_11" {
		HammingDecodeOutput = lab11.HammingDecode(HammingDecodeInput)
	}

	return HammingDecodeOutput
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

	inputBits := utils.InputWordToBits(inputWord, coderChoose)

	for i := 0; i < 10; i++ {
		alpha := float64(i) * 3.0 / 9.0

		outputBits := SimpleTransmissionSystem(inputWord, coderChoose, signalModulator, alpha)
		ber := BER(inputBits, outputBits)

		alphaValues = append(alphaValues, alpha)
		berValues = append(berValues, ber)
	}

	return alphaValues, berValues
}

func DrawBERChartExercise1(inputWord string, coderChoose string, signalModulator string) *charts.Line {
	alphaValues, berValues := GenerateBERDataExercise1(inputWord, coderChoose, signalModulator)

	chart := charts.NewLine()

	utils.SetChartOptions(chart, "Lab12 - Ćwiczenie 1", coderChoose+" - "+signalModulator, "Alpha")

	chart.SetXAxis(utils.AlphaAxisLabels(alphaValues)).
		AddSeries("BER", utils.FromSliceToLineData(berValues)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{
			Smooth:     opts.Bool(false),
			ShowSymbol: opts.Bool(true),
		}))

	return chart
}

func DrawExercise1Hamming_7_4(w http.ResponseWriter, r *http.Request) {
	inputWord := r.URL.Query().Get("word")

	chartFSK := DrawBERChartExercise1(inputWord, "Hamming_7_4", "FSK")
	chartASK := DrawBERChartExercise1(inputWord, "Hamming_7_4", "ASK")
	chartPSK := DrawBERChartExercise1(inputWord, "Hamming_7_4", "PSK")

	page := components.NewPage()
	page.AddCharts(chartFSK, chartASK, chartPSK)
	page.Render(w)
}

func DrawExercise1Hamming_15_11(w http.ResponseWriter, r *http.Request) {
	inputWord := r.URL.Query().Get("word")

	chartFSK := DrawBERChartExercise1(inputWord, "Hamming_15_11", "FSK")
	chartASK := DrawBERChartExercise1(inputWord, "Hamming_15_11", "ASK")
	chartPSK := DrawBERChartExercise1(inputWord, "Hamming_15_11", "PSK")

	page := components.NewPage()
	page.AddCharts(chartFSK, chartASK, chartPSK)
	page.Render(w)
}
