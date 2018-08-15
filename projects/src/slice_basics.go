package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
)

func separator() {

	fmt.Printf("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
}

func clear() {
        c := exec.Command("clear")
        c.Stdout = os.Stdout
        c.Run()
	return
}	
func slice_subject() {

	var subject = []string{"maths", "science", "civics", "hindi", "computer", "history", "economics"}
	
	separator()
	fmt.Printf("Total no. of subjects are: %v\n", len(subject))
	fmt.Printf("All subjects in the slice are: %v\n", subject)
	separator()
	fmt.Printf("\nSubject at 4th element of slice is: %v\n", subject[3])
	fmt.Printf("Subjects after 2nd element of slice are: %v\n", subject[2:])
	fmt.Printf("Subjects b/w 3rd element and 6th element of slice are: %v\n", subject[3:5])
	fmt.Printf("Last 3 elements of slice are: %v\n", subject[len(subject)-3:len(subject)])
	separator()
	sort.Strings(subject)
	fmt.Printf("Subjects sorted: %v\n", subject)
	separator()	
	return
}

func main() {
	
	clear()
	separator()
	fmt.Print("This program helps understand the basics of slice in golang.\n")
	separator()
	var user_input string
	fmt.Print("Do you want to play with slice of subjects? Y/N: \n")
	fmt.Scan(&user_input)
	if ( user_input == "Y" ) {
		clear()
		slice_subject()
	}
}
