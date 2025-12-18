package main

import (
	"fmt"
	"math"
)

const IMTPower = 2 

func main() {
	fmt.Println("___ Калькулятор индекса тела ___")
	userHeight, userWeight := userScanGet()
	// getUserInput("Введите свой рост в сантиметрах: ", &userHeight)
	// getUserInput("Введите свой вес в килограммах: ", &userWeight)
	IMT := calculateIMT(userHeight, userWeight)
	if IMT < 16 {
		fmt.Println("Сильный дефицит массы тела")
	}
	outputResult(IMT)
}

func userScanGet() (float64, float64) {
	var userHeight float64
	var userWeight float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес в килограммах: ")
	fmt.Scan(&userWeight)
	return userHeight, userWeight
}

func getUserInput(text string, Data *float64) {
	fmt.Print(text)
	fmt.Scan(Data)
}

func calculateIMT(Height float64, Weight float64) float64 {
	result := Weight / math.Pow(Height/100, IMTPower)
	return result
}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс тела: %.f", IMT)
	fmt.Print(result)
}
