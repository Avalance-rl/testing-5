package tests

func MultipleEachSliceElements(slice []int, multiplier int) []int {
	for i := 0; i < len(slice); i++ {
		slice[i] = slice[i] * multiplier
	}

	return slice
}
