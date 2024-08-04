package builders

import (
	"builder-pattern/models"
	"fmt"
)

type baseVariant struct {
	lightType          string
	wheelsType         string
	infotainmentSystem models.InfotainmentSystem
}

func (b *baseVariant) buildBody() {
	fmt.Println("Building Body")
}

func (b *baseVariant) addLights() {
	fmt.Println("Adding Lights")
	b.lightType = "Halogen Lamps"
}

func (b *baseVariant) addWheels() {
	fmt.Println("Adding Wheels")
	b.wheelsType = "Steel Wheels"
}

func (b *baseVariant) addInfotainmentSystem() {
	fmt.Println("Building Infotainment System")
	infotainmentSystem := models.InfotainmentSystem{
		NoOfSpeakers: 2,
		DisplaySize:  5,
		SpeakerMake:  "Basic",
	}
	b.infotainmentSystem = infotainmentSystem
}

func (b *baseVariant) getCar() models.Car {
	fmt.Println("Getting Car")
	fmt.Println(b.lightType)
	fmt.Println(b.wheelsType)
	return models.Car{
		LightType:          b.lightType,
		WheelsType:         b.wheelsType,
		InfotainmentSystem: b.infotainmentSystem,
	}
}

func NewBaseVariant() CarBuilder {
	return &baseVariant{}
}
