package main

import (
	"fmt"
	"example.com/practice/vehicle"
	"example.com/practice/car"
)




func main() {
	var mainCar car.Car = car.Car{"Kia", 2024, "Forte GT-line", 400, 305}

	gasType, gasPrice := getCurrentGasInformation() //Menu





}

func getCurrentGasInformation() (int, float64){
	var gasTypes [3]int = [3]int{87, 89, 91}
	var gasPrices [3]float64 = [3]float64{3.99, 4.19, 4.50}

	fmt.Println("Welcome to the Shell Gas Station:")
	fmt.Println("What gas type will you use? 87, 89, 91")
	var userChoice int;
	for {
		fmt.Scan(&userChoice)
		index, err := findGasResult(&gasTypes)



		switch userChoice {
		case 87:

			fmt.Printf("Your gas type is %d, Price: %.2f", gasTypes[])
		case 89:

		case 91:

		default:
			fmt.Println("Only options available 87, 89, 91")
			continue
		}
	}

}


func findGasResult(gasTypesArray *[]int, lookingFor int)  (int, error){
	for i:= 0; i < len(*gasTypesArray); i++ {
		currentItem := gasTypesArray[i]

		if (currentItem == lookingFor) {
			return 0, nil
		}
	}


	return 0, errors.New("Invalid Index")
}

