// it is idiomatic to return an error as the last return from functions and methods
// ex:
// func testConn(target string) (rspTime float64, err error)
// nil or zero error code indicates success.
// Error is a type in go.
package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("./test1.txt")

	if err != nil {
		fmt.Println("this is the error code:", err)
	}

	fmt.Println("this is the error code:", err)
}
