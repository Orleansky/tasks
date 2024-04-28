package main

import (
	"fmt"
	"os"
)

func findCommonDivisors(arr []int) []int {
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			min = arr[i]
		}
	}

	var divisorsArr []int

	for i := 2; i <= min; i++ {
		isDivisible := true
		for _, val := range arr {
			if val%i != 0 {
				isDivisible = false
				break
			}
		}
		if isDivisible {
			divisorsArr = append(divisorsArr, i)
		}
	}
	return divisorsArr
}

func main() {
	fmt.Print("Введите длину массива: ")
	var arrLen int
	_, err := fmt.Scanln(&arrLen)
	if err != nil || arrLen <= 0 {
		fmt.Println("Некорректный ввод!")
		os.Exit(1)
	}

	arr := make([]int, arrLen)

	//элементы массива вводить по одному

	fmt.Print("Введите элементы массива: ")
	for i := 0; i < len(arr); i++ {
		_, err := fmt.Scanln(&arr[i])
		if err != nil || arr[i] < 2 {
			fmt.Println("Некорректный ввод!")
			os.Exit(1)
		}
	}

	fmt.Println(findCommonDivisors(arr))
}
