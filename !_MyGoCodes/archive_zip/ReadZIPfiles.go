package main

import (
	"archive/zip"
	"fmt"
	"log"
	"unicode/utf8"

	"github.com/fragmentization/mahonia"
)

func ConvertToString(src string, srcCode string, tagCode string) string {
	if utf8.Valid([]byte(src)) { //判断src是否本身就是utf8
		return src
	}
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func main() {
	// 打开一个zip文件供阅读。
	r, err := zip.OpenReader("Day86.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	//中文文件名GBK乱码转UTF-8
	for _, f := range r.File {
		f_Name := ConvertToString(f.Name, "gbk", "utf-8")
		fmt.Printf("Contents of %s:\n", f_Name)
	}
}
