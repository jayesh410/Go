package main

import "fmt"

// Define a structure for Book
type Book struct {
	bookID int
	title  string
	author string
	price  float64
}

func main() {
	var n int

	
	fmt.Print("Enter number of books: ")
	fmt.Scan(&n)

	
	books := make([]Book, n)

	for i := 0; i < n; i++ {
		fmt.Println("\nEnter details of Book", i+1)

		fmt.Print("Book ID: ")
		fmt.Scan(&books[i].bookID)

		fmt.Print("Title: ")
		fmt.Scan(&books[i].title)

		fmt.Print("Author: ")
		fmt.Scan(&books[i].author)

		fmt.Print("Price: ")
		fmt.Scan(&books[i].price)
	}

	fmt.Println("\nBook Details:")
	for i := 0; i < n; i++ {
		fmt.Println("\nBook", i+1)
		fmt.Println("Book ID :", books[i].bookID)
		fmt.Println("Title   :", books[i].title)
		fmt.Println("Author  :", books[i].author)
		fmt.Println("Price   :", books[i].price)
	}
}

