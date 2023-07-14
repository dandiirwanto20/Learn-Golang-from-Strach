package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

// menghitung luas lingkaran
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// menghitung luas persegi panjang
func (r rectangle) area() float64 {
	return r.height * r.width
}

// menghitung keliling lingkaran
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// menghitung keliling persegi panjang
func (r rectangle) perimeter() float64 {
	return 2 * (r.height * r.width)
}

func print(t string, s shape) {
	fmt.Printf("%s area %v\n", t, s.area())
	fmt.Printf("%s perimeter %v\n", t, s.perimeter())
}

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}


func main() {
	// implementasi type data interface (shape)
	var c1 shape = circle{radius: 5}
	var r1 shape = rectangle{width: 3, height: 2}

	fmt.Printf("Type of c1: %T\n", c1)
	fmt.Printf("Type of r1: %T\n", r1)

	// fmt.Println("Circle area", c1.area())
	// fmt.Println("Circle perimeter", c1.perimeter())

	// fmt.Println("Rectangle area", r1.area())
	// fmt.Println("Rectangle perimeter", r1.perimeter())
	print("Circle", c1)
	print("Rectangle", r1)

	// interface type assertion
	var c12 shape = circle{radius: 5}

	// c12.volume() tidak dapat dideklarasikan seperti ini karena func volume() tidak dideklarasikan pada interface shape

	// sehingga perlu dideklarasikan dengan teknik type assertion
	// c12.(circle).volume() // formatnya namaVariabel.(type data struct).nama interface diluar dari deklarasi interface

	// dapat memeriksa apakah sebuah type assertion yang digunakan berhasil atau tidak dengan cara:
	value, ok := c12.(circle)

	if ok == true {
		fmt.Printf("Circle value: %+v\n", value)
		fmt.Printf("Circle volume: %f", value.volume())
	}

}
