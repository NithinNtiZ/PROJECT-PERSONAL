package main

import "fmt"

func main () {
	var username string 
	var firstname string 
	var lastname string 
	var  total int = 50
	var  num1 int
	var list []int
	
	
	for {
		
		fmt.Println("Enter  your first name: ")
		fmt.Scan(&firstname)
		fmt.Println("Enter  your last name: ")
		fmt.Scan(&lastname)
		fmt.Printf("Hello  %s\n", username)
		fmt.Printf("Hello  %s %s\n", firstname, lastname)
		fmt.Printf("enter no of tickets you want to buy: ")
		fmt.Scan(&num1)
		total = total - num1
		list = append(list, total)
		fmt.Printf("Your have booked %d tickets, Balance no of ticket is %d\n", num1 , total)
	}
	
} 