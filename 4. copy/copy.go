package main

import (
	"flag"
	"io"
	"os"
)

// Copy - копирование файла
func Copy(src, dst string) error {
	in, err := os.Open(src)

	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return out.Close()
}

var (
	from = flag.String("from", "", "From")
	to   = flag.String("to", "", "To")
)

// go run copy.go --from=lorem.txt --to=lorem2.txt

func main() {
	flag.Parse()

	Copy(*from, *to)
}
