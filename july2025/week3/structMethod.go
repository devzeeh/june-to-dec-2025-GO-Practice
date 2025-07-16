package week3

import (
	"fmt"
	"strconv"
)

// define struct 'laptop' to represent the specifications of a laptop
type laptop struct {
	cpu          string // Fields
	ram          int
	storage      int
	manufacturer string
}

// Value receiver method that returns a combination of the CPU and manufacturer
// Get the laptop cpu and manufacturer
func (l laptop) cpuManufacturer() string {
	res := l.cpu + ", " + l.manufacturer
	return res
}

// Pointer receiver method to modify the actual storage field of the laptop
// Increases the storage by the specified size (in GB)
func (l *laptop) upgradeStorage(size int) {
	l.storage += size // The size is equal to 100 [+= equivalent to x = x + 1
}

func StructFieldMethod() (string, string, string) {
	// Create a new laptop instance with initial specifications
	asus := laptop{"Intel", 16, 256, "Asus"}

	// Upgrade the laptop's storage by 100 GB
	asus.upgradeStorage(100)

	laptopStorage := fmt.Sprint("Laptop Storage: ", strconv.Itoa(asus.storage)+"GB") //  Laptop storage: 356GB
	laptopRam := fmt.Sprint(" Laptop Ram: ", strconv.Itoa(asus.ram)+"GB")            //Laptop Ram: 16GB
	laptopCpuManu := fmt.Sprint(" Cpu and Manufacturer: ", asus.cpuManufacturer())   //Cpu and Manufacturer: AMD Asus

	return laptopStorage, laptopRam, laptopCpuManu
}
