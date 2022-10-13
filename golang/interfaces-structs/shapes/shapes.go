package shapes

import (
	"fmt"
	"math"
)

type Cube struct {
	Length float64
}

type Box struct {
	Length float64
	Width  float64
	Height float64
}

type Sphere struct {
	Radius float64
}

type ofStructure interface {
	Volume() float64
}

func (c Cube) String() string {
	return fmt.Sprintf("Cube has length: %v", c.Length)
}

func (c Sphere) String() string {
	return fmt.Sprintf("Sphere has radius of: %v", c.Radius)
}

func (c Box) String() string {
	return fmt.Sprintf("Box has dimensions length: %v, Width: %v Height: %v", c.Length, c.Width, c.Height)
}

func (c Cube) Volume() float64 {
	return c.Length * c.Length * c.Length
}

func (s Sphere) Volume() float64 {
	return (4 * math.Pi * math.Pow(s.Radius, float64(3))) / 3
}

func (b Box) Volume() float64 {
	return b.Length * b.Width * b.Height
}

func CalculateVolume(kind ofStructure, called string) {
	fmt.Printf("The Volume calculated for our %s is: %f \n", called, kind.Volume())
}

func ReadStructure(shape fmt.Stringer) {
	fmt.Println(shape.String())
}

func main() {

	c := Cube{
		Length: 7,
	}

	b := Box{
		Length: 5.5,
		Width:  5.5,
		Height: 7.7,
	}

	s := Sphere{
		Radius: 7.14,
	}

	CalculateVolume(c, "Cube")
	CalculateVolume(b, "Box")
	CalculateVolume(s, "Sphere")
	ReadStructure(s)
}
