package week1

func slices() (result []string) {
	array := []string{"Freya", "Athena", "Angel", "Mika", "Kazumi"}
	result = array[0:3] // index 1 to length of 5 []
	return result
}

func Slices() []string{
	return slices()
}