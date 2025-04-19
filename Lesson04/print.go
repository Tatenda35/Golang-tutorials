package main

import "fmt"

func main(){
	// Print is like Println but does not print a new line
	nameOne := "Tatenda"
	age := 30
	height := 1.567
	fmt.Print("Hello, ", nameOne,".\n")

	// Printf is like in c programming where we can format the string and does not print a new line
	fmt.Printf("Hello, %s. How does it feel to turn %d? And is it true that you are %.1fm tall.\n", nameOne, age, height)

	// Println automatically prints a new line
	// But an extra space is inserted after the variable
	// So best to use Printf
	fmt.Println("Hello,", nameOne,". How does it feel to turn,", age,"? And is it true that you are ", height, "m tall?")
}