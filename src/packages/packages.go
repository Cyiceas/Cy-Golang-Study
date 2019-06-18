package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf("My favorite number is %d.", rand.Intn(10))
	fmt.Println()
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("My favorite number is %d.", rand.Intn(10))
}
