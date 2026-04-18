package utils

import "math"

func Spectrum(xRe, xIm []float64) []float64 {
	N := len(xRe)
	M := make([]float64, 0, N/2)

	for k := 0; k < N/2; k++ {
		value := math.Sqrt(xRe[k]*xRe[k]+xIm[k]*xIm[k]) / float64(N)
		M = append(M, value)
	}

	return M
}

func MzHatSpectrum(alpha, beta float64, Mx, My []float64) []float64 {
	sliceOfData := make([]float64, 0, len(Mx))

	for i := 0; i < len(Mx); i++ {
		value := alpha*Mx[i] + beta*My[i]
		sliceOfData = append(sliceOfData, value)
	}
	return sliceOfData
}
