package models

import (
	"fmt"
	"log"
	"time"
)

type Member struct {
	joinId           int
	firstname        string
	lastname         string
	age              int64
	dateRegistration time.Time
	status           string
}

var Members = make(map[int]Member)

func (m *Member) CreateMembers() {

	scanStrInput("Enter your firstname:", &m.firstname)

	scanStrInput("Enter your lastname:", &m.lastname)

	scanIntInput("Enter your age:", &m.age)

	var existingMember *Member
	for _, member := range Members {
		if member.firstname == m.firstname && member.lastname == m.lastname && member.age == m.age {
			existingMember = &member
			break
		}
	}
	if existingMember != nil {
		CallClear()
		fmt.Printf("A member with the same details already exists.\n")

	} else {
		m.joinId = len(Members) + 1
		m.status = "available"
		m.dateRegistration = time.Now()
		Members[m.joinId] = *m

		//Transfer information to the database
		db, err := ConnectToDB()
		defer db.Close()
		if err != nil {
			log.Fatal(err)
		}
		stmt, err := db.Prepare("INSERT INTO `library`.`members` (joinId, firstname, lastname, age, dateRegistration, status) VALUES(?, ?, ?, ?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		_, err = stmt.Exec(m.joinId, m.firstname, m.lastname, m.age, m.dateRegistration.Format("2006-01-02 15:04:05"), m.status)
		if err != nil {
			log.Fatal(err)
		}

		CallClear()
		fmt.Printf("The member was successfully created.\nmemberId:%d, firstname:%s, lastname:%s, age:%d, dateRegistration:%q, status:%s\n", m.joinId, m.firstname, m.lastname, m.age, m.dateRegistration.Format("2006-01-02 15:04:05"), m.status)
	}
}

func (m Member) EditMembers() {
	for {
		CallClear()
		fmt.Println("If you know the ID of member you want to edit, please enter it. otherwise enter number 0.\nenter -1 return to the previous menu.")
		fmt.Scanln(&inputID)
		if inputID == -1 {
			break
		} else if inputID == 0 {
			CallClear()
			for key, value := range Members {
				fmt.Printf("memberId:%d, firstname:%s, lastname:%s, age:%d, dateRegistration:%q, status:%s\n", key, value.firstname, value.lastname, value.age, value.dateRegistration.Format("2006-01-02 15:04:05"), value.status)
			}
			pressEnter()
			continue
		} else if member, exists := Members[inputID]; exists {
			if member.status != "deleted" {
				CallClear()
				fmt.Printf("Editing member with ID %d\n", inputID)

				scanNewStrInput("Enter new firstname (press Enter to keep previous value): ", &member.firstname)

				scanNewStrInput("Enter new lastname (press Enter to keep previous value): ", &member.lastname)

				scanNewIntInput("Enter new age (press Enter to keep previous value): ", &member.age)

				Members[inputID] = member

				//Transfer information to the database
				db, err := ConnectToDB()
				defer db.Close()
				if err != nil {
					log.Fatal(err)
				}
				stmt, err := db.Prepare("UPDATE `library`.`members` SET firstname=?, lastname=?, age=?, dateRegistration=?, status=? WHERE joinId=?")
				if err != nil {
					log.Fatal(err)
				}
				_, err = stmt.Exec(member.firstname, member.lastname, member.age, member.dateRegistration.Format("2006-01-02 15:04:05"), member.status, inputID)
				if err != nil {
					log.Fatal(err)
				}

				CallClear()
				fmt.Printf("The member was successfully edited.\nmemberId:%d, firstname:%s, lastname:%s, age:%d, dateRegistration:%q, status:%s\n", inputID, member.firstname, member.lastname, member.age, member.dateRegistration.Format("2006-01-02 15:04:05"), member.status)
				pressEnter()
			} else {
				CallClear()
				fmt.Println("The desired member has been deleted")
			}
		} else {
			CallClear()
			fmt.Println("There is no member with this ID")
		}
	}
}

func (m Member) DeleteMembers() {
	for {
		CallClear()
		fmt.Println("If you know the ID of member you want to delete, please enter it. otherwise enter number 0.\nenter -1 return to the previous menu.")
		fmt.Scanln(&inputID)
		if inputID == -1 {
			break
		} else if inputID == 0 {
			CallClear()
			for key, value := range Members {
				fmt.Printf("memberId:%d, firstname:%s, lastname:%s, age:%d, dateRegistration:%q, status:%s\n", key, value.firstname, value.lastname, value.age, value.dateRegistration.Format("2006-01-02 15:04:05"), value.status)
			}
			pressEnter()
			continue
		} else if member, exists := Members[inputID]; exists {
			if member.status != "deleted" {
				CallClear()
				fmt.Printf("Deleted member with ID %d\n", inputID)
				member.status = "deleted"
				Members[inputID] = member

				//Transfer information to the database
				db, err := ConnectToDB()
				defer db.Close()
				if err != nil {
					log.Fatal(err)
				}
				stmt, err := db.Prepare("UPDATE `library`.`members` SET status=? WHERE joinId=?")
				if err != nil {
					log.Fatal(err)
				}
				_, err = stmt.Exec(member.status, inputID)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("The member was successfully deleted.\nmemberId:%d, firstname:%s, lastname:%s, age:%d, dateRegistration:%q, status:%s\n", inputID, member.firstname, member.lastname, member.age, member.dateRegistration.Format("2006-01-02 15:04:05"), member.status)
				pressEnter()
			} else {
				CallClear()
				fmt.Println("The desired member has been deleted")
			}
		} else {
			CallClear()
			fmt.Println("There is no member with this ID")
		}
	}
}

func (m Member) ReadMembers() {
	for {
		var memberFound bool = false
		CallClear()
		fmt.Println("Enter number 1 to view all library members\nEnter number 2 to view members by name filter\nEnter number 3 to view members by last name filter\nEnter number 4 to view members by age filter\nEnter number 5 to view members by ID filter\nEnter number 6 to return to the previous menu")
		var inputSixthMenu string
		fmt.Scanln(&inputSixthMenu)

		if inputSixthMenu == "1" {
			CallClear()
			for key, value := range Members {
				fmt.Printf("memberId:%d, firstname:%s, lastname:%s, age:%d, dateRegistration:%q, status:%s\n", key, value.firstname, value.lastname, value.age, value.dateRegistration.Format("2006-01-02 15:04:05"), value.status)
			}
			//Transfer information to the database
			db, err := ConnectToDB()
			defer db.Close()
			if err != nil {
				log.Fatal(err)
			}
			// Execute a SELECT query
			rows, err := db.Query("SELECT * FROM library.members ")
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()
			var (
				joinId, age                 int
				firstName, lastName, status string
				dateRegistration            time.Time
				dateRegistrationRaw         []byte
			)

			for rows.Next() {
				err = rows.Scan(&joinId, &firstName, &lastName, &age, &dateRegistrationRaw, &status)
				if err != nil {
					log.Fatal(err)
				}
				dateRegistrationStr := string(dateRegistrationRaw)
				dateRegistration, err = time.Parse("2006-01-02 15:04:05", dateRegistrationStr)
				if err != nil {
					log.Fatal(err)
				}
				log.Println(joinId, firstName, lastName, age, dateRegistration, status)
			}
			err = rows.Err()
			if err != nil {
				log.Fatal(err)
			}
			pressEnter()
		} else if inputSixthMenu == "2" {
			fmt.Println("Please write the first name of the member:")
			var inputMemberFirstName string
			fmt.Scanln(&inputMemberFirstName)
			for key, value := range Members {
				if value.firstname == inputMemberFirstName {
					fmt.Printf("memberId:%d, firstname:%s, lastname:%s, age:%d, dateRegistration:%q, status:%s\n", key, value.firstname, value.lastname, value.age, value.dateRegistration.Format("2006-01-02 15:04:05"), value.status)
					pressEnter()
					memberFound = true
				}
			}
			if memberFound == false {
				fmt.Println("There is no member with this first name")
				pressEnter()
			}
		} else if inputSixthMenu == "3" {
			fmt.Println("Please write the last name of the member:")
			var inputLastName string
			fmt.Scanln(&inputLastName)
			for key, value := range Members {
				if value.lastname == inputLastName {
					fmt.Printf("memberId:%d, firstname:%s, lastname:%s, age:%d, dateRegistration:%q, status:%s\n", key, value.firstname, value.lastname, value.age, value.dateRegistration.Format("2006-01-02 15:04:05"), value.status)
					pressEnter()
					memberFound = true
				}
			}
			if memberFound == false {
				fmt.Println("There is no member with this last name")
				pressEnter()
			}
		} else if inputSixthMenu == "4" {
			fmt.Println("Please write the age of the member:")
			var inputMemberAge int64
			fmt.Scanln(&inputMemberAge)
			for key, value := range Members {
				if value.age == inputMemberAge {
					fmt.Printf("memberId:%d, firstname:%s, lastname:%s, age:%d, dateRegistration:%q, status:%s\n", key, value.firstname, value.lastname, value.age, value.dateRegistration.Format("2006-01-02 15:04:05"), value.status)
					pressEnter()
					memberFound = true
				}
			}
			if memberFound == false {
				fmt.Println("There is no member with this age")
				pressEnter()
			}
		} else if inputSixthMenu == "5" {
			fmt.Println("Please enter the member ID:")
			var inputMemberID int
			fmt.Scanln(&inputMemberID)
			for key, value := range Members {
				if value.joinId == inputMemberID {
					fmt.Printf("memberId:%d, firstname:%s, lastname:%s, age:%d, dateRegistration:%q, status:%s\n", key, value.firstname, value.lastname, value.age, value.dateRegistration.Format("2006-01-02 15:04:05"), value.status)
					pressEnter()
					memberFound = true
				}
			}
			if memberFound == false {
				fmt.Println("There is no member with this ID")
				pressEnter()
			}
		} else if inputSixthMenu == "6" {
			return
		} else {
			CallClear()
			fmt.Println("invalid selection")
			pressEnter()
			continue
		}
	}
}

func MembersMenu() {
	for {
		CallClear()
		fmt.Println("Enter number 1 to create a new member\nEnter number 2 to edit members\nEnter number 3 to delete member\nEnter the number 4 to view and search for members\nEnter number 5 to return to the previous menu")
		var inputThirdMenu string
		fmt.Scanln(&inputThirdMenu)

		if inputThirdMenu == "1" {
			var newMember Member
			newMember.CreateMembers()
			pressEnter()
			continue
		} else if inputThirdMenu == "2" {
			var editMember Member
			editMember.EditMembers()
			continue
		} else if inputThirdMenu == "3" {
			var deleteMember Member
			deleteMember.DeleteMembers()
			continue
		} else if inputThirdMenu == "4" {
			var readMember Member
			readMember.ReadMembers()
			pressEnter()
			continue
		} else if inputThirdMenu == "5" {
			return
		} else {
			CallClear()
			fmt.Println("invalid selection")
			pressEnter()
			continue
		}
	}
}
