package main

import (
	"fmt"
	"time"
)

func main() {
	for timer := 10; timer >= 0; timer-- {
		if timer%2 == 0 {
			fmt.Println("Boom!")
			// so we dont print the 0:
			// breaks out of the current loop
			break
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}
}
