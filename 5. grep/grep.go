package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// GrepString - поиск в файле и вывод на экран строк в которых есть pattern
func GrepString(fileName string, pattern string) {
	file, err := os.Open(fileName)

	if err != nil {
		return
	}
	defer file.Close()

	// читаем построчно
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// ищем есть ли нужный pattern, если есть, выводим
		if strings.Contains(scanner.Text(), pattern) {
			fmt.Println(scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

var (
	file    = flag.String("file", "", "File")
	pattern = flag.String("pattern", "", "Pattern")
)

// go run grep.go --file=lorem.txt --pattern=ipsum

func main() {

	flag.Parse()

	GrepString(*file, *pattern)
}
