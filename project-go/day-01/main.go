package main

import (
	"fmt"
)

func userinput() (string, string) {
	fmt.Println("Enter your first name : ")
	var firstname string
	fmt.Scanln(&firstname)
	fmt.Println("Enter you last name : ")
	var lastname string
	fmt.Scanln(&lastname)
	return firstname, lastname	
}


func main() {
	
	total := 50
	firstname, lastname := userinput()
	fmt.Printf("Hello %s %s\n", firstname, lastname)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Printf("Enter No of ticket : ")
		var ticket int
		fmt.Scan(&ticket)
		fmt.Printf("Enter No of ticket : ")
		fmt.Printf("thank you for booking %v tickets\n", ticket)
		total -= ticket
		fmt.Printf("Remaining tickets %v\n", total)
		if total <= 0 {
			fmt.Println("No more tickets available")
			break
		} 
	}
}