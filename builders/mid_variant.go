package builders

import (
	"builder-pattern/models"
	"fmt"
)

type midVariant struct {
	lightType          string
	wheelsType         string
	infotainmentSystem models.InfotainmentSystem
}

func (m *midVariant) buildBody() {
	fmt.Println("Building Body")
}

func (m *midVariant) addLights() {
	fmt.Println("Adding Lights")
	m.lightType = "Halogen Lamps with LED DRLs"
}

func (m *midVariant) addWheels() {
	fmt.Println("Adding Wheels")
	m.wheelsType = "Styled Steel Wheels"
}

func (m *midVariant) addInfotainmentSystem() {
	fmt.Println("Building Infotainment System")
	infotainmentSystem := models.InfotainmentSystem{
		NoOfSpeakers:   4,
		DisplaySize:    5,
		SpeakerMake:    "Basic",
		IsTouchEnabled: true,
	}
	m.infotainmentSystem = infotainmentSystem
}

func (m *midVariant) getCar() models.Car {
	return models.Car{
		LightType:          m.lightType,
		WheelsType:         m.wheelsType,
		InfotainmentSystem: m.infotainmentSystem,
	}
}

func NewMidVariant() CarBuilder {
	return &midVariant{}
}
