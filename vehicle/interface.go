package vehicle

type Vehicle interface {
	StartEngine()
	TurnOffEngine()
	Drive()
	Brake()
	CurrentGasoline() int
	//Refuel()
}

type Refueler interface {
	Refuel(float64)
}
