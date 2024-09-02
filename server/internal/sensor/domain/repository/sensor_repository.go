package repository

type SensorRepository interface {
	Store() error
	FindMany()
}
