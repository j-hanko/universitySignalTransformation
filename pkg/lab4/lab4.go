package lab4

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
	Tc   = 1
	fs   = 2000
	step = 200
	fn   = 50
	fm   = 10
)

func signalGenerationExerise1(Tc, fs, fn, fm, useCase float64, formula string) []float64 {
	N := int(math.Round(Tc * fs))
	sliceOfData := make([]float64, 0)
	if formula == "M" {
		for n := 0; n <= N; n++ {
			t := float64(n) / fs
			m := math.Sin(float64(2) * math.Pi * fm * t)
			sliceOfData = append(sliceOfData, m)
		}
	} else if formula == "Z_A" {
		for n := 0; n <= N; n++ {
			t := float64(n) / fs
			z_a := useCase * math.Sin(float64(2)*math.Pi*fm*t) * math.Cos(float64(2)*math.Pi*fn*t)
			sliceOfData = append(sliceOfData, z_a)
		}
	} else if formula == "Z_P" {
		for n := 0; n <= N; n++ {
			t := float64(n) / fs
			z_p := math.Cos(float64(2)*math.Pi*fn*t + useCase*math.Sin(float64(2)*math.Pi*fm*t))
			sliceOfData = append(sliceOfData, z_p)
		}
	} else if formula == "Z_F" {
		for n := 0; n <= N; n++ {
			t := float64(n) / fs
			z_f := math.Cos(float64(2)*math.Pi*fn*t + useCase/fm*math.Sin(float64(2)*math.Pi*fm*t))
			sliceOfData = append(sliceOfData, z_f)
		}
	} else {
		fmt.Println("Wrong formula argument")
		return nil
	}
	return sliceOfData
}

func DrawExercise_Za(w http.ResponseWriter, _ *http.Request) {
	Za_a := signalGenerationExerise1(Tc, fs, fn, fm, 0.5, "Z_A")
	Za_b := signalGenerationExerise1(Tc, fs, fn, fm, 7.5, "Z_A")
	Za_c := signalGenerationExerise1(Tc, fs, fn, fm, 25.5, "Z_A")

	Re_a, Im_a := utils.DFT(Za_a)
	Re_b, Im_b := utils.DFT(Za_b)
	Re_c, Im_c := utils.DFT(Za_c)

	Ma := utils.Spectrum(Re_a, Im_a)
	Mb := utils.Spectrum(Re_b, Im_b)
	Mc := utils.Spectrum(Re_c, Im_c)

	chart1 := charts.NewLine()
	utils.SetChartOptions(chart1, "Laboratorium 4", "Zadanie 1 przypadek a dla warotości k = 0.5", "Czas [s]")
	chart1.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("Z_a", utils.FromSliceToLineData(Za_a)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart2 := charts.NewLine()
	utils.SetChartOptions(chart2, "Laboratorium 4", "Zadanie 1 przypadek b dla warotości k = 7.5", "Czas [s]")
	chart2.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("Z_a", utils.FromSliceToLineData(Za_b)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart3 := charts.NewLine()
	utils.SetChartOptions(chart3, "Laboratorium 4", "Zadanie 1 przypadek c dla warotości k = 25.5", "Czas [s]")
	chart3.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("Z_a", utils.FromSliceToLineData(Za_c)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart4 := charts.NewLine()
	utils.SetChartOptions(chart4, "Laboratorium 4", "Zadanie 2 spektrum dla przypadku a", "Częstotliwość [Hz]")
	chart4.SetXAxis(utils.FrequencyAxisLabels(Tc, fs, step)).
		AddSeries("M_a", utils.FromSliceToLineData(Ma)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart5 := charts.NewLine()
	utils.SetChartOptions(chart5, "Laboratorium 4", "Zadanie 2 spektrum dla przypadku b", "Częstotliwość [Hz]")
	chart5.SetXAxis(utils.FrequencyAxisLabels(Tc, fs, step)).
		AddSeries("M_a", utils.FromSliceToLineData(Mb)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart6 := charts.NewLine()
	utils.SetChartOptions(chart6, "Laboratorium 4", "Zadanie 2 spektrum dla przypadku c", "Częstotliwość [Hz]")
	chart6.SetXAxis(utils.FrequencyAxisLabels(Tc, fs, step)).
		AddSeries("M_a", utils.FromSliceToLineData(Mc)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	page := components.NewPage()
	page.AddCharts(chart1, chart2, chart3, chart4, chart5, chart6)
	page.Render(w)

}

func DrawExercise_Zf(w http.ResponseWriter, _ *http.Request) {
	Zf_a := signalGenerationExerise1(Tc, fs, fn, fm, 0.5, "Z_F")
	Zf_b := signalGenerationExerise1(Tc, fs, fn, fm, 2.3, "Z_F")
	Zf_c := signalGenerationExerise1(Tc, fs, fn, fm, 11.3, "Z_F")

	Re_a, Im_a := utils.DFT(Zf_a)
	Re_b, Im_b := utils.DFT(Zf_b)
	Re_c, Im_c := utils.DFT(Zf_c)

	Ma := utils.Spectrum(Re_a, Im_a)
	Mb := utils.Spectrum(Re_b, Im_b)
	Mc := utils.Spectrum(Re_c, Im_c)

	chart1 := charts.NewLine()
	utils.SetChartOptions(chart1, "Laboratorium 4", "Zadanie 1 przypadek a dla warotości k = 0.5", "Czas [s]")
	chart1.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("Z_f", utils.FromSliceToLineData(Zf_a)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart2 := charts.NewLine()
	utils.SetChartOptions(chart2, "Laboratorium 4", "Zadanie 1 przypadek b dla warotości k = 2.3", "Czas [s]")
	chart2.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("Z_f", utils.FromSliceToLineData(Zf_b)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart3 := charts.NewLine()
	utils.SetChartOptions(chart3, "Laboratorium 4", "Zadanie 1 przypadek b dla warotości k = 11.3", "Czas [s]")
	chart3.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("Z_f", utils.FromSliceToLineData(Zf_c)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart4 := charts.NewLine()
	utils.SetChartOptions(chart4, "Laboratorium 4", "Zadanie 2 spektrum dla przypadku a", "Częstotliwość [Hz]")
	chart4.SetXAxis(utils.FrequencyAxisLabels(Tc, fs, step)).
		AddSeries("M_f", utils.FromSliceToLineData(Ma)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart5 := charts.NewLine()
	utils.SetChartOptions(chart5, "Laboratorium 4", "Zadanie 2 spektrum dla przypadku b", "Częstotliwość [Hz]")
	chart5.SetXAxis(utils.FrequencyAxisLabels(Tc, fs, step)).
		AddSeries("M_f", utils.FromSliceToLineData(Mb)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart6 := charts.NewLine()
	utils.SetChartOptions(chart6, "Laboratorium 4", "Zadanie 2 spektrum dla przypadku c", "Częstotliwość [Hz]")
	chart6.SetXAxis(utils.FrequencyAxisLabels(Tc, fs, step)).
		AddSeries("M_f", utils.FromSliceToLineData(Mc)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	page := components.NewPage()
	page.AddCharts(chart1, chart2, chart3, chart4, chart5, chart6)
	page.Render(w)
}

func DrawExercise_Zp(w http.ResponseWriter, _ *http.Request) {
	Zp_a := signalGenerationExerise1(Tc, fs, fn, fm, 0.5, "Z_P")
	Zp_b := signalGenerationExerise1(Tc, fs, fn, fm, 2.7, "Z_P")
	Zp_c := signalGenerationExerise1(Tc, fs, fn, fm, 9.7, "Z_P")

	Re_a, Im_a := utils.DFT(Zp_a)
	Re_b, Im_b := utils.DFT(Zp_b)
	Re_c, Im_c := utils.DFT(Zp_c)

	Ma := utils.Spectrum(Re_a, Im_a)
	Mb := utils.Spectrum(Re_b, Im_b)
	Mc := utils.Spectrum(Re_c, Im_c)

	chart1 := charts.NewLine()
	utils.SetChartOptions(chart1, "Laboratorium 4", "Zadanie 1 przypadek a dla warotości k = 0.7", "Czas [s]")
	chart1.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("Z_p", utils.FromSliceToLineData(Zp_a)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart2 := charts.NewLine()
	utils.SetChartOptions(chart2, "Laboratorium 4", "Zadanie 1 przypadek b dla warotości k = 2.7", "Czas [s]")
	chart2.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("Z_p", utils.FromSliceToLineData(Zp_b)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart3 := charts.NewLine()
	utils.SetChartOptions(chart3, "Laboratorium 4", "Zadanie 1 przypadek b dla warotości k = 9.7", "Czas [s]")
	chart3.SetXAxis(utils.TimeAxisLabels(Tc, fs, step)).
		AddSeries("Z_p", utils.FromSliceToLineData(Zp_c)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart4 := charts.NewLine()
	utils.SetChartOptions(chart4, "Laboratorium 4", "Zadanie 2 spektrum dla przypadku a", "Częstotliwość [Hz]")
	chart4.SetXAxis(utils.FrequencyAxisLabels(Tc, fs, step)).
		AddSeries("M_p", utils.FromSliceToLineData(Ma)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart5 := charts.NewLine()
	utils.SetChartOptions(chart5, "Laboratorium 4", "Zadanie 2 spektrum dla przypadku b", "Częstotliwość [Hz]")
	chart5.SetXAxis(utils.FrequencyAxisLabels(Tc, fs, step)).
		AddSeries("M_p", utils.FromSliceToLineData(Mb)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	chart6 := charts.NewLine()
	utils.SetChartOptions(chart6, "Laboratorium 4", "Zadanie 2 spektrum dla przypadku c", "Częstotliwość [Hz]")
	chart6.SetXAxis(utils.FrequencyAxisLabels(Tc, fs, step)).
		AddSeries("M_p", utils.FromSliceToLineData(Mc)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	page := components.NewPage()
	page.AddCharts(chart1, chart2, chart3, chart4, chart5, chart6)
	page.Render(w)
}
