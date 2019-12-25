package structs

import "math"

// interface declaration
// interface resolution is implicit.

// If the type you pass in matches what the interface is asking for,
// it will compile
type Shape interface {
	Area() float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Height + r.Width)
}

//  A struct is just a named collection of fields
type Rectangle struct {
	Width  float64
	Height float64
}

//  No overloading in Go

// A method is a function with a receiver
// Method declaration: binds identifier to method
//	& associates with receiver base type
// func (receiverName RecieverType) MethodName(args)
// convention: have the receiver var. = first letter of the type
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Width  float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Height * t.Width / 2
}
