package main

import (
	"errors"
	"fmt"
	"math"
)

const imtPower = 2;

func main() {
	fmt.Println("___ BMI Calculator ___")
	for {
		height, weight := getUserInput()
		IMT, err := calculateIMT(height, weight)
		if err != nil {
			fmt.Println("Invalid input")
			continue
			// panic("Invalid input")
		}
		outputResult(IMT)
		if !checkRepeatCalculation() {
			break
		}
	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Your BMI: %.0f", imt)
	fmt.Println(result)
	switch {
	case imt < 16:
        fmt.Println("You have a severe weight deficiency")
	case imt < 18.5:
        fmt.Println("You have a weight deficiency")
	case imt < 25:
        fmt.Println("You have a normal weight")
	default:
        fmt.Println("You have a degree of obesity")
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
	fmt.Print("Enter your height in centimeters: ")
	fmt.Scan(&height)
    fmt.Print("Enter your weight: ")
	fmt.Scan(&weight)
	return height, weight
}

func checkRepeatCalculation() bool {
	var repeat string
	fmt.Println("Do you want to calculate again? (y/n): ")
	fmt.Scan(&repeat)
	if repeat == "y" || repeat == "Y" {
		return true
	}
	return false
}
