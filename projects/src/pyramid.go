package main

import (
	"fmt"
	"os"
	"os/exec"
)

func dot(){
	fmt.Print("********************************************************\n")
	return
}

func pyramid_print() int{
	
	var number int
	fmt.Print("Enter the no. of rows to create a triangle: ")
	fmt.Scan(&number)
	for row := 1; row <= number; row++ {
		
		for count :=1; count <=row; count++ {
			fmt.Print(row)
		}
		fmt.Print("\n")
	}
	return number 
}
func main() {
	
        c := exec.Command("clear")
        c.Stdout = os.Stdout
        c.Run()
	
	dot()
	fmt.Print("This program creates triangle numbers based on input.\n")
	dot()
	pyramid_print()	

}
