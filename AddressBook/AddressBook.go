package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// tPhone - тип телефонного номера
type tPhone string

// tPerson - тип человек для телефонного справочника
type tPerson struct {
	FirstName   string // имя
	LastName    string // фамилия
	CompanyName string // дата рождения
}

// tAddressBook - тип машина
type tAddressBook map[tPhone]tPerson

// main - точка входа
func main() {

	addressBook := make(tAddressBook)

	addressBook[tPhone("7-900-0000-000")] = tPerson{
		FirstName:   "Bill",
		LastName:    "Gates",
		CompanyName: "Microsoft",
	}

	addressBook[tPhone("7-900-0000-001")] = tPerson{
		FirstName:   "Jeffrey",
		LastName:    "Bezos",
		CompanyName: "Amazon",
	}

	addressBook[tPhone("7-900-0000-002")] = tPerson{
		FirstName:   "Sergey",
		LastName:    "Brin",
		CompanyName: "Google",
	}

	fmt.Println(addressBook)
	fmt.Println()

	wFile, _ := json.Marshal(addressBook)
	//wFile, err := json.MarshalIndent(addressBook, "", " ")

	//fmt.Println("Error", err)

	fmt.Println(string(wFile))
	fmt.Println()

	_ = ioutil.WriteFile("AddressBook.json", wFile, 0644)

	// теперь прочитаем

	rFile, _ := ioutil.ReadFile("AddressBook.json")
	fmt.Println(string(rFile))
	fmt.Println()

	var newAddressBook tAddressBook

	_ = json.Unmarshal(rFile, &newAddressBook)

	fmt.Println(newAddressBook)
	fmt.Println()

}
