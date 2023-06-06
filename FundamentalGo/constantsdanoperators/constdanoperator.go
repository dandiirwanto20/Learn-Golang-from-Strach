package main

import "fmt"

func main() {
	// constant digunakan untuk mendeklarasikan variabel dengan nilai tetap dan tidak dapat diubah atau diinilisasi ulang
	// const full_name string = "Dandi Irwanto"

	const full_name = "Dandi"

	fmt.Printf("Hello %s\n", full_name)

	// Operator

	// Aritmatika
	var value = 2 + 2*3
	fmt.Println(value)

	// Relasi
	var firstCondition bool = 2 < 3
	var secondCondition bool = "joey" == "Joey"
	var thirdCondition bool = 10 != 2.3
	var fourthCondition bool = 11 <= 11

	fmt.Println("first condition:", firstCondition) // kondisi dimana akan bernilai true karena nilai 2 kurang dari 3
	fmt.Println("second condition:", secondCondition) // menghasilkan false karena tidak sama
	fmt.Println("third condition:", thirdCondition) // true karena tidak sama dengan
	fmt.Println("fourth condition:", fourthCondition) // true karena kurang dari sama dengan

	//  Logika
	var right = true
	var wrong = false

	var wrongAndRight = wrong && right
	fmt.Printf("wrong && right \t(%t) \n", wrongAndRight)

	var wrongOrRight = wrong || right
	fmt.Printf("worng || right \t(%t) \n", wrongOrRight)

	var wrongReverse = !wrong
	fmt.Printf("!wrong \t\t(%t) \n", wrongReverse)
}