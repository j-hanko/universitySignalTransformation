package lab1

import (
	"fmt"
	"math"

	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

//Numery wybranych funkcji 13,3,13,13

func signalGenerationExerise1(Tc, f, fs, phi float64, formula string) []opts.LineData {
	N := int(math.Round(Tc * fs))
	sliceOfData := make([]opts.LineData, 0)
	for n := 0; n <= N; n++ {
		t := float64(n) / fs
		x := 0.9*math.Sin(2*math.Pi*f*t+phi)*math.Cos(21*math.Pi*t) + (t - 0.66*t)
		y := (math.Pow(t, 3) - 1) + math.Cos(4*math.Pow(t, 2)*math.Pi)*t
		z := x / (math.Abs(y*math.Cos(5*t)-x*y) + 3)
		v := (x * 662) / (math.Abs(x-y) + 0.5)

		if formula == "x" {
			sliceOfData = append(sliceOfData, opts.LineData{Value: x})
		} else if formula == "y" {
			sliceOfData = append(sliceOfData, opts.LineData{Value: y})
		} else if formula == "z" {
			sliceOfData = append(sliceOfData, opts.LineData{Value: z})
		} else if formula == "v" {
			sliceOfData = append(sliceOfData, opts.LineData{Value: v})
		} else {
			fmt.Println("Wrong formula argument")
			return nil
		}

	}
	return sliceOfData
}

func signalGenerationExerise2(Tc, fs, h float64, formula string) []opts.LineData {
	N := int(math.Round(Tc * fs))
	sliceOfData := make([]opts.LineData, 0)
	for n := 0; n <= N; n++ {
		t := float64(n) / fs
		if formula == "u" {
			if t >= 0 && t < 0.3 {
				u := 0.1*math.Cos(36*math.Pi*t) + math.Sin(22*math.Pi*t)
				sliceOfData = append(sliceOfData, opts.LineData{Value: u})
			} else if t >= 0.3 && t < 0.8 {
				u := (t - 0.3) * math.Cos(26*math.Pi*t+math.Sin(12*t))
				sliceOfData = append(sliceOfData, opts.LineData{Value: u})
			} else if t >= 0.8 && t < 1 {
				u := 0.1 * (math.Log2(t+2)*math.Sin(6*math.Pi*t) + math.Log2(math.Cos(44*math.Pi*t)+2))
				sliceOfData = append(sliceOfData, opts.LineData{Value: u})
			}
		} else if formula == "b" {
			b1 := float64(0)
			for i := 0; i < int(h); i++ {
				calc := math.Sin((6*math.Pi+1)*t*math.Pi) * math.Sin(math.Pow(h, 3)) / (6*h + 2)
				b1 = b1 + calc
			}
			sliceOfData = append(sliceOfData, opts.LineData{Value: b1})

		} else {
			fmt.Println("Wrong formula argument")
			return nil
		}

	}
	return sliceOfData
}

func drawExercise1(w http.ResponseWriter, _ *http.Request) {
	chart1 := charts.NewLine()
	setChartOptions(chart1, "Laboratorium 1", "Zadanie 1", "Czas [s]")
	chart1.SetXAxis(timeAxisLabels(Tc, fs, step)).
		AddSeries("x(t)", fromSliceToLineData(x)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))

	page := components.NewPage()
	page.AddCharts(chart1)
	page.Render(w)
}

func drawZad1(w http.ResponseWriter, _ *http.Request) {
	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWonderland}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Transmisja danych - instrukcja nr 1",
			Subtitle: "Zadanie 1",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name:         "Czas [s]",
			NameLocation: "middle",
			NameGap:      30,
			AxisLabel: &opts.AxisLabel{
				Show:     opts.Bool(true),
				Interval: "0",
			},
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name:         "Amplituda",
			NameLocation: "middle",
			NameGap:      50,
			Position:     "left",
			AxisLabel: &opts.AxisLabel{
				Show: opts.Bool(true),
			},
		}))

	// Put data into instance
	line.SetXAxis(timeAxisLabels(1.7, 8000, 800)).
		AddSeries("X", lab1Zad1(1.7, 50, 8000, math.Pi/2, "x")).
		AddSeries("Y", lab1Zad1(1.7, 50, 8000, math.Pi/2, "y")).
		AddSeries("Z", lab1Zad1(1.7, 50, 8000, math.Pi/2, "z")).
		AddSeries("V", lab1Zad1(1.7, 50, 8000, math.Pi/2, "v")).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))
	line.Render(w)

}

func drawZad2(w http.ResponseWriter, _ *http.Request) {
	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWonderland}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Transmisja danych - instrukcja nr 1",
			Subtitle: "Zadanie 2",
		}),
		charts.WithGridOpts(opts.Grid{
			Bottom: "60",
			Right:  "80",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name:         "Czas [s]",
			NameLocation: "middle",
			NameGap:      30,
			AxisLabel: &opts.AxisLabel{
				Show:     opts.Bool(true),
				Interval: "0",
			},
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name:         "Amplituda",
			NameLocation: "middle",
			NameGap:      50,
			Position:     "left",
			AxisLabel: &opts.AxisLabel{
				Show: opts.Bool(true),
			},
		}))

	// Put data into instance
	line.SetXAxis(timeAxisLabels(1.7, 8000, 800)).
		AddSeries("U", lab1Zad2(1.7, 8000, 0, "u")).
		AddSeries("b1", lab1Zad2(1, 22050, 2, "b")).
		AddSeries("b2", lab1Zad2(1, 22050, 4, "b")).
		AddSeries("b3", lab1Zad2(1, 22050, 8, "b")).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(false)}))
	line.Render(w)

}

func main() {
	http.HandleFunc("/zad1", drawZad1)
	http.HandleFunc("/zad2", drawZad2)
	http.ListenAndServe(":8081", nil)
	fmt.Println("Server started on port 8081")
}
