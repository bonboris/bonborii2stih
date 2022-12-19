package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("stih.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("can't close file", err)
			os.Exit(1)
		}
	}(file)

	data := make([]byte, 64)
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}

		var column []string
		for i := 0; i < n; i++ {
			column = append(column, string(data[i]))
			fmt.Println(column[i])
		}
	}
}