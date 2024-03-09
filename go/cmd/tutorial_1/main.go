package main

import (
	"fmt"
	"unicode/utf8"
)

// data types
func main() {
	fmt.Println("Hello, world!")

	var intNum int = 323767
	fmt.Println(intNum)

	var floatNum float32 = 3.1415926
	fmt.Println(floatNum)

	var result float32 = floatNum + float32(intNum)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(float32(intNum1 / intNum2))

	var myString string = "Hello" + " " + "World"
	fmt.Println(myString)

	fmt.Println(utf8.RuneCountInString(myString))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var mybool bool = true
	fmt.Println(mybool)

	myVar := "text"
	fmt.Println(myVar)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

}
