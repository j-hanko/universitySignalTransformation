package utils

import "gonum.org/v1/gonum/dsp/fourier"

func FFT(x []float64) ([]float64, []float64) {
	fft := fourier.NewFFT(len(x))
	coeff := fft.Coefficients(nil, x)

	xRe := make([]float64, len(coeff))
	xIm := make([]float64, len(coeff))

	for i := 0; i < len(coeff); i++ {
		xRe[i] = real(coeff[i])
		xIm[i] = imag(coeff[i])
	}

	return xRe, xIm
}
