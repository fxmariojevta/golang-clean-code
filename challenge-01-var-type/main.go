package main

import "fmt"

func main() {
	var i int = 21
	fmt.Printf("menampilakan nilai i => %d\n", i)
	fmt.Printf("menampilkan tipe data dari variabel i => %T\n", i)
	fmt.Printf("menampilkan tanda %% => %%\n")

	var j bool = false
	fmt.Printf("menampilkan nilai boolean j => %t\n", j)
	fmt.Printf("menampilkan nilai boolean j => %t\n", !j)

	fmt.Printf("menampilkan unicode russia : Я (ya) => %U\n", 'Я')

	fmt.Printf("menampilkan nilai base 10 : 21 => %d\n", 21)
	fmt.Printf("menampilkan nilai base 8 :25 => %o\n", 25)
	fmt.Printf("menampilkan nilai base 16 : f => %x\n", 'f')
	fmt.Printf("menampilkan nilai base 16 : F => %X\n", 'F')
	fmt.Printf("menampilkan unicode karakter Я : U+042F => %c\n", '\u042F')

	var k float64 = 123.456
	fmt.Printf("menampilkan float : 123.456000 => %.6f\n", k)
	fmt.Printf("menampilkan float scientific : 1.234560E+02 => %E\n", k)
}
