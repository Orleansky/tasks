package main

import (
	"fmt"
	"os"
)

func findPrimeNums(min, max int) []int {
	arr := []int{}

	for i := min; i <= max; i++ {
		isPrime := true

		if i <= 1 {
			isPrime = false
		}
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
			}
		}
		if isPrime {
			arr = append(arr, i)
		}
	}
	return arr
}

func main() {
	var min, max int
	fmt.Print("Введите минимальную границу диапазона: ")
	_, err := fmt.Scanln(&min)
	if err != nil {
		fmt.Println("Некорректный ввод!")
		os.Exit(1)
	}

	fmt.Print("Введите максимальную границу диапазона: ")
	_, err = fmt.Scanln(&max)
	if err != nil {
		fmt.Println("Некорректный ввод!")
		os.Exit(1)
	}

	fmt.Println(findPrimeNums(min, max))
}
