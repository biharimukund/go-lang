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
func floyd_triangle() {
	
        c := exec.Command("clear")
        c.Stdout = os.Stdout
        c.Run()	
	separator()
	fmt.Print("This program prints Floyd's triangle upto n rows.\n")
	separator()

	var row, tmp int
	fmt.Print("Enter the no. of rows for the floyd's triangle: ")
	fmt.Scan(&row)
	tmp=1
	for i := 0; i <= row; i++ {
		
		for p :=1; p <=i; p++ {
			fmt.Print(tmp, " ")
			tmp ++
		}
		fmt.Print("\n")
	}
	separator()
}

func main() {
	floyd_triangle()
}
