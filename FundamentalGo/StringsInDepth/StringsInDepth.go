package main

import (
	"fmt"
	// "unicode/utf8"
)

func main() {
	// Looping Over String (byte-by-byte)
	
	// city := "Jakarta"

	// for i := 0; i < len(city); i++ {
	// 	fmt.Printf("index: %d, byte: %d\n", i, city[i])
	// }

	// var city []byte = []byte{74, 97, 107, 97, 114, 116, 97}

	// fmt.Println(string(city))

	// rune-by-rune
	str1 := "makan"
	_ = str1
	str2 := "mânca"

	// fmt.Printf("str1 byte length => %d\n", len(str1))
	// fmt.Printf("str2 byte length => %d\n", len(str2)) // dibaca 6 karena bahasa romania tidak mendukung ASCII accented-character "â" memiliki 2 byte

	// Ketika kita hendak mencari jumlah karakter nya, dan bukan jumlah bytenya, maka kita perlu merubah string tersebutmenjadi rune terlebih dahulu.
	// Tipe data rune merupakan tipe data alias dari int32

	// function RuneCountInString dari package utf8 untuk merubah string menjadi rune sekaligusmencari panjang karakternya.
	// fmt.Printf("str1 byte length => %d\n", utf8.RuneCountInString(str1))
	// fmt.Printf("str2 byte length => %d\n", utf8.RuneCountInString(str2))

	// Looping Over String (rune-by-rune)
	for i, s := range str2 {
		fmt.Printf("index => %d, string => %s\n", i, string(s)) // dibaca dari index 0 dan dilompati index 2 karena â berjumlah 2 byte maka 0, 1, 3 dst
	}
}