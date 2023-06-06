package main

import "fmt"

func main() {
	// map seperti array dan slice namun bedanya datanya disimpan dalam key:value pair
	// key dalam map haruslah unik sedangkan value tidak mesti
	// termasuk tipe data reference type seperti tipe data slice
	// zero value adalah nil
	

	// var person map[string]string // deklarasi
	// person = map[string]string{} // inisialisasi

	// person["name"] = "Dandi"

	// person["age"] = "22"

	// person["address"] = "Jalan Sudirman"

	// fmt.Println("name:", person["name"])
	// fmt.Println("age:", person["age"])
	// fmt.Println("address:", person["address"])


	// Dapat diberikan value beserta dengan keynya dengan langsung
	// var person1 = map[string]string{
	// 	"nama": "Dondo",
	// 	"umur": "22",
	// 	"alamat": "Rembang",
	// }
	
	// fmt.Println("nama:", person1["nama"])
	// fmt.Println("umur:", person1["umur"])
	// fmt.Println("alamat:", person1["alamat"])


	// // Looping with map
	// // for key, value := range person1 {
	// // 	fmt.Println(key, ":", value)
	// // }


	// // map Deleting value
	// fmt.Println("Before Deleting:", person1)

	// delete(person1, "umur")

	// fmt.Println("After Deleting:", person1)


	// Map detecting a value
	// var person = map[string]string{
	// 	"name":		"Dandi",
	// 	"age":		"22",
	// 	"address":	"Rembang",
	// }

	// value, exist := person["age"]

	// if exist {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("Value does'nt exist")
	// }

	// delete(person, "age")

	// value, exist = person["age"]

	// if exist {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("Value has been deleted")
	// }

	// Combining slice with map
	var people = []map[string]string{
		{"name": "Dandi", "age": "22"},
		{"name": "Irwanto", "age": "20"},
		{"name": "Mei", "age": "20"},
	}

	for i, person := range people {
		fmt.Printf("Index: %d, name: %s, age: %s\n", i, person["name"], person["age"]) // i = index dan person adalah key = value
	}
}