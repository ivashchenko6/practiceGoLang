package gasStation

import (
	"errors"

	"example.com/practice/vehicle"
)

type GasStation struct {
	GasTypes  []int
	GasPrices []float64
}

func (gasStation GasStation) RefuelVehicle(payment float64, gasType int, vehicle vehicle.Vehicle) {
	fuelPrice, err := gasStation.getFuelPrice(gasType)

	if err != nil {
		panic(err)
	}
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
