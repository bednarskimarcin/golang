package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Only functions which names starts with UpperCase letter are visible outside the package.
func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0, errors.New("file not found")
	}

	valueText := string(data) //byte array to string
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New("failed to parse the stored value")
	}
	return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)                  //string from float
	os.WriteFile(fileName, []byte(valueText), 0644) //conversion with []byte(string)
}
