package main

import "fmt"

func ExmplValue() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1=", 1+1)
	fmt.Println("7.0/3.0", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func ExmplVar() {
	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)
}

func main() {
	ExmplValue()
	ExmplVar()
}
