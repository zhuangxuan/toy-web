package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	f.WriteString("hello, world ddd")
	fmt.Println("done")
	defer f.Close()
	//	strings.HasPrefix(str, "h")

	file, err := os.OpenFile("test.txt", os.O_APPEND, 0666)
	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])

}
