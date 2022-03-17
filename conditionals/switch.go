package main

import (
	"fmt"
)

// NO IMPLICIT FALLTHROUGH. every case has an implicit break
func main() {
	switch "Kubernetes" {
	case "kubernetes":
		fmt.Println("Case 1. kubenetes with lower-case \"k\".")
	case "Kubernetes":
		fmt.Println("case 2. Kubernetes with a capital \"K\".")
		fallthrough
	case "K8s":
		fmt.Println("Case 3. Kubernetes short name \"Kates\".")
		fallthrough
	case "Istio":
		fmt.Println("Case 4. Service mesh is the future.")
		fallthrough
	default:
		fmt.Println("Hit the default")
	}
}
