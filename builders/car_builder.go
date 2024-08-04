package builders

import "builder-pattern/models"

type CarBuilder interface {
	buildBody()
	addLights()
	addWheels()
	addInfotainmentSystem()
	getCar() models.Car
}
