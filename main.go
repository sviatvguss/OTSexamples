package main

import (
	"fmt"
	"guss.com/examples/libs/iolib"
	"guss.com/examples/libs/strlib"
	"guss.com/examples/libs/timelib"
	"os"
)

func main() {
	var choice int
	fmt.Print("What step to execute: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Good luck!")
		fmt.Printf("Time is %s\n", timelib.GetTime())
	case 2:
		fmt.Println("Введите набор строк:")
		err := iolib.Uniq(os.Stdin, os.Stdout)
		if err != nil {
			panic(err.Error())
		}
	case 3:
		res := strlib.ReplaceAtoBinC("#", "o", "G# r#cks!")
		fmt.Println(res)
	case 4:
		f, _ := os.Open("input.txt")
		iolib.ReadStringAndPrint(f)
	case 5:
		fmt.Println(timelib.GetRandomVal(1000))
	}

}
