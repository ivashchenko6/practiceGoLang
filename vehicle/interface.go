package vehicle

type Vehicle interface {
	Drive()
	CurrentGasoline() int
	MaximumFuelCapacity() int
	Refuel(int)
}
