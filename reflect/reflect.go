package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	// reflect.ValueOf() -> untuk check informasi yang berhubungan dengan nilai pada variabel
	// reflect.ValueOf() -> untuk check informasi yang berhubungan dengan type data variabel yang dicari
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("tipe variabel :", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int { // mememeriksaterlebih dahulu apa jenis tipe datanya menggunakanmethod Kind()
		fmt.Println("nilai Variabel :", reflectValue.Int())
	}

	fmt.Println("tipe variabel :", reflectValue.Type())
	// Accessing Value Using Interface Method -> Jika nilai hanya diperlukan untuk ditampilkan ke output,
	fmt.Println("nilai variabel :", reflectValue.Interface())

	/*
		Accessing Value Using Interface Method
		Method Interface() mengembalikan nilai interface kosong atau interface{}.
		Nilai aslinya sendiri bisa diakses dengan meng-casting interface kosong tersebut.
		example: var nilai = reflectValue.Interface().(int)
	*/

	/*
		Identifying Method Information
		Informasi mengenai method juga bisa diakses lewat reflect,
		syaratnya masih sama seperti padapengaksesan property, yaitu harus bermodifier public.

		Pada contoh dibawah ini informasi method SetName akan diambil lewat reflection.

		Siapkan method baru di struct student, dengan nama SetName.

		func (s *student) SetName(name string) {
			s.Name = name
		}
	*/

	// Identifying Method Information penerapan di main
	var s1 = &student{Name: "Dandi Irwanto", Grade: 2}
	fmt.Println("nama :", s1.Name)

	var reflectValue1 = reflect.ValueOf(s1)

	var method = reflectValue1.MethodByName("SetName")
	method.Call([]reflect.Value {
		reflect.ValueOf("Irwanto"),
	})

	fmt.Println("nama :", s1.Name)
}
