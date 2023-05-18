package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		strconv包：字符串和基本类型之前的转换
			string convert
	*/
	//fmt.Println("aa"+100)
	//1.bool类型
	s1 := "true"
	b1, err := strconv.ParseBool(s1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%t\n", b1, b1)

	ss1 := strconv.FormatBool(b1)
	fmt.Printf("%T,%s\n", ss1, ss1)

	//2.整数
	//ParseInt(s2, 10, 64),10进制，int64
	s2 := "100"
	i2, err := strconv.ParseInt(s2, 2, 64)
	ii2, err := strconv.ParseInt(s2, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%d\n", i2, i2)
	fmt.Printf("%T,%d\n", ii2, ii2)

	ss2 := strconv.FormatInt(i2, 10)
	fmt.Printf("%T,%s\n", ss2, ss2)

	//itoa(),atoi()
	i3, err := strconv.Atoi("-42") //转为int类型
	fmt.Printf("%T,%d\n", i3, i3)
	//Itoa int转string，不是int转ascii
	sss := 66
	ss3 := strconv.Itoa(sss)
	fmt.Printf("%T,%v\n", sss, sss)
	fmt.Printf("%T,%c\n", ss3, ss3)
}
