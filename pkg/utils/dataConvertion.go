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
