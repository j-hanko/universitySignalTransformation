package utils

import "math"

func DFT(x []float64) (x_re []float64, x_im []float64) {
	N := len(x)
	x_re = make([]float64, N)
	x_im = make([]float64, N)

	for k := 0; k < N; k++ {
		a_sum := 0.0
		b_sum := 0.0

		for n := 0; n < N; n++ {
			alpha := (-2 * math.Pi * float64(n*k)) / float64(N)
			a_sum += x[n] * math.Cos(alpha)
			b_sum += x[n] * math.Sin(alpha)
		}

		x_re[k] = a_sum
		x_im[k] = b_sum
	}

	return x_re, x_im
}
