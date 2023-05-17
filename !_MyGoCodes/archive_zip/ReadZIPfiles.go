package main

import (
	"archive/zip"
	"fmt"
	"log"
	"unicode/utf8"
    "github.com/axgle/mahonia"
)

func main() {
	// 打开一个zip文件供阅读。
	r, err := zip.OpenReader("Day86.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// 遍历存档中的文件，
	// 打印他们的一些内容。
	for _, f := range r.File {
		data := ConvertToString(str, "gbk", "utf-8")
		fmt.Printf("Contents of %s:\n", f.Name)
	}
}
