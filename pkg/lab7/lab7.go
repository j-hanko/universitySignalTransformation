package lab7

import (
	"math"
	"universitySignalTransformation/pkg/lab6"
	"universitySignalTransformation/pkg/utils"
)

const (
	A  = 1
	W  = 2
	Tc = 1
)

func SaveAllExercise1Data(fileName string, db float64) {
	bits := utils.ASCII_to_bit("KOSIARKA")

	B := len(bits)
	Tb := Tc / float64(B)
	fn := W / Tb
	fs := int(math.Round(1000 * fn))

	Z_a := lab6.SignalGenerationExercise(bits, "ASK")
	Z_p := lab6.SignalGenerationExercise(bits, "PSK")
	Z_f := lab6.SignalGenerationExercise(bits, "FSK")

	utils.CountBandwidthAndWriteToFile("Z_a", utils.Bandwidth(Z_a, fs, db), fileName)
	utils.CountBandwidthAndWriteToFile("Z_p", utils.Bandwidth(Z_p, fs, db), fileName)
	utils.CountBandwidthAndWriteToFile("Z_f", utils.Bandwidth(Z_f, fs, db), fileName)
}
