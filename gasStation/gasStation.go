package gasStation

import (
	"errors"
	"fmt"

	"example.com/practice/vehicle"
)

type GasStation struct {
	GasTypes  []int
	GasPrices []float64
}

func (gasStation GasStation) RefuelVehicle(payment *float64, gasType int, vehicle vehicle.Vehicle) {
	fuelPrice, err := gasStation.getFuelPrice(gasType)

	if err != nil {
		panic(err)
	}

	//Maximum gallons is 13
	fmt.Printf("Budget: %.1f\n", *payment)
	fmt.Println("Start refueling")
	fmt.Printf("%d Gasoline, Price: %.2f\n", gasType, fuelPrice)
	var refueledTotal int
	for {
		if vehicle.CurrentGasoline() >= vehicle.MaximumFuelCapacity() {
			fmt.Println("Vehicle already has maximum fuel amount")
			break
		}
		//TODO:  Fix a bug which allows vehicle to fuel more that maxCapacity
		//Todo: Have to have check to do strictly less than or equal  <= MaxCapacity.
		vehicle.Refuel(25)
		refueledTotal++
		*payment -= fuelPrice

		fmt.Printf("-%.2f$, Refueled(gallons): %d\n", fuelPrice, refueledTotal)
	}
	fmt.Println("Vehicle refueled")

}

func (gasStation GasStation) getFuelPrice(gasType int) (float64, error) {
	for i := 0; i < len(gasStation.GasTypes); i++ {
		current := gasStation.GasTypes[i]

		if current == gasType {
			return gasStation.GasPrices[i], nil
		}
	}

	return 0.0, errors.New("Invalid fuel type. Error occured")

}
