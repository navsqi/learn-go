package main

import "fmt"

// Deklarasi interface
type Shape interface {
	Area() float64
}

// Deklarasi struct Circle
type Circle struct {
	Radius float64
}

// Metode untuk menghitung area Circle
// (c Circle) --> seperti membuat method pada class di OOP
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Deklarasi struct Rectangle
type Rectangle struct {
	Width, Height float64
}

// Metode untuk menghitung area Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	// Membuat instance Circle dan Rectangle
	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 3}

	// Memanggil metode Area melalui interface Shape
	shapes := []Shape{c, r}
	for _, shape := range shapes {
		fmt.Println("Area:", shape.Area())
	}
}
