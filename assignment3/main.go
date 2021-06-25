package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1]) // For read access.
	if err == nil {
		io.Copy(os.Stdout, file)
	}
}
