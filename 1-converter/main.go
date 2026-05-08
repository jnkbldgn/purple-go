package main

import "fmt"

func main() {
	var USD_EUR float64
	var USD_RUB float64

	readInput("Введите курс USD к EUR: ", &USD_EUR)
	readInput("Введите курс USD к RUB: ", &USD_RUB)

	var EUR_RUB = (1 / USD_EUR) * USD_RUB

	printValue(EUR_RUB)
}

func readInput(preview string, numeric *float64) {
	fmt.Print(preview)
	_, err := fmt.Scan(numeric)

	if err != nil {
		fmt.Println("Ошибка ввода")
	}
}

func printValue(value float64) {
	fmt.Printf("Курс EUR к RUB: %.2f\n", value)
}
