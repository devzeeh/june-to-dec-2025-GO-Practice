package july
import (
	"fmt"
	"strconv"
)

func Inventory() (string, string, string, string, string, string, string) {
	var a, initialInventory, initial, addInv, addStock, updateStock, totalQty string
	inventory := map[string]string{"Nescafe": "21", "Kopiko": "50", "Zesto": "10"}

	a += fmt.Sprintln("Inventory!!!")
	for items, qty := range inventory {
		initialInventory += fmt.Sprintln(items, qty)
	}

	// sum all the quantity
	var total1 int
	for _, qty1 := range inventory {
		num1, _ := strconv.Atoi(qty1)
		
		/*if err1 != nil {
			fmt.Sprintln("Errors", qty1)
			continue
		}*/
		total1 += num1 // += equavalent to  x = x + num
	}

	initial += fmt.Sprintln("Initial Stock:", total1)

	//Add to inventory
	addInv += fmt.Sprintln("Add inventory... \nDove 34 \nHead & Shoulder 9")
	inventory["Dove Soap"] = "34"
	inventory["Head & Shoulder"] = "9"

	for stock, quantity := range inventory {
		addStock += fmt.Sprintln(stock, quantity)
	}

	//update quantity
	updateStock += fmt.Sprintln("Updated the stock quantity... \nUpdate Kopiko to 23")
	inventory["Kopiko"] = "23"
	updatedStock := inventory

	// sum all the quantity
	var total int
	for _, qty := range updatedStock {
		num, _ := strconv.Atoi(qty)
		total += num // += equavalent to  x = x + num
	}

	totalQty += fmt.Sprintln("Total quantity:", total)


	return a, initialInventory, initial, addInv, addStock, updateStock, totalQty
}
