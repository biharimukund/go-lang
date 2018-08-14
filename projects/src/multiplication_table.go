package main

import "fmt"

func main() {
	fmt.Println("Enter the number you want to generate the multiplication table:")
	var number int
	fmt.Scan(&number)
	fmt.Println("Below is the table of", number)
	value :=0
	for num :=1; num<=10; num++ {
		value = number * num
		fmt.Println( number, "x", num, "=", value)
	}
}
