package main

import (
	"fmt"
	"strings"
)

func main() {
	// slice seperti array namun memiliki perbedaan di tipe datanya. slice tidak memiliki sifat fixed-length yang berarti panjang dari
	// slice tidak tetap sehingga kita bisa dengan leluasa menentukan panjang dari slicenya
	// slice masuk dalam reference type yang dimana jika kita melakukan copy terhadap suatu slice, dan kita ubah element dari yang kita copy tersebut
	// maka slice semula akan ikut berubah
	var fruits = []string{"apple", "banana", "mango"}

	_ = fruits

	// fmt.Printf("%#v", fruits)

	// slice fungsi make
	var buah = make([]string, 3)
	
	_ = buah
	
	// untuk memasukan nilai pada slice kosong
	buah[0] = "Jeruk Bali"
	buah[1] = "Banana"
	buah[2] = "Sirsak"

	fmt.Printf("%#v\n", buah)

	// cara kedua adalah dengan menggunakan append
	var fruit = make([]string, 3)

	fruit = append(fruit, "apple", "banana", "mango") //param pertama merupakan slice, kedua merupakan elemen

	fmt.Printf("%#v\n", fruit)

	// Slice append function with ellipsis atau (...)
	var fruits1 = []string{"apple", "banana", "mango"}
	
	var fruits2 = []string{"durian", "pineapple", "starfruit"}

	var buah1 = append(fruits1, fruits2...)

	fmt.Printf("%#v\n", buah1)

	// Slice copy function
	nn := copy(fruits1, fruits2) // dalam function copy ini variabel fruits2 yang dicopy dan direplace ke fruits1

	fmt.Println("Fruits1 =>", fruits1)
	fmt.Println("Fruits2 =>", fruits2)
	fmt.Println("Copied Element =>", nn)

	// Slicing Slice
	var fruits3 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	var fruits4 = fruits3[1:4] //[start:stop]
	fmt.Printf("%#v\n", fruits4)

	var fruits5 = fruits3[0:]
	fmt.Printf("%#v\n", fruits5)

	var fruits6 = fruits3[:3]
	fmt.Printf("%#v\n", fruits6)

	var fruits7 = fruits3[:] // sama dengan fruits3[:len(fruits3)]
	fmt.Printf("%#v\n", fruits7)

	// Slice Combining slicing and append
	fruits3 = append(fruits3[:3], "rambutan")
	fmt.Printf("%#v\n", fruits3)

	// Backing array
	var buah2 = []string{"apple", "mango", "durian", "banana", "starfruit"}

	var buah3 = buah2[2:4]

	buah3[0] = "rambutan"

	fmt.Println("fruits1 =>", buah2)
	fmt.Println("fruits1 =>", buah3)

	// Slice cap function untuk mengetahui kapasitas dari sebuah array maupun slice
	var buahan = []string{"apple", "mango", "durian", "banana"}

	fmt.Println("Fruits1 cap:", cap(buahan))
	fmt.Println("Fruits1 len:", len(buahan))

	fmt.Println(strings.Repeat("#", 20))

	var buahan1 = buahan[0:3]

	fmt.Println("Buahan cap:", cap(buahan1))
	fmt.Println("Buahan len:", len(buahan1))

	fmt.Println(strings.Repeat("#", 20))

	var buahan2 = buahan[1:]

	fmt.Println("Buahan cap:", cap(buahan2))
	fmt.Println("Buahan len:", len(buahan2))

	// Slice creating a new backing array
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"
	fmt.Println("cars:", cars)
	fmt.Println("newCars:", newCars)
}