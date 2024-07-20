package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float32 = 10.1
	var intNum32 int32 = 2
	var result = floatNum + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	var myString string = "Hello, World!"
	fmt.Println(myString)
	fmt.Println(utf8.RuneCountInString(myString))
}
