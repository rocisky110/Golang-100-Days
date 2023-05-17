package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		strings包下的关于字符串的函数

	*/
	s1 := "helloworld"
	//1.是否包含指定的内容-->bool
	fmt.Println(strings.Contains(s1, "abc"))
	//2.是否包含chars中任意的一个字符即可
	fmt.Println(strings.ContainsAny(s1, "abcd"))
	//3.统计substr在s中出现的次数
	fmt.Println(strings.Count(s1, "lloo"))

	//4.以xxx前缀开头，以xxx后缀结尾
	s2 := "20190525课堂笔记.txt"
	if strings.HasPrefix(s2, "201905") {
		fmt.Println("19年5月的文件。。")
	}
	if strings.HasSuffix(s2, ".txt") {
		fmt.Println("文本文档。。")
	}

	//索引
	//helloworld
	fmt.Println(strings.Index(s1, "lloo"))     //查找substr在s中的位置，如果不存在就返回-1
	fmt.Println(strings.IndexAny(s1, "abcdh")) //查找chars中任意的一个字符，出现在s中的位置
	fmt.Println(strings.LastIndex(s1, "l"))    //查找substr在s中最后一次出现的位置

	//字符串的拼接
	ss1 := []string{"abc", "world", "hello", "ruby"}
	s3 := strings.Join(ss1, "-")
	fmt.Println(s3)

	//切割
	s4 := "123,4563,aaa,49595,45"
	ss2 := strings.Split(s4, ",")
	//fmt.Println(ss2)
	for i := 0; i < len(ss2); i++ {
		fmt.Println(ss2[i])
	}

	// range 在迭代过程中返回的是迭代值的拷贝，如果每次迭代的元素的内存占用很低，
	// 那么 for 和 range 的性能几乎是一样，例如 []int。但是如果迭代的元素内存占
	// 用较高，例如一个包含很多属性的 struct 结构体，那么 for 的性能将显著地高于
	// range，有时候甚至会有上千倍的性能差异。对于这种场景，建议使用 for，如果使
	// 用 range，建议只迭代下标，通过下标访问迭代值，这种使用方式和 for 就没有区
	// 别了。如果想使用 range 同时迭代下标和值，则需要将切片/数组的元素改为指针，
	// 才能不影响性能。

	// 代替上面的for循环
	// for _, i := range ss2 {
	// 	fmt.Println(i)
	// }

	//重复，自己拼接自己count次
	s5 := strings.Repeat("hello", 5)
	fmt.Println(s5)

	//替换
	//helloworld
	s6 := strings.Replace(s1, "l", "*", -1)
	fmt.Println(s6)
	//fmt.Println(strings.Repeat("hello",5))

	s7 := "heLLo WOrlD**123.."
	fmt.Println(strings.ToLower(s7))
	fmt.Println(strings.ToUpper(s7))

	/*
		截取子串：
		substring(start,end)-->substr
		str[start:end]-->substr
			包含start，不包含end下标
	*/
	fmt.Println(s1)
	s8 := s1[:5]
	fmt.Println(s8)
	fmt.Println(s1[5:])
}
