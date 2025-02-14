package main

import (
	"errors"
	"fmt"

	"example.com/practice/car"
)

func main() {
	var mainCar car.Car = car.Car{"Kia", 2024, "Forte GT-line", 400, 305}

	gasType, gasPrice := getCurrentGasInformation() //Возвращает тип бензина и стоимость

}

func getCurrentGasInformation() (int, float64) { //Возвращает выбранный Бензин и Стоимость
	var gasTypes [3]int = [3]int{87, 89, 91}
	var gasPrices [3]float64 = [3]float64{3.99, 4.19, 4.50}

	fmt.Println("Welcome to the Shell Gas Station:")
	fmt.Println("What gas type will you use? 87, 89, 91  or -1 to exit: ")
	var userChoice int
	for {
		fmt.Scan(&userChoice)

		if userChoice == -1 {
			fmt.Println("Goodbye!")
			panic("")
		} else if userChoice != 87 || userChoice == 89 || userChoice == 91 {
			fmt.Println("Only options available 87, 89, 91")
			continue
		}

		index, err := findGasResult(&gasTypes, userChoice)

		if err != nil {
			fmt.Print(err)
			panic(err)
		}
		fmt.Printf("Your gas type is %d, Price: %.2f", gasTypes[index], gasPrices[index])
		return gasTypes[index], gasPrices[index]
	}

}

func findGasResult(gasTypesArray *[3]int, lookingFor int) (int, error) { //Ищет тип бензина в массиве и возвращает его индекс
	for i := 0; i < len(*gasTypesArray); i++ {
		currentItem := gasTypesArray[i]

		if currentItem == lookingFor {
			return i, nil
		}
	}

	return 0, errors.New("Invalid Index")
}
