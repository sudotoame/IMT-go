package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	for {
		fmt.Println("___ Калькулятор индекса тела ___")
		userHeight, userWeight := getUserInput()
		IMT := calculateIMT(userHeight, userWeight)
		outputResult(IMT)
		printIMTStatus(IMT)
		isReapeatCalculation := isContinue()
		if !isReapeatCalculation {
			break
		}
	}
}

func isContinue() bool {
	var checkIsContinue string
	fmt.Println("Вычислить еще раз?(y/n)")
	fmt.Scan(&checkIsContinue)
	if checkIsContinue == "y" || checkIsContinue == "Y" {
		return true
	}
	return false
}

func printIMTStatus(IMT float64) {
	switch {
	case IMT < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case IMT < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case IMT < 25:
		fmt.Println("У вас норма массы тела")
	case IMT < 30:
		fmt.Println("У вас избыточная масса тела")
	default:
		fmt.Println("У вас какая-то степень ожирение")
	}
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userWeight float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес в килограммах: ")
	fmt.Scan(&userWeight)
	return userHeight, userWeight
}

func calculateIMT(Height float64, Weight float64) float64 {
	result := Weight / math.Pow(Height/100, IMTPower)
	return result
}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс тела: %.f", IMT)
	fmt.Println(result)
}
