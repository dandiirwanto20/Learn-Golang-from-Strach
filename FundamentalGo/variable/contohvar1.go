package main

import "fmt"

func main() {
	// Deklarasi variabel dengan type data
	var name string = "Dandi"
	var age int = 22

	fmt.Println("Ini adalah namanya ==> ", name)
	fmt.Println("Ini adalah Umur nya ==> ", age)

	// Deklarasi variabel dengan type data dengan cara reassign

	var tinggi int
	tinggi = 170

	var namaKedua string
	namaKedua = "Dinda"

	tinggi = 171

	fmt.Println("Ini adalah tinggi ==>", tinggi)
	fmt.Println("Ini adalah Nama Kedua ==>", namaKedua)

	// Deklarasi variabel tanpa data type / type interface
	var namaPertama = "Dandi"
	var umur = 22

	fmt.Println("Nama :", namaPertama)
	fmt.Printf("%T, %T \n", namaPertama, umur)
	fmt.Println("Umur :", umur)

	// Deklarasi Variabel Short Declaration

	hewanPertama := "Kucing"

	fmt.Println("Nama Hewan:", hewanPertama)
	
	// Deklarasi Variabel Multiple

	var student1, student2, student3 string = "Dandi", "Irwanto", "Taiga"

	var first, second, third = 1, 2, 3

	fmt.Println(student1, student2, student3)
	fmt.Println(first, second, third)

	// Deklarasi Undescore Variabel

	var fisrtName string

	var nam, ag, address = "Dandi", 22, "Jalan Bogorejo"

	_, _, _, _ = fisrtName, nam, ag, address

	/*
	fmt.Prinf Usage

	flag %T digunakan untuk menampilkan type data

	jika langsung memasukan dalam kalimat bisa menggunakan beberapa flag:
	%s untuk data string
	%d untuk data int
	
	*/

}