package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CalculateAmountOfSymbolsInFile(t *testing.T) {
	t.Parallel()
	path := "task2_testfile.txt"
	want := 703

	res, err := CalculateAmountOfSymbolsInFile(path)
	assert.NoError(t, err)
	assert.Equal(t, res, want)

	path = "task_testfile.txt"

	res, err = CalculateAmountOfSymbolsInFile(path)
	assert.Error(t, err)
}
