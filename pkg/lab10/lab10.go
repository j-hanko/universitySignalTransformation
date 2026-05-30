package lab10

import (
	"fmt"
	"universitySignalTransformation/pkg/utils"
)

func HammingCode(inputWord string) (result []byte) {
	fmt.Println("Kodowanie Hamminga\n")
	inputTMP := utils.ASCII_to_bit(inputWord)
	input := utils.FromIntSliceToBitSlice(inputTMP)
	if r := len(input) % 4; r != 0 {
		input = append(make([]byte, 4-r), input...)
	}
	fmt.Println("Bity wejsciowe: ", input)
	fmt.Println("Długość wejciowa: ", len(input), "\n---")
	for i := 0; i < len(input); i = i + 4 {
		byteTMP := make([]byte, 7)
		x3 := input[i]
		x5 := input[i+1]
		x6 := input[i+2]
		x7 := input[i+3]

		x1 := x3 ^ x5 ^ x7
		x2 := x3 ^ x6 ^ x7
		x4 := x5 ^ x6 ^ x7

		byteTMP[0] = x1
		byteTMP[1] = x2
		byteTMP[2] = x3
		byteTMP[3] = x4
		byteTMP[4] = x5
		byteTMP[5] = x6
		byteTMP[6] = x7

		result = append(result, byteTMP...)
		fmt.Println("Bity wejściowe: ", x3, x5, x6, x7, " Bity wyjściowe:", byteTMP)

	}
	fmt.Println("---\nBity wyjściowe: ", result)
	fmt.Println("Długość wyjściowa: ", len(result))
	if (len(input) / 4 * 7) == len(result) {
		fmt.Println("Info: Poprawnie przeliczono długość\n\n")
	}
	return result
}

func HammingDecode(input []byte) (result []byte) {
	fmt.Println("Dekodowanie Hamminga\n")
	for i := 0; i < len(input); i = i + 7 {
		block := make([]byte, 7)
		copy(block, input[i:i+7])

		x1prim := input[i+2] ^ input[i+4] ^ input[i+6]
		x2prim := input[i+2] ^ input[i+5] ^ input[i+6]
		x4prim := input[i+4] ^ input[i+5] ^ input[i+6]

		x1m := input[i] ^ x1prim
		x2m := input[i+1] ^ x2prim
		x4m := input[i+3] ^ x4prim

		S := x1m*1 + x2m*2 + x4m*4

		if S != 0 {
			block[S-1] ^= 1
		}
		result = append(result, block[2], block[4], block[5], block[6])

	}
	fmt.Println("Długość zdekodowanego kodu: ", len(result))
	return result
}
