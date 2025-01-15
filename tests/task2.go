package tests

import (
	"io"
	"os"
	"unicode/utf8"
)

func CalculateAmountOfSymbolsInFile(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return 0, err
	}

	return utf8.RuneCount(data), nil
}
