package main

import (
	"fmt"
	"guss.com/examples/libs/iolib"
	"guss.com/examples/libs/timelib"
	"os"
)

func main() {
	fmt.Println("Good luck!")
	fmt.Printf("Time is %s\n", timelib.GetTime())

	fmt.Println("Введите набор строк:")
	err := iolib.Uniq(os.Stdin, os.Stdout)
	if err != nil {
		panic(err.Error())
	}
}
