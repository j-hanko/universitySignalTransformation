package lab5

import (
	"net/http"
	"universitySignalTransformation/pkg/lab4"
	"universitySignalTransformation/pkg/utils"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

const (
	Tc   = 2
	fs   = 2000
	step = 10
	fn   = 50
	fm   = 10
)

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

	utils.CountBandwidthAndWriteToFile("Za_a", utils.Bandwidth(Za_a, fs, db), fileName)
	utils.CountBandwidthAndWriteToFile("Za_b", utils.Bandwidth(Za_b, fs, db), fileName)
	utils.CountBandwidthAndWriteToFile("Za_c", utils.Bandwidth(Za_c, fs, db), fileName)

	utils.CountBandwidthAndWriteToFile("Zf_a", utils.Bandwidth(Zf_a, fs, db), fileName)
	utils.CountBandwidthAndWriteToFile("Zf_b", utils.Bandwidth(Zf_b, fs, db), fileName)
	utils.CountBandwidthAndWriteToFile("Zf_c", utils.Bandwidth(Zf_c, fs, db), fileName)

	utils.CountBandwidthAndWriteToFile("Zp_a", utils.Bandwidth(Zp_a, fs, db), fileName)
	utils.CountBandwidthAndWriteToFile("Zp_b", utils.Bandwidth(Zp_b, fs, db), fileName)
	utils.CountBandwidthAndWriteToFile("Zp_c", utils.Bandwidth(Zp_c, fs, db), fileName)
}

func DrawExercise_Ma(w http.ResponseWriter, _ *http.Request) {
	Za_a := lab4.SignalGenerationExerise1(Tc, fs, fn, fm, 0.5, "Z_A")
	Za_b := lab4.SignalGenerationExerise1(Tc, fs, fn, fm, 7.5, "Z_A")
	Za_c := lab4.SignalGenerationExerise1(Tc, fs, fn, fm, 25.5, "Z_A")

	aRe, aIm := utils.FFT(Za_a)
	bRe, bIm := utils.FFT(Za_b)
	cRe, cIm := utils.FFT(Za_c)

	Ma_a := utils.DB_Spectrum(utils.Spectrum(aRe, aIm))
	Ma_b := utils.DB_Spectrum(utils.Spectrum(bRe, bIm))
	Ma_c := utils.DB_Spectrum(utils.Spectrum(cRe, cIm))

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

	Mf_a := utils.DB_Spectrum(utils.Spectrum(aRe, aIm))
	Mf_b := utils.DB_Spectrum(utils.Spectrum(bRe, bIm))
	Mf_c := utils.DB_Spectrum(utils.Spectrum(cRe, cIm))

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

	Mp_a := utils.DB_Spectrum(utils.Spectrum(aRe, aIm))
	Mp_b := utils.DB_Spectrum(utils.Spectrum(bRe, bIm))
	Mp_c := utils.DB_Spectrum(utils.Spectrum(cRe, cIm))

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
