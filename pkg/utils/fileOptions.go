package utils

import (
	"fmt"
	"os"
)

func WriteToFile(data, fileName string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, data)
	if err != nil {
		panic(err)
	}
}
