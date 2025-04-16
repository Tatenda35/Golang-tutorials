package main

import "fmt"

func main(){
	// sttrings
	var nameOne string = "Tatenda"
	// Or, this will mean even in future this remains a string
	var nameTwo = "David"
	// Or declare for later use, but it will get a default empty value we can't see
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	// Now we update, and can only be updated as string

	nameOne = "Ishe"
	nameTwo = "Mark"
	nameThree = "Matthew"

	fmt.Println(nameOne, nameTwo, nameThree)

	// Initializing variable with shortkey, 
	// Later when you update the variable there is no need to use colon.
	// Colon shortkey can only be used in local variable
	nameFour := "John"
	fmt.Println(nameFour)
}