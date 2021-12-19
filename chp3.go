package main
import "fmt"

func main(){
  farenheitToCelius(70)
  feetToMeters(100)
}
/**
1. What are two ways to create a new variable?
Either 
x := 2
or 
var x int = 2

2. What is the value of x after running x := 5; x += 1?
6

3. What is scope? How do you determine the scope of a variable in Go?
Scope is determined by the closest curly brackets that the variable is a part of.

4. What is the difference between var and const?
The variables assigned with var can be reassigned.

5. Using the example program as a starting point, write a program that converts
from Fahrenheit into Celsius (C = (F − 32) * 5/9).
**/
func farenheitToCelius( fahrenheit int  ) {
  fmt.Printf("Converting %d fahrenheit to celcius...", fahrenheit)
  celcius := (float64(fahrenheit) - 32) * (5.0/9)
  fmt.Printf("equals %f\n", celcius)
}
/**
6. Write another program that converts from feet into meters (1 ft = 0.3048 m).
**/
func feetToMeters( feet float64 ){
  fmt.Printf("Converting %f feet to meters...", feet)
  meters := feet * 0.3048
  fmt.Printf("equals %f\n", meters)
}
