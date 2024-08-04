package main

import (
	"builder-pattern/builders"
	"fmt"
)

func main() {
	fmt.Println("Base Variant Car ...")
	baseVariantBuilder := builders.NewBaseVariant()
	director := builders.NewDirector(baseVariantBuilder)
	baseVariantCar := director.BuildCar()
	fmt.Println(fmt.Sprintf("%+v", baseVariantCar))

	fmt.Println("Mid Variant Car ...")
	midVariantBuilder := builders.NewMidVariant()
	director.SetBuilder(midVariantBuilder)
	midVariantCar := director.BuildCar()
	fmt.Println(fmt.Sprintf("%+v", midVariantCar))

	fmt.Println("Top Variant Car ...")
	topVariantBuilder := builders.NewTopVariant()
	director.SetBuilder(topVariantBuilder)
	topVariantCar := director.BuildCar()
	fmt.Println(fmt.Sprintf("%+v", topVariantCar))
}
