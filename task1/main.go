package main

import (
	"fmt"
	"os"
	"strconv"
)

func computerBuilder(amount int) string {
	var computers string

	if amount < 0 {
		return "Введите число больше 0!"
	}

	if (amount%100 >= 11 && amount%100 <= 14) || amount%10 == 0 || (amount%10 >= 5 && amount%10 <= 9) {
		computers = strconv.Itoa(amount) + " компьютеров"
	} else if amount%10 >= 2 && amount%10 <= 4 {
		computers = strconv.Itoa(amount) + " компьютера"
	} else {
		computers = strconv.Itoa(amount) + " компьютер"
	}
	return computers
}

func main() {
	fmt.Print("Введите количество компьютеров: ")
	var amount int
	_, err := fmt.Scanln(&amount)
	if err != nil {
		fmt.Println("Некорректный ввод!")
		os.Exit(1)
	}

	fmt.Println(computerBuilder(amount))
}
