package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Getenv("USER")
	course := "Go Fundamentals"

	fmt.Println("\nHi", name, "your current course is", course)
	updateCourse(&course)

	fmt.Println("you're currently watching", course)
}

func updateCourse(course *string) string {
	*course = "Getting started with Docker"
	fmt.Println("Updated course to", *course)
	return *course
}
