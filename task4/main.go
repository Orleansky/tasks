package main

import (
	"fmt"
	"os"
)

func printMultTable(max int) {
	for i := 1; i <= max; i++ {
		fmt.Print("\t", i)
	}

	for i := 1; i <= max; i++ {
		fmt.Print("\n", i)
		for j := 1; j <= max; j++ {
			fmt.Print("\t", j*i)
		}
	}
}

func main() {
	fmt.Print("Введите максимальное число для таблицы: ")
	var max int
	_, err := fmt.Scanln(&max)
	if err != nil || max <= 0 {
		fmt.Println("Некорректный ввод!")
		os.Exit(1)
	}

	printMultTable(max)
}
