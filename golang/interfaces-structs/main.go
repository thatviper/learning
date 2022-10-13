package main

import (
	"fmt"
	shapes "interfaces-structs/shapes"
)


func main() {
	fmt.Println("hello interfaces & structs")
	s := shapes.Sphere{
		Radius: 7.14,
	}
	fmt.Println(s.Volume())
	shapes.ReadStructure(s)
}
