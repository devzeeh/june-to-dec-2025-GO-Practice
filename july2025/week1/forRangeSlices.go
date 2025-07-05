package week1

func RangeSlice() (res string) {
	fruits := []string{"Freya", "Kazumi", "Angel", "Mika", "Athena"}

	for _, val := range fruits {
		res += val + " "
	}
	return res
}