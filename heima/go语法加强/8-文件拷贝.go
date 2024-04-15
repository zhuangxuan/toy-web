package main

import (
	"fmt"
	"os"
)

func main() {
	srcFile := "简历.pdf.zip"
	dstFile := "简历.pdf2.zip"
	written, err := CopyFile(srcFile, dstFile)
	if err != nil {
		fmt.Println("copy file failed, err:", err)
		return
	}
	fmt.Println("copy done, total bytes:", written)
}

// 将上面的方法封装成函数
func CopyFile(srcFile, dstFile string) (written int64, err error) {
	f_r, err := os.Open(srcFile)
	if err != nil {
		panic(err)
	}
	defer f_r.Close()
	f_w, err := os.Create(dstFile)
	defer f_w.Close()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 1024)
	for {
		n, err := f_r.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		// 读多少，写多少
		f_w.Write(buf[:n])
		written += int64(n)
	}
	return
}
