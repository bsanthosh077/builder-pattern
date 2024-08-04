package builders

import "builder-pattern/models"

type Director struct {
	carBuilder CarBuilder
}

func NewDirector(carBuilder CarBuilder) *Director {
	return &Director{carBuilder}
}

func (d *Director) BuildCar() models.Car {
	d.carBuilder.buildBody()
	d.carBuilder.addLights()
	d.carBuilder.addWheels()
	d.carBuilder.addInfotainmentSystem()
	return d.carBuilder.getCar()
}

func (d *Director) SetBuilder(builder CarBuilder) {
	d.carBuilder = builder
}
