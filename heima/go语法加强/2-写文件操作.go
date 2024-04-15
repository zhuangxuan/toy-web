package main

import (
	"io"
	"os"
)

func main() {
	f, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	//_, err = f.Write([]byte("hello world"))
	// 使用writeString

	// 修改读写位置
	offSet, err := f.Seek(0, io.SeekEnd)
	if err != nil {
		panic(err)
	}
	//n, err := f.WriteString("hello world fffssss")
	n, err := f.WriteAt([]byte("hello world"), offSet)
	if err != nil {
		panic(err)
	}
	println(" bytes written\n", n)
}
