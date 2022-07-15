//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
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
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkIn.String()
		}
		fmt.Println(member.name, ":", title, ":", audit.checkOut.String(), "through", returnTime)
	}
}

func printMembersAudit(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}
}

func printLibraryBooks(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println(title, "/ total:", book.total, "/ lended:", book.lended)
	}
}

func checkOutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of library")
		return false
	}

	if book.lended == book.total {
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
		fmt.Println("Book not part of library")
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
		books:   make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}

	library.books["Webapps in Go"] = BookEntry{
		total:  4,
		lended: 0,
	}
	library.books["Webapps in JavaScript"] = BookEntry{
		total:  9,
		lended: 0,
	}
	library.books["Webapps in NodeJs"] = BookEntry{
		total:  2,
		lended: 0,
	}
	library.books["Webapps in Go"] = BookEntry{
		total:  6,
		lended: 0,
	}

	library.members["Jason"] = Member{"Jason", make(map[Title]LendAudit)}
	library.members["Shubham"] = Member{"Shubham", make(map[Title]LendAudit)}
	library.members["Ruchi"] = Member{"Ruchi", make(map[Title]LendAudit)}
	library.members["Vineet"] = Member{"Vineet", make(map[Title]LendAudit)}

	fmt.Println("\n Initial:")
	printLibraryBooks(&library)
	printMembersAudit(&library)

	memeber := library.members["Jason"]
	checkedout := checkOutBook(&library, "Webapps in Go", &memeber)
	fmt.Printf("\nCheckout a book")

	if checkedout {
		printLibraryBooks(&library)
		printMembersAudit(&library)
	}

	retured := returnBook(&library, "Webapps in Go", &memeber)
	fmt.Printf("\nCheckin a book")
	if retured {
		printLibraryBooks(&library)
		printMembersAudit(&library)
	}
}
