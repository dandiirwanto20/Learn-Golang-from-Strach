package main

import (
	"fmt"
	"os"
	"strconv"

	"challange/helpers"
)

func main() {
	
	args := os.Args
	if len(args) < 2 {
		panic("Error: Argumen harus berupa nomor teman")
	}

	arg := os.Args[1]

	no, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("Argumen harus berupa angka")
		return
	}

	friendList := helpers.GetFriendList()
	helpers.ShowFriendData(friendList, no)
}