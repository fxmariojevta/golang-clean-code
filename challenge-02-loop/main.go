package main

import "fmt"

var ARR_VALUES = []rune("\u0421\u0410\u0428\u0410\u0420\u0412\u041E")

func main() {
	loopOne(0, "i", 4)
	loopOne(0, "j", 10)
}

func loopOne(start int, variable string, itter int) {
	for i := start; i <= itter; i++ {
		if i == 5 {
			loopTwo(ARR_VALUES)
			continue
		}
		fmt.Printf("Value %s = %d\n", variable, i)
	}
}

func loopTwo(arrValues []rune) {
	for index, arrValue := range arrValues {
		if index > 0 {
			index *= 2
		}
		fmt.Printf("Character %U '%c' starts at byte position %d\n",
			arrValue,
			arrValue,
			index,
		)
	}
}
