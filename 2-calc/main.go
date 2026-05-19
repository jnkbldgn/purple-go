package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var availableOperations = [3]string{"AVG", "SUM", "MED"}

func main() {

	operation := inputOpertion()
	numbers := inputNumbers()

	result := calcResult(operation, numbers)

	fmt.Println("Результат: ", result)
}

func inputOpertion() string {
	var operation string

	for {
		fmt.Printf("Введите тип вычесления %v: ", availableOperations)
		_, errOperation := fmt.Scan(&operation)

		if errOperation != nil {
			fmt.Println("Ошибка ввода типа операции.")
			continue
		}

		operation := strings.ToUpper(operation)
		isHas := slices.Contains(availableOperations[:], operation)

		if !isHas {
			fmt.Println("Ошибка ввода типа операции. Введите одно из значений: ", availableOperations)
			continue
		}

		break
	}

	return operation
}

func inputNumbers() []float64 {
	arr := make([]float64, 0)

	for {
		fmt.Print("Введите числа через запятую: ")
		var numbersInput string
		_, errorNumbers := fmt.Scan(&numbersInput)

		if errorNumbers != nil {
			fmt.Println("Ошибка ввода чисел.")
			continue
		}

		numbers := strings.Split(numbersInput, ",")
		isError := false

		for _, value := range numbers {
			value := strings.TrimSpace(value)
			newValue, err := strconv.ParseFloat(value, 64)

			if err != nil {
				isError = true
				break
			}

			arr = append(arr, newValue)
		}

		if isError {
			fmt.Println("Ошибка конвертации числа.")
			arr = make([]float64, 0)
			continue
		}

		break
	}

	return arr
}

func calcSum(sli []float64) float64 {
	result := 0.0

	for _, value := range sli {
		result += value
	}

	return result
}

func calcAvg(sli []float64) float64 {
	sum := calcSum(sli)

	return sum / float64(len(sli))
}

func calcMedian(sli []float64) float64 {
	if len(sli) == 1 {
		return sli[0]
	}

	slices.Sort(sli)

	if len(sli)%2 == 0 {
		index := len(sli) / 2

		return (sli[index] + sli[index+1]) / 2
	}

	return sli[len(sli)/2]
}

func calcResult(operation string, numbers []float64) float64 {
	switch strings.ToUpper(operation) {
	case "SUM":
		return calcSum(numbers)
	case "AVG":
		return calcAvg(numbers)
	case "MED":
		return calcMedian(numbers)
	}

	panic("Ошибка. Операция не поддерживается")
}
