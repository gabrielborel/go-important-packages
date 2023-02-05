package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// WRITE
	fileToWrite, err := os.Create("file.txt")
	defer fileToWrite.Close()
	if err != nil {
		panic(err)
	}

	size, err := fileToWrite.WriteString("Hello World! Golang is very nice to learn.")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Size of the file: %d bytes\n\n", size)

	// READ
	fileToRead, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(fileToRead))

	// CHUNK BY CHUNK READING
	fileToReadChunkByChunk, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(fileToReadChunkByChunk)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}

	// REMOVE
	err = os.Remove("file.txt")
	if err != nil {
		panic(err)
	}
}
