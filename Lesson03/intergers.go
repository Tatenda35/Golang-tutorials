// intergers
// Same as with strings
package main

import "fmt"

func main(){
	var ageOne int = 10
	var ageTwo = 20
	var ageThree int
	ageFour := 30

	fmt.Println(ageOne, ageTwo, ageThree, ageFour)

	// updating
	ageOne = ageOne + 5
	ageTwo = ageTwo + 5
	ageThree = 5
	ageFour = ageFour + 5	

	fmt.Println(ageOne, ageTwo, ageThree, ageFour)

	// bits and memory: integer type 8bits, 16bits etc
	var numOne int8 = 127
	var wordOne string = "Tatenda"
	var length float32 = 1.5
	decimalOne := 2.543
	fmt.Println(numOne)
	fmt.Printf("%s, got a gpa of %.1f, though he knows to count to %d, not to mention that he is %.2f tall!\n", wordOne, decimalOne, numOne, length)

	// gpa()
}


// floats

// func gpa(){
// 	var studentOne floats = 2.5
// 	var studentTwo = 1.5
// 	var studentThree floats
// 	studentFour := 0.5

// 	fmt.Println(studentOne, studentTwo, studentThree, studentFour)

// }