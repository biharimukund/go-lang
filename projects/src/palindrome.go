package main

import (
        "fmt"
        "os"
        "os/exec"
)

func main() {

        c := exec.Command("clear")
        c.Stdout = os.Stdout
        c.Run()
	var number, remainder, reverse_no, original_no int	
	fmt.Print("Enter the number to be tested: ")
	fmt.Scan(&number)
	reverse_no = 0
	original_no = number
	for{
		remainder = number % 10
		reverse_no =  reverse_no*10 + remainder
		number = number /10
		if ( number == 0 ) {
			break
		}
	}
	if ( reverse_no == original_no ) {
		fmt.Print("Your number ", original_no, " is pallindrome!\n")
	} else {
		fmt.Print("Your number ", original_no, " is not pallindrome.\n")
	}
}
