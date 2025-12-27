package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1, str2 string

	
	fmt.Print("Enter main string: ")
	fmt.Scan(&str1)

	fmt.Print("Enter substring to search: ")
	fmt.Scan(&str2)

	
	if strings.Contains(str1, str2) {
		fmt.Println("The first string contains the second string (substring found).")
	} else {
		fmt.Println("Substring not found.")
	}
}

