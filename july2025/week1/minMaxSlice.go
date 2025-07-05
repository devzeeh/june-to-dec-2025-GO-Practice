package week1

func MinMaxSlice() (int, int) {
	number := []int{90, 938, 78, 69, 90, 100, 1000, 20}

	min := number[0]
	max := number[0]
	for _, v := range number {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}
