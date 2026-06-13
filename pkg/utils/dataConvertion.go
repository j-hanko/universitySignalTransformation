package utils

import (
	"github.com/go-echarts/go-echarts/v2/opts"
)

func FromSliceToLineData(slice []float64) []opts.LineData {
	sliceOfData := make([]opts.LineData, 0, len(slice))
	for _, v := range slice {
		sliceOfData = append(sliceOfData, opts.LineData{Value: v})
	}
	return sliceOfData
}

func FromSliceToScatterData(data []float64) []opts.ScatterData {
	items := make([]opts.ScatterData, 0, len(data))
	for _, v := range data {
		items = append(items, opts.ScatterData{Value: v})
	}
	return items
}

func FromIntSliceToBitSlice(slice []int) []byte {
	result := make([]byte, len(slice))

	for i, v := range slice {
		result[i] = byte(v)
	}

	return result
}

func FromBitSliceToIntSlice(bits []byte) []int {
	result := make([]int, 0, len(bits))

	for _, v := range bits {
		result = append(result, int(v))
	}

	return result
}

func FromFloatSliceToIntSlice(f []float64) []int {
	result := make([]int, 0, len(f))
	for _, v := range f {
		result = append(result, int(v))
	}

	return result
}

func InputWordToBits(inputWord string, coderChoose string) []byte {
	inputTMP := ASCII_to_bit(inputWord)
	inputBits := FromIntSliceToBitSlice(inputTMP)

	if coderChoose == "Hamming_7_4" {
		if r := len(inputBits) % 4; r != 0 {
			inputBits = append(make([]byte, 4-r), inputBits...)
		}
	} else if coderChoose == "Hamming_15_11" {
		if r := len(inputBits) % 11; r != 0 {
			inputBits = append(make([]byte, 11-r), inputBits...)
		}
	}

	return inputBits
}
