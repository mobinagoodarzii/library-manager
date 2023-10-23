package main

import (
	"fmt"
	"github.com/mobinagoodarzii/models"
)

func main() {
	//Check database
	models.CheckInitialization()

	//Library operations
	for {
		models.CallClear()
		fmt.Println("Welcome to the Central Library!\nEnter number 1 to view the books menu\nEnter number 2 to view the members menu\nEnter number 3 to exchange books\nAnd enter 0 to exit")

		var inputFirstMenu string
		fmt.Scanln(&inputFirstMenu)
		models.CallClear()
		if inputFirstMenu == "1" {
			models.BooksMenu()
			continue
		} else if inputFirstMenu == "2" {
			models.MembersMenu()
			continue
		} else if inputFirstMenu == "0" {
			fmt.Println("Goodbye")
			break
		} else {
			fmt.Println("invalid selection")
			continue
		}

	}
}
