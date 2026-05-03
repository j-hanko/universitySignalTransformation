package lab6

import (
	"math"
	"net/http"
)

const (
	A1 = 0.5
	A2 = 1.0
	W  = 2
	Tc = 1
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
				value = math.Sin(math.Pi * fn1 * t)
			} else if bit == 1 {
				value = math.Sin(math.Pi * fn2 * t)
			}
		} else if formula == "PSK" {
			if bit == 0 {
				value = math.Sin(math.Pi * fn * t)
			} else if bit == 1 {
				value = math.Sin(math.Pi*fn*t + math.Pi)
			}
		}
		sliceOfData = append(sliceOfData, value)
	}
	return sliceOfData
}

func DrawExercise(w http.ResponseWriter, _ *http.Request) {

}
