package structsmethinter

import "math"

// Create a simple type using a struct
type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

// Two choices:
// Have functions with the samee name declared in different packages.
// We can define methods on our newly defined types instead

func Perimeter(rect Rectangle) float64 {
	return (2 * rect.Width) + (2 * rect.Height)
}

func Area(rect Rectangle) float64 {
	return rect.Height * rect.Width
}

// A method is function w/ a receiver. eg. t.Errorf -> method
// Methods are similar to functions but they are called by invoking them
// on an instance of a particular type.

// tldr: you can only call methods on "things", whereas functions can be called anywhere

// Syntax for declaring methods, -> func(receiverName ReceiverType) MethodName(args)
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// Next, wee want to be able to write some kind of checkArea function
// that we can pass both Rectangle s and Circle s to, but fail to compile if
// we try to pass something that isn't a shape

// Interfaces are a powerful concept in statically typed languages like Go
// because they allow you to make functions that can be used w/ different types
// and create highly-decoupled code

// How does something become a shapee? We just tell Go what a Shape is using
// an interface declaration

type Shape interface {
	Area() float64
}
