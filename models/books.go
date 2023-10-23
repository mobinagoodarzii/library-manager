package models

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Book struct {
	bookId           int
	title            string
	name             string
	author           string
	pages            int64
	price            float64
	numberOfBook     int64
	dateRegistration time.Time
	status           string
}

var inputID int
var Books = make(map[int]Book)
var enterKey string

func pressEnter() {
	fmt.Println("\n\nPress Enter to continue...")
	_, err := fmt.Scanln(&enterKey)
	if err != nil {
		return
	}
}

func scanStrInput(msg string, input *string) {
	//The fmt package filters whitespaces
	//Instead use bufio.Scanner to read lines
	CallClear()
	fmt.Println(msg)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		*input = scanner.Text()
	}
}
func scanIntInput(msg string, input *int64) {
	CallClear()
	fmt.Println(msg)
	_, err := fmt.Scanln(input)
	if err != nil {
		return
	}
}

func (b *Book) CreateBooks() {

	b.numberOfBook = 1

	scanStrInput("Enter title:", &b.title)

	scanStrInput("Enter name:", &b.name)

	scanStrInput("Enter author:", &b.author)

	scanIntInput("Enter pages:", &b.pages)

	CallClear()
	fmt.Println("Enter price:")
	_, err := fmt.Scanln(&b.price)
	if err != nil {
		fmt.Println("There is a problem here.")
		os.Exit(1)
		return
	}

	scanIntInput("Enter numberOfBook:", &b.numberOfBook)

	var existingBook *Book
	for _, book := range Books {
		if book.title == b.title && book.name == b.name && book.author == b.author && book.pages == b.pages && book.price == b.price {
			existingBook = &book
			break
		}
	}
	if existingBook != nil {
		CallClear()

		existingBook.numberOfBook += b.numberOfBook
		Books[existingBook.bookId] = *existingBook

		fmt.Printf("A book with the same details already exists.\n If you want to add more numbers to your desired book, select the edit option in the main menu")

	} else {
		b.bookId = len(Books) + 1
		b.status = "available"
		b.dateRegistration = time.Now()
		Books[b.bookId] = *b

		//Transfer information to the database
		db, err := ConnectToDB()
		defer db.Close()
		if err != nil {
			log.Fatal(err)
		}
		stmt, err := db.Prepare("INSERT INTO `library`.`books` (bookId, title, bookName, author, pages, price, numberOfBook, dateRegistration, status) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		_, err = stmt.Exec(b.bookId, b.title, b.name, b.author, b.pages, b.price, b.numberOfBook, b.dateRegistration.Format("2006-01-02 15:04:05"), b.status)
		if err != nil {
			log.Fatal(err)
		}

		CallClear()
		fmt.Printf("The book was successfully created.\nbookId:%d, title:%s, name:%s, author:%s, pages:%d, price:%g, numberOfBook:%d, dateRegistration:%q, status:%s\n", b.bookId, b.title, b.name, b.author, b.pages, b.price, b.numberOfBook, b.dateRegistration.Format("2006-01-02 15:04:05"), b.status)
	}
}

func scanNewStrInput(msg string, input *string) {
	var newInput string
	fmt.Println(msg)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		newInput = scanner.Text()
	}
	if newInput != "" {
		*input = newInput
	}
}

func scanNewIntInput(msg string, input *int64) {
	var newInput int64
	fmt.Println(msg)
	_, err := fmt.Scanln(input)
	if err != nil {
		return
	}
	if newInput != 0 {
		*input = newInput
	}
}

func (b *Book) EditBooks() {
	for {
		CallClear()
		fmt.Println("If you know the ID of the book you want to edit, please enter it. otherwise enter number 0.\nenter -1 return to the previous menu.")
		fmt.Scanln(&inputID)
		if inputID == -1 {
			break
		} else if inputID == 0 {
			CallClear()
			for key, value := range Books {
				fmt.Printf("bookId:%d, title:%s, name:%s, author:%s, pages:%d, price:%g, numberOfBook:%d, dateRegistration:%q, status:%s\n", key, value.title, value.name, value.author, value.pages, value.price, value.numberOfBook, value.dateRegistration.Format("2006-01-02 15:04:05"), value.status)
			}
			pressEnter()
			continue
		} else if book, exists := Books[inputID]; exists {
			if book.numberOfBook != 0 && book.status != "deleted" {
				CallClear()
				fmt.Printf("Editing book with ID %d\n", inputID)
				var newPrice float64

				scanNewStrInput("Enter new title (press Enter to keep previous value): ", &book.title)

				scanNewStrInput("Enter new name (press Enter to keep previous value): ", &book.name)

				scanNewStrInput("Enter new author (press Enter to keep previous value): ", &book.author)

				scanNewIntInput("Enter new pages (press Enter to keep previous value): ", &book.pages)

				fmt.Println("Enter new price (press Enter to keep previous value): ")
				n, err := fmt.Scanf("%f\n", &newPrice)
				if err != nil || n != 1 {
					// handle invalid input
					fmt.Println(n, err)
				}

				if newPrice != 0 {
					book.price = newPrice
				}

				scanNewIntInput("Enter new number of book (press Enter to keep previous value): ", &book.numberOfBook)

				Books[inputID] = book

				//Transfer information to the database
				db, err := ConnectToDB()
				defer db.Close()
				if err != nil {
					log.Fatal(err)
				}
				stmt, err := db.Prepare("UPDATE `library`.`books` SET title=?, bookName=?, author=?, pages=?, price=?, numberOfBook=?, dateRegistration=?, status=? WHERE bookId=?")
				if err != nil {
					log.Fatal(err)
				}
				_, err = stmt.Exec(book.title, book.name, book.author, book.pages, book.price, book.numberOfBook, book.dateRegistration.Format("2006-01-02 15:04:05"), book.status, inputID)
				if err != nil {
					log.Fatal(err)
				}

				CallClear()
				fmt.Printf("The book was successfully edited.\nbookId:%d, title:%s, name:%s, author:%s, pages:%d, price:%g, numberOfBook:%d, dateRegistration:%q, status:%s\n", inputID, book.title, book.name, book.author, book.pages, book.price, book.numberOfBook, book.dateRegistration.Format("2006-01-02 15:04:05"), book.status)
				pressEnter()
			} else {
				CallClear()
				fmt.Println("The desired book has been deleted")
			}
		} else {
			CallClear()
			fmt.Println("There is no book with this ID")
		}
	}
}

func (b Book) DeleteBooks() {
	for {
		CallClear()
		fmt.Println("If you know the ID of the book you want to delete, please enter it. otherwise enter number 0.\nenter -1 return to the previous menu.")
		fmt.Scanln(&inputID)
		if inputID == -1 {
			break
		} else if inputID == 0 {
			CallClear()
			for key, value := range Books {
				fmt.Printf("bookId:%d, title:%s, name:%s, author:%s, pages:%d, price:%g, numberOfBook:%d, dateRegistration:%q, status:%s\n", key, value.title, value.name, value.author, value.pages, value.price, value.numberOfBook, value.dateRegistration.Format("2006-01-02 15:04:05"), value.status)
			}
			pressEnter()
			continue
		} else if book, exists := Books[inputID]; exists {
			if book.numberOfBook != 0 && book.status != "deleted" {
				CallClear()
				fmt.Printf("Deleted book with ID %d\n", inputID)
				book.numberOfBook -= 1
				if book.numberOfBook < 1 {
					book.status = "deleted"
				}
				Books[inputID] = book

				//Transfer information to the database
				db, err := ConnectToDB()
				defer db.Close()
				if err != nil {
					log.Fatal(err)
				}
				stmt, err := db.Prepare("UPDATE `library`.`books` SET numberOfBook=?, status=? WHERE bookId=?")
				if err != nil {
					log.Fatal(err)
				}
				_, err = stmt.Exec(book.numberOfBook, book.status, inputID)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("The book was successfully deleted.\nbookId:%d, title:%s, name:%s, author:%s, pages:%d, price:%g, numberOfBook:%d, dateRegistration:%q, status:%s\n", inputID, book.title, book.name, book.author, book.pages, book.price, book.numberOfBook, book.dateRegistration.Format("2006-01-02 15:04:05"), book.status)
				pressEnter()
			} else {
				CallClear()
				fmt.Println("The desired book is not available!")
				pressEnter()
			}
		} else {
			CallClear()
			fmt.Println("There is no book with this ID")
		}
	}
}

func (b Book) ReadBooks() {
	for {
		var bookFound bool = false
		CallClear()
		fmt.Println("Enter number 1 to see all the books in the library\nEnter number 2 to view books by name filter\nEnter number 3 to view books by ID filter\nEnter number 4 to view books by title filter\nEnter number 5 to view books by author filter\nEnter number 6 to return to the previous menu")
		var inputFifthMenu string
		fmt.Scanln(&inputFifthMenu)

		if inputFifthMenu == "1" {
			CallClear()
			for key, value := range Books {
				fmt.Printf("bookId:%d, title:%s, name:%s, author:%s, pages:%d, price:%g, numberOfBook:%d, dateRegistration:%q, status:%s\n", key, value.title, value.name, value.author, value.pages, value.price, value.numberOfBook, value.dateRegistration.Format("2006-01-02 15:04:05"), value.status)
			}

			//Transfer information to the database
			db, err := ConnectToDB()
			defer db.Close()
			if err != nil {
				log.Fatal(err)
			}
			// Execute a SELECT query
			rows, err := db.Query("SELECT * FROM library.books ")
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()
			var (
				bookId, pages, prices, numberOfBook int
				title, bookName, author, status     string
				dateRegistration                    time.Time
				dateRegistrationRaw                 []byte
			)

			for rows.Next() {
				err = rows.Scan(&bookId, &title, &bookName, &author, &pages, &prices, &numberOfBook, &dateRegistrationRaw, &status)
				if err != nil {
					log.Fatal(err)
				}
				dateRegistrationStr := string(dateRegistrationRaw)
				dateRegistration, err = time.Parse("2006-01-02 15:04:05", dateRegistrationStr)
				if err != nil {
					log.Fatal(err)
				}
				log.Println(bookId, title, bookName, author, pages, prices, numberOfBook, dateRegistration, status)
			}
			err = rows.Err()
			if err != nil {
				log.Fatal(err)
			}

			pressEnter()

		} else if inputFifthMenu == "2" {
			fmt.Println("Please write the name of the book:")
			var inputBookName string
			fmt.Scanln(&inputBookName)
			for key, value := range Books {
				if value.name == inputBookName {
					fmt.Printf("bookId:%d, title:%s, name:%s, author:%s, pages:%d, price:%g, numberOfBook:%d, dateRegistration:%q, status:%s\n", key, value.title, value.name, value.author, value.pages, value.price, value.numberOfBook, value.dateRegistration.Format("2006-01-02 15:04:05"), value.status)
					pressEnter()
					bookFound = true
				}
			}
			if bookFound == false {
				fmt.Println("There is no book with this name")
				pressEnter()
			}
		} else if inputFifthMenu == "3" {
			fmt.Println("Please write the ID of the book:")
			var inputBookID int
			fmt.Scanln(&inputBookID)
			for key, value := range Books {
				if value.bookId == inputBookID {
					fmt.Printf("bookId:%d, title:%s, name:%s, author:%s, pages:%d, price:%g, numberOfBook:%d, dateRegistration:%q, status:%s\n", key, value.title, value.name, value.author, value.pages, value.price, value.numberOfBook, value.dateRegistration.Format("2006-01-02 15:04:05"), value.status)
					pressEnter()
					bookFound = true
				}
			}
			if bookFound == false {
				fmt.Println("There is no book with this ID")
				pressEnter()
			}
		} else if inputFifthMenu == "4" {
			fmt.Println("Please write the title of the book:")
			var inputBookTitle string
			fmt.Scanln(&inputBookTitle)
			for key, value := range Books {
				if strings.Contains(value.title, inputBookTitle) {
					fmt.Printf("bookId:%d, title:%s, name:%s, author:%s, pages:%d, price:%g, numberOfBook:%d, dateRegistration:%q, status:%s\n", key, value.title, value.name, value.author, value.pages, value.price, value.numberOfBook, value.dateRegistration.Format("2006-01-02 15:04:05"), value.status)
					pressEnter()
					bookFound = true
				}
			}
			if bookFound == false {
				fmt.Println("There is no book with this title")
				pressEnter()
			}
		} else if inputFifthMenu == "5" {
			fmt.Println("Please write author of the book:")
			var inputBookAuthor string
			fmt.Scanln(&inputBookAuthor)
			for key, value := range Books {
				if value.author == inputBookAuthor {
					fmt.Printf("bookId:%d, title:%s, name:%s, author:%s, pages:%d, price:%g, numberOfBook:%d, dateRegistration:%q, status:%s\n", key, value.title, value.name, value.author, value.pages, value.price, value.numberOfBook, value.dateRegistration.Format("2006-01-02 15:04:05"), value.status)
					pressEnter()
					bookFound = true
				}
			}
			if bookFound == false {
				fmt.Println("There is no book with this author")
				pressEnter()
			}
		} else if inputFifthMenu == "6" {
			return
		} else {
			CallClear()
			fmt.Println("invalid selection")
			pressEnter()
			continue
		}
	}
}

func BooksMenu() {
	for {
		CallClear()
		fmt.Println("Enter number 1 to create a new book\nEnter number 2 to edit the Books\nEnter number 3 to delete the book\nEnter number 4 to view library and search for Books\nEnter number 5 to return to the previous menu")
		var inputSecondMenu string
		fmt.Scanln(&inputSecondMenu)

		if inputSecondMenu == "1" {
			var newBook Book
			newBook.CreateBooks()
			pressEnter()
			continue
		} else if inputSecondMenu == "2" {
			var editBook Book
			editBook.EditBooks()
			continue
		} else if inputSecondMenu == "3" {
			var deleteBook Book
			deleteBook.DeleteBooks()
			continue
		} else if inputSecondMenu == "4" {
			var readBook Book
			readBook.ReadBooks()
			continue
		} else if inputSecondMenu == "5" {
			return
		} else {
			CallClear()
			fmt.Println("invalid selection")
			pressEnter()
			continue
		}
	}
}
