package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

type CurrencyListMap = map[string]map[string]float64

const usdInEur = 0.85
const usdInRub = 74.3

const eurInUsd = 1 / usdInEur
const eurInRub = eurInUsd * usdInRub

const rubInUsd = 1 / usdInRub
const rubInEur = 1 / eurInRub

func main() {
	var currencyList = CurrencyListMap{
		"USD": {
			"EUR": usdInEur,
			"RUB": usdInRub,
		},
		"EUR": {
			"USD": eurInUsd,
			"RUB": eurInRub,
		},
		"RUB": {
			"USD": rubInUsd,
			"EUR": rubInEur,
		},
	}

	inCurrency, outCurrency, sumCurrency := readInput(&currencyList)

	result := convertSum(sumCurrency, inCurrency, outCurrency, &currencyList)

	fmt.Printf("Итого: %.2f %s\n", result, outCurrency)
}

func readInputCurrency(welcome string, availableList []string) string {
	var currency string
	for {
		fmt.Printf("%s %v ", welcome, availableList)
		_, errorInput := fmt.Scan(&currency)

		if errorInput != nil {
			fmt.Println("Ошибка ввода! Повторите еще раз.")
			continue
		}

		currency = strings.ToUpper(currency)
		isHas := slices.Contains(availableList, currency)

		if !isHas {
			fmt.Printf("Ошибка ввода! Список доступных валют: %v\n", availableList)
			continue
		}
		break
	}

	return currency
}

func readInputSum(welcome string) float64 {
	var sum float64

Loop:
	for {
		fmt.Print(welcome)
		_, errorInput := fmt.Scan(&sum)

		if errorInput != nil {
			fmt.Println("Ошибка ввода! Повторите еще раз.")
			continue
		}

		break Loop
	}

	return sum
}

func readInput(currencyList *CurrencyListMap) (string, string, float64) {
	availableList := slices.Collect(maps.Keys(*currencyList))

	inCurrency := readInputCurrency("Введите исходную валюту:", availableList)
	sumCurrency := readInputSum("Введите сумму для конвертации: ")

	inCurrencyIndex := slices.Index(availableList, inCurrency)

	if inCurrencyIndex >= 0 {
		availableList = slices.Delete(availableList, inCurrencyIndex, inCurrencyIndex+1)
	}
	outCurrency := readInputCurrency("Введите целевую валюту:", availableList)

	return inCurrency, outCurrency, sumCurrency
}

func convertSum(sum float64, inCurrency string, outCurrency string, currencyList *CurrencyListMap) float64 {
	if (*currencyList)[inCurrency] == nil {
		panic("Неверная исходная валюте")
	}

	if (*currencyList)[inCurrency][outCurrency] == 0 {
		panic("Неверная целевая валюта")
	}

	return sum * (*currencyList)[inCurrency][outCurrency]
}
