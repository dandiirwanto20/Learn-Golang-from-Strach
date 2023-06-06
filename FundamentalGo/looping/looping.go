package main

import "fmt"

func main() {
	// perulangan (first way of looping)
	for i := 0; i < 3; i++ {
		fmt.Println("Angka", i)
	}

	// perulangan (second way of looping) memasukan kondisional menggunakan while
	var a = 0

	for a < 3 {
		fmt.Println("Angka", a)
		a++
	}

	// Perulangan (third way of looping) tanpa memberi kondisi apapun
	var b = 0

	for {
		fmt.Println("Angka", b)

		b++
		if b == 3 {
			break
		}
	}

	// Perulangan break and continue keywords
	for c := 1; c <= 10; c++ {
		if c%2 == 1 {
			continue
		}

		if c > 8 {
			break
		}

		fmt.Println("Angka", c)
	}

	// Perulangan (nested looping)
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

	// Perulangan (Label)
	outerLoop:
	for i := 0; i < 3; i++ {
		fmt.Println("Perulangan ke - ", i+1)
		for j := 0; j < 3; j++ {
			if i == 2 {
				break outerLoop
			}
			fmt.Print(j, " ")
		}
		fmt.Print("\n")
	}
}