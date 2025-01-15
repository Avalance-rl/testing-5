package tests

func PositivePow(number int, degree uint) int {
	res := 1
	for range degree {
		res *= number
	}
	return res
}
