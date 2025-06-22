package week4

// This function returns multiple values: an integer, a string, and another string
func multipleReturn(name string) (names string, message string) {
	names = "Hello " + name
	message = "Welcome to Go progamming"
	return
}
func Multiple(names string) (mess string) {
	name, message := multipleReturn(names)
	mess = name + ", " + message
	return mess
}