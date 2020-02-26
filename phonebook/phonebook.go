package main

import (
	"fmt"
	"sort"
)

// Person - тип человек для телефонного справочника
type Person struct {
	Phone     string // телефон
	FirstName string // имя
	LastName  string // фамилия
}

// PhoneBook - тип телефонная книга
type PhoneBook []Person

func (p PhoneBook) Len() int {
	return len(p)
}

func (p PhoneBook) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PhoneBook) Less(i, j int) bool {
	return p[i].Phone < p[j].Phone
}

// main - точка входа
func main() {

	var phoneBook PhoneBook

	phoneBook = append(phoneBook, Person{
		Phone:     "7-900-0000-000",
		FirstName: "Bill",
		LastName:  "Gates",
	})

	phoneBook = append(phoneBook, Person{
		Phone:     "7-900-0000-002",
		FirstName: "Jeffrey",
		LastName:  "Bezos",
	})

	phoneBook = append(phoneBook, Person{
		Phone:     "7-900-0000-001",
		FirstName: "Sergey",
		LastName:  "Brin",
	})

	fmt.Println(phoneBook)

	sort.Sort(phoneBook)

	fmt.Println(phoneBook)
}
