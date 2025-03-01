package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 0.0, errors.New("failed to read file")
	}

	valueText := string(data)
	// float64() will not work with data directly from os.ReadFile so we convert to string and use strconv.ParseFloat
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 0.0, errors.New("failed to parse stored value")
	}

	return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
