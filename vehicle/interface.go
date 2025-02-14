package vehicle

type Vehicle interface {
	StartEngine()
	TurnOffEngine()
	Drive()
	Brake()
	CurrentGasoline() int
}

type Refueler interface {
	Refuel(float64)
}
