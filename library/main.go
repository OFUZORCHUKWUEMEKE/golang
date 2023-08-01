package main

import (
	// "fmt"
	"fmt"
	"time"
)

type Title string
type Name string

type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}

type Member struct {
	name  Name
	books map[Title]LendAudit
}

type BookEntry struct {
	total  int
	lended int
}

type Library struct {
	member map[Name]Member
	books  map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "not returned Yet"
		} else {
			returnTime = audit.checkIn.String()
		}
		fmt.Println(member.name, ":", title, ":", audit.checkOut.String(), "through", ":", returnTime)
	}
}

func printMemberAudits(library *Library) {
	for _, member := range library.member {
		printMemberAudit(&member)
	}
}

func PrintLibarryBooks(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println(title, "/total:", book.total, "/ lended", book.lended)
	}
}

func checkoutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of library")
		return false
	}
	if book.lended == book.lended {
		fmt.Println("No more books available to lend")
		return false
	}
	book.lended += 1
	library.books[title] = book
	member.books[title] = LendAudit{checkOut: time.Now()}
	return true
}

func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of Library")
		return false
	}
	audit, found := member.books[title]
	if !found {
		fmt.Println("Member did not check out this book")
		return false
	}
	book.lended -= 1
	library.books[title] = book
	audit.checkIn = time.Now()
	member.books[title] = audit
	return true
}

func main() {
	library := Library{
		books:  make(map[Title]BookEntry),
		member: make(map[Name]Member),
	}
	library.books["Webapp in Go"] = BookEntry{
		total:  4,
		lended: 0,
	}
	library.books["Webapp in Go"] = BookEntry{
		total:  3,
		lended: 0,
	}
	library.books["Webapp in Go"] = BookEntry{
		total:  2,
		lended: 0,
	}
	library.books["Webapp in Go"] = BookEntry{
		total:  1,
		lended: 0,
	}

	library.member["Emeke"] = Member{"Emeke", make(map[Title]LendAudit)}
	library.member["Ekenem"] = Member{"Ekenem", make(map[Title]LendAudit)}
	library.member["Faith"] = Member{"Faith", make(map[Title]LendAudit)}

	fmt.Println("\nInitial")
	PrintLibarryBooks(&library)
	printMemberAudits(&library)
	member := library.member["Emeke"]

	checkedOut := checkoutBook(&library, "Webapp in Go", &member)

	fmt.Println("\nCheck out a book")

	if checkedOut {
		printMemberAudits(&library)
		PrintLibarryBooks(&library)

	}
	returned := returnBook(&library, "Webapp in Go", &member)

	fmt.Println("\nChecked in books")

	if returned {
		PrintLibarryBooks(&library)
		printMemberAudits(&library)
	}

}
