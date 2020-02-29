package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// ReadFileNotOK - функция чтения из файла которая не очень хороша
func ReadFileNotOK(fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		return
	}
	defer file.Close()

	// getting size of file
	stat, err := file.Stat()
	if err != nil {
		return
	}

	// reading file
	// V тут читаем весь файл разом независимо от размера, не очень хорошо
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	fmt.Print(string(bs))
}

// ReadFileOK - функция чтения из файла получше
func ReadFileOK(fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		return
	}
	defer file.Close()

	// читаем построчно
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	fileName := "lorem.txt"

	ReadFileNotOK(fileName)

	fmt.Println("\n ------------------------- ")

	ReadFileOK(fileName)

}
