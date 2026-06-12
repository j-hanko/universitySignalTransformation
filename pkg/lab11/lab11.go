package lab11

import (
	"fmt"
	"universitySignalTransformation/pkg/utils"
)

const (
	n = 15
	k = 11
	m = 4
)

func createPMatrix() [][]byte {
	P := [][]byte{
		{1, 1, 0, 0},
		{1, 0, 1, 0},
		{0, 1, 1, 0},
		{1, 1, 1, 0},
		{1, 0, 0, 1},
		{0, 1, 0, 1},
		{1, 1, 0, 1},
		{0, 0, 1, 1},
		{1, 0, 1, 1},
		{0, 1, 1, 1},
		{1, 1, 1, 1},
	}

	return P
}

func HammingCode(inputWord string) (result []byte) {
	fmt.Println("Kodowanie Hamminga\n")

	inputTMP := utils.ASCII_to_bit(inputWord)
	input := utils.FromIntSliceToBitSlice(inputTMP)

	if r := len(input) % k; r != 0 {
		input = append(make([]byte, k-r), input...)
	}

	fmt.Println("Bity wejsciowe: ", input)
	fmt.Println("Długość wejciowa: ", len(input), "\n---")

	P := createPMatrix()

	for i := 0; i < len(input); i = i + k {
		byteTMP := make([]byte, n)

		x3 := input[i]
		x5 := input[i+1]
		x6 := input[i+2]
		x7 := input[i+3]
		x9 := input[i+4]
		x10 := input[i+5]
		x11 := input[i+6]
		x12 := input[i+7]
		x13 := input[i+8]
		x14 := input[i+9]
		x15 := input[i+10]

		for column := 0; column < m; column++ {
			var p byte = 0

			for row := 0; row < k; row++ {
				multiplication := input[i+row] * P[row][column]
				p = p ^ multiplication
			}

			byteTMP[column] = p
		}

		for j := 0; j < k; j++ {
			byteTMP[m+j] = input[i+j]
		}

		result = append(result, byteTMP...)

		fmt.Print("Bity wejściowe: ")
		fmt.Println(x3, x5, x6, x7, x9, x10, x11, x12, x13, x14, x15, " Bity wyjściowe:", byteTMP)
	}

	fmt.Println("---\nBity wyjściowe: ", result)
	fmt.Println("Długość wyjściowa: ", len(result))

	if (len(input) / k * n) == len(result) {
		fmt.Println("Info: Poprawnie przeliczono długość\n\n")
	}

	return result
}

func HammingDecode(input []byte) (result []byte) {
	fmt.Println("Dekodowanie Hamminga\n")

	P := createPMatrix()

	for i := 0; i < len(input); i = i + n {
		x1 := input[i]
		x2 := input[i+1]
		x4 := input[i+2]
		x8 := input[i+3]
		x3 := input[i+4]
		x5 := input[i+5]
		x6 := input[i+6]
		x7 := input[i+7]
		x9 := input[i+8]
		x10 := input[i+9]
		x11 := input[i+10]
		x12 := input[i+11]
		x13 := input[i+12]
		x14 := input[i+13]
		x15 := input[i+14]

		s := make([]byte, m)

		s[0] = x1
		s[1] = x2
		s[2] = x4
		s[3] = x8

		data := []byte{x3, x5, x6, x7, x9, x10, x11, x12, x13, x14, x15}

		for column := 0; column < m; column++ {
			for row := 0; row < k; row++ {
				multiplication := data[row] * P[row][column]
				s[column] = s[column] ^ multiplication
			}
		}

		S := s[0]*1 + s[1]*2 + s[2]*4 + s[3]*8

		if S != 0 {
			input[i+int(S)-1] ^= 1

			x3 = input[i+4]
			x5 = input[i+5]
			x6 = input[i+6]
			x7 = input[i+7]
			x9 = input[i+8]
			x10 = input[i+9]
			x11 = input[i+10]
			x12 = input[i+11]
			x13 = input[i+12]
			x14 = input[i+13]
			x15 = input[i+14]
		}

		result = append(
			result,
			x3, x5, x6, x7, x9, x10, x11, x12, x13, x14, x15,
		)
	}

	fmt.Println("Bity zdekodowane: ", result)
	fmt.Println("Długość zdekodowanego kodu: ", len(result))

	return result
}
