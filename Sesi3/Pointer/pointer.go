package main

import (
	"fmt"
	"strings"
)

func main() {
	// var firstNumber *int
	// var secondNumber *int

	// _,_ = firstNumber, secondNumber

	// Pointer Memory Address

	/*
		Kita dapat mendapatkan alamat memory dari variable biasa dengan
		menggunakan tanda ampersand &. Kita juga dapat mendapatkan nilai
		asli dari variable pointer dengan cara menggunakan tanda asterisk *
	*/
	var firstNumber int = 4
	var secondNumber *int = &firstNumber
	fmt.Println("firstNumber (value) :", firstNumber)
	fmt.Println("firstNumber (memori address) :", &firstNumber) //tanda ampersand &

	fmt.Println("secondNumber (value) :", *secondNumber)
	fmt.Println("secondNumber (memori address) :", &secondNumber)

	// Pointer (Changing value through pointer)
	var firstPerson string = "John"
	var secondPerson *string = &firstPerson

	fmt.Println("firstPerson (value) :", firstPerson)
	fmt.Println("firstPerson (memori address) :", &firstPerson)
	fmt.Println("secondPerson (value) :", *secondPerson)
	fmt.Println("secondPerson (memori address) :", secondPerson)

	fmt.Print("\n", strings.Repeat("#", 30), "\n")

	*secondPerson = "Doe" //change value

	fmt.Println("firstPerson (value) :", firstPerson)
	fmt.Println("firstPerson (memori address) :", &firstPerson)
	fmt.Println("secondPerson (value) :", *secondPerson)
	fmt.Println("secondPerson (memori address) :", secondPerson)

	// Pointer as Parameter
	var a int = 10

	fmt.Println("Before :", a)

	changeValue(&a)

	fmt.Println("After :", a)

}

func changeValue(number *int) {
	*number = 20
}
