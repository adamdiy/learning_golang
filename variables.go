package main

import (
	"fmt"
)

func main() {
	
	course := "magic go course"

	fmt.Println(course)

	changeCourse(&course)

	fmt.Println(course)
}

func changeCourse(course *string) string {
	*course = "new course"
	fmt.Println("trying to change your course to ", *course)
	return *course
}
