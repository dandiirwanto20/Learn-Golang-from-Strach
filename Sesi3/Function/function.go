package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// greet("Dandi", 23)
	// println()
	// ucapan("Dandi", "Rembang")

	var names = []string{"Dandi", "Irwanto"}
	var printMsg = greeting("Heii", names)

	fmt.Println(printMsg)

	fmt.Println()

	var diameter float64 = 15

	var area, circumference = calculate(diameter)

	fmt.Println("Area:", area)
	fmt.Println("Circumference:", circumference)

	// func lingkaran untuk predefined
	var diameter1 float64 = 15

	var luas, keliling = lingkaran(diameter1)

	fmt.Println("Luas:", luas)
	fmt.Println("Keliling:", keliling)

	// Func Variadic Map
	studentList := print("Dan", "Nada", "Alex", "Lia", "Mei")

	fmt.Printf("%v", studentList)

	fmt.Println()

	// Func Variadic slice
	numberList := []int{1, 2, 3, 4, 5, 6, 7, 8}

	result := sum(numberList...)

	fmt.Println("Result:", result)

	// Variadic 3 func merge basic params with variadic params
	profile("Dandi", "Geprek", "Nasi Goreng", "Sate", "Mie Ayam")

}

// func greet(name string, age int8) {
// 	fmt.Printf("Hello there! My Name is %s and I'm %d years old", name, age)
// }

// // ketika tipe data dalam params sama maka tidak perlu dideklarasikan ulang
// func ucapan(nama, alamat string) {
// 	fmt.Printf("Halo perkenalkan saya %s dan saya tinggal di %s", nama, alamat)
// }

// function yang menggunakan return
func greeting(msg string, names []string) string {
	var joinStr = strings.Join(names, " ") // menggunakan Function Join untuk menghubungkan string di array maupun slice

	var result string = fmt.Sprintf("%s %s", msg, joinStr) // menggunakan Sprintf karena bisa mereturn nilai karena menggunakan printf tidak bisa

	return result
}

// function mengembalikan banyak value (return multiple values)
func calculate(d float64) (float64, float64) {
	// Menghitung luas
	var area float64 = math.Pi * math.Pow(d/2, 2)

	// Menghitung keliling
	var circumference = math.Pi * d

	return area, circumference
}

// Predefined return value
func lingkaran(dia float64) (luas float64, keliling float64) {

	// Menghitung Luas
	luas = math.Pi * math.Pow(dia/2, 2)

	// Menghitung keliling
	keliling = math.Pi * dia

	return
}

// Variadic Function (...) menerima argument yang tak terbatas dengan menggunakan map
func print(names ...string) []map[string]string {
	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("student%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}

	return result
}

// Variadic Function dengan menggunakan slice datatype
func sum(numbers ...int) int {
	total := 0

	for _, v := range numbers {
		total += v
	}
	return total
}

// Variadic Func menggabungkan params biasa dengan params variadic
func profile(name1 string, favFoods ...string) {
	mergeFavFoods := strings.Join(favFoods, ",")

	fmt.Println("Hello there!!! I'm", name1)
	fmt.Println("I really love to eat", mergeFavFoods)
}
