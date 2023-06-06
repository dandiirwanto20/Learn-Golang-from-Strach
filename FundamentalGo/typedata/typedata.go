package main

import "fmt"

func main() {
	/*
		Tipe data Go terbagi menjadi 4 kategori dengan detail:
		1. Basic type: number, string, boolean
		2. Aggregate type: array dan struct
		3. Reference type: slice, pointer, map, function, channel
		4. Interface type: interface
	*/

	// Number non desimal/non floating menggunakan integer (int) sedangkan desimal/flioating menggunakan float
	// integer (int)
	var first = 89
	var second = -12

	fmt.Printf("type data first: %T\n", first)
	fmt.Printf("type data second: %T\n", second)

	// Untuk menghindari boros memory maka kita dapat mendeklarasikan type datanya
	var firsttype uint8 = 127
	var secondtype int8 = -12
	
	fmt.Printf("type data first: %T\n", firsttype)
	fmt.Printf("type data second: %T\n", secondtype)

	// Number (float) float32 dan float64
	var decimalNumber float32 = 3.63

	fmt.Printf("type data first: %f\n", decimalNumber) //pada dasarnya flag %f akan memformat angka desimal menjadi 6 digit
	fmt.Printf("type data second: %.3f\n", decimalNumber) //sedangkan pada flag %.nf akan mengecilkan angka desimal sesuai custom

	// Boolean (bool)
	var condition bool = true

	fmt.Printf("Is it permited? %t \n", condition)

	// String
	// Menggunakan Double Quote
	var message = "Hallo"

	fmt.Println(message)

	// Menggunakan backtics `` atau grave accent

	var messageDua = `Selamat datang di "Hacktiv8",
	salam kenal :).
	Mari Belajar "Scalable Web Service With Go".`

	fmt.Println(messageDua)

	// nil dan zero value
	/*
		nil bukan merupakan type data melainkan sebuah nilai, merupakan variabel yang isinya nilanya nil berarti memiliki nilai kosong

		 Semua tipe data memiliki zero value (nilai default tipe data)
		 Zero value dari string adalah "" (string kosong).
		 ●Zero value dari bool adalah false.
		 ●Zero value dari tipe numerik non-desimal adalah 0.
		 ●Zero value dari tipe numerik desimal adalah 0.0.
		 
		 Zero value berbeda dengan nil. Nil adalah nilai kosong, benar-benar kosong. 
		 Nil tidak bisa digunakan pada tipe data yang sudah dibahas sebelumnya. 
		 
		 Ada beberapa tipe data yang bisa di-set nilainya dengan nil, diantaranya:
		 ● pointer
		 ● tipe data function
		 ● slice
		 ● map
		 ● channel
		 ● interface kosong atau interface{}
	*/
	
}
