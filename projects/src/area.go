package main

import (
	"fmt"
	"os"
	"os/exec"
)

func rectangle_area() float32{
	var length, breadth float32
        fmt.Print("Enter the length of rectangle: ")
        fmt.Scan(&length)
        fmt.Print("Enter the breadth of rectangle: ")
        fmt.Scan(&breadth)
        return length * breadth
}

func square_area() float32{
        var length float32
        fmt.Print("Enter the length of square: ")
        fmt.Scan(&length)
        return length * length 
}
func circle_area() float32{
	var radius float32
	fmt.Print("Enter the radius of circle: ")
	fmt.Scan(&radius)
	return 3.14 * radius * radius

}
func calculation() {
	
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	fmt.Print("********************************************************************\n")
	fmt.Print("You can calculate area of rectangle|circle|square using this script.\n")
	fmt.Print("********************************************************************\n")	
	var shape string	
	fmt.Print("Enter the shape for which you want to calculate the area: ")
	fmt.Scan(&shape)
	
	if shape == "rectangle" {
	
		fmt.Print("The area of rectangle is: ", rectangle_area(), "\n")
	
	} else if shape == "circle" {
		fmt.Print("The area of circle is: ", circle_area(), "\n")
	
	} else if shape == "square" {
		fmt.Print("The area of sqaure is: ", square_area(), "\n")
	} else {
		fmt.Print("Wrong shape entered!!\n")
	}
	return	
}

func main() {
	calculation()
	var user_input string
	fmt.Print("Do you want to continue? Y or N: ")
	fmt.Scan(&user_input)
	if user_input == "Y" {
		main()
	} else {
		fmt.Print("Script exited!\n")
	}

}
