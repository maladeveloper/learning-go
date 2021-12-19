package main
import (
  "fmt"
  "math"
)

func main(){
  fmt.Println("Nothing to print.")
  r := Rectangle{5, 0, 10, 10}
  c := Circle{0, 0, 5}
  fmt.Println("The perimeter of the rectangle is:", r.perimeter())
  fmt.Println("The perimeter of the circle is:", c.perimeter())
}

/**
1. What’s the difference between a method and a function?
A method is a function that belongs to a struct. In Go terms, 
a method has a receiver whilst a function does not.

2. Why would you use an embedded anonymous field instead of a normal named
field?
Would be used for 'is a' relationships such that the methods of the embedded type
may be called directly from the struct containing it.

3. Add a new perimeter method to the Shape interface to calculate the perimeter of
a shape. Implement the method for Circle and Rectangle.
**/
type Circle struct {
  x, y, r float64
}

type Rectangle struct {
  x1, y1, x2, y2 float64
}

type Shape interface {
  perimeter() float64
}

func (r *Rectangle) perimeter() float64{
  return math.Abs(r.x1 - r.x2) * 2 + (math.Abs(r.y1-r.y2)) * 2
}

func (c *Circle) perimeter() float64{
  return 2 * math.Pi * c.r
}

