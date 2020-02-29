package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// Person - структура
type Person struct {
	Phone int
	Name  string
	Money float64
}

func main() {

	csvFile, err := os.Open("peoples.csv")

	if err != nil {
		return
	}
	defer csvFile.Close()

	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = ';'

	var people []Person

	var onePerson Person

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		onePerson.Phone, _ = strconv.Atoi(line[0])
		onePerson.Name = line[1]
		onePerson.Money, _ = strconv.ParseFloat(line[2], 64)

		people = append(people, onePerson)

	}

	fmt.Println(people)

}

// https://medium.com/@barunthapa/working-with-csv-in-go-50a4f540e623
