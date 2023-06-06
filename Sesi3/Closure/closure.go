package main

import (
	"fmt"
	"strings"
)

type isOddNum func(int) bool // Alias untuk mempersingkat penulisan function callback(Closure)

func main () {
	// Deklarasi Closure dalam variabel
	var evenNumbers = func(numbers ...int) []int {
		var result []int

		for _, v := range numbers {
			if v % 2 == 0 {
				result = append(result, v)
			}
		}

		return result
	}

	var numbers = []int{4, 56, 43, 323, 22, 121, 34}

	fmt.Println(evenNumbers(numbers...))

	// Closure (IIFE) immediately-invoked-function-expression
	var isPalindrome = func(str string) bool {
		var temp string = ""

		for i := len(str) - 1; i >= 0; i-- {
			temp += string(byte(str[i]))
		}
		return temp == str
	}("katak")

	fmt.Println(isPalindrome)

	// closure dapat menjadi value return
	var studentList = []string{"Dandi", "Lia", "Agung", "Lai"}

	var find = findStudent(studentList)

	fmt.Println(find("dandi"))


	// Closure Callback
	var angka = []int{2, 5, 8, 10, 3, 99, 23}

	var find1 = findOddNumbers(angka, func (number int) bool {
		return number%2 != 0
	})

	fmt.Println("Total odd numbers:", find1)

}

// closure as a return value
func findStudent(students []string) func(string) string {

	return func(s string) string {
		var student string
		var position int

		for i, v := range students {
			if strings.ToLower(v) == strings.ToLower(s) {
				student = v
				position = i
				break
			}
		}
		if student == "" {
			return fmt.Sprintf("%s doesnt exist!!!", s)
		}
		return fmt.Sprintf("We found %s at positon %d", s, position+1)
	}
}

// func Closure untuk callback
func findOddNumbers(numbers []int, callback isOddNum) int {
	var totalOddNumbers int

	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}