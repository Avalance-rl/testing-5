package tests

import "errors"

var ErrNegativeNumber = errors.New("negative number not allowed")

func CheckPositiveNumber(n int) error {
	if n < 0 {
		return ErrNegativeNumber
	}
	return nil
}
