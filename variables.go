package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Environ())

	name := os.Getenv("AIRE_HOME")
	course := "magic go course"

	fmt.Println(name, course)

	changeCourse(&course)

	fmt.Println(name, course)
}

func changeCourse(course *string) string {
	*course = "new look"
	fmt.Println("trying to change your couse to ", *course)
	return *course
}
