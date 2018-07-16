package main

import (
	"fmt"
	"math"
	"time"
)

const s string = "constant"

func ExmplSwitch() {
	i := 2
	fmt.Print("Write", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}
}

func ExmplIfElse() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

func ExmplFor() {
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}
}

func ExmplConst() {
	fmt.Println(s)
	const n = 500000000

	const d = 3e20 / n

	fmt.Println(d)
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
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

func ExmplValue() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1=", 1+1)
	fmt.Println("7.0/3.0", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func main() {

	ExmplValue()
	ExmplVar()
	ExmplConst()
	ExmplFor()
	ExmplIfElse()
	ExmplSwitch()
}
