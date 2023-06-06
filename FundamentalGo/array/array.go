package main

import (
	"fmt"
 	"strings"
)
func main() {
	// array go bersifat fixed-length atau memiliki panjang yang tetap yang harus kita tentukan dari awal kita membuat array
	// var numbers [4]int
	// numbers = [4]int{1, 2, 3, 4}

	// var strings = [3]string{"Dandi", "Irwanto", "Super"}

	// fmt.Printf("%#v\n", numbers) // %#v untuk menampilkan type array
	// fmt.Printf("%#v\n", strings)

	// array (Modify element through index)
	var fruits = [3]string{"apel", "pisang", "mangga"}
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "mango"

	fmt.Printf("%#v\n", fruits)

	// array (loop through elemen)
	// cara pertama dengan menggunakan range loop
	for i, v := range fruits {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}

	fmt.Println(strings.Repeat("#", 25))

	// cara kedua menggunakan loop biasa dengan len() 
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Index: %d, Value: %s\n", i, fruits[i])
	}

	// array multidimensional
	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}

}
