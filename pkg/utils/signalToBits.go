package utils

func SignalToBits(c []float64, samplesPerBit int) []int {
	bits := make([]int, 0)

	for i := samplesPerBit - 1; i < len(c); i += samplesPerBit {
		if c[i] >= 0.5 {
			bits = append(bits, 1)
		} else {
			bits = append(bits, 0)
		}
	}

	return bits
}
