package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	for {
		fmt.Println("___ Калькулятор индекса тела ___")
		userHeight, userWeight, errGetUserInput := getUserInput()
		if errGetUserInput != nil {
			fmt.Println("INPUT ERROR: ", errGetUserInput)
			continue
		}
		IMT, errCalculateIMT := calculateIMT(userHeight, userWeight)
		if errCalculateIMT != nil {
			fmt.Println(errCalculateIMT)
			continue
		}
		outputResult(IMT)
		printIMTStatus(IMT)
		isReapeatCalculation, errIsRepeat := isContinue()
		if errIsRepeat != nil {
			continue
		} else if !isReapeatCalculation {
			break
		}
	}
}

func isContinue() (bool, error) {
	var checkIsContinue string
	fmt.Println("Вычислить еще раз?(y/n)")
	_, errCheck := fmt.Scan(&checkIsContinue)
	if errCheck != nil {
		fmt.Println("INPUT ERROR: ", errCheck)
		return true, errCheck
	}
	if checkIsContinue == "y" || checkIsContinue == "Y" {
		return true, nil
	}
	return false, nil
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

func getUserInput() (float64, float64, error) {
	var userHeight, userWeight float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	_, errHeight := fmt.Scan(&userHeight)
	if errHeight != nil {
		return 0, 0, errHeight
	}
	if userHeight <= 0 {
		return 0, 0, errors.New("ERROR: HEIGHT NOT BE LESS 0 OR EQUAL 0")
	}
	fmt.Print("Введите свой вес в килограммах: ")
	weight, errWeigth := fmt.Scan(&userWeight)
	if errWeigth != nil {
		return 0, 0, errWeigth
	}
	if weight <= 0 {
		return 0, 0, errors.New("ERROR: WEIGHT NOT BE LESS 0 OR EQUAL 0")
	}
	return userHeight, userWeight, nil
}

func calculateIMT(Height float64, Weight float64) (float64, error) {
	if Height <= 0 || Weight <= 0 {
		return 0, errors.New("ОШИБКА ЗНАЧЕНИИ ПАРАМЕТРА ФУНКЦИИ ВЫЧИСЛЕНИИ ИМТ")
	}
	result := Weight / math.Pow(Height/100, IMTPower)
	return result, nil
}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс тела: %.f", IMT)
	fmt.Println(result)
}
