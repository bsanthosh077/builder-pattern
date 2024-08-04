package builders

import (
	"builder-pattern/models"
	"fmt"
)

type topVariant struct {
	lightType          string
	wheelsType         string
	infotainmentSystem models.InfotainmentSystem
}

func (t *topVariant) buildBody() {
	fmt.Println("Building Body")
}

func (t *topVariant) addLights() {
	fmt.Println("Adding Lights")
	t.lightType = "LED Lamps with LED DRLs"
}

func (t *topVariant) addWheels() {
	fmt.Println("Adding Wheels")
	t.wheelsType = "Alloy Wheels"
}

func (t *topVariant) addInfotainmentSystem() {
	fmt.Println("Building Infotainment System")
	infotainmentSystem := models.InfotainmentSystem{
		NoOfSpeakers:   7,
		DisplaySize:    7,
		SpeakerMake:    "Bose",
		IsTouchEnabled: true,
	}
	t.infotainmentSystem = infotainmentSystem
}

func (t *topVariant) getCar() models.Car {
	return models.Car{
		LightType:          t.lightType,
		WheelsType:         t.wheelsType,
		InfotainmentSystem: t.infotainmentSystem,
	}
}

func NewTopVariant() CarBuilder {
	return &topVariant{}
}
