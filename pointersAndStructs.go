package main

import (
	"fmt"
)

type AnimalType string

const (
	dog     AnimalType = "dog"
	giraffe AnimalType = "giraffe"
)

type ZooSection struct {
	animalType AnimalType
	population float64
}

type Zoo struct {
	sections []ZooSection
}

func (z *ZooSection) Jailbreak() {
	z.population = 0
}

func (z ZooSection) Headcount() {
	fmt.Print(z.animalType, ":", z.population)

}

func main() {
	adamsZoo := []ZooSection{
		{dog, 300},
		{giraffe, 3},
	}

	adamsZoo[1].Jailbreak()
	fmt.Println(adamsZoo)
}
