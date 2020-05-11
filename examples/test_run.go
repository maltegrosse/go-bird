package main

import (
	"fmt"
	"github.com/maltegrosse/go-bird"
)

func main() {

	b, err := bird.NewBird(50.11162202402973, 0.9965422973539708, 820, 0.3, 1.5, 0.08, 0.85, 0.2, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("direct normal solar irradiance: ", b.GetDirectNormal(), "[W/m^2]")
	fmt.Println("global horizontal solar irradiance: ", b.GetGlobalHoriz(), "[W/m^2]")
	fmt.Println("diffuse horizontal solar irradiance: ", b.GetDiffuseHoriz(), "[W/m^2]")

	/*
		OUTPUT.
		direct normal solar irradiance:  874.5065757929946 [W/m^2]
		global horizontal solar irradiance:  664.8749080114459 [W/m^2]
		diffuse horizontal solar irradiance:  104.05908396215341 [W/m^2]

	*/

}
