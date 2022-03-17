package main

import (
	"fmt"
)

func main() {
	// courses := make([]string, 5, 10)
	// courses[0] = "Docker & Kubernetes: The Big Picture"
	// courses[1] = "Getting Started with Docker"
	// courses[2] = "Getting Started with Kubernetes"
	// Refactor the above to use one line:

	courses := []string{
		"Docker & Kubernetes: The Big Picture",
		"Getting Started with Docker",
		"Getting Started with Kubernetes",
	}

	// len() and cap() are built in functions that return the length and capacity of slices
	fmt.Printf("Length of slice is %d and capacity is %d\n",
		len(courses), cap(courses))

	for _, i := range courses {
		fmt.Println(i)
	}
}
