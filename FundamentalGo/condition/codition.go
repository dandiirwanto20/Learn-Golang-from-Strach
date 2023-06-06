package main

import "fmt"

func main() {
	// terdapat 2 macan kondisional yakni if else dan switch

	// variabel temporer terbuat dari scope block
	var currentYear = 2021

	if age := currentYear - 1998; age < 17 {
		fmt.Println("Kamu belum boleh membuat kartu sim")
	} else {
		fmt.Println("Kamu sudah boleh membuat kartu sim")
	}

	// switch
	var score = 8

	switch score {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// switch with relasional variabel
	var score1 = 5

	switch {
	case score1 == 8:
		fmt.Println("perfect")
	case (score1 < 8) && (score1 > 3):
		fmt.Println("not bad")
	default:
		{
			fmt.Println("study harder")
			fmt.Println("you need to learn more")
		}
	}

	// switch (fallthrough keyword)
	var score2 = 6

	switch {
	case score2 == 8:
		fmt.Println("perfect")
	case (score2 < 8) && (score2 > 3):
		fmt.Println("not bad")
		fallthrough
	case score2 < 5:
		fmt.Println("It is ok, but please study harder")
	default:
		{
			fmt.Println("study harder")
			fmt.Println("You don't have a good score yet")
		}
	}

	// nested condition
	var nilai = 10

	if nilai > 7 {
		switch nilai {
		case 10:
			fmt.Println("perfet!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if nilai == 5 {
			fmt.Println("not bad")
		} else if nilai == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if nilai == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}