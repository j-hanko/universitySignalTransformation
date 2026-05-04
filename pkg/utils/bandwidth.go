package utils

import (
	"fmt"
	"math"
)

func Bandwidth(x []float64, fs int, dB float64) float64 {
	re, im := FFT(x)
	M := Spectrum(re, im)

	N := len(x)
	df := float64(fs) / float64(N)

	max := M[1]
	for i := 2; i < len(M); i++ {
		if M[i] > max {
			max = M[i]
		}
	}

	level := max * math.Pow(10, -dB/20.0)

	first := -1
	last := -1

	for i := 1; i < len(M); i++ {
		if M[i] >= level {
			if first == -1 {
				first = i
			}
			last = i
		}
	}

	if first == -1 {
		return 0
	}

	B := float64(last-first) * df
	return B
}

func CountBandwidthAndWriteToFile(bandwidthTitleName string, data float64, fileName string) {
	WriteToFile(fmt.Sprintf("%s: %.2f", bandwidthTitleName, data), fileName)
}
