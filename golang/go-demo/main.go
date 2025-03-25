package main

import (
	"errors"
	"fmt"
	"math"
)

const imtPower = 2;

func main() {
	fmt.Println("___ Калькулятор индекса массы тела ___")
	for {
		height, weight := getUserInput()
		IMT, err := calculateIMT(height, weight)
		if err != nil {
			fmt.Println("Некорректные данные")
			continue
			// panic("Нfекорректные данные")
		}
		outputResult(IMT)
		if !checkRepeatCalculation() {
			break
		}
	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Println(result)
	switch {
	case imt < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case imt < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case imt < 25:
		fmt.Println("У вас нормальный вес")
	default:
		fmt.Println("У вас степень ожирения")
	}
}

func calculateIMT(height float64, weight float64) (float64, error) {
	if height <= 0 || weight <= 0 {
		return 0, errors.New("NOT_VALID_INPUT")
	}
	IMT := weight / math.Pow(height/100, imtPower)
	return IMT, nil
}

func getUserInput() (float64, float64) {
	var height float64
	var weight float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&height)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&weight)
	return height, weight
}

func checkRepeatCalculation() bool {
	var repeat string
	fmt.Println("Хотите продолжить? (y/n): ")
	fmt.Scan(&repeat)
	if repeat == "y" || repeat == "Y" {
		return true
	}
	return false
}
