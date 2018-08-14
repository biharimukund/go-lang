package main

import (
	"fmt"
	"os"
	"os/exec"
)
	
func separator() {

        fmt.Print("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
        return
}

func main() {
	
	
        c := exec.Command("clear")
        c.Stdout = os.Stdout
        c.Run()

	test_array := [7]int{23,235,435,674,969,20,924}
	var largest, tmp int

	separator()	

	fmt.Print("This script finds the largest number in the array.\n")

	separator()

	for i :=0; i < len(test_array); i++ {
		fmt.Println("The value stored in test_array[",i, "] is: ", test_array[i])
	}

	separator()

	tmp = test_array[0]
	for i :=1; i < len(test_array); i++ {
		if (  tmp > test_array[i] ) {
			largest = tmp
		} else {
			largest = test_array[i]
			tmp = largest
		}
	}	

	fmt.Println("The largest number in the array is: ", largest)
	
	separator()
}
