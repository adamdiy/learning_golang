package main

import (
	"errors"
	"fmt"
)

func main() {

	plants := []PowerPlant{
		PowerPlant{hydro, 300, active},
		PowerPlant{wind, 300, active},
		PowerPlant{hydro, 20, active},
		PowerPlant{wind, 300, inactive},
	}

	grid := PowerGrid{300, plants}

	//var gridLoad = 75.

	if option, err := requestOption(); err == nil {

		fmt.Println("")
		switch option {
		case "1":
			grid.generatePR()
		case "2":
			fmt.Println("TBA")
			//			generatePGR(activatePlants, plantCapacities, gridLoad)
		}
	} else {
		fmt.Println(err.Error())
	}
}

func requestOption() (option string, err error) {
	fmt.Println("1) Generate Power Plant Report")

	fmt.Scanln(&option)

	if option != "1" {
		err = errors.New("invalid")
	}
	return
}

type PlantType string

const (
	hydro PlantType = "hydro"
	wind  PlantType = "wind"
)

type PlantStatus string

const (
	active   PlantStatus = "active"
	inactive PlantStatus = "inactive"
)

type PowerPlant struct {
	plantType PlantType
	capacity  float64
	status    PlantStatus
}

type PowerGrid struct {
	load   float64
	plants []PowerPlant
}

func (pg *PowerGrid) generatePR() {
	for idx, p := range pg.plants {
		label := fmt.Sprintf("%s%d", "Plant #", idx)
		fmt.Println(label)
		fmt.Println("----------")
		fmt.Println(p.plantType)
		fmt.Println(p.capacity)
		fmt.Println(p.status)
		fmt.Println("----------")
	}
}
