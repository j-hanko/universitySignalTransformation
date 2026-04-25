package lab5

import (
	"fmt"
	"math"
	"net/http"
	"slices"
	"universitySignalTransformation/pkg/lab4"
	"universitySignalTransformation/pkg/utils"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

const (
	Tc   = 1
	fs   = 2000
	step = 10
	fn   = 50
	fm   = 10
)

func dB_Spectrum(M []float64) []float64 {
	sliceOfData := make([]float64, 0, len(M))
	for i := 0; i < len(M); i++ {
		M_db := 20 * math.Log10(M[i])
		sliceOfData = append(sliceOfData, M_db)
	}
	return sliceOfData
}

func bandwidth(x []float64, fs int, dB float64) float64 {
	re, im := utils.FFT(x)
	M := utils.Spectrum(re, im)
	MdB := dB_Spectrum(M)

	maxdB := slices.Max(MdB)
	threshold := maxdB - dB

	N := len(x)
	half := len(MdB) / 2

	var fMin float64
	var fMax float64

	for i := 0; i < half; i++ {
		if MdB[i] >= threshold {
			fMin = float64(i) * float64(fs) / float64(N)
			break
		}
	}

	for i := half - 1; i >= 0; i-- {
		if MdB[i] >= threshold {
			fMax = float64(i) * float64(fs) / float64(N)
			break
		}
	}

	B := fMax - fMin

	return B
}

func countBandwidthAndWriteToFile(bandwidthTitleName string, data float64, fileName string) {
	utils.WriteToFile(fmt.Sprintf("%s: %.2f", bandwidthTitleName, data), fileName)
}

func SaveAllExercise1Data(fileName string, db float64) {
	Za_a := lab4.SignalGenerationExerise1(Tc, fs, fn, fm, 0.5, "Z_A")
	Za_b := lab4.SignalGenerationExerise1(Tc, fs, fn, fm, 7.5, "Z_A")
	Za_c := lab4.SignalGenerationExerise1(Tc, fs, fn, fm, 25.5, "Z_A")

	Zf_a := lab4.SignalGenerationExerise1(Tc, fs, fn, fm, 0.5, "Z_F")
	Zf_b := lab4.SignalGenerationExerise1(Tc, fs, fn, fm, 2.3, "Z_F")
	Zf_c := lab4.SignalGenerationExerise1(Tc, fs, fn, fm, 11.3, "Z_F")

	Zp_a := lab4.SignalGenerationExerise1(Tc, fs, fn, fm, 0.5, "Z_P")
	Zp_b := lab4.SignalGenerationExerise1(Tc, fs, fn, fm, 2.7, "Z_P")
	Zp_c := lab4.SignalGenerationExerise1(Tc, fs, fn, fm, 9.7, "Z_P")

	countBandwidthAndWriteToFile("Za_a", bandwidth(Za_a, fs, db), fileName)
	countBandwidthAndWriteToFile("Za_b", bandwidth(Za_b, fs, db), fileName)
	countBandwidthAndWriteToFile("Za_c", bandwidth(Za_c, fs, db), fileName)

	countBandwidthAndWriteToFile("Zf_a", bandwidth(Zf_a, fs, db), fileName)
	countBandwidthAndWriteToFile("Zf_b", bandwidth(Zf_b, fs, db), fileName)
	countBandwidthAndWriteToFile("Zf_c", bandwidth(Zf_c, fs, db), fileName)

	countBandwidthAndWriteToFile("Zp_a", bandwidth(Zp_a, fs, db), fileName)
	countBandwidthAndWriteToFile("Zp_b", bandwidth(Zp_b, fs, db), fileName)
	countBandwidthAndWriteToFile("Zp_c", bandwidth(Zp_c, fs, db), fileName)
}

func DrawExercise_Ma(w http.ResponseWriter, _ *http.Request) {
	Za_a := lab4.SignalGenerationExerise1(Tc, fs, fn, fm, 0.5, "Z_A")
	Za_b := lab4.SignalGenerationExerise1(Tc, fs, fn, fm, 7.5, "Z_A")
	Za_c := lab4.SignalGenerationExerise1(Tc, fs, fn, fm, 25.5, "Z_A")

	aRe, aIm := utils.FFT(Za_a)
	bRe, bIm := utils.FFT(Za_b)
	cRe, cIm := utils.FFT(Za_c)

	Ma_a := dB_Spectrum(utils.Spectrum(aRe, aIm))
	Ma_b := dB_Spectrum(utils.Spectrum(bRe, bIm))
	Ma_c := dB_Spectrum(utils.Spectrum(cRe, cIm))

	chart1 := charts.NewLine()
	utils.SetSpectrumChartOptions(chart1, "Laboratorium 5", "Zadanie 1 widmo amplitudowe dla Z_a z wartością k = 0.5")
	chart1.SetXAxis(utils.FrequencyAxisLabels(Tc, fs, step)).
		AddSeries("Ma_a", utils.FromSliceToLineData(Ma_a)).
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}),
			charts.WithLineStyleOpts(opts.LineStyle{Width: 1}))

	chart2 := charts.NewLine()
	utils.SetSpectrumChartOptions(chart2, "Laboratorium 5", "Zadanie 1 widmo amplitudowe dla Z_a z wartością k = 7.5")
	chart2.SetXAxis(utils.FrequencyAxisLabels(Tc, fs, step)).
		AddSeries("Ma_b", utils.FromSliceToLineData(Ma_b)).
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}),
			charts.WithLineStyleOpts(opts.LineStyle{Width: 1}))

	chart3 := charts.NewLine()
	utils.SetSpectrumChartOptions(chart3, "Laboratorium 5", "Zadanie 1 widmo amplitudowe dla Z_a z wartością k = 25.5")
	chart3.SetXAxis(utils.FrequencyAxisLabels(Tc, fs, step)).
		AddSeries("Ma_c", utils.FromSliceToLineData(Ma_c)).
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}),
			charts.WithLineStyleOpts(opts.LineStyle{Width: 1}))

	page := components.NewPage()
	page.AddCharts(chart1, chart2, chart3)
	page.Render(w)
}

func DrawExercise_Mf(w http.ResponseWriter, _ *http.Request) {
	Zf_a := lab4.SignalGenerationExerise1(Tc, fs, fn, fm, 0.5, "Z_F")
	Zf_b := lab4.SignalGenerationExerise1(Tc, fs, fn, fm, 2.3, "Z_F")
	Zf_c := lab4.SignalGenerationExerise1(Tc, fs, fn, fm, 11.3, "Z_F")

	aRe, aIm := utils.FFT(Zf_a)
	bRe, bIm := utils.FFT(Zf_b)
	cRe, cIm := utils.FFT(Zf_c)

	Mf_a := dB_Spectrum(utils.Spectrum(aRe, aIm))
	Mf_b := dB_Spectrum(utils.Spectrum(bRe, bIm))
	Mf_c := dB_Spectrum(utils.Spectrum(cRe, cIm))

	chart1 := charts.NewLine()
	utils.SetSpectrumChartOptions(chart1, "Laboratorium 5", "Zadanie 1 widmo amplitudowe dla Z_f z wartością k = 0.5")
	chart1.SetXAxis(utils.FrequencyAxisLabels(Tc, fs, step)).
		AddSeries("Mf_a", utils.FromSliceToLineData(Mf_a)).
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}),
			charts.WithLineStyleOpts(opts.LineStyle{Width: 1}))

	chart2 := charts.NewLine()
	utils.SetSpectrumChartOptions(chart2, "Laboratorium 5", "Zadanie 1 widmo amplitudowe dla Z_f z wartością k = 2.3")
	chart2.SetXAxis(utils.FrequencyAxisLabels(Tc, fs, step)).
		AddSeries("Mf_b", utils.FromSliceToLineData(Mf_b)).
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}),
			charts.WithLineStyleOpts(opts.LineStyle{Width: 1}))

	chart3 := charts.NewLine()
	utils.SetSpectrumChartOptions(chart3, "Laboratorium 5", "Zadanie 1 widmo amplitudowe dla Z_f z wartością k = 11.3")
	chart3.SetXAxis(utils.FrequencyAxisLabels(Tc, fs, step)).
		AddSeries("Mf_c", utils.FromSliceToLineData(Mf_c)).
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}),
			charts.WithLineStyleOpts(opts.LineStyle{Width: 1}))

	page := components.NewPage()
	page.AddCharts(chart1, chart2, chart3)
	page.Render(w)
}

func DrawExercise_Mp(w http.ResponseWriter, _ *http.Request) {
	Zp_a := lab4.SignalGenerationExerise1(Tc, fs, fn, fm, 0.5, "Z_P")
	Zp_b := lab4.SignalGenerationExerise1(Tc, fs, fn, fm, 2.7, "Z_P")
	Zp_c := lab4.SignalGenerationExerise1(Tc, fs, fn, fm, 9.7, "Z_P")

	aRe, aIm := utils.DFT(Zp_a)
	bRe, bIm := utils.DFT(Zp_b)
	cRe, cIm := utils.DFT(Zp_c)

	Mp_a := dB_Spectrum(utils.Spectrum(aRe, aIm))
	Mp_b := dB_Spectrum(utils.Spectrum(bRe, bIm))
	Mp_c := dB_Spectrum(utils.Spectrum(cRe, cIm))

	chart1 := charts.NewLine()
	utils.SetSpectrumChartOptions(chart1, "Laboratorium 5", "Zadanie 1 widmo amplitudowe dla Z_p z wartością k = 0.5")
	chart1.SetXAxis(utils.FrequencyAxisLabels(Tc, fs, step)).
		AddSeries("Mp_a", utils.FromSliceToLineData(Mp_a)).
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}),
			charts.WithLineStyleOpts(opts.LineStyle{Width: 1}))

	chart2 := charts.NewLine()
	utils.SetSpectrumChartOptions(chart2, "Laboratorium 5", "Zadanie 1 widmo amplitudowe dla Z_p z wartością k = 2.7")
	chart2.SetXAxis(utils.FrequencyAxisLabels(Tc, fs, step)).
		AddSeries("Mp_b", utils.FromSliceToLineData(Mp_b)).
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}),
			charts.WithLineStyleOpts(opts.LineStyle{Width: 1}))

	chart3 := charts.NewLine()
	utils.SetSpectrumChartOptions(chart3, "Laboratorium 5", "Zadanie 1 widmo amplitudowe dla Z_p z wartością k = 9.7")
	chart3.SetXAxis(utils.FrequencyAxisLabels(Tc, fs, step)).
		AddSeries("Mp_c", utils.FromSliceToLineData(Mp_c)).
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}),
			charts.WithLineStyleOpts(opts.LineStyle{Width: 1}))

	page := components.NewPage()
	page.AddCharts(chart1, chart2, chart3)
	page.Render(w)
}
