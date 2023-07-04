package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

// type Employee struct {
// 	division string
// 	person   Person
// }

// type Employee struct {
// 	name     string
// 	age      int
// 	division string
// }

func main() {
	// Struct (Giving value to struct)

	// var employee Employee

	// employee.name = "Dandi"

	// employee.age = 23

	// employee.division = "IT Developer"

	// fmt.Println(employee.name)
	// fmt.Println(employee.age)
	// fmt.Println(employee.division)

	// Initializing struct

	// var employee1 = Employee{}
	// employee1.name = "Dandi"
	// employee1.age = 23
	// employee1.division = "IT Dev"

	// var employee2 = Employee{name: "Anda", age: 40, division: "Case"}

	// fmt.Printf("Employee1: %+v\n", employee1)
	// fmt.Printf("Employee2: %+v\n", employee2)

	// Pointer to a struct

	// var employee1 = Employee{name: "Dandi", age: 22, division: "Developer"}

	// var employee2 *Employee = &employee1

	// fmt.Println("Employee1 name:", employee1.name)
	// fmt.Println("Employee2 name:", employee2.name)

	// fmt.Println(strings.Repeat("#", 20))

	// employee2.name = "Ananda"

	// fmt.Println("Employee1 name:", employee1.name)
	// fmt.Println("Employee2 name:", employee2.name)

	// Embedded struct
	// var employee1 = Employee{}

	// employee1.person.name = "Dandi"
	// employee1.person.age = 23
	// employee1.division = "IT Dev"

	// fmt.Printf("%+v", employee1)

	// Anonymous struct tanpa pengisian field

	// var employee1 = struct {
	// 	person   Person
	// 	division string
	// }{}

	// employee1.person = Person{name: "Dandi", age: 23}
	// employee1.division = "IT Dev"

	// Anonymous struct dengan pengisian field

	// var employee2 = struct {
	// 	person   Person
	// 	division string
	// }{
	// 	person:   Person{name: "Dian Sastro", age: 23},
	// 	division: "Finance",
	// }

	// fmt.Printf("Employee1: %+v\n", employee1)
	// fmt.Printf("Employee2: %+v\n", employee2)

	// Slice of struct

	// var people = []Person{
	// 	{name: "Dand", age: 20},
	// 	{name: "Irwan", age: 21},
	// 	{name: "Dev", age: 12},
	// }

	// for _, v := range people {
	// 	fmt.Printf("%+v\n", v)
	// }

	// Slice of anonymous struct dengan field tidak di isi
	var employee = []struct {
		person   Person
		division string
	}{
		{person: Person{name: "Dandi", age: 23}, division: "IT Dev"},
		{person: Person{name: "Dian", age: 23}, division: "Finance"},
		{person: Person{name: "Mailo", age: 21}, division: "Marketing"},
	}

	for _, v := range employee {
		fmt.Printf("%+v\n", v)
	}

}
