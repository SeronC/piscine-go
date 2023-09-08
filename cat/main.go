package main

import (
	"fmt"
	"io"
	"os"
)

func catFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("ERROR: open %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	io.Copy(os.Stdout, file)
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	for _, filename := range args {
		catFile(filename)
	}
}
