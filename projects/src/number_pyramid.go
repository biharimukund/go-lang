package main

import (
	"fmt"
	"os"
	"os/exec"
)

func separator() {
	fmt.Print("**************************************************\n")
	return
}
func pascal_triangle() {
	
        c := exec.Command("clear")
        c.Stdout = os.Stdout
        c.Run()	
	separator()
	fmt.Print("This program prints pyramid of numbers upto n rows.\n")
	separator()

	var row, tmp int
	fmt.Print("Enter the no. of rows to print pyramid of numbers: ")
	fmt.Scan(&row)
	
	if ( row >= 10 ) {
		fmt.Print("Pyramid won't be symmetrical for double digit numbers!\n")
		os.Exit(1)
	}
	
	for i := 1; i <= row; i++ {

		for j := 1; j <= row-i; j++ {
		 	fmt.Print(" ") 
		}
		tmp=1	
		for k := 1; k <= i; k++ {
			fmt.Print(tmp)
			tmp++
		}
			for n := tmp-2; n >=1; n-- {
				fmt.Print(n)
			}			
		fmt.Print("\n")
	}
}

func main() {
	pascal_triangle()
}
