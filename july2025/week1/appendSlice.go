package week1

func appending() []string{
	// add "Athena" to array
	array := []string{"Freya"}
	add := append(array, "Athena")

	// Add "Angel", "Mika", "Kazumi"
	add = append(add, "Angel", "Mika", "Kazumi")
	return add
}

func AppendSlice() []string{
	return appending()
	/*array := []string{"Freya"}
	add := append(array, "Athena")
	//more := []string{"Angel", "Mika", "Kazumi"}
	add = append(add, "Angel", "Mika", "Kazumi")
	remove := append(add[:3], add[4:]...)
	//return add
	fmt.Println(add)
	fmt.Println("remove: ", remove)
	fmt.Println("length: ", len(add))*/
	
}