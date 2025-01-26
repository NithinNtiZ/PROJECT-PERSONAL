package main

import "fmt"

func main () {
	var username string 
	fmt.Println("Enter  your name:")
	fmt.Scan(&username)
	fmt.Printf("%s Hello World \n", username)



	num1 := 1
	num2 := 2
	sum := num1 + num2
	var list [50]int
	list[0] = sum
	fmt.Printf("Sum of %d and %d is %d", num1, num2, list[0])
}