package week2

func KeyCheck() string {
	laptopDetails := map[string]string{
		"Brand": "Asus",
		"Model": "E410k",
		"Color": "Black",
		"Year":  "2021",
	}

	//check the key only not a value (the "_" means value)
	_, ok := laptopDetails["Brand"] // will output boolean
	//check the key if existing
	_, ok2 := laptopDetails["Spec"]

	var result string

	if ok == true {
		value := laptopDetails["Brand"]
		result += "Yes the brand is " + value
	} else {
		result += "The brand is not existing"
	}

	if ok2 == true {
		value := laptopDetails["Spec"]
		result += "; Yes the Spec is " + value
	} else {
		result += "; The Spec is not existing"
	}
	return result
}
