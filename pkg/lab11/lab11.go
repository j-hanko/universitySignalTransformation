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

func createIdentityMatrix(size int) [][]byte {
	I := make([][]byte, size)

	for row := 0; row < size; row++ {
		I[row] = make([]byte, size)

		for column := 0; column < size; column++ {
			if row == column {
				I[row][column] = 1
			} else {
				I[row][column] = 0
			}
		}
	}

	return I
}

func createGMatrix() [][]byte {
	P := createPMatrix()
	I := createIdentityMatrix(k)

	G := make([][]byte, k)

	for row := 0; row < k; row++ {
		G[row] = make([]byte, n)

		for column := 0; column < m; column++ {
			G[row][column] = P[row][column]
		}

		for column := 0; column < k; column++ {
			G[row][m+column] = I[row][column]
		}
	}

	return G
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

	G := createGMatrix()

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

		b := []byte{x3, x5, x6, x7, x9, x10, x11, x12, x13, x14, x15}

		for column := 0; column < n; column++ {
			var sum byte = 0

			for row := 0; row < k; row++ {
				multiplication := b[row] * G[row][column]
				sum = sum ^ multiplication
			}

			byteTMP[column] = sum
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
