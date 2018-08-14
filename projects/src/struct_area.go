package main

import (
	"fmt"
	"os"
	"os/exec"
)

type rectangle struct {

	length float32
	breadth float32
	
	geometry struct {
		area float32
		perimeter float32
	}
}

type circle struct {

	radius float32
	
	geometry struct {
		area float32
		circumference float32
	}
}

func rect_area() {

	var rect rectangle
	fmt.Print("Enter the length of rectangle: ")
	fmt.Scan(&rect.length)
	fmt.Print("Enter the breadth of rectangle: ")
	fmt.Scan(&rect.breadth)

	rect.geometry.area = rect.length * rect.breadth
	rect.geometry.perimeter = 2 * ( rect.length + rect.breadth )

	fmt.Print("The area of rectangle is: ", rect.geometry.area, "\n")
	fmt.Print("The perimeter of rectangle is: ", rect.geometry.perimeter, "\n")
	return
}

func circle_area() {
	var circ circle
	fmt.Print("Enter the radius of circle: ")
	fmt.Scan(&circ.radius)

	circ.geometry.area = 3.14 * circ.radius * circ.radius
	circ.geometry.circumference = 2 * 3.14 * circ.radius

	fmt.Print("The area of circle is: ", circ.geometry.area, "\n")
	fmt.Print("The circumference of circle is: ", circ.geometry.circumference, "\n")
	return
}

func separator() {

	fmt.Print("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	return
}

func repeat() {

	fmt.Print("Do you want to continue further? Y/N: ")
	var input string
	fmt.Scan(&input)
	if ( input == "Y" ) {
		main ()
	} else {
		fmt.Print("Exiting the script!\n")
		separator()
	}
	return
}
func main() {

        c := exec.Command("clear")
        c.Stdout = os.Stdout
        c.Run()

	separator()
	fmt.Print("This program calculates area of rectangle and circle\n")	
	separator()
	
	var shape string
	fmt.Print("Enter the shape, rectangle or circle: ")
	fmt.Scan(&shape)

	if ( shape == "rectangle" ){
		rect_area()
		separator()
		repeat()
	} else if ( shape == "circle" ) {
		circle_area()
		separator()
		repeat()
	} else {
		fmt.Print("You have entered wrong shape!\n")
		separator()
		repeat()
	}
}
