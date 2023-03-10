package main

import "fmt"

var ARR_VALUES = []rune("\u0421\u0410\u0428\u0410\u0420\u0412\u041E")

func main() {
	loopOne(0, 4)
	loopTwo(0, 10)
}

func loopOne(start int, itter int) {
	for i := start; i <= itter; i++ {
		fmt.Println("Nilai i =", i)
	}
}

func loopTwo(start int, itter int) {
	for j := start; j <= itter; j++ {
		if j == 5 {
			loopThree(ARR_VALUES)
			continue
		}
		fmt.Println("Nilai j =", j)
	}
}

func loopThree(arrValues []rune) {
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
