package july

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//func Week1Result1() {
func Week1Result1() {
	var movies []string

	scan := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Manage your movies")
		fmt.Println("1. Add Movie")
		fmt.Println("2. Display Movie")
		fmt.Println("3. Exit")
		fmt.Print("Choice option: ")
		
		scan.Scan()
		choice := strings.TrimSpace(scan.Text())

		switch choice {
		case "1":
			fmt.Print("Enter Movie name: ")
			scan.Scan()
			name := strings.TrimSpace(scan.Text())

			if name != "" {
				movies = append(movies, name)
				fmt.Println("Movie added successfully.")
			} else {
				fmt.Println("Movie cannot be empty.")
			}
		case "2":
			if len(movies) == 0 {
				fmt.Println("No favorite movies added yet.")
			} else {
				fmt.Println("\nYour Favorite Movies:")
				for i, movie := range movies {
					fmt.Printf("%d. %s\n", i+1, movie)
				}
			}
		case "3":
			fmt.Println("Exiting the program")
			return
		default:
			fmt.Println("Invalid option. Try again")
		}
	}
	
}
