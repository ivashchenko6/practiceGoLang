package car

import "fmt"

type Car struct {
	CarBrand        string
	CarYear         int
	CarModel        string
	MaxFuelCapacity int
	CurrentFuel     int
}

func (car Car) Drive() {
	fmt.Println("Drive")
}
func (car Car) CurrentGasoline() int {
	return car.CurrentFuel
}

func (car Car) MaximumFuelCapacity() int {
	return car.MaxFuelCapacity
}

func (car *Car) Refuel(amount int) {
	car.CurrentFuel += amount
}
