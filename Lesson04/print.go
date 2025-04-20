package main

import "fmt"

func main(){
	// Print is like Println but does not print a new line
	nameOne := "Tatenda"
	age := 30
	height := 1.567
	fmt.Print("Hello, ", nameOne,".\n")

	// Println automatically prints a new line
	// But an extra space is inserted after the variable
	// So best to use Printf
	fmt.Println("Hello,", nameOne,". How does it feel to turn,", age,"? And is it true that you are ", height, "m tall?")

	// Printf is like in c programming where we can format the string and does not print a new line
	// using %v for variable known as format specifier to represent where the variable will go
	// or you can use %s for string, %d for int, %f for float
	// %q puts qoutes
	fmt.Printf("Hello, %v. How does it feel to turn %v? And is it true that you are %.1fm tall.\n", nameOne, age, height)
	qoute := "By his stripes we were healed."
	fmt.Printf("My favaourate qoute is, %q \n", qoute)

	// getting the variable type
	fmt.Printf("age is of type %T \n", age)

	// printing floats
	// percent sign is double %%
	grade := 99.55
	fmt.Printf("You scored %.1f%% \n", grade)
}