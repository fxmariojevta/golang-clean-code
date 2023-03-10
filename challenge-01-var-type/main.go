package main

import "fmt"

func main() {
	var i int = 21
	fmt.Printf("print value of variable i => %d\n", i)
	fmt.Printf("print data type of variable i => %T\n", i)
	fmt.Printf("print symbol of %% => %%\n")

	var j bool = false
	fmt.Printf("print value of boolean j => %t\n", j)
	fmt.Printf("print value of boolean j => %t\n", !j)

	fmt.Printf("print value of russian unicode : Я (ya) => %U\n", 'Я')

	fmt.Printf("print value of base 10 : 21 => %d\n", 21)
	fmt.Printf("print value of base 8 :25 => %o\n", 25)
	fmt.Printf("print value of base 16 : f => %x\n", 'f')
	fmt.Printf("print value of base 16 : F => %X\n", 'F')
	fmt.Printf("print value of russian character unicode Я : U+042F => %c\n", '\u042F')

	var k float64 = 123.456
	fmt.Printf("print value of float : 123.456000 => %.6f\n", k)
	fmt.Printf("print value of float scientific : 1.234560E+02 => %E\n", k)
}
